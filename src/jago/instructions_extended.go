package jago

import "fmt"

/*196 (0xC4)*/
func WIDE(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*197 (0xC5)*/
func MULTIANEWARRAY(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	//index := f.index16()
	//dimensions := m.code[f.pc+3]
	//
	//arrayType := c.constantPool[index].(*ClassRef).class.(*ArrayClass)
	//componentType := arrayType.componentType.(ClassType)
	//counts := make([]jint, dimensions)
	//for i := jint(dimensions-1); i >= 0; i-- {
	//	counts[i] = f.pop().(jint)
	//}
	//var arrayref *Array
	//componentArray := newArray(componentType, counts[0])
	//for j := uint8(1); j < dimensions-1; j++ {
	//	arrayref = newArray(componentType, counts[j])
	//	for k := jint(0); k < counts[j]; k++ {
	//		arrayref.elements[k] = componentArray
	//	}
	//	componentArray = arrayref
	//}
	//f.push(arrayref)
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*198 (0xC6)*/
func IFNULL(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	offset := f.offset16()

	value := f.pop().(Reference)
	if value.IsNull() {
		f.pc += int(offset)
		*jumped = true
	}
}

/*199 (0xC7)*/
func IFNONNULL(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	offset := f.offset16()

	value := f.pop().(Reference)
	if !value.IsNull() {
		f.pc += int(offset)
		*jumped = true
	}
}

/*200 (0xC8)*/
func GOTO_W(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*201 (0xC9)*/
func JSR_W(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}
