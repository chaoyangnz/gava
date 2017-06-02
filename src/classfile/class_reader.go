package classfile

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

func (this *ClassReader) ReadBytes(len int) []byte {
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
	cpInfoCount := this.ReadUint16()
	//fmt.Printf("cp_info count: %d\n", cpInfoCount)
	this.classfile.constantPool = make([]ConstantPoolInfo, cpInfoCount)
	for i := uint16(1); i < cpInfoCount; i++ {
		//fmt.Printf(" #%2d = ", i)
		tag := this.ReadUint8()
		var cpInfo ConstantPoolInfo
		switch tag {
		case CONSTANT_Class:
			cpInfo = &ConstantClassInfo{}
		case CONSTANT_Fieldref:
			cpInfo = &ConstantFieldrefInfo{}
		case CONSTANT_Methodref:
			cpInfo = &ConstantMethodrefInfo{}
		case CONSTANT_InterfaceMethodref:
			cpInfo = &ConstantInterfaceMethodrefInfo{}
		case CONSTANT_String:
			cpInfo = &ConstantStringInfo{}
		case CONSTANT_Integer:
			cpInfo = &ConstantIntegerInfo{}
		case CONSTANT_Float:
			cpInfo = &ConstantFloatInfo{}
		case CONSTANT_Long:
			// occupy two entries
			cpInfo = &ConstantLongInfo{}
			this.classfile.constantPool[i] = cpInfo
			i++
		case CONSTANT_Double:
			// occupy two entries
			cpInfo = &ConstantDoubleInfo{}
			this.classfile.constantPool[i] = cpInfo
			i++
		case CONSTANT_NameAndType:
			cpInfo = &ConstantNameAndTypeInfo{}
		case CONSTANT_Utf8:
			cpInfo = &ConstantUtf8Info{}
		case CONSTANT_MethodHandle:
			cpInfo = &ConstantMethodHandleInfo{}
		case CONSTANT_MethodType:
			cpInfo = &ConstantMethodTypeInfo{}
		case CONSTANT_InvokeDynamic:
			cpInfo = &ConstantInvokeDynamicInfo{}
		default:
			// ignore
		}

		if cpInfo != nil {
			cpInfo.Read(this)
		}
		this.classfile.constantPool[i] = cpInfo
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
	this.classfile.fields = make([]FieldInfo, fieldsCount)
	for i := uint16(0); i < fieldsCount; i++ {
		fieldInfo := FieldInfo{}
		fieldInfo.accessFlags = this.ReadUint16()
		fieldInfo.nameIndex = this.ReadUint16()
		fieldInfo.descriptorIndex = this.ReadUint16()
		var attributesCount = this.ReadUint16()
		fieldInfo.attributes = make([]AttributeInfo, attributesCount)
		for i := uint16(0); i < attributesCount; i++ {
			fieldInfo.attributes[i] = this.ReadAsAttribute()
		}
		this.classfile.fields[i] = fieldInfo
	}
}

func (this *ClassReader) readMethods() {
	methodsCount := this.ReadUint16()
	this.classfile.methods = make([]MethodInfo, methodsCount)
	for i := uint16(0); i < methodsCount; i++ {
		methodInfo := MethodInfo{}
		methodInfo.accessFlags = this.ReadUint16()
		methodInfo.nameIndex = this.ReadUint16()
		methodInfo.descriptorIndex = this.ReadUint16()
		var attributesCount = this.ReadUint16()
		methodInfo.attributes = make([]AttributeInfo, attributesCount)
		for i := uint16(0); i < attributesCount; i++ {
			methodInfo.attributes[i] = this.ReadAsAttribute()
		}
		this.classfile.methods[i] = methodInfo
	}
}

func (this *ClassReader) read(reader *ClassReader) {

}


func (this *ClassReader) readAttributes() {
	attributesCount := this.ReadUint16()
	this.classfile.attributes = make([]AttributeInfo, attributesCount)
	for i := uint16(0); i < attributesCount; i++ {
		this.classfile.attributes[i] = this.ReadAsAttribute()
	}
}

func (this *ClassReader) ReadAsAttribute() AttributeInfo  {
	var attributeInfo AttributeInfo
	attributeNameIndex := this.ReadUint16()
	attributeLength := this.ReadUint32()
	var cpInfo = this.classfile.constantPool[attributeNameIndex]
	utf8 := cpInfo.(*ConstantUtf8Info)

	attributeName := string(utf8.bytes)
	switch attributeName {
	case "Code":
		attributeInfo = &CodeAttribute{}
	case "SourceFile":
		attributeInfo = &SourceFileAttribue{}
	case "LineNumberTable":
		attributeInfo = &LineNumberTableAttribute{}
	default:
		this.ReadBytes(int(attributeLength)) // just read out
	}

	if attributeInfo != nil {
		attributeInfo.ReadInfo(this)
	}

	return attributeInfo
}