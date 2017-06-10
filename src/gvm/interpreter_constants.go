package gvm

/*00 (0X00)*/
func NOP(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
}

/*01 (0X01)*/
func ACONST_NULL(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	f.push(nil)
}

/*02 (0X02)*/
func ICONST_M1(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	f.push(j_int(-1))
}

/*03 (0X03)*/
func ICONST_0(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	f.push(j_int(0))
}

/*04 (0X04)*/
func ICONST_1(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	f.push(j_int(1))
}

/*05 (0X05)*/
func ICONST_2(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	f.push(j_int(2))
}

/*06 (0X06)*/
func ICONST_3(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	f.push(j_int(3))
}

/*07 (0X07)*/
func ICONST_4(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	f.push(j_int(4))
}

/*08 (0X08)*/
func ICONST_5(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	f.push(j_int(5))
}

/*09 (0X09)*/
func LCONST_0(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	f.push(j_long(0))
}

/*10 (0X0A)*/
func LCONST_1(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	f.push(j_long(1))
}

/*11 (0X0B)*/
func FCONST_0(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	f.push(j_float(0.0))
}

/*12 (0X0C)*/
func FCONST_1(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	f.push(j_float(1.0))
}

/*13 (0X0D)*/
func FCONST_2(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	f.push(j_float(2.0))
}

/*14 (0X0E)*/
func DCONST_0(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	f.push(j_double(0.0))
}

/*15 (0X0F)*/
func DCONST_1(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	f.push(j_double(1.0))
}

/*16 (0X10)*/
func BIPUSH(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	b := j_byte(m.code[f.pc+1])
	f.push(j_int(b))
}

/*17 (0X11)*/
func SIPUSH(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	s := j_short((m.code[f.pc+1] << 8) | m.code[f.pc+2])
	f.push(j_int(s))
}

/*18 (0X12)*/
func LDC(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	index := m.code[f.pc+1]
	cpInfo := c.constantPool[index].resolve()
	switch cpInfo.(type) {
	case *RuntimeConstantIntegerInfo:
		f.push(cpInfo.(*RuntimeConstantIntegerInfo).value)
	case *RuntimeConstantFloatInfo:
		f.push(cpInfo.(*RuntimeConstantFloatInfo).value)
	case *RuntimeConstantStringInfo:
		f.push(cpInfo.(*RuntimeConstantStringInfo).value)
	default:
		panic("Not int, float or string in constant pool")
	}
}

/*19 (0X13)*/
func LDC_W(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	index := uint16((m.code[f.pc+1] << 8) | m.code[f.pc+2])
	cpInfo := c.constantPool[index].resolve()
	switch cpInfo.(type) {
	case *RuntimeConstantIntegerInfo:
		f.push(cpInfo.(*RuntimeConstantIntegerInfo).value)
	case *RuntimeConstantFloatInfo:
		f.push(cpInfo.(*RuntimeConstantFloatInfo).value)
	case *RuntimeConstantStringInfo:
		f.push(cpInfo.(*RuntimeConstantStringInfo).value)
	default:
		panic("Not int, float or string in constant pool")
	}
}

/*20 (0X14)*/
func LDC2_W(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	index := uint16((m.code[f.pc+1] << 8) | m.code[f.pc+2])
	cpInfo := c.constantPool[index].resolve()
	switch cpInfo.(type) {
	case *RuntimeConstantLongInfo:
		f.push(cpInfo.(*RuntimeConstantLongInfo).value)
	case *RuntimeConstantDoubleInfo:
		f.push(cpInfo.(*RuntimeConstantDoubleInfo).value)
	default:
		panic("Not long or double in constant pool")
	}
}