package jago

import (
	"os"
	"unsafe"
	"bufio"
)

func register_java_io_FieOutputStream() {
	VM.RegisterNative("java/io/FileOutputStream.initIDs()V", JDK_java_io_FileOutputStream_initIDs)
	VM.RegisterNative("java/io/FileOutputStream.writeBytes([BIIZ)V", JDK_java_io_FileOutputStream_writeBytes)
}

func JDK_java_io_FileOutputStream_initIDs() {
	// TODO
}

func JDK_java_io_FileOutputStream_writeBytes(this Reference, byteArr ArrayRef, off Int, len Int, append Boolean) {
	bytes := make([]int8, byteArr.ArrayLength())
	for i := 0; i < int(byteArr.ArrayLength()); i++ {
		bytes[i] = int8(byteArr.GetArrayElement(Int(i)).(Byte))
	}

	bytes = bytes[off : off+len]
	ptr := unsafe.Pointer(&bytes)

	var file *os.File

	fileDescriptor := this.GetInstanceVariableByName("fd", "Ljava/io/FileDescriptor;").(Reference)
	if !fileDescriptor.IsNull() {
		fd := fileDescriptor.GetInstanceVariableByName("fd", "I").(Int)
		switch fd {
		case 0: file = os.Stdin
		case 1: file = os.Stdout
		case 2: file = os.Stderr
		}
	} else {
		path := this.GetInstanceVariableByName("path", "Ljava/lang/String;").(JavaLangString)
		f, err := os.Open(path.toNativeString())
		if err != nil {
			Fatal("Write file %s failed", path)
		}
		file = f
	}

	if append.IsTrue() {
		file.Chmod(os.ModeAppend)
	}

	f := bufio.NewWriter(file)
	defer f.Flush()
	f.Write(*((*[]byte)(ptr)))
}
