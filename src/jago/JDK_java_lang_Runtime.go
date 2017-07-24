package jago

import "runtime"

func register_java_lang_Runtime() {
	register("java/lang/Runtime.availableProcessors()I", Java_jang_lang_Runtime_availableProcessors)
}

func Java_jang_lang_Runtime_availableProcessors(this Reference) Int {
	return Int(runtime.NumCPU())
}

