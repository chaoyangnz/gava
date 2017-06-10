/* Created by "go tool cgo" - DO NOT EDIT. */

/* package gvm */

/* Start of preamble from import "C" comments.  */


#line 7 "/Users/Charvis/GoglandProjects/gvm/src/gvm/gni.go"


#include <jni.h>

#line 1 "cgo-generated-wrapper"


/* End of preamble from import "C" comments.  */


/* Start of boilerplate cgo prologue.  */
#line 1 "cgo-gcc-export-header-prolog"

#ifndef GO_CGO_PROLOGUE_H
#define GO_CGO_PROLOGUE_H

typedef signed char GoInt8;
typedef unsigned char GoUint8;
typedef short GoInt16;
typedef unsigned short GoUint16;
typedef int GoInt32;
typedef unsigned int GoUint32;
typedef long long GoInt64;
typedef unsigned long long GoUint64;
typedef GoInt64 GoInt;
typedef GoUint64 GoUint;
typedef __SIZE_TYPE__ GoUintptr;
typedef float GoFloat32;
typedef double GoFloat64;
typedef float _Complex GoComplex64;
typedef double _Complex GoComplex128;

/*
  static assertion to make sure the file is being used on architecture
  at least with matching size of GoInt.
*/
typedef char _check_for_64_bit_pointer_matching_GoInt[sizeof(void*)==64/8 ? 1:-1];

typedef struct { const char *p; GoInt n; } GoString;
typedef void *GoMap;
typedef void *GoChan;
typedef struct { void *t; void *v; } GoInterface;
typedef struct { void *data; GoInt len; GoInt cap; } GoSlice;

#endif

/* End of boilerplate cgo prologue.  */

#ifdef __cplusplus
extern "C" {
#endif


//jint(*GetVersion)(JNIEnv *env);

extern jint GNI_GetVersion(JNIEnv* p0);

//jclass(*DefineClass)(JNIEnv *env, const char *name, jobject loader, const jbyte *buf,jsize len);

extern jclass GNI_DefineClass(JNIEnv* p0, char* p1, jobject p2, jbyte* p3, jsize p4);

//jclass(*FindClass)(JNIEnv *env, const char *name);

extern jclass GNI_FindClass(JNIEnv* p0, char* p1);

//jmethodID(*FromReflectedMethod)(JNIEnv *env, jobject method);

extern jmethodID GNI_FromReflectedMethod(JNIEnv* p0, jobject p1);

//jfieldID(*FromReflectedField)(JNIEnv *env, jobject field);

extern jfieldID GNI_FromReflectedField(JNIEnv* p0, jobject p1);

//jobject(*ToReflectedMethod)(JNIEnv *env, jclass cls, jmethodID methodID, jboolean isStatic);

extern jobject GNI_ToReflectedMethod(JNIEnv* p0, jclass p1, jmethodID p2, jboolean p3);

//jclass(*GetSuperclass)(JNIEnv *env, jclass sub);

extern jclass GNI_GetSuperclass(JNIEnv* p0, jclass p1);

//jboolean(*IsAssignableFrom)(JNIEnv *env, jclass sub, jclass sup);

extern jboolean GNI_IsAssignableFrom(JNIEnv* p0, jclass p1, jclass p2);

//jobject(*ToReflectedField)(JNIEnv *env, jclass cls, jfieldID fieldID, jboolean isStatic);

extern jobject GNI_ToReflectedField(JNIEnv* p0, jclass p1, jfieldID p2, jboolean p3);

//jint(*Throw)(JNIEnv *env, jthrowable obj);

extern jint GNI_Throw(JNIEnv* p0, jthrowable p1);

//jint(*ThrowNew)(JNIEnv *env, jclass clazz, const char *msg);

extern jint GNI_ThrowNew(JNIEnv* p0, jclass p1, char* p2);

//jthrowable(*ExceptionOccurred)(JNIEnv *env);

extern jthrowable GNI_ExceptionOccurred(JNIEnv* p0);

//void(*ExceptionDescribe)(JNIEnv *env);

extern void GNI_ExceptionDescribe(JNIEnv* p0);

//void(*ExceptionClear)(JNIEnv *env);

extern void GNI_ExceptionClear(JNIEnv* p0);

//void(*FatalError)(JNIEnv *env, const char *msg);

extern void GNI_FatalError(JNIEnv* p0, char* p1);

//jint(*PushLocalFrame)(JNIEnv *env, jint capacity);

extern jint GNI_PushLocalFrame(JNIEnv* p0, jint p1);

//jobject(*PopLocalFrame)(JNIEnv *env, jobject result);

extern jobject GNI_PopLocalFrame(JNIEnv* p0, jobject p1);

//jobject(*NewGlobalRef)(JNIEnv *env, jobject lobj);

extern jobject GNI_NewGlobalRef(JNIEnv* p0, jobject p1);

//void(*DeleteGlobalRef)(JNIEnv *env, jobject gref);

extern void GNI_DeleteGlobalRef(JNIEnv* p0, jobject p1);

//void(*DeleteLocalRef)(JNIEnv *env, jobject obj);

extern void GNI_DeleteLocalRef(JNIEnv* p0, jobject p1);

//jboolean(*IsSameObject)(JNIEnv *env, jobject obj1, jobject obj2);

extern jboolean GNI_IsSameObject(JNIEnv* p0, jobject p1, jobject p2);

//jobject(*NewLocalRef)(JNIEnv *env, jobject ref);

extern jobject GNI_NewLocalRef(JNIEnv* p0, jobject p1);

//jint(*EnsureLocalCapacity)(JNIEnv *env, jint capacity);

extern jint GNI_EnsureLocalCapacity(JNIEnv* p0, jint p1);

//jobject(*AllocObject)(JNIEnv *env, jclass clazz);

extern jobject GNI_AllocObject(JNIEnv* p0, jclass p1);

//jobject(*NewObject)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//func GNI_NewObject(env *C.JNIEnv, clazz C.jclass, methodID ...C.jmethodID) C.jobject {
//  fatal("Not implemented")
//  return nil
//}
//jobject(*NewObjectV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);

extern jobject GNI_NewObjectV(JNIEnv* p0, jclass p1, jmethodID p2, va_list* p3);

//jobject(*NewObjectA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

extern jobject GNI_NewObjectA(JNIEnv* p0, jclass p1, jmethodID p2, jvalue* p3);

//jclass(*GetObjectClass)(JNIEnv *env, jobject obj);

extern jclass GNI_GetObjectClass(JNIEnv* p0, jobject p1);

//jboolean(*IsInstanceOf)(JNIEnv *env, jobject obj, jclass clazz);

extern jboolean GNI_IsInstanceOf(JNIEnv* p0, jobject p1, jclass p2);

//jmethodID(*GetMethodID)(JNIEnv *env, jclass clazz, const char *name, const char *sig);

extern jmethodID GNI_GetMethodID(JNIEnv* p0, jclass p1, char* p2, char* p3);

//jobject(*CallObjectMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//func GNI_CallObjectMethod(env *C.JNIEnv, obj C.jobject, methodID ...C.jmethodID) C.jobject {
//  fatal("Not implemented")
//  return nil
//}
//jobject(*CallObjectMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);

extern jobject GNI_CallObjectMethodV(JNIEnv* p0, jobject p1, jmethodID p2, va_list* p3);

//jobject(*CallObjectMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue * args);

extern jobject GNI_CallObjectMethodA(JNIEnv* p0, jobject p1, jmethodID p2, jvalue* p3);

//jboolean(*CallBooleanMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//func GNI_CallBooleanMethod(env *C.JNIEnv, obj C.jobject, methodID ...C.jmethodID) C.jboolean {
//  fatal("Not implemented")
//  return nil
//}
//jboolean(*CallBooleanMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);

extern jboolean GNI_CallBooleanMethodV(JNIEnv* p0, jobject p1, jmethodID p2, va_list* p3);

//jboolean(*CallBooleanMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue * args);

extern jboolean GNI_CallBooleanMethodA(JNIEnv* p0, jobject p1, jmethodID p2, jvalue* p3);

//jbyte(*CallByteMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//func GNI_CallByteMethod(env *C.JNIEnv, obj C.jobject, methodID ...C.jmethodID) C.jbyte {
//  fatal("Not implemented")
//  return nil
//}
//jbyte(*CallByteMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);

extern jbyte GNI_CallByteMethodV(JNIEnv* p0, jobject p1, jmethodID p2, va_list* p3);

//jbyte(*CallByteMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);

extern jbyte GNI_CallByteMethodA(JNIEnv* p0, jobject p1, jmethodID p2, jvalue* p3);

//jchar(*CallCharMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//func GNI_CallCharMethod(env *C.JNIEnv, obj C.jobject, methodID ...C.jmethodID) C.jchar {
//  fatal("Not implemented")
//  return nil
//}
//jchar(*CallCharMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);

extern jchar GNI_CallCharMethodV(JNIEnv* p0, jobject p1, jmethodID p2, va_list* p3);

//jchar(*CallCharMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);

extern jchar GNI_CallCharMethodA(JNIEnv* p0, jobject p1, jmethodID p2, jvalue* p3);

//jshort(*CallShortMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//func GNI_CallShortMethod(env *C.JNIEnv, obj C.jobject, methodID ...C.jmethodID) C.jshort {
//  fatal("Not implemented")
//  return nil
//}
//jshort(*CallShortMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);

extern jshort GNI_CallShortMethodV(JNIEnv* p0, jobject p1, jmethodID p2, va_list* p3);

//jshort(*CallShortMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);

extern jshort GNI_CallShortMethodA(JNIEnv* p0, jobject p1, jmethodID p2, jvalue* p3);

//jint(*CallIntMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//func GNI_CallIntMethod(env *C.JNIEnv, obj C.jobject, methodID ...C.jmethodID) C.jint {
//  fatal("Not implemented")
//  return nil
//}
//jint(*CallIntMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);

extern jint GNI_CallIntMethodV(JNIEnv* p0, jobject p1, jmethodID p2, va_list* p3);

//jint(*CallIntMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);

extern jint GNI_CallIntMethodA(JNIEnv* p0, jobject p1, jmethodID p2, jvalue* p3);

//jlong(*CallLongMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//func GNI_CallLongMethod(env *C.JNIEnv, obj C.jobject, methodID ...C.jmethodID) C.jlong {
//  fatal("Not implemented")
//  return nil
//}
//jlong(*CallLongMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);

extern jlong GNI_CallLongMethodV(JNIEnv* p0, jobject p1, jmethodID p2, va_list* p3);

//jlong(*CallLongMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);

extern jlong GNI_CallLongMethodA(JNIEnv* p0, jobject p1, jmethodID p2, jvalue* p3);

//jfloat(*CallFloatMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//func GNI_CallFloatMethod(env *C.JNIEnv, obj C.jobject, methodID ...C.jmethodID) C.jfloat {
//  fatal("Not implemented")
//  return nil
//}
//jfloat(*CallFloatMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);

extern jfloat GNI_CallFloatMethodV(JNIEnv* p0, jobject p1, jmethodID p2, va_list* p3);

//jfloat(*CallFloatMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);

extern jfloat GNI_CallFloatMethodA(JNIEnv* p0, jobject p1, jmethodID p2, jvalue* p3);

//jdouble(*CallDoubleMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//func GNI_CallDoubleMethod(env *C.JNIEnv, obj C.jobject, methodID ...C.jmethodID) C.jdouble {
//  fatal("Not implemented")
//  return nil
//}
//jdouble(*CallDoubleMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);

extern jdouble GNI_CallDoubleMethodV(JNIEnv* p0, jobject p1, jmethodID p2, va_list* p3);

//jdouble(*CallDoubleMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);

extern jdouble GNI_CallDoubleMethodA(JNIEnv* p0, jobject p1, jmethodID p2, jvalue* p3);

//void(*CallVoidMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//func GNI_CallVoidMethod(env *C.JNIEnv, obj C.jobject, methodID ...C.jmethodID) {
//  fatal("Not implemented")
//}
//void(*CallVoidMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);

extern void GNI_CallVoidMethodV(JNIEnv* p0, jobject p1, jmethodID p2, va_list* p3);

//void(*CallVoidMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue * args);

extern void GNI_CallVoidMethodA(JNIEnv* p0, jobject p1, jmethodID p2, jvalue* p3);

//jobject(*CallNonvirtualObjectMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//func GNI_CallNonvirtualObjectMethod(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID ...C.jmethodID) C.jobject {
//  fatal("Not implemented")
//  return nil
//}
//jobject(*CallNonvirtualObjectMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,va_list args);

extern jobject GNI_CallNonvirtualObjectMethodV(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, va_list* p4);

//jobject(*CallNonvirtualObjectMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,const jvalue * args);

extern jobject GNI_CallNonvirtualObjectMethodA(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, jvalue* p4);

//jboolean(*CallNonvirtualBooleanMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//func GNI_CallNonvirtualBooleanMethod(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID ...C.jmethodID) C.jboolean {
//  fatal("Not implemented")
//  return nil
//}
//jboolean(*CallNonvirtualBooleanMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,va_list args);

extern jboolean GNI_CallNonvirtualBooleanMethodV(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, va_list* p4);

//jboolean(*CallNonvirtualBooleanMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,const jvalue * args);

extern jboolean GNI_CallNonvirtualBooleanMethodA(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, jvalue* p4);

//jbyte(*CallNonvirtualByteMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//func GNI_CallNonvirtualByteMethod(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID ...C.jmethodID) C.jbyte {
//  fatal("Not implemented")
//  return nil
//}
//jbyte(*CallNonvirtualByteMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,va_list args);

extern jbyte GNI_CallNonvirtualByteMethodV(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, va_list* p4);

//jbyte(*CallNonvirtualByteMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,const jvalue *args);

extern jbyte GNI_CallNonvirtualByteMethodA(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, jvalue* p4);

//jchar(*CallNonvirtualCharMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//func GNI_CallNonvirtualCharMethod(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID ...C.jmethodID) C.jchar {
//  fatal("Not implemented")
//  return nil
//}
//jchar(*CallNonvirtualCharMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,va_list args);

extern jchar GNI_CallNonvirtualCharMethodV(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, va_list* p4);

//jchar(*CallNonvirtualCharMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,const jvalue *args);

extern jchar GNI_CallNonvirtualCharMethodA(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, jvalue* p4);

//jshort(*CallNonvirtualShortMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//func GNI_CallNonvirtualShortMethod(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID ...C.jmethodID) C.jshort {
//  fatal("Not implemented")
//  return nil
//}
//jshort(*CallNonvirtualShortMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,va_list args);

extern jshort GNI_CallNonvirtualShortMethodV(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, va_list* p4);

//jshort(*CallNonvirtualShortMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,const jvalue *args);

extern jshort GNI_CallNonvirtualShortMethodA(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, jvalue* p4);

//jint(*CallNonvirtualIntMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//func GNI_CallNonvirtualIntMethod(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID ...C.jmethodID) C.jint {
//  fatal("Not implemented")
//  return nil
//}
//jint(*CallNonvirtualIntMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,va_list args);

extern jint GNI_CallNonvirtualIntMethodV(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, va_list* p4);

//jint(*CallNonvirtualIntMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,const jvalue *args);

extern jint GNI_CallNonvirtualIntMethodA(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, jvalue* p4);

//jlong(*CallNonvirtualLongMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//func GNI_CallNonvirtualLongMethod(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID ...C.jmethodID) C.jlong {
//  fatal("Not implemented")
//  return nil
//}
//jlong(*CallNonvirtualLongMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,va_list args);

extern jlong GNI_CallNonvirtualLongMethodV(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, va_list* p4);

//jlong(*CallNonvirtualLongMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,const jvalue *args);

extern jlong GNI_CallNonvirtualLongMethodA(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, jvalue* p4);

//jfloat(*CallNonvirtualFloatMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//func GNI_CallNonvirtualFloatMethod(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID ...C.jmethodID) C.jfloat {
//  fatal("Not implemented")
//  return nil
//}
//jfloat(*CallNonvirtualFloatMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,va_list args);

extern jfloat GNI_CallNonvirtualFloatMethodV(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, va_list* p4);

//jfloat(*CallNonvirtualFloatMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,const jvalue *args);

extern jfloat GNI_CallNonvirtualFloatMethodA(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, jvalue* p4);

//jdouble(*CallNonvirtualDoubleMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//func GNI_CallNonvirtualDoubleMethod(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID ...C.jmethodID) C.jdouble {
//  fatal("Not implemented")
//  return nil
//}
//jdouble(*CallNonvirtualDoubleMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,va_list args);

extern jdouble GNI_CallNonvirtualDoubleMethodV(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, va_list* p4);

//jdouble(*CallNonvirtualDoubleMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,const jvalue *args);

extern jdouble GNI_CallNonvirtualDoubleMethodA(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, jvalue* p4);

//void(*CallNonvirtualVoidMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//func GNI_CallNonvirtualVoidMethod(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID ...C.jmethodID) {
//  fatal("Not implemented")
//}
//void(*CallNonvirtualVoidMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,va_list args);

extern void GNI_CallNonvirtualVoidMethodV(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, va_list* p4);

//void(*CallNonvirtualVoidMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID,const jvalue * args);

extern void GNI_CallNonvirtualVoidMethodA(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, jvalue* p4);

//jfieldID(*GetFieldID)(JNIEnv *env, jclass clazz, const char *name, const char *sig);

extern jfieldID GNI_GetFieldID(JNIEnv* p0, jclass p1, char* p2, char* p3);

//jobject(*GetObjectField)(JNIEnv *env, jobject obj, jfieldID fieldID);

extern jobject GNI_GetObjectField(JNIEnv* p0, jobject p1, jfieldID p2);

//jboolean(*GetBooleanField)(JNIEnv *env, jobject obj, jfieldID fieldID);

extern jboolean GNI_GetBooleanField(JNIEnv* p0, jobject p1, jfieldID p2);

//jbyte(*GetByteField)(JNIEnv *env, jobject obj, jfieldID fieldID);

extern jbyte GNI_GetByteField(JNIEnv* p0, jobject p1, jfieldID p2);

//jchar(*GetCharField)(JNIEnv *env, jobject obj, jfieldID fieldID);

extern jchar GNI_GetCharField(JNIEnv* p0, jobject p1, jfieldID p2);

//jshort(*GetShortField)(JNIEnv *env, jobject obj, jfieldID fieldID);

extern jshort GNI_GetShortField(JNIEnv* p0, jobject p1, jfieldID p2);

//jint(*GetIntField)(JNIEnv *env, jobject obj, jfieldID fieldID);

extern jint GNI_GetIntField(JNIEnv* p0, jobject p1, jfieldID p2);

//jlong(*GetLongField)(JNIEnv *env, jobject obj, jfieldID fieldID);

extern jlong GNI_GetLongField(JNIEnv* p0, jobject p1, jfieldID p2);

//jfloat(*GetFloatField)(JNIEnv *env, jobject obj, jfieldID fieldID);

extern jfloat GNI_GetFloatField(JNIEnv* p0, jobject p1, jfieldID p2);

//jdouble(*GetDoubleField)(JNIEnv *env, jobject obj, jfieldID fieldID);

extern jdouble GNI_GetDoubleField(JNIEnv* p0, jobject p1, jfieldID p2);

//void(*SetObjectField)(JNIEnv *env, jobject obj, jfieldID fieldID, jobject val);

extern void GNI_SetObjectField(JNIEnv* p0, jobject p1, jfieldID p2, jobject p3);

//void(*SetBooleanField)(JNIEnv *env, jobject obj, jfieldID fieldID, jboolean val);

extern void GNI_SetBooleanField(JNIEnv* p0, jobject p1, jfieldID p2, jboolean p3);

//void(*SetByteField)(JNIEnv *env, jobject obj, jfieldID fieldID, jbyte val);

extern void GNI_SetByteField(JNIEnv* p0, jobject p1, jfieldID p2, jbyte p3);

//void(*SetCharField)(JNIEnv *env, jobject obj, jfieldID fieldID, jchar val);

extern void GNI_SetCharField(JNIEnv* p0, jobject p1, jfieldID p2, jchar p3);

//void(*SetShortField)(JNIEnv *env, jobject obj, jfieldID fieldID, jshort val);

extern void GNI_SetShortField(JNIEnv* p0, jobject p1, jfieldID p2, jshort p3);

//void(*SetIntField)(JNIEnv *env, jobject obj, jfieldID fieldID, jint val);

extern void GNI_SetIntField(JNIEnv* p0, jobject p1, jfieldID p2, jint p3);

//void(*SetLongField)(JNIEnv *env, jobject obj, jfieldID fieldID, jlong val);

extern void GNI_SetLongField(JNIEnv* p0, jobject p1, jfieldID p2, jlong p3);

//void(*SetFloatField)(JNIEnv *env, jobject obj, jfieldID fieldID, jfloat val);

extern void GNI_SetFloatField(JNIEnv* p0, jobject p1, jfieldID p2, jfloat p3);

//void(*SetDoubleField)(JNIEnv *env, jobject obj, jfieldID fieldID, jdouble val);

extern void GNI_SetDoubleField(JNIEnv* p0, jobject p1, jfieldID p2, jdouble p3);

//jmethodID(*GetStaticMethodID)(JNIEnv *env, jclass clazz, const char *name, const char *sig);

extern jmethodID GNI_GetStaticMethodID(JNIEnv* p0, jclass p1, char* p2, char* p3);

//jobject(*CallStaticObjectMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//func GNI_CallStaticObjectMethod(env *C.JNIEnv, clazz C.jclass, methodID ...C.jmethodID) C.jobject {
//  fatal("Not implemented")
//  return nil
//}
//jobject(*CallStaticObjectMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);

extern jobject GNI_CallStaticObjectMethodV(JNIEnv* p0, jclass p1, jmethodID p2, va_list* p3);

//jobject(*CallStaticObjectMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

extern jobject GNI_CallStaticObjectMethodA(JNIEnv* p0, jclass p1, jmethodID p2, jvalue* p3);

//jboolean(*CallStaticBooleanMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//func GNI_CallStaticBooleanMethod(env *C.JNIEnv, clazz C.jclass, methodID ...C.jmethodID) C.jboolean {
//  fatal("Not implemented")
//  return nil
//}
//jboolean(*CallStaticBooleanMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);

extern jboolean GNI_CallStaticBooleanMethodV(JNIEnv* p0, jclass p1, jmethodID p2, va_list* p3);

//jboolean(*CallStaticBooleanMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

extern jboolean GNI_CallStaticBooleanMethodA(JNIEnv* p0, jclass p1, jmethodID p2, jvalue* p3);

//jbyte(*CallStaticByteMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//func GNI_CallStaticByteMethod(env *C.JNIEnv, clazz C.jclass, methodID ...C.jmethodID) C.jbyte {
//  fatal("Not implemented")
//  return nil
//}
//jbyte(*CallStaticByteMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);

extern jbyte GNI_CallStaticByteMethodV(JNIEnv* p0, jclass p1, jmethodID p2, va_list* p3);

//jbyte(*CallStaticByteMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

extern jbyte GNI_CallStaticByteMethodA(JNIEnv* p0, jclass p1, jmethodID p2, jvalue* p3);

//jchar(*CallStaticCharMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//func GNI_CallStaticCharMethod(env *C.JNIEnv, clazz C.jclass, methodID ...C.jmethodID) C.jchar {
//  fatal("Not implemented")
//  return nil
//}
//jchar(*CallStaticCharMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);

extern jchar GNI_CallStaticCharMethodV(JNIEnv* p0, jclass p1, jmethodID p2, va_list* p3);

//jchar(*CallStaticCharMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

extern jchar GNI_CallStaticCharMethodA(JNIEnv* p0, jclass p1, jmethodID p2, jvalue* p3);

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

extern jshort GNI_CallStaticShortMethodA(JNIEnv* p0, jclass p1, jmethodID p2, jvalue* p3);

//jint(*CallStaticIntMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//func GNI_CallStaticIntMethod(env *C.JNIEnv, clazz C.jclass, methodID ...C.jmethodID) C.jint {
//  fatal("Not implemented")
//  return nil
//}
//jint(*CallStaticIntMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);

extern jint GNI_CallStaticIntMethodV(JNIEnv* p0, jclass p1, jmethodID p2, va_list* p3);

//jint(*CallStaticIntMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

extern jint GNI_CallStaticIntMethodA(JNIEnv* p0, jclass p1, jmethodID p2, jvalue* p3);

//jlong(*CallStaticLongMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//func GNI_CallStaticLongMethod(env *C.JNIEnv, clazz C.jclass, methodID ...C.jmethodID) C.jlong {
//  fatal("Not implemented")
//  return nil
//}
//jlong(*CallStaticLongMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);

extern jlong GNI_CallStaticLongMethodV(JNIEnv* p0, jclass p1, jmethodID p2, va_list* p3);

//jlong(*CallStaticLongMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

extern jlong GNI_CallStaticLongMethodA(JNIEnv* p0, jclass p1, jmethodID p2, jvalue* p3);

//jfloat(*CallStaticFloatMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//func GNI_CallStaticFloatMethod(env *C.JNIEnv, clazz C.jclass, methodID ...C.jmethodID) C.jfloat {
//  fatal("Not implemented")
//  return nil
//}
//jfloat(*CallStaticFloatMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);

extern jfloat GNI_CallStaticFloatMethodV(JNIEnv* p0, jclass p1, jmethodID p2, va_list* p3);

//jfloat(*CallStaticFloatMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

extern jfloat GNI_CallStaticFloatMethodA(JNIEnv* p0, jclass p1, jmethodID p2, jvalue* p3);

//jdouble(*CallStaticDoubleMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//func GNI_CallStaticDoubleMethod(env *C.JNIEnv, clazz C.jclass, methodID ...C.jmethodID) C.jdouble {
//  fatal("Not implemented")
//  return nil
//}
//jdouble(*CallStaticDoubleMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);

extern jdouble GNI_CallStaticDoubleMethodV(JNIEnv* p0, jclass p1, jmethodID p2, va_list* p3);

//jdouble(*CallStaticDoubleMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

extern jdouble GNI_CallStaticDoubleMethodA(JNIEnv* p0, jclass p1, jmethodID p2, jvalue* p3);

//void(*CallStaticVoidMethod)(JNIEnv *env, jclass cls, jmethodID methodID, ...);
//func GNI_CallStaticVoidMethod(env *C.JNIEnv, cls C.jclass, methodID ...C.jmethodID) {
//  fatal("Not implemented")
//}
//void(*CallStaticVoidMethodV)(JNIEnv *env, jclass cls, jmethodID methodID, va_list args);

extern void GNI_CallStaticVoidMethodV(JNIEnv* p0, jclass p1, jmethodID p2, va_list* p3);

//void(*CallStaticVoidMethodA)(JNIEnv *env, jclass cls, jmethodID methodID, const jvalue * args);

extern void GNI_CallStaticVoidMethodA(JNIEnv* p0, jclass p1, jmethodID p2, jvalue* p3);

//jfieldID(*GetStaticFieldID)(JNIEnv *env, jclass clazz, const char *name, const char *sig);

extern jfieldID GNI_GetStaticFieldID(JNIEnv* p0, jclass p1, char* p2, char* p3);

//jobject(*GetStaticObjectField)(JNIEnv *env, jclass clazz, jfieldID fieldID);

extern jobject GNI_GetStaticObjectField(JNIEnv* p0, jclass p1, jfieldID p2);

//jboolean(*GetStaticBooleanField)(JNIEnv *env, jclass clazz, jfieldID fieldID);

extern jboolean GNI_GetStaticBooleanField(JNIEnv* p0, jclass p1, jfieldID p2);

//jbyte(*GetStaticByteField)(JNIEnv *env, jclass clazz, jfieldID fieldID);

extern jbyte GNI_GetStaticByteField(JNIEnv* p0, jclass p1, jfieldID p2);

//jchar(*GetStaticCharField)(JNIEnv *env, jclass clazz, jfieldID fieldID);

extern jchar GNI_GetStaticCharField(JNIEnv* p0, jclass p1, jfieldID p2);

//jshort(*GetStaticShortField)(JNIEnv *env, jclass clazz, jfieldID fieldID);

extern jshort GNI_GetStaticShortField(JNIEnv* p0, jclass p1, jfieldID p2);

//jint(*GetStaticIntField)(JNIEnv *env, jclass clazz, jfieldID fieldID);

extern jint GNI_GetStaticIntField(JNIEnv* p0, jclass p1, jfieldID p2);

//jlong(*GetStaticLongField)(JNIEnv *env, jclass clazz, jfieldID fieldID);

extern jlong GNI_GetStaticLongField(JNIEnv* p0, jclass p1, jfieldID p2);

//jfloat(*GetStaticFloatField)(JNIEnv *env, jclass clazz, jfieldID fieldID);

extern jfloat GNI_GetStaticFloatField(JNIEnv* p0, jclass p1, jfieldID p2);

//jdouble(*GetStaticDoubleField)(JNIEnv *env, jclass clazz, jfieldID fieldID);

extern jdouble GNI_GetStaticDoubleField(JNIEnv* p0, jclass p1, jfieldID p2);

//void(*SetStaticObjectField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jobject value);

extern void GNI_SetStaticObjectField(JNIEnv* p0, jclass p1, jfieldID p2, jobject p3);

//void(*SetStaticBooleanField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jboolean value);

extern void GNI_SetStaticBooleanField(JNIEnv* p0, jclass p1, jfieldID p2, jboolean p3);

//void(*SetStaticByteField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jbyte value);

extern void GNI_SetStaticByteField(JNIEnv* p0, jclass p1, jfieldID p2, jbyte p3);

//void(*SetStaticCharField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jchar value);

extern void GNI_SetStaticCharField(JNIEnv* p0, jclass p1, jfieldID p2, jchar p3);

//void(*SetStaticShortField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jshort value);

extern void GNI_SetStaticShortField(JNIEnv* p0, jclass p1, jfieldID p2, jshort p3);

//void(*SetStaticIntField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jint value);

extern void GNI_SetStaticIntField(JNIEnv* p0, jclass p1, jfieldID p2, jint p3);

//void(*SetStaticLongField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jlong value);

extern void GNI_SetStaticLongField(JNIEnv* p0, jclass p1, jfieldID p2, jlong p3);

//void(*SetStaticFloatField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jfloat value);

extern void GNI_SetStaticFloatField(JNIEnv* p0, jclass p1, jfieldID p2, jfloat p3);

//void(*SetStaticDoubleField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jdouble value);

extern void GNI_SetStaticDoubleField(JNIEnv* p0, jclass p1, jfieldID p2, jdouble p3);

//jstring(*NewString)(JNIEnv *env, const jchar *unicode, jsize len);

extern jstring GNI_NewString(JNIEnv* p0, jchar* p1, jsize p2);

//jsize(*GetStringLength)(JNIEnv *env, jstring str);

extern jsize GNI_GetStringLength(JNIEnv* p0, jstring p1);

//const jchar *(*GetStringChars)(JNIEnv *env, jstring str, jboolean *isCopy);

extern jchar* GNI_GetStringChars(JNIEnv* p0, jstring p1, jboolean* p2);

//void(*ReleaseStringChars)(JNIEnv *env, jstring str, const jchar *chars);

extern void GNI_ReleaseStringChars(JNIEnv* p0, jstring p1, jchar p2);

//jstring(*NewStringUTF)(JNIEnv *env, const char *utf);

extern jstring GNI_NewStringUTF(JNIEnv* p0, char* p1);

//jsize(*GetStringUTFLength)(JNIEnv *env, jstring str);

extern jsize GNI_GetStringUTFLength(JNIEnv* p0, jstring p1);

//const char*(*GetStringUTFChars)(JNIEnv *env, jstring str, jboolean *isCopy);

extern char* GNI_GetStringUTFChars(JNIEnv* p0, jstring p1, jboolean* p2);

//void(*ReleaseStringUTFChars)(JNIEnv *env, jstring str, const char* chars);

extern void GNI_ReleaseStringUTFChars(JNIEnv* p0, jstring p1, char* p2);

//jsize(*GetArrayLength)(JNIEnv *env, jarray array);

extern jsize GNI_GetArrayLength(JNIEnv* p0, jarray* p1);

//jobjectArray(*NewObjectArray)(JNIEnv *env, jsize len, jclass clazz, jobject init);

extern jobjectArray GNI_NewObjectArray(JNIEnv* p0, jsize p1, jclass p2, jobject p3);

//jobject(*GetObjectArrayElement)(JNIEnv *env, jobjectArray array, jsize index);

extern jobject GNI_GetObjectArrayElement(JNIEnv* p0, jobjectArray p1, jsize p2);

//void(*SetObjectArrayElement)(JNIEnv *env, jobjectArray array, jsize index, jobject val);

extern void GNI_SetObjectArrayElement(JNIEnv* p0, jobjectArray p1, jsize p2, jobject p3);

//jbooleanArray(*NewBooleanArray)(JNIEnv *env, jsize len);

extern jbooleanArray GNI_NewBooleanArray(JNIEnv* p0, jsize p1);

//jbyteArray(*NewByteArray)(JNIEnv *env, jsize len);

extern jbyteArray GNI_NewByteArray(JNIEnv* p0, jsize p1);

//jcharArray(*NewCharArray)(JNIEnv *env, jsize len);

extern jcharArray GNI_NewCharArray(JNIEnv* p0, jsize p1);

//jshortArray(*NewShortArray)(JNIEnv *env, jsize len);

extern jshortArray GNI_NewShortArray(JNIEnv* p0, jsize p1);

//jintArray(*NewIntArray)(JNIEnv *env, jsize len);

extern jintArray GNI_NewIntArray(JNIEnv* p0, jsize p1);

//jlongArray(*NewLongArray)(JNIEnv *env, jsize len);

extern jlongArray GNI_NewLongArray(JNIEnv* p0, jsize p1);

//jfloatArray(*NewFloatArray)(JNIEnv *env, jsize len);

extern jfloatArray GNI_NewFloatArray(JNIEnv* p0, jsize p1);

//jdoubleArray(*NewDoubleArray)(JNIEnv *env, jsize len);

extern jdoubleArray GNI_NewDoubleArray(JNIEnv* p0, jsize p1);

//jboolean *(*GetBooleanArrayElements)(JNIEnv *env, jbooleanArray array, jboolean *isCopy);

extern jboolean GNI_GetBooleanArrayElements(JNIEnv* p0, jbooleanArray p1, jboolean* p2);

//jbyte *(*GetByteArrayElements)(JNIEnv *env, jbyteArray array, jboolean *isCopy);

extern jbyte GNI_GetByteArrayElements(JNIEnv* p0, jbyteArray p1, jboolean* p2);

//jchar *(*GetCharArrayElements)(JNIEnv *env, jcharArray array, jboolean *isCopy);

extern jchar GNI_GetCharArrayElements(JNIEnv* p0, jcharArray p1, jboolean* p2);

//jshort *(*GetShortArrayElements)(JNIEnv *env, jshortArray array, jboolean *isCopy);

extern jshort GNI_GetShortArrayElements(JNIEnv* p0, jshortArray p1, jboolean* p2);

//jint *(*GetIntArrayElements)(JNIEnv *env, jintArray array, jboolean *isCopy);

extern jint GNI_GetIntArrayElements(JNIEnv* p0, jintArray p1, jboolean* p2);

//jlong *(*GetLongArrayElements)(JNIEnv *env, jlongArray array, jboolean *isCopy);

extern jlong GNI_GetLongArrayElements(JNIEnv* p0, jlongArray p1, jboolean* p2);

//jfloat *(*GetFloatArrayElements)(JNIEnv *env, jfloatArray array, jboolean *isCopy);

extern jfloat GNI_GetFloatArrayElements(JNIEnv* p0, jfloatArray p1, jboolean* p2);

//jdouble *(*GetDoubleArrayElements)(JNIEnv *env, jdoubleArray array, jboolean *isCopy);

extern jdouble GNI_GetDoubleArrayElements(JNIEnv* p0, jdoubleArray p1, jboolean* p2);

//void(*ReleaseBooleanArrayElements)(JNIEnv *env, jbooleanArray array, jboolean *elems, jint mode);

extern void GNI_ReleaseBooleanArrayElements(JNIEnv* p0, jbooleanArray p1, jboolean* p2, jint p3);

//void(*ReleaseByteArrayElements)(JNIEnv *env, jbyteArray array, jbyte *elems, jint mode);

extern void GNI_ReleaseByteArrayElements(JNIEnv* p0, jbyteArray p1, jbyte* p2, jint p3);

//void(*ReleaseCharArrayElements)(JNIEnv *env, jcharArray array, jchar *elems, jint mode);

extern void GNI_ReleaseCharArrayElements(JNIEnv* p0, jcharArray p1, jchar* p2, jint p3);

//void(*ReleaseShortArrayElements)(JNIEnv *env, jshortArray array, jshort *elems, jint mode);

extern void GNI_ReleaseShortArrayElements(JNIEnv* p0, jshortArray p1, jshort* p2, jint p3);

//void(*ReleaseIntArrayElements)(JNIEnv *env, jintArray array, jint *elems, jint mode);

extern void GNI_ReleaseIntArrayElements(JNIEnv* p0, jintArray p1, jint* p2, jint p3);

//void(*ReleaseLongArrayElements)(JNIEnv *env, jlongArray array, jlong *elems, jint mode);

extern void GNI_ReleaseLongArrayElements(JNIEnv* p0, jlongArray p1, jlong* p2, jint p3);

//void(*ReleaseFloatArrayElements)(JNIEnv *env, jfloatArray array, jfloat *elems, jint mode);

extern void GNI_ReleaseFloatArrayElements(JNIEnv* p0, jfloatArray p1, jfloat* p2, jint p3);

//void(*ReleaseDoubleArrayElements)(JNIEnv *env, jdoubleArray array, jdouble *elems, jint mode);

extern void GNI_ReleaseDoubleArrayElements(JNIEnv* p0, jdoubleArray p1, jdouble* p2, jint p3);

//void(*GetBooleanArrayRegion)(JNIEnv *env, jbooleanArray array, jsize start, jsize l, jboolean *buf);

extern void GNI_GetBooleanArrayRegion(JNIEnv* p0, jbooleanArray p1, jsize p2, jsize p3, jboolean* p4);

//void(*GetByteArrayRegion)(JNIEnv *env, jbyteArray array, jsize start, jsize len, jbyte *buf);

extern void GNI_GetByteArrayRegion(JNIEnv* p0, jbyteArray p1, jsize p2, jsize p3, jbyte* p4);

//void(*GetCharArrayRegion)(JNIEnv *env, jcharArray array, jsize start, jsize len, jchar *buf);

extern void GNI_GetCharArrayRegion(JNIEnv* p0, jcharArray p1, jsize p2, jsize p3, jchar* p4);

//void(*GetShortArrayRegion)(JNIEnv *env, jshortArray array, jsize start, jsize len, jshort *buf);

extern void GNI_GetShortArrayRegion(JNIEnv* p0, jshortArray p1, jsize p2, jsize p3, jshort* p4);

//void(*GetIntArrayRegion)(JNIEnv *env, jintArray array, jsize start, jsize len, jint *buf);

extern void GNI_GetIntArrayRegion(JNIEnv* p0, jintArray p1, jsize p2, jsize p3, jint* p4);

//void(*GetLongArrayRegion)(JNIEnv *env, jlongArray array, jsize start, jsize len, jlong *buf);

extern void GNI_GetLongArrayRegion(JNIEnv* p0, jlongArray p1, jsize p2, jsize p3, jlong* p4);

//void(*GetFloatArrayRegion)(JNIEnv *env, jfloatArray array, jsize start, jsize len, jfloat *buf);

extern void GNI_GetFloatArrayRegion(JNIEnv* p0, jfloatArray p1, jsize p2, jsize p3, jfloat* p4);

//void(*GetDoubleArrayRegion)(JNIEnv *env, jdoubleArray array, jsize start, jsize len, jdouble *buf);

extern void GNI_GetDoubleArrayRegion(JNIEnv* p0, jdoubleArray p1, jsize p2, jsize p3, jdouble* p4);

//void(*SetBooleanArrayRegion)(JNIEnv *env, jbooleanArray array, jsize start, jsize l, const jboolean *buf);

extern void GNI_SetBooleanArrayRegion(JNIEnv* p0, jbooleanArray p1, jsize p2, jsize p3, jboolean* p4);

//void(*SetByteArrayRegion)(JNIEnv *env, jbyteArray array, jsize start, jsize len, const jbyte *buf);

extern void GNI_SetByteArrayRegion(JNIEnv* p0, jbyteArray p1, jsize p2, jsize p3, jbyte* p4);

//void(*SetCharArrayRegion)(JNIEnv *env, jcharArray array, jsize start, jsize len, const jchar *buf);

extern void GNI_SetCharArrayRegion(JNIEnv* p0, jcharArray p1, jsize p2, jsize p3, jchar* p4);

//void(*SetShortArrayRegion)(JNIEnv *env, jshortArray array, jsize start, jsize len, const jshort *buf);

extern void GNI_SetShortArrayRegion(JNIEnv* p0, jshortArray p1, jsize p2, jsize p3, jshort* p4);

//void(*SetIntArrayRegion)(JNIEnv *env, jintArray array, jsize start, jsize len, const jint *buf);

extern void GNI_SetIntArrayRegion(JNIEnv* p0, jintArray p1, jsize p2, jsize p3, jint* p4);

//void(*SetLongArrayRegion)(JNIEnv *env, jlongArray array, jsize start, jsize len, const jlong *buf);

extern void GNI_SetLongArrayRegion(JNIEnv* p0, jlongArray p1, jsize p2, jsize p3, jlong* p4);

//void(*SetFloatArrayRegion)(JNIEnv *env, jfloatArray array, jsize start, jsize len, const jfloat *buf);

extern void GNI_SetFloatArrayRegion(JNIEnv* p0, jfloatArray p1, jsize p2, jsize p3, jfloat* p4);

//void(*SetDoubleArrayRegion)(JNIEnv *env, jdoubleArray array, jsize start, jsize len, const jdouble *buf);

extern void GNI_SetDoubleArrayRegion(JNIEnv* p0, jdoubleArray p1, jsize p2, jsize p3, jdouble* p4);

//jint(*RegisterNatives)(JNIEnv *env, jclass clazz, const JNINativeMethod *methods, jint nMethods);

extern jint GNI_RegisterNatives(JNIEnv* p0, jclass p1, JNINativeMethod* p2, jint p3);

//jint(*UnregisterNatives)(JNIEnv *env, jclass clazz);

extern jint GNI_UnregisterNatives(JNIEnv* p0, jclass p1);

//jint(*MonitorEnter)(JNIEnv *env, jobject obj);

extern jint GNI_MonitorEnter(JNIEnv* p0, jobject p1);

//jint(*MonitorExit)(JNIEnv *env, jobject obj);

extern jint GNI_MonitorExit(JNIEnv* p0, jobject p1);

//jint(*GetJavaVM)(JNIEnv *env, JavaVM **vm);

extern jint GNI_GetJavaVM(JNIEnv* p0, void* p1);

//void(*GetStringRegion)(JNIEnv *env, jstring str, jsize start, jsize len, jchar *buf);

extern void GNI_GetStringRegion(JNIEnv* p0, jstring p1, jsize p2, jsize p3, jchar* p4);

//void(*GetStringUTFRegion)(JNIEnv *env, jstring str, jsize start, jsize len, char *buf);

extern void GNI_GetStringUTFRegion(JNIEnv* p0, jstring p1, jsize p2, jsize p3, char* p4);

//void *(*GetPrimitiveArrayCritical)(JNIEnv *env, jarray array, jboolean *isCopy);

extern void GNI_GetPrimitiveArrayCritical(JNIEnv* p0, jarray p1, jboolean* p2);

//void(*ReleasePrimitiveArrayCritical)(JNIEnv *env, jarray array, void *carray, jint mode);

extern void GNI_ReleasePrimitiveArrayCritical(JNIEnv* p0, jarray p1, void* p2, jint p3);

//const jchar *(*GetStringCritical)(JNIEnv *env, jstring string, jboolean *isCopy);

extern jchar* GNI_GetStringCritical(JNIEnv* p0, jstring p1, jboolean* p2);

//void(*ReleaseStringCritical)(JNIEnv *env, jstring string, const jchar *cstring);

extern void GNI_ReleaseStringCritical(JNIEnv* p0, jstring p1, jchar* p2);

//jweak(*NewWeakGlobalRef)(JNIEnv *env, jobject obj);

extern jweak GNI_NewWeakGlobalRef(JNIEnv* p0, jobject p1);

//void(*DeleteWeakGlobalRef)(JNIEnv *env, jweak ref);

extern void GNI_DeleteWeakGlobalRef(JNIEnv* p0, jweak p1);

//jboolean(*ExceptionCheck)(JNIEnv *env);

extern jboolean GNI_ExceptionCheck(JNIEnv* p0);

//jobject(*NewDirectByteBuffer)(JNIEnv* env, void* address, jlong capacity);

extern jobject GNI_NewDirectByteBuffer(JNIEnv* p0, void* p1, jlong p2);

//void*(*GetDirectBufferAddress)(JNIEnv* env, jobject buf);

extern void* GNI_GetDirectBufferAddress(JNIEnv* p0, jobject p1);

//jlong(*GetDirectBufferCapacity)(JNIEnv* env, jobject buf);

extern jlong GNI_GetDirectBufferCapacity(JNIEnv* p0, jobject p1);

/* New JNI 1.6 Features */
//jobjectRefType(*GetObjectRefType)(JNIEnv* env, jobject obj);

extern jobjectRefType GNI_GetObjectRefType(JNIEnv* p0, jobject p1);

#ifdef __cplusplus
}
#endif
