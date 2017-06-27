package jago

import "fmt"

/*87 (0x57)*/
func POP(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.pop()
}

/*88 (0x58)*/
func POP2(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*89 (0x59)*/
func DUP(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value := f.pop()
	f.push(value)
	f.push(value)
}

/*90 (0x5A)*/
func DUP_X1(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value1 := f.pop()
	value2 := f.pop()
	f.push(value1)
	f.push(value2)
	f.push(value1)
}

/*91 (0x5B)*/
func DUP_X2(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*92 (0x5C)*/
func DUP2(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*93 (0x5D)*/
func DUP2_X1(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*94 (0x5E)*/
func DUP2_X2(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*95 (0x5F)*/
func SWAP(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value1 := f.pop()
	value2 := f.pop()
	f.push(value1)
	f.push(value2)
}
