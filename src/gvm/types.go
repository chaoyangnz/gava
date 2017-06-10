package gvm

const (
	METHOD_ACC_PUBLIC	    = 0x0001
	METHOD_ACC_PRIVATE	    = 0x0002
	METHOD_ACC_PROTECTED	= 0x0004
	METHOD_ACC_STATIC	    = 0x0008
	METHOD_ACC_FINAL	    = 0x0010
	METHOD_ACC_SYNCHRONIZED	= 0x0020
	METHOD_ACC_BRIDGE	    = 0x0040
	METHOD_ACC_VARARGS	    = 0x0080
	METHOD_ACC_NATIVE	    = 0x0100
	METHOD_ACC_ABSTRACT 	= 0x0400
	METHOD_ACC_STRICT	    = 0x0800
	METHOD_ACC_SYNTHETIC	= 0x1000
)

const (
	FIELD_ACC_PUBLIC	    = 0x0001
	FIELD_ACC_PRIVATE	    = 0x0002
	FIELD_ACC_PROTECTED	    = 0x0004
	FIELD_ACC_STATIC	    = 0x0008
	FIELD_ACC_FINAL	        = 0x0010
	FIELD_ACC_VOLATILE	    = 0x0040
	FIELD_ACC_TRANSIENT	    = 0x0080
	FIELD_ACC_SYNTHETIC	    = 0x1000
	FIELD_ACC_ENUM	        = 0x4000
)

const (
	CLASS_ACC_PUBLIC	    = 0x0001
	CLASS_ACC_FINAL	        = 0x0010
	CLASS_ACC_SUPER	        = 0x0020
	CLASS_ACC_INTERFACE	    = 0x0200
	CLASS_ACC_ABSTRACT	    = 0x0400
	CLASS_ACC_SYNTHETIC	    = 0x1000
	CLASS_ACC_ANNOTATION	= 0x2000
	CLASS_ACC_ENUM	        = 0x4000
)

/**
VM internal types:

Type (interface)
	|- *ByteType
	|- *ShortType
	|- *CharType
	|- *IntType
	|- *LongType
	|- *FloatType
	|- *DoubleType
	|- *BooleanType
	|- j_reference (interface)
		|- *ClassType
		|- *ArrayType

 */

type Type interface {
	isReferenceType() bool
	descriptor()  string
	isElementType() bool // as final component
}

type ReferenceType interface {
	isReferenceType() bool
	descriptor()  string
	isElementType() bool

	isArrayType() bool
}

type ArrayType struct {
	componentType Type
}

func (this *ArrayType) isElementType() bool {
	return false
}

func (this *ArrayType) descriptor() string {
	return "[" + this.componentType.descriptor()
}

func (this *ArrayType) isReferenceType() bool {
	return true
}

func (this *ArrayType) isArrayType() bool {
	return true
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
)

// Singleton types
var (
	BYTE_TYPE = &ByteType{}
	SHORT_TYPE = &ShortType{}
	INT_TYPE = &IntType{}
	CHAR_TYPE = &CharType{}
	LONG_TYPE = &LongType{}
	FLOAT_TYPE = &FloatType{}
	DOUBLE_TYPE = &DoubleType{}
	BOOLEAN_TYPE = &BooleanType{}
)

func (this *ByteType) isElementType() bool {
	return true
}

func (this *ByteType) descriptor() string {
	return "B"
}

func (this *ByteType) isReferenceType() bool {
	return false
}

func (this *ShortType) isElementType() bool {
	return true
}

func (this *ShortType) descriptor() string {
	return "S"
}

func (this *ShortType) isReferenceType() bool {
	return false
}

func (this *CharType) isElementType() bool {
	return true
}

func (this *CharType) descriptor() string {
	return "C"
}

func (this *CharType) isReferenceType() bool {
	return false
}

func (this *IntType) isElementType() bool {
	return true
}

func (this *IntType) descriptor() string {
	return "I"
}

func (this *IntType) isReferenceType() bool {
	return false
}

func (this *LongType) isElementType() bool {
	return true
}

func (this *LongType) descriptor() string {
	return "J"
}

func (this *LongType) isReferenceType() bool {
	return false
}

func (this *FloatType) isElementType() bool {
	return true
}

func (this *FloatType) descriptor() string {
	return "F"
}

func (this *FloatType) isReferenceType() bool {
	return false
}

func (this *DoubleType) isElementType() bool {
	return true
}

func (this *DoubleType) descriptor() string {
	return "D"
}

func (this *DoubleType) isReferenceType() bool {
	return false
}

func (this *BooleanType) isElementType() bool {
	return true
}

func (this *BooleanType) descriptor() string {
	return "Z"
}

func (this *BooleanType) isReferenceType() bool {
	return false
}

type ClassType struct {
	constantPool    []RuntimeConstantPoolInfo
	accessFlags     uint16
	thisClass       uint16
	thisClassName   string
	superClass      uint16
	superClassName  string
	interfaces      []string
	fields          []*Field
	methods         []*Method
	fieldsMap       map[string]*Field
	methodsMap      map[string]*Method
	instanceFieldsStart uint16
	instanceFileds  []*Field
	staticFields    []j_any
	//attributes   []Attribute

	// bridge java world
	classLoader     java_lang_classloader
	classObject     java_lang_class // pointer to heap: instance of java/lang/Class
}

func (this *ClassType) isElementType() bool {
	return true
}

func (this *ClassType) descriptor() string {
	return this.thisClassName
}

func (this *ClassType) isReferenceType() bool {
	return false
}

func (this *ClassType) isArrayType() bool {
	return false
}

type Field struct {
	class           *ClassType
	accessFlags     uint16
	name            string
	descriptor      string
	/**
	index of instanceFields or staticFields
	for instance fields, it is the global index considering superclass hierarchy
	 */
	index           uint16
}

func (this *Field)defaultValue() j_any {
	ch := this.descriptor[:1]
	var value j_any
	switch ch {
	case "B": value = j_byte(0) //byte
	case "C": value = j_char(0) //char
	case "D": value = j_double(0.0) //double
	case "F": value = j_float(0.0) //float
	case "I": value = j_int(0) //int
	case "J": value = j_long(0) //long
	case "S": value = j_short(0) //short
	case "Z": value = boolean_false //boolean
	case "L": value = (*j_object)(nil) //reference
	case "[": value = (*j_array)(nil) //array
	default:
		fatal("Not a valid vm type")
	}
	return value
}

func (this *Field)isStatic() bool {
	return this.accessFlags & FIELD_ACC_STATIC > 0
}

type Method struct {
	class           *ClassType
	accessFlags     uint16
	name            string
	descriptor      string
	maxStack        uint16
	maxLocals       uint16
	code            []uint8               //u4 code_length
	exceptions      []ExceptionTableEntry //u2 exception_table_length
	localVariables  []LocalVariable
	parameterDescriptors   []string
	returnDescriptor    string
}

func (this *Method) isStatic() bool {
	return this.accessFlags & METHOD_ACC_STATIC > 0
}

func (this *Method) isNative() bool {
	return this.accessFlags & METHOD_ACC_NATIVE > 0
}

type LocalVariable struct {
	method              *Method
	startPc             uint16
	length              uint16
	index               uint16
	name                string
	descriptor          string
}

//func (this *Method) localVariablesSize() uint {
//	sum := uint(0)
//	for i := 0; i < len(this.localVariables); i++ {
//		switch this.localVariables[i].descriptor[:1] {
//		case "B": sum += 1 //byte
//		case "C": sum += 2 //char
//		case "D": sum += 8 //double
//		case "F": sum += 4 //float
//		case "I": sum += 4 //int
//		case "J": sum += 8 //long
//		case "S": sum += 2 //short
//		case "Z": sum += 4 //boolean
//		case "L": sum += 4 //reference
//		case "[": sum += 4 //array
//		}
//	}
//	return sum
//}

/**
create a java instance: return the vm representation
 */
func (this *ClassType) newObject() *j_object {
	fields := make([]j_any, this.instanceFieldsStart + uint16(len(this.instanceFileds)))

	// initialize fields to default values
	clazz := this
	for ;; {
		instanceFields := clazz.instanceFileds
		for i := 0; i < len(instanceFields); i++ {
			fields[instanceFields[i].index] = instanceFields[i].defaultValue()
		}
		if clazz.superClass == 0 {
			break
		}
		clazz = clazz.constantPool[clazz.superClass].resolve().(*RuntimeConstantClassInfo).referenceType.(*ClassType)
	}

	return &j_object{
		class: this,
		fields: fields}
}

/*
All fields: static, instance, if not found, find its superclass and upper
 */
func (this *ClassType) findField(signature string) *Field {
	field, found :=  this.fieldsMap[signature]
	if !found {
		field = this.constantPool[this.superClass].resolve().(*RuntimeConstantClassInfo).referenceType.(*ClassType).findField(signature)
	}
	return field
}


func (this *ClassType) findMethod(signature string) *Method {
	method, found := this.methodsMap[signature]
	if !found {
		method = this.constantPool[this.superClass].resolve().(*RuntimeConstantClassInfo).referenceType.(*ClassType).findMethod(signature)
	}
	return method
}