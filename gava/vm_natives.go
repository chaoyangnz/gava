package gava

import (
	"reflect"
)

type NativeMethodRegistry map[string]reflect.Value

func (this NativeMethodRegistry) RegisterNative(qualifier string, function interface{}) {
	this[qualifier] = reflect.ValueOf(function)
}

func (this NativeMethodRegistry) unregister(qualifier string, function interface{}) {
	delete(this, qualifier)
}

func (this NativeMethodRegistry) FindNative(qualifier string) (reflect.Value, bool) {
	fun, found := this[qualifier]
	return fun, found
}

func (this NativeMethodRegistry) RegisterNatives() {
	register_java_lang_System()
	register_java_lang_Object()
	register_java_lang_Class()
	register_java_lang_ClassLoader()
	register_java_lang_Package()

	register_java_lang_String()
	register_java_lang_Float()
	register_java_lang_Double()

	register_java_lang_Thread()
	register_java_lang_Throwable()
	register_java_lang_Runtime()

	register_java_lang_StrictMath()

	register_java_security_AccessController()

	register_java_lang_reflect_Array()

	register_sun_misc_VM()
	register_sun_misc_Unsafe()
	register_sun_reflect_Reflection()
	register_sun_reflect_NativeConstructorAccessorImpl()
	register_sun_misc_URLClassPath()

	register_java_io_FileDescriptor()
	register_java_io_FieInputStream()
	register_java_io_FieOutputStream()
	register_java_io_UnixFileSystem()

	register_java_util_concurrent_atomic_AtomicLong()

	register_java_util_zip_ZipFile()
	register_java_util_TimeZone()
}
