package jago

func register_jago_Jago() {
	register("jago/Jago.print(Ljava/lang/String;)V", Java_jago_Jago_print)
}

func Java_jago_Jago_print(str JavaLangString) {
	JavaOutPrintf("%s", str.toNativeString())
}

