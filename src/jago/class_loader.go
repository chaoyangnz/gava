package jago

type BootstrapClassLoader struct {
	classPath     *ClassPath
	*Logger
}

type ClassTriggerReason struct {
	code string
	desc string
}

var (
	TRIGGER_BY_JAVA_REFLECTION    = &ClassTriggerReason{"JR", "java reflection from name"}
	TRIGGER_BY_JAVA_CLASSLOADER   = &ClassTriggerReason{"JR", "user defined classloader except bootstrap classloader"}
	TRIGGER_BY_CHECK_OBJECT_TYPE  = &ClassTriggerReason{"CT", "check java object type"}
	TRIGGER_BY_AS_SUPERCLASS      = &ClassTriggerReason{"SC", "as superclass"}
	TRIGGER_BY_AS_SUPERINTERFACE  = &ClassTriggerReason{"SI", "as superinterface"}
	TRIGGER_BY_AS_ARRAY_COMPONENT = &ClassTriggerReason{"AC", "as array component"}
	TRIGGER_BY_RESOLVE_CLASS_REF  = &ClassTriggerReason{"RR", "revolve symbol_ref in constant pool"}
	TRIGGER_BY_NEW_INSTANCE       = &ClassTriggerReason{"NI", "new instance"}
	TRIGGER_BY_ACCESS_MEMBER      = &ClassTriggerReason{"AM", "access field or method"}
)

func (this *BootstrapClassLoader) LoadClass(N string, triggerReason *ClassTriggerReason) *Class {
	if C, ok := VM.GetInitiatedClass(N, NULL); ok {
		return C
	}

	// search for a purported representation of C, typically, a class or interface will be represented using a file in a hierachical file system
	bytecode := this.findClass(N)

	C := VM.deriveClass(N, NULL, bytecode, triggerReason)

	//this.depth--
	return C
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









