package jago

func register_java_lang_System()  {
	VM.RegisterNative("java/lang/System.registerNatives()V", JDK_java_lang_System_registerNatives)
	VM.RegisterNative("java/lang/System.setIn0(Ljava/io/InputStream;)V", JDK_java_lang_System_setIn0)
	VM.RegisterNative("java/lang/System.setOut0(Ljava/io/PrintStream;)V", JDK_java_lang_System_setOut0)
	VM.RegisterNative("java/lang/System.setErr0(Ljava/io/PrintStream;)V", JDK_java_lang_System_setErr0)
	VM.RegisterNative("java/lang/System.currentTimeMillis()J", JDK_java_lang_System_currentTimeMillis)
	VM.RegisterNative("java/lang/System.nanoTime()J", JDK_java_lang_System_nanoTime)
	VM.RegisterNative("java/lang/System.arraycopy(Ljava/lang/Object;ILjava/lang/Object;II)V", JDK_java_lang_System_arraycopy)
	VM.RegisterNative("java/lang/System.identityHashCode(Ljava/lang/Object;)I", JDK_java_lang_System_identityHashCode)
	VM.RegisterNative("java/lang/System.initProperties(Ljava/util/Properties;)Ljava/util/Properties;", JDK_java_lang_System_initProperties)
	VM.RegisterNative("java/lang/System.mapLibraryName(Ljava/lang/String;)Ljava/lang/String;", JDK_java_lang_System_mapLibraryName)
}

// private static void registers()
func JDK_java_lang_System_registerNatives() {}

// private static void setIn0(InputStream is)
func JDK_java_lang_System_setIn0(is ObjectRef) {
	VM.SetStaticVariable(VM.GetClass("java/lang/System"), "in", "Ljava/io/InputStream;", is)
}

// private static void setOut0(PrintStream ps)
func JDK_java_lang_System_setOut0(ps ObjectRef)  {
	VM.SetStaticVariable(VM.GetClass("java/lang/System"), "out", "Ljava/io/PrintStream;", ps)
}

// private static void setErr0(PrintStream ps)
func JDK_java_lang_System_setErr0(ps ObjectRef)  {
	VM.SetStaticVariable(VM.GetClass("java/lang/System"), "err", "Ljava/io/PrintStream;", ps)
}

// public static long currentTimeMillis()
func JDK_java_lang_System_currentTimeMillis() Long {
	return VM.CurrentTimeMillis()
}

// public static long nanoTime()
func JDK_java_lang_System_nanoTime() Long {
	return VM.CurrentTimeNano()
}

// public static void arraycopy(Object fromArray, int fromIndex, Object toArray, int toIndex, int length)
func JDK_java_lang_System_arraycopy(src Reference, srcPos Int, dest Reference, destPos Int, length Int) {
	if !src.Class().IsArray() || !dest.Class().IsArray() {
		VM.Throw("java/lang/ArrayStoreException", "")
	}
	srcArrayRef := src.AsArrayRef()
	dstArrayRef := dest.AsArrayRef()

	if srcPos + length > srcArrayRef.ArrayLength() || destPos + length > dstArrayRef.ArrayLength() {
		VM.Throw("java/lang/ArrayIndexOutOfBoundsException", "")
	}
	for i := Int(0); i < length; i++ {
		dstArrayRef.SetArrayElement(destPos + i, srcArrayRef.GetArrayElement(srcPos + i))
	}
}

// public static int identityHashCode(Object object)
func JDK_java_lang_System_identityHashCode(object Reference) Int {
	return VM.IHashCode(object)
}

// private static Properties initProperties(Properties properties)
func JDK_java_lang_System_initProperties(properties ObjectRef) ObjectRef {

	//TODO
	setProperty := properties.Class().GetMethod("setProperty", "(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")
	for key, val := range VM.ListSystemProperties() {
		VM.InvokeMethod(setProperty, properties, VM.NewJavaLangString(key), VM.NewJavaLangString(val))
	}

	return properties
}

func JDK_java_lang_System_mapLibraryName(name JavaLangString) JavaLangString {
	return name
}
