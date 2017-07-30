package jago

func register_java_io_UnixFileSystem() {
	VM.RegisterNative("java/io/UnixFileSystem.initIDs()V", JDK_java_io_UnixFileSystem_initIDs)
}

func JDK_java_io_UnixFileSystem_initIDs() {
	// do nothing
}
