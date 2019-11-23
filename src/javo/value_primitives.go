package javo

type (
	Byte int8
	Short int16
	Char uint16
	Int int32
	Long int64
	Float float32
	Double float64
	Boolean Byte // for boolean array, store as byte array. For other instruction, regarded as int

	ReturnAddress uint32
)

type Void struct{} // void value has not type
func (this Void) Type() Type { return nil }

var (
	TRUE  = Boolean(1)
	FALSE = Boolean(0)

	VOID = Void{}
)

func (this Byte) Type() Type          { return BYTE_TYPE }
func (this Short) Type() Type         { return SHORT_TYPE }
func (this Char) Type() Type          { return CHAR_TYPE }
func (this Int) Type() Type           { return INT_TYPE }
func (this Long) Type() Type          { return LONG_TYPE }
func (this Float) Type() Type         { return FLOAT_TYPE }
func (this Double) Type() Type        { return DOUBLE_TYPE }
func (this Boolean) Type() Type       { return BOOLEAN_TYPE }
func (this ReturnAddress) Type() Type { return RETURN_ADDRESS_TYPE }

func (this Boolean) IsTrue() bool {
	return this == TRUE
}

func (this Boolean) IsFalse() bool {
	return this == FALSE
}

func (this Boolean) ToInt() Int {
	return Int(int32(this))
}

func (this Int) ToBoolean() Boolean {
	return Boolean(this)
}
