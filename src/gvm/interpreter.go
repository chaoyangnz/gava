package gvm

import "fmt"

const MAIN_METHOD = "main([Ljava/lang/String;)V"

const DEFAULT_VM_STACK_SIZE  = 512

func run(mainClass *JavaClass)  {
	mainMethod := mainClass.findMethod(MAIN_METHOD)
	thread := &Thread{vmStack: VMStack{stackFrames: make([]*StackFrame, DEFAULT_VM_STACK_SIZE), size:0, capacity: DEFAULT_VM_STACK_SIZE}}
	thread.vmStack.push(NewStackFrame(mainMethod))
	for thread.vmStack.size != 0 { // per stack frame
		f := thread.vmStack.peek()
		for f.pc < uint32(len(f.method.code)) {
				next, progress := interpret(f, thread, f.method, f.method.class)
				f.pc += progress
				if !next {
					break
				}
			}
		}
}

func interpret(f *StackFrame, thread *Thread, method *JavaMethod, class *JavaClass) (bool, uint32) {
	bytecode := f.method.code
	opcode := bytecode[f.pc]
	switch opcode {
	case ICONST_1:
		f.pushInt(1)
		return true, 1
	case ICONST_2:
		f.pushInt(2)
		return true, 1
	case ISTORE_0:
		f.storeIntVar(0, f.popInt())
		return true, 1
	case ISTORE_1:
		f.storeIntVar(1, f.popInt())
		return true, 1
	case ISTORE_2:
		f.storeIntVar(2, f.popInt())
		return true, 1
	case ISTORE_3:
		f.storeIntVar(3, f.popInt())
		return true, 1
	case ISTORE:
		index := bytecode[f.pc+1]
		f.storeIntVar(uint(index), f.popInt())
		return true, 2
	case ILOAD_0:
		f.pushInt(f.loadIntVar(0))
		return true, 1
	case ILOAD_1:
		f.pushInt(f.loadIntVar(1))
		return true, 1
	case ILOAD_2:
		f.pushInt(f.loadIntVar(2))
		return true, 1
	case ILOAD_3:
		f.pushInt(f.loadIntVar(3))
		return true, 1
	case IADD:
		f.pushInt(f.popInt() + f.popInt())
		return true, 1
	case INVOKESTATIC:
		index := bytes2uint16(bytecode[f.pc+1:f.pc+3])

		methodRef := f.method.class.constantPool[index].resolve(f.method.class).(*RuntimeConstantMethodrefInfo)
		method := methodRef.method
		frame := NewStackFrame(method)
		// pass parameters
		f.passParameters(frame)

		thread.vmStack.push(frame)
		return false, 3
	case IRETURN:
		thread.vmStack.pop()
		// return value
		f.passReturn(thread.vmStack.peek())
		return false, 1
	case POP:
		thread.vmStack.pop()
		f.pc++
	case RETURN:
		thread.vmStack.pop()
		return false, 1
	case NEW:
		index := bytes2uint16(bytecode[f.pc+1:f.pc+3])
		classInfo := class.constantPool[index].resolve(class).(*RuntimeConstantClassInfo)
		jreference := &JavaObject{class: classInfo.class}
		jreference.fields = make([]java_any, len(classInfo.class.fields))
		f.pushReference(jreference)
		return true, 3
	case ANEWARRAY:
		index := bytes2uint16(bytecode[f.pc+1:f.pc+3])
		count := f.popInt()

		classInfo := class.constantPool[index].resolve(class).(*RuntimeConstantClassInfo)
		jreference := &JavaArray{aclass: classInfo.class, size: uint32(count)}
		jreference.elements = make([]java_any, uint32(count))
		f.pushArray(jreference)
		return true, 3
	case NEWARRAY:
		atype := uint8(bytecode[f.pc+1])
		count := f.popInt()
		jreference := &JavaArray{atype: atype, size: uint32(count)}
		jreference.elements = make([]java_any, uint32(count))
		f.pushArray(jreference)
		return true, 2
	default:
		//ignore
		fmt.Printf("No implementation for %d", opcode)
		panic("Abort")
		return true, 1
	}
	return true, 1
}


