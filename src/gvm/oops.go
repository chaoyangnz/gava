package gvm

type (
	java_any           interface{}

	java_byte          int8
	java_char          uint16
	java_short         int16
	java_int           int32
	java_long          int64
	java_boolean       int32
	java_float         float32
	java_double        float64
	java_object        *JavaObject
	java_array         *JavaArray

	java_int_array      []java_int
	java_char_array     []java_char

	java_string        *JavaObject
)

const (
	java_true           = java_boolean(0x1)
	java_false          = java_boolean(0x0)
	//java_null           = nil
)

type JavaObject struct {
	//header part
	class *JavaClass
	flags uint32
	locks uint32
	//fields
	fields []java_any
}

func (this *JavaObject) getField(caller *JavaClass, index uint16) java_any {
	i := caller.constantPool[index].resolve().(*RuntimeConstantFieldrefInfo).field.index
	return this.fields[i]
}

func (this *JavaObject) putField(caller *JavaClass, index uint16, value java_any) {
	i := caller.constantPool[index].resolve().(*RuntimeConstantFieldrefInfo).field.index
	this.fields[i] = value
}

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



type JavaArray struct {
	//header part
	/*
	Array Type	atype
	T_BOOLEAN	4
	T_CHAR	5
	T_FLOAT	6
	T_DOUBLE	7
	T_BYTE	8
	T_SHORT	9
	T_INT	10
	T_LONG	11
	*/
	atype   uint8
	aclass  *JavaClass
	flags   uint32
	locks   uint32
	size    uint32
	//fields
	elements []java_any
}


/*
cp_info {
    u1 tag;
    u1 info[];
}
 */
type RuntimeConstantPoolInfo interface {
	resolve() RuntimeConstantPoolInfo
}

/*
CONSTANT_Class_info {
    u1 tag;
    u2 name_index;
}
 */
type RuntimeConstantClassInfo struct {
	hostClass       *JavaClass
	nameIndex       u2
	resolved        bool
	class           *JavaClass
}

func (this *RuntimeConstantClassInfo) resolve() RuntimeConstantPoolInfo {
	if !this.resolved {
		name := this.hostClass.constantPool[this.nameIndex].resolve().(*RuntimeConstantUtf8Info).value
		clazz := classCache[name]
		if clazz == nil {
			if this == this.hostClass.constantPool[this.hostClass.thisClass] {
				clazz = this.hostClass
			} else {
				cr := NewClassReader(name)
				cf := cr.ReadAsClassFile()

				clazz = &JavaClass{}
				clazz.Load(cf)
			}
		}

		this.class = clazz
		this.resolved = true
	}
	return this
}

/*
CONSTANT_Fieldref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
 */
type RuntimeConstantFieldrefInfo struct {
	hostClass           *JavaClass
	classIndex          u2
	nameAndTypeIndex    u2
	resolved            bool
	field               *JavaField
}

func (this *RuntimeConstantFieldrefInfo) resolve() RuntimeConstantPoolInfo  {
	if !this.resolved {
		rcpClass := this.hostClass.constantPool[this.classIndex].resolve().(*RuntimeConstantClassInfo)

		nameAndType := this.hostClass.constantPool[this.nameAndTypeIndex].resolve().(*RuntimeConstantNameAndTypeInfo)
		this.field = rcpClass.class.findField(nameAndType.toString())
		this.resolved = true
	}
	return this
}

/*
CONSTANT_Methodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
 */
type RuntimeConstantMethodrefInfo struct {
	hostClass       *JavaClass
	classIndex       u2
	nameAndTypeIndex u2
	resolved         bool
	method           *JavaMethod
}

func (this *RuntimeConstantMethodrefInfo) resolve() RuntimeConstantPoolInfo  {
	if !this.resolved {
		rcpClass := this.hostClass.constantPool[this.classIndex].resolve().(*RuntimeConstantClassInfo)

		nameAndType := this.hostClass.constantPool[this.nameAndTypeIndex].resolve().(*RuntimeConstantNameAndTypeInfo)
		this.method = rcpClass.class.findMethod(nameAndType.toString())
		this.resolved = true
	}
	return this
}


/*
CONSTANT_InterfaceMethodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
 */
type RuntimeConstantInterfaceMethodrefInfo struct {
	hostClass       *JavaClass
	classIndex       u2
	nameAndTypeIndex u2
	resolved         bool
	method           *JavaMethod
}

func (this *RuntimeConstantInterfaceMethodrefInfo) resolve() RuntimeConstantPoolInfo  {
	if !this.resolved {
		//TODO

		this.resolved = true
	}
	return this
}

/*
CONSTANT_String_info {
    u1 tag;
    u2 string_index;
}
 */
type RuntimeConstantStringInfo struct {
	hostClass       *JavaClass
	stringIndex     u2
	resolved        bool
	//value           []java_char
	value           java_string
}

func NewJavaString(chars java_char_array) java_string  {
	stringClass := classCache["java/lang/String"]
	if stringClass == nil {
		classReader := NewClassReader("java/lang/String")
		stringClass = &JavaClass{}
		stringClass.Load(classReader.ReadAsClassFile())
	}
	object := stringClass.new()
	object.fields[0] = chars
	return java_string(object)
}

func JavaString2Utf8(jstring java_string) string {
	runes := []rune{}
	chars := jstring.fields[0].(java_char_array)
	for i:=0; i < len(chars); i++ {
		char := chars[i]
		if char >= 0xD800 && char <= 0xDBFF {
			h := char
			if i+1 < len(chars) && chars[i+1] >= 0xDC00 && chars[i+1] <= 0xDFFF {
				l := chars[i+1]
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

func (this *RuntimeConstantStringInfo) resolve() RuntimeConstantPoolInfo {
	if !this.resolved {
		utf8string := this.hostClass.constantPool[this.stringIndex].resolve().(*RuntimeConstantUtf8Info).value
		jstring := stringTable[utf8string]
		if jstring == nil {
			jstring = NewJavaString(string2JavaChars(utf8string))
			stringTable[utf8string] = jstring
		}
		this.value = jstring
		this.resolved = true
	}
	return this
}

/*
CONSTANT_Integer_info {
    u1 tag;
    u4 bytes;
}
 */
type RuntimeConstantIntegerInfo struct {
	hostClass       *JavaClass
	bytes           u4
	resolved        bool
	value           java_int
}


func (this *RuntimeConstantIntegerInfo) resolve() RuntimeConstantPoolInfo  {
	if !this.resolved {
		this.value = java_int(this.bytes)
		this.resolved = true
	}
	return this
}

/*
CONSTANT_Float_info {
    u1 tag;
    u4 bytes;
}
 */
type RuntimeConstantFloatInfo struct {
	hostClass       *JavaClass
	bytes           u4
	resolved        bool
	value           java_float
}

func (this *RuntimeConstantFloatInfo) resolve() RuntimeConstantPoolInfo  {
	if !this.resolved {
		this.value = java_float(this.bytes)
		this.resolved = true
	}
	return this
}

/*
CONSTANT_Long_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
 */
type RuntimeConstantLongInfo struct {
	hostClass       *JavaClass
	highBytes       u4
	lowBytes        u4
	resolved        bool
	value           java_long
}

func (this *RuntimeConstantLongInfo) resolve() RuntimeConstantPoolInfo  {
	if !this.resolved {
		this.value = java_long(this.highBytes << 8 | this.lowBytes)
		this.resolved = true
	}
	return this
}

/*
CONSTANT_Double_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
 */
type RuntimeConstantDoubleInfo struct {
	hostClass       *JavaClass
	highBytes       u4
	lowBytes        u4
	resolved        bool
	value           java_double
}


func (this *RuntimeConstantDoubleInfo) resolve() RuntimeConstantPoolInfo  {
	if !this.resolved {
		this.value = java_double(this.highBytes << 8 | this.lowBytes)
		this.resolved = true
	}
	return this
}

/*
CONSTANT_NameAndType_info {
    u1 tag;
    u2 name_index;
    u2 descriptor_index;
}
 */
type RuntimeConstantNameAndTypeInfo struct {
	hostClass       *JavaClass
	nameIndex       u2
	descriptorIndex u2
	resolved        bool
	name            string
	descriptor      string
}

func (this *RuntimeConstantNameAndTypeInfo) resolve() RuntimeConstantPoolInfo  {
	if !this.resolved {
		this.name = this.hostClass.constantPool[this.nameIndex].resolve().(*RuntimeConstantUtf8Info).value
		this.descriptor = this.hostClass.constantPool[this.descriptorIndex].resolve().(*RuntimeConstantUtf8Info).value
		this.resolved = true
	}
	return this
}

func (this *RuntimeConstantNameAndTypeInfo) toString() string  {
	return this.name + this.descriptor
}

/*
CONSTANT_Utf8_info {
    u1 tag;
    u2 length;
    u1 bytes[length];
}
 */
type RuntimeConstantUtf8Info struct {
	hostClass       *JavaClass
	length          u2
	bytes           []u1
	resolved        bool
	value           string
}

func (this *RuntimeConstantUtf8Info) resolve() RuntimeConstantPoolInfo {
	if !this.resolved {
		this.value = u2s(this.bytes)
		this.resolved = true
	}
	return this
}

/*
CONSTANT_MethodHandle_info {
    u1 tag;
    u1 reference_kind;
    u2 reference_index;
}
 */
type RuntimeConstantMethodHandleInfo struct {
	hostClass       *JavaClass
	referenceKind   u1
	referenceIndex  u2
	resolved        bool
	//TODO
}

func (this *RuntimeConstantMethodHandleInfo) resolve() RuntimeConstantPoolInfo {
	if !this.resolved {
		//TODO
	}
	return this
}

/*
CONSTANT_MethodType_info {
    u1 tag;
    u2 descriptor_index;
}
 */
type RuntimeConstantMethodTypeInfo struct {
	hostClass       *JavaClass
	descriptorIndex u2
	resolved        bool
	descriptor      string
}

func (this *RuntimeConstantMethodTypeInfo) resolve() RuntimeConstantPoolInfo {
	if !this.resolved {
		this.descriptor = this.hostClass.constantPool[this.descriptorIndex].resolve().(*RuntimeConstantUtf8Info).value
		this.resolved = true
	}
	return this
}

/*
CONSTANT_InvokeDynamic_info {
    u1 tag;
    u2 bootstrap_method_attr_index;
    u2 name_and_type_index;
}
 */
type RuntimeConstantInvokeDynamicInfo struct {
	hostClass                *JavaClass
	bootstrapMethodAttrIndex u2
	nameAndTypeIndex         u2
	resolved                 bool
	//TODO
}


func (this *RuntimeConstantInvokeDynamicInfo) resolve() RuntimeConstantPoolInfo {
	if !this.resolved {
		//TODO
	}
	return this
}


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
	staticFields    []java_any
	//attributes   []Attribute
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

func (this *JavaMethod) localVariablesSize() uint {
	sum := uint(0)
	for i := 0; i < len(this.localVariables); i++ {
		switch this.localVariables[i].descriptor[:1] {
		case "B": sum += 1 //byte
		case "C": sum += 2 //char
		case "D": sum += 8 //double
		case "F": sum += 4 //float
		case "I": sum += 4 //int
		case "J": sum += 8 //long
		case "S": sum += 2 //short
		case "Z": sum += 4 //boolean
		case "L": sum += 4 //reference
		case "[": sum += 4 //array
		}
	}
	return sum
}

var classCache = make(map[string] *JavaClass)

func (this *JavaClass) Load(classfile *ClassFile)  {
	this.constantPool = make([]RuntimeConstantPoolInfo, classfile.constantPoolCount)
	for i := u2(1); i < classfile.constantPoolCount; i++ {
		this.constantPool[i] = classfile.constantPool[i].runtime(this)
	}
	//for i := u2(1); i < classfile.constantPoolCount; i++ {
	//	runtimeConstantPoolInfo := this.constantPool[i]
	//	switch runtimeConstantPoolInfo.(type) {
	//	case *RuntimeConstantIntegerInfo, *RuntimeConstantLongInfo, *RuntimeConstantFloatInfo, *RuntimeConstantDoubleInfo,
	//		 *RuntimeConstantStringInfo, *RuntimeConstantUtf8Info, *RuntimeConstantNameAndTypeInfo, *RuntimeConstantMethodTypeInfo:
	//		runtimeConstantPoolInfo.resolve()
	//	}
	//}
	this.accessFlags = uint16(classfile.accessFlags)
	// resolve this class
	this.thisClass = uint16(classfile.thisClass)
	this.thisClassName = classfile.cpUtf8(this.constantPool[this.thisClass].resolve().(*RuntimeConstantClassInfo).nameIndex)

	// resolve super class
	this.superClass = uint16(classfile.superClass)
	if this.superClass != 0 {
		this.superClassName = classfile.cpUtf8(this.constantPool[this.superClass].resolve().(*RuntimeConstantClassInfo).nameIndex)
	}

	this.fields = make([]*JavaField, len(classfile.fields))
	this.fieldsMap = make(map[string]*JavaField)
	this.staticFields = []java_any{}
	maxInstanceFieldIndex := uint16(0)
	maxStaticFieldIndex   := uint16(0)
	if this.superClass == 0 { // jdk/lang/Object
		this.instanceFieldsStart = 0
	} else {
		superClass := this.constantPool[this.superClass].(*RuntimeConstantClassInfo).class
		this.instanceFieldsStart = superClass.instanceFieldsStart + uint16(len(superClass.instanceFileds))
	}
	for i := 0; i < len(classfile.fields); i++ {
		fieldInfo := classfile.fields[i]
		javaFiled := &JavaField{class: this,
			accessFlags: uint16(fieldInfo.accessFlags),
			name: classfile.cpUtf8(fieldInfo.nameIndex),
			descriptor: classfile.cpUtf8(fieldInfo.descriptorIndex)}
		this.fields[i] = javaFiled
		this.fieldsMap[javaFiled.name + javaFiled.descriptor] = javaFiled
		if javaFiled.isStatic() {
			javaFiled.index = maxStaticFieldIndex
			this.staticFields = append(this.staticFields, 0)
			maxStaticFieldIndex++
		} else {
			javaFiled.index = maxInstanceFieldIndex + this.instanceFieldsStart
			this.instanceFileds = append(this.instanceFileds, javaFiled)
			maxInstanceFieldIndex++
		}
	}


	this.methods = make([]*JavaMethod, len(classfile.methods))
	this.methodsMap = make(map[string]*JavaMethod)
	for i := 0; i < len(classfile.methods); i++ {
		methodInfo := &classfile.methods[i]
		javaMethod := &JavaMethod{class: this,
			accessFlags: uint16(methodInfo.accessFlags),
			name: classfile.cpUtf8(methodInfo.nameIndex),
			descriptor: classfile.cpUtf8(methodInfo.descriptorIndex)}

		javaMethod.parameterDescriptors, javaMethod.returnDescriptor = parametersAndReturn(javaMethod.descriptor)
		for j := u2(0); j < methodInfo.attributeCount; j++ {
			attributeInfo := methodInfo.attributes[j]
			switch attributeInfo.(type) {
			case *CodeAttribute:
				codeAttribute := attributeInfo.(*CodeAttribute)
				javaMethod.maxStack = uint16(codeAttribute.maxStack)
				javaMethod.maxLocals = uint16(codeAttribute.maxLocals)
				javaMethod.code = u2b(codeAttribute.code)
				for k := u2(0); k < codeAttribute.attributesCount; k++ {
					codeAttributeAttribute := codeAttribute.attributes[k]
					switch codeAttributeAttribute.(type) {
					case *LocalVariableTableAttribute:
						localVariableTableAttribute := codeAttributeAttribute.(*LocalVariableTableAttribute)
						javaMethod.localVariables = make([]LocalVariable, localVariableTableAttribute.localVariableTableLength)
						for m := u2(0); m < localVariableTableAttribute.localVariableTableLength; m++ {
							javaMethod.localVariables[m] = LocalVariable{
								javaMethod,
								uint16(localVariableTableAttribute.localVariableTable[m].startPc),
								uint16(localVariableTableAttribute.localVariableTable[m].length),
								uint16(localVariableTableAttribute.localVariableTable[m].index),
								classfile.cpUtf8(localVariableTableAttribute.localVariableTable[m].nameIndex),
								classfile.cpUtf8(localVariableTableAttribute.localVariableTable[m].descriptorIndex)}
						}
					}
				}
			}
		}
		this.methods[i] = javaMethod
		this.methodsMap[javaMethod.name + javaMethod.descriptor] = javaMethod
	}

	classCache[this.thisClassName] = this
}

func (this *JavaClass) new() java_object  {
	return java_object(&JavaObject{
		class: this,
		fields: make([]java_any, this.instanceFieldsStart + uint16(len(this.instanceFileds)))})
}

func (this *JavaClass) findField(signature string) *JavaField {
	field :=  this.fieldsMap[signature]
	if field == nil {
		field = this.constantPool[this.superClass].resolve().(*RuntimeConstantClassInfo).class.findField(signature)
	}
	return field
}


func (this *JavaClass) findMethod(signature string) *JavaMethod {
	method := this.methodsMap[signature]
	if method == nil {
		method = this.constantPool[this.superClass].resolve().(*RuntimeConstantClassInfo).class.findMethod(signature)
	}
	return method
}