package jago

import (
	"unsafe"
	. "encoding/binary"
)

type (
	u1 = uint8
	u2 = uint16
	u4 = uint32
	u8 = uint64
)

type ClassReader struct {
	bytecode  []uint8
	classfile *ClassFile
}

func (this *ClassReader) readU4() u4 {
	value := BigEndian.Uint32(this.bytecode[:4])
	this.bytecode = this.bytecode[4:]
	return value
}

func (this *ClassReader) readU2() u2 {
	value := BigEndian.Uint16(this.bytecode[:2])
	this.bytecode = this.bytecode[2:]
	return value
}

func (this *ClassReader) readU1() u1 {
	return this.readU1s(1)[0]
}

func (this *ClassReader) readU1s(length uint32) []u1 {
	bytes := this.bytecode[:length]
	this.bytecode = this.bytecode[length:]
	return bytes
}

func (this *ClassReader) length() int {
	return len(this.bytecode)
}

// read all attributes for class, field or method or attribute
func (this *ClassReader) readAttributes() (u2, []AttributeInfo) {
	attributesCount := this.readU2()
	attributes := make([]AttributeInfo, attributesCount)
	for i := u2(0); i < attributesCount; i++ {
		attributes[i] = this.readAttribute()
	}

	return attributesCount, attributes
}

// read any attribute
func (this *ClassReader) readAttribute() AttributeInfo {
	attributeNameIndex := this.readU2()
	attributeLength := this.readU4()
	attributeName := this.classfile.constantPool[attributeNameIndex].(*ConstantUtf8Info).value()
	var attributeInfo AttributeInfo
	switch attributeName {
	case "Code":
		attributeInfo = &CodeAttribute{attributeNameIndex: attributeNameIndex, attributeLength: attributeLength}
	case "SourceFile":
		attributeInfo = &SourceFileAttribute{attributeNameIndex: attributeNameIndex, attributeLength: attributeLength}
	case "LineNumberTable":
		attributeInfo = &LineNumberTableAttribute{attributeNameIndex: attributeNameIndex, attributeLength: attributeLength}
	case "LocalVariableTable":
		attributeInfo = &LocalVariableTableAttribute{attributeNameIndex: attributeNameIndex, attributeLength: attributeLength}
	default:
		VM.BootstrapClassLoader.All("No reader for attribute: %s, skip\n", attributeName)
		this.readU1s(attributeLength) // just skip out
	}

	if attributeInfo != nil {
		attributeInfo.readInfo(this)
	}

	return attributeInfo
}

/*
ClassFile {
	u4				magic;
	u2 				minor_version;
	u2 				major_version;
	u2 				constant_pool_count;
	cp_info 		constant_pool[constant_pool_count-1];
	u2 				access_flags;
	u2 				this_class;
	u2 				super_class;
	u2 				interfaces_count;
	u2 				interfaces[interfaces_count];
	u2 				fields_count;
	field_info 		fields[fields_count];
	u2 				methods_count;
	method_info 	methods[methods_count];
	u2 				attributes_count;
	attribute_info 	attributes[attributes_count];
}
*/
type ClassFile struct {
	magic             u4
	minorVersion      u2
	majorVersion      u2
	constantPoolCount u2
	constantPool      []ConstantPoolInfo
	accessFlags       u2
	thisClass         u2
	superClass        u2
	interfaceCount    u2
	interfaces        []u2
	fieldsCount       u2
	fields            []FieldInfo
	methodsCount      u2
	methods           []MethodInfo
	attributeCount    u2
	attributes        []AttributeInfo
}

func (this *ClassFile) read(bytes []byte) {
	reader := &ClassReader{bytes, this}

	this.readMagic(reader)
	this.readMinorVersion(reader)
	this.readMajorVersion(reader)
	this.readConstantPool(reader)
	this.readAccessFlags(reader)
	this.readThisClass(reader)
	this.readSuperClass(reader)
	this.readInterfaces(reader)
	this.readFields(reader)
	this.readMethods(reader)

	this.attributeCount, this.attributes = reader.readAttributes()
}

func (this *ClassFile) readMagic(reader *ClassReader) {
	this.magic = reader.readU4()
}

func (this *ClassFile) readMinorVersion(reader *ClassReader) {
	this.minorVersion = reader.readU2()
}

func (this *ClassFile) readMajorVersion(reader *ClassReader) {
	this.majorVersion = reader.readU2()
}

func (this *ClassFile) readConstantPool(reader *ClassReader) {
	var constantPoolCount = reader.readU2()
	this.constantPoolCount = constantPoolCount
	this.constantPool = make([]ConstantPoolInfo, constantPoolCount)
	for i := u2(1); i < constantPoolCount; i++ {
		//LOG.Trace(" #%2d = ", i)
		tag := reader.readU1()

		var cpInfo ConstantPoolInfo
		switch tag {
		case CONSTANT_Class:
			cpInfo = &ConstantClassInfo{tag: tag}
		case CONSTANT_Fieldref:
			cpInfo = &ConstantFieldrefInfo{tag: tag}
		case CONSTANT_Methodref:
			cpInfo = &ConstantMethodrefInfo{tag: tag}
		case CONSTANT_InterfaceMethodref:
			cpInfo = &ConstantInterfaceMethodrefInfo{tag: tag}
		case CONSTANT_String:
			cpInfo = &ConstantStringInfo{tag: tag}
		case CONSTANT_Integer:
			cpInfo = &ConstantIntegerInfo{tag: tag}
		case CONSTANT_Float:
			cpInfo = &ConstantFloatInfo{tag: tag}
		case CONSTANT_Long:
			// occupy two entries
			cpInfo = &ConstantLongInfo{tag: tag}
			this.constantPool[i] = cpInfo
			i++
		case CONSTANT_Double:
			// occupy two entries
			cpInfo = &ConstantDoubleInfo{tag: tag}
			this.constantPool[i] = cpInfo
			i++
		case CONSTANT_NameAndType:
			cpInfo = &ConstantNameAndTypeInfo{tag: tag}
		case CONSTANT_Utf8:
			cpInfo = &ConstantUtf8Info{tag: tag}
		case CONSTANT_MethodHandle:
			cpInfo = &ConstantMethodHandleInfo{tag: tag}
		case CONSTANT_MethodType:
			cpInfo = &ConstantMethodTypeInfo{tag: tag}
		case CONSTANT_InvokeDynamic:
			cpInfo = &ConstantInvokeDynamicInfo{tag: tag}
		default:
			// ignore
			VM.Throw("java/lang/ClassFormatError", "Not a supported constant tag: %d", tag)
		}
		cpInfo.readInfo(reader)
		this.constantPool[i] = cpInfo
	}
}

func (this *ClassFile) readAccessFlags(reader *ClassReader) {
	this.accessFlags = reader.readU2()
}

func (this *ClassFile) readThisClass(reader *ClassReader) {
	this.thisClass = reader.readU2()
}

func (this *ClassFile) readSuperClass(reader *ClassReader) {
	this.superClass = reader.readU2()
}

func (this *ClassFile) readInterfaces(reader *ClassReader) {
	var interfacesCount = reader.readU2()
	this.interfaceCount = interfacesCount
	this.interfaces = make([]u2, interfacesCount)
	for i := u2(0); i < interfacesCount; i++ {
		this.interfaces[i] = reader.readU2()
	}
}

func (this *ClassFile) readFields(reader *ClassReader) {
	var fieldsCount = reader.readU2()
	this.fieldsCount = fieldsCount
	this.fields = make([]FieldInfo, fieldsCount)
	for i := u2(0); i < fieldsCount; i++ {
		fieldInfo := FieldInfo{}
		fieldInfo.accessFlags = reader.readU2()
		fieldInfo.nameIndex = reader.readU2()
		fieldInfo.descriptorIndex = reader.readU2()
		fieldInfo.attributeCount, fieldInfo.attributes = reader.readAttributes()
		this.fields[i] = fieldInfo
	}
}

func (this *ClassFile) readMethods(reader *ClassReader) {
	methodsCount := reader.readU2()
	this.methodsCount = methodsCount
	this.methods = make([]MethodInfo, methodsCount)
	for i := u2(0); i < methodsCount; i++ {
		methodInfo := MethodInfo{}
		methodInfo.accessFlags = reader.readU2()
		methodInfo.nameIndex = reader.readU2()
		methodInfo.descriptorIndex = reader.readU2()
		methodInfo.attributeCount, methodInfo.attributes = reader.readAttributes()
		this.methods[i] = methodInfo
	}
}

func (this *ClassFile) cpUtf8(index u2) string {
	u1s := this.constantPool[index].(*ConstantUtf8Info).bytes
	bytes := *(*[]byte)(unsafe.Pointer(&u1s))
	return string(bytes)
}

/*
cp_info {
    u1 tag;
    u1 info[];
}
 */
type ConstantPoolInfo interface {
	readInfo(reader *ClassReader)
}

/*
CONSTANT_Class_info {
    u1 tag;
    u2 name_index;
}
 */
type ConstantClassInfo struct {
	tag       u1
	nameIndex u2
}

func (this *ConstantClassInfo) readInfo(reader *ClassReader) {
	this.nameIndex = reader.readU2()
}

/*
CONSTANT_Fieldref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
 */
type ConstantFieldrefInfo struct {
	tag              u1
	classIndex       u2
	nameAndTypeIndex u2
}

func (this *ConstantFieldrefInfo) readInfo(reader *ClassReader) {
	this.classIndex = reader.readU2()
	this.nameAndTypeIndex = reader.readU2()
}

/*
CONSTANT_Methodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
 */
type ConstantMethodrefInfo struct {
	tag              u1
	classIndex       u2
	nameAndTypeIndex u2
}

func (this *ConstantMethodrefInfo) readInfo(reader *ClassReader) {
	this.classIndex = reader.readU2()
	this.nameAndTypeIndex = reader.readU2()
}

/*
CONSTANT_InterfaceMethodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
 */
type ConstantInterfaceMethodrefInfo struct {
	tag              u1
	classIndex       u2
	nameAndTypeIndex u2
}

func (this *ConstantInterfaceMethodrefInfo) readInfo(reader *ClassReader) {
	this.classIndex = reader.readU2()
	this.nameAndTypeIndex = reader.readU2()
}

/*
CONSTANT_String_info {
    u1 tag;
    u2 string_index;
}
 */
type ConstantStringInfo struct {
	tag         u1
	stringIndex u2
}

func (this *ConstantStringInfo) readInfo(reader *ClassReader) {
	this.stringIndex = reader.readU2()
}

/*
CONSTANT_Integer_info {
    u1 tag;
    u4 bytes;
}
 */
type ConstantIntegerInfo struct {
	tag   u1
	bytes u4
}

func (this *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	this.bytes = reader.readU4()
}

/*
CONSTANT_Float_info {
    u1 tag;
    u4 bytes;
}
 */
type ConstantFloatInfo struct {
	tag   u1
	bytes u4
}

func (this *ConstantFloatInfo) readInfo(reader *ClassReader) {
	this.bytes = reader.readU4()
}

/*
CONSTANT_Long_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
 */
type ConstantLongInfo struct {
	tag       u1
	highBytes u4
	lowBytes  u4
}

func (this *ConstantLongInfo) readInfo(reader *ClassReader) {
	this.highBytes = reader.readU4()
	this.lowBytes = reader.readU4()
}

/*
CONSTANT_Double_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
 */
type ConstantDoubleInfo struct {
	tag       u1
	highBytes u4
	lowBytes  u4
}

func (this *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	this.highBytes = reader.readU4()
	this.lowBytes = reader.readU4()
}

/*
CONSTANT_NameAndType_info {
    u1 tag;
    u2 name_index;
    u2 descriptor_index;
}
 */
type ConstantNameAndTypeInfo struct {
	tag             u1
	nameIndex       u2
	descriptorIndex u2
}

func (this *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	this.nameIndex = reader.readU2()
	this.descriptorIndex = reader.readU2()
}

/*
CONSTANT_Utf8_info {
    u1 tag;
    u2 length;
    u1 bytes[length];
}
 */
type ConstantUtf8Info struct {
	tag    u1
	length u2
	bytes  []u1 //u2 length
}

func (this *ConstantUtf8Info) readInfo(reader *ClassReader) {
	this.length = reader.readU2()
	this.bytes = reader.readU1s(uint32(this.length))
}

func (this *ConstantUtf8Info) value() string {
	return string(*(*[]byte)(&this.bytes))
}

/*
CONSTANT_MethodHandle_info {
    u1 tag;
    u1 reference_kind;
    u2 reference_index;
}
 */
type ConstantMethodHandleInfo struct {
	tag            u1
	referenceKind  u1
	referenceIndex u2
}

func (this *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {
	this.referenceKind = reader.readU1()
	this.referenceIndex = reader.readU2()
}

/*
CONSTANT_MethodType_info {
    u1 tag;
    u2 descriptor_index;
}
 */
type ConstantMethodTypeInfo struct {
	tag             u1
	descriptorIndex u2
}

func (this *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	this.descriptorIndex = reader.readU2()
}

/*
CONSTANT_InvokeDynamic_info {
    u1 tag;
    u2 bootstrap_method_attr_index;
    u2 name_and_type_index;
}
 */
type ConstantInvokeDynamicInfo struct {
	tag                      u1
	bootstrapMethodAttrIndex u2
	nameAndTypeIndex         u2
}

func (this *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	this.bootstrapMethodAttrIndex = reader.readU2()
	this.nameAndTypeIndex = reader.readU2()
}

/*
field_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type FieldInfo struct {
	accessFlags     u2
	nameIndex       u2
	descriptorIndex u2
	attributeCount  u2
	attributes      []AttributeInfo
}

/*
method_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type MethodInfo struct {
	accessFlags     u2
	nameIndex       u2
	descriptorIndex u2
	attributeCount  u2
	attributes      []AttributeInfo
}

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/
type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

/*
Code_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 max_stack;
    u2 max_locals;
    u4 code_length;
    u1 code[code_length];
    u2 exception_table_length;
    {   u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type CodeAttribute struct {
	attributeNameIndex   u2
	attributeLength      u4
	maxStack             u2
	maxLocals            u2
	codeLength           u4
	code                 []u1 //u4 code_length
	exceptionTableLength u2
	exceptionTable       []ExceptionTableEntry //u2 exception_table_length
	attributesCount      u2
	attributes           []AttributeInfo //u2 attributes_count
}

func (this *CodeAttribute) readInfo(reader *ClassReader) {
	this.maxStack = reader.readU2()
	this.maxLocals = reader.readU2()
	codeLength := reader.readU4()
	this.codeLength = codeLength
	this.code = reader.readU1s(uint32(codeLength))
	exceptionTableLength := reader.readU2()
	this.exceptionTableLength = exceptionTableLength
	this.exceptionTable = make([]ExceptionTableEntry, exceptionTableLength)
	for i := u2(0); i < exceptionTableLength; i++ {
		exceptionTable := ExceptionTableEntry{}
		exceptionTable.ReadInfo(reader)
		this.exceptionTable[i] = exceptionTable
	}

	this.attributesCount, this.attributes = reader.readAttributes()
}

type ExceptionTableEntry struct {
	startPc   u2
	endPc     u2
	handlerPc u2
	catchType u2
}

func (this *ExceptionTableEntry) ReadInfo(reader *ClassReader) {
	this.startPc = reader.readU2()
	this.endPc = reader.readU2()
	this.handlerPc = reader.readU2()
	this.catchType = reader.readU2()
}

/*
LineNumberTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 line_number_table_length;
    {   u2 start_pc;
        u2 line_number;
    } line_number_table[line_number_table_length];
}
*/
type LineNumberTableAttribute struct {
	attributeNameIndex    u2
	attributeLength       u4
	lineNumberTableLength u2
	lineNumberTable       []LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc    u2
	lineNumber u2
}

func (this *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readU2()
	this.lineNumberTableLength = lineNumberTableLength
	this.lineNumberTable = make([]LineNumberTableEntry, lineNumberTableLength)
	for i := u2(0); i < lineNumberTableLength; i++ {
		this.lineNumberTable[i] = LineNumberTableEntry{reader.readU2(), reader.readU2()}
	}
}

/*
LocalVariableTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 local_variable_table_length;
    {   u2 start_pc;
        u2 length;
        u2 name_index;
        u2 descriptor_index;
        u2 index;
    } local_variable_table[local_variable_table_length];
}
 */
type LocalVariableTableAttribute struct {
	attributeNameIndex       u2
	attributeLength          u4
	localVariableTableLength u2
	localVariableTable       []LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	startPc         u2
	length          u2
	nameIndex       u2
	descriptorIndex u2
	index           u2
}

func (this *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	localVariableTableLength := reader.readU2()
	this.localVariableTableLength = localVariableTableLength
	this.localVariableTable = make([]LocalVariableTableEntry, localVariableTableLength)
	for i := u2(0); i < localVariableTableLength; i++ {
		this.localVariableTable[i] = LocalVariableTableEntry{
			reader.readU2(),
			reader.readU2(),
			reader.readU2(),
			reader.readU2(),
			reader.readU2()}
	}
}

/*
SourceFile_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 sourcefile_index;
}
 */

type SourceFileAttribute struct {
	attributeNameIndex u2
	attributeLength    u4
	sourceFileIndex    u2
}

func (this *SourceFileAttribute) readInfo(reader *ClassReader) {
	this.sourceFileIndex = reader.readU2()
}

/*
RuntimeVisibleAnnotations_attribute {
    u2         attribute_name_index;
    u4         attribute_length;
    u2         num_annotations;
    annotation annotations[num_annotations];
}
 */

type RuntimeVisibleAnnotationsAttribute struct {
	attributeNameIndex u2
	attributeLength    u2
	numAnnotations     u2
	annotations        []Annotation
}

/*
annotation {
    u2 type_index;
    u2 num_element_value_pairs;
    {   u2            element_name_index;
        element_value value;
    } element_value_pairs[num_element_value_pairs];
}
 */
type Annotation struct {
	typeIndex u2
	elementValuePairs []struct {
		element_name_index u2
		value              ElementValue
	}
}

/*
element_value {
    u1 tag;
    union {
        u2 const_value_index;

        {   u2 type_name_index;
            u2 const_name_index;
        } enum_const_value;

        u2 class_info_index;

        annotation annotation_value;

        {   u2            num_values;
            element_value values[num_values];
        } array_value;
    } value;
}
 */

type ElementValue struct {
	tag               u1
	const_value_index u2
	enum_const_value struct {
		type_name_index  u2
		const_name_index u2
	}
	class_info_index u2
	array_value struct {
		num_values u2
		values     []ElementValue
	}
}
