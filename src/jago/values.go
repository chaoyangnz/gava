package jago


/*
Value system:

<Value>
	|- Byte
	|- Short
	|- Char
	|- Int
	|- Long
	|- Float
	|- Double
	|- Boolean
	|- <ObjectRef>                      \
	    |- <JavaLangString>              \
	    |- <JavaLangClass>                Reference ( -> *Object)
	    |- <JavaLangClassLoader>         /
	|- <ArrayRef>                       /

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

type ObjectHeader struct {}

type Object struct {
	header       ObjectHeader
	class        *Class
	values      []Value
}

type ObjectRef interface {
	Value
	IsNull() bool
	IsEqual(reference Reference) bool
	Class() *Class

	GetInstanceVariable(index Int) Value
	SetInstanceVariable(index Int, value Value)
	GetInstanceVariableByName(name string, descriptor string) Value
	SetInstanceVariableByName(name string, descriptor string, value Value)
}

type ArrayRef interface {
	Value
	IsNull() bool
	IsEqual(reference Reference) bool
	Class() *Class

	Elements() []Value
	Length() Int
	GetElement(index Int) Value
	SetElement(index Int, value Value)
}

type Reference struct {
	oop *Object
}

func (this Reference) Type() Type       {
	return this.Class()
}
func (this Reference) IsNull() bool {return this.oop == nil}
func (this Reference) IsEqual(reference Reference) bool {
	if this.IsNull() && reference.IsNull() {
		return true
	}
	if !this.IsNull() && !reference.IsNull() {
		return this.oop == reference.oop
	}
	return false
}


func (this Reference) assertObject()  {
	if this.IsNull() {
		Throw("NullPointerException", "")
	}
	if this.Class().IsArray() {
		Bug("It is not an ObjectRef")
	}
}
func (this Reference) Class() *Class {
	return this.oop.class
}
func (this Reference) GetInstanceVariable(index Int) Value {
	return this.oop.values[index]
}
func (this Reference) SetInstanceVariable(index Int, value Value) {
	this.assertObject()
	this.oop.values[index] = value
}
func (this Reference) GetInstanceVariableByName(name string, descriptor string) Value {
	this.assertObject()

	objectref := this.oop
	field := objectref.class.FindField(name, descriptor)
	if field.isStatic() {
		Fatal("not a instance variable")
	}
	return objectref.values[field.index]
}
func (this Reference) SetInstanceVariableByName(name string, descriptor string, value Value) {
	this.assertObject()

	objectref := this.oop
	field := objectref.class.FindField(name, descriptor)
	if field.isStatic() {
		Fatal("not a instance variable")
	}
	objectref.values[field.index] = value
}

func (this Reference) assertArray()  {
	if this.IsNull() {
		Throw("NullPointerException", "")
	}
	if !this.Class().IsArray() {
		Bug("It is not an ArrayRef")
	}
}
func (this Reference) Elements() []Value  {
	this.assertArray()
	return this.oop.values
}
func (this Reference) Length() Int  {
	this.assertArray()
	return Int(len(this.oop.values))
}
func (this Reference) GetElement(index Int) Value {
	this.assertArray()
	return this.oop.values[index]
}
func (this Reference) SetElement(index Int, value Value) {
	this.assertArray()
	this.oop.values[index] = value
}

func (this Reference) AsObjectRef() ObjectRef {
	return this
}
func (this Reference) AsArrayRef() ArrayRef {
	return this
}

var null = Reference{nil}


type JavaLangString interface {
	ObjectRef
	toString() string
}
type JavaLangClass interface {ObjectRef}

func (this Reference) toString() string  {
	if this.IsNull() {
		Throw("NullPointerException", "")
	}
	this.assertObject()
	if this.AsObjectRef().Class().name != "java/lang/String" {
		Bug("It is not a java/lang/String")
	}

	runes := []rune{}
	chars := this.GetInstanceVariableByName("value", "[C").(ArrayRef).Elements()
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

func NewJavaLangClass() JavaLangClass {
	class := BOOTSTRAP_CLASSLOADER.CreateClass("java/lang/Class")
	object := class.NewObject()

	return object
}

func NewJavaLangString(str string) JavaLangString {
	// check string table
	if strObj, found := STRING_TABLE[str]; found {
		return strObj
	}
	class := BOOTSTRAP_CLASSLOADER.CreateClass("java/lang/String")
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
	charArrayClass := BOOTSTRAP_CLASSLOADER.CreateClass("[C")
	values := charArrayClass.NewArray(Int(len(chars)))
	for i, _ := range values.Elements() {
		values.SetElement(Int(i), chars[i])
	}
	object.SetInstanceVariableByName("value", "[C", values)

	// create hashing field?
	// TODO

	return object.(Reference)
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