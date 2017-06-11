package gvm

import (
	"io/ioutil"
)

type ClassLoader interface {
	load(classname string) *ClassType
}

var __indention = 0
func __times(t int, str string) string {
	ret := ""
	for i := 0; i < t; i++ {
		ret += str
	}
	return ret
}

func loadClass(classpath []string, classname string) *ClassType {
	//signature := classNameToSignature(classname)
	//classType, found := typeCache[signature]
	//if found {
	//	return classType.(*ClassType)
	//}
	__indention++
	trace(__times(__indention, "  ") + __times(50-2*__indention, "‧"))
	//pushTrace(ClassLoadTrace{'S', classname})
	trace(__times(__indention, "  ") + "⤈ %s", classname)

	// create classreader
	var classreader *ClassReader
	for x := 0; x < len(classpath); x++ {
		bytecode, err := ioutil.ReadFile(classpath[x] + "/" + classname + ".class")
		if err == nil {
			classreader = &ClassReader{bytecode: bytecode}
			break
		}
	}
	if classreader == nil {
		fatal("Cannot find class %s in class path: %s", classname, classpath)
	}

	// read as binary representation of class file
	classfile := classreader.Read()

	// create java class
	clazz := &ClassType{}
	// start loading
	clazz.constantPool = make([]RuntimeConstantPoolInfo, classfile.constantPoolCount)
	for i := u2(1); i < classfile.constantPoolCount; i++ {
		clazz.constantPool[i] = classfile.constantPool[i].runtime(clazz)
	}
	//for i := u2(1); i < classfile.constantPoolCount; i++ {
	//	runtimeConstantPoolInfo := this.constantPool[i]
	//	switch runtimeConstantPoolInfo.(type) {
	//	case *RuntimeConstantIntegerInfo, *RuntimeConstantLongInfo, *RuntimeConstantFloatInfo, *RuntimeConstantDoubleInfo,
	//		 *RuntimeConstantStringInfo, *RuntimeConstantUtf8Info, *RuntimeConstantNameAndTypeInfo, *RuntimeConstantMethodTypeInfo:
	//		runtimeConstantPoolInfo.resolve()
	//	}
	//}
	clazz.accessFlags = uint16(classfile.accessFlags)
	// resolve this class
	clazz.thisClass = uint16(classfile.thisClass)
	clazz.thisClassName = classfile.cpUtf8(classfile.constantPool[classfile.thisClass].(*ConstantClassInfo).nameIndex)
	typeCache[clazz.signature()] = clazz // immediately put into class cache to prevent forever loop
	clazz.constantPool[clazz.thisClass].resolve()

	// resolve super class
	clazz.superClass = uint16(classfile.superClass)
	if clazz.superClass != 0 {
		clazz.superClassName = classfile.cpUtf8(classfile.constantPool[classfile.superClass].(*ConstantClassInfo).nameIndex)
		clazz.constantPool[clazz.superClass].resolve()
	}

	clazz.fields = make([]*Field, len(classfile.fields))
	clazz.fieldsMap = make(map[string]*Field)
	clazz.staticFields = []j_any{}
	maxInstanceFieldIndex := uint16(0)
	maxStaticFieldIndex   := uint16(0)
	if clazz.superClass == 0 { // jdk/lang/Object
		clazz.instanceFieldsStart = 0
	} else {
		superClass := clazz.constantPool[clazz.superClass].resolve().(*RuntimeConstantClassInfo).referenceType.(*ClassType)
		clazz.instanceFieldsStart = superClass.instanceFieldsStart + uint16(len(superClass.instanceFileds))
	}
	for i := 0; i < len(classfile.fields); i++ {
		fieldInfo := classfile.fields[i]
		javaField := &Field{class: clazz,
			accessFlags:           uint16(fieldInfo.accessFlags),
			name:                  classfile.cpUtf8(fieldInfo.nameIndex),
			descriptor:            classfile.cpUtf8(fieldInfo.descriptorIndex)}
		clazz.fields[i] = javaField
		clazz.fieldsMap[javaField.name + javaField.descriptor] = javaField
		if javaField.isStatic() {
			javaField.index = maxStaticFieldIndex
			clazz.staticFields = append(clazz.staticFields, javaField.defaultValue()) // initialize static variables to default values
			maxStaticFieldIndex++
		} else {
			javaField.index = maxInstanceFieldIndex + clazz.instanceFieldsStart
			clazz.instanceFileds = append(clazz.instanceFileds, javaField)
			maxInstanceFieldIndex++
		}
	}


	clazz.methods = make([]*Method, len(classfile.methods))
	clazz.methodsMap = make(map[string]*Method)
	for i := 0; i < len(classfile.methods); i++ {
		methodInfo := &classfile.methods[i]
		javaMethod := &Method{class: clazz,
			accessFlags:             uint16(methodInfo.accessFlags),
			name:                    classfile.cpUtf8(methodInfo.nameIndex),
			descriptor:              classfile.cpUtf8(methodInfo.descriptorIndex)}

		javaMethod.parameterDescriptors, javaMethod.returnDescriptor = parametersAndReturn(javaMethod.descriptor)
		for j := u2(0); j < methodInfo.attributeCount; j++ {
			attributeInfo := methodInfo.attributes[j]
			switch attributeInfo.(type) {
			case *CodeAttribute:
				codeAttribute := attributeInfo.(*CodeAttribute)
				javaMethod.maxStack = uint16(codeAttribute.maxStack)
				javaMethod.maxLocals = uint16(codeAttribute.maxLocals)
				javaMethod.code = u2b(codeAttribute.code)
				for k := u2(0); k < codeAttribute.attributesCount; k++ {
					codeAttributeAttribute := codeAttribute.attributes[k]
					switch codeAttributeAttribute.(type) {
					case *LocalVariableTableAttribute:
						localVariableTableAttribute := codeAttributeAttribute.(*LocalVariableTableAttribute)
						javaMethod.localVariables = make([]LocalVariable, localVariableTableAttribute.localVariableTableLength)
						for m := u2(0); m < localVariableTableAttribute.localVariableTableLength; m++ {
							javaMethod.localVariables[m] = LocalVariable{
								javaMethod,
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
		clazz.methods[i] = javaMethod
		clazz.methodsMap[javaMethod.name + javaMethod.descriptor] = javaMethod
	}

	clazz.classObject = newJavaLangClass()

	trace(__times(__indention, "  ") + "⤉ %s", classname)
	trace(__times(__indention, "  ") + __times(50-2*__indention, "‧"))
	__indention--

	return clazz
}

type BootstrapClassLoader struct {}

/**
this method should only be called by ofClassType(className)
 */
func (this *BootstrapClassLoader) load(classname string) *ClassType {
	class := loadClass(coreClassPath, classname)
	class.classLoader = java_lang_classloader{object_null} // nil for bootstrap loader
	return class
}