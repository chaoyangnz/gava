package gvm

import (
	"encoding/binary"
)

var bigEndian = binary.BigEndian

type ClassReader struct {
	bytecode []byte
	classfile *ClassFile
}

func NewClassReader(bytecode []byte) *ClassReader {
	return &ClassReader{bytecode: bytecode}
}

func (this *ClassReader) ReadUint32() uint32 {
	value := bigEndian.Uint32(this.bytecode[:4])
	this.bytecode = this.bytecode[4:]
	return value
}

func (this *ClassReader) ReadUint16() uint16 {
	value := bigEndian.Uint16(this.bytecode[:2])
	this.bytecode = this.bytecode[2:]
	return value
}

func (this *ClassReader) ReadUint8() uint8 {
	return uint8(this.ReadBytes(1)[0])
}

func (this *ClassReader) ReadBytes(len uint32) []byte {
	bytes := this.bytecode[:len]
	this.bytecode = this.bytecode[len:]
	return bytes
}

func (this *ClassReader) Length() int {
	return len(this.bytecode)
}

func (this *ClassReader) ReadAsClassFile() *ClassFile  {
	this.classfile = &ClassFile{}
	this.classfile.size = this.Length()
	this.readMagic()
	this.readMinorVersion()
	this.readMajorVersion()
	this.readConstantPool()
	this.readAccessFlags()
	this.readThisClass()
	this.readSuperClass()
	this.readInterfaces()
	this.readFields()
	this.readMethods()
	this.readAttributes()

	return this.classfile
}

func (this *ClassReader) readMagic() {
	this.classfile.magic = this.ReadUint32()
}

func (this *ClassReader) readMinorVersion() {
	this.classfile.minorVersion = this.ReadUint16()
}

func (this *ClassReader) readMajorVersion() {
	this.classfile.majorVersion = this.ReadUint16()
}

func (this *ClassReader) readConstantPool() {
	var constantPoolCount = this.ReadUint16()
	this.classfile.constantPoolCount = constantPoolCount
	//fmt.Printf("cp_info count: %d\n", constantPoolCount)
	this.classfile.constantPool = make([]ConstantPoolInfo, constantPoolCount)
	for i := uint16(1); i < constantPoolCount; i++ {
		//fmt.Printf(" #%2d = ", i)
		tag := this.ReadUint8()

		switch tag {
		case CONSTANT_Class:
			cpInfo := &ConstantClassInfo{tag: tag}
			this.readConstantClassInfo(cpInfo)
			this.classfile.constantPool[i] = cpInfo
		case CONSTANT_Fieldref:
			cpInfo := &ConstantFieldrefInfo{tag: tag}
			this.readConstantFieldrefInfo(cpInfo)
			this.classfile.constantPool[i] = cpInfo
		case CONSTANT_Methodref:
			cpInfo := &ConstantMethodrefInfo{tag: tag}
			this.readConstantMethodrefInfo(cpInfo)
			this.classfile.constantPool[i] = cpInfo
		case CONSTANT_InterfaceMethodref:
			cpInfo := &ConstantInterfaceMethodrefInfo{tag: tag}
			this.readConstantInterfaceMethodrefInfo(cpInfo)
			this.classfile.constantPool[i] = cpInfo
		case CONSTANT_String:
			cpInfo := &ConstantStringInfo{tag: tag}
			this.readConstantStringInfo(cpInfo)
			this.classfile.constantPool[i] = cpInfo
		case CONSTANT_Integer:
			cpInfo := &ConstantIntegerInfo{tag: tag}
			this.readConstantIntegerInfo(cpInfo)
			this.classfile.constantPool[i] = cpInfo
		case CONSTANT_Float:
			cpInfo := &ConstantFloatInfo{tag: tag}
			this.readConstantFloatInfo(cpInfo)
			this.classfile.constantPool[i] = cpInfo
		case CONSTANT_Long:
			// occupy two entries
			cpInfo := &ConstantLongInfo{tag: tag}
			this.readConstantLongInfo(cpInfo)
			this.classfile.constantPool[i] = cpInfo
			i++
			this.classfile.constantPool[i] = cpInfo
		case CONSTANT_Double:
			// occupy two entries
			cpInfo := &ConstantDoubleInfo{tag: tag}
			this.readConstantDoubleInfo(cpInfo)
			this.classfile.constantPool[i] = cpInfo
			i++
			this.classfile.constantPool[i] = cpInfo
		case CONSTANT_NameAndType:
			cpInfo := &ConstantNameAndTypeInfo{tag: tag}
			this.readConstantNameAndTypeInfo(cpInfo)
			this.classfile.constantPool[i] = cpInfo
		case CONSTANT_Utf8:
			cpInfo := &ConstantUtf8Info{tag: tag}
			this.readConstantUtf8Info(cpInfo)
			this.classfile.constantPool[i] = cpInfo
		case CONSTANT_MethodHandle:
			cpInfo := &ConstantMethodHandleInfo{tag: tag}
			this.readConstantMethodHandleInfo(cpInfo)
			this.classfile.constantPool[i] = cpInfo
		case CONSTANT_MethodType:
			cpInfo := &ConstantMethodTypeInfo{tag: tag}
			this.readConstantMethodTypeInfo(cpInfo)
			this.classfile.constantPool[i] = cpInfo
		case CONSTANT_InvokeDynamic:
			cpInfo := &ConstantInvokeDynamicInfo{tag: tag}
			this.readConstantInvokeDynamicInfo(cpInfo)
			this.classfile.constantPool[i] = cpInfo
		default:
			// ignore
		}
	}
}

func (this *ClassReader) readAccessFlags() {
	this.classfile.accessFlags = this.ReadUint16()
}

func (this *ClassReader) readThisClass() {
	this.classfile.thisClass = this.ReadUint16()
}

func (this *ClassReader) readSuperClass() {
	this.classfile.superClass = this.ReadUint16()
}

func (this *ClassReader) readInterfaces() {
	var interfacesCount = this.ReadUint16()
	this.classfile.interfaces = make([]uint16, interfacesCount)
	for i := uint16(0); i < interfacesCount; i++ {
		this.classfile.interfaces[i] = this.ReadUint16()
	}
}

func (this *ClassReader) readFields() {
	var fieldsCount = this.ReadUint16()
	this.classfile.fieldsCount = fieldsCount
	this.classfile.fields = make([]FieldInfo, fieldsCount)
	for i := uint16(0); i < fieldsCount; i++ {
		fieldInfo := FieldInfo{}
		fieldInfo.accessFlags = this.ReadUint16()
		fieldInfo.nameIndex = this.ReadUint16()
		fieldInfo.descriptorIndex = this.ReadUint16()
		var attributesCount = this.ReadUint16()
		fieldInfo.attributeCount = attributesCount
		fieldInfo.attributes = make([]AttributeInfo, attributesCount)
		for i := uint16(0); i < attributesCount; i++ {
			 this.readAttribute(&fieldInfo.attributes[i])
		}
		this.classfile.fields[i] = fieldInfo
	}
}

func (this *ClassReader) readMethods() {
	methodsCount := this.ReadUint16()
	this.classfile.methodsCount = methodsCount
	this.classfile.methods = make([]MethodInfo, methodsCount)
	for i := uint16(0); i < methodsCount; i++ {
		methodInfo := MethodInfo{}
		methodInfo.accessFlags = this.ReadUint16()
		methodInfo.nameIndex = this.ReadUint16()
		methodInfo.descriptorIndex = this.ReadUint16()
		var attributesCount = this.ReadUint16()
		methodInfo.attributeCount = attributesCount
		methodInfo.attributes = make([]AttributeInfo, attributesCount)
		for i := uint16(0); i < attributesCount; i++ {
			this.readAttribute(&methodInfo.attributes[i])
		}
		this.classfile.methods[i] = methodInfo
	}
}


func (this *ClassReader) readAttributes() {
	attributesCount := this.ReadUint16()
	this.classfile.attributes = make([]AttributeInfo, attributesCount)
	for i := uint16(0); i < attributesCount; i++ {
		this.readAttribute(&this.classfile.attributes[i])
	}
}

func (this *ClassReader) readAttribute(attributeInfo *AttributeInfo)  {
	attributeNameIndex := this.ReadUint16()
	attributeLength := this.ReadUint32()

	attributeName := this.classfile.cpUtf8(attributeNameIndex)
	switch attributeName {
	case "Code":
		*attributeInfo = &CodeAttribute{attributeNameIndex: attributeNameIndex, attributeLength: attributeLength}
	case "SourceFile":
		*attributeInfo = &SourceFileAttribue{attributeNameIndex: attributeNameIndex, attributeLength: attributeLength}
	case "LineNumberTableEntry":
		*attributeInfo = &LineNumberTableAttribute{attributeNameIndex: attributeNameIndex, attributeLength: attributeLength}
	case "LocalVariableTableEntry":
		*attributeInfo = &LocalVariableTableAttribute{attributeNameIndex: attributeNameIndex, attributeLength: attributeLength}
	default:
		this.ReadBytes(uint32(attributeLength)) // just skip out
	}

	if *attributeInfo != nil {
		(*attributeInfo).ReadInfo(this)
	}
}

func (this *ClassReader) readConstantClassInfo(cp *ConstantClassInfo) {
	cp.nameIndex = this.ReadUint16()
}



func (this *ClassReader) readConstantFieldrefInfo(cp *ConstantFieldrefInfo) {
	cp.classIndex = this.ReadUint16()
	cp.nameAndTypeIndex = this.ReadUint16()
}



func (this *ClassReader) readConstantMethodrefInfo(cp *ConstantMethodrefInfo) {
	cp.classIndex = this.ReadUint16()
	cp.nameAndTypeIndex = this.ReadUint16()
}



func (this *ClassReader) readConstantStringInfo(cp *ConstantStringInfo) {
	cp.stringIndex = this.ReadUint16()
}



func (this *ClassReader) readConstantInterfaceMethodrefInfo(cp *ConstantInterfaceMethodrefInfo) {
	cp.classIndex = this.ReadUint16()
	cp.nameAndTypeIndex = this.ReadUint16()
}



func (this *ClassReader) readConstantIntegerInfo(cp *ConstantIntegerInfo) {
	cp.bytes = this.ReadUint32()
}



func (this *ClassReader) readConstantFloatInfo(cp *ConstantFloatInfo) {
	cp.bytes = this.ReadUint32()
}



func (this *ClassReader) readConstantLongInfo(cp *ConstantLongInfo) {
	cp.highBytes = this.ReadUint32()
	cp.lowBytes = this.ReadUint32()
}



func (this *ClassReader) readConstantDoubleInfo(cp *ConstantDoubleInfo) {
	cp.highBytes = this.ReadUint32()
	cp.lowBytes = this.ReadUint32()
}



func (this *ClassReader) readConstantNameAndTypeInfo(cp *ConstantNameAndTypeInfo) {
	cp.nameIndex = this.ReadUint16()
	cp.descriptorIndex = this.ReadUint16()
}



func (this *ClassReader) readConstantUtf8Info(cp *ConstantUtf8Info) {
	cp.length = this.ReadUint16()
	cp.bytes = this.ReadBytes(uint32(cp.length))
}



func (this *ClassReader) readConstantMethodHandleInfo(cp *ConstantMethodHandleInfo) {
	cp.referenceKind = this.ReadBytes(1)[0]
	cp.referenceIndex = this.ReadUint16()
}



func (this *ClassReader) readConstantMethodTypeInfo(cp *ConstantMethodTypeInfo) {
	cp.descriptorIndex = this.ReadUint16()
}



func (this *ClassReader) readConstantInvokeDynamicInfo(cp *ConstantInvokeDynamicInfo) {
	cp.bootstrapMethodAttrIndex = this.ReadUint16()
	cp.nameAndTypeIndex = this.ReadUint16()
}