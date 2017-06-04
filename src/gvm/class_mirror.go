package gvm

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

func (this *MethodMirror) localVariablesSize() uint {
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



func NewClassMirror() *ClassMirror {
	return &ClassMirror{}
}

func (this *ClassMirror) Load(classfile *ClassFile)  {
	this.thisClass = classfile.cpClass(classfile.thisClass)
	this.superClass = classfile.cpClass(classfile.superClass)
	this.fields = make(map[string]*FieldMirror)
	for i := 0; i < len(classfile.fields); i++ {
		filedMirror := &FieldMirror{class: this}
		fieldInfo := classfile.fields[i]
		filedMirror.signature = Signature{classfile.cpUtf8(fieldInfo.nameIndex), classfile.cpUtf8(fieldInfo.descriptorIndex)}
		this.fields[filedMirror.signature.toString()] = filedMirror
	}

	this.methods = make(map[string]*MethodMirror)
	for i := 0; i < len(classfile.methods); i++ {
		methodMirror := &MethodMirror{class: this}
		methodInfo := &classfile.methods[i]
		methodMirror.signature = Signature{classfile.cpUtf8(methodInfo.nameIndex), classfile.cpUtf8(methodInfo.descriptorIndex)}
		for j := uint16(0); j < methodInfo.attributeCount; j++ {
			attributeInfo := methodInfo.attributes[j]
			switch attributeInfo.(type) {
			case *CodeAttribute:
				codeAttribute := attributeInfo.(*CodeAttribute)
				methodMirror.code = codeAttribute.code
				for k := uint16(0); k < codeAttribute.attributesCount; k++ {
					codeAttributeAttribute := codeAttribute.attributes[k]
					switch codeAttributeAttribute.(type) {
					case *LocalVariableTableAttribute:
						localVariableTableAttribute := codeAttributeAttribute.(*LocalVariableTableAttribute)
						methodMirror.localVariables = make([]Signature, localVariableTableAttribute.localVariableTableLength)
						for m := uint16(0); m < localVariableTableAttribute.localVariableTableLength; m++ {
							methodMirror.localVariables[m] = Signature{
								classfile.cpUtf8(localVariableTableAttribute.localVariableTable[m].nameIndex),
								classfile.cpUtf8(localVariableTableAttribute.localVariableTable[m].descriptorIndex)}
						}
					}
				}
			}
		}
		this.methods[methodMirror.signature.toString()] = methodMirror
	}
}

func (this *ClassMirror) FindMethod(signature string) *MethodMirror {
	return  this.methods[signature]
}
