package jago

/*196 (0xC4)*/
func WIDE(t *Thread, f *Frame, c *Class, m *Method) {
	wide_opcode := f.operandUByte()

	instruction := VM.GetInstruction(wide_opcode)
	t.Trace("\t%s", instruction.mnemonic)

	switch instruction.mnemonic {
	case "iload", "fload", "aload", "lload", "dload":
		index := f.operandIndex16()

		f.push(f.loadVar(uint(index)))
		f.offsetPc(4)
	case "istore", "fstore", "astore", "lstore", "dstore":
		index := f.operandIndex16()

		f.storeVar(uint(index), f.pop())
		f.offsetPc(4)
	case "ret":
		/*index :=*/ f.operandIndex16()

		// IGNORE
		f.offsetPc(4)
	case "iinc":
		index := f.operandIndex16()
		const_value := f.operandConst16()

		value := f.loadVar(uint(index)).(Int)
		f.storeVar(uint(index), value+Int(const_value))
		f.offsetPc(6)
	default:
		Fatal("Not support opcode % for wide", instruction.mnemonic)
	}
}

/*197 (0xC5)*/
func MULTIANEWARRAY(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.operandIndex16()
	dimensions := f.operandUByte()

	if dimensions < 1 {
		Fatal("The dimension of multi-dimensional array must be >= 1")
	}

	t.Trace("\t%d\t dim %d:", index, dimensions)

	counts := make([]Int, dimensions)
	for j := int(dimensions - 1); j >= 0; j-- {
		counts[j] = f.pop().(Int)
		if counts[j] < 0 {
			VM.Throw("java/lang/NegativeArraySizeException", "Array size cannot be negative")
		}
		t.Trace("\t%d", counts[j])
	}

	class := c.constantPool[index].(*ClassRef).ResolvedClass()
	if !class.IsArray() {
		Fatal("Non-Array class %s cannot be used to new multi-dimensional array", class.name)
	}

	f.push(VM.NewMultiDimensioalArray(class, counts))
}

/*198 (0xC6)*/
func IFNULL(t *Thread, f *Frame, c *Class, m *Method) {
	offset := f.operandOffset16()

	value := f.pop().(Reference)

	if value.IsNull() {
		f.offsetPc(offset)
	}
}

/*199 (0xC7)*/
func IFNONNULL(t *Thread, f *Frame, c *Class, m *Method) {
	offset := f.operandOffset16()

	value := f.pop().(Reference)

	if !value.IsNull() {
		f.offsetPc(offset)
	}
}

/*200 (0xC8)*/
func GOTO_W(t *Thread, f *Frame, c *Class, m *Method) {
	Fatal("Not implemented for opcode %d\n", f.opcode())
}

/*201 (0xC9)*/
func JSR_W(t *Thread, f *Frame, c *Class, m *Method) {
	Fatal("Not implemented for opcode %d\n", f.opcode())
}
