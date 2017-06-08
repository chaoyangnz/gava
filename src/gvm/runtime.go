package gvm


const DEFAULT_VM_STACK_SIZE  = 512

type Thread struct {
	id          uint
	name        string
	vmStack     VMStack
}

func newThread(name string) *Thread  {
	return &Thread{
		name: name,
		vmStack: VMStack{stackFrames: make([]*StackFrame, DEFAULT_VM_STACK_SIZE), size:0, capacity: DEFAULT_VM_STACK_SIZE}}
}


type StackFrame struct {
	method *JavaMethod
	// if this frame is current frame, the pc is for the pc of this thread;
	// otherwise, it is a snapshot one since the last time
	pc uint32
	localVariables      []t_any
	// operand stack
	operandStack        []t_any
	operandStackSize    uint
}

func NewStackFrame(method *JavaMethod) *StackFrame {
	stackFrame := &StackFrame{
		method: method,
		pc: 0,
		localVariables: make([]t_any, method.maxLocals),
		operandStack: make([]t_any, method.maxStack),
		operandStackSize: 0}
	return stackFrame
}

func (this *StackFrame) loadVar(index uint) t_any {
	return this.localVariables[index]
}

func (this *StackFrame) storeVar(index uint, value t_any)  {
	this.localVariables[index] = value
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

func (this *StackFrame) push(jany t_any)  {
	this.operandStack[this.operandStackSize] = jany
	this.operandStackSize++
}

func (this *StackFrame) pop() t_any {
	jany := this.operandStack[this.operandStackSize-1]
	this.operandStack[this.operandStackSize-1] = nil
	this.operandStackSize--
	return jany
}

func (this *StackFrame) peek() t_any {
	jany := this.operandStack[this.operandStackSize-1]
	return jany
}



type VMStack struct {
	stackFrames []*StackFrame
	size uint
	capacity uint
}

func (this *VMStack) push(stackFrame *StackFrame)  {
	this.stackFrames[this.size] = stackFrame
	this.size++
}

func (this *VMStack) pop()  {
	if this.size == 0 {
		return
	}
	this.stackFrames[this.size-1] = nil
	this.size--
}

func (this *VMStack) peek() *StackFrame  {
	return this.stackFrames[this.size-1]
}