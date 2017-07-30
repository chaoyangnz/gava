package jago


type InstructionRegistry []Instruction

func (this InstructionRegistry) GetInstruction(opcode uint8) Instruction {
	return this[opcode]
}

type Instruction struct {
	mnemonic  string
	interpret func(*Thread, *Frame, *Class, *Method)
}

func assertIntCompatible(value Value) Value {
	switch value.(type) {
	case Boolean, Byte, Short, Char, Int:
	default:
		Fatal("%s is not int compatible", value.Type().Name())
	}
	return value
}

func intCompatible(value Value) Int {
	assertIntCompatible(value)

	switch value.(type) {
	case Boolean:
		return Int(value.(Boolean))
	case Byte:
		return Int(value.(Byte))
	case Short:
		return Int(value.(Short))
	case Char:
		return Int(value.(Char))
	case Int:
		return value.(Int)
	default:
		Fatal("%s is not int compatible", value.Type().Name())
	}

	// never be here
	return 0
}


