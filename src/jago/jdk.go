package jago

import (
	"reflect"
)

var NATIVE_METHODS = map[string]reflect.Value {
	//"Java_GVM_println": reflect.ValueOf(Java_GVM_println),
	//"Java_GVM_toUpper": reflect.ValueOf(Java_GVM_toUpper),
	//"Java_java_lang_Object_registerNatives": reflect.ValueOf(Java_java_lang_Object_registerNatives),
	//"Java_java_lang_System_registerNatives": reflect.ValueOf(Java_java_lang_System_registerNatives),
	//"Java_java_lang_Class_registerNatives": reflect.ValueOf(Java_java_lang_Class_registerNatives),
	//"Java_java_lang_Class_getPrimitiveClass": reflect.ValueOf(Java_java_lang_Class_getPrimitiveClass),
	//"Java_java_lang_Class_desiredAssertionStatus0": reflect.ValueOf(Java_java_lang_Class_desiredAssertionStatus0),
	//"Java_java_lang_Float_floatToRawIntBits": reflect.ValueOf(Java_java_lang_Float_floatToRawIntBits),
	//"Java_java_lang_Double_doubleToRawLongBits": reflect.ValueOf(Java_java_lang_Double_doubleToRawLongBits),
	//"Java_java_lang_Double_longBitsToDouble": reflect.ValueOf(Java_java_lang_Double_longBitsToDouble),
	//"Java_java_lang_System_arraycopy": reflect.ValueOf(Java_java_lang_System_arraycopy),
	//"Java_java_lang_Object_getClass": reflect.ValueOf(Java_java_lang_Object_getClass),
	//"Java_java_lang_Class_getName0": reflect.ValueOf(Java_java_lang_Class_getName0),
	//"Java_java_lang_Object_hashCode": reflect.ValueOf(Java_java_lang_Object_hashCode),
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

	///
	register_jago_Jago()
}
//
//
///* ---------- JDK Native methods implementation ---*/
//
//func Java_GVM_println(o ObjectRef, s JavaLangString) {
//	println(s.toGoString())
//}
//
//func Java_GVM_toUpper(s JavaLangString) JavaLangString {
//	return NewJavaLangString(strings.ToUpper(s.toGoString()))
//}
//
//func Java_java_lang_Object_registerNatives() {
//
//}
//
//func Java_java_lang_System_registerNatives() {
//
//}
//
//func Java_java_lang_Class_registerNatives()  {
//
//}
//
//func Java_java_lang_Class_getPrimitiveClass(jobject ObjectRef) ObjectRef {
//	return NewJavaLangClass()
//}
//
//func Java_java_lang_Class_desiredAssertionStatus0(clazz ObjectRef) Int {
//	return TRUE.ToInt()
//}
//
//func Java_java_lang_Float_floatToRawIntBits(f Float) Int {
//	return 1
//}
//
//func Java_java_lang_Double_doubleToRawLongBits(d Double) Long {
//	return Long(d)
//}
//
//func Java_java_lang_Double_longBitsToDouble(l Long) Double  {
//	return Double(l)
//}
//
//func Java_java_lang_System_arraycopy(src ObjectRef, srcPos Int, dest ObjectRef, destPos Int, length Int) {
//	srcObjectref := src.(Reference)
//	dstObjectref := dest.(Reference)
//	if !srcObjectref.Class().IsArray() || !dstObjectref.Class().IsArray() {
//		Throw("ArrayStoreException", "")
//	}
//	srcArrayRef := srcObjectref.AsArrayRef()
//	dstArrayRef := dstObjectref.AsArrayRef()
//
//	if srcPos + length > srcArrayRef.Length() || destPos + length > dstArrayRef.Length() {
//		Throw("ArrayIndexOutOfBoundsException", "")
//	}
//	for i := Int(0); i < length; i++ {
//		dstArrayRef.SetElement(destPos + i, srcArrayRef.GetElement(srcPos + i))
//	}
//}
//
//func Java_java_lang_Object_getClass(this Reference) JavaLangClass  {
//	return this.Class().classObject
//}
//
//func Java_java_lang_Class_getName0(this JavaLangClass) JavaLangString  {
//	for k, v := range BOOTSTRAP_CLASSLOADER.classCache {
//		if v.ClassObject() == this {
//			return NewJavaLangString(k)
//		}
//	}
//	return NewJavaLangString("")
//}
//
//func Java_java_lang_Object_hashCode(this Reference) Int  {
//	return 1
//}
