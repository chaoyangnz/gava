package gvm

import (
	"io/ioutil"
)

type (
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
	fields []uint64
}



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
	elements []uint64
}


/*
cp_info {
    u1 tag;
    u1 info[];
}
 */
type RuntimeConstantPoolInfo interface {
	//resolved() bool
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
	class     *JavaClass
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





type JavaClass struct {
	constantPool    []RuntimeConstantPoolInfo
	accessFlags     uint16
	thisClass       string
	superClass      string
	interfaces      []string
	fields          []*JavaField
	methods         []*JavaMethod
	fieldsMap       map[string]*JavaField
	methodsMap      map[string]*JavaMethod
	instanceFileds  []*JavaField
	staticFields    []uint64
	//attributes   []Attribute
}

type JavaField struct {
	class           *JavaClass
	accessFlags     uint16
	name            string
	descriptor      string
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
			this.constantPool[i] = &RuntimeConstantUtf8Info{
				cp.length,
				cp.bytes,
				true,
				u2s(cp.bytes)}
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
		case *ConstantClassInfo:
			cp := cpInfo.(*ConstantClassInfo)
			this.constantPool[i] = &RuntimeConstantClassInfo{
				cp.nameIndex,
				false,
				nil}
		}

	}
	this.thisClass = classfile.cpUtf8(classfile.constantPool[classfile.thisClass].(*ConstantClassInfo).nameIndex)
	this.superClass = classfile.cpUtf8(classfile.constantPool[classfile.superClass].(*ConstantClassInfo).nameIndex)
	this.fields = make([]*JavaField, len(classfile.fields))
	this.fieldsMap = make(map[string]*JavaField)
	for i := 0; i < len(classfile.fields); i++ {
		filedMirror := &JavaField{class: this}
		fieldInfo := classfile.fields[i]
		filedMirror.name = classfile.cpUtf8(fieldInfo.nameIndex)
		filedMirror.descriptor = classfile.cpUtf8(fieldInfo.descriptorIndex)
		this.fields[i] = filedMirror
		this.fieldsMap[filedMirror.name + filedMirror.descriptor] = filedMirror
	}

	this.methods = make([]*JavaMethod, len(classfile.methods))
	this.methodsMap = make(map[string]*JavaMethod)
	for i := 0; i < len(classfile.methods); i++ {
		methodMirror := &JavaMethod{class: this}
		methodInfo := &classfile.methods[i]
		methodMirror.name = classfile.cpUtf8(methodInfo.nameIndex)
		methodMirror.descriptor = classfile.cpUtf8(methodInfo.descriptorIndex)
		methodMirror.parameterDescriptors, methodMirror.returnDescriptor = parametersAndReturn(methodMirror.descriptor)
		for j := u2(0); j < methodInfo.attributeCount; j++ {
			attributeInfo := methodInfo.attributes[j]
			switch attributeInfo.(type) {
			case *CodeAttribute:
				codeAttribute := attributeInfo.(*CodeAttribute)
				methodMirror.maxStack = uint16(codeAttribute.maxStack)
				methodMirror.maxLocals = uint16(codeAttribute.maxLocals)
				methodMirror.code = u2b(codeAttribute.code)
				for k := u2(0); k < codeAttribute.attributesCount; k++ {
					codeAttributeAttribute := codeAttribute.attributes[k]
					switch codeAttributeAttribute.(type) {
					case *LocalVariableTableAttribute:
						localVariableTableAttribute := codeAttributeAttribute.(*LocalVariableTableAttribute)
						methodMirror.localVariables = make([]LocalVariable, localVariableTableAttribute.localVariableTableLength)
						for m := u2(0); m < localVariableTableAttribute.localVariableTableLength; m++ {
							methodMirror.localVariables[m] = LocalVariable{
								methodMirror,
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
		this.methods[i] = methodMirror
		this.methodsMap[methodMirror.name + methodMirror.descriptor] = methodMirror
	}

	classCache[this.thisClass] = this
}

func (this *JavaClass) resolveClass(classInfo *RuntimeConstantClassInfo) {
	classdescriptor := this.constantPool[classInfo.nameIndex].(*RuntimeConstantUtf8Info).value
	if !classInfo.resolved {
		class := classCache[classdescriptor]
		if class == nil {
			bytes, _ := ioutil.ReadFile(classdescriptor + ".class")
			cr := NewClassReader(bytes)
			cf := cr.ReadAsClassFile()

			class = NewClassMirror()
			class.Load(cf)
		}

		classInfo.class = class
		classInfo.resolved = true
	}
}

func (this *JavaClass) resolveMethodRef(methodRef *RuntimeConstantMethodrefInfo)  {
	if !methodRef.resolved {
		rcpClass := this.constantPool[methodRef.classIndex].(*RuntimeConstantClassInfo)
		this.resolveClass(rcpClass)

		nameAndType := this.constantPool[methodRef.nameAndTypeIndex].(*RuntimeConstantNameAndTypeInfo)
		methodRef.method = rcpClass.class.findMethod(nameAndType.toString())
		methodRef.resolved = true
	}
}


func (this *JavaClass) findMethod(signature string) *JavaMethod {
	return  this.methodsMap[signature]
}
