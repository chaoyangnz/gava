package jago

import (
	//"os"
	"strconv"
	"github.com/jtolds/gls"
	"github.com/orcaman/concurrent-map"
	"fmt"
	"time"
)

var thread_id_key = gls.GenSym()
type ThreadManager struct {
	//contextManager *gls.ContextManager
	threads cmap.ConcurrentMap // [string]*Thread
}

func tid2str(tid uint64) string {
	return fmt.Sprintf("%v", tid)
}

func (this *ThreadManager) register(thread *Thread)  {
	this.threads.Set(tid2str(thread.id), thread)
}

func (this *ThreadManager) unregister(thread *Thread)  {
	this.threads.Remove(tid2str(thread.id))
}

func (this *ThreadManager) exitDaemonThreads()  {
	for _, t := range this.threads.Items() {
		thread := t.(*Thread)
		if thread.daemon {
			LOG.Info("[thread]Thread '%s' interrupt \n", thread.name)
			//thread.interrupt()
		}
	}
}

func (this *ThreadManager) current() *Thread {
	gid := getGID()
	if t, ok := this.threads.Get(tid2str(gid)); ok {
		return t.(*Thread)
	}

	Bug("Cannot get current thread")
	return nil
}

func (this *ThreadManager) RunBootstrapThread(run func()) {
	thread := precreateThread("bootstrap", run, func(){})
	thread.id = getGID()
	THREAD_MANAGER.register(thread) // register to THREAD_MANAGER
	// this java.lang.Thread object hasn't been initialized, defer after System.initializeSystemClasses(..)
	thread.threadObject = NewJavaLangThread(thread.name)

	thread.started = true
	LOG.Info("[thread]Start to run bootstrap thread %s #%d\n", VM_currentThread().name, VM_currentThread().id)

	// start Java code execution
	thread.runBlock.Do()
}

func precreateThread(name string, run func(), exitHook func()) *Thread {
	thread := &Thread{
		name: name,
		vmStack: make([]*Frame, 0, DEFAULT_VM_STACK_SIZE),
		started: false,
		//cond: sync.NewCond(&sync.Mutex{}),
		ch: make(chan int),
	}

	thread.runBlock = Block{
		func() { run() },
		func(throwable Reference)  {
			EXEC_LOG.Info("\n💥Exception uncaught  in thread \"%s\", thus print stacktrace on stdout", thread.name)
			detailMessage := throwable.GetInstanceVariableByName("detailMessage", "Ljava/lang/String;").(JavaLangString)
			detailMessageStr := ""
			if !detailMessage.IsNull() {
				detailMessageStr = detailMessage.toNativeString()
			}
			VM_stderrPrintf("\nException in thread \"%s\" #%d %s: %s\n", thread.name, thread.id, throwable.Class().Name(), detailMessageStr)

			stacktrace := throwable.GetExtra()
			if stacktrace != nil {
				for _, stacktraceelement := range stacktrace.([]string) {
					VM_stderrPrintf("\t at %s\n", stacktraceelement)
				}
			}
		},
		func() {
			exitHook()
		},
	}

	return thread
}

const (
	_start  = 0
	_interrupt = 1
	_timeout = 2
)

func (this *ThreadManager) NewThread(name string, run func(), exitHook func()) *Thread {
	thread := precreateThread(name, run, exitHook)

	VM_WG.Add(1)
	go func() {
		// prepare the thread environment
		thread.id = getGID() // set go routine id as thread id
		THREAD_MANAGER.register(thread) // register to THREAD_MANAGER

		// pause here to wait start command
		for !thread.started {
			LOG.Info("[thread]Created thread '%s' #%d but wait to start\n", VM_currentThread().name, VM_currentThread().id)
			thread.started = <- thread.ch == _start
		}

		LOG.Info("[thread]Start thread '%s' #%d\n", VM_currentThread().name, VM_currentThread().id)
		// start Java code execution
		thread.runBlock.Do()

		// prepare to exit
		LOG.Info("[thread]Exit thread '%s' %d \n", thread.name, thread.id)
		this.unregister(thread)
		VM_WG.Done()
	}()

	return thread
}

// We choose fix-sized stack size
const DEFAULT_VM_STACK_SIZE  = 512

type Thread struct {
	id          uint64
	name        string
	vmStack     []*Frame
	runBlock    Block
	started     bool

	ch          chan int

	daemon      bool

	//cond        *sync.Cond

	threadObject JavaLangThread
}

// start thread
func (this *Thread) start()  {
	if this.started {
		return
	}


	if this.threadObject == nil {
		// ----- prepare java.lang.Thread object ----
		// this construct object must be run in parent thread

		// this java.lang.Thread object hasn't been initialized, defer after System.initializeSystemClasses(..)
		this.threadObject = NewJavaLangThread(this.name)
		// once system classes are initialized, set up the java.lang.Thread object
		threadConstructor := this.threadObject.Class().GetConstructor("(Ljava/lang/String;)V")
		VM_invokeMethod(threadConstructor, this.threadObject, NewJavaLangString(this.name))
	}

	// get daemon
	this.daemon = this.threadObject.GetInstanceVariableByName("daemon", "Z").(Boolean).IsTrue()

	// signal to run thread
	this.ch <- _start
}

func (this *Thread) interrupt()  {
	this.ch <- _interrupt
}

func (this *Thread) sleep(timeout time.Duration) bool {
	go func() {
		time.Sleep(timeout)

		this.ch <- _timeout
	}()

	interrupted := <- this.ch == _interrupt

	return interrupted
}

/*
  _____
  |   |   <-  a new frame added on top, no touch existing. just run this frame
----------
  |   |   <- existing
  ______
  |   |
  ______

Only run one single frame
*/
func (this *Thread) ExecuteFrame() /* this return is throwable if this method is return exceptionally, otherwise nil */{
	f := this.current()
	bytecode := f.method.code
	if f.pc == 0 {
		EXEC_LOG.Debug("\n%s🔹[%s]%s", repeat("\t", this.indexOf(f)),this.name, f.method.Qualifier())
	}

	for f.pc < len(f.method.code) {
		pc := f.pc
		opcode := bytecode[pc]
		instruction := instructions[opcode]
		EXEC_LOG.Trace("\n%s%04d ➢ %-18s", repeat("\t", this.indexOf(f)), int(pc), instruction.mnemonic)
		this.interceptBefore(f)
		// each instruction execution can throw exception
		Block{
			func() { instruction.interpret(this, f, f.method.class, f.method) },
			func(throwable Reference) {
				// try catch exception
				caught := false
				handlePc := -1
				for _, exception := range f.method.exceptions {
					if f.pc >= exception.startPc && f.pc < exception.endPc {
						if exception.catchType == 0 { // catch-all
							caught, handlePc = true, exception.handlerPc
							break
						} else {
							catchType := f.method.class.constantPool[int32(exception.catchType)].(*ClassRef).ResolvedClass()
							if catchType.IsAssignableFrom(throwable.Class()) {
								caught, handlePc = true, exception.handlerPc
								break
							}
						}
					}
				}

				if caught {
					EXEC_LOG.Info("\n%s💧Exception caught: %s at %s", repeat("\t", this.indexOf(f)+1), throwable.Class().name, f.method.Qualifier())
					f.jumpPc(handlePc) // move pc
					f.clear()
					f.push(throwable)
					return
				} else {
					// set exception here
					f.exception = throwable
					this.pop()
				}
			},
			func() {},
		}.Do()
		this.interceptAfter(f)

		if len(this.vmStack) == 0 || f != this.current() || !f.exception.IsNull() {
			// 1) normal return and it's the last method
			// 2) call another method
			// 3) current method throws exception but has not caught this exception. To be rethrow in its caller
			break
		}
		if pc == f.pc { // not jump
			f.nextPc()
		}
		// otherwise normal jump or exception handle in current method: just go to!
		// jump instruction can operate pc; some instruction also have variable length: tableswitch...
		// these instructions will control pc themselves, if instruction operates the stack, we follow it
	}
}

func (this *Thread) interceptBefore(frame *Frame)  {

}

func (this *Thread) interceptAfter(frame *Frame)  {

}

type Frame struct {
	method *Method
	// if this this is current this, the pc is for the pc of this thread;
	// otherwise, it is a snapshot one since the last time
	pc int
	pos int // operand pos: internal use only. For an instruction, initially it always starts from pc. Each time read an operand, it advanced.
	// long and double will occupy two variable indexes
	localVariables      []Value
	// operand stack
	// a value of type `long` or `double` contributes two units to the indexOf and a value of any other type contributes one unit
	// but here we use long and double only use one unit. There is not any violation
	operandStack        []Value

	exception    Reference
}

/*
Should be called only by Run() method
 */
func (this *Frame) nextPc()  {
	opcode := this.method.code[this.pc]
	this.pc += JVM_OPCODE_LENGTH_INITIALIZER[opcode]
	this.pos = this.pc + 1
}

func (this *Frame) offsetPc(offset int16)  {
	this.jumpPc(this.pc + int(offset))
}

func (this *Frame) offsetPc32(offset int32)  {
	this.jumpPc(this.pc + int(offset))
}

func (this *Frame) jumpPc(to int)  {
	if this.pc == to {
		VM_throw("java/lang/InternalError", "Dead loop with empty body has been detected. Possibly you write: for(;;){} or while(true){}")
	}
	this.pc = to
	this.pos = this.pc + 1
}

func NewStackFrame(method *Method) *Frame {
	stackFrame := &Frame{
		method: method,
		pc: 0,
		pos: 1, // all opcode is 1 byte, then operand starts with 1
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

func (this *Frame) opcode() uint8 {
	return this.method.code[this.pc]
}

/*
padding operand start pos to multiply of 4
 */
func (this *Frame) operandPadding() {
	var start int
	for i:= 0; i <=3; i++ {
		if (this.pos+i) % 4 == 0 {
			start = this.pos +i
			break;
		}
	}
	this.pos = start
}

func (this *Frame) operandConst8() int8 {
	constant := int8(this.method.code[this.pos])
	EXEC_LOG.Trace("\t%d", constant)
	this.pos += 1
	return constant
}

func (this *Frame) operandConst16() int16 {
	constbyte1 := this.method.code[this.pos]
	constbyte2 := this.method.code[this.pos+1]
	constant := int16(uint16(constbyte1) << 8 | uint16(constbyte2))
	EXEC_LOG.Trace("\t%d", constant)
	this.pos += 2
	return constant
}

func (this *Frame) operandConst32() int32 {
	constbyte1 := this.method.code[this.pos]
	constbyte2 := this.method.code[this.pos+1]
	constbyte3 := this.method.code[this.pos+2]
	constbyte4 := this.method.code[this.pos+3]
	constant := int32((uint32(constbyte1) << 24) | (uint32(constbyte2) << 16) | (uint32(constbyte3) << 8) | uint32(constbyte4))
	EXEC_LOG.Trace("\t%d", constant)
	this.pos += 4
	return constant
}

func (this *Frame) operandIndex8() uint8 {
	index := uint8(this.method.code[this.pos])
	EXEC_LOG.Trace("\t♯%d", index)
	this.pos += 1
	return index
}

func (this *Frame) operandIndex16() uint16 {
	indexbyte1 := this.method.code[this.pos]
	indexbyte2 := this.method.code[this.pos+1]
	index := (uint16(indexbyte1) << 8) | uint16(indexbyte2)
	EXEC_LOG.Trace("\t♯%d", index)
	this.pos += 2
	return index
}

func (this *Frame) operandOffset16() int16 {
	offsetbyte1 := this.method.code[this.pos]
	offsetbyte2 := this.method.code[this.pos+1]

	offset := int16((uint16(offsetbyte1) << 8) | uint16(offsetbyte2))
	EXEC_LOG.Trace("\t»%d (%s)", this.pc + int(offset), numberWithSign(int32(offset)))
	this.pos += 2
	return offset
}

func (this *Frame) operandOffset32() int32 {
	offsetbyte1 := this.method.code[this.pos]
	offsetbyte2 := this.method.code[this.pos+1]
	offsetbyte3 := this.method.code[this.pos+2]
	offsetbyte4 := this.method.code[this.pos+3]

	offset := int32((uint32(offsetbyte1) << 24) | (uint32(offsetbyte2) << 16) | (uint32(offsetbyte3) << 8) | uint32(offsetbyte4))
	EXEC_LOG.Trace("\t»%d (%s)", this.pc + int(offset), numberWithSign(offset))

	this.pos += 4
	return offset
}

//func (this *Frame) operandByte() int32 {
//	byte := uint32(this.method.code[this.pos])
//	this.pos += 1
//	return int32(byte)
//}

func (this *Frame) operandUByte() uint8 {
	byte := this.method.code[this.pos]
	this.pos += 1
	return byte
}

func (this *Frame) loadParameters(callee *Method) []Value {
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

	return params
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
	return objectref.GetInstanceVariable(Int(i))
}

func (this *Frame) putField(objectref ObjectRef, index uint16, value Value) {
	field := this.method.class.constantPool[index].(*FieldRef).ResolvedField()
	i := field.index
	if field.descriptor == "Z" {
		value = value.(Int).ToBoolean()
	}
	objectref.SetInstanceVariable(Int(i), value)
}

func (this *Frame) getSourceFileAndLineNumber() string {
	sourceFile := this.method.class.sourceFile
	if sourceFile == "" {
		sourceFile = "<Unknow>"
	}
	lineNumber := -1
	lineNumbers := this.method.lineNumbers
	for i := len(lineNumbers)-1; i >= 0; i-- {
		entry := lineNumbers[i]
		if this.pc >= entry.startPc {
			lineNumber = int(entry.lineNumber)
			break
		}
	}

	linenum := ""
	if lineNumber >= 0 {
		linenum = strconv.FormatInt(int64(lineNumber), 10)
	}
	return "(" + sourceFile + ":" + linenum + ")"
}

func (this *Frame) push(value Value)  {
	if value == nil { // check not-null
		Bug("Operand stack should never contain nil")
	}
	if value == VOID {
		return
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

//func (this *Frame) peek() Value {
//	operandStackSize := len(this.operandStack)
//	Value := this.operandStack[operandStackSize-1]
//	return Value
//}

/*
Should be called only by VM_invokeMethod(..)
 */
func (this *Thread) push(stackFrame *Frame)  {
	size := len(this.vmStack)
	if size == DEFAULT_VM_STACK_SIZE {
		VM_throw("java/lang/StackOverflowError", "Exceed the maximum stack size")
	}
	this.vmStack = this.vmStack[:size+1]
	this.vmStack[size] = stackFrame
}

func (this *Thread) pop()  {
	size := len(this.vmStack)
	if size == 0 {
		return
	}
	this.vmStack = this.vmStack[:size-1]
}

func (this *Thread) current() *Frame {
	size := len(this.vmStack)
	if size == 0 {
		return nil
	}
	return this.vmStack[size-1]
}

func (this *Thread) indexOf(frame *Frame) int {
	for i, f := range this.vmStack {
		if f == frame {
			return i
		}
	}

	Bug("Frame is not in the VM stack")
	return -1
}