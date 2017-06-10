package gvm

import (
	"strings"
	"reflect"
	"errors"
)

const DEFAULT_VM_STACK_SIZE  = 512

type Thread struct {
	id          uint
	name        string
	vmStack     []*StackFrame
}

func newThread(name string) *Thread  {
	return &Thread{
		name: name,
		vmStack: make([]*StackFrame, 0, DEFAULT_VM_STACK_SIZE)}
}


type StackFrame struct {
	method *Method
	// if this frame is current frame, the pc is for the pc of this thread;
	// otherwise, it is a snapshot one since the last time
	pc uint32
	localVariables      []t_any
	// operand stack
	operandStack        []t_any
}

func NewStackFrame(method *Method) *StackFrame {
	stackFrame := &StackFrame{
		method: method,
		pc: 0,
		localVariables: make([]t_any, method.maxLocals), // local variables have no initial values
		operandStack: make([]t_any, 0, method.maxStack)}
	return stackFrame
}

func (this *StackFrame) loadVar(index uint) t_any {
	value := this.localVariables[index]
	return value
}

func (this *StackFrame) storeVar(index uint, value t_any)  {
	this.localVariables[index] = value
}

func (this *Thread) invokeStaticMethod(index uint16) {
	f := this.peekFrame()
	m := f.method
	c := m.class

	method := c.constantPool[index].resolve().(*RuntimeConstantMethodrefInfo).method
	parameterCount := len(method.parameterDescriptors)
	params := make([]t_any, parameterCount+1)
	// read parameters
	for i := parameterCount-1; i >= 0; i-- {
		params[i] = f.pop()
	}

	if method.isNative() {
		this.invokeNativeMethod(method, params...)
	} else {
		frame := NewStackFrame(method)
		// pass parameters
		for j := parameterCount-1; j >= 0; j--  {
			frame.storeVar(uint(j), params[j])
		}

		this.pushFrame(frame)
	}

}

func (this *Thread) invokeVitrualMethod(index uint16) {
	f := this.peekFrame()
	m := f.method
	c := m.class
	method := c.constantPool[index].resolve().(*RuntimeConstantMethodrefInfo).method
	if method.isStatic() {
		fatal("Not an instance method")
	}
	parameterCount := len(method.parameterDescriptors) + 1 // with an extra objectref: this
	params := make([]t_any, parameterCount)
	for i := parameterCount-1; i >= 0; i-- {
		params[i] = f.pop()
	}
	// get objectref and target method
	objectref := params[0].(*j_object)
	overridenMethod := objectref.class.findMethod(method.name + method.descriptor)


	if method.isNative() {
		this.invokeNativeMethod(overridenMethod, params...)
	} else {
		frame := NewStackFrame(overridenMethod)
		// pass parameters
		for j := 0; j < parameterCount; j++ {
			frame.storeVar(uint(j), params[j])
		}

		this.pushFrame(frame)
	}
}

// like invokevirtual with objectref, but don't find along the inheritance
func (this *Thread) invokeSpecialMethod(index uint16)  {
	f := this.peekFrame()
	m := f.method
	c := m.class
	method := c.constantPool[index].resolve().(*RuntimeConstantMethodrefInfo).method
	parameterCount := len(method.parameterDescriptors) + 1 // with an extra objectref: this
	params := make([]t_any, parameterCount)
	for i := parameterCount-1; i >= 0; i-- {
		params[i] = f.pop()
	}

	if method.isNative() {
		this.invokeNativeMethod(method, params...)
	} else {
		frame := NewStackFrame(method)
		// pass parameters
		for j := 0; j < parameterCount; j++ {
			frame.storeVar(uint(j), params[j])
		}

		this.pushFrame(frame)
	}
}

func (this *Thread) invokeNativeMethod(method *Method, params ... t_any) (result []reflect.Value, err interface{}) {
	if !method.isNative() {
		fatal("Not a native method")
	}
	debug("🍺invoke native method %s#%s%s 😎\n", method.class.thisClassName, method.name, method.descriptor)
	name := "Java_" + strings.Replace(method.class.thisClassName + "_" + method.name, "/", "_", -1)
	funcs := NativeFunctions
	if _, ok := funcs[name]; !ok {
		err = errors.New(name + " does not exist.")
		fatal( "%s does not exist.", name)
		return
	}
	if len(params) != funcs[name].Type().NumIn() {
		err = errors.New("The number of params is not adapted.")
		fatal( "The number of params is not adapted.")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = funcs[name].Call(in)
	return
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

func (this *StackFrame) getField(objectref *j_object, index uint16) t_any {
	i := this.method.class.constantPool[index].resolve().(*RuntimeConstantFieldrefInfo).field.index
	return objectref.fields[i]
}

func (this *StackFrame) putField(objectref *j_object, index uint16, value t_any) {
	i := this.method.class.constantPool[index].resolve().(*RuntimeConstantFieldrefInfo).field.index
	objectref.fields[i] = value
}

func (this *StackFrame) push(jany t_any)  {
	operandStackSize := len(this.operandStack)
	this.operandStack = this.operandStack[:operandStackSize+1]
	this.operandStack[operandStackSize] = jany
}

func (this *StackFrame) pop() t_any {
	operandStackSize := len(this.operandStack)
	jany := this.operandStack[operandStackSize-1]
	this.operandStack[operandStackSize-1] = nil
	this.operandStack = this.operandStack[:operandStackSize-1]
	return jany
}

func (this *StackFrame) peek() t_any {
	operandStackSize := len(this.operandStack)
	jany := this.operandStack[operandStackSize-1]
	return jany
}

func (this *Thread) pushFrame(stackFrame *StackFrame)  {
	size := len(this.vmStack)
	if size == DEFAULT_VM_STACK_SIZE {
		fatal("Stack Overflow")
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