package jago

import (
	"strings"
	"fmt"
)

func register_java_lang_Throwable() {
	VM.RegisterNative("java/lang/Throwable.getStackTraceDepth()I", JDK_jang_lang_Throwable_getStackTraceDepth)
	VM.RegisterNative("java/lang/Throwable.fillInStackTrace(I)Ljava/lang/Throwable;", JDK_jang_lang_Throwable_fillInStackTrace)
	//VM.RegisterNative("java/lang/Throwable.getStackTraceElement(I)Ljava/lang/StackTraceElement;", JDK_jang_lang_Throwable_getStackTraceElement)
}

func JDK_jang_lang_Throwable_getStackTraceDepth(this Reference) Int {
	thread := VM.CurrentThread()
	return Int(len(thread.vmStack) - this.Class().inheritanceDepth()) // skip how many frames
}

func JDK_jang_lang_Throwable_fillInStackTrace(this Reference, dummy Int) Reference {
	thread := VM.CurrentThread()

	depth := len(thread.vmStack) - this.Class().inheritanceDepth() // skip how many frames
	//backtrace := NewArray("[Ljava/lang/String;", Int(depth))
	//
	//for i, frame := range thread.vmStack[:depth] {
	//	javaClassName := strings.Replace(frame.method.class.name, "/", ".", -1)
	//	str := NewJavaLangString(javaClassName + "." + frame.method.name + frame.getSourceFileAndLineNumber(this))
	//	backtrace.SetArrayElement(Int(depth-1-i), str)
	//}
	//
	//this.SetInstanceVariableByName("backtrace", "Ljava/lang/Object;", backtrace)

	backtrace := make([]string, depth)

	for i, frame := range thread.vmStack[:depth] {
		javaClassName := strings.Replace(frame.method.class.name, "/", ".", -1)
		backtrace[depth-1-i] = fmt.Sprintf("%s.%s%s", javaClassName , frame.method.name, frame.getSourceFileAndLineNumber())
	}

	this.oop.header.vmBackTrace = backtrace

	return this
}

//func JDK_jang_lang_Throwable_getStackTraceElement(i Int) Reference {
//
//}


