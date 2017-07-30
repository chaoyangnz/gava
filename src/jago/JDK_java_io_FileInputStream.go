package jago

func register_java_io_FieInputStream() {
	VM.RegisterNative("java/io/FileInputStream.initIDs()V", JDK_java_io_FileInputStream_initIDs)
}

func JDK_java_io_FileInputStream_initIDs() {
	// TODO
}
