package jago

func register_java_lang_String() {
	register("java/lang/String.intern()Ljava/lang/String;", JDK_jang_lang_String_intern)
}

func JDK_jang_lang_String_intern(this JavaLangString) JavaLangString {
	return VM_intern_String(this)
}

