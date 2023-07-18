package gava

import "math"

/*96 (0x60)*/
func IADD(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(value1 + value2)
}

/*97 (0x61)*/
func LADD(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Long)
	value1 := f.pop().(Long)
	f.push(value1 + value2)
}

/*98 (0x62)*/
func FADD(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Float)
	value1 := f.pop().(Float)
	f.push(value1 + value2)
}

/*99 (0x63)*/
func DADD(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Double)
	value1 := f.pop().(Double)
	f.push(value1 + value2)
}

/*100 (0x64)*/
func ISUB(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(value1 - value2)
}

/*101 (0x65)*/
func LSUB(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Long)
	value1 := f.pop().(Long)
	f.push(Long(value1 - value2))
}

/*102 (0x66)*/
func FSUB(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Float)
	value1 := f.pop().(Float)
	f.push(Float(value1 - value2))
}

/*103 (0x67)*/
func DSUB(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Double)
	value1 := f.pop().(Double)
	f.push(Double(value1 - value2))
}

/*104 (0x68)*/
func IMUL(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(Int(value1 * value2))
}

/*105 (0x69)*/
func LMUL(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Long)
	value1 := f.pop().(Long)
	f.push(Long(value1 * value2))
}

/*106 (0x6A)*/
func FMUL(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Float)
	value1 := f.pop().(Float)
	f.push(Float(value1 * value2))
}

/*107 (0x6B)*/
func DMUL(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Double)
	value1 := f.pop().(Double)
	f.push(Double(value1 * value2))
}

/*108 (0x6C)*/
func IDIV(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(Int(value1 / value2))
}

/*109 (0x6D)*/
func LDIV(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Long)
	value1 := f.pop().(Long)
	f.push(Long(value1 / value2))
}

/*110 (0x6E)*/
func FDIV(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Float)
	value1 := f.pop().(Float)
	f.push(Float(value1 / value2))
}

/*111 (0x6F)*/
func DDIV(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Double)
	value1 := f.pop().(Double)
	f.push(Double(value1 / value2))
}

/*112 (0x70)*/
func IREM(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(Int(value1 % value2))
}

/*113 (0x71)*/
func LREM(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Long)
	value1 := f.pop().(Long)
	f.push(Long(value1 % value2))
}

/*114 (0x72)*/
func FREM(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Float)
	value1 := f.pop().(Float)
	f.push(Float(math.Remainder(float64(value1), float64(value2))))
}

/*115 (0x73)*/
func DREM(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Double)
	value1 := f.pop().(Double)
	f.push(Float(math.Remainder(float64(value1), float64(value2))))
}

/*116 (0x74)*/
func INEG(t *Thread, f *Frame, c *Class, m *Method) {
	value := f.pop().(Int)
	f.push(Int(-value))
}

/*117 (0x75)*/
func LNEG(t *Thread, f *Frame, c *Class, m *Method) {
	value := f.pop().(Long)
	f.push(Long(-value))
}

/*118 (0x76)*/
func FNEG(t *Thread, f *Frame, c *Class, m *Method) {
	value := f.pop().(Float)
	f.push(Float(-value))
}

/*119 (0x77)*/
func DNEG(t *Thread, f *Frame, c *Class, m *Method) {
	value := f.pop().(Double)
	f.push(Double(-value))
}

/*120 (0x78)*/
func ISHL(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)

	shift := uint32(value2) & 0x1F // low 5 bits
	f.push(Int(value1 << shift))
}

/*121 (0x79)*/
func LSHL(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Long)

	shift := uint32(value2) & 0x3F // low 6 bits
	f.push(Long(value1 << shift))
}

/*122 (0x7A)*/
func ISHR(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)

	shift := uint32(value2) & 0x1F // low 5 bits
	f.push(Int(value1 >> shift))
}

/*123 (0x7B)*/
func LSHR(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Long)

	shift := uint32(value2) & 0x3F // low 6 bits
	f.push(Long(value1 >> shift))
}

/*124 (0x7C)*/
func IUSHR(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)

	shift := uint32(value2) & 0x1F
	f.push(Int(int32(uint32(value1) >> shift))) // low 5 bits
}

/*125 (0x7D)*/
func LUSHR(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Long)

	shift := uint32(value2) & 0x3F // // low 6 bits
	f.push(Long(int64(uint64(value1) >> shift)))
}

/*126 (0x7E)*/
func IAND(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(Int(value1 & value2))
}

/*127 (0x7F)*/
func LAND(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Long)
	value1 := f.pop().(Long)
	f.push(Long(value1 & value2))
}

/*128 (0x80)*/
func IOR(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(Int(value1 | value2))
}

/*129 (0x81)*/
func LOR(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Long)
	value1 := f.pop().(Long)
	f.push(Long(value1 | value2))
}

/*130 (0x82)*/
func IXOR(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(Int(value1 ^ value2))
}

/*131 (0x83)*/
func LXOR(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Long)
	value1 := f.pop().(Long)
	f.push(Long(value1 ^ value2))
}

/*132 (0x84)*/
func IINC(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.operandIndex8()
	const_value := f.operandConst8()

	value := f.loadVar(uint(index)).(Int)
	f.storeVar(uint(index), value+Int(const_value))
}
