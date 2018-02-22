package jago

func register_java_lang_String() {
	VM.RegisterNative("java/lang/String.intern()Ljava/lang/String;", JDK_jang_lang_String_intern)
}

func JDK_jang_lang_String_intern(this JavaLangString) JavaLangString {
	return VM.InternString(this)
}
