package jago

import "fmt"

/*167 (0xA7)*/
func GOTO(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	offset := f.offset16()

	f.pc += int(offset)
	*jumped = true
}

/*168 (0xA8)*/
func JSR(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*169 (0xA9)*/
func RET(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*170 (0xAA)*/
func TABLESWITCH(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*171 (0xAB)*/
func LOOKUPSWITCH(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	var start int
	for i:= 0; i <=3; i++ {
		if ((f.pc+1)+i) % 4 == 0 {
			start = f.pc+1 +i
			break;
		}
	}

	defaultByte1 := m.code[start]
	defaultByte2 := m.code[start+1]
	defaultByte3 := m.code[start+2]
	defaultByte4 := m.code[start+3]
	npairs1 := m.code[start+4]
	npairs2 := m.code[start+5]
	npairs3 := m.code[start+6]
	npairs4 := m.code[start+7]

	defaultOffset := fourUBytesToInt(defaultByte1, defaultByte2, defaultByte3, defaultByte4)
	npairs := fourUBytesToInt(npairs1, npairs2, npairs3, npairs4)

	LOG.Trace("[lookupswitch] npairs: %d\n", npairs)

	matches := make([]int32, npairs)
	offsets := make([]int32, npairs)
	for n :=0; n < int(npairs); n++ {
		matchOffsetStart := start+8*(n+1)
		match1 := m.code[matchOffsetStart]
		match2 := m.code[matchOffsetStart+1]
		match3 := m.code[matchOffsetStart+2]
		match4 := m.code[matchOffsetStart+3]
		match := fourUBytesToInt(match1, match2, match3, match4)
		matches[n] = match
		offset1 := m.code[matchOffsetStart+4]
		offset2 := m.code[matchOffsetStart+5]
		offset3 := m.code[matchOffsetStart+6]
		offset4 := m.code[matchOffsetStart+7]
		offset := fourUBytesToInt(offset1, offset2, offset3, offset4)
		offsets[n] = offset

		LOG.Trace("[lookupswitch] match: %d -> offset: %d\n", match, offset)
	}
	LOG.Trace("[lookupswitch] default: %d\n", defaultOffset)

	key := f.pop().(Int)

	matched := false
	for j:= 0; j < int(npairs); j++ {
		if int32(key) == matches[j] {
			f.pc += int(offsets[j])
			matched = true
			break
		}
	}

	if !matched {
		f.pc += int(defaultOffset)
	}
	*jumped = true
}

func fourUBytesToInt(byte1 uint8, byte2 uint8, byte3 uint8, byte4 uint8) int32 {
	return int32((uint32(byte1) << 24) | (uint32(byte2) << 16) | (uint32(byte3) << 8) | uint32(byte4))
}

/*172 (0xAC)*/
func IRETURN(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	t.popFrame()
	// return value
	f.passReturn(t.peekFrame())
}

/*173 (0xAD)*/
func LRETURN(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	t.popFrame()
	// return value
	f.passReturn(t.peekFrame())
}

/*174 (0xAE)*/
func FRETURN(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	t.popFrame()
	// return value
	f.passReturn(t.peekFrame())
}

/*175 (0xAF)*/
func DRETURN(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	t.popFrame()
	// return value
	f.passReturn(t.peekFrame())
}

/*176 (0xB0)*/
func ARETURN(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	t.popFrame()
	// return value
	f.passReturn(t.peekFrame())
}

/*177 (0xB1)*/
func RETURN(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	t.popFrame()
}
