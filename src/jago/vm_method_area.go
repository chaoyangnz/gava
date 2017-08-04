package jago

import (
	"strings"
	"math"
	"unsafe"
	"sync"
)

type NL struct {
	N string // binary class name, NO L; for non-array class
	L *Object
}

type MethodArea struct {
	DefinedClasses   map[NL]*Class // TODO lock
	InitiatedClasses map[NL]*Class
	StringPool
	*BootstrapClassLoader //bootstrap classloader
}

func (this MethodArea) GetInitiatedClass(className string, classLoader JavaLangClassLoader) (*Class, bool)  {
	if class, ok := this.InitiatedClasses[NL{className, classLoader.(Reference).oop}]; ok {
		return class, true
	}
	return nil, false
}

func (this MethodArea) SetInitiatedClass(className string, classLoader JavaLangClassLoader, class *Class) {
	VM.Info("-> Set initiated class (%s, %p): %p \n", className, classLoader.(Reference).oop, class)
	this.InitiatedClasses[NL{className, classLoader.(Reference).oop}] = class
}

func (this MethodArea) GetDefinedClass(className string, classLoader JavaLangClassLoader) (*Class, bool)  {
	if class, ok := this.DefinedClasses[NL{className, classLoader.(Reference).oop}]; ok {
		return class, true
	}
	return nil, false
}

func (this MethodArea) GetDefiningClassLoader(class *Class) JavaLangClassLoader {
	for cn, cl := range this.DefinedClasses {
		if cl == class && cn.N == class.name {
			return Reference{cn.L}
		}
	}
	return NULL
}

func (this MethodArea) GetInitiatingClassLoader(class *Class) JavaLangClassLoader {
	for cn, cl := range this.InitiatedClasses {
		if cl == class && cn.N == class.name {
			return Reference{cn.L}
		}
	}
	return NULL
}

func (this MethodArea) SetDefinedClass(className string, classLoader JavaLangClassLoader, class *Class) {
	VM.Info("-> Set defined class (%s, %p): %p \n", className, classLoader.(Reference).oop, class)
	this.DefinedClasses[NL{className, classLoader.(Reference).oop}] = class
}



var depth = 0

func (this *MethodArea) CreateClass0(N string, Ld JavaLangClassLoader, triggerReason *ClassTriggerReason) *Class {
	var C *Class
	if N[:1] != JVM_SIGNATURE_ARRAY {
		if Ld.IsNull() {
			// bootstrap classloader
			C = VM.LoadClass(N, triggerReason)
		} else {
			// user-defined classloader
			C = VM.LoadClassUd(N, Ld, triggerReason)
		}
	} else {
		C = this.createArrayClass(N, Ld, triggerReason)
	}

	VM.link(C) // eagerly link

	return C
}

func (this *MethodArea) CreateClass(N string, D *Class, triggerReason *ClassTriggerReason) *Class {
	var Ld JavaLangClassLoader
	if D == nil {
		Ld = NULL
	} else {
		Ld = VM.GetDefiningClassLoader(D)
	}

	return this.CreateClass0(N, Ld, triggerReason)
}

func (this *MethodArea) LoadClassUd(N string, L JavaLangClassLoader, triggerReason *ClassTriggerReason) *Class {
	if C, ok := VM.GetInitiatedClass(N, L); ok {
		return C
	}
	loadClassMethod := L.Class().FindMethod("loadClass", "(Ljava/lang/String;)Ljava/lang/Class;")
	VM.Info("==before java.lang.ClassLoader#loadClass %s in LoadClassUd \n", N)
	javaname := binaryName2JavaName(N)
	classObject := VM.InvokeMethod(loadClassMethod, L, javaname).(JavaLangClass)

	VM.Info("==after java.lang.ClassLoader#loadClass %s %s in LoadClassUd jc=%p \n", N, classObject.Class().name, classObject.(Reference).oop)
	C := classObject.retrieveType().(*Class)
	//
	//VM.SetInitiatedClass(N, L, C)
	return C
}

func componentAndElementTypeName(arrayClassName string) (string, string, int) {
	componentDescriptor := string(arrayClassName[1])
	switch componentDescriptor {
	case JVM_SIGNATURE_BYTE,
		JVM_SIGNATURE_CHAR,
		JVM_SIGNATURE_SHORT,
		JVM_SIGNATURE_INT,
		JVM_SIGNATURE_LONG,
		JVM_SIGNATURE_FLOAT,
		JVM_SIGNATURE_DOUBLE,
		JVM_SIGNATURE_BOOLEAN:
		{
			return componentDescriptor, componentDescriptor, 1
		}
	case JVM_SIGNATURE_CLASS:
		{
			className := arrayClassName[2:len(arrayClassName)-1]
			return className, className, 1
		}
	case JVM_SIGNATURE_ARRAY:
		{
			componentArrayClassName := arrayClassName[1:]
			_, elementTypeName, dimension := componentAndElementTypeName(componentArrayClassName)
			return componentArrayClassName, elementTypeName, dimension+1
		}
	}
	VM.Throw("java/lang/IllegalArgumentException", "%s is not a array class name")
	return "", "", -1
}

func (this *MethodArea) createArrayClass(N string, L JavaLangClassLoader, triggerReason *ClassTriggerReason) *Class {
	if C, ok := VM.GetInitiatedClass(N, L); ok {
		return C //???
		// TODO If L has already been recorded as an intiating loader of an array class with the same component type as N,
		// the class is C. No array class creation is neccessry
	}

	if C, ok := VM.GetDefinedClass(N, NULL); ok { // this lookup is not mentioned in jvms
		return C
	}

	// ------------------
	VM.BootstrapClassLoader.Debug(repeat("\t", depth) + "↳ %s ", N)
	VM.BootstrapClassLoader.Debug("(reason: %s", triggerReason.code)
	VM.BootstrapClassLoader.Trace(" - %s", triggerReason.desc)
	VM.BootstrapClassLoader.Debug(") ")
	var cll string
	if L.IsNull()  {
		cll = "<bootstrap>"
	} else {
		cll = L.Class().name
	}
	VM.BootstrapClassLoader.Debug("[%s] (array)\n", cll)
	depth++
	// ------------------

	C := &Class{
		name:           N,
		superClassName: "java/lang/Object",
		interfaceNames: []string{"java/io/Serializable", "java/lang/Cloneable"},
		LC: sync.NewCond(&sync.Mutex{})}

	C.superClass = VM.CreateClass0("java/lang/Object", NULL, TRIGGER_BY_AS_SUPERCLASS)
	C.interfaces = []*Class{
		VM.CreateClass0("java/io/Serializable", NULL, TRIGGER_BY_AS_SUPERINTERFACE),
		VM.CreateClass0("java/lang/Cloneable", NULL, TRIGGER_BY_AS_SUPERINTERFACE)}


	componentDescriptor := string(N[1])
	switch componentDescriptor {
	case JVM_SIGNATURE_BYTE:
		{
			C.componentType = BYTE_TYPE
			C.elementType = BYTE_TYPE
			C.dimensions = 1
		}
	case JVM_SIGNATURE_CHAR:
		{
			C.componentType = CHAR_TYPE
			C.elementType = CHAR_TYPE
			C.dimensions = 1
		}
	case JVM_SIGNATURE_SHORT:
		{
			C.componentType = SHORT_TYPE
			C.elementType = SHORT_TYPE
			C.dimensions = 1
		}
	case JVM_SIGNATURE_INT:
		{
			C.componentType = INT_TYPE
			C.elementType = INT_TYPE
			C.dimensions = 1
		}
	case JVM_SIGNATURE_LONG:
		{
			C.componentType = LONG_TYPE
			C.elementType = LONG_TYPE
			C.dimensions = 1
		}
	case JVM_SIGNATURE_FLOAT:
		{
			C.componentType = FLOAT_TYPE
			C.elementType = FLOAT_TYPE
			C.dimensions = 1
		}
	case JVM_SIGNATURE_DOUBLE:
		{
			C.componentType = DOUBLE_TYPE
			C.elementType = DOUBLE_TYPE
			C.dimensions = 1
		}
	case JVM_SIGNATURE_BOOLEAN:
		{
			C.componentType = BOOLEAN_TYPE
			C.elementType = BOOLEAN_TYPE
			C.dimensions = 1
		}
	case JVM_SIGNATURE_CLASS:
		{
			componentTypeName := N[2:len(N)-1]
			C.componentType = VM.CreateClass0(componentTypeName, L, TRIGGER_BY_AS_ARRAY_COMPONENT)
			C.elementType = C.componentType
			C.dimensions = 1
		}
	case JVM_SIGNATURE_ARRAY:
		{
			componentTypeName := N[2:len(N)-1]
			componentArrayClass := VM.CreateClass0(componentTypeName, L, TRIGGER_BY_AS_ARRAY_COMPONENT)
			C.componentType = componentArrayClass
			C.elementType = componentArrayClass.elementType
			C.dimensions = componentArrayClass.dimensions + 1
		}
	}

	switch C.componentType.(type) {
	case *Class:
		componentClass := C.componentType.(*Class)
		C.accessFlags = componentClass.accessFlags
		VM.SetDefinedClass(N, VM.GetDefiningClassLoader(componentClass), C)
	default:
		C.accessFlags = JVM_ACC_PUBLIC
		VM.SetDefinedClass(N, NULL, C)
	}

	VM.SetInitiatedClass(N, L, C)
	C.initialized = &InitializeSate{UNINITIALIZED, nil}
	C.classObject = VM.NewJavaLangClass(C)

	depth--

	return C
}

// jvms 5.4.3.1
func (this *MethodArea) resolveClass(N string, D *Class, triggerReason *ClassTriggerReason) *Class {
	C := VM.CreateClass(N, D, triggerReason)
	if C.IsArray() {
		_, elementTypeName, _ := componentAndElementTypeName(C.Name())
		 if elementTypeName[:1] == JVM_SIGNATURE_CLASS { // cannot be array for an element type
			 C.elementType = this.resolveClass(elementTypeName, D, TRIGGER_BY_AS_ARRAY_COMPONENT)
		 }
	}

	return C
}


// jvms 5.3.5
func (this *MethodArea) deriveClass(N string, L JavaLangClassLoader, bytecode []byte, triggerReason *ClassTriggerReason) *Class  {

	if _, ok := VM.GetInitiatedClass(N, L); ok {
		VM.Throw("java/lang/LinkageError", "Classloader %s is an initiating loader of class %s", L.Class().Name(), N)
	}

	// ------------------
	VM.BootstrapClassLoader.Debug(repeat("\t", depth) + "↳ %s ", N)
	VM.BootstrapClassLoader.Debug("(reason: %s", triggerReason.code)
	VM.BootstrapClassLoader.Trace(" - %s", triggerReason.desc)
	VM.BootstrapClassLoader.Debug(") ")
	var cll string
	if L.IsNull()  {
		cll = "<bootstrap>"
	} else {
		cll = L.Class().name
	}
	VM.BootstrapClassLoader.Debug("[%s] (derive)\n", cll)
	depth++
	// ------------------

	classfile := &ClassFile{}
	classfile.read(bytecode)

	// TODO purported representation is not of a supported major or minor version -> java/lang.UnsupportedClassVersionError

	// TODO class file does not atually represent a class name N -> java/langNoClassDefFoundError

	C := &Class{LC: sync.NewCond(&sync.Mutex{})}

	C.accessFlags = uint16(classfile.accessFlags)
	C.name = classfile.cpUtf8(classfile.constantPool[classfile.thisClass].(*ConstantClassInfo).nameIndex)

	// constant pool
	constantPool := make([]Constant, classfile.constantPoolCount+1)
	C.constantPool = constantPool

	for i, _ := range classfile.constantPool {
		constInfo := classfile.constantPool[i]
		var constant Constant
		switch constInfo.(type) {
		case *ConstantClassInfo:
			constantClassInfo := constInfo.(*ConstantClassInfo)
			constant = &ClassRef{
				C,
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
					C,
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
					C,
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
					C,
					className,
					nil,
					name,
					descriptor},
				nil}
		case *ConstantUtf8Info:
			constantUtf8Info := constInfo.(*ConstantUtf8Info)
			constant = &UTF8Constant{C,u2s(constantUtf8Info.bytes)}
		case *ConstantStringInfo:
			constantStringInfo := constInfo.(*ConstantStringInfo)
			constant = &StringConstant{C, classfile.cpUtf8(constantStringInfo.stringIndex), NULL}
		case *ConstantIntegerInfo:
			constantIntegerInfo := constInfo.(*ConstantIntegerInfo)
			constant = &IntegerConstant{C, Int(constantIntegerInfo.bytes)}
		case *ConstantFloatInfo:
			constantFloatInfo := constInfo.(*ConstantFloatInfo)
			constant = &FloatConstant{C, Float(constantFloatInfo.bytes)}
		case *ConstantLongInfo:
			constantLongInfo := constInfo.(*ConstantLongInfo)
			l := int64(uint64(constantLongInfo.highBytes) << 32 + uint64(constantLongInfo.lowBytes))
			constant = &LongConstant{C, Long(l)}
		case *ConstantDoubleInfo:
			constantDoubleInfo := constInfo.(*ConstantDoubleInfo)
			bits := uint64(constantDoubleInfo.highBytes) << 32 + uint64(constantDoubleInfo.lowBytes)
			d := math.Float64frombits(bits)
			constant = &DoubleConstant{C, Double(d)}
		case *ConstantNameAndTypeInfo:
			constantNameAndTypeInfo := constInfo.(*ConstantNameAndTypeInfo)
			constant = &NameAndTypeConstant{C,classfile.cpUtf8(constantNameAndTypeInfo.nameIndex),classfile.cpUtf8(constantNameAndTypeInfo.descriptorIndex)}
		case *ConstantMethodTypeInfo:
			constantMethodTypeInfo := constInfo.(*ConstantMethodTypeInfo)
			constant = &MethodTypeConstant{C,classfile.cpUtf8(constantMethodTypeInfo.descriptorIndex)}
		case *ConstantInvokeDynamicInfo:
			constantInvokeDynamicInfo := constInfo.(*ConstantInvokeDynamicInfo)
			nameAndTypeInfo := classfile.constantPool[constantInvokeDynamicInfo.nameAndTypeIndex].(*ConstantNameAndTypeInfo)
			name := classfile.cpUtf8(nameAndTypeInfo.nameIndex)
			descriptor := classfile.cpUtf8(nameAndTypeInfo.descriptorIndex)
			constant = &InvokeDynamicConstant{C,
			                                  "",//TODO classfile.cpUtf8(constantInvokeDynamicInfo.bootstrapMethodAttrIndex),
			                                  name,
			                                  descriptor}
		}

		C.constantPool[i] = constant
	}

	for k := u2(0); k < classfile.attributeCount; k++ {
		attributeInfo := classfile.attributes[k]
		switch attributeInfo.(type) {
		case *SourceFileAttribue:
			sourceFileAttribue := attributeInfo.(*SourceFileAttribue)
			C.sourceFile = classfile.cpUtf8(sourceFileAttribue.sourceFileIndex)
		}
	}


	C.fields = make([]*Field, len(classfile.fields))
	//class.staticVars = ??

	for i, fieldInfo := range classfile.fields {
		field := &Field{}
		field.accessFlags = uint16(fieldInfo.accessFlags)
		field.name = classfile.cpUtf8(fieldInfo.nameIndex)
		field.descriptor = classfile.cpUtf8(fieldInfo.descriptorIndex)
		field.class = C
		//field.index = ??
		C.fields[i] = field
	}

	C.methods = make([]*Method, len(classfile.methods))

	for i, methodInfo := range classfile.methods {
		method := &Method{}
		method.accessFlags = uint16(methodInfo.accessFlags)
		method.name = classfile.cpUtf8(methodInfo.nameIndex)
		method.descriptor = classfile.cpUtf8(methodInfo.descriptorIndex)
		method.class = C

		for j := u2(0); j < methodInfo.attributeCount; j++ {
			attributeInfo := methodInfo.attributes[j]
			switch attributeInfo.(type) {
			case *CodeAttribute:
				codeAttribute := attributeInfo.(*CodeAttribute)
				method.maxStack = uint(codeAttribute.maxStack)
				method.maxLocals = uint(codeAttribute.maxLocals)
				method.code = *(*[]uint8)(unsafe.Pointer(&codeAttribute.code))
				method.exceptions = make([]*ExceptionHandler, codeAttribute.exceptionTableLength)
				for ei, ete := range codeAttribute.exceptionTable {
					exceptionHandler := &ExceptionHandler{int(ete.startPc), int(ete.endPc), int(ete.handlerPc), int(ete.catchType)}
					method.exceptions[ei] = exceptionHandler
				}
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
					case *LineNumberTableAttribute:
						lineNumberTableAttribute := codeAttributeAttribute.(*LineNumberTableAttribute)
						method.lineNumbers = make([]*LineNumber, lineNumberTableAttribute.lineNumberTableLength)
						for i, lineNumberTableEntry := range lineNumberTableAttribute.lineNumberTable {
							method.lineNumbers[i] = &LineNumber{int(lineNumberTableEntry.startPc), int(lineNumberTableEntry.lineNumber)}
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

		C.methods[i] = method
	}

	if classfile.superClass == 0 {
		C.superClassName = ""
	} else {
		C.superClassName = classfile.cpUtf8(classfile.constantPool[classfile.superClass].(*ConstantClassInfo).nameIndex)

		// if C has a direct superclass, the symbolic reference from C to its direct superclass is resolved using the algorithm of 5.4.3.1
		superClass := VM.resolveClass(C.superClassName, C, TRIGGER_BY_AS_SUPERCLASS)
		// if the class or interface name as direct superclass of C is in fact an interface -> java/lang/IncompatibleClassChangeError
		if superClass.IsInterface() {
			VM.Throw("java/lang/IncompatibleClassChangeError", "Superclass of %s cannot be interface", N)
		}
		// if any of the superclasses of C is C itself -> java/lang/ClassCircularityError
		if superClass.Name() == N || superClass == C {
			VM.Throw("java/lang/ClassCircularityError", "Superclass of %s is itself", N)
		}
		C.superClass = superClass
	}

	C.interfaceNames = make([]string, len(classfile.interfaces))
	C.interfaces = make([]*Class, len(C.interfaceNames))
	for i, interfaceIndex := range classfile.interfaces {
		superinterfaceName := classfile.cpUtf8(classfile.constantPool[interfaceIndex].(*ConstantClassInfo).nameIndex)
		C.interfaceNames[i] = superinterfaceName

		// if C has any direct superinterfaces, the symbolic references from c to its direct superinterfaces are resolved using the algorithm of 5.4.3.1
		superinterface := VM.resolveClass(superinterfaceName, C, TRIGGER_BY_AS_SUPERINTERFACE)
		// if the class or interface name as direct superinterface of C is not in fact an interface -> java/lang/IncompatibleClassChangeError
		if !superinterface.IsInterface() {
			VM.Throw("java/lang/IncompatibleClassChangeError", "Superinterface of %s must be interface", N)
		}
		// if any of the superinterfaces of C is C itself -> java/lang/ClassCircularityError
		if superinterface.Name() == N || superinterface == C {
			VM.Throw("java/lang/ClassCircularityError", "Superinterface of %s is itself", N)
		}
		C.interfaces[i] = superinterface
	}


	// calculate static variables and instance variable count
	// must be immediately after resolving super class
	maxInstanceVars := 0  // include all the ancestry
	maxStaticVars := 0

	instanceVarFields := make([]*Field, 0)
	staticVarFields := make([]*Field, 0)


	if C.superClass != nil {
		maxInstanceVars = C.superClass.maxInstanceVars

		instanceVarFields = append(instanceVarFields, C.superClass.instanceVarFields...)
	}
	for _, field := range C.fields {
		if field.IsStatic() {
			field.slot = maxStaticVars
			staticVarFields = append(staticVarFields, field)
			maxStaticVars++
		} else {
			field.slot = maxInstanceVars
			instanceVarFields = append(instanceVarFields, field)
			maxInstanceVars++
		}
	}
	C.maxInstanceVars = maxInstanceVars
	C.maxStaticVars = maxStaticVars

	C.instanceVarFields = instanceVarFields
	C.staticVarFields = staticVarFields

	// marks C as having L as its defining class loader and records that L is an initiating loader of C
	VM.SetDefinedClass(N, L, C)
	VM.SetInitiatedClass(N, L, C)

	C.initialized = &InitializeSate{UNINITIALIZED, nil}
	C.classObject = VM.NewJavaLangClass(C)

	depth--

	return C
}

func (this *MethodArea) link(class *Class)  {
	if class.linked {
		return
	}
	VM.verify(class)
	VM.prepare(class)

	class.linked = true
	// we resolve each symbolic class in a class or interface individually when it is used ("lazy" or "late" resolution)
	// So SymbolRef all implements a method PrimitiveType resolve()
}

func (this *MethodArea) initializePrimitives()  {
	BYTE_TYPE.classObject = VM.NewJavaLangClass(BYTE_TYPE)
	SHORT_TYPE.classObject = VM.NewJavaLangClass(SHORT_TYPE)
	CHAR_TYPE.classObject = VM.NewJavaLangClass(CHAR_TYPE)
	INT_TYPE.classObject = VM.NewJavaLangClass(INT_TYPE)
	LONG_TYPE.classObject = VM.NewJavaLangClass(LONG_TYPE)
	FLOAT_TYPE.classObject = VM.NewJavaLangClass(FLOAT_TYPE)
	DOUBLE_TYPE.classObject = VM.NewJavaLangClass(DOUBLE_TYPE)
	BOOLEAN_TYPE.classObject = VM.NewJavaLangClass(BOOLEAN_TYPE)
}

// invoke <clinit> to execute initialization code
func (this *MethodArea) initialize(class *Class) {
	thread := VM.CurrentThread()

	LC := class.LC
	LOCK := LC.L
	LOCK.Lock()

	for {
		state := class.initialized.state
		if state == INITIALIZED {
			LOCK.Unlock()
			return
		} else if state == INITIALIZING {
			if class.initialized.thread == thread {
				LOCK.Unlock()
				return
			} else {
				LC.Wait()
			}
		} else if state == UNINITIALIZED {
			class.initialized.state = INITIALIZING
			class.initialized.thread = thread
			LOCK.Unlock()

			// initialize its supper class and interfaces first
			if class.superClass != nil {
				this.initialize(class.superClass)
			}
			for _, iface := range class.interfaces {
				this.initialize(iface)
			}

			if class.IsArray() {
				if elementClass, ok := class.elementType.(*Class); ok {
					this.initialize(elementClass)
				}
			}

			clinit := class.GetClassInitializer()
			if clinit != nil {
				VM.BootstrapClassLoader.Debug(repeat("\t", depth) + "⇉ %s \n", clinit.Qualifier())
				VM.InvokeMethod(clinit)
			}

			LOCK.Lock()
			class.initialized.state = INITIALIZED
			LC.Broadcast()
			LOCK.Unlock()

			return
		}
	}
}

func (this *MethodArea) verify(class *Class) {
	//TODO
}

// initialize static variables to default values: no need to execute code
func (this *MethodArea) prepare(class *Class)  {
	// Initialize static variables
	class.staticVars = make([]Value, class.maxStaticVars)
	for _, field := range class.fields {
		if field.IsStatic() {
			class.staticVars[field.slot] = field.defaultValue()
		}
	}
}


func (this *MethodArea) GetClass(name string) *Class {
	return VM.CreateClass(name, VM.CallerClass(), TRIGGER_BY_JAVA_REFLECTION)
}

func (this *MethodArea) GetTypeClass(descriptor string) JavaLangClass {
	var typeClass JavaLangClass
	switch string(descriptor[0]) {
	case JVM_SIGNATURE_CLASS: {
		fieldTypeClassName := descriptor[1:len(descriptor)-1]
		typeClass = VM.CreateClass(fieldTypeClassName, VM.CallerClass(), TRIGGER_BY_JAVA_REFLECTION).ClassObject()
	}
	case JVM_SIGNATURE_ARRAY: {
		fieldTypeClassName := descriptor
		typeClass = VM.CreateClass(fieldTypeClassName, VM.CallerClass(), TRIGGER_BY_JAVA_REFLECTION).ClassObject()
	}
	case JVM_SIGNATURE_BYTE: typeClass = BYTE_TYPE.ClassObject()
	case JVM_SIGNATURE_SHORT: typeClass = SHORT_TYPE.ClassObject()
	case JVM_SIGNATURE_CHAR: typeClass = CHAR_TYPE.ClassObject()
	case JVM_SIGNATURE_INT: typeClass = INT_TYPE.ClassObject()
	case JVM_SIGNATURE_LONG: typeClass = LONG_TYPE.ClassObject()
	case JVM_SIGNATURE_FLOAT: typeClass = FLOAT_TYPE.ClassObject()
	case JVM_SIGNATURE_DOUBLE: typeClass = DOUBLE_TYPE.ClassObject()
	case JVM_SIGNATURE_BOOLEAN: typeClass = BOOLEAN_TYPE.ClassObject()
	default:
		Fatal("type %s is not a unsupported type", descriptor)
	}

	return typeClass
}



