package jago

func register_java_lang_Thread() {
	register("java/lang/Thread.registerNatives()V", JDK_jang_lang_Thread_registerNatives)
	register("java/lang/Thread.currentThread()Ljava/lang/Thread;", JDK_java_lang_Thread_currentThread)
	register("java/lang/Thread.setPriority0(I)V", JDK_java_lang_Thread_setPriority0)
	register("java/lang/Thread.isAlive()Z", JDK_java_lang_Thread_isAlive)
	register("java/lang/Thread.start0()V", JDK_java_lang_Thread_start0)
	register("java/lang/Thread.sleep(J)V", JDK_java_lang_Thread_sleep)
	register("java/lang/Thread.interrupt0()V", JDK_java_lang_Thread_interrupt0)
}

// private static void registerNatives()
func JDK_jang_lang_Thread_registerNatives()  {}

func JDK_java_lang_Thread_currentThread() JavaLangThread  {
	return VM_currentThread().threadObject
}

func JDK_java_lang_Thread_setPriority0(this Reference, priority Int) {
	if priority < 1 {
		this.SetInstanceVariableByName("priority", "I", Int(5))
	}

}

func JDK_java_lang_Thread_isAlive(this Reference) Boolean {
	return FALSE
}

func JDK_java_lang_Thread_start0(this Reference) {
	if this.Class().name == "java/lang/ref/Reference$ReferenceHandler" {
		return // TODO hack: ignore these threads
	}
	name := this.GetInstanceVariableByName("name", "Ljava/lang/String;").(JavaLangString).toNativeString()
	runMethod := this.Class().GetMethod("run", "()V")

	thread := THREAD_MANAGER.NewThread(name, func() {
		VM_invokeMethod(runMethod, this)
	}, func(){})

	thread.threadObject = this
	thread.start()
}

func JDK_java_lang_Thread_sleep(millis Long)  {
	thread := VM_currentThread()
	interrupted := thread.sleep(int64(millis))
	if interrupted {
		VM_throw("java/lang/InterruptedException", "sleep interrupted")
	}
}

func JDK_java_lang_Thread_interrupt0(this JavaLangThread) {
	thread := this.GetExtra().(*Thread)
	thread.interrupt()
}

