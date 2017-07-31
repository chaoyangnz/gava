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

	var file *os.File

	fileDescriptor := this.GetInstanceVariableByName("fd", "Ljava/io/FileDescriptor;").(Reference)
	path := this.GetInstanceVariableByName("path", "Ljava/lang/String;").(JavaLangString)

	if !path.IsNull() {
		f, err := os.Open(path.toNativeString())
		if err != nil {
			VM.Throw("java/lang/IOException", "Cannot open file: %s")
		}
		file = f
	} else if !fileDescriptor.IsNull() {
		fd := fileDescriptor.GetInstanceVariableByName("fd", "I").(Int)
		switch fd {
		case 0: file = os.Stdin
		case 1: file = os.Stdout
		case 2: file = os.Stderr
		default:
			file = os.NewFile(uintptr(fd), "")
		}
	}

	if file == nil {
		VM.Throw("java/lang/IOException", "File cannot open")
	}

	if append.IsTrue() {
		file.Chmod(os.ModeAppend)
	}

	bytes := make([]int8, byteArr.ArrayLength())
	for i := 0; i < int(byteArr.ArrayLength()); i++ {
		bytes[i] = int8(byteArr.GetArrayElement(Int(i)).(Byte))
	}

	bytes = bytes[off : off+len]
	ptr := unsafe.Pointer(&bytes)

	f := bufio.NewWriter(file)
	defer f.Flush()
	_, err := f.Write(*((*[]byte)(ptr)))
	if err == nil {
		return
	}
	VM.Throw("java/lang/IOException", "Cannot write to file: %s", file.Name())
}
