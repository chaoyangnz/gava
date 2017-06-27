package jago

import (
	"unsafe"
	"strings"
)

type ClassLoader struct {
	classPath     *ClassPath
	classCache    map[string]*Class
	parent        *ClassLoader
}

func NewClassLoader(str string, parent *ClassLoader) *ClassLoader {
	return &ClassLoader{
		NewClassPath(str),
		make(map[string]*Class),
		parent}
}

func (this *ClassLoader) CreateClass(className string) *Class {
	if class, found := this.classCache[className]; found {
		return class
	}

	__indention++
	if __indention == 1 {
		Trace("\n")
	}

	Trace(__times(__indention, "  ") + __times(50-2*__indention, "‧") + "\n")
	Trace(__times(__indention, "  ") + "↳ %s \n", className)

	var class *Class
	if string(className[0]) == JVM_SIGNATURE_ARRAY {
		class = this.createArrayClass(className)
	} else {
		clazz := this.loadClass(className)
		// eager linkage
		this.link(clazz)

		// attach a java.lang.Class object
		clazz.classObject = NewJavaLangClass()

		class = clazz
	}

	Trace(__times(__indention, "  ") + "↱ %s \n", className)
	Trace(__times(__indention, "  ") + __times(50-2*__indention, "‧") + "\n")
	__indention--
	return class
}

var __indention = 0
func __times(t int, str string) string {
	ret := ""
	for i := 0; i < t; i++ {
		ret += str
	}
	return ret
}

func (this *ClassLoader) createArrayClass(className string) *Class {

	arrayClass := &Class{
			name: className,
			superClassName: "java/lang/Object",
			interfaceNames: []string{"java/io/Serializable", "java/lang/Cloneable"}}

	this.classCache[className] = arrayClass

	arrayClass.accessFlags = 0
	arrayClass.superClass = BOOTSTRAP_CLASSLOADER.CreateClass("java/lang/Object")
	arrayClass.interfaces = []*Class{
		BOOTSTRAP_CLASSLOADER.CreateClass("java/io/Serializable"),
		BOOTSTRAP_CLASSLOADER.CreateClass("java/lang/Cloneable")}

	componentDescriptor := string(className[1])
	switch componentDescriptor {
	case JVM_SIGNATURE_BYTE:
		{
			arrayClass.componentType = BYTE_TYPE
			arrayClass.elementType = BYTE_TYPE
		}
	case JVM_SIGNATURE_CHAR:
		{
			arrayClass.componentType = CHAR_TYPE
			arrayClass.elementType = CHAR_TYPE
		}
	case JVM_SIGNATURE_SHORT:
		{
			arrayClass.componentType = SHORT_TYPE
			arrayClass.elementType = SHORT_TYPE
		}
	case JVM_SIGNATURE_INT:
		{
			arrayClass.componentType = INT_TYPE
			arrayClass.elementType = INT_TYPE
		}
	case JVM_SIGNATURE_LONG:
		{
			arrayClass.componentType = LONG_TYPE
			arrayClass.elementType = LONG_TYPE
		}
	case JVM_SIGNATURE_FLOAT:
		{
			arrayClass.componentType = FLOAT_TYPE
			arrayClass.elementType = FLOAT_TYPE
		}
	case JVM_SIGNATURE_DOUBLE:
		{
			arrayClass.componentType = DOUBLE_TYPE
			arrayClass.elementType = DOUBLE_TYPE
		}
	case JVM_SIGNATURE_BOOLEAN:
		{
			arrayClass.componentType = BOOLEAN_TYPE
			arrayClass.elementType = BOOLEAN_TYPE
		}
	case JVM_SIGNATURE_CLASS:
		{
			arrayClass.componentType = BOOTSTRAP_CLASSLOADER.CreateClass(className[2:len(className)-1])
			arrayClass.elementType = arrayClass.componentType
		}
	case JVM_SIGNATURE_ARRAY:
		{
			arrayClass.componentType = BOOTSTRAP_CLASSLOADER.CreateClass(className[1:])
			arrayClass.elementType = arrayClass.componentType.(*Class).elementType
		}
	}
	dimensions := 1
	class := arrayClass
	for ;; {
		componentType, ok := class.componentType.(*Class)
		if ok {
			class = componentType
			dimensions++
		} else {
			break
		}
	}
	arrayClass.dimensions = dimensions

	// set classloader
	class.classLoader = this
	class.classObject = NewJavaLangClass()
	return arrayClass
}

func (this *ClassLoader) loadClass(className string) *Class  {
	bytecode := this.readClass(className)

	//If L creates C directly, we say that L defines C
	class := this.defineClass(bytecode)

	// TODO delegation

	// set classloader
	class.classLoader = this
	return class
}

func (this *ClassLoader) readClass(className string) []byte  {
	bytecode, err := this.classPath.ReadClass(className)
	if err != nil {
		Throw("java.lang.ClassNotFoundException", className)
	}
	return bytecode
}

func (this *ClassLoader) defineClass(bytecode []byte) *Class  {
	classfile := NewClassReader(bytecode).Read()

	class := &Class{}

	class.accessFlags = uint16(classfile.accessFlags)
	class.name = classfile.cpUtf8(classfile.constantPool[classfile.thisClass].(*ConstantClassInfo).nameIndex)
	// add to classcache
	this.classCache[class.name] = class

	if classfile.superClass == 0 {
		class.superClassName = ""
	} else {
		class.superClassName = classfile.cpUtf8(classfile.constantPool[classfile.superClass].(*ConstantClassInfo).nameIndex)
	}

	//class.superClass = ??
	class.interfaceNames = make([]string, len(classfile.interfaces))
	for i, interfaceIndex := range classfile.interfaces {
		class.interfaceNames[i] = classfile.cpUtf8(classfile.constantPool[interfaceIndex].(*ConstantClassInfo).nameIndex)
	}
	//class.interfaces = ??

	constantPool := make([]Constant, classfile.constantPoolCount+1)
	class.constantPool = constantPool

	// start loading
	for i, _ := range classfile.constantPool {
		constInfo := classfile.constantPool[i]
		var constant Constant
		switch constInfo.(type) {
		case *ConstantClassInfo:
			constantClassInfo := constInfo.(*ConstantClassInfo)
			constant = &ClassRef{
				class,
				classfile.cpUtf8(constantClassInfo.nameIndex),
				nil}
		case *ConstantFieldrefInfo:
			constantFieldrefInfo := constInfo.(*ConstantFieldrefInfo)
			nameAndTypeInfo := classfile.constantPool[constantFieldrefInfo.nameAndTypeIndex].(*ConstantNameAndTypeInfo)
			name := classfile.cpUtf8(nameAndTypeInfo.nameIndex)
			descriptor := classfile.cpUtf8(nameAndTypeInfo.descriptorIndex)
			className := classfile.cpUtf8(classfile.constantPool[constantFieldrefInfo.classIndex].(*ConstantClassInfo).nameIndex)
			constant = &FieldRef{
				MemberRef{
					class,
					className,
					nil,
					name,
				descriptor},
				nil}
		case *ConstantMethodrefInfo:
			constantMethodrefInfo := constInfo.(*ConstantMethodrefInfo)
			nameAndTypeInfo := classfile.constantPool[constantMethodrefInfo.nameAndTypeIndex].(*ConstantNameAndTypeInfo)
			name := classfile.cpUtf8(nameAndTypeInfo.nameIndex)
			descriptor := classfile.cpUtf8(nameAndTypeInfo.descriptorIndex)
			className := classfile.cpUtf8(classfile.constantPool[constantMethodrefInfo.classIndex].(*ConstantClassInfo).nameIndex)
			constant = &MethodRef{
				MemberRef{
					class,
					className,
					nil,
					name,
					descriptor},
				nil}
		case *ConstantInterfaceMethodrefInfo:
			constantInterfaceMethodrefInfo := constInfo.(*ConstantInterfaceMethodrefInfo)
			nameAndTypeInfo := classfile.constantPool[constantInterfaceMethodrefInfo.nameAndTypeIndex].(*ConstantNameAndTypeInfo)
			name := classfile.cpUtf8(nameAndTypeInfo.nameIndex)
			descriptor := classfile.cpUtf8(nameAndTypeInfo.descriptorIndex)
			className := classfile.cpUtf8(classfile.constantPool[constantInterfaceMethodrefInfo.classIndex].(*ConstantClassInfo).nameIndex)
			constant = &InterfaceMethodRef{
				MemberRef{
					class,
					className,
					nil,
					name,
					descriptor},
				nil}
		case *ConstantUtf8Info:
			constantUtf8Info := constInfo.(*ConstantUtf8Info)
			constant = &UTF8Constant{class,u2s(constantUtf8Info.bytes)}
		case *ConstantStringInfo:
			constantStringInfo := constInfo.(*ConstantStringInfo)
			constant = &StringConstant{class, classfile.cpUtf8(constantStringInfo.stringIndex), NULL}
		case *ConstantIntegerInfo:
			constantIntegerInfo := constInfo.(*ConstantIntegerInfo)
			constant = &IntegerConstant{class, Int(constantIntegerInfo.bytes)}
		case *ConstantFloatInfo:
			constantFloatInfo := constInfo.(*ConstantFloatInfo)
			constant = &FloatConstant{class, Float(constantFloatInfo.bytes)}
		case *ConstantLongInfo:
			constantLongInfo := constInfo.(*ConstantLongInfo)
			constant = &LongConstant{class, Long(constantLongInfo.highBytes << 32 | constantLongInfo.lowBytes)}
		case *ConstantDoubleInfo:
			constantDoubleInfo := constInfo.(*ConstantDoubleInfo)
			constant = &DoubleConstant{class, Double(constantDoubleInfo.highBytes << 32 | constantDoubleInfo.lowBytes)}
		case *ConstantNameAndTypeInfo:
			constantNameAndTypeInfo := constInfo.(*ConstantNameAndTypeInfo)
			constant = &NameAndTypeConstant{class,classfile.cpUtf8(constantNameAndTypeInfo.nameIndex),classfile.cpUtf8(constantNameAndTypeInfo.descriptorIndex)}
		case *ConstantMethodTypeInfo:
			constantMethodTypeInfo := constInfo.(*ConstantMethodTypeInfo)
			constant = &MethodTypeConstant{class,classfile.cpUtf8(constantMethodTypeInfo.descriptorIndex)}
		case *ConstantInvokeDynamicInfo:
			constantInvokeDynamicInfo := constInfo.(*ConstantInvokeDynamicInfo)
			nameAndTypeInfo := classfile.constantPool[constantInvokeDynamicInfo.nameAndTypeIndex].(*ConstantNameAndTypeInfo)
			name := classfile.cpUtf8(nameAndTypeInfo.nameIndex)
			descriptor := classfile.cpUtf8(nameAndTypeInfo.descriptorIndex)
			constant = &InvokeDynamicConstant{class,
				"",//TODO classfile.cpUtf8(constantInvokeDynamicInfo.bootstrapMethodAttrIndex),
				name,
				descriptor}
		}

		class.constantPool[i] = constant
	}


	class.fields = make([]*Field, len(classfile.fields))
	//class.staticVars = ??

	for i, fieldInfo := range classfile.fields {
		field := &Field{}
		field.accessFlags = uint16(fieldInfo.accessFlags)
		field.name = classfile.cpUtf8(fieldInfo.nameIndex)
		field.descriptor = classfile.cpUtf8(fieldInfo.descriptorIndex)
		field.class = class
		//field.index = ??
		class.fields[i] = field
	}

	class.methods = make([]*Method, len(classfile.methods))

	for i, methodInfo := range classfile.methods {
		method := &Method{}
		method.accessFlags = uint16(methodInfo.accessFlags)
		method.name = classfile.cpUtf8(methodInfo.nameIndex)
		method.descriptor = classfile.cpUtf8(methodInfo.descriptorIndex)
		method.class = class

		for j := u2(0); j < methodInfo.attributeCount; j++ {
			attributeInfo := methodInfo.attributes[j]
			switch attributeInfo.(type) {
			case *CodeAttribute:
				codeAttribute := attributeInfo.(*CodeAttribute)
				method.maxStack = uint(codeAttribute.maxStack)
				method.maxLocals = uint(codeAttribute.maxLocals)
				method.code = *(*[]uint8)(unsafe.Pointer(&codeAttribute.code))
				for k := u2(0); k < codeAttribute.attributesCount; k++ {
					codeAttributeAttribute := codeAttribute.attributes[k]
					switch codeAttributeAttribute.(type) {
					case *LocalVariableTableAttribute:
						localVariableTableAttribute := codeAttributeAttribute.(*LocalVariableTableAttribute)
						method.localVars = make([]*LocalVariable, localVariableTableAttribute.localVariableTableLength)
						for m := u2(0); m < localVariableTableAttribute.localVariableTableLength; m++ {
							method.localVars[m] = &LocalVariable{
								method,
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

		// start parse parameters and return
		arr := strings.Split(method.descriptor, ")")
		parametersStr := arr[0][1:]
		returnDescritor := arr[1]

		var parameterDescriptors []string

		for i := 0; i < len(parametersStr); {
			ch := rune(parametersStr[i])
			switch ch {
			case 'B', 'C', 'D', 'F', 'I', 'J', 'S', 'Z':
				parameterDescriptors = append(parameterDescriptors, string(ch))
				i++
			case 'L':
			Ref: for j := i+1; j < len(parametersStr); j++ {
				switch rune(parametersStr[j]) {
				case ';':
					parameterDescriptors = append(parameterDescriptors, string(parametersStr[i:j+1]))
					i = j+1
					break Ref
				}
			}
			case '[':
			Arr: for j := i+1; j < len(parametersStr); j++ {
				switch rune(parametersStr[j]) {
				case '[':
					continue
				case 'B', 'C', 'D', 'F', 'I', 'J', 'S', 'Z':
					parameterDescriptors = append(parameterDescriptors, string(parametersStr[i:j+1]))
					i = j+1
					break Arr
				case 'L':
					for k := j+1; j < len(parametersStr); k++ {
						switch rune(parametersStr[k]) {
						case ';':
							parameterDescriptors = append(parameterDescriptors, string(parametersStr[i:k+1]))
							i = k+1
							break Arr
						}
					}
				}
			}
			}
		}
		method.parameterDescriptors = parameterDescriptors
		method.returnDescriptor = returnDescritor

		class.methods[i] = method
	}


	// resolve super class
	if class.superClassName != "" {
		class.superClass = this.CreateClass(class.superClassName)
	}

	// calculate static variables and instance variable count
	// must be immediately after resolving super class
	maxInstanceVars := 0  // include all the ancestry
	maxStaticVars := 0

	if class.superClass != nil {
		maxInstanceVars = class.superClass.maxInstanceVars
	}
	for _, field := range class.fields {
		if field.isStatic() {
			field.index = maxStaticVars
			maxStaticVars++
		} else {
			field.index = maxInstanceVars
			maxInstanceVars++
		}
	}
	class.maxInstanceVars = maxInstanceVars
	class.maxStaticVars = maxStaticVars

	// resolve interfaces
	class.interfaces = make([]*Class, len(class.interfaceNames))
	for i, interfaceName := range class.interfaceNames {
		class.interfaces[i] = this.CreateClass(interfaceName)
	}

	return class
}

func (this *ClassLoader) link(class *Class)  {
	if class.linked {
		return
	}
	this.verify(class)
	this.prepare(class)

	class.linked = true
	// we resolve each symbolic class in a class or interface individually when it is used ("lazy" or "late" resolution)
	// So SymbolRef all implements a method PrimitiveType resolve()
	//for _, constant := range this.constantPool {
	//	switch constant.(type) {
	//	case SymbolRef:
	//		constant.(SymbolRef).resolve()
	//	}
	//}
	//this.resolve(class)
}

func (this *ClassLoader) verify(class *Class) {
	//TODO
}

// initialize static variables to default values: no need to execute code
func (this *ClassLoader) prepare(class *Class)  {
	// Initialize static variables
	class.staticVars = make([]Value, class.maxStaticVars)
	for _, field := range class.fields {
		if field.isStatic() {
			class.staticVars[field.index] = field.defaultValue()
		}
	}
}







