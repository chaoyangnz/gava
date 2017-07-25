package jago

func register_java_lang_ClassLoader() {
	register("java/lang/ClassLoader.registerNatives()V", Java_java_lang_ClassLoader_registerNatives)
}

func Java_java_lang_ClassLoader_registerNatives() {
	// TODO
}
