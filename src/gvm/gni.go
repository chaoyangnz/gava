// Copyright 2016 Tim O'Brien. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gvm

/*
#cgo LDFLAGS: -I${SRCDIR}/gvm
#include <jni.h>
*/
import "C"
import "unsafe"

/* ---------- JDK Native methods implementation ---*/
func GVM_print(s java_lang_string) {
	println(s.toString())
}


//jint(*GetVersion)(JNIEnv *env);
//export GNI_GetVersion
func GNI_GetVersion(env *C.JNIEnv) C.jint {
	fatal("Not implemented")
	return C.jint(0)
}
//jclass(*DefineClass)(JNIEnv *env, const char *name, jobject loader, const jbyte *buf,jsize len);
//export GNI_DefineClass
func GNI_DefineClass(env *C.JNIEnv, name *C.char, loader C.jobject, buf *C.jbyte, len C.jsize) C.jclass {
	fatal("Not implemented")
	return nil
}
//jclass(*FindClass)(JNIEnv *env, const char *name);
//export GNI_FindClass
func GNI_FindClass(env *C.JNIEnv, name *C.char) C.jclass {
	fatal("Not implemented")
	return nil
}
//jmethodID(*FromReflectedMethod)(JNIEnv *env, jobject method);
//export GNI_FromReflectedMethod
func GNI_FromReflectedMethod(env *C.JNIEnv, method C.jobject) C.jmethodID {
	fatal("Not implemented")
	return nil
}
//jfieldID(*FromReflectedField)(JNIEnv *env, jobject field);
//export GNI_FromReflectedField
func GNI_FromReflectedField(env *C.JNIEnv, field C.jobject) C.jfieldID {
	fatal("Not implemented")
	return nil
}
//jobject(*ToReflectedMethod)(JNIEnv *env, jclass cls, jmethodID methodID, jboolean isStatic);
//export GNI_ToReflectedMethod
func GNI_ToReflectedMethod(env *C.JNIEnv, cls C.jclass, methodID C.jmethodID, isStatic C.jboolean) C.jobject {
	fatal("Not implemented")
	return nil
}
//jclass(*GetSuperclass)(JNIEnv *env, jclass sub);
//export GNI_GetSuperclass
func GNI_GetSuperclass(env *C.JNIEnv, sub C.jclass) C.jclass {
	fatal("Not implemented")
	return nil
}
//jboolean(*IsAssignableFrom)(JNIEnv *env, jclass sub, jclass sup);
//export GNI_IsAssignableFrom
func GNI_IsAssignableFrom(env *C.JNIEnv,sub C.jclass, sup C.jclass) C.jboolean {
	fatal("Not implemented")
	return C.jboolean(0)
}
//jobject(*ToReflectedField)(JNIEnv *env, jclass cls, jfieldID fieldID, jboolean isStatic);
//export GNI_ToReflectedField
func GNI_ToReflectedField(env *C.JNIEnv, cls C.jclass, fieldID C.jfieldID, isStatic C.jboolean) C.jobject {
	fatal("Not implemented")
	return nil
}
//jint(*Throw)(JNIEnv *env, jthrowable obj);
//export GNI_Throw
func GNI_Throw(env *C.JNIEnv, obj C.jthrowable) C.jint {
	fatal("Not implemented")
	return C.jint(0)
}
//jint(*ThrowNew)(JNIEnv *env, jclass clazz, const char *msg);
//export GNI_ThrowNew
func GNI_ThrowNew(env *C.JNIEnv, clazz C.jclass, msg *C.char) C.jint {
	fatal("Not implemented")
	return C.jint(0)
}
//jthrowable(*ExceptionOccurred)(JNIEnv *env);
//export GNI_ExceptionOccurred
func GNI_ExceptionOccurred(env *C.JNIEnv) C.jthrowable {
	fatal("Not implemented")
	return nil
}
//void(*ExceptionDescribe)(JNIEnv *env);
//export GNI_ExceptionDescribe
func GNI_ExceptionDescribe(env *C.JNIEnv) {
	fatal("Not implemented")
}
//void(*ExceptionClear)(JNIEnv *env);
//export GNI_ExceptionClear
func GNI_ExceptionClear(env *C.JNIEnv) {
	fatal("Not implemented")
}
//void(*FatalError)(JNIEnv *env, const char *msg);
//export GNI_FatalError
func GNI_FatalError(env *C.JNIEnv, msg  *C.char) {
	fatal("Not implemented")
}
//jint(*PushLocalFrame)(JNIEnv *env, jint capacity);
//export GNI_PushLocalFrame
func GNI_PushLocalFrame(env *C.JNIEnv, capacity C.jint) C.jint {
	fatal("Not implemented")
	return C.jint(0)
}
//jobject(*PopLocalFrame)(JNIEnv *env, jobject result);
//export GNI_PopLocalFrame
func GNI_PopLocalFrame(env *C.JNIEnv, result C.jobject) C.jobject {
	fatal("Not implemented")
	return nil
}
//jobject(*NewGlobalRef)(JNIEnv *env, jobject lobj);
//export GNI_NewGlobalRef
func GNI_NewGlobalRef(env *C.JNIEnv, lobj C.jobject) C.jobject {
	fatal("Not implemented")
	return nil
}
//void(*DeleteGlobalRef)(JNIEnv *env, jobject gref);
//export GNI_DeleteGlobalRef
func GNI_DeleteGlobalRef(env *C.JNIEnv, gref C.jobject) {
	fatal("Not implemented")
}
//void(*DeleteLocalRef)(JNIEnv *env, jobject obj);
//export GNI_DeleteLocalRef
func GNI_DeleteLocalRef(env *C.JNIEnv, obj C.jobject) {
	fatal("Not implemented")
}
//jboolean(*IsSameObject)(JNIEnv *env, jobject obj1, jobject obj2);
//export GNI_IsSameObject
func GNI_IsSameObject(env *C.JNIEnv, obj1 C.jobject, obj2 C.jobject) C.jboolean {
	fatal("Not implemented")
	return C.jboolean(0)
}
//jobject(*NewLocalRef)(JNIEnv *env, jobject ref);
//export GNI_NewLocalRef
func GNI_NewLocalRef(env *C.JNIEnv, ref C.jobject) C.jobject {
	fatal("Not implemented")
	return nil
}
//jint(*EnsureLocalCapacity)(JNIEnv *env, jint capacity);
//export GNI_EnsureLocalCapacity
func GNI_EnsureLocalCapacity(env *C.JNIEnv, capacity C.jint) C.jint {
	fatal("Not implemented")
	return C.jint(0)
}
//jobject(*AllocObject)(JNIEnv *env, jclass clazz);
//export GNI_AllocObject
func GNI_AllocObject(env *C.JNIEnv, clazz C.jclass) C.jobject {
	fatal("Not implemented")
	return nil
}
//jobject(*NewObject)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//func GNI_NewObject(env *C.JNIEnv, clazz C.jclass, methodID ...C.jmethodID) C.jobject {
//  fatal("Not implemented")
//  return nil
//}
//jobject(*NewObjectV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//export GNI_NewObjectV
func GNI_NewObjectV(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID,  args *C.va_list) C.jobject {
	fatal("Not implemented")
	return nil
}
//jobject(*NewObjectA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
//export GNI_NewObjectA
func GNI_NewObjectA(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jobject {
	fatal("Not implemented")
	return nil
}
//jclass(*GetObjectClass)(JNIEnv *env, jobject obj);
//export GNI_GetObjectClass
func GNI_GetObjectClass(env *C.JNIEnv, obj C.jobject) C.jclass {
	fatal("Not implemented")
	return nil
}
//jboolean(*IsInstanceOf)(JNIEnv *env, jobject obj, jclass clazz);
//export GNI_IsInstanceOf
func GNI_IsInstanceOf(env *C.JNIEnv, obj C.jobject, clazz C.jclass) C.jboolean {
	fatal("Not implemented")
	return C.jboolean(0)
}
//jmethodID(*GetMethodID)(JNIEnv *env, jclass clazz, const char *name, const char *sig);
//export GNI_GetMethodID
func GNI_GetMethodID(env *C.JNIEnv, clazz C.jclass, name *C.char, sig *C.char) C.jmethodID {
	fatal("Not implemented")
	return nil
}
//jobject(*CallObjectMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//func GNI_CallObjectMethod(env *C.JNIEnv, obj C.jobject, methodID ...C.jmethodID) C.jobject {
//  fatal("Not implemented")
//  return nil
//}
//jobject(*CallObjectMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//export GNI_CallObjectMethodV
func GNI_CallObjectMethodV(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID,  args *C.va_list) C.jobject {
	fatal("Not implemented")
	return nil
}
//jobject(*CallObjectMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue * args);
//export GNI_CallObjectMethodA
func GNI_CallObjectMethodA(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID,  args *C.jvalue) C.jobject {
	fatal("Not implemented")
	return nil
}
//jboolean(*CallBooleanMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//func GNI_CallBooleanMethod(env *C.JNIEnv, obj C.jobject, methodID ...C.jmethodID) C.jboolean {
//  fatal("Not implemented")
//  return nil
//}
//jboolean(*CallBooleanMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//export GNI_CallBooleanMethodV
func GNI_CallBooleanMethodV(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args *C.va_list) C.jboolean {
	fatal("Not implemented")
	return C.jboolean(0)
}
//jboolean(*CallBooleanMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue * args);
//export GNI_CallBooleanMethodA
func GNI_CallBooleanMethodA(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args *C.jvalue) C.jboolean {
	fatal("Not implemented")
	return C.jboolean(0)
}
//jbyte(*CallByteMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//func GNI_CallByteMethod(env *C.JNIEnv, obj C.jobject, methodID ...C.jmethodID) C.jbyte {
//  fatal("Not implemented")
//  return nil
//}
//jbyte(*CallByteMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//export GNI_CallByteMethodV
func GNI_CallByteMethodV(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args *C.va_list) C.jbyte {
	fatal("Not implemented")
	return C.jbyte(0)
}
//jbyte(*CallByteMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);
//export GNI_CallByteMethodA
func GNI_CallByteMethodA(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args *C.jvalue) C.jbyte {
	fatal("Not implemented")
	return C.jbyte(0)
}
//jchar(*CallCharMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//func GNI_CallCharMethod(env *C.JNIEnv, obj C.jobject, methodID ...C.jmethodID) C.jchar {
//  fatal("Not implemented")
//  return nil
//}
//jchar(*CallCharMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//export GNI_CallCharMethodV
func GNI_CallCharMethodV(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args *C.va_list) C.jchar {
	fatal("Not implemented")
	return C.jchar(0)
}
//jchar(*CallCharMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);
//export GNI_CallCharMethodA
func GNI_CallCharMethodA(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args *C.jvalue) C.jchar {
	fatal("Not implemented")
	return C.jchar(0)
}
//jshort(*CallShortMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//func GNI_CallShortMethod(env *C.JNIEnv, obj C.jobject, methodID ...C.jmethodID) C.jshort {
//  fatal("Not implemented")
//  return nil
//}
//jshort(*CallShortMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//export GNI_CallShortMethodV
func GNI_CallShortMethodV(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args *C.va_list) C.jshort {
	fatal("Not implemented")
	return C.jshort(0)
}
//jshort(*CallShortMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);
//export GNI_CallShortMethodA
func GNI_CallShortMethodA(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args *C.jvalue) C.jshort {
	fatal("Not implemented")
	return C.jshort(0)
}
//jint(*CallIntMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//func GNI_CallIntMethod(env *C.JNIEnv, obj C.jobject, methodID ...C.jmethodID) C.jint {
//  fatal("Not implemented")
//  return nil
//}
//jint(*CallIntMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//export GNI_CallIntMethodV
func GNI_CallIntMethodV(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID,  args *C.va_list) C.jint {
	fatal("Not implemented")
	return C.jint(0)
}
//jint(*CallIntMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);
//export GNI_CallIntMethodA
func GNI_CallIntMethodA(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args *C.jvalue) C.jint {
	fatal("Not implemented")
	return C.jint(0)
}
//jlong(*CallLongMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//func GNI_CallLongMethod(env *C.JNIEnv, obj C.jobject, methodID ...C.jmethodID) C.jlong {
//  fatal("Not implemented")
//  return nil
//}
//jlong(*CallLongMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//export GNI_CallLongMethodV
func GNI_CallLongMethodV(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID,  args *C.va_list) C.jlong {
	fatal("Not implemented")
	return C.jlong(0)
}
//jlong(*CallLongMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);
//export GNI_CallLongMethodA
func GNI_CallLongMethodA(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args *C.jvalue) C.jlong {
	fatal("Not implemented")
	return C.jlong(0)
}
//jfloat(*CallFloatMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//func GNI_CallFloatMethod(env *C.JNIEnv, obj C.jobject, methodID ...C.jmethodID) C.jfloat {
//  fatal("Not implemented")
//  return nil
//}
//jfloat(*CallFloatMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//export GNI_CallFloatMethodV
func GNI_CallFloatMethodV(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID,  args *C.va_list) C.jfloat {
	fatal("Not implemented")
	return C.jfloat(0)
}
//jfloat(*CallFloatMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);
//export GNI_CallFloatMethodA
func GNI_CallFloatMethodA(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args *C.jvalue) C.jfloat {
	fatal("Not implemented")
	return C.jfloat(0)
}
//jdouble(*CallDoubleMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//func GNI_CallDoubleMethod(env *C.JNIEnv, obj C.jobject, methodID ...C.jmethodID) C.jdouble {
//  fatal("Not implemented")
//  return nil
//}
//jdouble(*CallDoubleMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//export GNI_CallDoubleMethodV
func GNI_CallDoubleMethodV(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID,  args *C.va_list) C.jdouble {
	fatal("Not implemented")
	return C.jdouble(0)
}
//jdouble(*CallDoubleMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);
//export GNI_CallDoubleMethodA
func GNI_CallDoubleMethodA(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args *C.jvalue) C.jdouble {
	fatal("Not implemented")
	return C.jdouble(0)
}
//void(*CallVoidMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//func GNI_CallVoidMethod(env *C.JNIEnv, obj C.jobject, methodID ...C.jmethodID) {
//  fatal("Not implemented")
//}
//void(*CallVoidMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//export GNI_CallVoidMethodV
func GNI_CallVoidMethodV(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID,  args *C.va_list) {
	fatal("Not implemented")
}
//void(*CallVoidMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue * args);
//export GNI_CallVoidMethodA
func GNI_CallVoidMethodA(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args *C.jvalue) {
	fatal("Not implemented")
}
//jobject(*CallNonvirtualObjectMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//func GNI_CallNonvirtualObjectMethod(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID ...C.jmethodID) C.jobject {
//  fatal("Not implemented")
//  return nil
//}
//jobject(*CallNonvirtualObjectMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,va_list args);
//export GNI_CallNonvirtualObjectMethodV
func GNI_CallNonvirtualObjectMethodV(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args *C.va_list) C.jobject {
	fatal("Not implemented")
	return nil
}
//jobject(*CallNonvirtualObjectMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,const jvalue * args);
//export GNI_CallNonvirtualObjectMethodA
func GNI_CallNonvirtualObjectMethodA(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID,args *C.jvalue) C.jobject {
	fatal("Not implemented")
	return nil
}
//jboolean(*CallNonvirtualBooleanMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//func GNI_CallNonvirtualBooleanMethod(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID ...C.jmethodID) C.jboolean {
//  fatal("Not implemented")
//  return nil
//}
//jboolean(*CallNonvirtualBooleanMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,va_list args);
//export GNI_CallNonvirtualBooleanMethodV
func GNI_CallNonvirtualBooleanMethodV(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args *C.va_list) C.jboolean {
	fatal("Not implemented")
	return C.jboolean(0)
}
//jboolean(*CallNonvirtualBooleanMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,const jvalue * args);
//export GNI_CallNonvirtualBooleanMethodA
func GNI_CallNonvirtualBooleanMethodA(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID,args *C.jvalue) C.jboolean {
	fatal("Not implemented")
	return C.jboolean(0)
}
//jbyte(*CallNonvirtualByteMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//func GNI_CallNonvirtualByteMethod(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID ...C.jmethodID) C.jbyte {
//  fatal("Not implemented")
//  return nil
//}
//jbyte(*CallNonvirtualByteMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,va_list args);
//export GNI_CallNonvirtualByteMethodV
func GNI_CallNonvirtualByteMethodV(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args *C.va_list) C.jbyte {
	fatal("Not implemented")
	return C.jbyte(0)
}
//jbyte(*CallNonvirtualByteMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,const jvalue *args);
//export GNI_CallNonvirtualByteMethodA
func GNI_CallNonvirtualByteMethodA(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID,args *C.jvalue) C.jbyte {
	fatal("Not implemented")
	return C.jbyte(0)
}
//jchar(*CallNonvirtualCharMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//func GNI_CallNonvirtualCharMethod(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID ...C.jmethodID) C.jchar {
//  fatal("Not implemented")
//  return nil
//}
//jchar(*CallNonvirtualCharMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,va_list args);
//export GNI_CallNonvirtualCharMethodV
func GNI_CallNonvirtualCharMethodV(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args *C.va_list) C.jchar {
	fatal("Not implemented")
	return C.jchar(0)
}
//jchar(*CallNonvirtualCharMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,const jvalue *args);
//export GNI_CallNonvirtualCharMethodA
func GNI_CallNonvirtualCharMethodA(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID,args *C.jvalue) C.jchar {
	fatal("Not implemented")
	return C.jchar(0)
}
//jshort(*CallNonvirtualShortMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//func GNI_CallNonvirtualShortMethod(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID ...C.jmethodID) C.jshort {
//  fatal("Not implemented")
//  return nil
//}
//jshort(*CallNonvirtualShortMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,va_list args);
//export GNI_CallNonvirtualShortMethodV
func GNI_CallNonvirtualShortMethodV(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args *C.va_list) C.jshort {
	fatal("Not implemented")
	return C.jshort(0)
}
//jshort(*CallNonvirtualShortMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,const jvalue *args);
//export GNI_CallNonvirtualShortMethodA
func GNI_CallNonvirtualShortMethodA(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID,args *C.jvalue) C.jshort {
	fatal("Not implemented")
	return C.jshort(0)
}
//jint(*CallNonvirtualIntMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//func GNI_CallNonvirtualIntMethod(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID ...C.jmethodID) C.jint {
//  fatal("Not implemented")
//  return nil
//}
//jint(*CallNonvirtualIntMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,va_list args);
//export GNI_CallNonvirtualIntMethodV
func GNI_CallNonvirtualIntMethodV(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args *C.va_list) C.jint {
	fatal("Not implemented")
	return C.jint(0)
}
//jint(*CallNonvirtualIntMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,const jvalue *args);
//export GNI_CallNonvirtualIntMethodA
func GNI_CallNonvirtualIntMethodA(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID,args *C.jvalue) C.jint {
	fatal("Not implemented")
	return C.jint(0)
}
//jlong(*CallNonvirtualLongMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//func GNI_CallNonvirtualLongMethod(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID ...C.jmethodID) C.jlong {
//  fatal("Not implemented")
//  return nil
//}
//jlong(*CallNonvirtualLongMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,va_list args);
//export GNI_CallNonvirtualLongMethodV
func GNI_CallNonvirtualLongMethodV(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args *C.va_list) C.jlong {
	fatal("Not implemented")
	return C.jlong(0)
}
//jlong(*CallNonvirtualLongMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,const jvalue *args);
//export GNI_CallNonvirtualLongMethodA
func GNI_CallNonvirtualLongMethodA(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID,args *C.jvalue) C.jlong {
	fatal("Not implemented")
	return C.jlong(0)
}
//jfloat(*CallNonvirtualFloatMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//func GNI_CallNonvirtualFloatMethod(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID ...C.jmethodID) C.jfloat {
//  fatal("Not implemented")
//  return nil
//}
//jfloat(*CallNonvirtualFloatMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,va_list args);
//export GNI_CallNonvirtualFloatMethodV
func GNI_CallNonvirtualFloatMethodV(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args *C.va_list) C.jfloat {
	fatal("Not implemented")
	return C.jfloat(0)
}
//jfloat(*CallNonvirtualFloatMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,const jvalue *args);
//export GNI_CallNonvirtualFloatMethodA
func GNI_CallNonvirtualFloatMethodA(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID,args *C.jvalue) C.jfloat {
	fatal("Not implemented")
	return C.jfloat(0)
}
//jdouble(*CallNonvirtualDoubleMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//func GNI_CallNonvirtualDoubleMethod(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID ...C.jmethodID) C.jdouble {
//  fatal("Not implemented")
//  return nil
//}
//jdouble(*CallNonvirtualDoubleMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,va_list args);
//export GNI_CallNonvirtualDoubleMethodV
func GNI_CallNonvirtualDoubleMethodV(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args *C.va_list) C.jdouble {
	fatal("Not implemented")
	return C.jdouble(0)
}
//jdouble(*CallNonvirtualDoubleMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,const jvalue *args);
//export GNI_CallNonvirtualDoubleMethodA
func GNI_CallNonvirtualDoubleMethodA(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID,args *C.jvalue) C.jdouble {
	fatal("Not implemented")
	return C.jdouble(0)
}
//void(*CallNonvirtualVoidMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//func GNI_CallNonvirtualVoidMethod(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID ...C.jmethodID) {
//  fatal("Not implemented")
//}
//void(*CallNonvirtualVoidMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,va_list args);
//export GNI_CallNonvirtualVoidMethodV
func GNI_CallNonvirtualVoidMethodV(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args *C.va_list) {
	fatal("Not implemented")
}
//void(*CallNonvirtualVoidMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,const jvalue * args);
//export GNI_CallNonvirtualVoidMethodA
func GNI_CallNonvirtualVoidMethodA(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID,args *C.jvalue) {
	fatal("Not implemented")
}
//jfieldID(*GetFieldID)(JNIEnv *env, jclass clazz, const char *name, const char *sig);
//export GNI_GetFieldID
func GNI_GetFieldID(env *C.JNIEnv, clazz C.jclass, name *C.char, sig *C.char) C.jfieldID {
	fatal("Not implemented")
	return nil
}
//jobject(*GetObjectField)(JNIEnv *env, jobject obj, jfieldID fieldID);
//export GNI_GetObjectField
func GNI_GetObjectField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID) C.jobject {
	fatal("Not implemented")
	return nil
}
//jboolean(*GetBooleanField)(JNIEnv *env, jobject obj, jfieldID fieldID);
//export GNI_GetBooleanField
func GNI_GetBooleanField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID) C.jboolean {
	fatal("Not implemented")
	return C.jboolean(0)
}
//jbyte(*GetByteField)(JNIEnv *env, jobject obj, jfieldID fieldID);
//export GNI_GetByteField
func GNI_GetByteField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID) C.jbyte {
	fatal("Not implemented")
	return C.jbyte(0)
}
//jchar(*GetCharField)(JNIEnv *env, jobject obj, jfieldID fieldID);
//export GNI_GetCharField
func GNI_GetCharField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID) C.jchar {
	fatal("Not implemented")
	return C.jchar(0)
}
//jshort(*GetShortField)(JNIEnv *env, jobject obj, jfieldID fieldID);
//export GNI_GetShortField
func GNI_GetShortField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID) C.jshort {
	fatal("Not implemented")
	return C.jshort(0)
}
//jint(*GetIntField)(JNIEnv *env, jobject obj, jfieldID fieldID);
//export GNI_GetIntField
func GNI_GetIntField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID) C.jint {
	fatal("Not implemented")
	return C.jint(0)
}
//jlong(*GetLongField)(JNIEnv *env, jobject obj, jfieldID fieldID);
//export GNI_GetLongField
func GNI_GetLongField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID) C.jlong {
	fatal("Not implemented")
	return C.jlong(0)
}
//jfloat(*GetFloatField)(JNIEnv *env, jobject obj, jfieldID fieldID);
//export GNI_GetFloatField
func GNI_GetFloatField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID) C.jfloat {
	fatal("Not implemented")
	return C.jfloat(0)
}
//jdouble(*GetDoubleField)(JNIEnv *env, jobject obj, jfieldID fieldID);
//export GNI_GetDoubleField
func GNI_GetDoubleField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID) C.jdouble {
	fatal("Not implemented")
	return C.jdouble(0)
}
//void(*SetObjectField)(JNIEnv *env, jobject obj, jfieldID fieldID, jobject val);
//export GNI_SetObjectField
func GNI_SetObjectField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID, val C.jobject) {
	fatal("Not implemented")
}
//void(*SetBooleanField)(JNIEnv *env, jobject obj, jfieldID fieldID, jboolean val);
//export GNI_SetBooleanField
func GNI_SetBooleanField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID, val C.jboolean) {
	fatal("Not implemented")
}
//void(*SetByteField)(JNIEnv *env, jobject obj, jfieldID fieldID, jbyte val);
//export GNI_SetByteField
func GNI_SetByteField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID, val C.jbyte) {
	fatal("Not implemented")
}
//void(*SetCharField)(JNIEnv *env, jobject obj, jfieldID fieldID, jchar val);
//export GNI_SetCharField
func GNI_SetCharField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID, val C.jchar) {
	fatal("Not implemented")
}
//void(*SetShortField)(JNIEnv *env, jobject obj, jfieldID fieldID, jshort val);
//export GNI_SetShortField
func GNI_SetShortField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID, val C.jshort) {
	fatal("Not implemented")
}
//void(*SetIntField)(JNIEnv *env, jobject obj, jfieldID fieldID, jint val);
//export GNI_SetIntField
func GNI_SetIntField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID, val C.jint) {
	fatal("Not implemented")
}
//void(*SetLongField)(JNIEnv *env, jobject obj, jfieldID fieldID, jlong val);
//export GNI_SetLongField
func GNI_SetLongField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID, val C.jlong) {
	fatal("Not implemented")
}
//void(*SetFloatField)(JNIEnv *env, jobject obj, jfieldID fieldID, jfloat val);
//export GNI_SetFloatField
func GNI_SetFloatField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID, val C.jfloat) {
	fatal("Not implemented")
}
//void(*SetDoubleField)(JNIEnv *env, jobject obj, jfieldID fieldID, jdouble val);
//export GNI_SetDoubleField
func GNI_SetDoubleField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID, val C.jdouble) {
	fatal("Not implemented")
}
//jmethodID(*GetStaticMethodID)(JNIEnv *env, jclass clazz, const char *name, const char *sig);
//export GNI_GetStaticMethodID
func GNI_GetStaticMethodID(env *C.JNIEnv, clazz C.jclass, name *C.char, sig *C.char) C.jmethodID {
	fatal("Not implemented")
	return nil
}
//jobject(*CallStaticObjectMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//func GNI_CallStaticObjectMethod(env *C.JNIEnv, clazz C.jclass, methodID ...C.jmethodID) C.jobject {
//  fatal("Not implemented")
//  return nil
//}
//jobject(*CallStaticObjectMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//export GNI_CallStaticObjectMethodV
func GNI_CallStaticObjectMethodV(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID,  args *C.va_list) C.jobject {
	fatal("Not implemented")
	return nil
}
//jobject(*CallStaticObjectMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
//export GNI_CallStaticObjectMethodA
func GNI_CallStaticObjectMethodA(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jobject {
	fatal("Not implemented")
	return nil
}
//jboolean(*CallStaticBooleanMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//func GNI_CallStaticBooleanMethod(env *C.JNIEnv, clazz C.jclass, methodID ...C.jmethodID) C.jboolean {
//  fatal("Not implemented")
//  return nil
//}
//jboolean(*CallStaticBooleanMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//export GNI_CallStaticBooleanMethodV
func GNI_CallStaticBooleanMethodV(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID,  args *C.va_list) C.jboolean {
	fatal("Not implemented")
	return C.jboolean(0)
}
//jboolean(*CallStaticBooleanMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
//export GNI_CallStaticBooleanMethodA
func GNI_CallStaticBooleanMethodA(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jboolean {
	fatal("Not implemented")
	return C.jboolean(0)
}
//jbyte(*CallStaticByteMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//func GNI_CallStaticByteMethod(env *C.JNIEnv, clazz C.jclass, methodID ...C.jmethodID) C.jbyte {
//  fatal("Not implemented")
//  return nil
//}
//jbyte(*CallStaticByteMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//export GNI_CallStaticByteMethodV
func GNI_CallStaticByteMethodV(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID,  args *C.va_list) C.jbyte {
	fatal("Not implemented")
	return C.jbyte(0)
}
//jbyte(*CallStaticByteMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
//export GNI_CallStaticByteMethodA
func GNI_CallStaticByteMethodA(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jbyte {
	fatal("Not implemented")
	return C.jbyte(0)
}
//jchar(*CallStaticCharMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//func GNI_CallStaticCharMethod(env *C.JNIEnv, clazz C.jclass, methodID ...C.jmethodID) C.jchar {
//  fatal("Not implemented")
//  return nil
//}
//jchar(*CallStaticCharMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//export GNI_CallStaticCharMethodV
func GNI_CallStaticCharMethodV(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID,  args *C.va_list) C.jchar {
	fatal("Not implemented")
	return C.jchar(0)
}
//jchar(*CallStaticCharMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
//export GNI_CallStaticCharMethodA
func GNI_CallStaticCharMethodA(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jchar {
	fatal("Not implemented")
	return C.jchar(0)
}
//jshort(*CallStaticShortMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//func GNI_CallStaticShortMethod(env *C.JNIEnv, clazz C.jclass, methodID ...C.jmethodID) C.jshort {
//	fatal("Not implemented")
//	return C.jshort(0)
//}
//jshort(*CallStaticShortMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//func GNI_CallStaticShortMethodV(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID,  args *C.va_list) C.jshort {
//  fatal("Not implemented")
//  return nil
//}
//jshort(*CallStaticShortMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
//export GNI_CallStaticShortMethodA
func GNI_CallStaticShortMethodA(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jshort {
	fatal("Not implemented")
	return C.jshort(0)
}
//jint(*CallStaticIntMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//func GNI_CallStaticIntMethod(env *C.JNIEnv, clazz C.jclass, methodID ...C.jmethodID) C.jint {
//  fatal("Not implemented")
//  return nil
//}
//jint(*CallStaticIntMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//export GNI_CallStaticIntMethodV
func GNI_CallStaticIntMethodV(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID,  args *C.va_list) C.jint {
	fatal("Not implemented")
	return C.jint(0)
}
//jint(*CallStaticIntMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
//export GNI_CallStaticIntMethodA
func GNI_CallStaticIntMethodA(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jint {
	fatal("Not implemented")
	return C.jint(0)
}
//jlong(*CallStaticLongMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//func GNI_CallStaticLongMethod(env *C.JNIEnv, clazz C.jclass, methodID ...C.jmethodID) C.jlong {
//  fatal("Not implemented")
//  return nil
//}
//jlong(*CallStaticLongMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//export GNI_CallStaticLongMethodV
func GNI_CallStaticLongMethodV(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID,  args *C.va_list) C.jlong {
	fatal("Not implemented")
	return C.jlong(0)
}
//jlong(*CallStaticLongMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
//export GNI_CallStaticLongMethodA
func GNI_CallStaticLongMethodA(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jlong {
	fatal("Not implemented")
	return C.jlong(0)
}
//jfloat(*CallStaticFloatMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//func GNI_CallStaticFloatMethod(env *C.JNIEnv, clazz C.jclass, methodID ...C.jmethodID) C.jfloat {
//  fatal("Not implemented")
//  return nil
//}
//jfloat(*CallStaticFloatMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//export GNI_CallStaticFloatMethodV
func GNI_CallStaticFloatMethodV(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID,  args *C.va_list) C.jfloat {
	fatal("Not implemented")
	return C.jfloat(0)
}
//jfloat(*CallStaticFloatMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
//export GNI_CallStaticFloatMethodA
func GNI_CallStaticFloatMethodA(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jfloat {
	fatal("Not implemented")
	return C.jfloat(0)
}
//jdouble(*CallStaticDoubleMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//func GNI_CallStaticDoubleMethod(env *C.JNIEnv, clazz C.jclass, methodID ...C.jmethodID) C.jdouble {
//  fatal("Not implemented")
//  return nil
//}
//jdouble(*CallStaticDoubleMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//export GNI_CallStaticDoubleMethodV
func GNI_CallStaticDoubleMethodV(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID,  args *C.va_list) C.jdouble {
	fatal("Not implemented")
	return C.jdouble(0)
}
//jdouble(*CallStaticDoubleMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
//export GNI_CallStaticDoubleMethodA
func GNI_CallStaticDoubleMethodA(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jdouble {
	fatal("Not implemented")
	return C.jdouble(0)
}
//void(*CallStaticVoidMethod)(JNIEnv *env, jclass cls, jmethodID methodID, ...);
//func GNI_CallStaticVoidMethod(env *C.JNIEnv, cls C.jclass, methodID ...C.jmethodID) {
//  fatal("Not implemented")
//}
//void(*CallStaticVoidMethodV)(JNIEnv *env, jclass cls, jmethodID methodID, va_list args);
//export GNI_CallStaticVoidMethodV
func GNI_CallStaticVoidMethodV(env *C.JNIEnv, cls C.jclass, methodID C.jmethodID,  args *C.va_list) {
	fatal("Not implemented")
}
//void(*CallStaticVoidMethodA)(JNIEnv *env, jclass cls, jmethodID methodID, const jvalue * args);
//export GNI_CallStaticVoidMethodA
func GNI_CallStaticVoidMethodA(env *C.JNIEnv, cls C.jclass, methodID C.jmethodID, args *C.jvalue) {
	fatal("Not implemented")
}
//jfieldID(*GetStaticFieldID)(JNIEnv *env, jclass clazz, const char *name, const char *sig);
//export GNI_GetStaticFieldID
func GNI_GetStaticFieldID(env *C.JNIEnv, clazz C.jclass, name *C.char, sig *C.char) C.jfieldID {
	fatal("Not implemented")
	return nil
}
//jobject(*GetStaticObjectField)(JNIEnv *env, jclass clazz, jfieldID fieldID);
//export GNI_GetStaticObjectField
func GNI_GetStaticObjectField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID) C.jobject {
	fatal("Not implemented")
	return nil
}
//jboolean(*GetStaticBooleanField)(JNIEnv *env, jclass clazz, jfieldID fieldID);
//export GNI_GetStaticBooleanField
func GNI_GetStaticBooleanField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID) C.jboolean {
	fatal("Not implemented")
	return C.jboolean(0)
}
//jbyte(*GetStaticByteField)(JNIEnv *env, jclass clazz, jfieldID fieldID);
//export GNI_GetStaticByteField
func GNI_GetStaticByteField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID) C.jbyte {
	fatal("Not implemented")
	return C.jbyte(0)
}
//jchar(*GetStaticCharField)(JNIEnv *env, jclass clazz, jfieldID fieldID);
//export GNI_GetStaticCharField
func GNI_GetStaticCharField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID) C.jchar {
	fatal("Not implemented")
	return C.jchar(0)
}
//jshort(*GetStaticShortField)(JNIEnv *env, jclass clazz, jfieldID fieldID);
//export GNI_GetStaticShortField
func GNI_GetStaticShortField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID) C.jshort {
	fatal("Not implemented")
	return C.jshort(0)
}
//jint(*GetStaticIntField)(JNIEnv *env, jclass clazz, jfieldID fieldID);
//export GNI_GetStaticIntField
func GNI_GetStaticIntField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID) C.jint {
	fatal("Not implemented")
	return C.jint(0)
}
//jlong(*GetStaticLongField)(JNIEnv *env, jclass clazz, jfieldID fieldID);
//export GNI_GetStaticLongField
func GNI_GetStaticLongField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID) C.jlong {
	fatal("Not implemented")
	return C.jlong(0)
}
//jfloat(*GetStaticFloatField)(JNIEnv *env, jclass clazz, jfieldID fieldID);
//export GNI_GetStaticFloatField
func GNI_GetStaticFloatField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID) C.jfloat {
	fatal("Not implemented")
	return C.jfloat(0)
}
//jdouble(*GetStaticDoubleField)(JNIEnv *env, jclass clazz, jfieldID fieldID);
//export GNI_GetStaticDoubleField
func GNI_GetStaticDoubleField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID) C.jdouble {
	fatal("Not implemented")
	return C.jdouble(0)
}
//void(*SetStaticObjectField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jobject value);
//export GNI_SetStaticObjectField
func GNI_SetStaticObjectField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID, value C.jobject) {
	fatal("Not implemented")
}
//void(*SetStaticBooleanField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jboolean value);
//export GNI_SetStaticBooleanField
func GNI_SetStaticBooleanField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID, value C.jboolean) {
	fatal("Not implemented")
}
//void(*SetStaticByteField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jbyte value);
//export GNI_SetStaticByteField
func GNI_SetStaticByteField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID, value C.jbyte) {
	fatal("Not implemented")
}
//void(*SetStaticCharField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jchar value);
//export GNI_SetStaticCharField
func GNI_SetStaticCharField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID, value C.jchar) {
	fatal("Not implemented")
}
//void(*SetStaticShortField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jshort value);
//export GNI_SetStaticShortField
func GNI_SetStaticShortField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID, value C.jshort) {
	fatal("Not implemented")
}
//void(*SetStaticIntField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jint value);
//export GNI_SetStaticIntField
func GNI_SetStaticIntField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID, value C.jint) {
	fatal("Not implemented")
}
//void(*SetStaticLongField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jlong value);
//export GNI_SetStaticLongField
func GNI_SetStaticLongField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID, value C.jlong) {
	fatal("Not implemented")
}
//void(*SetStaticFloatField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jfloat value);
//export GNI_SetStaticFloatField
func GNI_SetStaticFloatField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID, value C.jfloat) {
	fatal("Not implemented")
}
//void(*SetStaticDoubleField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jdouble value);
//export GNI_SetStaticDoubleField
func GNI_SetStaticDoubleField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID, value C.jdouble) {
	fatal("Not implemented")
}
//jstring(*NewString)(JNIEnv *env, const jchar *unicode, jsize len);
//export GNI_NewString
func GNI_NewString(env *C.JNIEnv, unicode *C.jchar, len C.jsize) C.jstring {
	fatal("Not implemented")
	return nil
}
//jsize(*GetStringLength)(JNIEnv *env, jstring str);
//export GNI_GetStringLength
func GNI_GetStringLength(env *C.JNIEnv, str C.jstring) C.jsize {
	fatal("Not implemented")
	return C.jsize(0)
}
//const jchar *(*GetStringChars)(JNIEnv *env, jstring str, jboolean *isCopy);
//export GNI_GetStringChars
func GNI_GetStringChars(env *C.JNIEnv, str C.jstring, isCopy *C.jboolean) *C.jchar {
	fatal("Not implemented")
	return nil
}
//void(*ReleaseStringChars)(JNIEnv *env, jstring str, const jchar *chars);
//export GNI_ReleaseStringChars
func GNI_ReleaseStringChars(env *C.JNIEnv, str C.jstring, chars C.jchar) {
	fatal("Not implemented")
}
//jstring(*NewStringUTF)(JNIEnv *env, const char *utf);
//export GNI_NewStringUTF
func GNI_NewStringUTF(env *C.JNIEnv, utf *C.char) C.jstring {
	fatal("Not implemented")
	return nil
}
//jsize(*GetStringUTFLength)(JNIEnv *env, jstring str);
//export GNI_GetStringUTFLength
func GNI_GetStringUTFLength(env *C.JNIEnv, str C.jstring) C.jsize {
	fatal("Not implemented")
	return C.jsize(0)
}
//const char*(*GetStringUTFChars)(JNIEnv *env, jstring str, jboolean *isCopy);
//export GNI_GetStringUTFChars
func GNI_GetStringUTFChars(env *C.JNIEnv, str C.jstring, isCopy *C.jboolean) *C.char {
	fatal("Not implemented")
	return nil
}
//void(*ReleaseStringUTFChars)(JNIEnv *env, jstring str, const char* chars);
//export GNI_ReleaseStringUTFChars
func GNI_ReleaseStringUTFChars(env *C.JNIEnv, str C.jstring, chars *C.char) {
	fatal("Not implemented")
}
//jsize(*GetArrayLength)(JNIEnv *env, jarray array);
//export GNI_GetArrayLength
func GNI_GetArrayLength(env *C.JNIEnv, array *C.jarray) C.jsize {
	fatal("Not implemented")
	return C.jsize(0)
}
//jobjectArray(*NewObjectArray)(JNIEnv *env, jsize len, jclass clazz, jobject init);
//export GNI_NewObjectArray
func GNI_NewObjectArray(env *C.JNIEnv, len C.jsize, clazz C.jclass, init C.jobject) C.jobjectArray {
	fatal("Not implemented")
	return nil
}
//jobject(*GetObjectArrayElement)(JNIEnv *env, jobjectArray array, jsize index);
//export GNI_GetObjectArrayElement
func GNI_GetObjectArrayElement(env *C.JNIEnv, array C.jobjectArray, index C.jsize) C.jobject {
	fatal("Not implemented")
	return nil
}
//void(*SetObjectArrayElement)(JNIEnv *env, jobjectArray array, jsize index, jobject val);
//export GNI_SetObjectArrayElement
func GNI_SetObjectArrayElement(env *C.JNIEnv, array C.jobjectArray, index C.jsize, val C.jobject) {
	fatal("Not implemented")
}
//jbooleanArray(*NewBooleanArray)(JNIEnv *env, jsize len);
//export GNI_NewBooleanArray
func GNI_NewBooleanArray(env *C.JNIEnv, len C.jsize) C.jbooleanArray {
	fatal("Not implemented")
	return nil
}
//jbyteArray(*NewByteArray)(JNIEnv *env, jsize len);
//export GNI_NewByteArray
func GNI_NewByteArray(env *C.JNIEnv, len C.jsize) C.jbyteArray {
	fatal("Not implemented")
	return nil
}
//jcharArray(*NewCharArray)(JNIEnv *env, jsize len);
//export GNI_NewCharArray
func GNI_NewCharArray(env *C.JNIEnv, len C.jsize) C.jcharArray {
	fatal("Not implemented")
	return nil
}
//jshortArray(*NewShortArray)(JNIEnv *env, jsize len);
//export GNI_NewShortArray
func GNI_NewShortArray(env *C.JNIEnv, len C.jsize) C.jshortArray {
	fatal("Not implemented")
	return nil
}
//jintArray(*NewIntArray)(JNIEnv *env, jsize len);
//export GNI_NewIntArray
func GNI_NewIntArray(env *C.JNIEnv, len C.jsize) C.jintArray {
	fatal("Not implemented")
	return nil
}
//jlongArray(*NewLongArray)(JNIEnv *env, jsize len);
//export GNI_NewLongArray
func GNI_NewLongArray(env *C.JNIEnv, len C.jsize) C.jlongArray {
	fatal("Not implemented")
	return nil
}
//jfloatArray(*NewFloatArray)(JNIEnv *env, jsize len);
//export GNI_NewFloatArray
func GNI_NewFloatArray(env *C.JNIEnv, len C.jsize) C.jfloatArray {
	fatal("Not implemented")
	return nil
}
//jdoubleArray(*NewDoubleArray)(JNIEnv *env, jsize len);
//export GNI_NewDoubleArray
func GNI_NewDoubleArray(env *C.JNIEnv, len C.jsize) C.jdoubleArray {
	fatal("Not implemented")
	return nil
}
//jboolean *(*GetBooleanArrayElements)(JNIEnv *env, jbooleanArray array, jboolean *isCopy);
//export GNI_GetBooleanArrayElements
func GNI_GetBooleanArrayElements(env *C.JNIEnv, array C.jbooleanArray, isCopy *C.jboolean) C.jboolean {
	fatal("Not implemented")
	return C.jboolean(0)
}
//jbyte *(*GetByteArrayElements)(JNIEnv *env, jbyteArray array, jboolean *isCopy);
//export GNI_GetByteArrayElements
func GNI_GetByteArrayElements(env *C.JNIEnv, array C.jbyteArray, isCopy *C.jboolean) C.jbyte {
	fatal("Not implemented")
	return C.jbyte(0)
}
//jchar *(*GetCharArrayElements)(JNIEnv *env, jcharArray array, jboolean *isCopy);
//export GNI_GetCharArrayElements
func GNI_GetCharArrayElements(env *C.JNIEnv, array C.jcharArray, isCopy *C.jboolean) C.jchar {
	fatal("Not implemented")
	return C.jchar(0)
}
//jshort *(*GetShortArrayElements)(JNIEnv *env, jshortArray array, jboolean *isCopy);
//export GNI_GetShortArrayElements
func GNI_GetShortArrayElements(env *C.JNIEnv, array C.jshortArray, isCopy *C.jboolean) C.jshort {
	fatal("Not implemented")
	return C.jshort(0)
}
//jint *(*GetIntArrayElements)(JNIEnv *env, jintArray array, jboolean *isCopy);
//export GNI_GetIntArrayElements
func GNI_GetIntArrayElements(env *C.JNIEnv, array C.jintArray, isCopy *C.jboolean) C.jint {
	fatal("Not implemented")
	return C.jint(0)
}
//jlong *(*GetLongArrayElements)(JNIEnv *env, jlongArray array, jboolean *isCopy);
//export GNI_GetLongArrayElements
func GNI_GetLongArrayElements(env *C.JNIEnv, array C.jlongArray, isCopy *C.jboolean) C.jlong {
	fatal("Not implemented")
	return C.jlong(0)
}
//jfloat *(*GetFloatArrayElements)(JNIEnv *env, jfloatArray array, jboolean *isCopy);
//export GNI_GetFloatArrayElements
func GNI_GetFloatArrayElements(env *C.JNIEnv, array C.jfloatArray, isCopy *C.jboolean) C.jfloat {
	fatal("Not implemented")
	return C.jfloat(0)
}
//jdouble *(*GetDoubleArrayElements)(JNIEnv *env, jdoubleArray array, jboolean *isCopy);
//export GNI_GetDoubleArrayElements
func GNI_GetDoubleArrayElements(env *C.JNIEnv, array C.jdoubleArray, isCopy *C.jboolean) C.jdouble {
	fatal("Not implemented")
	return C.jdouble(0)
}
//void(*ReleaseBooleanArrayElements)(JNIEnv *env, jbooleanArray array, jboolean *elems, jint mode);
//export GNI_ReleaseBooleanArrayElements
func GNI_ReleaseBooleanArrayElements(env *C.JNIEnv,array C.jbooleanArray, elems *C.jboolean, mode C.jint) {
	fatal("Not implemented")
}
//void(*ReleaseByteArrayElements)(JNIEnv *env, jbyteArray array, jbyte *elems, jint mode);
//export GNI_ReleaseByteArrayElements
func GNI_ReleaseByteArrayElements(env *C.JNIEnv,array C.jbyteArray, elems *C.jbyte, mode C.jint) {
	fatal("Not implemented")
}
//void(*ReleaseCharArrayElements)(JNIEnv *env, jcharArray array, jchar *elems, jint mode);
//export GNI_ReleaseCharArrayElements
func GNI_ReleaseCharArrayElements(env *C.JNIEnv,array C.jcharArray, elems *C.jchar, mode C.jint) {
	fatal("Not implemented")
}
//void(*ReleaseShortArrayElements)(JNIEnv *env, jshortArray array, jshort *elems, jint mode);
//export GNI_ReleaseShortArrayElements
func GNI_ReleaseShortArrayElements(env *C.JNIEnv,array C.jshortArray, elems *C.jshort, mode C.jint) {
	fatal("Not implemented")
}
//void(*ReleaseIntArrayElements)(JNIEnv *env, jintArray array, jint *elems, jint mode);
//export GNI_ReleaseIntArrayElements
func GNI_ReleaseIntArrayElements(env *C.JNIEnv,array C.jintArray, elems *C.jint, mode C.jint) {
	fatal("Not implemented")
}
//void(*ReleaseLongArrayElements)(JNIEnv *env, jlongArray array, jlong *elems, jint mode);
//export GNI_ReleaseLongArrayElements
func GNI_ReleaseLongArrayElements(env *C.JNIEnv,array C.jlongArray, elems *C.jlong, mode C.jint) {
	fatal("Not implemented")
}
//void(*ReleaseFloatArrayElements)(JNIEnv *env, jfloatArray array, jfloat *elems, jint mode);
//export GNI_ReleaseFloatArrayElements
func GNI_ReleaseFloatArrayElements(env *C.JNIEnv,array C.jfloatArray, elems *C.jfloat, mode C.jint) {
	fatal("Not implemented")
}
//void(*ReleaseDoubleArrayElements)(JNIEnv *env, jdoubleArray array, jdouble *elems, jint mode);
//export GNI_ReleaseDoubleArrayElements
func GNI_ReleaseDoubleArrayElements(env *C.JNIEnv,array C.jdoubleArray, elems *C.jdouble, mode C.jint) {
	fatal("Not implemented")
}
//void(*GetBooleanArrayRegion)(JNIEnv *env, jbooleanArray array, jsize start, jsize l, jboolean *buf);
//export GNI_GetBooleanArrayRegion
func GNI_GetBooleanArrayRegion(env *C.JNIEnv, array C.jbooleanArray,start C.jsize, l C.jsize, buf *C.jboolean) {
	fatal("Not implemented")
}
//void(*GetByteArrayRegion)(JNIEnv *env, jbyteArray array, jsize start, jsize len, jbyte *buf);
//export GNI_GetByteArrayRegion
func GNI_GetByteArrayRegion(env *C.JNIEnv,array C.jbyteArray, start C.jsize, len C.jsize, buf *C.jbyte) {
	fatal("Not implemented")
}
//void(*GetCharArrayRegion)(JNIEnv *env, jcharArray array, jsize start, jsize len, jchar *buf);
//export GNI_GetCharArrayRegion
func GNI_GetCharArrayRegion(env *C.JNIEnv,array C.jcharArray, start C.jsize, len C.jsize, buf *C.jchar) {
	fatal("Not implemented")
}
//void(*GetShortArrayRegion)(JNIEnv *env, jshortArray array, jsize start, jsize len, jshort *buf);
//export GNI_GetShortArrayRegion
func GNI_GetShortArrayRegion(env *C.JNIEnv,array C.jshortArray, start C.jsize, len C.jsize, buf *C.jshort) {
	fatal("Not implemented")
}
//void(*GetIntArrayRegion)(JNIEnv *env, jintArray array, jsize start, jsize len, jint *buf);
//export GNI_GetIntArrayRegion
func GNI_GetIntArrayRegion(env *C.JNIEnv,array C.jintArray, start C.jsize, len C.jsize, buf *C.jint) {
	fatal("Not implemented")
}
//void(*GetLongArrayRegion)(JNIEnv *env, jlongArray array, jsize start, jsize len, jlong *buf);
//export GNI_GetLongArrayRegion
func GNI_GetLongArrayRegion(env *C.JNIEnv,array C.jlongArray, start C.jsize, len C.jsize, buf *C.jlong) {
	fatal("Not implemented")
}
//void(*GetFloatArrayRegion)(JNIEnv *env, jfloatArray array, jsize start, jsize len, jfloat *buf);
//export GNI_GetFloatArrayRegion
func GNI_GetFloatArrayRegion(env *C.JNIEnv,array C.jfloatArray, start C.jsize, len C.jsize, buf *C.jfloat) {
	fatal("Not implemented")
}
//void(*GetDoubleArrayRegion)(JNIEnv *env, jdoubleArray array, jsize start, jsize len, jdouble *buf);
//export GNI_GetDoubleArrayRegion
func GNI_GetDoubleArrayRegion(env *C.JNIEnv,array C.jdoubleArray, start C.jsize, len C.jsize, buf *C.jdouble) {
	fatal("Not implemented")
}
//void(*SetBooleanArrayRegion)(JNIEnv *env, jbooleanArray array, jsize start, jsize l, const jboolean *buf);
//export GNI_SetBooleanArrayRegion
func GNI_SetBooleanArrayRegion(env *C.JNIEnv, array C.jbooleanArray,start C.jsize, l C.jsize, buf *C.jboolean) {
	fatal("Not implemented")
}
//void(*SetByteArrayRegion)(JNIEnv *env, jbyteArray array, jsize start, jsize len, const jbyte *buf);
//export GNI_SetByteArrayRegion
func GNI_SetByteArrayRegion(env *C.JNIEnv,array C.jbyteArray, start C.jsize, len C.jsize, buf *C.jbyte) {
	fatal("Not implemented")
}
//void(*SetCharArrayRegion)(JNIEnv *env, jcharArray array, jsize start, jsize len, const jchar *buf);
//export GNI_SetCharArrayRegion
func GNI_SetCharArrayRegion(env *C.JNIEnv,array C.jcharArray, start C.jsize, len C.jsize, buf *C.jchar) {
	fatal("Not implemented")
}
//void(*SetShortArrayRegion)(JNIEnv *env, jshortArray array, jsize start, jsize len, const jshort *buf);
//export GNI_SetShortArrayRegion
func GNI_SetShortArrayRegion(env *C.JNIEnv,array C.jshortArray, start C.jsize, len C.jsize, buf *C.jshort) {
	fatal("Not implemented")
}
//void(*SetIntArrayRegion)(JNIEnv *env, jintArray array, jsize start, jsize len, const jint *buf);
//export GNI_SetIntArrayRegion
func GNI_SetIntArrayRegion(env *C.JNIEnv,array C.jintArray, start C.jsize, len C.jsize, buf *C.jint) {
	fatal("Not implemented")
}
//void(*SetLongArrayRegion)(JNIEnv *env, jlongArray array, jsize start, jsize len, const jlong *buf);
//export GNI_SetLongArrayRegion
func GNI_SetLongArrayRegion(env *C.JNIEnv,array C.jlongArray, start C.jsize, len C.jsize, buf *C.jlong) {
	fatal("Not implemented")
}
//void(*SetFloatArrayRegion)(JNIEnv *env, jfloatArray array, jsize start, jsize len, const jfloat *buf);
//export GNI_SetFloatArrayRegion
func GNI_SetFloatArrayRegion(env *C.JNIEnv,array C.jfloatArray, start C.jsize, len C.jsize, buf *C.jfloat) {
	fatal("Not implemented")
}
//void(*SetDoubleArrayRegion)(JNIEnv *env, jdoubleArray array, jsize start, jsize len, const jdouble *buf);
//export GNI_SetDoubleArrayRegion
func GNI_SetDoubleArrayRegion(env *C.JNIEnv,array C.jdoubleArray, start C.jsize, len C.jsize, buf *C.jdouble) {
	fatal("Not implemented")
}
//jint(*RegisterNatives)(JNIEnv *env, jclass clazz, const JNINativeMethod *methods, jint nMethods);
//export GNI_RegisterNatives
func GNI_RegisterNatives(env *C.JNIEnv, clazz C.jclass, methods *C.JNINativeMethod, nMethods C.jint) C.jint {
	fatal("Not implemented")
	return C.jint(0)
}
//jint(*UnregisterNatives)(JNIEnv *env, jclass clazz);
//export GNI_UnregisterNatives
func GNI_UnregisterNatives(env *C.JNIEnv, clazz C.jclass) C.jint {
	fatal("Not implemented")
	return C.jint(0)
}
//jint(*MonitorEnter)(JNIEnv *env, jobject obj);
//export GNI_MonitorEnter
func GNI_MonitorEnter(env *C.JNIEnv, obj C.jobject) C.jint {
	fatal("Not implemented")
	return C.jint(0)
}
//jint(*MonitorExit)(JNIEnv *env, jobject obj);
//export GNI_MonitorExit
func GNI_MonitorExit(env *C.JNIEnv, obj C.jobject) C.jint {
	fatal("Not implemented")
	return C.jint(0)
}
//jint(*GetJavaVM)(JNIEnv *env, JavaVM **vm);
//export GNI_GetJavaVM
func GNI_GetJavaVM(env *C.JNIEnv, vm unsafe.Pointer) C.jint {
	fatal("Not implemented")
	return C.jint(0)
}
//void(*GetStringRegion)(JNIEnv *env, jstring str, jsize start, jsize len, jchar *buf);
//export GNI_GetStringRegion
func GNI_GetStringRegion(env *C.JNIEnv,str C.jstring, start C.jsize, len C.jsize, buf *C.jchar) {
	fatal("Not implemented")
}
//void(*GetStringUTFRegion)(JNIEnv *env, jstring str, jsize start, jsize len, char *buf);
//export GNI_GetStringUTFRegion
func GNI_GetStringUTFRegion(env *C.JNIEnv,str C.jstring, start C.jsize, len C.jsize, buf *C.char) {
	fatal("Not implemented")
}
//void *(*GetPrimitiveArrayCritical)(JNIEnv *env, jarray array, jboolean *isCopy);
//export GNI_GetPrimitiveArrayCritical
func GNI_GetPrimitiveArrayCritical(env *C.JNIEnv, array C.jarray, isCopy *C.jboolean) {
	fatal("Not implemented")
}
//void(*ReleasePrimitiveArrayCritical)(JNIEnv *env, jarray array, void *carray, jint mode);
//export GNI_ReleasePrimitiveArrayCritical
func GNI_ReleasePrimitiveArrayCritical(env *C.JNIEnv,array C.jarray, carray unsafe.Pointer, mode C.jint) {
	fatal("Not implemented")
}
//const jchar *(*GetStringCritical)(JNIEnv *env, jstring string, jboolean *isCopy);
//export GNI_GetStringCritical
func GNI_GetStringCritical(env *C.JNIEnv, string C.jstring, isCopy *C.jboolean) *C.jchar {
	fatal("Not implemented")
	return nil
}
//void(*ReleaseStringCritical)(JNIEnv *env, jstring string, const jchar *cstring);
//export GNI_ReleaseStringCritical
func GNI_ReleaseStringCritical(env *C.JNIEnv, string C.jstring, cstring *C.jchar) {
	fatal("Not implemented")
}
//jweak(*NewWeakGlobalRef)(JNIEnv *env, jobject obj);
//export GNI_NewWeakGlobalRef
func GNI_NewWeakGlobalRef(env *C.JNIEnv, obj C.jobject) C.jweak {
	fatal("Not implemented")
	return nil
}
//void(*DeleteWeakGlobalRef)(JNIEnv *env, jweak ref);
//export GNI_DeleteWeakGlobalRef
func GNI_DeleteWeakGlobalRef(env *C.JNIEnv, ref C.jweak) {
	fatal("Not implemented")
}
//jboolean(*ExceptionCheck)(JNIEnv *env);
//export GNI_ExceptionCheck
func GNI_ExceptionCheck(env *C.JNIEnv) C.jboolean {
	fatal("Not implemented")
	return C.jboolean(0)
}
//jobject(*NewDirectByteBuffer)(JNIEnv* env, void* address, jlong capacity);
//export GNI_NewDirectByteBuffer
func GNI_NewDirectByteBuffer(env *C.JNIEnv, address unsafe.Pointer, capacity C.jlong) C.jobject {
	fatal("Not implemented")
	return nil
}
//void*(*GetDirectBufferAddress)(JNIEnv* env, jobject buf);
//export GNI_GetDirectBufferAddress
func GNI_GetDirectBufferAddress(env *C.JNIEnv, buf C.jobject) unsafe.Pointer {
	fatal("Not implemented")
	return nil
}
//jlong(*GetDirectBufferCapacity)(JNIEnv* env, jobject buf);
//export GNI_GetDirectBufferCapacity
func GNI_GetDirectBufferCapacity(env *C.JNIEnv, buf C.jobject) C.jlong {
	fatal("Not implemented")
	return C.jlong(0)
}
/* New JNI 1.6 Features */
//jobjectRefType(*GetObjectRefType)(JNIEnv* env, jobject obj);
//export GNI_GetObjectRefType
func GNI_GetObjectRefType(env *C.JNIEnv, obj C.jobject) C.jobjectRefType {
	fatal("Not implemented")
	return C.jobjectRefType(0)
}