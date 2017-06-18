package jago

/*21 (0X15)*/
func ILOAD(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	index := f.index8()
	f.push(f.loadVar(uint(index)).(jint))
}

/*22 (0X16)*/
func LLOAD(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	index := f.index8()
	f.push(f.loadVar(uint(index)).(jlong))
}

/*23 (0X17)*/
func FLOAD(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	index := f.index8()
	f.push(f.loadVar(uint(index)).(jfloat))
}

/*24 (0X18)*/
func DLOAD(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	index := f.index8()
	f.push(f.loadVar(uint(index)).(jdouble))
}

/*25 (0X19)*/
func ALOAD(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	index := f.index8()
	f.push(f.loadVar(uint(index)).(Reference))
}

/*26 (0X1A)*/
func ILOAD_0(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	f.push(f.loadVar(0).(jint))
}

/*27 (0X1B)*/
func ILOAD_1(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	f.push(f.loadVar(1).(jint))
}

/*28 (0X1C)*/
func ILOAD_2(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	f.push(f.loadVar(2).(jint))
}

/*29 (0X1D)*/
func ILOAD_3(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	f.push(f.loadVar(3).(jint))
}

/*30 (0X1E)*/
func LLOAD_0(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	f.push(f.loadVar(uint(0)).(jlong))
}

/*31 (0X1F)*/
func LLOAD_1(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	f.push(f.loadVar(uint(1)).(jlong))
}

/*32 (0X20)*/
func LLOAD_2(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	f.push(f.loadVar(uint(2)).(jlong))
}

/*33 (0X21)*/
func LLOAD_3(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	f.push(f.loadVar(uint(3)).(jlong))
}

/*34 (0X22)*/
func FLOAD_0(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	f.push(f.loadVar(uint(0)).(jfloat))
}

/*35 (0X23)*/
func FLOAD_1(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	f.push(f.loadVar(uint(1)).(jfloat))
}

/*36 (0X24)*/
func FLOAD_2(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	f.push(f.loadVar(uint(2)).(jfloat))
}

/*37 (0X25)*/
func FLOAD_3(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	f.push(f.loadVar(uint(3)).(jfloat))
}

/*38 (0X26)*/
func DLOAD_0(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	f.push(f.loadVar(uint(0)).(jdouble))
}

/*39 (0X27)*/
func DLOAD_1(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	f.push(f.loadVar(uint(1)).(jdouble))
}

/*40 (0X28)*/
func DLOAD_2(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	f.push(f.loadVar(uint(2)).(jdouble))
}

/*41 (0X29)*/
func DLOAD_3(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	f.push(f.loadVar(uint(3)).(jdouble))
}

/*42 (0X2A)*/
func ALOAD_0(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	f.push(f.loadVar(0).(jobject))
}

/*43 (0X2B)*/
func ALOAD_1(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	f.push(f.loadVar(1).(jobject))
}

/*44 (0X2C)*/
func ALOAD_2(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	f.push(f.loadVar(2).(jobject))
}

/*45 (0X2D)*/
func ALOAD_3(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	f.push(f.loadVar(3))
}

/*46 (0X2E)
Run-time Exceptions

If arrayref is null, iaload throws a NullPointerException.

Otherwise, if index is not within the bounds of the array referenced by arrayref, the iaload instruction throws an ArrayIndexOutOfBoundsException.
*/
func IALOAD(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	index := f.pop().(jint)
	arrayref := f.pop().(jarray)
	if arrayref.IsNull() {
		Throw("NullPointerException", "")
	}
	if arrayref.class.componentType != INT_TYPE {
		Fatal("Not an int array")
	}
	f.push(arrayref.elements[index])
}

/*47 (0X2F)*/
func LALOAD(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	index := f.pop().(jint)
	arrayref := f.pop().(jarray)
	if arrayref.IsNull() {
		Throw("NullPointerException", "")
	}
	if arrayref.class.componentType != LONG_TYPE {
		Fatal("Not a long array")
	}
	f.push(arrayref.elements[index])
}

/*48 (0X30)*/
func FALOAD(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	index := f.pop().(jint)
	arrayref := f.pop().(jarray)
	if arrayref.IsNull() {
		Throw("NullPointerException", "")
	}
	if arrayref.class.componentType != FLOAT_TYPE {
		Fatal("Not an float array")
	}
	f.push(arrayref.elements[index])
}

/*49 (0X31)*/
func DALOAD(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	index := f.pop().(jint)
	arrayref := f.pop().(jarray)
	if arrayref.IsNull() {
		Throw("NullPointerException", "")
	}
	if arrayref.class.componentType != DOUBLE_TYPE {
		Fatal("Not an double array")
	}
	f.push(arrayref.elements[index])
}

/*50 (0X32)*/
func AALOAD(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	index := f.pop().(jint)
	arrayref := f.pop().(jarray)
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
func BALOAD(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	index := f.pop().(jint)
	arrayref := f.pop().(jarray)
	if arrayref.IsNull() {
		Throw("NullPointerException", "")
	}
	if arrayref.class.componentType != BOOLEAN_TYPE {
		Fatal("Not a boolean array")
	}
	f.push(arrayref.elements[index])
}

/*52 (0X34)*/
func CALOAD(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	index := f.pop().(jint)
	arrayref := f.pop().(jarray)
	if arrayref.IsNull() {
		Throw("NullPointerException", "")
	}
	if arrayref.class.componentType != CHAR_TYPE {
		Fatal("Not a char array")
	}
	f.push(arrayref.elements[index])
}

/*53 (0X35)*/
func SALOAD(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	index := f.pop().(jint)
	arrayref := f.pop().(jarray)
	if arrayref.IsNull() {
		Throw("NullPointerException", "")
	}
	if arrayref.class.componentType != SHORT_TYPE {
		Fatal("Not a short array")
	}
	f.push(arrayref.elements[index])
}