package classfile

import (
	"fmt"
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
	size 	     int
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool []ConstantPoolInfo
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []FieldInfo
	methods      []MethodInfo
	attributes   []AttributeInfo
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
		fmt.Printf("\t%s", this.utf8(fieldInfo.nameIndex))
	}

	fmt.Printf("method: %d\n", len(this.methods))
	for i := 0; i < len(this.methods); i++  {
		methodInfo := this.methods[i]
		fmt.Printf("\t%s\n", this.utf8(methodInfo.nameIndex))
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
		fmt.Printf("\t\tSourceFile: %v\n", this.utf8(sourceFileAttribute.sourceFileIndex))
	case *LineNumberTableAttribute:
		lineNumberTableAttribute := attribute.(*LineNumberTableAttribute)
		fmt.Printf("\t\tlineNumberTableAttribute: %v\n", lineNumberTableAttribute)
	}
}

func (this *ClassFile) utf8(index uint16) string  {
	utf8 := this.constantPool[index].(*ConstantUtf8Info)
	return string(utf8.bytes)
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
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
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
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}


