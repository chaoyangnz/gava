package jago

func register_java_lang_ClassLoader() {
	VM.RegisterNative("java/lang/ClassLoader.registerNatives()V", JDK_java_lang_ClassLoader_registerNatives)
	VM.RegisterNative("java/lang/ClassLoader.findBuiltinLib(Ljava/lang/String;)Ljava/lang/String;", JDK_java_lang_ClassLoader_findBuiltinLib)
	VM.RegisterNative("java/lang/ClassLoader$NativeLibrary.load(Ljava/lang/String;Z)V", JDK_java_lang_ClassLoader_NativeLibrary_load)
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
