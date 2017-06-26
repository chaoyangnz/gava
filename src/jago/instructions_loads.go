package jago

/*21 (0X15)*/
func ILOAD(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	index := f.index8()
	f.push(f.loadVar(uint(index)).(Int))
}

/*22 (0X16)*/
func LLOAD(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	index := f.index8()
	f.push(f.loadVar(uint(index)).(Long))
}

/*23 (0X17)*/
func FLOAD(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	index := f.index8()
	f.push(f.loadVar(uint(index)).(Float))
}

/*24 (0X18)*/
func DLOAD(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	index := f.index8()
	f.push(f.loadVar(uint(index)).(Double))
}

/*25 (0X19)*/
func ALOAD(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	index := f.index8()
	f.push(f.loadVar(uint(index)).(Reference))
}

/*26 (0X1A)*/
func ILOAD_0(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(f.loadVar(0).(Int))
}

/*27 (0X1B)*/
func ILOAD_1(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(f.loadVar(1).(Int))
}

/*28 (0X1C)*/
func ILOAD_2(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(f.loadVar(2).(Int))
}

/*29 (0X1D)*/
func ILOAD_3(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(f.loadVar(3).(Int))
}

/*30 (0X1E)*/
func LLOAD_0(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(f.loadVar(uint(0)).(Long))
}

/*31 (0X1F)*/
func LLOAD_1(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(f.loadVar(uint(1)).(Long))
}

/*32 (0X20)*/
func LLOAD_2(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(f.loadVar(uint(2)).(Long))
}

/*33 (0X21)*/
func LLOAD_3(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(f.loadVar(uint(3)).(Long))
}

/*34 (0X22)*/
func FLOAD_0(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(f.loadVar(uint(0)).(Float))
}

/*35 (0X23)*/
func FLOAD_1(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(f.loadVar(uint(1)).(Float))
}

/*36 (0X24)*/
func FLOAD_2(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(f.loadVar(uint(2)).(Float))
}

/*37 (0X25)*/
func FLOAD_3(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(f.loadVar(uint(3)).(Float))
}

/*38 (0X26)*/
func DLOAD_0(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(f.loadVar(uint(0)).(Double))
}

/*39 (0X27)*/
func DLOAD_1(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(f.loadVar(uint(1)).(Double))
}

/*40 (0X28)*/
func DLOAD_2(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(f.loadVar(uint(2)).(Double))
}

/*41 (0X29)*/
func DLOAD_3(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(f.loadVar(uint(3)).(Double))
}

/*42 (0X2A)*/
func ALOAD_0(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(f.loadVar(0).(Reference))
}

/*43 (0X2B)*/
func ALOAD_1(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(f.loadVar(1).(Reference))
}

/*44 (0X2C)*/
func ALOAD_2(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(f.loadVar(2).(Reference))
}

/*45 (0X2D)*/
func ALOAD_3(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(f.loadVar(3).(Reference))
}

/*46 (0X2E)
Run-time Exceptions

If arrayref is null, iaload throws a NullPointerException.

Otherwise, if index is not within the bounds of the array referenced by arrayref, the iaload instruction throws an ArrayIndexOutOfBoundsException.
*/
func IALOAD(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)
	if arrayref.IsNull() {
		Throw("NullPointerException", "")
	}
	if arrayref.class.componentType != INT_TYPE {
		Fatal("Not an int array")
	}
	f.push(arrayref.elements[index])
}

/*47 (0X2F)*/
func LALOAD(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)
	if arrayref.IsNull() {
		Throw("NullPointerException", "")
	}
	if arrayref.class.componentType != LONG_TYPE {
		Fatal("Not a long array")
	}
	f.push(arrayref.elements[index])
}

/*48 (0X30)*/
func FALOAD(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)
	if arrayref.IsNull() {
		Throw("NullPointerException", "")
	}
	if arrayref.class.componentType != FLOAT_TYPE {
		Fatal("Not an float array")
	}
	f.push(arrayref.elements[index])
}

/*49 (0X31)*/
func DALOAD(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)
	if arrayref.IsNull() {
		Throw("NullPointerException", "")
	}
	if arrayref.class.componentType != DOUBLE_TYPE {
		Fatal("Not an double array")
	}
	f.push(arrayref.elements[index])
}

/*50 (0X32)*/
func AALOAD(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)
	if arrayref.IsNull() {
		Throw("NullPointerException", "")
	}
	_, ok := arrayref.class.componentType.(ClassType)
	if !ok {
		Fatal("Not an reference array")
	}
	f.push(arrayref.elements[index])
}

/*51 (0X33)*/
func BALOAD(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)
	if arrayref.IsNull() {
		Throw("NullPointerException", "")
	}
	if arrayref.class.componentType != BOOLEAN_TYPE {
		Fatal("Not a boolean array")
	}
	f.push(arrayref.elements[index])
}

/*52 (0X34)*/
func CALOAD(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)
	if arrayref.IsNull() {
		Throw("NullPointerException", "")
	}
	if arrayref.class.componentType != CHAR_TYPE {
		Fatal("Not a char array")
	}
	//zero-extended to an int value
	value := Int(arrayref.elements[index].(Char))
	f.push(value)
}

/*53 (0X35)*/
func SALOAD(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)
	if arrayref.IsNull() {
		Throw("NullPointerException", "")
	}
	if arrayref.class.componentType != SHORT_TYPE {
		Fatal("Not a short array")
	}
	f.push(arrayref.elements[index])
}