package gvm

import "fmt"

/*178 (0XB2)*/
func GETSTATIC(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	index := m.code[f.pc+1] << 8 | m.code[f.pc+2]
	field := c.constantPool[index].resolve().(*RuntimeConstantFieldrefInfo).field
	f.push(c.staticFields[field.index])
}

/*179 (0XB3)*/
func PUTSTATIC(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	index := m.code[f.pc+1] << 8 | m.code[f.pc+2]
	value := f.pop()
	field := c.constantPool[index].resolve().(*RuntimeConstantFieldrefInfo).field
	c.staticFields[field.index] = value
}

/*180 (0XB4)*/
func GETFIELD(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	index := bytes2uint16(m.code[f.pc+1:f.pc+3])
	objectref := f.pop().(*t_object)

	f.push(f.getField(objectref, index))
}

/*181 (0XB5)*/
func PUTFIELD(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	index := bytes2uint16(m.code[f.pc+1:f.pc+3])
	value := f.pop()
	objectref := f.pop().(*t_object)

	f.putField(objectref, index, value)
}

/*182 (0XB6)*/
func INVOKEVIRTUAL(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	index := uint16((m.code[f.pc+1] << 8) | m.code[f.pc+2])

	t.invokeVitrualMethod(index)
}

/*183 (0XB7)*/
func INVOKESPECIAL(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	index := uint16((m.code[f.pc+1] << 8) | m.code[f.pc+2])

	t.invokeSpecialMethod(index)
}

/*184 (0XB8)*/
func INVOKESTATIC(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	index := uint16((m.code[f.pc+1] << 8) | m.code[f.pc+2])

	t.invokeStaticMethod(index)
}

/*185 (0XB9)*/
func INVOKEINTERFACE(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*186 (0XBA)*/
func INVOKEDYNAMIC(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*187 (0XBB)*/
func NEW(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	index := bytes2uint16(m.code[f.pc+1:f.pc+3])
	class := c.constantPool[index].resolve().(*RuntimeConstantClassInfo).class
	objectreference := class.newObject()
	f.push(objectreference)
}

/*188 (0XBC)*/
func NEWARRAY(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	atype := uint8(m.code[f.pc+1])
	count := f.pop().(t_int)
	jreference := &t_array{atype: atype, length: count}
	jreference.elements = make([]t_any, count)
	f.push(jreference)
}

/*189 (0XBD)*/
func ANEWARRAY(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	index := bytes2uint16(m.code[f.pc+1:f.pc+3])
	count := f.pop().(t_int)

	classInfo := c.constantPool[index].resolve().(*RuntimeConstantClassInfo)
	jreference := newArray(classInfo.class, count)
	f.push(jreference)
}

/*190 (0XBE)*/
func ARRAYLENGTH(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	f.push(f.pop().(*t_array).length)
}

/*191 (0XBF)*/
func ATHROW(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*192 (0XC0)*/
func CHECKCAST(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*193 (0XC1)*/
func INSTANCEOF(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*194 (0XC2)*/
func MONITORENTER(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*195 (0XC3)*/
func MONITOREXIT(opcode uint8, f *StackFrame, t *Thread, c *JavaClass, m *JavaMethod) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}
