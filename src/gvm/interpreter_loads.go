package gvm

/*21 (0X15)*/
func ILOAD(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	index := m.code[f.pc+1]
	f.push(f.loadVar(uint(index)).(t_int))
}

/*22 (0X16)*/
func LLOAD(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	index := m.code[f.pc+1]
	f.push(f.loadVar(uint(index)).(t_long))
}

/*23 (0X17)*/
func FLOAD(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	index := m.code[f.pc+1]
	f.push(f.loadVar(uint(index)).(t_float))
}

/*24 (0X18)*/
func DLOAD(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	index := m.code[f.pc+1]
	f.push(f.loadVar(uint(index)).(t_double))
}

/*25 (0X19)*/
func ALOAD(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	index := m.code[f.pc+1]
	f.push(f.loadVar(uint(index)).(t_object))
}

/*26 (0X1A)*/
func ILOAD_0(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	f.push(f.loadVar(0).(t_int))
}

/*27 (0X1B)*/
func ILOAD_1(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	f.push(f.loadVar(1).(t_int))
}

/*28 (0X1C)*/
func ILOAD_2(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	f.push(f.loadVar(2).(t_int))
}

/*29 (0X1D)*/
func ILOAD_3(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	f.push(f.loadVar(3).(t_int))
}

/*30 (0X1E)*/
func LLOAD_0(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	f.push(f.loadVar(uint(0)).(t_long))
}

/*31 (0X1F)*/
func LLOAD_1(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	f.push(f.loadVar(uint(1)).(t_long))
}

/*32 (0X20)*/
func LLOAD_2(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	f.push(f.loadVar(uint(2)).(t_long))
}

/*33 (0X21)*/
func LLOAD_3(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	f.push(f.loadVar(uint(3)).(t_long))
}

/*34 (0X22)*/
func FLOAD_0(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	f.push(f.loadVar(uint(0)).(t_float))
}

/*35 (0X23)*/
func FLOAD_1(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	f.push(f.loadVar(uint(1)).(t_float))
}

/*36 (0X24)*/
func FLOAD_2(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	f.push(f.loadVar(uint(2)).(t_float))
}

/*37 (0X25)*/
func FLOAD_3(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	f.push(f.loadVar(uint(3)).(t_float))
}

/*38 (0X26)*/
func DLOAD_0(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	f.push(f.loadVar(uint(0)).(t_double))
}

/*39 (0X27)*/
func DLOAD_1(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	f.push(f.loadVar(uint(1)).(t_double))
}

/*40 (0X28)*/
func DLOAD_2(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	f.push(f.loadVar(uint(2)).(t_double))
}

/*41 (0X29)*/
func DLOAD_3(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	f.push(f.loadVar(uint(3)).(t_double))
}

/*42 (0X2A)*/
func ALOAD_0(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	f.push(f.loadVar(0).(t_object))
}

/*43 (0X2B)*/
func ALOAD_1(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	f.push(f.loadVar(1).(t_object))
}

/*44 (0X2C)*/
func ALOAD_2(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	f.push(f.loadVar(2).(t_object))
}

/*45 (0X2D)*/
func ALOAD_3(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	f.push(f.loadVar(3).(t_object))
}

/*46 (0X2E)
Run-time Exceptions

If arrayref is null, iaload throws a NullPointerException.

Otherwise, if index is not within the bounds of the array referenced by arrayref, the iaload instruction throws an ArrayIndexOutOfBoundsException.
*/
func IALOAD(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	index := f.pop().(t_int)
	arrayref := f.pop().(t_array)
	if arrayref.atype != T_INT {
		panic("Not an int array")
	}
	f.push(arrayref.elements[index])
}

/*47 (0X2F)*/
func LALOAD(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	index := f.pop().(t_int)
	arrayref := f.pop().(t_array)
	if arrayref.atype != T_LONG {
		panic("Not an long array")
	}
	f.push(arrayref.elements[index])
}

/*48 (0X30)*/
func FALOAD(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	index := f.pop().(t_int)
	arrayref := f.pop().(t_array)
	if arrayref.atype != T_FLOAT {
		panic("Not an long array")
	}
	f.push(arrayref.elements[index])
}

/*49 (0X31)*/
func DALOAD(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	index := f.pop().(t_int)
	arrayref := f.pop().(t_array)
	if arrayref.atype != T_DOUBLE {
		panic("Not an long array")
	}
	f.push(arrayref.elements[index])
}

/*50 (0X32)*/
func AALOAD(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	index := f.pop().(t_int)
	arrayref := f.pop().(t_array)
	if arrayref.atype != 0 || arrayref.aclass == nil {
		panic("Not an object array")
	}
	f.push(arrayref.elements[index])
}

/*51 (0X33)*/
func BALOAD(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	index := f.pop().(t_int)
	arrayref := f.pop().(t_array)
	if arrayref.atype != T_BOOLEAN {
		panic("Not an boolean array")
	}
	f.push(arrayref.elements[index])
}

/*52 (0X34)*/
func CALOAD(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	index := f.pop().(t_int)
	arrayref := f.pop().(t_array)
	if arrayref.atype != T_CHAR {
		panic("Not an char array")
	}
	f.push(arrayref.elements[index])
}

/*53 (0X35)*/
func SALOAD(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	index := f.pop().(t_int)
	arrayref := f.pop().(t_array)
	if arrayref.atype != T_SHORT {
		panic("Not an short array")
	}
	f.push(arrayref.elements[index])
}