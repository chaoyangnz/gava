package jago

func register_java_util_zip_ZipFile() {
	VM.RegisterNative("java/util/zip/ZipFile.initIDs()V", JDK_java_util_zip_ZipFile_initIDs)
}

func JDK_java_util_zip_ZipFile_initIDs() {
	//DO NOTHING
}
