package gvm

import "strings"

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
	|- ReferenceType (interface)
		|- *ClassType
		|- *ArrayType
	|- *VoidType
 */

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

	VOID_TYPE = &VoidType{}
)

const (
	boolean_true  = j_boolean(0x1)
	boolean_false = j_boolean(0x0)

	byte_default = j_byte(0)
	short_default = j_short(0)
	char_default = j_char(0)
	int_default = j_int(0)
	long_default = j_long(0)
	float_default = j_float(0.0)
	double_default = j_double(0.0)
	boolean_default = boolean_false
)
var (
	object_null = (*j_object)(nil)
	array_null = (*j_array)(nil)

	reference_default j_reference = nil
	object_default    *j_object   = object_null
	array_default     *j_array    = array_null
)

type Type interface {
	isReferenceType() bool
	defaultValue()    j_any
	signature()  string
	isElementType() bool // as final component in array
}

func ofType(signature string) Type {
	t, found := typeCache[signature]
	if found {
		return t
	}
	switch string(signature[0]) {
	case JVM_SIGNATURE_BYTE:    t = BYTE_TYPE //byte
	case JVM_SIGNATURE_CHAR:    t = CHAR_TYPE //char
	case JVM_SIGNATURE_DOUBLE:  t = DOUBLE_TYPE //double
	case JVM_SIGNATURE_FLOAT:   t = FLOAT_TYPE //float
	case JVM_SIGNATURE_INT:     t = INT_TYPE //int
	case JVM_SIGNATURE_LONG:    t = LONG_TYPE //long
	case JVM_SIGNATURE_SHORT:   t = SHORT_TYPE //short
	case JVM_SIGNATURE_BOOLEAN: t = BOOLEAN_TYPE //boolean

	case JVM_SIGNATURE_VOID:    t = VOID_TYPE // void

	case JVM_SIGNATURE_CLASS:   t = bootstrapClassLoader.load(signatureToClassName(signature)) //class
	case JVM_SIGNATURE_ARRAY:   t = &ArrayType{ofType(signature[1:])} //array

	case JVM_SIGNATURE_FUNC:    {
		arr := strings.Split(signature[1:], JVM_SIGNATURE_ENDFUNC)
		parametersSignature := arr[0]
		returnSignature := arr[1]

		var parametersSignatures []string

		for i := 0; i < len(parametersSignature); {
			ch := string(parametersSignature[i])
			switch ch {
			case JVM_SIGNATURE_BYTE, JVM_SIGNATURE_CHAR,  JVM_SIGNATURE_SHORT, JVM_SIGNATURE_INT,
				 JVM_SIGNATURE_LONG, JVM_SIGNATURE_FLOAT, JVM_SIGNATURE_DOUBLE, JVM_SIGNATURE_BOOLEAN:
				parametersSignatures = append(parametersSignatures, string(ch))
				i++
			case JVM_SIGNATURE_CLASS:
			Ref: for j := i+1; j < len(parametersSignature); j++ {
				switch string(parametersSignature[j]) {
				case JVM_SIGNATURE_ENDCLASS:
					parametersSignatures = append(parametersSignatures, string(parametersSignature[i:j+1]))
					i = j+1
					break Ref
				}
			}
			case JVM_SIGNATURE_ARRAY:
			Arr: for j := i+1; j < len(parametersSignature); j++ {
				switch string(parametersSignature[j]) {
				case JVM_SIGNATURE_ARRAY:
					continue
				case JVM_SIGNATURE_BYTE, JVM_SIGNATURE_CHAR,  JVM_SIGNATURE_SHORT, JVM_SIGNATURE_INT,
					 JVM_SIGNATURE_LONG, JVM_SIGNATURE_FLOAT, JVM_SIGNATURE_DOUBLE, JVM_SIGNATURE_BOOLEAN:
					parametersSignatures = append(parametersSignatures, string(parametersSignature[i:j+1]))
					i = j+1
					break Arr
				case JVM_SIGNATURE_CLASS:
					for k := j+1; j < len(parametersSignature); k++ {
						switch rune(parametersSignature[k]) {
						case ';':
							parametersSignatures = append(parametersSignatures, string(parametersSignature[i:k+1]))
							i = k+1
							break Arr
						}
					}
				}
			}
			}
		}

		parameterTypes := make([]Type, len(parametersSignatures))
		for i, parameterSignature := range parametersSignatures {
			parameterTypes[i] = ofType(parameterSignature)
		}
		returnType := ofType(returnSignature)
		t = &FunctionType{parameterTypes, returnType}
	}
	default:
		fatal("Not a valid vm type signature")
	}
	typeCache[signature] = t
	return t
}

func ofClassType(className string) *ClassType {
	return ofType(classNameToSignature(className)).(*ClassType)
}

func ofArrayType(signature string) *ArrayType  {
	return ofType(signature).(*ArrayType)
}

func classNameToSignature(className string) string {
	return JVM_SIGNATURE_CLASS + className + JVM_SIGNATURE_ENDCLASS;
}

func signatureToClassName(signature string) string  {
	if string(signature[0]) != JVM_SIGNATURE_CLASS || string(signature[len(signature)-1]) != JVM_SIGNATURE_ENDCLASS {
		fatal("Not a class signature")
	}
	return signature[1:len(signature)-1]
}

type ReferenceType interface {
	isReferenceType() bool
	signature()  string
	isElementType() bool
	defaultValue()    j_any

	isArrayType() bool
}

type ArrayType struct {
	componentType Type
}

func (this *ArrayType) isElementType() bool {
	return false
}

func (this *ArrayType) signature() string {
	return JVM_SIGNATURE_ARRAY + this.componentType.signature()
}

func (this *ArrayType) isReferenceType() bool {
	return true
}

func (this *ArrayType) isArrayType() bool {
	return true
}

func (this *ArrayType) defaultValue() j_any {
	return array_default
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

	VoidType struct {}
)

func (this *ByteType) isElementType() bool {
	return true
}

func (this *ByteType) signature() string {
	return JVM_SIGNATURE_BYTE
}

func (this *ByteType) isReferenceType() bool {
	return false
}

func (this *ByteType) defaultValue() j_any {
	return byte_default
}

func (this *ShortType) isElementType() bool {
	return true
}

func (this *ShortType) signature() string {
	return JVM_SIGNATURE_SHORT
}

func (this *ShortType) isReferenceType() bool {
	return false
}

func (this *ShortType) defaultValue() j_any {
	return short_default
}

func (this *CharType) isElementType() bool {
	return true
}

func (this *CharType) signature() string {
	return JVM_SIGNATURE_CHAR
}

func (this *CharType) isReferenceType() bool {
	return false
}

func (this *CharType) defaultValue() j_any {
	return char_default
}

func (this *IntType) isElementType() bool {
	return true
}

func (this *IntType) signature() string {
	return JVM_SIGNATURE_INT
}

func (this *IntType) isReferenceType() bool {
	return false
}

func (this *IntType) defaultValue() j_any {
	return int_default
}

func (this *LongType) isElementType() bool {
	return true
}

func (this *LongType) signature() string {
	return JVM_SIGNATURE_LONG
}

func (this *LongType) isReferenceType() bool {
	return false
}

func (this *LongType) defaultValue() j_any {
	return long_default
}

func (this *FloatType) isElementType() bool {
	return true
}

func (this *FloatType) signature() string {
	return JVM_SIGNATURE_FLOAT
}

func (this *FloatType) isReferenceType() bool {
	return false
}

func (this *FloatType) defaultValue() j_any {
	return float_default
}

func (this *DoubleType) isElementType() bool {
	return true
}

func (this *DoubleType) signature() string {
	return JVM_SIGNATURE_DOUBLE
}

func (this *DoubleType) isReferenceType() bool {
	return false
}

func (this *DoubleType) defaultValue() j_any {
	return double_default
}

func (this *BooleanType) isElementType() bool {
	return true
}

func (this *BooleanType) signature() string {
	return JVM_SIGNATURE_BOOLEAN
}

func (this *BooleanType) isReferenceType() bool {
	return false
}

func (this *BooleanType) defaultValue() j_any {
	return boolean_default
}

func (this *VoidType) isElementType() bool {
	return false
}

func (this *VoidType) signature() string {
	return JVM_SIGNATURE_VOID
}

func (this *VoidType) isReferenceType() bool {
	return false
}

func (this *VoidType) defaultValue() j_any {
	fatal("Void type has no default value")
	return nil
}

type FunctionType struct {
	parameterTypes  []Type
	returnType      Type
}

func (this *FunctionType) isElementType() bool {
	return false
}

func (this *FunctionType) signature() string {
	parameterSignatures := make([]string, len(this.parameterTypes))
	for i, parameterType := range this.parameterTypes {
		parameterSignatures[i] = parameterType.signature()
	}
	return JVM_SIGNATURE_FUNC + strings.Join(parameterSignatures, ",") + JVM_SIGNATURE_ENDFUNC
}

func (this *FunctionType) isReferenceType() bool {
	return false
}

func (this *FunctionType) defaultValue() j_any {
	fatal("Void type has no default value")
	return nil
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

func (this *ClassType) signature() string {
	return classNameToSignature(this.thisClassName)
}

func (this *ClassType) isReferenceType() bool {
	return false
}

func (this *ClassType) isArrayType() bool {
	return false
}

func (this *ClassType) defaultValue() j_any {
	return object_default
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
//		switch this.localVariables[i].signature[:1] {
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