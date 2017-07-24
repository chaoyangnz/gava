package jago

type ThreadManager struct {
	currentThread *Thread
}

func (this *ThreadManager) NewThread(name string) *Thread {
	thread := &Thread{
			name: name,
			vmStack: make([]*Frame, 0, DEFAULT_VM_STACK_SIZE)}
	thread.threadObject = NewJavaLangThread()
	this.currentThread = thread
	return thread
}

// We choose fix-sized stack size
const DEFAULT_VM_STACK_SIZE  = 512

type Thread struct {
	id          uint
	name        string
	vmStack     []*Frame

	threadObject JavaLangThread
}

func (this *Thread) Run()  {

	for len(this.vmStack) > 0 { // per stack frame
		this.runFrame()
	}
}

func (this *Thread) RunTo(frame *Frame)  {

	for { // per stack frame
		f := this.peekFrame()
		if f == nil || f == frame {
			break
		}
		this.runFrame()
	}
}

func __indent(thread *Thread, frame *Frame) string {
	var index int
	for i, f := range thread.vmStack {
		if f == frame {
			index = i
		}
	}

	str := ""
	for i:=0; i < index; i++ {
		str += "\t"
	}
	return str
}

func (this *Thread) runFrame()  {
	f := this.peekFrame()

	bytecode := f.method.code
	LOG.Info("\n%s🍏%s", __indent(this, f), f.method.Qualifier())
	for f.pc < len(f.method.code) {
		pc := f.pc
		opcode := bytecode[pc]
		instruction := instructions[opcode]
		jumped := false
		LOG.Debug("\n%s%04d ➢ %-18s", __indent(this, f), int(pc), instruction.mnemonic)
		intercept(f)
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

func intercept(f *Frame)  {
	//if f.method.Qualifier() == "java/lang/String.valueOf(Ljava/lang/Object;)Ljava/lang/String;" && f.pc == 13 {
	//	print("breakpoint")
	//}
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
	switch value.(type) {
	case Long, Double:
		this.localVariables[index+1] = value //occupy two places
	}
	this.localVariables[index] = value
}

func (this *Frame) const8(pos int) int8 {
	constant := int8(this.method.code[this.pc + pos])
	LOG.Debug("\t%d", constant)
	return constant
}

func (this *Frame) index8() uint8 {
	index := uint8(this.method.code[this.pc+1])
	LOG.Debug("\t#%d", index)
	return index
}

func (this *Frame) index16() uint16 {
	index := (uint16(this.method.code[this.pc+1]) << 8) | uint16(this.method.code[this.pc+2])
	LOG.Debug("\t#%d", index)
	return index
}

func (this *Frame) offset16() int16 {
	offset := int16((uint16(this.method.code[this.pc+1]) << 8) | uint16(this.method.code[this.pc+2]))
	LOG.Debug("\t⤋%d", this.pc + int(offset))
	return offset
}

func (this *Frame) params(callee *Method) []Value {
	parameterCount := len(callee.parameterDescriptors)
	if !callee.isStatic() {
		parameterCount += 1 // with an extra objectref: this
	}
	params := make([]Value, parameterCount)
	// (objectref) (arg1, arg2 ...) ->
	for i := parameterCount-1; i >= 0; i-- {
		param := this.pop()

		if callee.isNative() {
			if !callee.isStatic()  {
				if (i != 0) {
					index := i-1
					if callee.parameterDescriptors[index] == JVM_SIGNATURE_BOOLEAN { //Boolean
						param = param.(Int).ToBoolean() // Int -> Boolean
					}
				}
			} else {
				index := i
				if callee.parameterDescriptors[index] == JVM_SIGNATURE_BOOLEAN { //Boolean
					param = param.(Int).ToBoolean() // Int -> Boolean
				}
			}
		}

		params[i] = param
	}
	if !callee.isStatic() {
		// get objectref and target method
		objectref := params[0].(Reference)
		if objectref.IsNull() {
			Fatal("NullPointerException")
		}
	}

	return params
}

//func (this *Thread) invokeNativeMethod(method *Method, params ... Value) Value {
//	if !method.isNative() {
//		Fatal("Not a native method")
//	}
//	Debug("\n🍺 invoke native method %s", method.Qualifier())
//	name := "Java_" + strings.Replace(method.class.name + "_" + method.name, "/", "_", -1)
//	funcs := NATIVE_METHODS
//	if _, ok := funcs[name]; !ok {
//		Fatal( "%s does not exist.", name)
//	}
//	if len(params) != funcs[name].Type().NumIn() {
//		Fatal( "The number of params is not adapted.")
//	}
//	in := make([]reflect.Value, len(params))
//	for k, param := range params {
//		in[k] = reflect.ValueOf(param)
//	}
//	result := funcs[name].Call(in)
//	if len(result) == 0 {
//		return Value(nil)
//	}
//	return result[0].Interface().(Value)
//}

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
	return objectref.GetInstanceVariable(Int(i))
}

func (this *Frame) putField(objectref ObjectRef, index uint16, value Value) {
	i := this.method.class.constantPool[index].(*FieldRef).ResolvedField().index
	objectref.SetInstanceVariable(Int(i), value)
}

func (this *Frame) push(value Value)  {
	if value == nil { // check not-null
		Bug("Operand stack should never contain nil")
	}
	if boolean, ok := value.(Boolean); ok {
		// because interpreter cannot recognise Boolean, always use Int
		value = boolean.ToInt()
	}

	operandStackSize := len(this.operandStack)
	this.operandStack = this.operandStack[:operandStackSize+1]
	this.operandStack[operandStackSize] = value
}

func (this *Frame) pop() Value {
	operandStackSize := len(this.operandStack)
	Value := this.operandStack[operandStackSize-1]
	this.operandStack[operandStackSize-1] = nil
	this.operandStack = this.operandStack[:operandStackSize-1]
	return Value
}

func (this *Frame) clear()  {
	this.operandStack = make([]Value, 0, this.method.maxStack)
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
	if size == 0 {
		return nil
	}
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