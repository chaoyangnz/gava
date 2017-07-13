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

	// support link and initialization
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

//func (this *Class) Link()  {
//	if this.linked {
//		return
//	}
//	this.verify()
//	this.prepare()
//
//	this.linked = true
//	// we resolve each symbolic class in a class or interface individually when it is used ("lazy" or "late" resolution)
//	// So SymbolRef all implements a method PrimitiveType resolve()
//	//for _, constant := range this.constantPool {
//	//	switch constant.(type) {
//	//	case SymbolRef:
//	//		constant.(SymbolRef).resolve()
//	//	}
//	//}
//	//this.resolve(class)
//}
//
//func (this *Class) verify() {
//	//TODO
//}
//
//// initialize static variables to default values: no need to execute code
//func (this *Class) prepare()  {
//	maxInstanceVars := 0  // include all the ancestry
//	maxStaticVars := 0
//
//	if this.superClass != nil {
//		this.superClass.Link()
//		maxInstanceVars = this.superClass.maxInstanceVars
//	}
//
//	// calculate static variables and instance variable count
//	for _, field := range this.fields {
//		if field.isStatic() {
//			field.index = maxStaticVars
//			maxStaticVars++
//		} else {
//			field.index = maxInstanceVars
//			maxInstanceVars++
//		}
//	}
//	this.maxInstanceVars = maxInstanceVars
//	this.maxStaticVars = maxStaticVars
//
//	// Initialize static variables
//	this.staticVars = make([]Value, this.maxStaticVars)
//	for _, field := range this.fields {
//		if field.isStatic() {
//			this.staticVars[field.index] = field.defaultValue()
//		}
//	}
//}

// invoke <clinit> to execute initialization code
func (this *Class) Initialize() []*Method {
	methods := []*Method{}
	if this.initialized {
		return methods
	}

	clinit := this.GetMethod("<clinit>", "()V")
	if clinit != nil {
		methods = append(methods, clinit)
	}

	this.initialized = true

	if this.superClass != nil {
		methods = append(methods, this.superClass.Initialize()...)
	}

	return methods
}

func (this *Class) IsInterface() bool  {
	return this.accessFlags & JVM_ACC_INTERFACE > 0
}

func (this *Class) IsArray() bool  {
	return string(this.name[0]) == JVM_SIGNATURE_ARRAY
}

func (this *Class) IsAssignableFrom(class *Class) bool  {
	// this is interface
	if this.IsInterface() { // check whether this is within class's interfaces list
		clazz := class
		for clazz != nil {
			interfaces := clazz.interfaces
			Info("class %s ifaces %d\n", clazz.Name(), len(interfaces))
			for _, interface0 := range interfaces {
				if interface0 == this {
					return true
				}
				interfaces = append(interfaces, interface0.interfaces...)
				Info("..interface %s ifaces %d\n", interface0.Name(), len(interfaces))
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

	object := &Object{class: this}
	object.header = ObjectHeader{hashCode: Int(fnv.New32a().Sum32())}
	object.values = make([]Value, this.maxInstanceVars)
	// Initialize instance variables
	class := this
	for class != nil {
		for _, field := range class.fields {
			if !field.isStatic() {
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

	return Reference{&Object{ObjectHeader{}, this, elements}}
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

/**
 * Find field with its name and descriptor in the class hierarchy
 * Java doesn't permit a static and instance method with same name + signature
 */
func (this *Class) FindMethod(name string, descriptor string) *Method {
	class := this
	for class != nil {
		for _, method := range class.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
		class = class.superClass
	}
	return nil
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
	index           int
}

func (this *Field) isStatic() bool {
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

type Method struct {
	accessFlags        uint16
	name               string
	descriptor         string
	class              *Class

	maxStack    uint
	maxLocals   uint
	code        []uint8               //u4 code_length
	exceptions  []ExceptionTableEntry //u2 exception_table_length
	localVars   []*LocalVariable

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


