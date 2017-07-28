package jago

import "runtime"

func register_java_lang_System()  {
	register("java/lang/System.registerNatives()V", JDK_java_lang_System_registerNatives)
	register("java/lang/System.setIn0(Ljava/io/InputStream;)V", JDK_java_lang_System_setIn0)
	register("java/lang/System.setOut0(Ljava/io/PrintStream;)V", JDK_java_lang_System_setOut0)
	register("java/lang/System.setErr0(Ljava/io/PrintStream;)V", JDK_java_lang_System_setErr0)
	register("java/lang/System.currentTimeMillis()J", JDK_java_lang_System_currentTimeMillis)
	register("java/lang/System.nanoTime()J", JDK_java_lang_System_nanoTime)
	register("java/lang/System.arraycopy(Ljava/lang/Object;ILjava/lang/Object;II)V", JDK_java_lang_System_arraycopy)
	register("java/lang/System.identityHashCode(Ljava/lang/Object;)I", JDK_java_lang_System_identityHashCode)
	register("java/lang/System.initProperties(Ljava/util/Properties;)Ljava/util/Properties;", JDK_java_lang_System_initProperties)
	register("java/lang/System.mapLibraryName(Ljava/lang/String;)Ljava/lang/String;", JDK_java_lang_System_mapLibraryName)
}

// private static void registers()
func JDK_java_lang_System_registerNatives() {}

// private static void setIn0(InputStream is)
func JDK_java_lang_System_setIn0(is ObjectRef) {
	VM_setStaticVariable(VM_getClass("java/lang/System"), "in", "Ljava/io/InputStream;", is)
}

// private static void setOut0(PrintStream ps)
func JDK_java_lang_System_setOut0(ps ObjectRef)  {
	VM_setStaticVariable(VM_getClass("java/lang/System"), "out", "Ljava/io/PrintStream;", ps)
}

// private static void setErr0(PrintStream ps)
func JDK_java_lang_System_setErr0(ps ObjectRef)  {
	VM_setStaticVariable(VM_getClass("java/lang/System"), "err", "Ljava/io/PrintStream;", ps)
}

// public static long currentTimeMillis()
func JDK_java_lang_System_currentTimeMillis() Long {
	return VM_currentTimeMillis()
}

// public static long nanoTime()
func JDK_java_lang_System_nanoTime() Long {
	return VM_currentTimeNano()
}

// public static void arraycopy(Object fromArray, int fromIndex, Object toArray, int toIndex, int length)
func JDK_java_lang_System_arraycopy(src Reference, srcPos Int, dest Reference, destPos Int, length Int) {
	if !src.Class().IsArray() || !dest.Class().IsArray() {
		VM_throw("java/lang/ArrayStoreException", "")
	}
	srcArrayRef := src.AsArrayRef()
	dstArrayRef := dest.AsArrayRef()

	if srcPos + length > srcArrayRef.Length() || destPos + length > dstArrayRef.Length() {
		VM_throw("java/lang/ArrayIndexOutOfBoundsException", "")
	}
	for i := Int(0); i < length; i++ {
		dstArrayRef.SetElement(destPos + i, srcArrayRef.GetElement(srcPos + i))
	}
}

// public static int identityHashCode(Object object)
func JDK_java_lang_System_identityHashCode(object Reference) Int {
	return VM_iHashCode(object)
}

// private static Properties initProperties(Properties properties)
func JDK_java_lang_System_initProperties(properties ObjectRef) ObjectRef {

	//TODO
	//setProperty := properties.Class().GetMethod("setProperty", "(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")
	//VM_invokeMethod0(THREAD_MANAGER.current, setProperty, properties, NewJavaLangString("file.encoding"), NewJavaLangString("UTF-8"))
	//for key, val := range _sysProps {
	//	VM_invokeMethod0(THREAD_MANAGER.current, setProperty, properties, NewJavaLangString(key), NewJavaLangString(val))
	//}

	return properties
}

var _sysProps map[string]string = map[string]string{
		"java.version":         "1.8.0",
		"java.vendor":          "jago",
		"java.vendor.url":      "https://yangchao.me",
		"java.home":            "todo",
		"java.class.version":   "52.0",
		"java.class.path":      "todo",
		"java.awt.graphicsenv": "sun.awt.CGraphicsEnvironment",
		"os.name":              runtime.GOOS,   // todo
		"os.arch":              runtime.GOARCH, // todo
		"os.version":           "",             // todo
		"file.separator":       "/",            // todo os.PathSeparator
		"path.separator":       ":",            // todo os.PathListSeparator
		"line.separator":       "\n",           // todo
		"user.name":            "",             // todo
		"user.home":            "",             // todo
		"user.dir":             ".",            // todo
		"user.country":         "CN",           // todo
		"file.encoding":        "UTF-8",
		"sun.stdout.encoding":  "UTF-8",
		"sun.stderr.encoding":  "UTF-8",
}

func JDK_java_lang_System_mapLibraryName(name JavaLangString) JavaLangString {
	return name
}
