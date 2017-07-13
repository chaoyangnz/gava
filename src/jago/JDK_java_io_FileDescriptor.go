package jago

func register_java_io_FileDescriptor()  {
	register("java/io/FileDescriptor.initIDs()V", Java_java_io_FileDescriptor_initIDs)
}

// private static void registers()
func Java_java_io_FileDescriptor_initIDs() {}
