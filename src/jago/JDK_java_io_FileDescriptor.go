package jago

func register_java_io_FileDescriptor() {
	VM.RegisterNative("java/io/FileDescriptor.initIDs()V", JDK_java_io_FileDescriptor_initIDs)
}

// private static void registers()
func JDK_java_io_FileDescriptor_initIDs() {}
