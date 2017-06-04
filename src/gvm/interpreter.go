package gvm

const MAIN_METHOD = "main:([Ljava/lang/String;)V"

const DEFAULT_OPERAND_STACK_SIZE = 512*1024
const DEFAULT_VM_STACK_SIZE  = 512

func run(mainClass *ClassMirror)  {
	mainMethod := mainClass.FindMethod(MAIN_METHOD)
	thread := &Thread{vmStack: VMStack{stackFrames: make([]*StackFrame, DEFAULT_VM_STACK_SIZE), size:0, capacity: DEFAULT_VM_STACK_SIZE}}
	stackFrame := NewStackFrame(mainMethod)
	thread.vmStack.push(stackFrame)
	for thread.vmStack.size != 0 { // per stack frame
		f := thread.vmStack.pop()
		bytecode := f.method.code
		for f.pc < len(bytecode) {
			opcode := bytecode[f.pc]
			switch opcode {
			case ICONST_1:
				f.pushInt(1)
			case ICONST_2:
				f.pushInt(2)
			case ISTORE_1:
				f.storeIntVar(1, f.popInt())
			case ISTORE_2:
				f.storeIntVar(2, f.popInt())
			case ISTORE_3:
				f.storeIntVar(3, f.popInt())
			case ILOAD_1:
				f.pushInt(f.loadIntVar(1))
			case ILOAD_2:
				f.pushInt(f.loadIntVar(2))
			case IADD:
				f.pushInt(f.popInt() + f.popInt())
			default:
				//ignore
			}
			f.pc++
		}
	}
}


