package gvm

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
	method *JavaMethod
	// if this frame is current frame, the pc is for the pc of this thread;
	// otherwise, it is a snapshot one since the last time
	pc uint32
	localVariables      []t_any
	// operand stack
	operandStack        []t_any
}

func NewStackFrame(method *JavaMethod) *StackFrame {
	stackFrame := &StackFrame{
		method: method,
		pc: 0,
		localVariables: make([]t_any, method.maxLocals), // local variables have no initial values
		operandStack: make([]t_any, 0, method.maxStack)}
	return stackFrame
}

func (this *StackFrame) loadVar(index uint) t_any {
	value := this.localVariables[index]
	return checkType(value)
}

func (this *StackFrame) storeVar(index uint, value t_any)  {
	checkType(value)
	this.localVariables[index] = value
}

func (this *Thread) invokeStaticMethod(index uint16) {
	f := this.peekFrame()
	m := f.method
	c := m.class
	method := c.constantPool[index].resolve().(*RuntimeConstantMethodrefInfo).method
	parameterCount := len(method.parameterDescriptors)
	if method.isNative() {
		debug("😎 invoke native method %s#%s%s \n", method.class.thisClassName, method.name, method.descriptor)
		GVM_print(f.pop().(t_object/*java_lang_string*/))
		return
	}
	frame := NewStackFrame(method)
	// pass parameters

	for i := parameterCount-1; i >= 0; i--  {
		argument := checkType(f.pop())
		frame.storeVar(uint(i), argument)
	}

	this.pushFrame(frame)
}

func (this *Thread) invokeVitrualMethod(index uint16) {
	f := this.peekFrame()
	m := f.method
	c := m.class
	method := c.constantPool[index].resolve().(*RuntimeConstantMethodrefInfo).method
	if method.isStatic() {
		fatal("Not an instance method")
	}
	parameterCount := len(method.parameterDescriptors)
	params := make([]t_any, parameterCount+1)
	for i := parameterCount; i >= 1; i-- {
		params[i] = f.pop()
	}
	// get objectref
	objectref := f.pop().(t_object)
	params[0] = objectref
	overridenMethod := objectref.class.findMethod(method.name + method.descriptor)
	frame := NewStackFrame(overridenMethod)
	// pass parameters
	for j := 0; j < parameterCount+1; j++ {
		frame.storeVar(uint(j), params[j])
	}

	this.pushFrame(frame)
}

func (this *Thread) invokeSpecialMethod(index uint16)  {
	f := this.peekFrame()
	m := f.method
	c := m.class
	method := c.constantPool[index].resolve().(*RuntimeConstantMethodrefInfo).method
	parameterCount := len(method.parameterDescriptors)
	frame := NewStackFrame(method)
	// pass parameters
	for i := parameterCount; i >= 1; i--  {
		argument := checkType(f.pop())
		frame.storeVar(uint(i), argument)
	}
	objectref := f.pop().(t_object)
	frame.storeVar(0, objectref) // this objectref

	this.pushFrame(frame)
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

func (this *StackFrame) getField(objectref t_object, index uint16) t_any {
	i := this.method.class.constantPool[index].resolve().(*RuntimeConstantFieldrefInfo).field.index
	return checkType(objectref.fields[i])
}

func (this *StackFrame) putField(objectref t_object, index uint16, value t_any) {
	i := this.method.class.constantPool[index].resolve().(*RuntimeConstantFieldrefInfo).field.index
	objectref.fields[i] = checkType(value)
}

func (this *StackFrame) push(jany t_any)  {
	operandStackSize := len(this.operandStack)
	this.operandStack = this.operandStack[:operandStackSize+1]
	this.operandStack[operandStackSize] = checkType(jany)
}

func (this *StackFrame) pop() t_any {
	operandStackSize := len(this.operandStack)
	jany := this.operandStack[operandStackSize-1]
	this.operandStack[operandStackSize-1] = nil
	this.operandStack = this.operandStack[:operandStackSize-1]
	return checkType(jany)
}

func (this *StackFrame) peek() t_any {
	operandStackSize := len(this.operandStack)
	jany := this.operandStack[operandStackSize-1]
	return checkType(jany)
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