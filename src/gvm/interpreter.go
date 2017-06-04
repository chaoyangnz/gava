package gvm

const MAIN_METHOD = "main([Ljava/lang/String;)V"

const DEFAULT_VM_STACK_SIZE  = 512

func run(mainClass *ClassMirror)  {
	mainMethod := mainClass.findMethod(MAIN_METHOD)
	thread := &Thread{vmStack: VMStack{stackFrames: make([]*StackFrame, DEFAULT_VM_STACK_SIZE), size:0, capacity: DEFAULT_VM_STACK_SIZE}}
	thread.vmStack.push(NewStackFrame(mainMethod))
	for thread.vmStack.size != 0 { // per stack frame
		f := thread.vmStack.peek()
		bytecode := f.method.code
	F: for f.pc < uint32(len(f.method.code)) {
			opcode := bytecode[f.pc]
			switch opcode {
			case ICONST_1:
				f.pushInt(1)
				f.pc++
			case ICONST_2:
				f.pushInt(2)
				f.pc++
			case ISTORE_0:
				f.storeIntVar(0, f.popInt())
				f.pc++
			case ISTORE_1:
				f.storeIntVar(1, f.popInt())
				f.pc++
			case ISTORE_2:
				f.storeIntVar(2, f.popInt())
				f.pc++
			case ISTORE_3:
				f.storeIntVar(3, f.popInt())
				f.pc++
			case ISTORE:
				index := bytecode[f.pc+1]
				f.storeIntVar(uint(index), f.popInt())
				f.pc += 2
			case ILOAD_0:
				f.pushInt(f.loadIntVar(0))
				f.pc++
			case ILOAD_1:
				f.pushInt(f.loadIntVar(1))
				f.pc++
			case ILOAD_2:
				f.pushInt(f.loadIntVar(2))
				f.pc++
			case ILOAD_3:
				f.pushInt(f.loadIntVar(3))
				f.pc++
			case IADD:
				f.pushInt(f.popInt() + f.popInt())
				f.pc++
			case INVOKESTATIC:
				index := (bytecode[f.pc+1] << 8) | bytecode[f.pc+2]
				f.pc += 3

				methodRef := f.method.class.constantPool[index].(*RuntimeConstantMethodrefInfo)
				f.method.class.resolveMethodRef(methodRef)
				method := methodRef.method
				frame := NewStackFrame(method)
				// pass parameters
				frame.storeIntVar(0, f.popInt())

				thread.vmStack.push(frame)
				break F
			case IRETURN:
				ret := f.popInt()
				thread.vmStack.pop()
				// return value
				thread.vmStack.peek().pushInt(ret)
				break F
			case POP:
				thread.vmStack.pop()
				f.pc++
			case RETURN:
				break F
			default:
				//ignore
				f.pc++
			}

		}
	}
}


