package gvm

import "fmt"

/*196 (0XC4)*/
func WIDE(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*197 (0XC5)*/
func MULTIANEWARRAY(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*198 (0XC6)*/
func IFNULL(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*199 (0XC7)*/
func IFNONNULL(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	offset := m.code[f.pc+1] << 8 | m.code[f.pc+2]

	value := f.pop().(Reference)
	if(value != nil) {
		f.pc += uint32(offset)
	}
}

/*200 (0XC8)*/
func GOTO_W(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*201 (0XC9)*/
func JSR_W(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}
