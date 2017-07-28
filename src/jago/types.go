package jago

import (
	"strings"
	"hash/fnv"
)

/*
Type system:

Type
  |- PrimitiveType
        |- *ByteType
        |- *ShortType
        |- *CharType
        |- *IntType
        |- *LongType
        |- *FloatType
        |- *DoubleType
        |- *BooleanType
  |- *ReturnAddressType
  |- *Class
 */

type Type interface {
	Name() string
	Descriptor() string
	ClassObject() JavaLangClass
}

type PrimitiveType interface {
	Type
}

type (
	ByteType struct {}
	ShortType struct {}
	CharType struct {}
	IntType struct {}
	LongType struct {}
	FloatType struct {}
	DoubleType struct {}
	BooleanType struct {}

	ReturnAddressType struct {}
)
// their singletons
var (
	BYTE_TYPE = &ByteType{}
	CHAR_TYPE = &CharType{}
	SHORT_TYPE = &ShortType{}
	INT_TYPE = &IntType{}
	LONG_TYPE = &LongType{}
	FLOAT_TYPE = &FloatType{}
	DOUBLE_TYPE = &DoubleType{}
	BOOLEAN_TYPE = &BooleanType{}

	RETURN_ADDRESS_TYPE = &ReturnAddressType{}
)

func (this *ByteType) Name() string    {return JVM_SIGNATURE_BYTE}
func (this *ShortType) Name() string   {return JVM_SIGNATURE_SHORT}
func (this *CharType) Name() string    {return JVM_SIGNATURE_CHAR}
func (this *IntType) Name() string     {return JVM_SIGNATURE_INT}
func (this *LongType) Name() string    {return JVM_SIGNATURE_LONG}
func (this *FloatType) Name() string   {return JVM_SIGNATURE_FLOAT}
func (this *DoubleType) Name() string  {return JVM_SIGNATURE_DOUBLE}
func (this *BooleanType) Name() string {return JVM_SIGNATURE_BOOLEAN}
func (this *ReturnAddressType) Name() string {return "->"}
func (this *ByteType) Descriptor() string    {return JVM_SIGNATURE_BYTE}
func (this *ShortType) Descriptor() string   {return JVM_SIGNATURE_SHORT}
func (this *CharType) Descriptor() string    {return JVM_SIGNATURE_CHAR}
func (this *IntType) Descriptor() string     {return JVM_SIGNATURE_INT}
func (this *LongType) Descriptor() string    {return JVM_SIGNATURE_LONG}
func (this *FloatType) Descriptor() string   {return JVM_SIGNATURE_FLOAT}
func (this *DoubleType) Descriptor() string  {return JVM_SIGNATURE_DOUBLE}
func (this *BooleanType) Descriptor() string {return JVM_SIGNATURE_BOOLEAN}
func (this *ReturnAddressType) Descriptor() string {return "->"}
func (this *ByteType) ClassObject() JavaLangClass    {return NewJavaLangClass(this)}
func (this *ShortType) ClassObject() JavaLangClass    {return NewJavaLangClass(this)}
func (this *CharType) ClassObject() JavaLangClass    {return NewJavaLangClass(this)}
func (this *IntType) ClassObject() JavaLangClass    {return NewJavaLangClass(this)}
func (this *LongType) ClassObject() JavaLangClass    {return NewJavaLangClass(this)}
func (this *FloatType) ClassObject() JavaLangClass    {return NewJavaLangClass(this)}
func (this *DoubleType) ClassObject() JavaLangClass    {return NewJavaLangClass(this)}
func (this *BooleanType) ClassObject() JavaLangClass    {return NewJavaLangClass(this)}
func (this *ReturnAddressType) ClassObject() JavaLangClass    {return NewJavaLangClass(this)}

//type ClassType interface {
//	Type
//	ClassObject() JavaLangClass
//	IsAssignableFrom(ClassType) bool
//	FindMethod(name string, descriptor string) *Method
//}

type Class struct {
	// shared
	name   string
	accessFlags         uint16
	superClassName      string
	interfaceNames      []string
	superClass          *Class
	interfaces          []*Class

	classObject         JavaLangClass
	classLoader         *ClassLoader

	// ---- these fields are only for non-array class ----
	constantPool        []Constant
	fields              []*Field
	methods             []*Method

	maxInstanceVars     int
	maxStaticVars       int
	staticVars          []Value

	sourceFile          string

	// support link and initialization
	defined     bool
	linked      bool
	initialized bool

	// ---- these fields are nly for array class -------
	componentType Type // any type
	elementType   Type // must be not array type
	dimensions    int
}

func (this *Class) Name() string {
	return this.name
}

func (this *Class) ClassObject() JavaLangClass {
	return this.classObject
}

func (this *Class) Descriptor() string  {
	return JVM_SIGNATURE_CLASS + this.Name() + JVM_SIGNATURE_ENDCLASS
}



func (this *Class) IsInterface() bool  {
	return this.accessFlags & JVM_ACC_INTERFACE > 0
}

func (this *Class) IsArray() bool  {
	return string(this.name[0]) == JVM_SIGNATURE_ARRAY
}

func (this *Class) IsAssignableFrom(class *Class) bool  {
	if this == class {
		return true
	}
	// this is interface
	if this.IsInterface() { // check whether this is within class's interfaces list
		clazz := class
		for clazz != nil {
			interfaces := clazz.interfaces
			for _, interface0 := range interfaces {
				if interface0 == this {
					return true
				}
				interfaces = append(interfaces, interface0.interfaces...)
			}
			clazz = clazz.superClass
		}

	} else if this.IsArray() {
		if class.IsArray() {
			thisComponentType, ok1 := this.componentType.(*Class)
			classComponentType, ok2 := class.componentType.(*Class)
			if ok1 && ok2 { // covariant
				return thisComponentType.IsAssignableFrom(classComponentType)
			}
		}
	} else { // non-array class
		if class.IsArray() {
			if class.superClass == this {
				return true
			}
		} else {
			clazz := class
			if clazz.IsInterface() {// interface disallow
				return false
			}
			for clazz != nil {
				if clazz == this {
					return true
				}
				clazz = clazz.superClass
			}
		}
	}

	return false
}

func (this *Class) NewObject() ObjectRef {
	//this.Link()

	object := &Object{}
	object.header = ObjectHeader{class: this, hashCode: Int(fnv.New32a().Sum32())}
	object.values = make([]Value, this.maxInstanceVars)
	// Initialize instance variables
	class := this
	for class != nil {
		for _, field := range class.fields {
			if !field.IsStatic() {
				object.values[field.index] = field.defaultValue()
			}
		}
		class = class.superClass
	}

	// verify initialization
	for _, instanceVar := range object.values {
		if instanceVar == nil {
			Fatal("Something wrong, unfinished instance variable initialization")
		}
	}

	return Reference{object}
}

func (this *Class) NewArray(length Int) ArrayRef {
	elements := make([]Value, length)
	for i, _ := range elements {
		switch this.componentType.(type) {
		case *ByteType:       elements[i] = Byte(0)
		case *ShortType:      elements[i] = Short(0)
		case *CharType:       elements[i] = Char(0)
		case *IntType:        elements[i] = Int(0)
		case *LongType:       elements[i] = Long(0)
		case *FloatType:      elements[i] = Float(0.0)
		case *DoubleType:     elements[i] = Long(0.0)
		case *BooleanType:    elements[i] = FALSE
		case *Class:          elements[i] = NULL
		default:
			Fatal("Not a valid component type")
		}
	}

	object := &Object{}
	object.header = ObjectHeader{class: this, hashCode: Int(fnv.New32a().Sum32())}
	object.values = elements

	return Reference{object}
}

/**
 * Find field with its name and descriptor in the class hierarchy
 * Java doesn't permit a static and instance field with same name + signature
 */
func (this *Class) FindField(name string, descriptor string) *Field {
	class := this
	for class != nil {
		for _, field := range class.fields {
			if field.name == name && field.descriptor == descriptor {
				return field
			}
		}
		class = class.superClass
	}
	return nil
}

// Only find its own defined method
func (this *Class) GetMethod(name string, descriptor string) *Method {
	for _, method := range this.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return nil
}

/*
<clinit>()V method
 */
func (this *Class) GetClassInitializer() *Method {
	clinit := this.GetMethod("<clinit>", "()V")
	if clinit != nil && clinit.isStatic() { // must be static
		return clinit
	}
	return nil
}

/*
<init>(..)V methods
 */
func (this *Class) GetConstructors(publicOnly bool) []*Method {
	constructors := []*Method{}
	for _,method := range this.methods {
		if method.name == "<init>" && method.returnDescriptor == JVM_SIGNATURE_VOID && !method.isStatic() { // non-static
			if publicOnly {
				if (method.accessFlags & JVM_ACC_PUBLIC) > 0 {
					constructors = append(constructors, method)
				}
			} else {
				constructors = append(constructors, method)
			}
		}
	}

	return constructors
}

func (this *Class) GetConstructor(descriptor string) *Method {
	method := this.GetMethod("<init>", descriptor)
	if method != nil && method.returnDescriptor == JVM_SIGNATURE_VOID && !method.isStatic() {
		return method
	}

	return nil
}

/**
 * Find field with its name and descriptor in the class hierarchy
 * Java doesn't permit a static and instance method with same name + signature
 */
func (this *Class) FindMethod(name string, descriptor string) *Method {
	for _, method := range this.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	if this.superClass != nil {
		method := this.superClass.FindMethod(name, descriptor)
		if method != nil {
			return method
		}
	}
	for _, iface := range this.interfaces {
		method := iface.FindMethod(name, descriptor)
		if method != nil {
			return method
		}
	}
	return nil
}

func (this *Class) GetDeclaredFields(publicOnly bool) []*Field {
	if !publicOnly {
		return this.fields
	}

	publicFields := []*Field{}
	for _, field := range this.fields {
		if (field.accessFlags & JVM_ACC_PUBLIC) > 0 {
			publicFields = append(publicFields, field)
		}
	}
	return publicFields
}

func (this *Class) inheritanceDepth() int {
	depth := 1
	for c := this.superClass; c != nil; c = c.superClass {
		depth++
	}
	return depth
}

type Field struct {
	accessFlags        uint16
	name               string
	descriptor         string
	class              *Class
	/**
	index of instanceFields or staticFields
	for instance fields, it is the global index considering superclass hierarchy
	 */
	index       int
}

func (this *Field) IsStatic() bool {
	return (this.accessFlags & JVM_ACC_STATIC) > 0
}

func  (this *Field) defaultValue() Value {
	var t Value
	switch string(this.descriptor[0]) {
	case JVM_SIGNATURE_BYTE: t = Byte(0)
	case JVM_SIGNATURE_SHORT: t = Short(0)
	case JVM_SIGNATURE_CHAR: t = Char(0)
	case JVM_SIGNATURE_INT: t = Int(0)
	case JVM_SIGNATURE_LONG: t = Long(0)
	case JVM_SIGNATURE_FLOAT: t = Float(0.0)
	case JVM_SIGNATURE_DOUBLE: t = Double(0.0)
	case JVM_SIGNATURE_BOOLEAN: t = FALSE
	case JVM_SIGNATURE_CLASS: t = NULL.AsObjectRef()
	case JVM_SIGNATURE_ARRAY: t = NULL.AsArrayRef()
	default:
		Fatal("Not a valid descriptor to get a default value")
	}

	return t
}

type ExceptionHandler struct {
	startPc     int
	endPc       int
	handlerPc   int
	catchType   int   // index of constant pool: ClassRef
}

type Method struct {
	accessFlags        uint16
	name               string
	descriptor         string
	class              *Class

	maxStack    uint
	maxLocals   uint
	code        []uint8               //u4 code_length
	exceptions  []*ExceptionHandler //u2 exception_table_length
	localVars   []*LocalVariable
	lineNumbers []*LineNumber

	parameterDescriptors  []string
	returnDescriptor string
}

func (this *Method) isStatic() bool {
	return (this.accessFlags & JVM_ACC_STATIC) > 0
}

func (this *Method) isNative() bool {
	return (this.accessFlags & JVM_ACC_NATIVE) > 0
}

func (this *Method) Signature() string  {
	return this.name + JVM_SIGNATURE_FUNC + strings.Join(this.parameterDescriptors, "") + JVM_SIGNATURE_ENDFUNC + this.returnDescriptor

}

func (this *Method) Qualifier() string  {
	return this.class.name + "." + this.Signature()
}

type LocalVariable struct {
	method              *Method
	startPc             uint16
	length              uint16
	index               uint16
	name                string
	descriptor          string
}

type LineNumber struct {
	startPc     int
	lineNumber  int
}


