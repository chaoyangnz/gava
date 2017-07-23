package jago

func register_java_lang_String() {
	register("java/lang/String.intern()Ljava/lang/String;", Java_jang_lang_String_intern)
}

func Java_jang_lang_String_intern(this JavaLangString) JavaLangString {
	return VM_intern_String(this)
}

