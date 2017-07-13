package jago

func register_java_lang_Class() {
	register("java/lang/Class.registerNatives()V", Java_jang_lang_Class_registerNatives)
	register("java/lang/Class.getPrimitiveClass(Ljava/lang/String;)Ljava/lang/Class;", Java_java_lang_Class_getPrimitiveClass)
	register("java/lang/Class.desiredAssertionStatus0(Ljava/lang/Class;)Z", Java_java_lang_Class_desiredAssertionStatus0)

}

// private static void registerNatives()
func Java_jang_lang_Class_registerNatives()  {}

// static Class getPrimitiveClass(String name)
func Java_java_lang_Class_getPrimitiveClass(name JavaLangString) JavaLangClass {
	// TODO??
	return NewJavaLangClass()
}

// private static boolean desiredAssertionStatus0(Class javaClass)
func Java_java_lang_Class_desiredAssertionStatus0(clazz JavaLangClass) Boolean {
	// TODO
	return TRUE
}
