package jago

import (
	"fmt"
	"math"
)

/*148 (0X94)*/
func LCMP(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*149 (0X95)*/
func FCMPL(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	value2 := f.pop().(jfloat)
	value1 := f.pop().(jfloat)

	if math.IsNaN(float64(value1)) || math.IsNaN(float64(value2)) {
		f.push(jint(-1))
		return
	}

	if value1 > value2 {
		f.push(jint(1))
	}
	if value1 == value2 {
		f.push(jint(0))
	}
	if value1 < value2 {
		f.push(jint(-1))
	}
}

/*150 (0X96)*/
func FCMPG(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	value2 := f.pop().(jfloat)
	value1 := f.pop().(jfloat)

	if math.IsNaN(float64(value1)) || math.IsNaN(float64(value2)) {
		f.push(jint(1))
		return
	}

	if value1 > value2 {
		f.push(jint(1))
	}
	if value1 == value2 {
		f.push(jint(0))
	}
	if value1 < value2 {
		f.push(jint(-1))
	}
}

/*151 (0X97)*/
func DCMPL(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*152 (0X98)*/
func DCMPG(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*153 (0X99)*/
func IFEQ(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	offset := (m.code[f.pc+1] << 8) | m.code[f.pc+2]
	value := f.pop().(jint)
	if value  == 0 {
		f.pc += int(offset)
	}
}

/*154 (0X9A)*/
func IFNE(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	offset := (m.code[f.pc+1] << 8) | m.code[f.pc+2]
	value := f.pop().(jint)
	if value != 0 {
		f.pc += int(offset)
	}
}

/*155 (0X9B)*/
func IFLT(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	offset := (m.code[f.pc+1] << 8) | m.code[f.pc+2]
	value := f.pop().(jint)
	if value < 0 {
		f.pc += int(offset)
	}
}

/*156 (0X9C)*/
func IFGE(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	offset := (m.code[f.pc+1] << 8) | m.code[f.pc+2]
	value := f.pop().(jint)
	if value >= 0 {
		f.pc += int(offset)
	}
}

/*157 (0X9D)*/
func IFGT(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	offset := (m.code[f.pc+1] << 8) | m.code[f.pc+2]
	value := f.pop().(jint)
	if value > 0 {
		f.pc += int(offset)
	}
}

/*158 (0X9E)*/
func IFLE(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	offset := (m.code[f.pc+1] << 8) | m.code[f.pc+2]
	value := f.pop().(jint)
	if value <= 0 {
		f.pc += int(offset)
	}
}

/*159 (0X9F)*/
func IF_ICMPEQ(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	offset := (m.code[f.pc+1] << 8) | m.code[f.pc+2]
	value2 := f.pop().(jint)
	value1 := f.pop().(jint)
	if value1 == value2 {
		f.pc += int(offset)
	}
}

/*160 (0XA0)*/
func IF_ICMPNE(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	offset := (m.code[f.pc+1] << 8) | m.code[f.pc+2]
	value2 := f.pop().(jint)
	value1 := f.pop().(jint)
	if value1 != value2 {
		f.pc += int(offset)
	}
}

/*161 (0XA1)*/
func IF_ICMPLT(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	offset := (m.code[f.pc+1] << 8) | m.code[f.pc+2]
	value2 := f.pop().(jint)
	value1 := f.pop().(jint)
	if value1 < value2 {
		f.pc += int(offset)
	}
}

/*162 (0XA2)*/
func IF_ICMPGE(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	offset := (m.code[f.pc+1] << 8) | m.code[f.pc+2]
	value2 := f.pop().(jint)
	value1 := f.pop().(jint)
	if value1 >= value2 {
		f.pc += int(offset)
	}
}

/*163 (0XA3)*/
func IF_ICMPGT(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	offset := (m.code[f.pc+1] << 8) | m.code[f.pc+2]
	value2 := f.pop().(jint)
	value1 := f.pop().(jint)
	if value1 > value2 {
		f.pc += int(offset)
	}
}

/*164 (0XA4)*/
func IF_ICMPLE(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	offset := (m.code[f.pc+1] << 8) | m.code[f.pc+2]
	value2 := f.pop().(jint)
	value1 := f.pop().(jint)
	if value1 <= value2 {
		f.pc += int(offset)
	}
}

/*165 (0XA5)*/
func IF_ACMPEQ(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	offset := (m.code[f.pc+1] << 8) | m.code[f.pc+2]
	value2 := f.pop().(jobject)
	value1 := f.pop().(jobject)
	if value1 == value2 {
		f.pc += int(offset)
	}
}

/*166 (0XA6)*/
func IF_ACMPNE(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	offset := (m.code[f.pc+1] << 8) | m.code[f.pc+2]
	value2 := f.pop().(jobject)
	value1 := f.pop().(jobject)
	if value1 != value2 {
		f.pc += int(offset)
	}
}
