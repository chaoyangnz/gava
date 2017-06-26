package jago

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
  |- ClassType
        |- *Class
        |- *ArrayClass
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

type ClassType interface {
	Type
	ClassObject() ObjectRef
	IsAssignableFrom(ClassType) bool
}

type class_type_shared struct {
	name   string
	accessFlags         uint16
	superClassName      string
	interfaceNames      []string
	superClass          *Class
	interfaces          []*Class

	// shared
	classObject ObjectRef
	classLoader *ClassLoader
}

func (this *class_type_shared) Name() string {
	return this.name
}

func (this *class_type_shared) ClassObject() ObjectRef {
	return this.classObject
}

type Class struct {
	class_type_shared
	// these fields are only for non-array class
	constantPool        []Constant
	fields              []*Field
	methods             []*Method

	maxInstanceVars     int
	maxStaticVars       int
	staticVars          []Value

	// support link and initialization
	linked      bool
	initialized bool
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

func (this *Class) IsAssignableFrom(class ClassType) bool  {
	// this is interface
	if this.IsInterface() {
		var interfaces []*Class
		switch class.(type) {
		case *Class:      interfaces = class.(*Class).interfaces
		case *ArrayClass: interfaces = class.(*ArrayClass).interfaces
		}
		for _, interface0 := range interfaces {
			if interface0 == this {
				return true
			}
			interfaces = append(interfaces, interface0.interfaces...)
		}
	} else { // non-interface class
		clazz := class.(*Class)
		switch class.(type) {
		case *Class:
			if clazz.IsInterface() {// interface disallow
				return false
			}
			for clazz != nil {
				if clazz == this {
					return true
				}
				clazz = clazz.superClass
			}
		case *ArrayClass:
			if this.name == "java/lang/Object" {
				return true
			}
		}
	}

	return false
}

func (this *Class) NewObject() ObjectRef {
	//this.Link()

	object := &Object{class: this}
	object.instanceVars = make([]Value, this.maxInstanceVars)
	// Initialize instance variables
	class := this
	for class != nil {
		for _, field := range class.fields {
			if !field.isStatic() {
				object.instanceVars[field.index] = field.defaultValue()
			}
		}
		class = class.superClass
	}

	// verify initialization
	for _, instanceVar := range object.instanceVars {
		if instanceVar == nil {
			Fatal("Something wrong, unfinished instance variable initialization")
		}
	}

	return ObjectRef{object}
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

type ArrayClass struct {
	class_type_shared
	// these fields are nly for array class
	componentType Type // any type
	elementType   Type // must be not array type
	dimensions    int
}

func (this *ArrayClass) Descriptor() string  {
	return this.Name()
}

func (this *ArrayClass) NewArray(length Int) ArrayRef {
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
		case *Class:      elements[i] = NULL_OBJECT
		case *ArrayClass: elements[i] = NULL_ARRAY
		default:
			Fatal("Not a valid component type")
		}
	}

	return ArrayRef{&Array{this, length, elements}}
}

func (this *ArrayClass) IsAssignableFrom(class ClassType) bool  {
	switch class.(type) {
	case *ArrayClass:
		clazz := class.(*ArrayClass)
		switch this.componentType.(type) {
		case ClassType:
			switch clazz.componentType.(type) {
			case ClassType:
				return this.componentType.(ClassType).IsAssignableFrom(clazz.componentType.(ClassType))
			}
		}
	}
	return false
}

//type Interface struct {
//	class_type_shared
//}
//
//func (this *Interface) IsAssignableFrom(class ClassType) bool  {
//	var interfaces []*Interface
//	switch class.(type) {
//	case *Class:      interfaces = class.(*Class).interfaces
//	case *ArrayClass: interfaces = class.(*ArrayClass).interfaces
//	case *Interface:  interfaces = class.(*Interface).interfaces
//	}
//	for _, interface0 := range interfaces {
//		if interface0 == this {
//			return true
//		}
//		interfaces = append(interfaces, interface0.interfaces...)
//	}
//	return false
//}

type class_member_shared struct {
	accessFlags        uint16
	name               string
	descriptor         string
	class              *Class
}

type Field struct {
	class_member_shared
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
	case JVM_SIGNATURE_CLASS: t = NULL_OBJECT
	case JVM_SIGNATURE_ARRAY: t = NULL_ARRAY
	default:
		Fatal("Not a valid descriptor to get a default value")
	}

	return t
}

type Method struct {
	class_member_shared

	maxStack    uint
	maxLocals   uint
	code        []byte               //u4 code_length
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

type LocalVariable struct {
	method              *Method
	startPc             uint16
	length              uint16
	index               uint16
	name                string
	descriptor          string
}


