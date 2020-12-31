package javo

import (
	"os"
)

func register_java_io_FieInputStream() {
	VM.RegisterNative("java/io/FileInputStream.initIDs()V", JDK_java_io_FileInputStream_initIDs)
	VM.RegisterNative("java/io/FileInputStream.open0(Ljava/lang/String;)V", JDK_java_io_FileInputStream_open0)
	VM.RegisterNative("java/io/FileInputStream.readBytes([BII)I", JDK_java_io_FileInputStream_readBytes)
	VM.RegisterNative("java/io/FileInputStream.close0()V", JDK_java_io_FileInputStream_close0)
}

func JDK_java_io_FileInputStream_initIDs() {
	// TODO
}

func JDK_java_io_FileInputStream_open0(this Reference, name JavaLangString) {
	_, error := os.Open(name.toNativeString())
	if error != nil {
		VM.Throw("java/io/IOException", "Cannot open file: %s", name.toNativeString())
	}
}

func JDK_java_io_FileInputStream_readBytes(this Reference, byteArr ArrayRef, offset Int, length Int) Int {

	var file *os.File

	fileDescriptor := this.GetInstanceVariableByName("fd", "Ljava/io/FileDescriptor;").(Reference)
	path := this.GetInstanceVariableByName("path", "Ljava/lang/String;").(JavaLangString)

	if !path.IsNull() {
		f, err := os.Open(path.toNativeString())
		if err != nil {
			VM.Throw("java/io/IOException", "Cannot open file: %s", path.toNativeString())
		}
		file = f
	} else if !fileDescriptor.IsNull() {
		fd := fileDescriptor.GetInstanceVariableByName("fd", "I").(Int)
		switch fd {
		case 0:
			file = os.Stdin
		case 1:
			file = os.Stdout
		case 2:
			file = os.Stderr
		default:
			file = os.NewFile(uintptr(fd), "")
		}
	}

	if file == nil {
		VM.Throw("java/io/IOException", "File cannot open")
	}

	bytes := make([]byte, length)

	file.Seek(int64(offset), 0)
	nsize, err := file.Read(bytes)
	VM.ExecutionEngine.ioLogger.Info("🅹 ⤆ %s - buffer size: %d, offset: %d, len: %d, actual read: %d \n", file.Name(), byteArr.ArrayLength(), offset, length, nsize)
	if err == nil || nsize == int(length) {
		for i := 0; i < int(length); i++ {
			byteArr.SetArrayElement(offset+Int(i), Byte(bytes[i]))
		}
		return Int(nsize)
	}

	VM.Throw("java/io/IOException", err.Error())
	return -1
}

func JDK_java_io_FileInputStream_close0(this Reference) {
	var file *os.File

	fileDescriptor := this.GetInstanceVariableByName("fd", "Ljava/io/FileDescriptor;").(Reference)
	path := this.GetInstanceVariableByName("path", "Ljava/lang/String;").(JavaLangString)
	if !fileDescriptor.IsNull() {
		fd := fileDescriptor.GetInstanceVariableByName("fd", "I").(Int)
		switch fd {
		case 0:
			file = os.Stdin
		case 1:
			file = os.Stdout
		case 2:
			file = os.Stderr
		}
	} else {
		f, err := os.Open(path.toNativeString())
		if err != nil {
			VM.Throw("java/io/IOException", "Cannot open file: %s", path.toNativeString())
		}
		file = f
	}

	err := file.Close()
	if err != nil {
		VM.Throw("java/io/IOException", "Cannot close file: %s", path)
	}
}
