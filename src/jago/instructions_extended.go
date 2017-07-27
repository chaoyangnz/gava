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

	index := uint16(indexbyte1) << 8 | uint16(indexbyte2)
	counts := make([]Int, dimensions)
	for i:=dimensions-1; i >= 0; i-- {
		counts[i] = f.pop().(Int)
	}

	class := c.constantPool[index].(ClassRef).ResolvedClass()

	arrayref := class.NewArray()
	//index := f.index16()
	//dimensions := m.code[f.pc+3]
	//
	//arrayType := c.constantPool[index].(*ClassRef).class.(*ArrayClass)
	//componentType := arrayType.componentType.(ClassType)
	//counts := make([]jint, dimensions)
	//for i := jint(dimensions-1); i >= 0; i-- {
	//	counts[i] = f.pop().(jint)
	//}
	//var arrayref *Array
	//componentArray := newArray(componentType, counts[0])
	//for j := uint8(1); j < dimensions-1; j++ {
	//	arrayref = newArray(componentType, counts[j])
	//	for k := jint(0); k < counts[j]; k++ {
	//		arrayref.elements[k] = componentArray
	//	}
	//	componentArray = arrayref
	//}
	//f.push(arrayref)
	Fatal("Not implemented for opcode %d\n", opcode)
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
