package gvm

import "fmt"

/*178 (0XB2)*/
func GETSTATIC(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	index := m.code[f.pc+1] << 8 | m.code[f.pc+2]
	field := c.constantPool[index].resolve().(*RuntimeConstantFieldrefInfo).field
	f.push(field.class.staticFields[field.index])
}

/*179 (0XB3)*/
func PUTSTATIC(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	index := m.code[f.pc+1] << 8 | m.code[f.pc+2]
	value := f.pop()
	field := c.constantPool[index].resolve().(*RuntimeConstantFieldrefInfo).field
	field.class.staticFields[field.index] = value
}

/*180 (0XB4)*/
func GETFIELD(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	index := bytes2uint16(m.code[f.pc+1:f.pc+3])
	objectref := f.pop().(*j_object)

	f.push(f.getField(objectref, index))
}

/*181 (0XB5)*/
func PUTFIELD(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	index := bytes2uint16(m.code[f.pc+1:f.pc+3])
	value := f.pop()
	objectref := f.pop().(*j_object)

	f.putField(objectref, index, value)
}

/*182 (0XB6)*/
func INVOKEVIRTUAL(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	index := uint16((m.code[f.pc+1] << 8) | m.code[f.pc+2])

	t.invokeVitrualMethod(index)
}

/*183 (0XB7)*/
func INVOKESPECIAL(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	index := uint16((m.code[f.pc+1] << 8) | m.code[f.pc+2])

	t.invokeSpecialMethod(index)
}

/*184 (0XB8)*/
func INVOKESTATIC(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	index := uint16((m.code[f.pc+1] << 8) | m.code[f.pc+2])

	t.invokeStaticMethod(index)
}

/*185 (0XB9)*/
func INVOKEINTERFACE(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*186 (0XBA)*/
func INVOKEDYNAMIC(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*187 (0XBB)*/
func NEW(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	index := bytes2uint16(m.code[f.pc+1:f.pc+3])
	class := c.constantPool[index].resolve().(*RuntimeConstantClassInfo).referenceType.(*ClassType)
	objectreference := class.newObject()
	f.push(objectreference)
}

// array component type
const (
	T_BOOLEAN	    = 4
	T_CHAR	        = 5
	T_FLOAT	        = 6
	T_DOUBLE	    = 7
	T_BYTE	        = 8
	T_SHORT	        = 9
	T_INT	        = 10
	T_LONG	        = 11
)

/*188 (0XBC)*/
func NEWARRAY(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	atype := uint8(m.code[f.pc+1])
	var componentType Type
	switch atype {
	case T_CHAR: componentType = CHAR_TYPE
	case T_BYTE: componentType = BYTE_TYPE
	case T_SHORT: componentType = SHORT_TYPE
	case T_INT: componentType = INT_TYPE
	case T_LONG: componentType = LONG_TYPE
	case T_FLOAT: componentType = FLOAT_TYPE
	case T_DOUBLE: componentType = DOUBLE_TYPE
	case T_BOOLEAN: componentType = BOOLEAN_TYPE
	default:
		fatal("Invalid atype value")
	}
	count := f.pop().(j_int)
	jreference := &j_array{atype: componentType, length: count}
	jreference.elements = make([]j_any, count)
	f.push(jreference)
}

/*189 (0XBD)*/
func ANEWARRAY(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	index := bytes2uint16(m.code[f.pc+1:f.pc+3])
	count := f.pop().(j_int)

	var reference j_reference
	cpInfo := c.constantPool[index].resolve()
	switch cpInfo.(type) {
	case *RuntimeConstantClassInfo: reference = newArray(cpInfo.(*RuntimeConstantClassInfo).referenceType, count)
		//TODO component type can be: interface an array
	}

	f.push(reference)
}

/*190 (0XBE)*/
func ARRAYLENGTH(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	f.push(f.pop().(*j_array).length)
}

/*191 (0XBF)*/
func ATHROW(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*192 (0XC0)*/
func CHECKCAST(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*193 (0XC1)*/
func INSTANCEOF(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*194 (0XC2)*/
func MONITORENTER(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*195 (0XC3)*/
func MONITOREXIT(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}
