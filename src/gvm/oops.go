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

// array component type
const (
	T_BOOLEAN	    = 4
	T_CHAR	        = 5
	T_FLOAT	        = 6
	T_DOUBLE	    = 7
	T_BYTE	        = 8
	T_SHORT	        = 9
	T_INT	        = 10
	T_LONG	        = 11
)







type JavaClass struct {
	constantPool    []RuntimeConstantPoolInfo
	accessFlags     uint16
	thisClass       uint16
	thisClassName   string
	superClass      uint16
	superClassName  string
	interfaces      []string
	fields          []*JavaField
	methods         []*JavaMethod
	fieldsMap       map[string]*JavaField
	methodsMap      map[string]*JavaMethod
	instanceFieldsStart uint16
	instanceFileds  []*JavaField
	staticFields    []t_any
	//attributes   []Attribute

	// bridge java world
	classLoader     *java_lang_classloader
	classObject     *java_lang_class // pointer to heap: instance of java/lang/Class
}



type JavaField struct {
	class           *JavaClass
	accessFlags     uint16
	name            string
	descriptor      string
	/**
	index of instanceFields or staticFields
	for instance fields, it is the global index considering superclass hierarchy
	 */
	index           uint16
}

func (this *JavaField)defaultValue() t_any {
	ch := this.descriptor[:1]
	var value t_any
	switch ch {
	case "B": value = t_byte(0) //byte
	case "C": value = t_char(0) //char
	case "D": value = t_double(0.0) //double
	case "F": value = t_float(0.0) //float
	case "I": value = t_int(0) //int
	case "J": value = t_long(0) //long
	case "S": value = t_short(0) //short
	case "Z": value = boolean_false //boolean
	case "L": value = (*t_object)(nil) //reference
	case "[": value = (*t_array)(nil) //array
	default:
		fatal("Not a valid vm type")
	}
	return value
}

func (this *JavaField)isStatic() bool {
	return this.accessFlags & FIELD_ACC_STATIC > 0
}

type JavaMethod struct {
	class           *JavaClass
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

func (this *JavaMethod) isStatic() bool {
	return this.accessFlags & METHOD_ACC_STATIC > 0
}

func (this *JavaMethod) isNative() bool {
	return this.accessFlags & METHOD_ACC_NATIVE > 0
}

type LocalVariable struct {
	method              *JavaMethod
	startPc             uint16
	length              uint16
	index               uint16
	name                string
	descriptor          string
}

//func (this *JavaMethod) localVariablesSize() uint {
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
func (this *JavaClass) newObject() *t_object {
	fields := make([]t_any, this.instanceFieldsStart + uint16(len(this.instanceFileds)))

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
		clazz = clazz.constantPool[clazz.superClass].resolve().(*RuntimeConstantClassInfo).class
	}

	return &t_object{
		class: this,
		fields: fields}
}

/*
All fields: static, instance, if not found, find its superclass and upper
 */
func (this *JavaClass) findField(signature string) *JavaField {
	field, found :=  this.fieldsMap[signature]
	if !found {
		field = this.constantPool[this.superClass].resolve().(*RuntimeConstantClassInfo).class.findField(signature)
	}
	return field
}


func (this *JavaClass) findMethod(signature string) *JavaMethod {
	method, found := this.methodsMap[signature]
	if !found {
		method = this.constantPool[this.superClass].resolve().(*RuntimeConstantClassInfo).class.findMethod(signature)
	}
	return method
}