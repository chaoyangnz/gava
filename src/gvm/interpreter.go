package gvm

import "fmt"

const MAIN_METHOD = "main([Ljava/lang/String;)V"

const DEFAULT_VM_STACK_SIZE  = 512

func run(mainClass *JavaClass)  {
	mainMethod := mainClass.findMethod(MAIN_METHOD)
	thread := &Thread{
		name: "main",
		vmStack: VMStack{stackFrames: make([]*StackFrame, DEFAULT_VM_STACK_SIZE), size:0, capacity: DEFAULT_VM_STACK_SIZE}}
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
		f.push(java_int(1))
		return true, 1
	case ICONST_2:
		f.push(java_int(2))
		return true, 1
	case ICONST_3:
		f.push(java_int(3))
		return true, 1
	case ISTORE_0:
		f.storeVar(0, f.pop())
		return true, 1
	case ISTORE_1:
		f.storeVar(1, f.pop())
		return true, 1
	case ISTORE_2:
		f.storeVar(2, f.pop())
		return true, 1
	case ISTORE_3:
		f.storeVar(3, f.pop())
		return true, 1
	case ISTORE:
		index := bytecode[f.pc+1]
		f.storeVar(uint(index), f.pop())
		return true, 2
	case ILOAD_0:
		f.push(f.loadVar(0).(java_int))
		return true, 1
	case ILOAD_1:
		f.push(f.loadVar(1).(java_int))
		return true, 1
	case ILOAD_2:
		f.push(f.loadVar(2).(java_int))
		return true, 1
	case ILOAD_3:
		f.push(f.loadVar(3).(java_int))
		return true, 1
	case ILOAD:
		index := bytecode[f.pc+1]
		f.push(f.loadVar(uint(index)).(java_int))
		return true, 2
	case IADD:
		f.push(f.pop().(java_int) + f.pop().(java_int))
		return true, 1
	case ALOAD_0:
		f.push(f.loadVar(0).(java_object))
		return true, 1
	case ASTORE:
		index := bytecode[f.pc+1]
		f.storeVar(uint(index), f.pop().(java_object))
		return true, 2
	case ALOAD:
		index := bytecode[f.pc+1]
		f.push(f.loadVar(uint(index)).(java_object))
		return true, 2
	case INVOKESTATIC:
		index := bytes2uint16(bytecode[f.pc+1:f.pc+3])

		methodRef := f.method.class.constantPool[index].resolve().(*RuntimeConstantMethodrefInfo)
		method := methodRef.method
		if method.isNative() {
			GVM_print(f.pop().(jstring))
			return true, 3
		}
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
		class := class.constantPool[index].resolve().(*RuntimeConstantClassInfo).class
		jreference := class.new()
		f.push(jreference)
		return true, 3
	case ANEWARRAY:
		index := bytes2uint16(bytecode[f.pc+1:f.pc+3])
		count := f.pop().(java_int)

		classInfo := class.constantPool[index].resolve().(*RuntimeConstantClassInfo)
		jreference := java_array(&JavaArray{aclass: classInfo.class, size: uint32(count)})
		jreference.elements = make([]java_any, uint32(count))
		f.push(jreference)
		return true, 3
	case NEWARRAY:
		atype := uint8(bytecode[f.pc+1])
		count := f.pop().(java_int)
		jreference := java_array(&JavaArray{atype: atype, size: uint32(count)})
		jreference.elements = make([]java_any, uint32(count))
		f.push(jreference)
		return true, 2
	case GETFIELD:
		index := bytes2uint16(bytecode[f.pc+1:f.pc+3])
		jreference := f.pop().(java_object)

		f.push((*jreference).getField(class, index))
		return true, 3
	case PUTFIELD:
		index := bytes2uint16(bytecode[f.pc+1:f.pc+3])

		value := f.pop()
		jreference := f.pop().(java_object)

		(*jreference).putField(class, index, value)
		return true, 3
	case DUP:
		value := f.peek()
		f.push(value)
		return true, 1
	case INVOKESPECIAL:
		index := bytes2uint16(bytecode[f.pc+1:f.pc+3])

		method := f.method.class.constantPool[index].resolve().(*RuntimeConstantMethodrefInfo).method
		frame := NewStackFrame(method)
		// pass parameters
		f.passParameters(frame)

		thread.vmStack.push(frame)
		return false, 3
	case BIPUSH:
		byte := bytecode[f.pc+1]
		f.push(java_int(byte))
		return true, 2
	case LDC:
		index := bytecode[f.pc+1]
		cpInfo := class.constantPool[index].resolve()
		switch cpInfo.(type) {
		case *RuntimeConstantIntegerInfo:
			f.push(cpInfo.(*RuntimeConstantIntegerInfo).value)
		case *RuntimeConstantFloatInfo:
			f.push(cpInfo.(*RuntimeConstantFloatInfo).value)
		case *RuntimeConstantStringInfo:
			f.push(cpInfo.(*RuntimeConstantStringInfo).value)
		}
		return true, 2
	case ICONST_5:
		cpInfo := class.constantPool[5].resolve()
		switch cpInfo.(type) {
		case *RuntimeConstantIntegerInfo:
			f.push(cpInfo.(*RuntimeConstantIntegerInfo).value)
		case *RuntimeConstantFloatInfo:
			f.push(cpInfo.(*RuntimeConstantFloatInfo).value)
		case *RuntimeConstantStringInfo:
			f.push(cpInfo.(*RuntimeConstantStringInfo).value)
		}
		return true, 1
	default:
		//ignore
		fmt.Printf("No implementation for opcode %d\n", opcode)
		panic("Abort")
		return true, 1
	}
	return true, 1
}


