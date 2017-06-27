package jago

/*54 (0x36)*/
func ISTORE(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	index := f.index8()
	f.storeVar(uint(index), f.pop().(Int))
}

/*55 (0x37)*/
func LSTORE(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	index := f.index8()
	f.storeVar(uint(index), f.pop().(Long))
}

/*56 (0x38)*/
func FSTORE(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	index := f.index8()
	f.storeVar(uint(index), f.pop().(Float))
}

/*57 (0x39)*/
func DSTORE(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	index := f.index8()
	f.storeVar(uint(index), f.pop().(Double))
}

/*58 (0x3A)*/
func ASTORE(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	index := f.index8()
	f.storeVar(uint(index), f.pop().(Reference))
}

/*59 (0x3B)*/
func ISTORE_0(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(0, f.pop().(Int))
}

/*60 (0x3C)*/
func ISTORE_1(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(1, f.pop().(Int))
}

/*61 (0x3D)*/
func ISTORE_2(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(2, f.pop().(Int))
}

/*62 (0x3E)*/
func ISTORE_3(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(3, f.pop().(Int))
}

/*63 (0x3F)*/
func LSTORE_0(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(0, f.pop().(Long))
}

/*64 (0x40)*/
func LSTORE_1(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(1, f.pop().(Long))
}

/*65 (0x41)*/
func LSTORE_2(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(2, f.pop().(Long))
}

/*66 (0x42)*/
func LSTORE_3(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(3, f.pop().(Long))
}

/*67 (0x43)*/
func FSTORE_0(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(0, f.pop().(Float))
}

/*68 (0x44)*/
func FSTORE_1(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(1, f.pop().(Float))
}

/*69 (0x45)*/
func FSTORE_2(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(2, f.pop().(Float))
}

/*70 (0x46)*/
func FSTORE_3(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(3, f.pop().(Float))
}

/*71 (0x47)*/
func DSTORE_0(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(0, f.pop().(Double))
}

/*72 (0x48)*/
func DSTORE_1(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(1, f.pop().(Double))
}

/*73 (0x49)*/
func DSTORE_2(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(2, f.pop().(Double))
}

/*74 (0x4A)*/
func DSTORE_3(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(3, f.pop().(Double))
}

/*75 (0x4B)*/
func ASTORE_0(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(0, f.pop().(Reference))
}

/*76 (0x4C)*/
func ASTORE_1(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(1, f.pop().(Reference))
}

/*77 (0x4D)*/
func ASTORE_2(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(2, f.pop().(Reference))
}

/*78 (0x4E)*/
func ASTORE_3(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(3, f.pop().(Reference))
}

/*79 (0x4F)*/
func IASTORE(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value := f.pop().(Int)
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)
	arrayref.elements[index] = value
	//TODO check component type and boundary
}

/*80 (0x50)*/
func LASTORE(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value := f.pop().(Long)
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)
	arrayref.elements[index] = value
	//TODO check component type and boundary
}

/*81 (0x51)*/
func FASTORE(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value := f.pop().(Float)
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)
	arrayref.elements[index] = value
	//TODO check component type and boundary
}

/*82 (0x52)*/
func DASTORE(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value := f.pop().(Double)
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)
	arrayref.elements[index] = value
	//TODO check component type and boundary
}

/*83 (0x53)*/
func AASTORE(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value := f.pop().(ArrayRef)
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)
	arrayref.elements[index] = value
	//TODO check component type, boundary and subtypes
}

/*84 (0x54)*/
func BASTORE(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value := f.pop().(Int)
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)
	arrayref.elements[index] = Byte(value)
	//TODO check component type and boundary
}

/*85 (0x55)*/
func CASTORE(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value := f.pop().(Int)
	index := f.pop().(Int)


	arrayref := f.pop().(ArrayRef)
	arrayref.elements[index] = Char(value)
	//TODO check component type and boundary
}

/*86 (0x56)*/
func SASTORE(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value := f.pop().(Int)
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)
	arrayref.elements[index] = Short(value)
	//TODO check component type and boundary
}
