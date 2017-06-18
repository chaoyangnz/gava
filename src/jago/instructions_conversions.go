package jago

import "fmt"

/*133 (0X85)*/
func I2L(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	value := f.pop().(jint)
	f.push(jlong(value))
}

/*134 (0X86)*/
func I2F(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	value := f.pop().(jint)
	f.push(jfloat(value))
}

/*135 (0X87)*/
func I2D(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	value := f.pop().(jint)
	f.push(jdouble(value))
}

/*136 (0X88)*/
func L2I(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*137 (0X89)*/
func L2F(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*138 (0X8A)*/
func L2D(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*139 (0X8B)*/
func F2I(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*140 (0X8C)*/
func F2L(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*141 (0X8D)*/
func F2D(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*142 (0X8E)*/
func D2I(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*143 (0X8F)*/
func D2L(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*144 (0X90)*/
func D2F(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*145 (0X91)*/
func I2B(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*146 (0X92)*/
func I2C(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*147 (0X93)*/
func I2S(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}