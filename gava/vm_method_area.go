package gava

import (
	"math"
	"strings"
	"sync"
	"unsafe"
)

type ClassTriggerReason struct {
	code string
	desc string
}

var (
	TRIGGER_BY_JAVA_REFLECTION    = &ClassTriggerReason{"JR", "java reflection from name"}
	TRIGGER_BY_JAVA_CLASSLOADER   = &ClassTriggerReason{"JC", "user defined classloader except bootstrap classloader"}
	TRIGGER_BY_CHECK_OBJECT_TYPE  = &ClassTriggerReason{"CT", "check java object type"}
	TRIGGER_BY_AS_SUPERCLASS      = &ClassTriggerReason{"SC", "as superclass"}
	TRIGGER_BY_AS_SUPERINTERFACE  = &ClassTriggerReason{"SI", "as superinterface"}
	TRIGGER_BY_AS_ARRAY_COMPONENT = &ClassTriggerReason{"AC", "as array component"}
	TRIGGER_BY_RESOLVE_CLASS_REF  = &ClassTriggerReason{"RR", "revolve symbol_ref in constant pool"}
	TRIGGER_BY_NEW_INSTANCE       = &ClassTriggerReason{"NI", "new instance"}
	TRIGGER_BY_ACCESS_MEMBER      = &ClassTriggerReason{"AM", "access field or method"}
)

/*
|--------------------------------------------------------------------------
| String pool in method area
|--------------------------------------------------------------------------
|
| String pool is actually a table and its key is internal native string
| associating with the Java heap String object (java.lang.String)
| Here, only intern() method is provided. All the String objects generated
| when implementation the VM are interned.
|
| The String objects in Java code can be interned by calling String.intern()
| method manually.
*/

type StringPool map[string]JavaLangString

func (this StringPool) GetStringInPool(str string) (JavaLangString, bool) {
	strObj, found := this[str]
	if !found {
		strObj = NULL
	}
	return strObj, found
}

func (this StringPool) InternString(stringobj JavaLangString) JavaLangString {
	str := stringobj.toNativeString()

	if strObj, found := this[str]; found {
		return strObj
	} else {
		this[str] = stringobj
		return stringobj
	}
}

/*
|--------------------------------------------------------------------------
| Method Area structure
|--------------------------------------------------------------------------
|
| Method Area includes all the classes (internal representation) defined and
| initiated by some Class Loaders.
|
| By default, it contains a Bootstrap Class Loader which is implemented within
| VM rather than in Java code.
|
| All its methods are related to class loading, linking and initialization.
| Please refer to JVM Specification: 5. Loading, Linking, and Initializing
*/

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

func (this MethodArea) getInitiatedClass(className string, classLoader JavaLangClassLoader) (*Class, bool) {
	if class, ok := this.InitiatedClasses[NL{className, classLoader.oop}]; ok {
		return class, true
	}
	return nil, false
}

func (this MethodArea) setInitiatedClass(className string, classLoader JavaLangClassLoader, class *Class) {
	VM.Info("-> Set initiated class (%s, %p): %p \n", className, classLoader.oop, class)
	this.InitiatedClasses[NL{className, classLoader.oop}] = class
}

func (this MethodArea) GetDefinedClass(className string, classLoader JavaLangClassLoader) (*Class, bool) {
	if class, ok := this.DefinedClasses[NL{className, classLoader.oop}]; ok {
		return class, true
	}
	return nil, false
}

func (this MethodArea) getDefiningClassLoader(class *Class) JavaLangClassLoader {
	for cn, cl := range this.DefinedClasses {
		if cl == class && cn.N == class.name {
			return Reference{cn.L}
		}
	}
	return NULL
}

func (this MethodArea) getInitiatingClassLoader(class *Class) JavaLangClassLoader {
	for cn, cl := range this.InitiatedClasses {
		if cl == class && cn.N == class.name {
			return Reference{cn.L}
		}
	}
	return NULL
}

func (this MethodArea) setDefinedClass(className string, classLoader JavaLangClassLoader, class *Class) {
	VM.Info("-> Set defined class (%s, %p): %p \n", className, classLoader.oop, class)
	this.DefinedClasses[NL{className, classLoader.oop}] = class
}

var depth = 0

func (this *MethodArea) createClass(N string, Ld JavaLangClassLoader, triggerReason *ClassTriggerReason) *Class {
	var C *Class
	if N[:1] != JVM_SIGNATURE_ARRAY {
		if Ld.IsNull() {
			// bootstrap classloader
			C = VM.LoadClass(N, triggerReason)
		} else {
			// user-defined classloader
			C = VM.loadClassUd(N, Ld, triggerReason)
		}
	} else {
		C = this.createArrayClass(N, Ld, triggerReason)
	}

	VM.link(C) // eagerly link

	return C
}

/*
Load class using user-defined class loader
*/
func (this *MethodArea) loadClassUd(N string, L JavaLangClassLoader, triggerReason *ClassTriggerReason) *Class {
	if C, ok := VM.getInitiatedClass(N, L); ok {
		return C
	}
	loadClassMethod := L.Class().FindMethod("loadClass", "(Ljava/lang/String;)Ljava/lang/Class;")
	VM.Info("==before java.lang.ClassLoader#loadClass %s in LoadClassUd \n", N)
	javaname := binaryNameToJavaName(N)
	classObject := VM.InvokeMethod(loadClassMethod, L, javaname).(JavaLangClass)

	VM.Info("==after java.lang.ClassLoader#loadClass %s %s in LoadClassUd jc=%p \n", N, classObject.Class().name, classObject.oop)
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
			className := arrayClassName[2 : len(arrayClassName)-1]
			return className, className, 1
		}
	case JVM_SIGNATURE_ARRAY:
		{
			componentArrayClassName := arrayClassName[1:]
			_, elementTypeName, dimension := componentAndElementTypeName(componentArrayClassName)
			return componentArrayClassName, elementTypeName, dimension + 1
		}
	}
	VM.Throw("java/lang/IllegalArgumentException", "%s is not a array class name")
	return "", "", -1
}

func (this *MethodArea) createArrayClass(N string, L JavaLangClassLoader, triggerReason *ClassTriggerReason) *Class {
	if C, ok := VM.getInitiatedClass(N, L); ok {
		return C //???
		// TODO If L has already been recorded as an intiating loader of an array class with the same component type as N,
		// the class is C. No array class creation is neccessry
	}

	if C, ok := VM.GetDefinedClass(N, NULL); ok { // this lookup is not mentioned in jvms
		return C
	}

	// ------------------
	VM.BootstrapClassLoader.Debug(repeat("\t", depth)+"↳ %s ", N)
	VM.BootstrapClassLoader.Debug("(reason: %s", triggerReason.code)
	VM.BootstrapClassLoader.Trace(" - %s", triggerReason.desc)
	VM.BootstrapClassLoader.Debug(") ")
	var cll string
	if L.IsNull() {
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
		LC:             sync.NewCond(&sync.Mutex{})}

	C.superClass = VM.createClass("java/lang/Object", NULL, TRIGGER_BY_AS_SUPERCLASS)
	C.interfaces = []*Class{
		VM.createClass("java/io/Serializable", NULL, TRIGGER_BY_AS_SUPERINTERFACE),
		VM.createClass("java/lang/Cloneable", NULL, TRIGGER_BY_AS_SUPERINTERFACE)}

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
			componentTypeName := N[2 : len(N)-1]
			C.componentType = VM.createClass(componentTypeName, L, TRIGGER_BY_AS_ARRAY_COMPONENT)
			C.elementType = C.componentType
			C.dimensions = 1
		}
	case JVM_SIGNATURE_ARRAY:
		{
			componentTypeName := N[2 : len(N)-1]
			componentArrayClass := VM.createClass(componentTypeName, L, TRIGGER_BY_AS_ARRAY_COMPONENT)
			C.componentType = componentArrayClass
			C.elementType = componentArrayClass.elementType
			C.dimensions = componentArrayClass.dimensions + 1
		}
	}

	switch C.componentType.(type) {
	case *Class:
		componentClass := C.componentType.(*Class)
		C.accessFlags = componentClass.accessFlags
		VM.setDefinedClass(N, VM.getDefiningClassLoader(componentClass), C)
	default:
		C.accessFlags = JVM_ACC_PUBLIC
		VM.setDefinedClass(N, NULL, C)
	}

	VM.setInitiatedClass(N, L, C)
	C.defined = true
	C.initialized = UNINITIALIZED

	depth--

	return C
}

// jvms 5.4.3.1
func (this *MethodArea) resolveClass(N string, D *Class, triggerReason *ClassTriggerReason) *Class {
	var Ld JavaLangClassLoader
	if D == nil {
		Ld = NULL
	} else {
		Ld = VM.getDefiningClassLoader(D)
	}
	C := this.createClass(N, Ld, triggerReason)
	if C.IsArray() {
		_, elementTypeName, _ := componentAndElementTypeName(C.Name())
		if elementTypeName[:1] == JVM_SIGNATURE_CLASS { // cannot be array for an element type
			C.elementType = this.resolveClass(elementTypeName, D, TRIGGER_BY_AS_ARRAY_COMPONENT)
		}
	}

	return C
}

// jvms 5.3.5
func (this *MethodArea) deriveClass(N string, L JavaLangClassLoader, bytecode []byte, triggerReason *ClassTriggerReason) *Class {

	if _, ok := VM.getInitiatedClass(N, L); ok {
		VM.Throw("java/lang/LinkageError", "Classloader %s is an initiating loader of class %s", L.Class().Name(), N)
	}

	// ------------------
	VM.BootstrapClassLoader.Debug(repeat("\t", depth)+"↳ %s ", N)
	VM.BootstrapClassLoader.Debug("(reason: %s", triggerReason.code)
	VM.BootstrapClassLoader.Trace(" - %s", triggerReason.desc)
	VM.BootstrapClassLoader.Debug(") ")
	var cll string
	if L.IsNull() {
		cll = "<bootstrap>"
	} else {
		cll = L.Class().name
	}
	VM.BootstrapClassLoader.Debug("[%s] (derive)\n", cll)
	depth++
	// ------------------

	classfile := &ClassFile{}
	classfile.read(bytecode)

	// purported representation is not of a supported major or minor version -> java/lang.UnsupportedClassVersionError
	if classfile.majorVersion > JVM_MAJOR_VERSION_JAVA_8 {
		Fatal("Incompatible class version, Java 8 or lower is supported")
	}
	// TODO class file does not actually represent a class name N -> java/langNoClassDefFoundError

	C := &Class{LC: sync.NewCond(&sync.Mutex{})}

	C.accessFlags = uint16(classfile.accessFlags)
	C.name = classfile.cpUtf8(classfile.constantPool[classfile.thisClass].(*ConstantClassInfo).nameIndex)

	// constant pool
	constantPool := make([]Constant, classfile.constantPoolCount+1)
	C.constantPool = constantPool

	for i := range classfile.constantPool {
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
			constant = &UTF8Constant{C, constantUtf8Info.value()}
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
			l := int64(uint64(constantLongInfo.highBytes)<<32 + uint64(constantLongInfo.lowBytes))
			constant = &LongConstant{C, Long(l)}
		case *ConstantDoubleInfo:
			constantDoubleInfo := constInfo.(*ConstantDoubleInfo)
			bits := uint64(constantDoubleInfo.highBytes)<<32 + uint64(constantDoubleInfo.lowBytes)
			d := math.Float64frombits(bits)
			constant = &DoubleConstant{C, Double(d)}
		case *ConstantNameAndTypeInfo:
			constantNameAndTypeInfo := constInfo.(*ConstantNameAndTypeInfo)
			constant = &NameAndTypeConstant{C, classfile.cpUtf8(constantNameAndTypeInfo.nameIndex), classfile.cpUtf8(constantNameAndTypeInfo.descriptorIndex)}
		case *ConstantMethodTypeInfo:
			constantMethodTypeInfo := constInfo.(*ConstantMethodTypeInfo)
			constant = &MethodTypeConstant{C, classfile.cpUtf8(constantMethodTypeInfo.descriptorIndex)}
		case *ConstantInvokeDynamicInfo:
			constantInvokeDynamicInfo := constInfo.(*ConstantInvokeDynamicInfo)
			nameAndTypeInfo := classfile.constantPool[constantInvokeDynamicInfo.nameAndTypeIndex].(*ConstantNameAndTypeInfo)
			name := classfile.cpUtf8(nameAndTypeInfo.nameIndex)
			descriptor := classfile.cpUtf8(nameAndTypeInfo.descriptorIndex)
			constant = &InvokeDynamicConstant{C,
				"", //TODO classfile.cpUtf8(constantInvokeDynamicInfo.bootstrapMethodAttrIndex),
				name,
				descriptor}
		}

		C.constantPool[i] = constant
	}

	for k := u2(0); k < classfile.attributeCount; k++ {
		attributeInfo := classfile.attributes[k]
		switch attributeInfo.(type) {
		case *SourceFileAttribute:
			sourceFileAttribue := attributeInfo.(*SourceFileAttribute)
			C.sourceFile = classfile.cpUtf8(sourceFileAttribue.sourceFileIndex)
		}
	}

	C.fields = make([]*Field, len(classfile.fields))
	//class.staticVars = ??

	for i, fieldInfo := range classfile.fields {
		field := &Field{}
		field.accessFlags = FieldAccessFlag(fieldInfo.accessFlags)
		field.name = classfile.cpUtf8(fieldInfo.nameIndex)
		field.descriptor = classfile.cpUtf8(fieldInfo.descriptorIndex)
		field.class = C
		//field.index = ??
		C.fields[i] = field
	}

	C.methods = make([]*Method, len(classfile.methods))

	for i, methodInfo := range classfile.methods {
		method := &Method{}
		method.accessFlags = MethodAccessFlag(methodInfo.accessFlags)
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
			Ref:
				for j := i + 1; j < len(parametersStr); j++ {
					switch rune(parametersStr[j]) {
					case ';':
						parameterDescriptors = append(parameterDescriptors, string(parametersStr[i:j+1]))
						i = j + 1
						break Ref
					}
				}
			case '[':
			Arr:
				for j := i + 1; j < len(parametersStr); j++ {
					switch rune(parametersStr[j]) {
					case '[':
						continue
					case 'B', 'C', 'D', 'F', 'I', 'J', 'S', 'Z':
						parameterDescriptors = append(parameterDescriptors, string(parametersStr[i:j+1]))
						i = j + 1
						break Arr
					case 'L':
						for k := j + 1; j < len(parametersStr); k++ {
							switch rune(parametersStr[k]) {
							case ';':
								parameterDescriptors = append(parameterDescriptors, string(parametersStr[i:k+1]))
								i = k + 1
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
	instanceVarsCount := 0 // include all the ancestry
	staticVarsCount := 0

	instanceVarFields := make([]*Field, 0)
	staticVarFields := make([]*Field, 0)

	if C.superClass != nil {
		instanceVarsCount = C.superClass.instanceVarsCount

		instanceVarFields = append(instanceVarFields, C.superClass.instanceVarFields...)
	}
	for _, field := range C.fields {
		if field.IsStatic() {
			field.slot = staticVarsCount
			staticVarFields = append(staticVarFields, field)
			staticVarsCount++
		} else {
			field.slot = instanceVarsCount
			instanceVarFields = append(instanceVarFields, field)
			instanceVarsCount++
		}
	}
	C.instanceVarsCount = instanceVarsCount
	C.staticVarsCount = staticVarsCount

	C.instanceVarFields = instanceVarFields
	C.staticVarFields = staticVarFields

	// marks C as having L as its defining class loader and records that L is an initiating loader of C
	VM.setDefinedClass(N, L, C)
	VM.setInitiatedClass(N, L, C)

	C.defined = true
	C.initialized = UNINITIALIZED

	depth--

	return C
}

func (this *MethodArea) link(class *Class) {
	if class.linked {
		return
	}
	VM.verify(class)
	VM.prepare(class)

	class.linked = true
	// we resolve each symbolic class in a class or interface individually when it is used ("lazy" or "late" resolution)
	// So SymbolRef all implements a method PrimitiveType resolve()
}

// invoke <clinit> to execute initialization code
func (this *MethodArea) initialize(class *Class) {
	thread := VM.CurrentThread()

	LC := class.LC
	LOCK := LC.L
	LOCK.Lock()

	for {
		state := class.initialized
		if state == INITIALIZED {
			LOCK.Unlock()
			return
		} else if state == INITIALIZING {
			if class.T == thread {
				LOCK.Unlock()
				return
			} else {
				LC.Wait()
			}
		} else if state == UNINITIALIZED {
			class.initialized = INITIALIZING
			class.T = thread
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
				VM.BootstrapClassLoader.Debug(repeat("\t", depth)+"⇉ %s \n", clinit.Qualifier())
				VM.InvokeMethod(clinit)
			}

			LOCK.Lock()
			class.initialized = INITIALIZED
			class.T = nil
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
func (this *MethodArea) prepare(class *Class) {
	// Initialize static variables
	class.staticVars = make([]Value, class.staticVarsCount)
	for _, field := range class.fields {
		if field.IsStatic() {
			class.staticVars[field.slot] = field.DefaultValue()
		}
	}
}

/*
Always use current class's class loader to resolve class
*/
func (this *MethodArea) ResolveClass(N string, reason *ClassTriggerReason) *Class {
	return VM.resolveClass(N, VM.CurrentClass(), reason)
}

func (this *MethodArea) GetTypeClass(descriptor string) JavaLangClass {
	var typeClass JavaLangClass
	switch string(descriptor[0]) {
	case JVM_SIGNATURE_CLASS:
		{
			fieldTypeClassName := descriptor[1 : len(descriptor)-1]
			typeClass = VM.ResolveClass(fieldTypeClassName, TRIGGER_BY_JAVA_REFLECTION).ClassObject()
		}
	case JVM_SIGNATURE_ARRAY:
		{
			fieldTypeClassName := descriptor
			typeClass = VM.ResolveClass(fieldTypeClassName, TRIGGER_BY_JAVA_REFLECTION).ClassObject()
		}
	case JVM_SIGNATURE_BYTE:
		typeClass = BYTE_TYPE.ClassObject()
	case JVM_SIGNATURE_SHORT:
		typeClass = SHORT_TYPE.ClassObject()
	case JVM_SIGNATURE_CHAR:
		typeClass = CHAR_TYPE.ClassObject()
	case JVM_SIGNATURE_INT:
		typeClass = INT_TYPE.ClassObject()
	case JVM_SIGNATURE_LONG:
		typeClass = LONG_TYPE.ClassObject()
	case JVM_SIGNATURE_FLOAT:
		typeClass = FLOAT_TYPE.ClassObject()
	case JVM_SIGNATURE_DOUBLE:
		typeClass = DOUBLE_TYPE.ClassObject()
	case JVM_SIGNATURE_BOOLEAN:
		typeClass = BOOLEAN_TYPE.ClassObject()
	default:
		Fatal("type %s is not a unsupported type", descriptor)
	}

	return typeClass
}

/*
|--------------------------------------------------------------------------
| Bootstrap class loader implemented in VM internally
|--------------------------------------------------------------------------
|
| The system class path is initialized when VM startup.
*/

type BootstrapClassLoader struct {
	classPath *ClassPath
	*Logger
}

func (this *BootstrapClassLoader) LoadClass(N string, triggerReason *ClassTriggerReason) *Class {
	if C, ok := VM.getInitiatedClass(N, NULL); ok {
		return C
	}

	// search for a purported representation of C, typically, a class or interface will be represented using a file in a hierachical file system
	bytecode := this.findClass(N)

	C := VM.deriveClass(N, NULL, bytecode, triggerReason)

	//this.depth--
	return C
}

/*
@throw java.lang.ClassNotFoundException
*/
func (this *BootstrapClassLoader) findClass(className string) []byte {
	bytecode, err := this.classPath.ReadClass(className)
	if err != nil {
		VM.Throw("java/lang/ClassNotFoundException", className)
	}

	return bytecode
}
