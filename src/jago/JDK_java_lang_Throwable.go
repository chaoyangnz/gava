package jago

import (
	"strings"
	"fmt"
)

func register_java_lang_Throwable() {
	register("java/lang/Throwable.fillInStackTrace(I)Ljava/lang/Throwable;", Java_jang_lang_Throwable_fillInStackTrace)
}

func Java_jang_lang_Throwable_fillInStackTrace(this Reference, dummy Int) Reference {
	thread := VM_getCurrentThread()

	size := len(thread.vmStack) - this.Class().inheritanceDepth() // skip how many frames
	//backtrace := NewArray("[Ljava/lang/String;", Int(size))
	//
	//for i, frame := range thread.vmStack[:size] {
	//	javaClassName := strings.Replace(frame.method.class.name, "/", ".", -1)
	//	str := NewJavaLangString(javaClassName + "." + frame.method.name + __getSourceFileAndLineNumber(frame))
	//	backtrace.SetElement(Int(size-1-i), str)
	//}
	//
	//this.SetInstanceVariableByName("backtrace", "Ljava/lang/Object;", backtrace)

	backtrace := make([]string, size)

	for i, frame := range thread.vmStack[:size] {
		javaClassName := strings.Replace(frame.method.class.name, "/", ".", -1)
		backtrace[size-1-i] = fmt.Sprintf("%s.%s%s", javaClassName , frame.method.name, frame.getSourceFileAndLineNumber())
	}

	this.SetExtra(backtrace)

	return this
}


