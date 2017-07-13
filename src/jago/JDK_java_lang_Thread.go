package jago

func register_java_lang_Thread() {
	register("java/lang/Thread.registerNatives()V", Java_jang_lang_Thread_registerNatives)
	register("java/lang/Thread.currentThread()Ljava/lang/Thread;", Java_java_lang_Thread_currentThread)
	register("java/lang/Thread.setPriority0(I)V", Java_java_lang_Thread_setPriority0)
	register("java/lang/Thread.isAlive()Z", Java_java_lang_Thread_isAlive)
	register("java/lang/Thread.start0()V", Java_java_lang_Thread_start0)
}

// private static void registerNatives()
func Java_jang_lang_Thread_registerNatives()  {}

func Java_java_lang_Thread_currentThread() JavaLangThread  {
	return VM_getCurrentThread().threadObject
}

func Java_java_lang_Thread_setPriority0(this Reference, priority Int) {
	//todo
}

func Java_java_lang_Thread_isAlive(this Reference) Boolean {
	return FALSE
}

func Java_java_lang_Thread_start0(this Reference) {
	//todo
}

