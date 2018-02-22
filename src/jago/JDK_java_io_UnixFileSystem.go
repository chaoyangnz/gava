package jago

import (
	"path/filepath"
	"os"
	"runtime"
	"log"
)

func register_java_io_UnixFileSystem() {
	VM.RegisterNative("java/io/UnixFileSystem.initIDs()V", JDK_java_io_UnixFileSystem_initIDs)
	VM.RegisterNative("java/io/UnixFileSystem.canonicalize0(Ljava/lang/String;)Ljava/lang/String;", JDK_java_io_UnixFileSystem_canonicalize0)
	VM.RegisterNative("java/io/UnixFileSystem.getBooleanAttributes0(Ljava/io/File;)I", JDK_java_io_UnixFileSystem_getBooleanAttributes0)
	VM.RegisterNative("java/io/UnixFileSystem.getLength(Ljava/io/File;)J", JDK_java_io_UnixFileSystem_getLength)
}

func JDK_java_io_UnixFileSystem_initIDs() {
	// do nothing
}

//@Native public static final int BA_EXISTS    = 0x01;
//@Native public static final int BA_REGULAR   = 0x02;
//@Native public static final int BA_DIRECTORY = 0x04;
//@Native public static final int BA_HIDDEN    = 0x08;
func JDK_java_io_UnixFileSystem_getBooleanAttributes0(this Reference, file Reference) Int {
	path := file.GetInstanceVariableByName("path", "Ljava/lang/String;").(JavaLangString).toNativeString()
	fileInfo, err := os.Stat(path)
	attr := 0
	if err == nil {
		attr |= 0x01
		if fileInfo.Mode().IsRegular() {
			attr |= 0x02
		}
		if fileInfo.Mode().IsDir() {
			attr |= 0x04
		}
		if hidden, err := IsHidden(path); hidden && err != nil {
			attr |= 0x08
		}
		return Int(attr)
	}

	VM.Throw("java/io/IOException", "Cannot get file attributes: %s", path)
	return -1

}

func IsHidden(filename string) (bool, error) {

	if runtime.GOOS != "windows" {

		// unix/linux file or directory that starts with . is hidden
		if filename[0:1] == "." {
			return true, nil

		} else {
			return false, nil
		}

	} else {
		log.Fatal("Unable to check if file is hidden under this OS")
	}
	return false, nil
}

func JDK_java_io_UnixFileSystem_canonicalize0(this Reference, path JavaLangString) JavaLangString {
	return VM.NewJavaLangString(filepath.Clean(path.toNativeString()))
}

func JDK_java_io_UnixFileSystem_getLength(this Reference, file Reference) Long {
	path := file.GetInstanceVariableByName("path", "Ljava/lang/String;").(JavaLangString).toNativeString()
	fileInfo, err := os.Stat(path)
	if err == nil {
		VM.ExecutionEngine.ioLogger.Info("📒    %s - length %d \n", path, fileInfo.Size())
		return Long(fileInfo.Size())
	}
	VM.Throw("java/io/IOException", "Cannot get file length: %s", path)
	return -1
}
