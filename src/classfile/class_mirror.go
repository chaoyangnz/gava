package classfile

import (
	"strings"
)

type ClassMirror struct {
	minorVersion uint16
	majorVersion uint16
	constantPool []ConstantPoolInfo
	accessFlags  uint16
	thisClass    string
	superClass   string
	interfaces   []string
	fields       map[string]*FieldMirror
	methods      map[string]*MethodMirror
	//attributes   []Attribute
}

type Signature struct {
	name            string
	descriptor      string
}

func (this *Signature)toString() string  {
	return this.name + ":" + this.descriptor
}

func SignatureOf(str string) Signature  {
	arr := strings.Split(str, ":")
	return Signature{arr[0], arr[1]}
}

type FieldMirror struct {
	class           *ClassMirror
	signature       Signature
}

type MethodMirror struct {
	class           *ClassMirror
	signature       Signature
	code            []byte
	localVariables  []Signature
}



func NewClassMirror(classfile *ClassFile) *ClassMirror {
	classMirror := &ClassMirror{}
	classMirror.thisClass = classfile.cpClass(classfile.thisClass)
	classMirror.superClass = classfile.cpClass(classfile.superClass)
	classMirror.fields = make(map[string]*FieldMirror)
	for i := 0; i < len(classfile.fields); i++ {
		filedMirror := &FieldMirror{class: classMirror}
		fieldInfo := classfile.fields[i]
		filedMirror.signature = Signature{classfile.cpUtf8(fieldInfo.nameIndex), classfile.cpUtf8(fieldInfo.descriptorIndex)}
		classMirror.fields[filedMirror.signature.toString()] = filedMirror
	}

	classMirror.methods = make(map[string]*MethodMirror)
	for i := 0; i < len(classfile.methods); i++ {
		methodMirror := &MethodMirror{class: classMirror}
		methodInfo := classfile.methods[i]
		methodMirror.signature = Signature{classfile.cpUtf8(methodInfo.nameIndex), classfile.cpUtf8(methodInfo.descriptorIndex)}
		for j := uint16(0); j < methodInfo.attributeCount; j++ {
			attributeInfo := methodInfo.attributes[j]
			switch attributeInfo.(type) {
			case *CodeAttribute:
				codeAttribute := attributeInfo.(*CodeAttribute)
				methodMirror.code = codeAttribute.code
				for k := uint16(0); k < codeAttribute.attributesCount; k++ {
					codeAttributeAttribute := codeAttribute.attributes[i]
					switch codeAttributeAttribute.(type) {
					case *LocalVariableTableAttribute:
						localVariableTableAttribute := codeAttributeAttribute.(*LocalVariableTableAttribute)
						methodMirror.localVariables = make([]Signature, localVariableTableAttribute.localVariableTableLength)
						for m := uint16(0); m < localVariableTableAttribute.localVariableTableLength; i++ {
							methodMirror.localVariables[m] = Signature{
								classfile.cpUtf8(localVariableTableAttribute.localVariableTable[m].nameIndex),
								classfile.cpUtf8(localVariableTableAttribute.localVariableTable[m].descriptorIndex)}
						}
					}
				}
			}
		}
		classMirror.methods[methodMirror.signature.toString()] = methodMirror
	}

	return classMirror
}

func (this *ClassMirror) FindMethod(signature Signature) *MethodMirror {
	return this.methods[signature.toString()]
}
