package jago

func register_jago_Jago() {
	register("jago/Jago.print(Ljava/lang/String;)V", Java_jago_Jago_print)
	register("jago/Jago.print(D)V", Java_jago_Jago_printDouble)
}

func Java_jago_Jago_print(str JavaLangString) {
	VM_stdoutPrintf("%s", str.toNativeString())
}

func Java_jago_Jago_printDouble(d Double)  {
	VM_stdoutPrintf("%v", d)
}

