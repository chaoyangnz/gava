package jago

func register_jago_Jago() {
	VM.RegisterNative("jago/Jago.print(Ljava/lang/String;)V", JDK_jago_Jago_print)
	VM.RegisterNative("jago/Jago.print(D)V", JDK_jago_Jago_printDouble)
}

func JDK_jago_Jago_print(str JavaLangString) {
	VM.StdoutPrintf("%s", str.toNativeString())
}

func JDK_jago_Jago_printDouble(d Double)  {
	VM.StdoutPrintf("%v", d)
}

