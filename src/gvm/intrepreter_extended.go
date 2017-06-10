package gvm

import "fmt"

/*196 (0XC4)*/
func WIDE(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*197 (0XC5)*/
func MULTIANEWARRAY(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	index := m.code[f.pc+1] << 8 | m.code[f.pc+2]
	dimensions := m.code[f.pc+3]

	arrayType := c.constantPool[index].resolve().(*RuntimeConstantClassInfo).referenceType.(*ArrayType)
	componentType := arrayType.componentType.(ReferenceType)
	counts := make([]j_int, dimensions)
	for i := j_int(dimensions-1); i >= 0; i-- {
		counts[i] = f.pop().(j_int)
	}
	var array *j_array
	componentArray := newArray(componentType, counts[0])
	for j := uint8(1); j < dimensions-1; j++ {
		array = newArray(componentType, counts[j])
		for k := j_int(0); k < counts[j]; k++ {
			array.elements[k] = componentArray
		}
		componentArray = array
	}
	f.push(array)
}

/*198 (0XC6)*/
func IFNULL(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*199 (0XC7)*/
func IFNONNULL(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	offset := m.code[f.pc+1] << 8 | m.code[f.pc+2]

	value := f.pop().(j_reference)
	if(value != nil) {
		f.pc += uint32(offset)
	}
}

/*200 (0XC8)*/
func GOTO_W(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*201 (0XC9)*/
func JSR_W(opcode uint8, f *StackFrame, t *Thread, c *ClassType, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}
