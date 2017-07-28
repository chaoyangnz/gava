package jago

func register_jago_Jago() {
	register("jago/Jago.print(Ljava/lang/String;)V", JDK_jago_Jago_print)
	register("jago/Jago.print(D)V", JDK_jago_Jago_printDouble)
}

func JDK_jago_Jago_print(str JavaLangString) {
	VM_stdoutPrintf("%s", str.toNativeString())
}

func JDK_jago_Jago_printDouble(d Double)  {
	VM_stdoutPrintf("%v", d)
}

