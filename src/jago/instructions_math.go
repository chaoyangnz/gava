package jago

import "math"

/*96 (0x60)*/
func IADD(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(value1 + value2)
	f.nextPc()
}

/*97 (0x61)*/
func LADD(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Long)
	value1 := f.pop().(Long)
	f.push(value1 + value2)
	f.nextPc()
}

/*98 (0x62)*/
func FADD(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Float)
	value1 := f.pop().(Float)
	f.push(value1 + value2)
	f.nextPc()
}

/*99 (0x63)*/
func DADD(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Double)
	value1 := f.pop().(Double)
	f.push(value1 + value2)
	f.nextPc()
}

/*100 (0x64)*/
func ISUB(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(value1 - value2)
	f.nextPc()
}

/*101 (0x65)*/
func LSUB(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Long)
	value1 := f.pop().(Long)
	f.push(Long(value1 - value2))
	f.nextPc()
}

/*102 (0x66)*/
func FSUB(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Float)
	value1 := f.pop().(Float)
	f.push(Float(value1 - value2))
	f.nextPc()
}

/*103 (0x67)*/
func DSUB(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Double)
	value1 := f.pop().(Double)
	f.push(Double(value1 - value2))
	f.nextPc()
}

/*104 (0x68)*/
func IMUL(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(Int(value1 * value2))
	f.nextPc()
}

/*105 (0x69)*/
func LMUL(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Long)
	value1 := f.pop().(Long)
	f.push(Long(value1 * value2))
	f.nextPc()
}

/*106 (0x6A)*/
func FMUL(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Float)
	value1 := f.pop().(Float)
	f.push(Float(value1 * value2))
	f.nextPc()
}

/*107 (0x6B)*/
func DMUL(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Double)
	value1 := f.pop().(Double)
	f.push(Double(value1 * value2))
	f.nextPc()
}

/*108 (0x6C)*/
func IDIV(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(Int(value1 / value2))
	f.nextPc()
}

/*109 (0x6D)*/
func LDIV(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Long)
	value1 := f.pop().(Long)
	f.push(Long(value1 / value2))
	f.nextPc()
}

/*110 (0x6E)*/
func FDIV(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Float)
	value1 := f.pop().(Float)
	f.push(Float(value1 / value2))
	f.nextPc()
}

/*111 (0x6F)*/
func DDIV(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Double)
	value1 := f.pop().(Double)
	f.push(Double(value1 / value2))
	f.nextPc()
}

/*112 (0x70)*/
func IREM(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(Int(value1 % value2))
	f.nextPc()
}

/*113 (0x71)*/
func LREM(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Long)
	value1 := f.pop().(Long)
	f.push(Long(value1 % value2))
	f.nextPc()
}

/*114 (0x72)*/
func FREM(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Float)
	value1 := f.pop().(Float)
	f.push(Float(math.Remainder(float64(value1), float64(value2))))
	f.nextPc()
}

/*115 (0x73)*/
func DREM(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Double)
	value1 := f.pop().(Double)
	f.push(Float(math.Remainder(float64(value1), float64(value2))))
	f.nextPc()
}

/*116 (0x74)*/
func INEG(t *Thread, f *Frame, c *Class, m *Method) {
	value := f.pop().(Int)
	f.push(Int(-value))
	f.nextPc()
}

/*117 (0x75)*/
func LNEG(t *Thread, f *Frame, c *Class, m *Method) {
	value := f.pop().(Long)
	f.push(Long(-value))
	f.nextPc()
}

/*118 (0x76)*/
func FNEG(t *Thread, f *Frame, c *Class, m *Method) {
	value := f.pop().(Float)
	f.push(Float(-value))
	f.nextPc()
}

/*119 (0x77)*/
func DNEG(t *Thread, f *Frame, c *Class, m *Method) {
	value := f.pop().(Double)
	f.push(Double(-value))
	f.nextPc()
}

/*120 (0x78)*/
func ISHL(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)

	shift := uint32(value2) & 0x1F // low 5 bits
	f.push(Int(value1 << shift))
	f.nextPc()
}

/*121 (0x79)*/
func LSHL(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Long)

	shift := uint32(value2) & 0x3F // low 6 bits
	f.push(Long(value1 << shift))
	f.nextPc()
}

/*122 (0x7A)*/
func ISHR(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)

	shift := uint32(value2) & 0x1F // low 5 bits
	f.push(Int(value1 >> shift))
	f.nextPc()
}

/*123 (0x7B)*/
func LSHR(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Long)

	shift := uint32(value2) & 0x3F // low 6 bits
	f.push(Long(value1 >> shift))
	f.nextPc()
}

/*124 (0x7C)*/
func IUSHR(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)

	shift := uint32(value2) & 0x1F
	f.push(Int(int32(uint32(value1) >> shift))) // low 5 bits
	f.nextPc()
}

/*125 (0x7D)*/
func LUSHR(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Long)

	shift := uint32(value2) & 0x3F // // low 6 bits
	f.push(Long(int64(uint64(value1) >> shift)))
	f.nextPc()
}

/*126 (0x7E)*/
func IAND(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(Int(value1 & value2))
	f.nextPc()
}

/*127 (0x7F)*/
func LAND(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Long)
	value1 := f.pop().(Long)
	f.push(Long(value1 & value2))
	f.nextPc()
}

/*128 (0x80)*/
func IOR(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(Int(value1 | value2))
	f.nextPc()
}

/*129 (0x81)*/
func LOR(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Long)
	value1 := f.pop().(Long)
	f.push(Long(value1 | value2))
	f.nextPc()
}

/*130 (0x82)*/
func IXOR(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(Int(value1 ^ value2))
	f.nextPc()
}

/*131 (0x83)*/
func LXOR(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Long)
	value1 := f.pop().(Long)
	f.push(Long(value1 ^ value2))
	f.nextPc()
}

/*132 (0x84)*/
func IINC(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.operandIndex8()
	const_value := f.operandConst8()

	value := f.loadVar(uint(index)).(Int)
	f.storeVar(uint(index), value + Int(const_value))
	f.nextPc()
}
