package jago

func register_java_lang_Throwable() {
	register("java/lang/Throwable.fillInStackTrace(I)Ljava/lang/Throwable;", Java_jang_lang_Throwable_fillInStackTrace)
}

func Java_jang_lang_Throwable_fillInStackTrace(this Reference, dummy Int) Reference {
	// TODO
	return this
}
