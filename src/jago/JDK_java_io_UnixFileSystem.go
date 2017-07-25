package jago

func register_java_io_UnixFileSystem() {
	register("java/io/UnixFileSystem.initIDs()V", Java_java_io_UnixFileSystem_initIDs)
}

func Java_java_io_UnixFileSystem_initIDs() {
	// do nothing
}
