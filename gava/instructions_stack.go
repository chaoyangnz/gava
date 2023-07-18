package gava

/*87 (0x57)*/
func POP(t *Thread, f *Frame, c *Class, m *Method) {
	f.pop()
}

/*88 (0x58)*/
func POP2(t *Thread, f *Frame, c *Class, m *Method) {
	Fatal("Not implemented for opcode %d\n", f.opcode())
}

/*89 (0x59)*/
func DUP(t *Thread, f *Frame, c *Class, m *Method) {
	value := f.pop()
	f.push(value)
	f.push(value)
}

/*90 (0x5A)*/
func DUP_X1(t *Thread, f *Frame, c *Class, m *Method) {
	value1 := f.pop()
	value2 := f.pop()
	f.push(value1)
	f.push(value2)
	f.push(value1)
}

/*91 (0x5B)*/
func DUP_X2(t *Thread, f *Frame, c *Class, m *Method) {
	Fatal("Not implemented for opcode %d\n", f.opcode())
}

/*92 (0x5C)*/
func DUP2(t *Thread, f *Frame, c *Class, m *Method) {
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
}

/*93 (0x5D)*/
func DUP2_X1(t *Thread, f *Frame, c *Class, m *Method) {
	Fatal("Not implemented for opcode %d\n", f.opcode())
}

/*94 (0x5E)*/
func DUP2_X2(t *Thread, f *Frame, c *Class, m *Method) {
	Fatal("Not implemented for opcode %d\n", f.opcode())
}

/*95 (0x5F)*/
func SWAP(t *Thread, f *Frame, c *Class, m *Method) {
	value1 := f.pop()
	value2 := f.pop()
	f.push(value1)
	f.push(value2)
}
