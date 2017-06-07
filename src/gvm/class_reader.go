package gvm

import (
	"encoding/binary"
	"io/ioutil"
	"fmt"
)

var bigEndian = binary.BigEndian

type ClassReader struct {
	bytecode []uint8
	classfile *ClassFile
}

func NewClassReader(classfile string) *ClassReader {
	classpath := ""
	for i := 0; i < len(Classpath); i++ {
		classpath += Classpath[i]
		bytecode, err := ioutil.ReadFile(Classpath[i] + "/" + classfile)
		if err == nil {
			return &ClassReader{bytecode: bytecode}
		}
	}
	panic(fmt.Sprintf("Cannot find class %s in class path: %s", classfile, Classpath))
}

func (this *ClassReader) readU4() u4 {
	value := bigEndian.Uint32(this.bytecode[:4])
	this.bytecode = this.bytecode[4:]
	return u4(value)
}

func (this *ClassReader) readU2() u2 {
	value := bigEndian.Uint16(this.bytecode[:2])
	this.bytecode = this.bytecode[2:]
	return u2(value)
}

func (this *ClassReader) readU1() u1 {
	return u1(this.readU1s(1)[0])
}

func (this *ClassReader) readU1s(length uint32) []u1 {
	bytes := this.bytecode[:length]
	this.bytecode = this.bytecode[length:]
	return b2u(bytes)
}

func (this *ClassReader) length() int {
	return len(this.bytecode)
}

func (this *ClassReader) ReadAsClassFile() *ClassFile  {
	this.classfile = &ClassFile{}
	this.classfile.size = this.length()
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
	this.classfile.magic = this.readU4()
}

func (this *ClassReader) readMinorVersion() {
	this.classfile.minorVersion = this.readU2()
}

func (this *ClassReader) readMajorVersion() {
	this.classfile.majorVersion = this.readU2()
}

func (this *ClassReader) readConstantPool() {
	var constantPoolCount = this.readU2()
	this.classfile.constantPoolCount = constantPoolCount
	//fmt.Printf("cp_info count: %d\n", constantPoolCount)
	this.classfile.constantPool = make([]ConstantPoolInfo, constantPoolCount)
	for i := u2(1); i < constantPoolCount; i++ {
		//fmt.Printf(" #%2d = ", i)
		tag := this.readU1()

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
	this.classfile.accessFlags = this.readU2()
}

func (this *ClassReader) readThisClass() {
	this.classfile.thisClass = this.readU2()
}

func (this *ClassReader) readSuperClass() {
	this.classfile.superClass = this.readU2()
}

func (this *ClassReader) readInterfaces() {
	var interfacesCount = this.readU2()
	this.classfile.interfaces = make([]u2, interfacesCount)
	for i := u2(0); i < interfacesCount; i++ {
		this.classfile.interfaces[i] = this.readU2()
	}
}

func (this *ClassReader) readFields() {
	var fieldsCount = this.readU2()
	this.classfile.fieldsCount = fieldsCount
	this.classfile.fields = make([]FieldInfo, fieldsCount)
	for i := u2(0); i < fieldsCount; i++ {
		fieldInfo := FieldInfo{}
		fieldInfo.accessFlags = this.readU2()
		fieldInfo.nameIndex = this.readU2()
		fieldInfo.descriptorIndex = this.readU2()
		var attributesCount = this.readU2()
		fieldInfo.attributeCount = attributesCount
		fieldInfo.attributes = make([]AttributeInfo, attributesCount)
		for i := u2(0); i < attributesCount; i++ {
			 this.readAttribute(&fieldInfo.attributes[i])
		}
		this.classfile.fields[i] = fieldInfo
	}
}

func (this *ClassReader) readMethods() {
	methodsCount := this.readU2()
	this.classfile.methodsCount = methodsCount
	this.classfile.methods = make([]MethodInfo, methodsCount)
	for i := u2(0); i < methodsCount; i++ {
		methodInfo := MethodInfo{}
		methodInfo.accessFlags = this.readU2()
		methodInfo.nameIndex = this.readU2()
		methodInfo.descriptorIndex = this.readU2()
		var attributesCount = this.readU2()
		methodInfo.attributeCount = attributesCount
		methodInfo.attributes = make([]AttributeInfo, attributesCount)
		for i := u2(0); i < attributesCount; i++ {
			this.readAttribute(&methodInfo.attributes[i])
		}
		this.classfile.methods[i] = methodInfo
	}
}


func (this *ClassReader) readAttributes() {
	attributesCount := this.readU2()
	this.classfile.attributes = make([]AttributeInfo, attributesCount)
	for i := u2(0); i < attributesCount; i++ {
		this.readAttribute(&this.classfile.attributes[i])
	}
}

func (this *ClassReader) readAttribute(attributeInfo *AttributeInfo)  {
	attributeNameIndex := this.readU2()
	attributeLength := this.readU4()

	attributeName := this.classfile.cpUtf8(attributeNameIndex)
	switch attributeName {
	case "Code":
		*attributeInfo = &CodeAttribute{attributeNameIndex: attributeNameIndex, attributeLength: attributeLength}
	case "SourceFile":
		*attributeInfo = &SourceFileAttribue{attributeNameIndex: attributeNameIndex, attributeLength: attributeLength}
	case "LineNumberTable":
		*attributeInfo = &LineNumberTableAttribute{attributeNameIndex: attributeNameIndex, attributeLength: attributeLength}
	case "LocalVariableTable":
		*attributeInfo = &LocalVariableTableAttribute{attributeNameIndex: attributeNameIndex, attributeLength: attributeLength}
	default:
		this.readU1s(uint32(attributeLength)) // just skip out
	}

	if *attributeInfo != nil {
		(*attributeInfo).ReadInfo(this)
	}
}

func (this *ClassReader) readConstantClassInfo(cp *ConstantClassInfo) {
	cp.nameIndex = this.readU2()
}



func (this *ClassReader) readConstantFieldrefInfo(cp *ConstantFieldrefInfo) {
	cp.classIndex = this.readU2()
	cp.nameAndTypeIndex = this.readU2()
}



func (this *ClassReader) readConstantMethodrefInfo(cp *ConstantMethodrefInfo) {
	cp.classIndex = this.readU2()
	cp.nameAndTypeIndex = this.readU2()
}



func (this *ClassReader) readConstantStringInfo(cp *ConstantStringInfo) {
	cp.stringIndex = this.readU2()
}



func (this *ClassReader) readConstantInterfaceMethodrefInfo(cp *ConstantInterfaceMethodrefInfo) {
	cp.classIndex = this.readU2()
	cp.nameAndTypeIndex = this.readU2()
}



func (this *ClassReader) readConstantIntegerInfo(cp *ConstantIntegerInfo) {
	cp.bytes = this.readU4()
}



func (this *ClassReader) readConstantFloatInfo(cp *ConstantFloatInfo) {
	cp.bytes = this.readU4()
}



func (this *ClassReader) readConstantLongInfo(cp *ConstantLongInfo) {
	cp.highBytes = this.readU4()
	cp.lowBytes = this.readU4()
}



func (this *ClassReader) readConstantDoubleInfo(cp *ConstantDoubleInfo) {
	cp.highBytes = this.readU4()
	cp.lowBytes = this.readU4()
}



func (this *ClassReader) readConstantNameAndTypeInfo(cp *ConstantNameAndTypeInfo) {
	cp.nameIndex = this.readU2()
	cp.descriptorIndex = this.readU2()
}



func (this *ClassReader) readConstantUtf8Info(cp *ConstantUtf8Info) {
	cp.length = this.readU2()
	cp.bytes = this.readU1s(uint32(cp.length))
}



func (this *ClassReader) readConstantMethodHandleInfo(cp *ConstantMethodHandleInfo) {
	cp.referenceKind = this.readU1s(1)[0]
	cp.referenceIndex = this.readU2()
}



func (this *ClassReader) readConstantMethodTypeInfo(cp *ConstantMethodTypeInfo) {
	cp.descriptorIndex = this.readU2()
}



func (this *ClassReader) readConstantInvokeDynamicInfo(cp *ConstantInvokeDynamicInfo) {
	cp.bootstrapMethodAttrIndex = this.readU2()
	cp.nameAndTypeIndex = this.readU2()
}