package jago

func register_java_lang_ClassLoader() {
	register("java/lang/ClassLoader.registerNatives()V", Java_java_lang_ClassLoader_registerNatives)
	register("java/lang/ClassLoader.findBuiltinLib(Ljava/lang/String;)Ljava/lang/String;", Java_java_lang_ClassLoader_findBuiltinLib)
	register("java/lang/ClassLoader$NativeLibrary.load(Ljava/lang/String;Z)V", Java_java_lang_ClassLoader_NativeLibrary_load)
}

func Java_java_lang_ClassLoader_registerNatives() {
	// TODO
}

func Java_java_lang_ClassLoader_findBuiltinLib(name JavaLangString) JavaLangString {
	return name
}

func Java_java_lang_ClassLoader_NativeLibrary_load(this JavaLangClassLoader, name JavaLangString, flag Boolean) {
	// DO NOTHING
}
