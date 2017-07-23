package jago

import "strings"

func register_java_lang_Throwable() {
	register("java/lang/Throwable.fillInStackTrace(I)Ljava/lang/Throwable;", Java_jang_lang_Throwable_fillInStackTrace)
}

func Java_jang_lang_Throwable_fillInStackTrace(this Reference, dummy Int) Reference {
	thread := THREAD_MANAGER.currentThread

	size := len(thread.vmStack) - 3 // skip 3 frames
	stacktrace := make([]string, size)


	for i, frame := range thread.vmStack[:size] {
		javaClassName := strings.Replace(frame.method.class.name, "/", ".", -1)
		stacktrace[size-1-i] = javaClassName + "." + frame.method.name + "()"
	}

	this.SetExtra(stacktrace)
	return this
}
