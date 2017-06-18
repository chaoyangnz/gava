package jago

import "fmt"

/*167 (0XA7)*/
func GOTO(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	offset := f.offset16()

	f.pc += int(offset)
}

/*168 (0XA8)*/
func JSR(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*169 (0XA9)*/
func RET(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*170 (0XAA)*/
func TABLESWITCH(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*171 (0XAB)*/
func LOOKUPSWITCH(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*172 (0XAC)*/
func IRETURN(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	t.popFrame()
	// return value
	f.passReturn(t.peekFrame())
}

/*173 (0XAD)*/
func LRETURN(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	t.popFrame()
	// return value
	f.passReturn(t.peekFrame())
}

/*174 (0XAE)*/
func FRETURN(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	t.popFrame()
	// return value
	f.passReturn(t.peekFrame())
}

/*175 (0XAF)*/
func DRETURN(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	t.popFrame()
	// return value
	f.passReturn(t.peekFrame())
}

/*176 (0XB0)*/
func ARETURN(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	t.popFrame()
	// return value
	f.passReturn(t.peekFrame())
}

/*177 (0XB1)*/
func RETURN(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	t.popFrame()
}
