package javo

import (
	"strings"
)

func register_java_lang_Throwable() {
	VM.RegisterNative("java/lang/Throwable.getStackTraceDepth()I", JDK_jang_lang_Throwable_getStackTraceDepth)
	VM.RegisterNative("java/lang/Throwable.fillInStackTrace(I)Ljava/lang/Throwable;", JDK_jang_lang_Throwable_fillInStackTrace)
	VM.RegisterNative("java/lang/Throwable.getStackTraceElement(I)Ljava/lang/StackTraceElement;", JDK_jang_lang_Throwable_getStackTraceElement)
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

	backtrace := make([]StackTraceElement, depth)

	for i, frame := range thread.vmStack[:depth] {
		javaClassName := strings.Replace(frame.method.class.name, "/", ".", -1)
		methodName := frame.method.name
		fileName := frame.getSourceFile()
		lineNumber := frame.getLineNumber()
		backtrace[depth-1-i] = StackTraceElement{javaClassName, methodName, fileName, lineNumber}
	}

	this.attachStacktrace(backtrace)

	return this
}

func JDK_jang_lang_Throwable_getStackTraceElement(this JavaLangThrowable, i Int) ObjectRef {
	stacktraceelement := this.retrieveStacktrace()[i]

	ste := VM.NewObjectOfName("java/lang/StackTraceElement")
	ste.SetInstanceVariableByName("declaringClass", "Ljava/lang/String;", VM.NewJavaLangString(stacktraceelement.declaringClass))
	ste.SetInstanceVariableByName("methodName", "Ljava/lang/String;", VM.NewJavaLangString(stacktraceelement.methodName))
	ste.SetInstanceVariableByName("fileName", "Ljava/lang/String;", VM.NewJavaLangString(stacktraceelement.fileName))
	ste.SetInstanceVariableByName("lineNumber", "I", Int(stacktraceelement.lineNumber))

	return ste
}
