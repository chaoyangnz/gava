package jago

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

func JDK_java_lang_ClassLoader_findLoadedClass0(this Reference, className JavaLangString) JavaLangClass {
	name := javaName2BinaryName(className)
	if classAndLoader, ok := VM.GetCachedClass(name); ok {
		return classAndLoader.class.classObject
	}
	return NULL
}

func JDK_java_lang_ClassLoader_findBootstrapClass(this JavaLangClassLoader, className JavaLangString) JavaLangClass {
	if classAndLoader, ok := VM.GetCachedClass(className.toNativeString()); ok {
		if classAndLoader.classLoader.IsNull() {
			return classAndLoader.class.classObject
		}
	}
	return NULL
}

func JDK_java_lang_ClassLoader_defineClass1(this Reference, className JavaLangString, byteArrRef ArrayRef, offset Int, length Int, pd Reference, source JavaLangString) JavaLangClass {
	byteArr := byteArrRef.ArrayElements()[offset : offset+length]
	bytes := make([]byte, length)
	for i, b := range byteArr {
		bytes[i] = byte(b.(Byte))
	}

	class := VM.defineClass(bytes)
	class.classObject = VM.NewJavaLangClass(class)
	class.sourceFile = source.toNativeString()
	VM.link(class)
	VM.initialize(class)

	return class.classObject
}

