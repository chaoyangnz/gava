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

type ObjectHeader struct {
	hashCode Int
	class        *Class
}

type Object struct {
	header       ObjectHeader
	values      []Value
	extra       interface{} // for vm use
}

type ObjectRef interface {
	Value
	IsNull() bool
	IsEqual(reference Reference) bool
	Class() *Class
	GetExtra() interface{}
	SetExtra(extra interface{})

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
	GetExtra() interface{}
	SetExtra(extra interface{})

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
		VM_throw("java/lang/NullPointerException", "")
	}
	if this.Class().IsArray() {
		Bug("It is not an ObjectRef")
	}
}
func (this Reference) Class() *Class {
	return this.oop.header.class
}
func (this Reference) GetExtra() interface{} {
	return this.oop.extra
}
func (this Reference) SetExtra(extra interface{}) {
	this.oop.extra = extra
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
	field := objectref.header.class.FindField(name, descriptor)
	if field.IsStatic() {
		Fatal("not a instance variable")
	}
	return objectref.values[field.index]
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
	objectref.values[field.index] = value
}

func (this Reference) assertArray()  {
	if this.IsNull() {
		VM_throw("java/lang/NullPointerException", "")
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
	if index >= this.Length() {
		VM_throw("java/lang/ArrayIndexOutOfBoundsException", "%d exceeded array boundary %d", index, this.Length())
	}
	this.oop.values[index] = value
}

func (this Reference) AsObjectRef() ObjectRef {
	return this
}
func (this Reference) AsArrayRef() ArrayRef {
	return this
}

func NewObject(className string) ObjectRef {
	class := BOOTSTRAP_CLASSLOADER.CreateClass(className, TRIGGER_BY_NEW_INSTANCE)

	return class.NewObject()
}

/*
arrayClassName is the full array class, not its component type
 */
func NewArray(arrayClassName string, length Int) ArrayRef {
	class := BOOTSTRAP_CLASSLOADER.CreateClass(arrayClassName, TRIGGER_BY_NEW_INSTANCE)

	return class.NewArray(length)
}



