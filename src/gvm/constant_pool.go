package gvm

import (
	//"fmt"
)

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
	Read(reader *ClassReader)
}

/*
CONSTANT_Class_info {
    u1 tag;
    u2 name_index;
}
 */
type ConstantClassInfo struct {
	tag       uint8
	nameIndex uint16
}

func (this *ConstantClassInfo) Read(reader *ClassReader) {
	this.nameIndex = reader.ReadUint16()

	//fmt.Printf("Class\t\t#%d\n", this.nameIndex)
}

/*
CONSTANT_Fieldref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
 */
type ConstantFieldrefInfo struct {
	tag              uint8
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (this *ConstantFieldrefInfo) Read(reader *ClassReader) {
	this.classIndex = reader.ReadUint16()
	this.nameAndTypeIndex = reader.ReadUint16()

	//fmt.Printf("Fieldref\t\t#%d.#%d\n", this.classIndex, this.nameAndTypeIndex)
}

/*
CONSTANT_Methodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
 */
type ConstantMethodrefInfo struct {
	tag              uint8
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (this *ConstantMethodrefInfo) Read(reader *ClassReader) {
	this.classIndex = reader.ReadUint16()
	this.nameAndTypeIndex = reader.ReadUint16()

	//fmt.Printf("Methodref\t#%d.#%d\n", this.classIndex, this.nameAndTypeIndex)
}

/*
CONSTANT_InterfaceMethodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
 */
type ConstantInterfaceMethodrefInfo struct {
	tag              uint8
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (this *ConstantInterfaceMethodrefInfo) Read(reader *ClassReader) {
	this.classIndex = reader.ReadUint16()
	this.nameAndTypeIndex = reader.ReadUint16()

	//fmt.Printf("InterfaceMethodref\t\t#%d.#%d\n", this.classIndex, this.nameAndTypeIndex)
}

/*
CONSTANT_String_info {
    u1 tag;
    u2 string_index;
}
 */
type ConstantStringInfo struct {
	tag         uint8
	stringIndex uint16
}

func (this *ConstantStringInfo) Read(reader *ClassReader) {
	this.stringIndex = reader.ReadUint16()

	//fmt.Printf("String\t\t#%d\n", this.stringIndex)
}

/*
CONSTANT_Integer_info {
    u1 tag;
    u4 bytes;
}
 */
type ConstantIntegerInfo struct {
	tag   uint8
	bytes uint32
}

func (this *ConstantIntegerInfo) Read(reader *ClassReader) {
	this.bytes = reader.ReadUint32()

	//fmt.Printf("Integer\t\t%s\n", this.bytes)
}

/*
CONSTANT_Float_info {
    u1 tag;
    u4 bytes;
}
 */
type ConstantFloatInfo struct {
	tag   uint8
	bytes uint32
}

func (this *ConstantFloatInfo) Read(reader *ClassReader) {
	this.bytes = reader.ReadUint32()

	//fmt.Printf("Float\t\t%s\n", this.bytes)
}

/*
CONSTANT_Long_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
 */
type ConstantLongInfo struct {
	tag       uint8
	highBytes uint32
	lowBytes  uint32
}

func (this *ConstantLongInfo) Read(reader *ClassReader) {
	this.highBytes = reader.ReadUint32()
	this.lowBytes = reader.ReadUint32()
	//fmt.Printf("Long\t\t%s%s\n", this.highBytes, this.lowBytes)
}

/*
CONSTANT_Double_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
 */
type ConstantDoubleInfo struct {
	tag       uint8
	highBytes uint32
	lowBytes  uint32
}

func (this *ConstantDoubleInfo) Read(reader *ClassReader) {
	this.highBytes = reader.ReadUint32()
	this.lowBytes = reader.ReadUint32()
	//fmt.Printf("Double\t\t%s%s\n", this.highBytes, this.lowBytes)
}

/*
CONSTANT_NameAndType_info {
    u1 tag;
    u2 name_index;
    u2 descriptor_index;
}
 */
type ConstantNameAndTypeInfo struct {
	tag             uint8
	nameIndex       uint16
	descriptorIndex uint16
}

func (this *ConstantNameAndTypeInfo) Read(reader *ClassReader) {
	this.nameIndex = reader.ReadUint16()
	this.descriptorIndex = reader.ReadUint16()
	//fmt.Printf("NameAndType\t#%d:#%d\n", this.nameIndex, this.descriptorIndex)
}

/*
CONSTANT_Utf8_info {
    u1 tag;
    u2 length;
    u1 bytes[length];
}
 */
type ConstantUtf8Info struct {
	tag    uint8
	bytes []byte //u2 length
}

func (this *ConstantUtf8Info) Read(reader *ClassReader) {
	this.bytes = reader.ReadBytes(int(reader.ReadUint16()))
	//fmt.Printf("Utf8\t\t%s\n", this.bytes)
}

/*
CONSTANT_MethodHandle_info {
    u1 tag;
    u1 reference_kind;
    u2 reference_index;
}
 */
type ConstantMethodHandleInfo struct {
	tag            uint8
	referenceKind  uint8
	referenceIndex uint16
}

func (this *ConstantMethodHandleInfo) Read(reader *ClassReader) {
	this.referenceKind = reader.ReadBytes(1)[0]
	this.referenceIndex = reader.ReadUint16()
	//fmt.Printf("MethodHandle\t\t%s%s\n", this.referenceKind, this.referenceIndex)
}

/*
CONSTANT_MethodType_info {
    u1 tag;
    u2 descriptor_index;
}
 */
type ConstantMethodTypeInfo struct {
	tag             uint8
	descriptorIndex uint16
}

func (this *ConstantMethodTypeInfo) Read(reader *ClassReader) {
	this.descriptorIndex = reader.ReadUint16()
	//fmt.Printf("MethodType\t%s\n", this.descriptorIndex)
}

/*
CONSTANT_InvokeDynamic_info {
    u1 tag;
    u2 bootstrap_method_attr_index;
    u2 name_and_type_index;
}
 */
type ConstantInvokeDynamicInfo struct {
	tag                      uint8
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (this *ConstantInvokeDynamicInfo) Read(reader *ClassReader) {
	this.bootstrapMethodAttrIndex = reader.ReadUint16()
	this.nameAndTypeIndex = reader.ReadUint16()
	//fmt.Printf("InvokeDynamic\t\t%s%s\n", this.bootstrapMethodAttrIndex, this.nameAndTypeIndex)
}