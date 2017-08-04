package jago

import (
	"os"
	"bufio"
)

func register_java_io_FieOutputStream() {
	VM.RegisterNative("java/io/FileOutputStream.initIDs()V", JDK_java_io_FileOutputStream_initIDs)
	VM.RegisterNative("java/io/FileOutputStream.writeBytes([BIIZ)V", JDK_java_io_FileOutputStream_writeBytes)
}

func JDK_java_io_FileOutputStream_initIDs() {
	// TODO
}

func JDK_java_io_FileOutputStream_writeBytes(this Reference, byteArr ArrayRef, offset Int, length Int, append Boolean) {

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

	bytes := make([]byte, byteArr.ArrayLength())
	for i := 0; i < int(byteArr.ArrayLength()); i++ {
		bytes[i] = byte(int8(byteArr.GetArrayElement(Int(i)).(Byte)))
	}

	bytes = bytes[offset: offset+length]
	//ptr := unsafe.Pointer(&bytes)

	f := bufio.NewWriter(file)
	defer f.Flush()
	nsize, err := f.Write(bytes)
	VM.ExecutionEngine.ioLogger.Info("🅹 ⤇ %s - buffer size: %d, offset: %d, len: %d, actual write: %d \n", file.Name(), byteArr.ArrayLength(), offset, length, nsize)
	if err == nil {
		return
	}
	VM.Throw("java/lang/IOException", "Cannot write to file: %s", file.Name())
}
