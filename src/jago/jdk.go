package jago

import (
	"reflect"
)

var NATIVE_METHODS = map[string]reflect.Value {
}

func register(qualifier string, function interface{})  {
	NATIVE_METHODS[qualifier] = reflect.ValueOf(function)
}

func unregister(qualifier string, function interface{})  {
	delete(NATIVE_METHODS, qualifier)
}

func findNative(qualifier string) (reflect.Value, bool) {
	fun, found := NATIVE_METHODS[qualifier]
	return fun, found
}

func RegisterNatives()  {
	register_java_lang_System()
	register_java_lang_Object()
	register_java_lang_Class()
	register_java_lang_ClassLoader()

	register_java_lang_String()
	register_java_lang_Float()
	register_java_lang_Double()

	register_java_lang_Thread()
	register_java_lang_Throwable()
	register_java_lang_Runtime()

	register_java_lang_StrictMath()

	register_java_security_AccessController()



	register_sun_misc_VM()
	register_sun_misc_Unsafe()
	register_sun_reflect_Reflection()
	register_sun_reflect_NativeConstructorAccessorImpl()

	register_java_io_FileDescriptor()
	register_java_io_FieInputStream()
	register_java_io_FieOutputStream()
	register_java_io_UnixFileSystem()

	register_java_util_concurrent_atomic_AtomicLong()
}
