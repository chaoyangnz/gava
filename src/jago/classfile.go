package jago

import (
	"fmt"
)

type (
	u1 uint8
	u2 uint16
	u4 uint32
	u8 uint64
)



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
	size 	            int
	magic               u4
	minorVersion        u2
	majorVersion        u2
	constantPoolCount   u2
	constantPool        []ConstantPoolInfo
	accessFlags         u2
	thisClass           u2
	superClass          u2
	interfaces          []u2
	fieldsCount         u2
	fields              []FieldInfo
	methodsCount        u2
	methods             []MethodInfo
	attributes          []AttributeInfo
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

func NewClassFile() *ClassFile {
	return &ClassFile{}
}

func (this *ClassFile) Print(){
	fmt.Printf("Size: %d bytes\n", this.size)
	fmt.Printf("magic: 0x%X\n", this.magic)
	fmt.Printf("minor version: %d\n", this.minorVersion)
	fmt.Printf("major version: %d\n", this.majorVersion)

	fmt.Printf("accessFlags: 0x%04x\n", this.accessFlags)
	fmt.Printf("thisClass: #%d\n", this.thisClass)
	fmt.Printf("superClass: #%d\n", this.superClass)
	fmt.Printf("interfaces: %d\n", len(this.interfaces))
	for i := 0; i < len(this.interfaces); i++  {
		fmt.Printf("\t#%d", this.interfaces[i])
	}

	fmt.Printf("fields: %d\n", len(this.fields))
	for i := 0; i < len(this.fields); i++  {
		fieldInfo := this.fields[i]
		fmt.Printf("\t%s", this.cpUtf8(fieldInfo.nameIndex))
	}

	fmt.Printf("method: %d\n", len(this.methods))
	for i := 0; i < len(this.methods); i++  {
		methodInfo := this.methods[i]
		fmt.Printf("\t%s\n", this.cpUtf8(methodInfo.nameIndex))
		for j :=0; j < len(methodInfo.attributes); j++ {
			attribute := methodInfo.attributes[j]
			this.printAttribute(attribute)
		}
	}

	fmt.Printf("attributes: %d\n", len(this.attributes))
	for j :=0; j < len(this.attributes); j++ {
		attribute := this.attributes[j]
		this.printAttribute(attribute)
	}
}

func (this *ClassFile) printAttribute(attribute AttributeInfo)  {
	switch attribute.(type) {
	case *CodeAttribute:
		codeAttribute := attribute.(*CodeAttribute)
		fmt.Printf("\t\tCode: %v\n", codeAttribute.code)
		fmt.Printf("\t\tMax locals: %d\n", codeAttribute.maxLocals)
		fmt.Printf("\t\tMax stack: %d\n", codeAttribute.maxStack)
	case *SourceFileAttribue:
		sourceFileAttribute := attribute.(*SourceFileAttribue)
		fmt.Printf("\t\tSourceFile: %v\n", this.cpUtf8(sourceFileAttribute.sourceFileIndex))
	case *LineNumberTableAttribute:
		lineNumberTableAttribute := attribute.(*LineNumberTableAttribute)
		fmt.Printf("\t\tlineNumberTableAttribute: %v\n", lineNumberTableAttribute)
	}
}


func (this *ClassFile) cpUtf8(index u2) string  {
	return u2s(this.constantPool[index].(*ConstantUtf8Info).bytes)
}

//func (this *ClassFile) cpClass(index u2) string  {
//	classInfo := this.constantPool[index].(*ConstantClassInfo)
//	return this.cpUtf8(classInfo.nameIndex)
//}
//
//func (this *ClassFile) cpNameAndType(index u2) (string, string)  {
//	nameAndTypeInfo := this.constantPool[index].(*ConstantNameAndTypeInfo)
//	return this.cpUtf8(nameAndTypeInfo.nameIndex), this.cpUtf8(nameAndTypeInfo.descriptorIndex)
//}
//
//// FieldRef, MethodRef, InterfaceMethodRef
//func (this *ClassFile) cpMemberRef(index u2) (string, string, string)  {
//	memberRefInfo := this.constantPool[index].(*ConstantFieldrefInfo)
//	name, signature := this.cpNameAndType(memberRefInfo.nameAndTypeIndex)
//	return this.cpClass(memberRefInfo.classIndex), name, signature
//}
//
//
//func (this *ClassFile) cpString(index u2) string  {
//	stringInfo := this.constantPool[index].(*ConstantStringInfo)
//	return this.cpUtf8(stringInfo.stringIndex)
//}
//
//func (this *ClassFile) cpInteger(index u2) int32  {
//	integerInfo := this.constantPool[index].(*ConstantIntegerInfo)
//	return int32(integerInfo.bytes)
//}
//
//func (this *ClassFile) cpLong(index u2) int64  {
//	longInfo := this.constantPool[index].(*ConstantLongInfo)
//	return int64((longInfo.highBytes << 32) | longInfo.lowBytes)
//}
//
//func (this *ClassFile) cpFloat(index u2) float32  {
//	floatInfo := this.constantPool[index].(*ConstantFloatInfo)
//	return float32(floatInfo.bytes)
//}
//
//func (this *ClassFile) cpDouble(index u2) float64  {
//	doubleInfo := this.constantPool[index].(*ConstantDoubleInfo)
//	return float64((doubleInfo.highBytes << 32) | doubleInfo.lowBytes)
//}


const (
	CONSTANT_Class              = 7
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_String             = 8
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_NameAndType        = 12
	CONSTANT_Utf8               = 1
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

/*
cp_info {
    u1 tag;
    u1 info[];
}
 */
type ConstantPoolInfo interface {

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


/*
CONSTANT_Utf8_info {
    u1 tag;
    u2 length;
    u1 bytes[length];
}
 */
type ConstantUtf8Info struct {
	tag     u1
	length  u2
	bytes   []u1 //u2 length
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

func (this *ConstantMethodHandleInfo) runtime(class *Class, classfile *ClassFile) Constant {
	return &MethodHandleConstant{
		referenceKind: this.referenceKind,
		referenceIndex: this.referenceIndex}
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




/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/
type AttributeInfo interface {
	ReadInfo(reader *ClassReader)
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
	code                 []u1               //u4 code_length
	exceptionTableLength u2
	exceptionTable       []ExceptionTableEntry //u2 exception_table_length
	attributesCount      u2
	attributes           []AttributeInfo  //u2 attributes_count
}

func (this *CodeAttribute) ReadInfo(reader *ClassReader) {
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
	attributesCount := reader.readU2()
	this.attributesCount = attributesCount
	this.attributes = make([]AttributeInfo, attributesCount)
	for i := u2(0); i < attributesCount; i++ {
		reader.readAttribute(&this.attributes[i])
	}
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

func (this *LineNumberTableAttribute) ReadInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readU2()
	this.lineNumberTableLength = lineNumberTableLength
	this.lineNumberTable = make([]LineNumberTableEntry, lineNumberTableLength)
	for i := u2(0); i < lineNumberTableLength; i++ {
		this.lineNumberTable[i] = LineNumberTableEntry{reader.readU2(), reader.readU2()}
	}
}

/*
SourceFile_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 sourcefile_index;
}
 */

type SourceFileAttribue struct {
	attributeNameIndex u2
	attributeLength    u4
	sourceFileIndex u2
}

func (this *SourceFileAttribue) ReadInfo(reader *ClassReader) {
	this.sourceFileIndex = reader.readU2()
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
	startPc             u2
	length              u2
	nameIndex           u2
	descriptorIndex     u2
	index               u2
}

func (this *LocalVariableTableAttribute) ReadInfo(reader *ClassReader) {
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
