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
	|- <ObjectRef>                     \
	    |- <JavaLangString>             \
	    |- <JavaLangThread>              \
	    |- <JavaLangClass>                Reference ( -> *Object)
	    |- <JavaLangClassLoader>         /
	    |- ...                          /
	|- <ArrayRef>                      /

ObjectRef and ArrayRef are only reference value holding a pointer to real heap object <Object> or <Array>.
The reference itself will be never nil, but its containing pointer can be nil, which means the reference is `NULL` in Java.
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

type Void struct {} // void value has not type
func (this Void) Type() Type {return nil}

var (
	TRUE = Boolean(1)
	FALSE = Boolean(0)

	VOID = Void{}
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

func (this Boolean) ToInt() Int {
	return Int(int32(this))
}

func (this Int) ToBoolean() Boolean {
	return Boolean(this)
}

type ObjectHeader struct {
	hashCode Int
	class        *Class
	monitor      *Monitor
	extra       interface{} // for vm use
}

type Object struct {
	header ObjectHeader
	slots  []Value
}

type Ref interface {
	Value
	IsNull() bool
	IsEqual(reference Reference) bool
	Class() *Class
	GetExtra() interface{}
	SetExtra(extra interface{})
	Monitor() *Monitor
}

type ObjectRef interface {
	Ref

	GetInstanceVariable(index Int) Value
	SetInstanceVariable(index Int, value Value)
	GetInstanceVariableByName(name string, descriptor string) Value
	SetInstanceVariableByName(name string, descriptor string, value Value)
}

type ArrayRef interface {
	Ref

	ArrayElements() []Value
	ArrayLength() Int
	GetArrayElement(index Int) Value
	SetArrayElement(index Int, value Value)
}

var NULL = Reference{nil}

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
		VM.Throw("java/lang/NullPointerException", "")
	}
	if this.Class().IsArray() {
		Bug("It is not an ObjectRef")
	}
}
func (this Reference) Class() *Class {
	return this.oop.header.class
}
func (this Reference) Monitor() *Monitor {
	return this.oop.header.monitor
}
func (this Reference) GetExtra() interface{} {
	return this.oop.header.extra
}
func (this Reference) SetExtra(extra interface{}) {
	this.oop.header.extra = extra
}
func (this Reference) GetInstanceVariable(index Int) Value {
	return this.oop.slots[index]
}
func (this Reference) SetInstanceVariable(index Int, value Value) {
	this.assertObject()
	this.oop.slots[index] = value
}
func (this Reference) GetInstanceVariableByName(name string, descriptor string) Value {
	this.assertObject()

	objectref := this.oop
	field := objectref.header.class.FindField(name, descriptor)
	if field == nil {
		print("break")
	}
	if field.IsStatic() {
		Fatal("not a instance variable")
	}
	return objectref.slots[field.slot]
}
func (this Reference) SetInstanceVariableByName(name string, descriptor string, value Value) {
	this.assertObject()

	objectref := this.oop
	field := objectref.header.class.FindField(name, descriptor)
	if field == nil {
		Fatal("Field (%s %s)  cannot be found in %s", name, descriptor, objectref.header.class.Name())
	}
	if field.IsStatic() {
		Fatal("not a instance variable")
	}
	objectref.slots[field.slot] = value
}

func (this Reference) dump() {
	if !this.IsNull() {
		VM.Debug("Dump object (%s): {", this.Class().Name())
		for _, field := range this.Class().fields {
			if !field.IsStatic() {
				value := this.GetInstanceVariable(Int(field.slot))
				VM.Debug("\t%s: %v", field.name, value)
			}
		}
		VM.Debug("}\n")
	}
}

func (this Reference) assertArray()  {
	if this.IsNull() {
		VM.Throw("java/lang/NullPointerException", "")
	}
	if !this.Class().IsArray() {
		Bug("It is not an ArrayRef")
	}
}
func (this Reference) ArrayElements() []Value  {
	this.assertArray()
	return this.oop.slots
}
func (this Reference) ArrayLength() Int  {
	this.assertArray()
	return Int(len(this.oop.slots))
}
func (this Reference) GetArrayElement(index Int) Value {
	this.assertArray()
	return this.oop.slots[index]
}
func (this Reference) SetArrayElement(index Int, value Value) {
	this.assertArray()
	if index >= this.ArrayLength() {
		VM.Throw("java/lang/ArrayIndexOutOfBoundsException", "%d exceeded array boundary %d", index, this.ArrayLength())
	}
	this.oop.slots[index] = value
}

func (this Reference) AsObjectRef() ObjectRef {
	return this
}
func (this Reference) AsArrayRef() ArrayRef {
	return this
}





