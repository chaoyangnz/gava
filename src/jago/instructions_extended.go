package jago

/*196 (0xC4)*/
func WIDE(opcode uint8, t *Thread, f *Frame, c *Class, m *Method) {
	wide_opcode := m.code[f.pc+1]
	instruction := instructions[wide_opcode]
	switch instruction.mnemonic {
	case "iload", "fload", "aload", "lload", "dload":
		indexbyte1 := m.code[f.pc+2]
		indexbyte2 := m.code[f.pc+3]

		index := uint16(indexbyte1) << 8 | uint16(indexbyte2)
		EXEC_LOG.Debug("\t%s\t#%d", instruction.mnemonic, index)

		f.push(f.loadVar(uint(index)))
		f.offsetPc(4)
	case "istore", "fstore", "astore", "lstore", "dstore":
		indexbyte1 := m.code[f.pc+2]
		indexbyte2 := m.code[f.pc+3]

		index := uint16(indexbyte1) << 8 | uint16(indexbyte2)
		EXEC_LOG.Debug("\t%s\t#%d", instruction.mnemonic, index)

		f.storeVar(uint(index), f.pop())
		f.offsetPc(4)
	case "ret":
		indexbyte1 := m.code[f.pc+2]
		indexbyte2 := m.code[f.pc+3]

		index := uint16(indexbyte1) << 8 | uint16(indexbyte2)
		EXEC_LOG.Debug("\t%s\t#%d", instruction.mnemonic, index)
		// IGNORE

		f.offsetPc(4)
	case "iinc":
		indexbyte1 := m.code[f.pc+2]
		indexbyte2 := m.code[f.pc+3]
		constbyte1 := m.code[f.pc+4]
		constbyte2 := m.code[f.pc+5]

		index := uint16(indexbyte1) << 8 | uint16(indexbyte2)
		const_value := int16(uint16(constbyte1) << 8 | uint16(constbyte2))

		EXEC_LOG.Debug("\t%s\t#%d\t%d", instruction.mnemonic, index, const_value)

		value := f.loadVar(uint(index)).(Int)
		f.storeVar(uint(index), value + Int(const_value))
		f.offsetPc(6)
	default:
		Fatal("Not support opcode % for wide", instruction.mnemonic)
	}
}

/*197 (0xC5)*/
func MULTIANEWARRAY(opcode uint8, t *Thread, f *Frame, c *Class, m *Method) {
	indexbyte1 := m.code[f.pc+1]
	indexbyte2 := m.code[f.pc+2]
	dimensions := m.code[f.pc+3]

	if dimensions < 1 {
		Fatal("The dimension of multi-dimensional array must be >= 1")
	}
	index := uint16(indexbyte1) << 8 | uint16(indexbyte2)

	EXEC_LOG.Debug("\t%d\t dim %d:", index, dimensions)

	counts := make([]Int, dimensions)
	for j:= int(dimensions-1); j >= 0; j-- {
		counts[j] = f.pop().(Int)
		if counts[j] < 0 {
			VM_throw("java/lang/NegativeArraySizeException", "Array size cannot be negative")
		}
		EXEC_LOG.Debug("\t%d", counts[j])
	}

	class := c.constantPool[index].(*ClassRef).ResolvedClass()
	if !class.IsArray() {
		Fatal("Non-Array class %s cannot be used to new multi-dimensional array", class.name)
	}

	f.push(newMultiDimensioalArray(counts, class))
	f.nextPc()
}

func newMultiDimensioalArray(counts []Int, class *Class) ArrayRef {
	count := counts[0]
	arr := class.NewArray(count)

	if len(counts) > 1 {
		elements := arr.Elements()
		for i := range elements {
			elements[i] = newMultiDimensioalArray(counts[1:], class.componentType.(*Class))
		}
	}

	return  arr
}

/*198 (0xC6)*/
func IFNULL(opcode uint8, t *Thread, f *Frame, c *Class, m *Method) {
	offset := f.offset16()

	value := f.pop().(Reference)

	if value.IsNull() {
		f.offsetPc(offset)
	} else {
		f.nextPc()
	}
}

/*199 (0xC7)*/
func IFNONNULL(opcode uint8, t *Thread, f *Frame, c *Class, m *Method) {
	offset := f.offset16()

	value := f.pop().(Reference)

	if !value.IsNull() {
		f.offsetPc(offset)
	} else {
		f.nextPc()
	}
}

/*200 (0xC8)*/
func GOTO_W(opcode uint8, t *Thread, f *Frame, c *Class, m *Method) {
	Fatal("Not implemented for opcode %d\n", opcode)
}

/*201 (0xC9)*/
func JSR_W(opcode uint8, t *Thread, f *Frame, c *Class, m *Method) {
	Fatal("Not implemented for opcode %d\n", opcode)
}
