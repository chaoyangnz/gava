package jago

import (
	"strings"
	"reflect"
)

type ThreadManager struct {
	currentThread *Thread
}

func (this *ThreadManager) CurrentThread() *Thread {
	if this.currentThread == nil {
		this.currentThread = &Thread{
			name: "main",
			vmStack: make([]*StackFrame, 0, DEFAULT_VM_STACK_SIZE)}
	}
	return this.currentThread
}

const DEFAULT_VM_STACK_SIZE  = 512

type Thread struct {
	id          uint
	name        string
	vmStack     []*StackFrame
}

type StackFrame struct {
	method *Method
	// if this frame is current frame, the pc is for the pc of this thread;
	// otherwise, it is a snapshot one since the last time
	pc uint32
	localVariables      []Value
	// operand stack
	operandStack        []Value
}

func NewStackFrame(method *Method) *StackFrame {
	stackFrame := &StackFrame{
		method: method,
		pc: 0,
		localVariables: make([]Value, method.maxLocals), // local variables have no initial values
		operandStack: make([]Value, 0, method.maxStack)}
	return stackFrame
}

func (this *StackFrame) loadVar(index uint) Value {
	value := this.localVariables[index]
	return value
}

func (this *StackFrame) storeVar(index uint, value Value)  {
	this.localVariables[index] = value
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
func (this *StackFrame) passParameters(callee *StackFrame)  {
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

func (this *StackFrame) passReturn(caller *StackFrame)  {
	caller.push(this.pop())
}

func (this *StackFrame) getField(objectref *Object, index uint16) Value {
	i := this.method.class.constantPool[index].(*FieldRef).ResolvedField().index
	return objectref.instanceVars[i]
}

func (this *StackFrame) putField(objectref *Object, index uint16, value Value) {
	i := this.method.class.constantPool[index].(*FieldRef).ResolvedField().index
	objectref.instanceVars[i] = value
}

func (this *StackFrame) push(Value Value)  {
	operandStackSize := len(this.operandStack)
	this.operandStack = this.operandStack[:operandStackSize+1]
	this.operandStack[operandStackSize] = Value
}

func (this *StackFrame) pop() Value {
	operandStackSize := len(this.operandStack)
	Value := this.operandStack[operandStackSize-1]
	this.operandStack[operandStackSize-1] = nil
	this.operandStack = this.operandStack[:operandStackSize-1]
	return Value
}

func (this *StackFrame) peek() Value {
	operandStackSize := len(this.operandStack)
	Value := this.operandStack[operandStackSize-1]
	return Value
}

func (this *Thread) pushFrame(stackFrame *StackFrame)  {
	size := len(this.vmStack)
	if size == DEFAULT_VM_STACK_SIZE {
		Fatal("Stack Overflow")
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

func (this *Thread) peekFrame() *StackFrame  {
	size := len(this.vmStack)
	return this.vmStack[size-1]
}
