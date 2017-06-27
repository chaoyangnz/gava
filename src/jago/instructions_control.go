package jago

import "fmt"

/*167 (0xA7)*/
func GOTO(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	offset := f.offset16()

	f.pc += int(offset)
	*jumped = true
}

/*168 (0xA8)*/
func JSR(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*169 (0xA9)*/
func RET(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*170 (0xAA)*/
func TABLESWITCH(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*171 (0xAB)*/
func LOOKUPSWITCH(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*172 (0xAC)*/
func IRETURN(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	t.popFrame()
	// return value
	f.passReturn(t.peekFrame())
}

/*173 (0xAD)*/
func LRETURN(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	t.popFrame()
	// return value
	f.passReturn(t.peekFrame())
}

/*174 (0xAE)*/
func FRETURN(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	t.popFrame()
	// return value
	f.passReturn(t.peekFrame())
}

/*175 (0xAF)*/
func DRETURN(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	t.popFrame()
	// return value
	f.passReturn(t.peekFrame())
}

/*176 (0xB0)*/
func ARETURN(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	t.popFrame()
	// return value
	f.passReturn(t.peekFrame())
}

/*177 (0xB1)*/
func RETURN(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	t.popFrame()
}
