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









