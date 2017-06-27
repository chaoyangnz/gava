package jago

import (
	"strings"
	"reflect"
)

type ThreadManager struct {
	currentThread *Thread
}

func (this *ThreadManager) NewThread(name string) *Thread {
	return &Thread{
			name: name,
			vmStack: make([]*Frame, 0, DEFAULT_VM_STACK_SIZE)}
}

// We choose fix-sized stack size
const DEFAULT_VM_STACK_SIZE  = 512

type Thread struct {
	id          uint
	name        string
	vmStack     []*Frame
}

func (this *Thread) Run()  {
	for len(this.vmStack) != 0 { // per stack frame
		f := this.peekFrame()
		bytecode := f.method.code
		Trace("🍷 %s.%s%s", f.method.class.name, f.method.name, f.method.descriptor)
		for f.pc < len(f.method.code) {
			pc := f.pc
			opcode := bytecode[pc]
			instruction := instructions[opcode]
			Trace("   %04d ➢ %s", int(pc), instruction.mnemonic)
			jumped := false
			instruction.interpret(opcode, this, f, f.method.class, f.method, &jumped)
			// jump instruction can operate pc
			// some instruction also have variable length: tableswitch...
			// these instructions will control pc themselves
			instruction_length := JVM_OPCODE_LENGTH_INITIALIZER[opcode]
			if !jumped {
				f.pc += instruction_length
			}

			// if instruction operates the stack, we follow it
			if len(this.vmStack) == 0 || f != this.peekFrame() {
				break
			}
		}
	}
}

type Frame struct {
	method *Method
	// if this frame is current frame, the pc is for the pc of this thread;
	// otherwise, it is a snapshot one since the last time
	pc int
	// long and double will occupy two variable indexes
	localVariables      []Value
	// operand stack
	// a value of type `long` or `double` contributes two units to the depth and a value of any other type contributes one unit
	// but here we use long and double only use one unit. There is not any violation
	operandStack        []Value
}

func NewStackFrame(method *Method) *Frame {
	stackFrame := &Frame{
		method: method,
		pc: 0,
		localVariables: make([]Value, method.maxLocals), // local variables have no initial values
		operandStack: make([]Value, 0, method.maxStack)}
	return stackFrame
}

func (this *Frame) loadVar(index uint) Value {
	value := this.localVariables[index]
	return value
}

func (this *Frame) storeVar(index uint, value Value)  {
	this.localVariables[index] = value
}

func (this *Frame) index8() uint8 {
	return uint8(this.method.code[this.pc+1])
}

func (this *Frame) index16() uint16 {
	return (uint16(this.method.code[this.pc+1]) << 8) | uint16(this.method.code[this.pc+2])
}

func (this *Frame) offset16() int16 {
	return int16((uint16(this.method.code[this.pc+1]) << 8) | uint16(this.method.code[this.pc+2]))
}

func (this *Thread) invokeNativeMethod(method *Method, params ... Value) Value {
	if !method.isNative() {
		Fatal("Not a native method")
	}
	Debug("🍺invoke native method %s#%s%s 😎\n", method.class.name, method.name, method.descriptor)
	name := "Java_" + strings.Replace(method.class.name + "_" + method.name, "/", "_", -1)
	funcs := NATIVE_FUNCTIONS
	if _, ok := funcs[name]; !ok {
		Fatal( "%s does not exist.", name)
	}
	if len(params) != funcs[name].Type().NumIn() {
		Fatal( "The number of params is not adapted.")
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result := funcs[name].Call(in)
	if len(result) == 0 {
		return Value(nil)
	}
	return result[0].Interface().(Value)
}

/*
Parameters are passed in a reversed order from operand stack in JVM
 */
func (this *Frame) passParameters(callee *Frame)  {
	method := callee.method
	start := len(method.parameterDescriptors) - 1
	end := 0
	if !method.isStatic() {
		start += 1
		end += 1
	}
	for i := start ;i >= end; i-- {
		callee.storeVar(uint(i), this.pop())
	}
	if !method.isStatic() {
		callee.storeVar(0, this.pop()) // this reference for instance method
	}
}

func (this *Frame) passReturn(caller *Frame)  {
	caller.push(this.pop())
}

func (this *Frame) getField(objectref ObjectRef, index uint16) Value {
	i := this.method.class.constantPool[index].(*FieldRef).ResolvedField().index
	return objectref.instanceVars[i]
}

func (this *Frame) putField(objectref ObjectRef, index uint16, value Value) {
	i := this.method.class.constantPool[index].(*FieldRef).ResolvedField().index
	objectref.instanceVars[i] = value
}

func (this *Frame) push(Value Value)  {
	operandStackSize := len(this.operandStack)
	this.operandStack = this.operandStack[:operandStackSize+1]
	this.operandStack[operandStackSize] = Value
}

func (this *Frame) pop() Value {
	operandStackSize := len(this.operandStack)
	Value := this.operandStack[operandStackSize-1]
	this.operandStack[operandStackSize-1] = nil
	this.operandStack = this.operandStack[:operandStackSize-1]
	return Value
}

func (this *Frame) peek() Value {
	operandStackSize := len(this.operandStack)
	Value := this.operandStack[operandStackSize-1]
	return Value
}

func (this *Thread) pushFrame(stackFrame *Frame)  {
	size := len(this.vmStack)
	if size == DEFAULT_VM_STACK_SIZE {
		Throw("java.lang.StackOverflowError", "Exceed the maximum stack size")
	}
	this.vmStack = this.vmStack[:size+1]
	this.vmStack[size] = stackFrame
}

func (this *Thread) popFrame()  {
	size := len(this.vmStack)
	if size == 0 {
		return
	}
	this.vmStack = this.vmStack[:size-1]
}

func (this *Thread) peekFrame() *Frame {
	size := len(this.vmStack)
	return this.vmStack[size-1]
}

func (this *Thread) pushFrames(stackFrames ...*Frame)  {
	for _, stackFrame := range stackFrames {
		this.pushFrame(stackFrame)
	}
}

/**
	Always add to tail: this can be used when system initialization
 */
func (this *Thread) enqueueFrame(stackFrame *Frame)  {
	size := len(this.vmStack)
	if size == DEFAULT_VM_STACK_SIZE {
		Fatal("Stack Overflow")
	}
	this.vmStack = this.vmStack[:size+1]
	for i := size; i >= 1; i-- {
		this.vmStack[i] = this.vmStack[i-1]
	}
	this.vmStack[0] = stackFrame
}

func (this *Thread) enqueueFrames(stackFrames ...*Frame)  {
	for _, stackFrame := range stackFrames {
		this.enqueueFrame(stackFrame)
	}
}