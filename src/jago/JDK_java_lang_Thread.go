package jago

func register_java_lang_Thread() {
	register("java/lang/Thread.registerNatives()V", JDK_jang_lang_Thread_registerNatives)
	register("java/lang/Thread.currentThread()Ljava/lang/Thread;", JDK_java_lang_Thread_currentThread)
	register("java/lang/Thread.setPriority0(I)V", JDK_java_lang_Thread_setPriority0)
	register("java/lang/Thread.isAlive()Z", JDK_java_lang_Thread_isAlive)
	register("java/lang/Thread.start0()V", JDK_java_lang_Thread_start0)
}

// private static void registerNatives()
func JDK_jang_lang_Thread_registerNatives()  {}

func JDK_java_lang_Thread_currentThread() JavaLangThread  {
	return VM_getCurrentThread().threadObject
}

func JDK_java_lang_Thread_setPriority0(this Reference, priority Int) {
	//todo
}

func JDK_java_lang_Thread_isAlive(this Reference) Boolean {
	return FALSE
}

func JDK_java_lang_Thread_start0(this Reference) {
	//todo
}

