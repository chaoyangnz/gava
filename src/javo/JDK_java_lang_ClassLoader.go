package javo

func register_java_lang_ClassLoader() {
	VM.RegisterNative("java/lang/ClassLoader.registerNatives()V", JDK_java_lang_ClassLoader_registerNatives)
	VM.RegisterNative("java/lang/ClassLoader.registerNatives()V", JDK_java_lang_ClassLoader_registerNatives)
	VM.RegisterNative("java/lang/ClassLoader.findBuiltinLib(Ljava/lang/String;)Ljava/lang/String;", JDK_java_lang_ClassLoader_findBuiltinLib)
	VM.RegisterNative("java/lang/ClassLoader$NativeLibrary.load(Ljava/lang/String;Z)V", JDK_java_lang_ClassLoader_NativeLibrary_load)
	VM.RegisterNative("java/lang/ClassLoader.findLoadedClass0(Ljava/lang/String;)Ljava/lang/Class;", JDK_java_lang_ClassLoader_findLoadedClass0)
	VM.RegisterNative("java/lang/ClassLoader.defineClass1(Ljava/lang/String;[BIILjava/security/ProtectionDomain;Ljava/lang/String;)Ljava/lang/Class;", JDK_java_lang_ClassLoader_defineClass1)
	VM.RegisterNative("java/lang/ClassLoader.findBootstrapClass(Ljava/lang/String;)Ljava/lang/Class;", JDK_java_lang_ClassLoader_findBootstrapClass)
}

func JDK_java_lang_ClassLoader_registerNatives() {
	// TODO
}

func JDK_java_lang_ClassLoader_findBuiltinLib(name JavaLangString) JavaLangString {
	return name
}

func JDK_java_lang_ClassLoader_NativeLibrary_load(this JavaLangClassLoader, name JavaLangString, flag Boolean) {
	// DO NOTHING
}

func JDK_java_lang_ClassLoader_findLoadedClass0(this JavaLangClassLoader, className JavaLangString) JavaLangClass {
	name := javaName2BinaryName(className)
	var C = NULL
	if class, ok := VM.getInitiatedClass(name, this); ok {
		C = class.ClassObject()
	}
	if C.IsNull() {
		VM.Info("  ***findLoadedClass0() %s fail [%s] \n", name, this.Class().name)
	} else {
		VM.Info("  ***findLoadedClass0() %s success [%s] \n", name, this.Class().name)
	}
	return C
}

func JDK_java_lang_ClassLoader_findBootstrapClass(this JavaLangClassLoader, className JavaLangString) JavaLangClass {
	name := javaName2BinaryName(className)
	var C = NULL
	if class, ok := VM.GetDefinedClass(name, NULL); ok {
		C = class.ClassObject()
		VM.Info("  ***findBootstrapClass() %s success [%s] *c=%p jc=%p \n", name, this.Class().name, class, class.classObject.oop)
	} else {
		c := VM.createClass(name, NULL, TRIGGER_BY_CHECK_OBJECT_TYPE)
		if c != nil {
			C = c.ClassObject()
		}
	}

	return C
}

func JDK_java_lang_ClassLoader_defineClass1(this JavaLangClassLoader, className JavaLangString, byteArrRef ArrayRef, offset Int, length Int, pd Reference, source JavaLangString) JavaLangClass {
	byteArr := byteArrRef.ArrayElements()[offset: offset+length]
	bytes := make([]byte, length)
	for i, b := range byteArr {
		bytes[i] = byte(b.(Byte))
	}

	C := VM.deriveClass(javaName2BinaryName(className), this, bytes, TRIGGER_BY_JAVA_CLASSLOADER)
	//VM.link(C)

	// associate JavaLangClass object
	//class.classObject = VM.NewJavaLangClass(class)
	//// specify its defining classloader
	C.ClassObject().SetInstanceVariableByName("classLoader", "Ljava/lang/ClassLoader;", this)
	VM.Info("  ==after java.lang.ClassLoader#defineClass1  %s *c=%p (derived) jc=%p \n", C.name, C, C.ClassObject().oop)

	//C.sourceFile = source.toNativeString() + C.Name() + ".java"

	return C.ClassObject()
}
