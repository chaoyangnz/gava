package gvm

const MAIN_METHOD = "main:([Ljava/lang/String;)V"

const DEFAULT_OPERAND_STACK_SIZE = 512*1024
const DEFAULT_VM_STACK_SIZE  = 512

func Run(mainClass *ClassMirror)  {
	mainMethod := mainClass.FindMethod(MAIN_METHOD)
	thread := &Thread{vmStack: VMStack{stackFrames: make([]*StackFrame, DEFAULT_VM_STACK_SIZE), size:0, capacity: DEFAULT_VM_STACK_SIZE}}
	stackFrame := NewStackFrame(mainMethod)
	thread.vmStack.push(stackFrame)
	for thread.vmStack.size != 0 { // per stack frame
		currentStackFrame := thread.vmStack.pop()
		bytecode := currentStackFrame.method.code
		for currentStackFrame.pc < len(bytecode) {
			opcode := bytecode[currentStackFrame.pc]
			switch opcode {
			case ICONST_1:
				currentStackFrame.operandStack.pushInt(1)
			case ISTORE_1:
				print(currentStackFrame.loadIntLocalVariable(1))
			default:
				//ignore
			}
			currentStackFrame.pc++
		}
	}
}

