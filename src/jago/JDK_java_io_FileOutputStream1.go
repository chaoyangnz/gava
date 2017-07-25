package jago

import (
	"os"
	"unsafe"
)

func register_java_io_FieOutputStream() {
	register("java/io/FileOutputStream.initIDs()V", Java_java_io_FileOutputStream_initIDs)
	register("java/io/FileOutputStream.writeBytes([BIIZ)V", Java_java_io_FileOutputStream_writeBytes)
}

func Java_java_io_FileOutputStream_initIDs() {
	// TODO
}

func Java_java_io_FileOutputStream_writeBytes(this Reference, byteArr ArrayRef, off Int, len Int, append Boolean) {
	bytes := make([]int8, byteArr.Length())
	for i := 0; i < int(byteArr.Length()); i++ {
		bytes[i] = int8(byteArr.GetElement(Int(i)).(Byte))
	}

	ptr := unsafe.Pointer(&bytes)
	os.Stdout.Write(*((*[]byte)(ptr)))
}
