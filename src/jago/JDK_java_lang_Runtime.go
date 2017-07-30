package jago

import "runtime"

func register_java_lang_Runtime() {
	VM.RegisterNative("java/lang/Runtime.availableProcessors()I", JDK_jang_lang_Runtime_availableProcessors)
}

func JDK_jang_lang_Runtime_availableProcessors(this Reference) Int {
	return Int(runtime.NumCPU())
}

