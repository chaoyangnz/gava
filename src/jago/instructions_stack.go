package jago

/*87 (0x57)*/
func POP(opcode uint8, t *Thread, f *Frame, c *Class, m *Method) {
	f.pop()
	f.nextPc()
}

/*88 (0x58)*/
func POP2(opcode uint8, t *Thread, f *Frame, c *Class, m *Method) {
	Fatal("Not implemented for opcode %d\n", opcode)
}

/*89 (0x59)*/
func DUP(opcode uint8, t *Thread, f *Frame, c *Class, m *Method) {
	value := f.pop()
	f.push(value)
	f.push(value)
	f.nextPc()
}

/*90 (0x5A)*/
func DUP_X1(opcode uint8, t *Thread, f *Frame, c *Class, m *Method) {
	value1 := f.pop()
	value2 := f.pop()
	f.push(value1)
	f.push(value2)
	f.push(value1)
	f.nextPc()
}

/*91 (0x5B)*/
func DUP_X2(opcode uint8, t *Thread, f *Frame, c *Class, m *Method) {
	Fatal("Not implemented for opcode %d\n", opcode)
}

/*92 (0x5C)*/
func DUP2(opcode uint8, t *Thread, f *Frame, c *Class, m *Method) {
	value1 := f.pop()
	switch value1.(type) {
	case Long, Double:
		f.push(value1)
		f.push(value1)
	default:
		value2 := f.pop()
		f.push(value2)
		f.push(value1)
		f.push(value2)
		f.push(value1)
	}
	f.nextPc()
}

/*93 (0x5D)*/
func DUP2_X1(opcode uint8, t *Thread, f *Frame, c *Class, m *Method) {
	Fatal("Not implemented for opcode %d\n", opcode)
}

/*94 (0x5E)*/
func DUP2_X2(opcode uint8, t *Thread, f *Frame, c *Class, m *Method) {
	Fatal("Not implemented for opcode %d\n", opcode)
}

/*95 (0x5F)*/
func SWAP(opcode uint8, t *Thread, f *Frame, c *Class, m *Method) {
	value1 := f.pop()
	value2 := f.pop()
	f.push(value1)
	f.push(value2)
	f.nextPc()
}
