package gvm

import (
	"io/ioutil"
)

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
	java_reference     *JavaObject
	java_array         *JavaArray
)

type JavaObject struct {
	//header part
	class *JavaClass
	flags uint32
	locks uint32
	//fields
	fields []java_any
}

func (this *JavaObject) getField(index uint16) java_any {
	i := this.class.constantPool[index].resolve(this.class).(*RuntimeConstantFieldrefInfo).field.index
	return this.fields[this.class.instanceFieldsStart + i]
}

func (this *JavaObject) putField(caller *JavaClass, index uint16, value java_any) {
	i := caller.constantPool[index].resolve(caller).(*RuntimeConstantFieldrefInfo).field.index
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
	resolve(class *JavaClass) RuntimeConstantPoolInfo
}

/*
CONSTANT_Class_info {
    u1 tag;
    u2 name_index;
}
 */
type RuntimeConstantClassInfo struct {
	nameIndex u2
	resolved  bool
	name      string
	class     *JavaClass
}

func (this *RuntimeConstantClassInfo) resolve(class *JavaClass) RuntimeConstantPoolInfo {
	if !this.resolved {
		name := class.constantPool[this.nameIndex].(*RuntimeConstantUtf8Info).value
		clazz := classCache[name]
		if clazz == nil {
			if this == class.constantPool[class.thisClass] {
				clazz = class
			} else {
				bytes, _ := ioutil.ReadFile("java/" + name + ".class")
				cr := NewClassReader(bytes)
				cf := cr.ReadAsClassFile()

				clazz = NewClassMirror()
				clazz.Load(cf)
			}

		}

		this.name = name
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
	classIndex       u2
	nameAndTypeIndex u2
	resolved         bool
	field            *JavaField
}

func (this *RuntimeConstantFieldrefInfo) resolve(class *JavaClass) RuntimeConstantPoolInfo  {
	if !this.resolved {
		rcpClass := class.constantPool[this.classIndex].resolve(class).(*RuntimeConstantClassInfo)

		nameAndType := class.constantPool[this.nameAndTypeIndex].(*RuntimeConstantNameAndTypeInfo)
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
	classIndex       u2
	nameAndTypeIndex u2
	resolved         bool
	method           *JavaMethod
}

func (this *RuntimeConstantMethodrefInfo) resolve(class *JavaClass) RuntimeConstantPoolInfo  {
	if !this.resolved {
		rcpClass := class.constantPool[this.classIndex].resolve(class).(*RuntimeConstantClassInfo)

		nameAndType := class.constantPool[this.nameAndTypeIndex].(*RuntimeConstantNameAndTypeInfo)
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
	classIndex       u2
	nameAndTypeIndex u2
	resolved         bool
	method           *JavaMethod
}

func (this *RuntimeConstantInterfaceMethodrefInfo) resolve(class *JavaClass) RuntimeConstantPoolInfo  {
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
	stringIndex     u2
	resolved        bool
	value           []java_char
}


func (this *RuntimeConstantStringInfo) resolve(class *JavaClass) RuntimeConstantPoolInfo {
	if !this.resolved {
		this.value = string2JavaChars(class.constantPool[this.stringIndex].resolve(class).(*RuntimeConstantUtf8Info).value)
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
	bytes       u4
	resolved    bool
	value       java_int
}


func (this *RuntimeConstantIntegerInfo) resolve(class *JavaClass) RuntimeConstantPoolInfo  {
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
	bytes       u4
	resolved    bool
	value       java_float
}

func (this *RuntimeConstantFloatInfo) resolve(class *JavaClass) RuntimeConstantPoolInfo  {
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
	highBytes   u4
	lowBytes    u4
	resolved    bool
	value       java_long
}

func (this *RuntimeConstantLongInfo) resolve(class *JavaClass) RuntimeConstantPoolInfo  {
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
	highBytes   u4
	lowBytes    u4
	resolved    bool
	value       java_double
}


func (this *RuntimeConstantDoubleInfo) resolve(class *JavaClass) RuntimeConstantPoolInfo  {
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
	nameIndex       u2
	descriptorIndex u2
	resolved        bool
	name            string
	descriptor      string
}

func (this *RuntimeConstantNameAndTypeInfo) resolve(class *JavaClass) RuntimeConstantPoolInfo  {
	if !this.resolved {
		this.name = class.constantPool[this.nameIndex].resolve(class).(*RuntimeConstantUtf8Info).value
		this.descriptor = class.constantPool[this.descriptorIndex].resolve(class).(*RuntimeConstantUtf8Info).value
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
	length          u2
	bytes           []u1
	resolved        bool
	value           string
}

func (this *RuntimeConstantUtf8Info) resolve(class *JavaClass) RuntimeConstantPoolInfo {
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
	referenceKind   u1
	referenceIndex  u2
	resolved        bool
	//TODO
}

func (this *RuntimeConstantMethodHandleInfo) resolve(class *JavaClass) RuntimeConstantPoolInfo {
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
	descriptorIndex u2
	resolved        bool
	descriptor      string
}

func (this *RuntimeConstantMethodTypeInfo) resolve(class *JavaClass) RuntimeConstantPoolInfo {
	if !this.resolved {
		this.descriptor = class.constantPool[this.descriptorIndex].resolve(class).(*RuntimeConstantUtf8Info).value
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
	bootstrapMethodAttrIndex u2
	nameAndTypeIndex         u2
	resolved                 bool
	//TODO
}


func (this *RuntimeConstantInvokeDynamicInfo) resolve(class *JavaClass) RuntimeConstantPoolInfo {
	if !this.resolved {
		//TODO
	}
	return this
}


type JavaClass struct {
	constantPool    []RuntimeConstantPoolInfo
	accessFlags     uint16
	thisClass       uint16
	superClass      uint16
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

func (this *JavaMethod)isStatic() bool {
	return this.accessFlags & METHOD_ACC_STATIC > 0
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

func NewClassMirror() *JavaClass {
	return &JavaClass{}
}

var classCache = make(map[string] *JavaClass)

func (this *JavaClass) Load(classfile *ClassFile)  {
	this.constantPool = make([]RuntimeConstantPoolInfo, classfile.constantPoolCount)
	for i := u2(1); i < classfile.constantPoolCount; i++ {
		cpInfo := classfile.constantPool[i]
		switch cpInfo.(type) {
		case *ConstantIntegerInfo:
			cp := cpInfo.(*ConstantIntegerInfo)
			this.constantPool[i] = &RuntimeConstantIntegerInfo{
				cp.bytes,
				true,
				java_int(cp.bytes)}
		case *ConstantLongInfo:
			cp := cpInfo.(*ConstantLongInfo)
			this.constantPool[i] = &RuntimeConstantLongInfo{
				cp.highBytes, cp.lowBytes,
				true,
				java_long((cp.highBytes << 32) | cp.lowBytes)}
		case *ConstantFloatInfo:
			cp := cpInfo.(*ConstantFloatInfo)
			this.constantPool[i] = &RuntimeConstantFloatInfo{
				cp.bytes,
				true,
				java_float(cp.bytes)}
		case *ConstantDoubleInfo:
			cp := cpInfo.(*ConstantDoubleInfo)
			this.constantPool[i] = &RuntimeConstantDoubleInfo{
				cp.highBytes,
				cp.lowBytes,
				true,
				java_double((cp.highBytes << 32) | cp.lowBytes)}
		case *ConstantStringInfo:
			cp := cpInfo.(*ConstantStringInfo)
			this.constantPool[i] = &RuntimeConstantStringInfo{
				cp.stringIndex,
				true,
				string2JavaChars(classfile.cpUtf8(cp.stringIndex))}
		case *ConstantUtf8Info:
			cp := cpInfo.(*ConstantUtf8Info)
			this.constantPool[i] = &RuntimeConstantUtf8Info{length: cp.length, bytes: cp.bytes}
			this.constantPool[i].resolve(this)
		case *ConstantNameAndTypeInfo:
			cp := cpInfo.(*ConstantNameAndTypeInfo)
			this.constantPool[i] = &RuntimeConstantNameAndTypeInfo{
				cp.nameIndex,
				cp.descriptorIndex,
				true,
				classfile.cpUtf8(cp.nameIndex),
				classfile.cpUtf8(cp.descriptorIndex)}
		case *ConstantMethodTypeInfo:
			cp := cpInfo.(*ConstantMethodTypeInfo)
			this.constantPool[i] = &RuntimeConstantMethodTypeInfo{
				cp.descriptorIndex,
				true,
				classfile.cpUtf8(cp.descriptorIndex)}

		case *ConstantMethodrefInfo:
			cp := cpInfo.(*ConstantMethodrefInfo)
			this.constantPool[i] = &RuntimeConstantMethodrefInfo{
				cp.classIndex,
				cp.nameAndTypeIndex,
				false,
				nil}
		case *ConstantFieldrefInfo:
			cp := cpInfo.(*ConstantFieldrefInfo)
			this.constantPool[i] = &RuntimeConstantFieldrefInfo{
				cp.classIndex,
				cp.nameAndTypeIndex,
				false,
				nil}
		case *ConstantClassInfo:
			cp := cpInfo.(*ConstantClassInfo)
			this.constantPool[i] = &RuntimeConstantClassInfo{
				nameIndex: cp.nameIndex,
				resolved: false}
		}
	}
	this.accessFlags = uint16(classfile.accessFlags)
	// resolve this class
	this.thisClass = uint16(classfile.thisClass)
	this.constantPool[classfile.thisClass].resolve(this)

	// resolve super class
	this.superClass = uint16(classfile.superClass)
	if this.superClass != 0 {
		this.constantPool[classfile.superClass].resolve(this)
	}

	this.fields = make([]*JavaField, len(classfile.fields))
	this.fieldsMap = make(map[string]*JavaField)
	this.staticFields = []java_any{}
	maxInstanceFieldIndex := uint16(0)
	maxStaticFieldIndex   := uint16(0)
	if this.superClass == 0 { // java/lang/Object
		this.instanceFieldsStart = 0
	} else {
		superClass := this.constantPool[this.superClass].resolve(this).(*RuntimeConstantClassInfo).class
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

	classCache[this.constantPool[this.thisClass].(*RuntimeConstantClassInfo).name] = this
}

func (this *JavaClass) findField(signature string) *JavaField {
	field :=  this.fieldsMap[signature]
	if field == nil {
		field = this.constantPool[this.superClass].resolve(this).(*RuntimeConstantClassInfo).class.findField(signature)
	}
	return field
}


func (this *JavaClass) findMethod(signature string) *JavaMethod {
	method := this.methodsMap[signature]
	if method == nil {
		method = this.constantPool[this.superClass].resolve(this).(*RuntimeConstantClassInfo).class.findMethod(signature)
	}
	return method
}