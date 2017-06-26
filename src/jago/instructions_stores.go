package jago

/*54 (0X36)*/
func ISTORE(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	index := f.index8()
	f.storeVar(uint(index), f.pop().(Int))
}

/*55 (0X37)*/
func LSTORE(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	index := f.index8()
	f.storeVar(uint(index), f.pop().(Long))
}

/*56 (0X38)*/
func FSTORE(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	index := f.index8()
	f.storeVar(uint(index), f.pop().(Float))
}

/*57 (0X39)*/
func DSTORE(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	index := f.index8()
	f.storeVar(uint(index), f.pop().(Double))
}

/*58 (0X3A)*/
func ASTORE(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	index := f.index8()
	f.storeVar(uint(index), f.pop().(Reference))
}

/*59 (0X3B)*/
func ISTORE_0(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(0, f.pop().(Int))
}

/*60 (0X3C)*/
func ISTORE_1(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(1, f.pop().(Int))
}

/*61 (0X3D)*/
func ISTORE_2(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(2, f.pop().(Int))
}

/*62 (0X3E)*/
func ISTORE_3(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(3, f.pop().(Int))
}

/*63 (0X3F)*/
func LSTORE_0(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(0, f.pop().(Long))
}

/*64 (0X40)*/
func LSTORE_1(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(1, f.pop().(Long))
}

/*65 (0X41)*/
func LSTORE_2(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(2, f.pop().(Long))
}

/*66 (0X42)*/
func LSTORE_3(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(3, f.pop().(Long))
}

/*67 (0X43)*/
func FSTORE_0(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(0, f.pop().(Float))
}

/*68 (0X44)*/
func FSTORE_1(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(1, f.pop().(Float))
}

/*69 (0X45)*/
func FSTORE_2(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(2, f.pop().(Float))
}

/*70 (0X46)*/
func FSTORE_3(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(3, f.pop().(Float))
}

/*71 (0X47)*/
func DSTORE_0(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(0, f.pop().(Double))
}

/*72 (0X48)*/
func DSTORE_1(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(1, f.pop().(Double))
}

/*73 (0X49)*/
func DSTORE_2(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(2, f.pop().(Double))
}

/*74 (0X4A)*/
func DSTORE_3(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(3, f.pop().(Double))
}

/*75 (0X4B)*/
func ASTORE_0(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(0, f.pop().(Reference))
}

/*76 (0X4C)*/
func ASTORE_1(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(1, f.pop().(Reference))
}

/*77 (0X4D)*/
func ASTORE_2(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(2, f.pop().(Reference))
}

/*78 (0X4E)*/
func ASTORE_3(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.storeVar(3, f.pop().(Reference))
}

/*79 (0X4F)*/
func IASTORE(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value := f.pop().(Int)
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)
	arrayref.elements[index] = value
	//TODO check component type and boundary
}

/*80 (0X50)*/
func LASTORE(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value := f.pop().(Long)
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)
	arrayref.elements[index] = value
	//TODO check component type and boundary
}

/*81 (0X51)*/
func FASTORE(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value := f.pop().(Float)
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)
	arrayref.elements[index] = value
	//TODO check component type and boundary
}

/*82 (0X52)*/
func DASTORE(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value := f.pop().(Double)
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)
	arrayref.elements[index] = value
	//TODO check component type and boundary
}

/*83 (0X53)*/
func AASTORE(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value := f.pop().(ArrayRef)
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)
	arrayref.elements[index] = value
	//TODO check component type, boundary and subtypes
}

/*84 (0X54)*/
func BASTORE(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value := f.pop().(Int)
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)
	arrayref.elements[index] = Byte(value)
	//TODO check component type and boundary
}

/*85 (0X55)*/
func CASTORE(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value := f.pop().(Int)
	index := f.pop().(Int)


	arrayref := f.pop().(ArrayRef)
	arrayref.elements[index] = Char(value)
	//TODO check component type and boundary
}

/*86 (0X56)*/
func SASTORE(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value := f.pop().(Int)
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)
	arrayref.elements[index] = Short(value)
	//TODO check component type and boundary
}
