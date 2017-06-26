package jago

/*00 (0X00)*/
func NOP(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
}

/*01 (0X01)*/
func ACONST_NULL(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(NULL_OBJECT)
}

/*02 (0X02)*/
func ICONST_M1(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(Int(-1))
}

/*03 (0X03)*/
func ICONST_0(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(Int(0))
}

/*04 (0X04)*/
func ICONST_1(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(Int(1))
}

/*05 (0X05)*/
func ICONST_2(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(Int(2))
}

/*06 (0X06)*/
func ICONST_3(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(Int(3))
}

/*07 (0X07)*/
func ICONST_4(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(Int(4))
}

/*08 (0X08)*/
func ICONST_5(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(Int(5))
}

/*09 (0X09)*/
func LCONST_0(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(Long(0))
}

/*10 (0X0A)*/
func LCONST_1(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(Long(1))
}

/*11 (0X0B)*/
func FCONST_0(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(Float(0.0))
}

/*12 (0X0C)*/
func FCONST_1(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(Float(1.0))
}

/*13 (0X0D)*/
func FCONST_2(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(Float(2.0))
}

/*14 (0X0E)*/
func DCONST_0(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(Double(0.0))
}

/*15 (0X0F)*/
func DCONST_1(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(Double(1.0))
}

/*16 (0X10)*/
func BIPUSH(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	b := Byte(m.code[f.pc+1])
	f.push(Int(b))
}

/*17 (0X11)*/
func SIPUSH(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	s := Short(f.offset16())
	f.push(Int(s))
}

/*18 (0X12)*/
func LDC(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	index := f.index8()
	cpInfo := c.constantPool[index]
	switch cpInfo.(type) {
	case *IntegerConstant:
		f.push(cpInfo.(*IntegerConstant).value)
	case *FloatConstant:
		f.push(cpInfo.(*FloatConstant).value)
	case *StringConstant:
		f.push(cpInfo.(*StringConstant).ResolvedString())
	case *ClassRef:
		f.push(cpInfo.(*ClassRef).ResolvedClass().ClassObject())
	case *MethodTypeConstant, *MethodHandleConstant:
		// TODO
	default:
		panic("Must be a run-time constant of type int or float, or a reference to a string literal, or a symbolic reference to a class, method type, or method handle")
	}
}

/*19 (0X13)*/
func LDC_W(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	index := f.index16()
	cpInfo := c.constantPool[index]
	switch cpInfo.(type) {
	case *IntegerConstant:
		f.push(cpInfo.(*IntegerConstant).value)
	case *FloatConstant:
		f.push(cpInfo.(*FloatConstant).value)
	case *StringConstant:
		f.push(cpInfo.(*StringConstant).ResolvedString())
	default:
		panic("Not int, float or string in constant pool")
	}
}

/*20 (0X14)*/
func LDC2_W(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	index := f.index16()

	cpInfo := c.constantPool[index]
	switch cpInfo.(type) {
	case *LongConstant:
		f.push(cpInfo.(*LongConstant).value)
	case *DoubleConstant:
		f.push(cpInfo.(*DoubleConstant).value)
	default:
		panic("Not long or double in constant pool")
	}
}