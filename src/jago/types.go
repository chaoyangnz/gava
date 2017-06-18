package jago

type Type interface {
	Name() string
}

type PremitiveType interface {
	Type
}

type (
	Byte struct {}
	Short struct {}
	Char struct {}
	Int struct {}
	Long struct {}
	Float struct {}
	Double struct {}
	Boolean struct {}
)
// their singletons
var (
	BYTE_TYPE = &Byte{}
	CHAR_TYPE = &Char{}
	SHORT_TYPE = &Short{}
	INT_TYPE = &Int{}
	LONG_TYPE = &Long{}
	FLOAT_TYPE = &Float{}
	DOUBLE_TYPE = &Double{}
	BOOLEAN_TYPE = &Boolean{}
)

func (this *Byte) Name() string  {return JVM_SIGNATURE_BYTE}
func (this *Short) Name() string  {return JVM_SIGNATURE_SHORT}
func (this *Char) Name() string  {return JVM_SIGNATURE_CHAR}
func (this *Int) Name() string  {return JVM_SIGNATURE_INT}
func (this *Long) Name() string  {return JVM_SIGNATURE_LONG}
func (this *Float) Name() string  {return JVM_SIGNATURE_FLOAT}
func (this *Double) Name() string  {return JVM_SIGNATURE_DOUBLE}
func (this *Boolean) Name() string  {return JVM_SIGNATURE_BOOLEAN}


type ClassType interface {
	Type
	SuperClassName() string
}

type class_type_shared struct {
	name   string
	accessFlags         uint16
	superClassName      string
	interfaceNames      []string
	superClass          *Class
	interfaces          []*Class
}

func (this *class_type_shared) Name() string {
	return this.name
}

func (this *class_type_shared) SuperClassName() string {
	return this.superClassName
}

type Class struct {
	class_type_shared
	// these fields are only for non-array class
	constantPool        []IConstant
	fields              []*Field
	methods             []*Method

	maxInstanceVars     int
	maxStaticVars       int
	staticVars          []Value

	// shared
	classObject         JavaLangClass
	classLoader         *ClassLoader

	// support link and initialization
	linked      bool
	initialized bool
}

func (this *Class) Link()  {
	if this.linked {
		return
	}
	this.verify()
	this.prepare()

	this.linked = true
	// we resolve each symbolic class in a class or interface individually when it is used ("lazy" or "late" resolution)
	// So SymbolRef all implements a method PremitiveType resolve()
	//for _, constant := range this.constantPool {
	//	switch constant.(type) {
	//	case ISymbolRef:
	//		constant.(ISymbolRef).resolve()
	//	}
	//}
	//this.resolve(class)
}

func (this *Class) verify() {
	//TODO
}

// initialize static variables to default values: no need to execute code
func (this *Class) prepare()  {
	maxInstanceVars := 0  // include all the ancestry
	maxStaticVars := 0

	if this.superClass != nil {
		this.superClass.Link()
		maxInstanceVars = this.superClass.maxInstanceVars
	}

	// calculate static variables and instance variable count
	for _, field := range this.fields {
		if field.isStatic() {
			field.index = maxStaticVars
			maxStaticVars++
		} else {
			field.index = maxInstanceVars
			maxInstanceVars++
		}
	}
	this.maxInstanceVars = maxInstanceVars
	this.maxStaticVars = maxStaticVars

	// Initialize static variables
	this.staticVars = make([]Value, this.maxStaticVars)
	for _, field := range this.fields {
		if field.isStatic() {
			this.staticVars[field.index] = field.defaultValue()
		}
	}
}

// invoke <clinit> to execute initialization code
func (this *Class) Initialize(thread *Thread)  {
	if this.initialized {
		return
	}

	class := this
	for class != nil {
		clinit := class.GetMethod("<clinit>", "()V")
		if clinit != nil {
			thread.pushFrame(NewStackFrame(clinit))
		}
		class.initialized = true

		class = class.superClass
	}
}

func (this *Class) NewObject() jobject {
	this.Link()

	object := &Object{class: this}
	object.instanceVars = make([]Value, this.maxInstanceVars)
	// Initialize instance variables
	for _, field := range this.fields {
		if !field.isStatic() {
			object.instanceVars[field.index] = field.defaultValue()
		}
	}

	return jobject{object}
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

func (this *ArrayClass) NewArray(length jint) jarray {
	elements := make([]Value, length)
	for i, _ := range elements {
		switch this.componentType.(type) {
		case *Byte:       elements[i] = jbyte(0)
		case *Short:      elements[i] = jshort(0)
		case *Char:       elements[i] = jchar(0)
		case *Int:        elements[i] = jint(0)
		case *Long:       elements[i] = jlong(0)
		case *Float:      elements[i] = jfloat(0.0)
		case *Double:     elements[i] = jlong(0.0)
		case *Boolean:    elements[i] = jboolean(0)
		case *Class:      elements[i] = NULL_OBJECT
		case *ArrayClass: elements[i] = NULL_ARRAY
		default:
			Fatal("Not a valid component type")
		}
	}

	return jarray{&Array{this, length, elements}}
}

type Interface struct {
	class_type_shared
}

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
	case JVM_SIGNATURE_BYTE: t = jbyte(0)
	case JVM_SIGNATURE_SHORT: t = jshort(0)
	case JVM_SIGNATURE_CHAR: t = jchar(0)
	case JVM_SIGNATURE_INT: t = jint(0)
	case JVM_SIGNATURE_LONG: t = jlong(0)
	case JVM_SIGNATURE_FLOAT: t = jfloat(0.0)
	case JVM_SIGNATURE_DOUBLE: t = jdouble(0.0)
	case JVM_SIGNATURE_BOOLEAN: t = jboolean(0)
	case JVM_SIGNATURE_CLASS: t = jobject{nil}
	case JVM_SIGNATURE_ARRAY: t = Reference(nil)
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


