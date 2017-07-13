package jago

func register_java_lang_Object() {
	register("java/lang/Object.registerNatives()V", Java_jang_lang_Object_registerNatives)
	register("java/lang/Object.hashCode()I", Java_java_lang_Object_hashCode)
}

// private static void registerNatives()
func Java_jang_lang_Object_registerNatives()  {}

func Java_java_lang_Object_hashCode(this Reference) Int  {
	return VM_iHashCode(this)
}
