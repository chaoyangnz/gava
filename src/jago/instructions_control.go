package jago

/*167 (0xA7)*/
func GOTO(t *Thread, f *Frame, c *Class, m *Method) {
	offset := f.operandOffset16()

	f.offsetPc(offset)
}

/*168 (0xA8)*/
func JSR(t *Thread, f *Frame, c *Class, m *Method) {
	Fatal("Not implemented for opcode %d\n", f.opcode())
}

/*169 (0xA9)*/
func RET(t *Thread, f *Frame, c *Class, m *Method) {
	/*index := */f.operandIndex8()
	// IGNORE
	f.nextPc()
}

/*170 (0xAA)*/
func TABLESWITCH(t *Thread, f *Frame, c *Class, m *Method) {
	f.operandPadding()

	EXEC_LOG.Debug(" default:")
	defaultOffset := f.operandOffset32()
	EXEC_LOG.Debug(" low-high: ")
	low := f.operandConst32()
	EXEC_LOG.Debug(" ~ ")
	high := f.operandConst32()

	offsets := make([]int32, high - low + 1)
	for n:=0; n < int(high - low + 1); n++ {
		EXEC_LOG.Debug(" %d: ", int(low) + n)
		offsets[n] = f.operandOffset32()
	}

	index := f.pop().(Int)

	if int32(index) < low || int32(index) > high {
		f.offsetPc32(defaultOffset)
	} else {
		f.offsetPc32(offsets[int32(index) - low])
	}
}

/*171 (0xAB)*/
func LOOKUPSWITCH(t *Thread, f *Frame, c *Class, m *Method) {
	f.operandPadding()

	EXEC_LOG.Debug(" default: ")
	defaultOffset := f.operandOffset32()
	EXEC_LOG.Debug(" pairs: ")
	npairs := f.operandConst32()
	EXEC_LOG.Debug("\t[")

	matches := make([]int32, npairs)
	offsets := make([]int32, npairs)
	for n :=0; n < int(npairs); n++ {

		match := f.operandConst32()
		matches[n] = match

		offset := f.operandOffset32()
		offsets[n] = offset
	}
	EXEC_LOG.Debug("\t]")

	key := f.pop().(Int)

	matched := false
	for j:= 0; j < int(npairs); j++ {
		if int32(key) == matches[j] {
			f.offsetPc32(offsets[j])
			matched = true
			break
		}
	}

	if !matched {
		f.offsetPc32(defaultOffset)
	}
}


/*172 (0xAC)*/
func IRETURN(t *Thread, f *Frame, c *Class, m *Method) {
	t.pop()
	// return value
	f.passReturn(t.current())
	f.nextPc()
}

/*173 (0xAD)*/
func LRETURN(t *Thread, f *Frame, c *Class, m *Method) {
	t.pop()
	// return value
	f.passReturn(t.current())
	f.nextPc()
}

/*174 (0xAE)*/
func FRETURN(t *Thread, f *Frame, c *Class, m *Method) {
	t.pop()
	// return value
	f.passReturn(t.current())
	f.nextPc()
}

/*175 (0xAF)*/
func DRETURN(t *Thread, f *Frame, c *Class, m *Method) {
	t.pop()
	// return value
	f.passReturn(t.current())
	f.nextPc()
}

/*176 (0xB0)*/
func ARETURN(t *Thread, f *Frame, c *Class, m *Method) {
	t.pop()
	// return value
	f.passReturn(t.current())
	f.nextPc()
}

/*177 (0xB1)*/
func RETURN(t *Thread, f *Frame, c *Class, m *Method) {
	t.pop()
	// no return
	f.nextPc()
}
