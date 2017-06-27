package jago

import "math"

/*96 (0x60)*/
func IADD(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(value1 + value2)
}

/*97 (0x61)*/
func LADD(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value2 := f.pop().(Long)
	value1 := f.pop().(Long)
	f.push(value1 + value2)
}

/*98 (0x62)*/
func FADD(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value2 := f.pop().(Float)
	value1 := f.pop().(Float)
	f.push(value1 + value2)
}

/*99 (0x63)*/
func DADD(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value2 := f.pop().(Double)
	value1 := f.pop().(Double)
	f.push(value1 + value2)
}

/*100 (0x64)*/
func ISUB(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(value1 - value2)
}

/*101 (0x65)*/
func LSUB(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value2 := f.pop().(Long)
	value1 := f.pop().(Long)
	f.push(Long(value1 - value2))
}

/*102 (0x66)*/
func FSUB(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value2 := f.pop().(Float)
	value1 := f.pop().(Float)
	f.push(Float(value1 - value2))
}

/*103 (0x67)*/
func DSUB(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value2 := f.pop().(Double)
	value1 := f.pop().(Double)
	f.push(Double(value1 - value2))
}

/*104 (0x68)*/
func IMUL(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(Int(value1 * value2))
}

/*105 (0x69)*/
func LMUL(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value2 := f.pop().(Long)
	value1 := f.pop().(Long)
	f.push(Long(value1 * value2))
}

/*106 (0x6A)*/
func FMUL(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value2 := f.pop().(Float)
	value1 := f.pop().(Float)
	f.push(Float(value1 * value2))
}

/*107 (0x6B)*/
func DMUL(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value2 := f.pop().(Double)
	value1 := f.pop().(Double)
	f.push(Double(value1 * value2))
}

/*108 (0x6C)*/
func IDIV(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(Int(value1 / value2))
}

/*109 (0x6D)*/
func LDIV(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value2 := f.pop().(Long)
	value1 := f.pop().(Long)
	f.push(Long(value1 / value2))
}

/*110 (0x6E)*/
func FDIV(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value2 := f.pop().(Float)
	value1 := f.pop().(Float)
	f.push(Float(value1 / value2))
}

/*111 (0x6F)*/
func DDIV(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value2 := f.pop().(Double)
	value1 := f.pop().(Double)
	f.push(Double(value1 / value2))
}

/*112 (0x70)*/
func IREM(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(Int(value1 % value2))
}

/*113 (0x71)*/
func LREM(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value2 := f.pop().(Long)
	value1 := f.pop().(Long)
	f.push(Long(value1 % value2))
}

/*114 (0x72)*/
func FREM(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value2 := f.pop().(Float)
	value1 := f.pop().(Float)
	f.push(Float(math.Remainder(float64(value1), float64(value2))))
}

/*115 (0x73)*/
func DREM(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value2 := f.pop().(Double)
	value1 := f.pop().(Double)
	f.push(Float(math.Remainder(float64(value1), float64(value2))))
}

/*116 (0x74)*/
func INEG(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value := f.pop().(Int)
	f.push(Int(-value))
}

/*117 (0x75)*/
func LNEG(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value := f.pop().(Long)
	f.push(Long(-value))
}

/*118 (0x76)*/
func FNEG(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value := f.pop().(Float)
	f.push(Float(-value))
}

/*119 (0x77)*/
func DNEG(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value := f.pop().(Double)
	f.push(Double(-value))
}

/*120 (0x78)*/
func ISHL(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(Int(value1 << uint(value2 & 0x1F))) // low 5 bits
}

/*121 (0x79)*/
func LSHL(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value2 := f.pop().(Int)

	value1 := f.pop().(Long)
	f.push(Long(value1 << uint(value2 & 0x3F))) // low 6 bits
}

/*122 (0x7A)*/
func ISHR(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(Int(value1 >> uint(value2 & 0x1F))) // low 5 bits
}

/*123 (0x7B)*/
func LSHR(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value2 := f.pop().(Int)

	value1 := f.pop().(Long)
	f.push(Long(value1 >> uint(value2 & 0x3F))) // low 6 bits
}

/*124 (0x7C)*/
func IUSHR(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(Int(uint32(value1) >> uint(value2 & 0x1F))) // low 5 bits
}

/*125 (0x7D)*/
func LUSHR(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value2 := f.pop().(Int)

	value1 := f.pop().(Long)
	f.push(Long(uint64(value1) >> uint(value2 & 0x3F))) // low 6 bits
}

/*126 (0x7E)*/
func IAND(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(Int(value1 & value2))
}

/*127 (0x7F)*/
func LAND(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value2 := f.pop().(Long)
	value1 := f.pop().(Long)
	f.push(Long(value1 & value2))
}

/*128 (0x80)*/
func IOR(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(Int(value1 | value2))
}

/*129 (0x81)*/
func LOR(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value2 := f.pop().(Long)
	value1 := f.pop().(Long)
	f.push(Long(value1 | value2))
}

/*130 (0x82)*/
func IXOR(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(Int(value1 ^ value2))
}

/*131 (0x83)*/
func LXOR(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value2 := f.pop().(Long)
	value1 := f.pop().(Long)
	f.push(Long(value1 ^ value2))
}

/*132 (0x84)*/
func IINC(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	index := f.index8()
	const_value := f.const8(2)

	value := f.loadVar(uint(index)).(Int)
	f.storeVar(uint(index), value + Int(const_value))
}
