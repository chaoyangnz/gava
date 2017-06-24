package jago


/*
Value system:

Value
  |- Byte
  |- Short
  |- Char
  |- Int
  |- Long
  |- Float
  |- Double
  |- Boolean
  |- Reference ( -> *Object / *Array)
        |- ObjectRef
        |- ArrayRef


ObjectRef and ArrayRef are only reference value holding a pointer to real heap object <Object> or <Array>.
The reference itself will be never nil, but its containing pointer can be nil, which means the reference is `null` in Java.
 */
type Value interface {
	Type() Type
}

type (
	Byte int8
	Short int16
	Char uint16
	Int int32
    Long int64
    Float float32
    Double float64
	Boolean Byte  // for boolean array, store as byte array. For other instruction, regarded as int

	ReturnAddress uint32
)

var (
	TRUE = Boolean(1)
	FALSE = Boolean(0)
)

func (this Byte) Type() Type   { return BYTE_TYPE }
func (this Short) Type() Type  { return SHORT_TYPE }
func (this Char) Type() Type   { return CHAR_TYPE }
func (this Int) Type() Type    { return INT_TYPE }
func (this Long) Type() Type   { return LONG_TYPE }
func (this Float) Type() Type  { return FLOAT_TYPE }
func (this Double) Type() Type { return DOUBLE_TYPE }
func (this Boolean) Type() Type { return BOOLEAN_TYPE }
func (this ReturnAddress) Type() Type { return RETURN_ADDRESS_TYPE }

func (this Boolean) IsTrue() bool {
	return this == TRUE
}

func (this Boolean) IsFalse() bool {
	return this == FALSE
}

type Object struct {
	class        *Class
	instanceVars []Value
}

func (this *Object) getInstanceVariable(name string, descriptor string) Value {
	field := this.class.FindField(name, descriptor)
	if field.isStatic() {
		Fatal("not a instance variable")
	}
	return this.instanceVars[field.index]
}

func (this *Object) setInstanceVariable(name string, descriptor string, value Value) {
	field := this.class.FindField(name, descriptor)
	if field.isStatic() {
		Fatal("not a instance variable")
	}
	this.instanceVars[field.index] = value
}

type Array struct {
	class    *ArrayClass
	length   Int
	elements []Value
}

type (
	 Reference interface {
		 Value
		 IsNull() bool
		 IsEqual(Reference) bool
		 Class() ClassType
	}
	ObjectRef struct {*Object}
	ArrayRef struct {*Array}
	//JavaLangClass struct {ObjectRef}
	//JavaLangObject struct {ObjectRef}
	//JavaLangString struct {ObjectRef}
)

func (this ObjectRef) Type() Type       { return this.Class() }
func (this ObjectRef) Class() ClassType { return this.class }
func (this ArrayRef) Type() Type        { return this.Class() }
func (this ArrayRef) Class() ClassType  { return this.class }
func (this ObjectRef) IsNull() bool     {return this.Object == nil}
func (this ArrayRef) IsNull() bool      {return this.Array == nil}
func (this ObjectRef) IsEqual(reference Reference) bool {
	switch reference.(type) {
	case ObjectRef:
		ref := reference.(ObjectRef)
		return this.Object == ref.Object
	}
	return false
}
func (this ArrayRef) IsEqual(reference Reference) bool {
	switch reference.(type) {
	case ObjectRef:
		ref := reference.(ArrayRef)
		return this.Array == ref.Array
	}
	return false
}

var (
	NULL_OBJECT = ObjectRef{nil}
	NULL_ARRAY  = ArrayRef{nil}
)

func (this ObjectRef) toString() string  {
	runes := []rune{}
	chars := this.instanceVars[0].(ArrayRef).elements
	for i:=0; i < len(chars); i++ {
		char := chars[i].(Char)
		if char >= 0xD800 && char <= 0xDBFF {
			h := char
			if i+1 < len(chars) && chars[i+1].(Char) >= 0xDC00 && chars[i+1].(Char) <= 0xDFFF {
				l := chars[i+1].(Char)
				//1000016 + (H − D80016) × 40016 + (L − DC0016)
				codepoint := 0x1000 + (h - 0xD800) * 0x400 + (l - 0xDC00)
				runes = append(runes, rune(codepoint))
			} else {
				panic("Illegal UTF-16 string: only high surrogate")
			}
			i++
		} else if char >= 0xDC00 && char <= 0xDFFF {
			panic("Illegal UTF-16 string: only low surrogate")
		} else {
			runes = append(runes, rune(char))
		}
	}
	return string(runes)
}

func NewJavaLangClass() ObjectRef {
	class := BOOTSTRAP_CLASSLOADER.CreateClass("java/lang/Class").(*Class)
	object := class.NewObject()

	return object
}

func NewJavaLangString(str string) ObjectRef {
	// check string table
	if strObj, found := STRING_TABLE[str]; found {
		return strObj
	}
	class := BOOTSTRAP_CLASSLOADER.CreateClass("java/lang/String").(*Class)
	object := class.NewObject()

	// convert rune (int32) to Java char (UTF-16 with surrogate)
	chars := []Char{}
	for _, codepoint := range str {
		if codepoint <= 0xFFFF {
			chars = append(chars, Char(codepoint))
		} else {
			/*
				H = (S - 10000) / 400 + D800
				L = (S - 10000) % 400 + DC00
			 */
			high_surrogate := Char((uint32(codepoint) - 0x10000) / 0x400 + 0xD800)
			low_surrogate := Char((uint32(codepoint) - 0x10000) % 0x400 + 0xDC00)
			chars = append(chars, high_surrogate, low_surrogate)
		}
	}
	// create value field
	charArrayClass := BOOTSTRAP_CLASSLOADER.CreateClass("[C").(*ArrayClass)
	values := charArrayClass.NewArray(Int(len(chars)))
	for i, _ := range values.elements {
		values.elements[i] = chars[i]
	}
	object.instanceVars[0] = values

	// create hashing field?
	// TODO

	return object
}

/////////////////////// Type Conversion //////////////////////////////////
func (this Boolean) ToInt() Int {
	return Int(this)
}

func (this Boolean) ToByte() Byte {
	return Byte(this)
}

func (this Int) ToBoolean() Boolean {
	return Boolean(this)
}

func (this Byte) ToBoolean() Boolean {
	return Boolean(this)
}