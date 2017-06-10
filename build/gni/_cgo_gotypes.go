// Created by cgo - DO NOT EDIT

package gvm

import "unsafe"

import _ "runtime/cgo"

import "syscall"

var _ syscall.Errno
func _Cgo_ptr(ptr unsafe.Pointer) unsafe.Pointer { return ptr }

//go:linkname _Cgo_always_false runtime.cgoAlwaysFalse
var _Cgo_always_false bool
//go:linkname _Cgo_use runtime.cgoUse
func _Cgo_use(interface{})
type _Ctype_JNIEnv *_Ctype_struct_JNINativeInterface_

type _Ctype_JNINativeMethod _Ctype_struct___0

type _Ctype___builtin_va_list [1]_Ctype_struct___va_list_tag

type _Ctype_char int8

type _Ctype_double float64

type _Ctype_float float32

type _Ctype_int int32

type _Ctype_jarray _Ctype_jobject

type _Ctype_jboolean _Ctype_uchar

type _Ctype_jbooleanArray _Ctype_jarray

type _Ctype_jbyte _Ctype_schar

type _Ctype_jbyteArray _Ctype_jarray

type _Ctype_jchar _Ctype_ushort

type _Ctype_jcharArray _Ctype_jarray

type _Ctype_jclass _Ctype_jobject

type _Ctype_jdouble _Ctype_double

type _Ctype_jdoubleArray _Ctype_jarray

type _Ctype_jfieldID *_Ctype_struct__jfieldID

type _Ctype_jfloat _Ctype_float

type _Ctype_jfloatArray _Ctype_jarray

type _Ctype_jint _Ctype_int

type _Ctype_jintArray _Ctype_jarray

type _Ctype_jlong _Ctype_long

type _Ctype_jlongArray _Ctype_jarray

type _Ctype_jmethodID *_Ctype_struct__jmethodID

type _Ctype_jobject *_Ctype_struct__jobject

type _Ctype_jobjectArray _Ctype_jarray

type _Ctype_jobjectRefType uint32

type _Ctype_jshort _Ctype_short

type _Ctype_jshortArray _Ctype_jarray

type _Ctype_jsize _Ctype_jint

type _Ctype_jstring _Ctype_jobject

type _Ctype_jthrowable _Ctype_jobject

type _Ctype_jvalue [8]byte

type _Ctype_jweak _Ctype_jobject

type _Ctype_long int64

type _Ctype_schar int8

type _Ctype_short int16

type _Ctype_struct_JNINativeInterface_ struct {
	reserved0			unsafe.Pointer
	reserved1			unsafe.Pointer
	reserved2			unsafe.Pointer
	reserved3			unsafe.Pointer
	GetVersion			*[0]byte
	DefineClass			*[0]byte
	FindClass			*[0]byte
	FromReflectedMethod		*[0]byte
	FromReflectedField		*[0]byte
	ToReflectedMethod		*[0]byte
	GetSuperclass			*[0]byte
	IsAssignableFrom		*[0]byte
	ToReflectedField		*[0]byte
	Throw				*[0]byte
	ThrowNew			*[0]byte
	ExceptionOccurred		*[0]byte
	ExceptionDescribe		*[0]byte
	ExceptionClear			*[0]byte
	FatalError			*[0]byte
	PushLocalFrame			*[0]byte
	PopLocalFrame			*[0]byte
	NewGlobalRef			*[0]byte
	DeleteGlobalRef			*[0]byte
	DeleteLocalRef			*[0]byte
	IsSameObject			*[0]byte
	NewLocalRef			*[0]byte
	EnsureLocalCapacity		*[0]byte
	AllocObject			*[0]byte
	NewObject			*[0]byte
	NewObjectV			*[0]byte
	NewObjectA			*[0]byte
	GetObjectClass			*[0]byte
	IsInstanceOf			*[0]byte
	GetMethodID			*[0]byte
	CallObjectMethod		*[0]byte
	CallObjectMethodV		*[0]byte
	CallObjectMethodA		*[0]byte
	CallBooleanMethod		*[0]byte
	CallBooleanMethodV		*[0]byte
	CallBooleanMethodA		*[0]byte
	CallByteMethod			*[0]byte
	CallByteMethodV			*[0]byte
	CallByteMethodA			*[0]byte
	CallCharMethod			*[0]byte
	CallCharMethodV			*[0]byte
	CallCharMethodA			*[0]byte
	CallShortMethod			*[0]byte
	CallShortMethodV		*[0]byte
	CallShortMethodA		*[0]byte
	CallIntMethod			*[0]byte
	CallIntMethodV			*[0]byte
	CallIntMethodA			*[0]byte
	CallLongMethod			*[0]byte
	CallLongMethodV			*[0]byte
	CallLongMethodA			*[0]byte
	CallFloatMethod			*[0]byte
	CallFloatMethodV		*[0]byte
	CallFloatMethodA		*[0]byte
	CallDoubleMethod		*[0]byte
	CallDoubleMethodV		*[0]byte
	CallDoubleMethodA		*[0]byte
	CallVoidMethod			*[0]byte
	CallVoidMethodV			*[0]byte
	CallVoidMethodA			*[0]byte
	CallNonvirtualObjectMethod	*[0]byte
	CallNonvirtualObjectMethodV	*[0]byte
	CallNonvirtualObjectMethodA	*[0]byte
	CallNonvirtualBooleanMethod	*[0]byte
	CallNonvirtualBooleanMethodV	*[0]byte
	CallNonvirtualBooleanMethodA	*[0]byte
	CallNonvirtualByteMethod	*[0]byte
	CallNonvirtualByteMethodV	*[0]byte
	CallNonvirtualByteMethodA	*[0]byte
	CallNonvirtualCharMethod	*[0]byte
	CallNonvirtualCharMethodV	*[0]byte
	CallNonvirtualCharMethodA	*[0]byte
	CallNonvirtualShortMethod	*[0]byte
	CallNonvirtualShortMethodV	*[0]byte
	CallNonvirtualShortMethodA	*[0]byte
	CallNonvirtualIntMethod		*[0]byte
	CallNonvirtualIntMethodV	*[0]byte
	CallNonvirtualIntMethodA	*[0]byte
	CallNonvirtualLongMethod	*[0]byte
	CallNonvirtualLongMethodV	*[0]byte
	CallNonvirtualLongMethodA	*[0]byte
	CallNonvirtualFloatMethod	*[0]byte
	CallNonvirtualFloatMethodV	*[0]byte
	CallNonvirtualFloatMethodA	*[0]byte
	CallNonvirtualDoubleMethod	*[0]byte
	CallNonvirtualDoubleMethodV	*[0]byte
	CallNonvirtualDoubleMethodA	*[0]byte
	CallNonvirtualVoidMethod	*[0]byte
	CallNonvirtualVoidMethodV	*[0]byte
	CallNonvirtualVoidMethodA	*[0]byte
	GetFieldID			*[0]byte
	GetObjectField			*[0]byte
	GetBooleanField			*[0]byte
	GetByteField			*[0]byte
	GetCharField			*[0]byte
	GetShortField			*[0]byte
	GetIntField			*[0]byte
	GetLongField			*[0]byte
	GetFloatField			*[0]byte
	GetDoubleField			*[0]byte
	SetObjectField			*[0]byte
	SetBooleanField			*[0]byte
	SetByteField			*[0]byte
	SetCharField			*[0]byte
	SetShortField			*[0]byte
	SetIntField			*[0]byte
	SetLongField			*[0]byte
	SetFloatField			*[0]byte
	SetDoubleField			*[0]byte
	GetStaticMethodID		*[0]byte
	CallStaticObjectMethod		*[0]byte
	CallStaticObjectMethodV		*[0]byte
	CallStaticObjectMethodA		*[0]byte
	CallStaticBooleanMethod		*[0]byte
	CallStaticBooleanMethodV	*[0]byte
	CallStaticBooleanMethodA	*[0]byte
	CallStaticByteMethod		*[0]byte
	CallStaticByteMethodV		*[0]byte
	CallStaticByteMethodA		*[0]byte
	CallStaticCharMethod		*[0]byte
	CallStaticCharMethodV		*[0]byte
	CallStaticCharMethodA		*[0]byte
	CallStaticShortMethod		*[0]byte
	CallStaticShortMethodV		*[0]byte
	CallStaticShortMethodA		*[0]byte
	CallStaticIntMethod		*[0]byte
	CallStaticIntMethodV		*[0]byte
	CallStaticIntMethodA		*[0]byte
	CallStaticLongMethod		*[0]byte
	CallStaticLongMethodV		*[0]byte
	CallStaticLongMethodA		*[0]byte
	CallStaticFloatMethod		*[0]byte
	CallStaticFloatMethodV		*[0]byte
	CallStaticFloatMethodA		*[0]byte
	CallStaticDoubleMethod		*[0]byte
	CallStaticDoubleMethodV		*[0]byte
	CallStaticDoubleMethodA		*[0]byte
	CallStaticVoidMethod		*[0]byte
	CallStaticVoidMethodV		*[0]byte
	CallStaticVoidMethodA		*[0]byte
	GetStaticFieldID		*[0]byte
	GetStaticObjectField		*[0]byte
	GetStaticBooleanField		*[0]byte
	GetStaticByteField		*[0]byte
	GetStaticCharField		*[0]byte
	GetStaticShortField		*[0]byte
	GetStaticIntField		*[0]byte
	GetStaticLongField		*[0]byte
	GetStaticFloatField		*[0]byte
	GetStaticDoubleField		*[0]byte
	SetStaticObjectField		*[0]byte
	SetStaticBooleanField		*[0]byte
	SetStaticByteField		*[0]byte
	SetStaticCharField		*[0]byte
	SetStaticShortField		*[0]byte
	SetStaticIntField		*[0]byte
	SetStaticLongField		*[0]byte
	SetStaticFloatField		*[0]byte
	SetStaticDoubleField		*[0]byte
	NewString			*[0]byte
	GetStringLength			*[0]byte
	GetStringChars			*[0]byte
	ReleaseStringChars		*[0]byte
	NewStringUTF			*[0]byte
	GetStringUTFLength		*[0]byte
	GetStringUTFChars		*[0]byte
	ReleaseStringUTFChars		*[0]byte
	GetArrayLength			*[0]byte
	NewObjectArray			*[0]byte
	GetObjectArrayElement		*[0]byte
	SetObjectArrayElement		*[0]byte
	NewBooleanArray			*[0]byte
	NewByteArray			*[0]byte
	NewCharArray			*[0]byte
	NewShortArray			*[0]byte
	NewIntArray			*[0]byte
	NewLongArray			*[0]byte
	NewFloatArray			*[0]byte
	NewDoubleArray			*[0]byte
	GetBooleanArrayElements		*[0]byte
	GetByteArrayElements		*[0]byte
	GetCharArrayElements		*[0]byte
	GetShortArrayElements		*[0]byte
	GetIntArrayElements		*[0]byte
	GetLongArrayElements		*[0]byte
	GetFloatArrayElements		*[0]byte
	GetDoubleArrayElements		*[0]byte
	ReleaseBooleanArrayElements	*[0]byte
	ReleaseByteArrayElements	*[0]byte
	ReleaseCharArrayElements	*[0]byte
	ReleaseShortArrayElements	*[0]byte
	ReleaseIntArrayElements		*[0]byte
	ReleaseLongArrayElements	*[0]byte
	ReleaseFloatArrayElements	*[0]byte
	ReleaseDoubleArrayElements	*[0]byte
	GetBooleanArrayRegion		*[0]byte
	GetByteArrayRegion		*[0]byte
	GetCharArrayRegion		*[0]byte
	GetShortArrayRegion		*[0]byte
	GetIntArrayRegion		*[0]byte
	GetLongArrayRegion		*[0]byte
	GetFloatArrayRegion		*[0]byte
	GetDoubleArrayRegion		*[0]byte
	SetBooleanArrayRegion		*[0]byte
	SetByteArrayRegion		*[0]byte
	SetCharArrayRegion		*[0]byte
	SetShortArrayRegion		*[0]byte
	SetIntArrayRegion		*[0]byte
	SetLongArrayRegion		*[0]byte
	SetFloatArrayRegion		*[0]byte
	SetDoubleArrayRegion		*[0]byte
	RegisterNatives			*[0]byte
	UnregisterNatives		*[0]byte
	MonitorEnter			*[0]byte
	MonitorExit			*[0]byte
	GetJavaVM			*[0]byte
	GetStringRegion			*[0]byte
	GetStringUTFRegion		*[0]byte
	GetPrimitiveArrayCritical	*[0]byte
	ReleasePrimitiveArrayCritical	*[0]byte
	GetStringCritical		*[0]byte
	ReleaseStringCritical		*[0]byte
	NewWeakGlobalRef		*[0]byte
	DeleteWeakGlobalRef		*[0]byte
	ExceptionCheck			*[0]byte
	NewDirectByteBuffer		*[0]byte
	GetDirectBufferAddress		*[0]byte
	GetDirectBufferCapacity		*[0]byte
	GetObjectRefType		*[0]byte
}

type _Ctype_struct___0 struct {
	name		*_Ctype_char
	signature	*_Ctype_char
	fnPtr		unsafe.Pointer
}

type _Ctype_struct___va_list_tag struct {
	gp_offset		_Ctype_uint
	fp_offset		_Ctype_uint
	overflow_arg_area	unsafe.Pointer
	reg_save_area		unsafe.Pointer
}

type _Ctype_struct__jfieldID struct{}

type _Ctype_struct__jmethodID struct{}

type _Ctype_struct__jobject struct{}

type _Ctype_uchar uint8

type _Ctype_uint uint32

type _Ctype_union_jvalue [8]byte

type _Ctype_ushort uint16

type _Ctype_va_list _Ctype___builtin_va_list

type _Ctype_void [0]byte

//go:linkname _cgo_runtime_cgocall runtime.cgocall
func _cgo_runtime_cgocall(unsafe.Pointer, uintptr) int32

//go:linkname _cgo_runtime_cgocallback runtime.cgocallback
func _cgo_runtime_cgocallback(unsafe.Pointer, unsafe.Pointer, uintptr, uintptr)

//go:linkname _cgoCheckPointer runtime.cgoCheckPointer
func _cgoCheckPointer(interface{}, ...interface{})

//go:linkname _cgoCheckResult runtime.cgoCheckResult
func _cgoCheckResult(interface{})

//go:cgo_export_dynamic GNI_GetVersion
//go:linkname _cgoexp_cf0433f838ae_GNI_GetVersion _cgoexp_cf0433f838ae_GNI_GetVersion
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetVersion
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetVersion(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetVersion
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetVersion(p0 *_Ctype_JNIEnv) (r0 _Ctype_jint) {
	return GNI_GetVersion(p0)
}
//go:cgo_export_dynamic GNI_DefineClass
//go:linkname _cgoexp_cf0433f838ae_GNI_DefineClass _cgoexp_cf0433f838ae_GNI_DefineClass
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_DefineClass
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_DefineClass(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_DefineClass
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_DefineClass(p0 *_Ctype_JNIEnv, p1 *_Ctype_char, p2 _Ctype_jobject, p3 *_Ctype_jbyte, p4 _Ctype_jsize) (r0 _Ctype_jclass) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_DefineClass(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_FindClass
//go:linkname _cgoexp_cf0433f838ae_GNI_FindClass _cgoexp_cf0433f838ae_GNI_FindClass
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_FindClass
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_FindClass(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_FindClass
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_FindClass(p0 *_Ctype_JNIEnv, p1 *_Ctype_char) (r0 _Ctype_jclass) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_FindClass(p0, p1)
}
//go:cgo_export_dynamic GNI_FromReflectedMethod
//go:linkname _cgoexp_cf0433f838ae_GNI_FromReflectedMethod _cgoexp_cf0433f838ae_GNI_FromReflectedMethod
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_FromReflectedMethod
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_FromReflectedMethod(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_FromReflectedMethod
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_FromReflectedMethod(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject) (r0 _Ctype_jmethodID) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_FromReflectedMethod(p0, p1)
}
//go:cgo_export_dynamic GNI_FromReflectedField
//go:linkname _cgoexp_cf0433f838ae_GNI_FromReflectedField _cgoexp_cf0433f838ae_GNI_FromReflectedField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_FromReflectedField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_FromReflectedField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_FromReflectedField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_FromReflectedField(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject) (r0 _Ctype_jfieldID) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_FromReflectedField(p0, p1)
}
//go:cgo_export_dynamic GNI_ToReflectedMethod
//go:linkname _cgoexp_cf0433f838ae_GNI_ToReflectedMethod _cgoexp_cf0433f838ae_GNI_ToReflectedMethod
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_ToReflectedMethod
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_ToReflectedMethod(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_ToReflectedMethod
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_ToReflectedMethod(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jmethodID, p3 _Ctype_jboolean) (r0 _Ctype_jobject) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_ToReflectedMethod(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_GetSuperclass
//go:linkname _cgoexp_cf0433f838ae_GNI_GetSuperclass _cgoexp_cf0433f838ae_GNI_GetSuperclass
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetSuperclass
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetSuperclass(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetSuperclass
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetSuperclass(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass) (r0 _Ctype_jclass) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_GetSuperclass(p0, p1)
}
//go:cgo_export_dynamic GNI_IsAssignableFrom
//go:linkname _cgoexp_cf0433f838ae_GNI_IsAssignableFrom _cgoexp_cf0433f838ae_GNI_IsAssignableFrom
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_IsAssignableFrom
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_IsAssignableFrom(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_IsAssignableFrom
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_IsAssignableFrom(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jclass) (r0 _Ctype_jboolean) {
	return GNI_IsAssignableFrom(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_ToReflectedField
//go:linkname _cgoexp_cf0433f838ae_GNI_ToReflectedField _cgoexp_cf0433f838ae_GNI_ToReflectedField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_ToReflectedField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_ToReflectedField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_ToReflectedField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_ToReflectedField(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jfieldID, p3 _Ctype_jboolean) (r0 _Ctype_jobject) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_ToReflectedField(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_Throw
//go:linkname _cgoexp_cf0433f838ae_GNI_Throw _cgoexp_cf0433f838ae_GNI_Throw
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_Throw
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_Throw(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_Throw
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_Throw(p0 *_Ctype_JNIEnv, p1 _Ctype_jthrowable) (r0 _Ctype_jint) {
	return GNI_Throw(p0, p1)
}
//go:cgo_export_dynamic GNI_ThrowNew
//go:linkname _cgoexp_cf0433f838ae_GNI_ThrowNew _cgoexp_cf0433f838ae_GNI_ThrowNew
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_ThrowNew
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_ThrowNew(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_ThrowNew
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_ThrowNew(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 *_Ctype_char) (r0 _Ctype_jint) {
	return GNI_ThrowNew(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_ExceptionOccurred
//go:linkname _cgoexp_cf0433f838ae_GNI_ExceptionOccurred _cgoexp_cf0433f838ae_GNI_ExceptionOccurred
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_ExceptionOccurred
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_ExceptionOccurred(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_ExceptionOccurred
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_ExceptionOccurred(p0 *_Ctype_JNIEnv) (r0 _Ctype_jthrowable) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_ExceptionOccurred(p0)
}
//go:cgo_export_dynamic GNI_ExceptionDescribe
//go:linkname _cgoexp_cf0433f838ae_GNI_ExceptionDescribe _cgoexp_cf0433f838ae_GNI_ExceptionDescribe
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_ExceptionDescribe
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_ExceptionDescribe(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_ExceptionDescribe
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_ExceptionDescribe(p0 *_Ctype_JNIEnv) {
	GNI_ExceptionDescribe(p0)
}
//go:cgo_export_dynamic GNI_ExceptionClear
//go:linkname _cgoexp_cf0433f838ae_GNI_ExceptionClear _cgoexp_cf0433f838ae_GNI_ExceptionClear
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_ExceptionClear
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_ExceptionClear(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_ExceptionClear
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_ExceptionClear(p0 *_Ctype_JNIEnv) {
	GNI_ExceptionClear(p0)
}
//go:cgo_export_dynamic GNI_FatalError
//go:linkname _cgoexp_cf0433f838ae_GNI_FatalError _cgoexp_cf0433f838ae_GNI_FatalError
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_FatalError
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_FatalError(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_FatalError
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_FatalError(p0 *_Ctype_JNIEnv, p1 *_Ctype_char) {
	GNI_FatalError(p0, p1)
}
//go:cgo_export_dynamic GNI_PushLocalFrame
//go:linkname _cgoexp_cf0433f838ae_GNI_PushLocalFrame _cgoexp_cf0433f838ae_GNI_PushLocalFrame
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_PushLocalFrame
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_PushLocalFrame(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_PushLocalFrame
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_PushLocalFrame(p0 *_Ctype_JNIEnv, p1 _Ctype_jint) (r0 _Ctype_jint) {
	return GNI_PushLocalFrame(p0, p1)
}
//go:cgo_export_dynamic GNI_PopLocalFrame
//go:linkname _cgoexp_cf0433f838ae_GNI_PopLocalFrame _cgoexp_cf0433f838ae_GNI_PopLocalFrame
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_PopLocalFrame
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_PopLocalFrame(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_PopLocalFrame
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_PopLocalFrame(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject) (r0 _Ctype_jobject) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_PopLocalFrame(p0, p1)
}
//go:cgo_export_dynamic GNI_NewGlobalRef
//go:linkname _cgoexp_cf0433f838ae_GNI_NewGlobalRef _cgoexp_cf0433f838ae_GNI_NewGlobalRef
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_NewGlobalRef
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_NewGlobalRef(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_NewGlobalRef
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_NewGlobalRef(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject) (r0 _Ctype_jobject) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_NewGlobalRef(p0, p1)
}
//go:cgo_export_dynamic GNI_DeleteGlobalRef
//go:linkname _cgoexp_cf0433f838ae_GNI_DeleteGlobalRef _cgoexp_cf0433f838ae_GNI_DeleteGlobalRef
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_DeleteGlobalRef
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_DeleteGlobalRef(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_DeleteGlobalRef
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_DeleteGlobalRef(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject) {
	GNI_DeleteGlobalRef(p0, p1)
}
//go:cgo_export_dynamic GNI_DeleteLocalRef
//go:linkname _cgoexp_cf0433f838ae_GNI_DeleteLocalRef _cgoexp_cf0433f838ae_GNI_DeleteLocalRef
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_DeleteLocalRef
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_DeleteLocalRef(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_DeleteLocalRef
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_DeleteLocalRef(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject) {
	GNI_DeleteLocalRef(p0, p1)
}
//go:cgo_export_dynamic GNI_IsSameObject
//go:linkname _cgoexp_cf0433f838ae_GNI_IsSameObject _cgoexp_cf0433f838ae_GNI_IsSameObject
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_IsSameObject
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_IsSameObject(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_IsSameObject
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_IsSameObject(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jobject) (r0 _Ctype_jboolean) {
	return GNI_IsSameObject(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_NewLocalRef
//go:linkname _cgoexp_cf0433f838ae_GNI_NewLocalRef _cgoexp_cf0433f838ae_GNI_NewLocalRef
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_NewLocalRef
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_NewLocalRef(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_NewLocalRef
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_NewLocalRef(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject) (r0 _Ctype_jobject) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_NewLocalRef(p0, p1)
}
//go:cgo_export_dynamic GNI_EnsureLocalCapacity
//go:linkname _cgoexp_cf0433f838ae_GNI_EnsureLocalCapacity _cgoexp_cf0433f838ae_GNI_EnsureLocalCapacity
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_EnsureLocalCapacity
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_EnsureLocalCapacity(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_EnsureLocalCapacity
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_EnsureLocalCapacity(p0 *_Ctype_JNIEnv, p1 _Ctype_jint) (r0 _Ctype_jint) {
	return GNI_EnsureLocalCapacity(p0, p1)
}
//go:cgo_export_dynamic GNI_AllocObject
//go:linkname _cgoexp_cf0433f838ae_GNI_AllocObject _cgoexp_cf0433f838ae_GNI_AllocObject
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_AllocObject
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_AllocObject(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_AllocObject
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_AllocObject(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass) (r0 _Ctype_jobject) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_AllocObject(p0, p1)
}
//go:cgo_export_dynamic GNI_NewObjectV
//go:linkname _cgoexp_cf0433f838ae_GNI_NewObjectV _cgoexp_cf0433f838ae_GNI_NewObjectV
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_NewObjectV
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_NewObjectV(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_NewObjectV
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_NewObjectV(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jmethodID, p3 *_Ctype_va_list) (r0 _Ctype_jobject) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_NewObjectV(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_NewObjectA
//go:linkname _cgoexp_cf0433f838ae_GNI_NewObjectA _cgoexp_cf0433f838ae_GNI_NewObjectA
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_NewObjectA
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_NewObjectA(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_NewObjectA
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_NewObjectA(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jmethodID, p3 *_Ctype_jvalue) (r0 _Ctype_jobject) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_NewObjectA(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_GetObjectClass
//go:linkname _cgoexp_cf0433f838ae_GNI_GetObjectClass _cgoexp_cf0433f838ae_GNI_GetObjectClass
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetObjectClass
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetObjectClass(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetObjectClass
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetObjectClass(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject) (r0 _Ctype_jclass) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_GetObjectClass(p0, p1)
}
//go:cgo_export_dynamic GNI_IsInstanceOf
//go:linkname _cgoexp_cf0433f838ae_GNI_IsInstanceOf _cgoexp_cf0433f838ae_GNI_IsInstanceOf
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_IsInstanceOf
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_IsInstanceOf(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_IsInstanceOf
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_IsInstanceOf(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jclass) (r0 _Ctype_jboolean) {
	return GNI_IsInstanceOf(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_GetMethodID
//go:linkname _cgoexp_cf0433f838ae_GNI_GetMethodID _cgoexp_cf0433f838ae_GNI_GetMethodID
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetMethodID
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetMethodID(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetMethodID
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetMethodID(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 *_Ctype_char, p3 *_Ctype_char) (r0 _Ctype_jmethodID) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_GetMethodID(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallObjectMethodV
//go:linkname _cgoexp_cf0433f838ae_GNI_CallObjectMethodV _cgoexp_cf0433f838ae_GNI_CallObjectMethodV
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallObjectMethodV
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallObjectMethodV(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallObjectMethodV
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallObjectMethodV(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jmethodID, p3 *_Ctype_va_list) (r0 _Ctype_jobject) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_CallObjectMethodV(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallObjectMethodA
//go:linkname _cgoexp_cf0433f838ae_GNI_CallObjectMethodA _cgoexp_cf0433f838ae_GNI_CallObjectMethodA
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallObjectMethodA
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallObjectMethodA(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallObjectMethodA
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallObjectMethodA(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jmethodID, p3 *_Ctype_jvalue) (r0 _Ctype_jobject) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_CallObjectMethodA(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallBooleanMethodV
//go:linkname _cgoexp_cf0433f838ae_GNI_CallBooleanMethodV _cgoexp_cf0433f838ae_GNI_CallBooleanMethodV
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallBooleanMethodV
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallBooleanMethodV(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallBooleanMethodV
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallBooleanMethodV(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jmethodID, p3 *_Ctype_va_list) (r0 _Ctype_jboolean) {
	return GNI_CallBooleanMethodV(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallBooleanMethodA
//go:linkname _cgoexp_cf0433f838ae_GNI_CallBooleanMethodA _cgoexp_cf0433f838ae_GNI_CallBooleanMethodA
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallBooleanMethodA
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallBooleanMethodA(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallBooleanMethodA
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallBooleanMethodA(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jmethodID, p3 *_Ctype_jvalue) (r0 _Ctype_jboolean) {
	return GNI_CallBooleanMethodA(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallByteMethodV
//go:linkname _cgoexp_cf0433f838ae_GNI_CallByteMethodV _cgoexp_cf0433f838ae_GNI_CallByteMethodV
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallByteMethodV
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallByteMethodV(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallByteMethodV
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallByteMethodV(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jmethodID, p3 *_Ctype_va_list) (r0 _Ctype_jbyte) {
	return GNI_CallByteMethodV(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallByteMethodA
//go:linkname _cgoexp_cf0433f838ae_GNI_CallByteMethodA _cgoexp_cf0433f838ae_GNI_CallByteMethodA
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallByteMethodA
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallByteMethodA(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallByteMethodA
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallByteMethodA(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jmethodID, p3 *_Ctype_jvalue) (r0 _Ctype_jbyte) {
	return GNI_CallByteMethodA(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallCharMethodV
//go:linkname _cgoexp_cf0433f838ae_GNI_CallCharMethodV _cgoexp_cf0433f838ae_GNI_CallCharMethodV
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallCharMethodV
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallCharMethodV(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallCharMethodV
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallCharMethodV(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jmethodID, p3 *_Ctype_va_list) (r0 _Ctype_jchar) {
	return GNI_CallCharMethodV(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallCharMethodA
//go:linkname _cgoexp_cf0433f838ae_GNI_CallCharMethodA _cgoexp_cf0433f838ae_GNI_CallCharMethodA
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallCharMethodA
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallCharMethodA(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallCharMethodA
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallCharMethodA(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jmethodID, p3 *_Ctype_jvalue) (r0 _Ctype_jchar) {
	return GNI_CallCharMethodA(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallShortMethodV
//go:linkname _cgoexp_cf0433f838ae_GNI_CallShortMethodV _cgoexp_cf0433f838ae_GNI_CallShortMethodV
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallShortMethodV
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallShortMethodV(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallShortMethodV
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallShortMethodV(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jmethodID, p3 *_Ctype_va_list) (r0 _Ctype_jshort) {
	return GNI_CallShortMethodV(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallShortMethodA
//go:linkname _cgoexp_cf0433f838ae_GNI_CallShortMethodA _cgoexp_cf0433f838ae_GNI_CallShortMethodA
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallShortMethodA
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallShortMethodA(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallShortMethodA
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallShortMethodA(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jmethodID, p3 *_Ctype_jvalue) (r0 _Ctype_jshort) {
	return GNI_CallShortMethodA(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallIntMethodV
//go:linkname _cgoexp_cf0433f838ae_GNI_CallIntMethodV _cgoexp_cf0433f838ae_GNI_CallIntMethodV
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallIntMethodV
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallIntMethodV(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallIntMethodV
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallIntMethodV(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jmethodID, p3 *_Ctype_va_list) (r0 _Ctype_jint) {
	return GNI_CallIntMethodV(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallIntMethodA
//go:linkname _cgoexp_cf0433f838ae_GNI_CallIntMethodA _cgoexp_cf0433f838ae_GNI_CallIntMethodA
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallIntMethodA
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallIntMethodA(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallIntMethodA
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallIntMethodA(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jmethodID, p3 *_Ctype_jvalue) (r0 _Ctype_jint) {
	return GNI_CallIntMethodA(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallLongMethodV
//go:linkname _cgoexp_cf0433f838ae_GNI_CallLongMethodV _cgoexp_cf0433f838ae_GNI_CallLongMethodV
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallLongMethodV
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallLongMethodV(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallLongMethodV
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallLongMethodV(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jmethodID, p3 *_Ctype_va_list) (r0 _Ctype_jlong) {
	return GNI_CallLongMethodV(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallLongMethodA
//go:linkname _cgoexp_cf0433f838ae_GNI_CallLongMethodA _cgoexp_cf0433f838ae_GNI_CallLongMethodA
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallLongMethodA
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallLongMethodA(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallLongMethodA
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallLongMethodA(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jmethodID, p3 *_Ctype_jvalue) (r0 _Ctype_jlong) {
	return GNI_CallLongMethodA(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallFloatMethodV
//go:linkname _cgoexp_cf0433f838ae_GNI_CallFloatMethodV _cgoexp_cf0433f838ae_GNI_CallFloatMethodV
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallFloatMethodV
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallFloatMethodV(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallFloatMethodV
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallFloatMethodV(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jmethodID, p3 *_Ctype_va_list) (r0 _Ctype_jfloat) {
	return GNI_CallFloatMethodV(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallFloatMethodA
//go:linkname _cgoexp_cf0433f838ae_GNI_CallFloatMethodA _cgoexp_cf0433f838ae_GNI_CallFloatMethodA
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallFloatMethodA
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallFloatMethodA(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallFloatMethodA
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallFloatMethodA(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jmethodID, p3 *_Ctype_jvalue) (r0 _Ctype_jfloat) {
	return GNI_CallFloatMethodA(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallDoubleMethodV
//go:linkname _cgoexp_cf0433f838ae_GNI_CallDoubleMethodV _cgoexp_cf0433f838ae_GNI_CallDoubleMethodV
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallDoubleMethodV
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallDoubleMethodV(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallDoubleMethodV
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallDoubleMethodV(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jmethodID, p3 *_Ctype_va_list) (r0 _Ctype_jdouble) {
	return GNI_CallDoubleMethodV(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallDoubleMethodA
//go:linkname _cgoexp_cf0433f838ae_GNI_CallDoubleMethodA _cgoexp_cf0433f838ae_GNI_CallDoubleMethodA
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallDoubleMethodA
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallDoubleMethodA(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallDoubleMethodA
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallDoubleMethodA(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jmethodID, p3 *_Ctype_jvalue) (r0 _Ctype_jdouble) {
	return GNI_CallDoubleMethodA(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallVoidMethodV
//go:linkname _cgoexp_cf0433f838ae_GNI_CallVoidMethodV _cgoexp_cf0433f838ae_GNI_CallVoidMethodV
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallVoidMethodV
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallVoidMethodV(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallVoidMethodV
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallVoidMethodV(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jmethodID, p3 *_Ctype_va_list) {
	GNI_CallVoidMethodV(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallVoidMethodA
//go:linkname _cgoexp_cf0433f838ae_GNI_CallVoidMethodA _cgoexp_cf0433f838ae_GNI_CallVoidMethodA
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallVoidMethodA
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallVoidMethodA(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallVoidMethodA
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallVoidMethodA(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jmethodID, p3 *_Ctype_jvalue) {
	GNI_CallVoidMethodA(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallNonvirtualObjectMethodV
//go:linkname _cgoexp_cf0433f838ae_GNI_CallNonvirtualObjectMethodV _cgoexp_cf0433f838ae_GNI_CallNonvirtualObjectMethodV
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallNonvirtualObjectMethodV
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallNonvirtualObjectMethodV(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualObjectMethodV
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualObjectMethodV(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jclass, p3 _Ctype_jmethodID, p4 *_Ctype_va_list) (r0 _Ctype_jobject) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_CallNonvirtualObjectMethodV(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_CallNonvirtualObjectMethodA
//go:linkname _cgoexp_cf0433f838ae_GNI_CallNonvirtualObjectMethodA _cgoexp_cf0433f838ae_GNI_CallNonvirtualObjectMethodA
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallNonvirtualObjectMethodA
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallNonvirtualObjectMethodA(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualObjectMethodA
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualObjectMethodA(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jclass, p3 _Ctype_jmethodID, p4 *_Ctype_jvalue) (r0 _Ctype_jobject) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_CallNonvirtualObjectMethodA(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_CallNonvirtualBooleanMethodV
//go:linkname _cgoexp_cf0433f838ae_GNI_CallNonvirtualBooleanMethodV _cgoexp_cf0433f838ae_GNI_CallNonvirtualBooleanMethodV
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallNonvirtualBooleanMethodV
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallNonvirtualBooleanMethodV(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualBooleanMethodV
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualBooleanMethodV(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jclass, p3 _Ctype_jmethodID, p4 *_Ctype_va_list) (r0 _Ctype_jboolean) {
	return GNI_CallNonvirtualBooleanMethodV(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_CallNonvirtualBooleanMethodA
//go:linkname _cgoexp_cf0433f838ae_GNI_CallNonvirtualBooleanMethodA _cgoexp_cf0433f838ae_GNI_CallNonvirtualBooleanMethodA
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallNonvirtualBooleanMethodA
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallNonvirtualBooleanMethodA(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualBooleanMethodA
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualBooleanMethodA(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jclass, p3 _Ctype_jmethodID, p4 *_Ctype_jvalue) (r0 _Ctype_jboolean) {
	return GNI_CallNonvirtualBooleanMethodA(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_CallNonvirtualByteMethodV
//go:linkname _cgoexp_cf0433f838ae_GNI_CallNonvirtualByteMethodV _cgoexp_cf0433f838ae_GNI_CallNonvirtualByteMethodV
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallNonvirtualByteMethodV
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallNonvirtualByteMethodV(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualByteMethodV
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualByteMethodV(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jclass, p3 _Ctype_jmethodID, p4 *_Ctype_va_list) (r0 _Ctype_jbyte) {
	return GNI_CallNonvirtualByteMethodV(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_CallNonvirtualByteMethodA
//go:linkname _cgoexp_cf0433f838ae_GNI_CallNonvirtualByteMethodA _cgoexp_cf0433f838ae_GNI_CallNonvirtualByteMethodA
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallNonvirtualByteMethodA
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallNonvirtualByteMethodA(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualByteMethodA
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualByteMethodA(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jclass, p3 _Ctype_jmethodID, p4 *_Ctype_jvalue) (r0 _Ctype_jbyte) {
	return GNI_CallNonvirtualByteMethodA(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_CallNonvirtualCharMethodV
//go:linkname _cgoexp_cf0433f838ae_GNI_CallNonvirtualCharMethodV _cgoexp_cf0433f838ae_GNI_CallNonvirtualCharMethodV
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallNonvirtualCharMethodV
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallNonvirtualCharMethodV(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualCharMethodV
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualCharMethodV(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jclass, p3 _Ctype_jmethodID, p4 *_Ctype_va_list) (r0 _Ctype_jchar) {
	return GNI_CallNonvirtualCharMethodV(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_CallNonvirtualCharMethodA
//go:linkname _cgoexp_cf0433f838ae_GNI_CallNonvirtualCharMethodA _cgoexp_cf0433f838ae_GNI_CallNonvirtualCharMethodA
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallNonvirtualCharMethodA
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallNonvirtualCharMethodA(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualCharMethodA
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualCharMethodA(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jclass, p3 _Ctype_jmethodID, p4 *_Ctype_jvalue) (r0 _Ctype_jchar) {
	return GNI_CallNonvirtualCharMethodA(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_CallNonvirtualShortMethodV
//go:linkname _cgoexp_cf0433f838ae_GNI_CallNonvirtualShortMethodV _cgoexp_cf0433f838ae_GNI_CallNonvirtualShortMethodV
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallNonvirtualShortMethodV
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallNonvirtualShortMethodV(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualShortMethodV
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualShortMethodV(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jclass, p3 _Ctype_jmethodID, p4 *_Ctype_va_list) (r0 _Ctype_jshort) {
	return GNI_CallNonvirtualShortMethodV(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_CallNonvirtualShortMethodA
//go:linkname _cgoexp_cf0433f838ae_GNI_CallNonvirtualShortMethodA _cgoexp_cf0433f838ae_GNI_CallNonvirtualShortMethodA
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallNonvirtualShortMethodA
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallNonvirtualShortMethodA(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualShortMethodA
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualShortMethodA(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jclass, p3 _Ctype_jmethodID, p4 *_Ctype_jvalue) (r0 _Ctype_jshort) {
	return GNI_CallNonvirtualShortMethodA(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_CallNonvirtualIntMethodV
//go:linkname _cgoexp_cf0433f838ae_GNI_CallNonvirtualIntMethodV _cgoexp_cf0433f838ae_GNI_CallNonvirtualIntMethodV
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallNonvirtualIntMethodV
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallNonvirtualIntMethodV(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualIntMethodV
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualIntMethodV(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jclass, p3 _Ctype_jmethodID, p4 *_Ctype_va_list) (r0 _Ctype_jint) {
	return GNI_CallNonvirtualIntMethodV(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_CallNonvirtualIntMethodA
//go:linkname _cgoexp_cf0433f838ae_GNI_CallNonvirtualIntMethodA _cgoexp_cf0433f838ae_GNI_CallNonvirtualIntMethodA
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallNonvirtualIntMethodA
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallNonvirtualIntMethodA(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualIntMethodA
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualIntMethodA(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jclass, p3 _Ctype_jmethodID, p4 *_Ctype_jvalue) (r0 _Ctype_jint) {
	return GNI_CallNonvirtualIntMethodA(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_CallNonvirtualLongMethodV
//go:linkname _cgoexp_cf0433f838ae_GNI_CallNonvirtualLongMethodV _cgoexp_cf0433f838ae_GNI_CallNonvirtualLongMethodV
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallNonvirtualLongMethodV
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallNonvirtualLongMethodV(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualLongMethodV
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualLongMethodV(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jclass, p3 _Ctype_jmethodID, p4 *_Ctype_va_list) (r0 _Ctype_jlong) {
	return GNI_CallNonvirtualLongMethodV(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_CallNonvirtualLongMethodA
//go:linkname _cgoexp_cf0433f838ae_GNI_CallNonvirtualLongMethodA _cgoexp_cf0433f838ae_GNI_CallNonvirtualLongMethodA
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallNonvirtualLongMethodA
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallNonvirtualLongMethodA(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualLongMethodA
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualLongMethodA(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jclass, p3 _Ctype_jmethodID, p4 *_Ctype_jvalue) (r0 _Ctype_jlong) {
	return GNI_CallNonvirtualLongMethodA(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_CallNonvirtualFloatMethodV
//go:linkname _cgoexp_cf0433f838ae_GNI_CallNonvirtualFloatMethodV _cgoexp_cf0433f838ae_GNI_CallNonvirtualFloatMethodV
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallNonvirtualFloatMethodV
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallNonvirtualFloatMethodV(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualFloatMethodV
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualFloatMethodV(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jclass, p3 _Ctype_jmethodID, p4 *_Ctype_va_list) (r0 _Ctype_jfloat) {
	return GNI_CallNonvirtualFloatMethodV(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_CallNonvirtualFloatMethodA
//go:linkname _cgoexp_cf0433f838ae_GNI_CallNonvirtualFloatMethodA _cgoexp_cf0433f838ae_GNI_CallNonvirtualFloatMethodA
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallNonvirtualFloatMethodA
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallNonvirtualFloatMethodA(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualFloatMethodA
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualFloatMethodA(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jclass, p3 _Ctype_jmethodID, p4 *_Ctype_jvalue) (r0 _Ctype_jfloat) {
	return GNI_CallNonvirtualFloatMethodA(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_CallNonvirtualDoubleMethodV
//go:linkname _cgoexp_cf0433f838ae_GNI_CallNonvirtualDoubleMethodV _cgoexp_cf0433f838ae_GNI_CallNonvirtualDoubleMethodV
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallNonvirtualDoubleMethodV
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallNonvirtualDoubleMethodV(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualDoubleMethodV
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualDoubleMethodV(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jclass, p3 _Ctype_jmethodID, p4 *_Ctype_va_list) (r0 _Ctype_jdouble) {
	return GNI_CallNonvirtualDoubleMethodV(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_CallNonvirtualDoubleMethodA
//go:linkname _cgoexp_cf0433f838ae_GNI_CallNonvirtualDoubleMethodA _cgoexp_cf0433f838ae_GNI_CallNonvirtualDoubleMethodA
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallNonvirtualDoubleMethodA
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallNonvirtualDoubleMethodA(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualDoubleMethodA
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualDoubleMethodA(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jclass, p3 _Ctype_jmethodID, p4 *_Ctype_jvalue) (r0 _Ctype_jdouble) {
	return GNI_CallNonvirtualDoubleMethodA(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_CallNonvirtualVoidMethodV
//go:linkname _cgoexp_cf0433f838ae_GNI_CallNonvirtualVoidMethodV _cgoexp_cf0433f838ae_GNI_CallNonvirtualVoidMethodV
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallNonvirtualVoidMethodV
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallNonvirtualVoidMethodV(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualVoidMethodV
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualVoidMethodV(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jclass, p3 _Ctype_jmethodID, p4 *_Ctype_va_list) {
	GNI_CallNonvirtualVoidMethodV(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_CallNonvirtualVoidMethodA
//go:linkname _cgoexp_cf0433f838ae_GNI_CallNonvirtualVoidMethodA _cgoexp_cf0433f838ae_GNI_CallNonvirtualVoidMethodA
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallNonvirtualVoidMethodA
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallNonvirtualVoidMethodA(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualVoidMethodA
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallNonvirtualVoidMethodA(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jclass, p3 _Ctype_jmethodID, p4 *_Ctype_jvalue) {
	GNI_CallNonvirtualVoidMethodA(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_GetFieldID
//go:linkname _cgoexp_cf0433f838ae_GNI_GetFieldID _cgoexp_cf0433f838ae_GNI_GetFieldID
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetFieldID
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetFieldID(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetFieldID
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetFieldID(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 *_Ctype_char, p3 *_Ctype_char) (r0 _Ctype_jfieldID) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_GetFieldID(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_GetObjectField
//go:linkname _cgoexp_cf0433f838ae_GNI_GetObjectField _cgoexp_cf0433f838ae_GNI_GetObjectField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetObjectField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetObjectField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetObjectField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetObjectField(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jfieldID) (r0 _Ctype_jobject) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_GetObjectField(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_GetBooleanField
//go:linkname _cgoexp_cf0433f838ae_GNI_GetBooleanField _cgoexp_cf0433f838ae_GNI_GetBooleanField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetBooleanField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetBooleanField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetBooleanField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetBooleanField(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jfieldID) (r0 _Ctype_jboolean) {
	return GNI_GetBooleanField(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_GetByteField
//go:linkname _cgoexp_cf0433f838ae_GNI_GetByteField _cgoexp_cf0433f838ae_GNI_GetByteField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetByteField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetByteField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetByteField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetByteField(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jfieldID) (r0 _Ctype_jbyte) {
	return GNI_GetByteField(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_GetCharField
//go:linkname _cgoexp_cf0433f838ae_GNI_GetCharField _cgoexp_cf0433f838ae_GNI_GetCharField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetCharField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetCharField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetCharField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetCharField(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jfieldID) (r0 _Ctype_jchar) {
	return GNI_GetCharField(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_GetShortField
//go:linkname _cgoexp_cf0433f838ae_GNI_GetShortField _cgoexp_cf0433f838ae_GNI_GetShortField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetShortField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetShortField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetShortField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetShortField(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jfieldID) (r0 _Ctype_jshort) {
	return GNI_GetShortField(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_GetIntField
//go:linkname _cgoexp_cf0433f838ae_GNI_GetIntField _cgoexp_cf0433f838ae_GNI_GetIntField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetIntField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetIntField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetIntField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetIntField(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jfieldID) (r0 _Ctype_jint) {
	return GNI_GetIntField(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_GetLongField
//go:linkname _cgoexp_cf0433f838ae_GNI_GetLongField _cgoexp_cf0433f838ae_GNI_GetLongField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetLongField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetLongField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetLongField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetLongField(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jfieldID) (r0 _Ctype_jlong) {
	return GNI_GetLongField(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_GetFloatField
//go:linkname _cgoexp_cf0433f838ae_GNI_GetFloatField _cgoexp_cf0433f838ae_GNI_GetFloatField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetFloatField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetFloatField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetFloatField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetFloatField(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jfieldID) (r0 _Ctype_jfloat) {
	return GNI_GetFloatField(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_GetDoubleField
//go:linkname _cgoexp_cf0433f838ae_GNI_GetDoubleField _cgoexp_cf0433f838ae_GNI_GetDoubleField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetDoubleField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetDoubleField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetDoubleField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetDoubleField(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jfieldID) (r0 _Ctype_jdouble) {
	return GNI_GetDoubleField(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_SetObjectField
//go:linkname _cgoexp_cf0433f838ae_GNI_SetObjectField _cgoexp_cf0433f838ae_GNI_SetObjectField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_SetObjectField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_SetObjectField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_SetObjectField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_SetObjectField(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jfieldID, p3 _Ctype_jobject) {
	GNI_SetObjectField(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_SetBooleanField
//go:linkname _cgoexp_cf0433f838ae_GNI_SetBooleanField _cgoexp_cf0433f838ae_GNI_SetBooleanField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_SetBooleanField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_SetBooleanField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_SetBooleanField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_SetBooleanField(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jfieldID, p3 _Ctype_jboolean) {
	GNI_SetBooleanField(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_SetByteField
//go:linkname _cgoexp_cf0433f838ae_GNI_SetByteField _cgoexp_cf0433f838ae_GNI_SetByteField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_SetByteField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_SetByteField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_SetByteField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_SetByteField(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jfieldID, p3 _Ctype_jbyte) {
	GNI_SetByteField(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_SetCharField
//go:linkname _cgoexp_cf0433f838ae_GNI_SetCharField _cgoexp_cf0433f838ae_GNI_SetCharField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_SetCharField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_SetCharField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_SetCharField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_SetCharField(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jfieldID, p3 _Ctype_jchar) {
	GNI_SetCharField(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_SetShortField
//go:linkname _cgoexp_cf0433f838ae_GNI_SetShortField _cgoexp_cf0433f838ae_GNI_SetShortField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_SetShortField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_SetShortField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_SetShortField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_SetShortField(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jfieldID, p3 _Ctype_jshort) {
	GNI_SetShortField(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_SetIntField
//go:linkname _cgoexp_cf0433f838ae_GNI_SetIntField _cgoexp_cf0433f838ae_GNI_SetIntField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_SetIntField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_SetIntField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_SetIntField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_SetIntField(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jfieldID, p3 _Ctype_jint) {
	GNI_SetIntField(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_SetLongField
//go:linkname _cgoexp_cf0433f838ae_GNI_SetLongField _cgoexp_cf0433f838ae_GNI_SetLongField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_SetLongField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_SetLongField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_SetLongField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_SetLongField(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jfieldID, p3 _Ctype_jlong) {
	GNI_SetLongField(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_SetFloatField
//go:linkname _cgoexp_cf0433f838ae_GNI_SetFloatField _cgoexp_cf0433f838ae_GNI_SetFloatField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_SetFloatField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_SetFloatField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_SetFloatField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_SetFloatField(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jfieldID, p3 _Ctype_jfloat) {
	GNI_SetFloatField(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_SetDoubleField
//go:linkname _cgoexp_cf0433f838ae_GNI_SetDoubleField _cgoexp_cf0433f838ae_GNI_SetDoubleField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_SetDoubleField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_SetDoubleField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_SetDoubleField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_SetDoubleField(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject, p2 _Ctype_jfieldID, p3 _Ctype_jdouble) {
	GNI_SetDoubleField(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_GetStaticMethodID
//go:linkname _cgoexp_cf0433f838ae_GNI_GetStaticMethodID _cgoexp_cf0433f838ae_GNI_GetStaticMethodID
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetStaticMethodID
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetStaticMethodID(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetStaticMethodID
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetStaticMethodID(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 *_Ctype_char, p3 *_Ctype_char) (r0 _Ctype_jmethodID) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_GetStaticMethodID(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallStaticObjectMethodV
//go:linkname _cgoexp_cf0433f838ae_GNI_CallStaticObjectMethodV _cgoexp_cf0433f838ae_GNI_CallStaticObjectMethodV
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallStaticObjectMethodV
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallStaticObjectMethodV(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallStaticObjectMethodV
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallStaticObjectMethodV(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jmethodID, p3 *_Ctype_va_list) (r0 _Ctype_jobject) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_CallStaticObjectMethodV(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallStaticObjectMethodA
//go:linkname _cgoexp_cf0433f838ae_GNI_CallStaticObjectMethodA _cgoexp_cf0433f838ae_GNI_CallStaticObjectMethodA
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallStaticObjectMethodA
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallStaticObjectMethodA(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallStaticObjectMethodA
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallStaticObjectMethodA(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jmethodID, p3 *_Ctype_jvalue) (r0 _Ctype_jobject) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_CallStaticObjectMethodA(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallStaticBooleanMethodV
//go:linkname _cgoexp_cf0433f838ae_GNI_CallStaticBooleanMethodV _cgoexp_cf0433f838ae_GNI_CallStaticBooleanMethodV
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallStaticBooleanMethodV
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallStaticBooleanMethodV(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallStaticBooleanMethodV
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallStaticBooleanMethodV(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jmethodID, p3 *_Ctype_va_list) (r0 _Ctype_jboolean) {
	return GNI_CallStaticBooleanMethodV(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallStaticBooleanMethodA
//go:linkname _cgoexp_cf0433f838ae_GNI_CallStaticBooleanMethodA _cgoexp_cf0433f838ae_GNI_CallStaticBooleanMethodA
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallStaticBooleanMethodA
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallStaticBooleanMethodA(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallStaticBooleanMethodA
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallStaticBooleanMethodA(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jmethodID, p3 *_Ctype_jvalue) (r0 _Ctype_jboolean) {
	return GNI_CallStaticBooleanMethodA(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallStaticByteMethodV
//go:linkname _cgoexp_cf0433f838ae_GNI_CallStaticByteMethodV _cgoexp_cf0433f838ae_GNI_CallStaticByteMethodV
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallStaticByteMethodV
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallStaticByteMethodV(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallStaticByteMethodV
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallStaticByteMethodV(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jmethodID, p3 *_Ctype_va_list) (r0 _Ctype_jbyte) {
	return GNI_CallStaticByteMethodV(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallStaticByteMethodA
//go:linkname _cgoexp_cf0433f838ae_GNI_CallStaticByteMethodA _cgoexp_cf0433f838ae_GNI_CallStaticByteMethodA
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallStaticByteMethodA
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallStaticByteMethodA(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallStaticByteMethodA
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallStaticByteMethodA(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jmethodID, p3 *_Ctype_jvalue) (r0 _Ctype_jbyte) {
	return GNI_CallStaticByteMethodA(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallStaticCharMethodV
//go:linkname _cgoexp_cf0433f838ae_GNI_CallStaticCharMethodV _cgoexp_cf0433f838ae_GNI_CallStaticCharMethodV
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallStaticCharMethodV
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallStaticCharMethodV(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallStaticCharMethodV
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallStaticCharMethodV(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jmethodID, p3 *_Ctype_va_list) (r0 _Ctype_jchar) {
	return GNI_CallStaticCharMethodV(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallStaticCharMethodA
//go:linkname _cgoexp_cf0433f838ae_GNI_CallStaticCharMethodA _cgoexp_cf0433f838ae_GNI_CallStaticCharMethodA
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallStaticCharMethodA
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallStaticCharMethodA(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallStaticCharMethodA
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallStaticCharMethodA(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jmethodID, p3 *_Ctype_jvalue) (r0 _Ctype_jchar) {
	return GNI_CallStaticCharMethodA(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallStaticShortMethodA
//go:linkname _cgoexp_cf0433f838ae_GNI_CallStaticShortMethodA _cgoexp_cf0433f838ae_GNI_CallStaticShortMethodA
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallStaticShortMethodA
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallStaticShortMethodA(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallStaticShortMethodA
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallStaticShortMethodA(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jmethodID, p3 *_Ctype_jvalue) (r0 _Ctype_jshort) {
	return GNI_CallStaticShortMethodA(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallStaticIntMethodV
//go:linkname _cgoexp_cf0433f838ae_GNI_CallStaticIntMethodV _cgoexp_cf0433f838ae_GNI_CallStaticIntMethodV
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallStaticIntMethodV
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallStaticIntMethodV(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallStaticIntMethodV
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallStaticIntMethodV(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jmethodID, p3 *_Ctype_va_list) (r0 _Ctype_jint) {
	return GNI_CallStaticIntMethodV(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallStaticIntMethodA
//go:linkname _cgoexp_cf0433f838ae_GNI_CallStaticIntMethodA _cgoexp_cf0433f838ae_GNI_CallStaticIntMethodA
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallStaticIntMethodA
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallStaticIntMethodA(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallStaticIntMethodA
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallStaticIntMethodA(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jmethodID, p3 *_Ctype_jvalue) (r0 _Ctype_jint) {
	return GNI_CallStaticIntMethodA(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallStaticLongMethodV
//go:linkname _cgoexp_cf0433f838ae_GNI_CallStaticLongMethodV _cgoexp_cf0433f838ae_GNI_CallStaticLongMethodV
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallStaticLongMethodV
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallStaticLongMethodV(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallStaticLongMethodV
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallStaticLongMethodV(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jmethodID, p3 *_Ctype_va_list) (r0 _Ctype_jlong) {
	return GNI_CallStaticLongMethodV(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallStaticLongMethodA
//go:linkname _cgoexp_cf0433f838ae_GNI_CallStaticLongMethodA _cgoexp_cf0433f838ae_GNI_CallStaticLongMethodA
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallStaticLongMethodA
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallStaticLongMethodA(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallStaticLongMethodA
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallStaticLongMethodA(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jmethodID, p3 *_Ctype_jvalue) (r0 _Ctype_jlong) {
	return GNI_CallStaticLongMethodA(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallStaticFloatMethodV
//go:linkname _cgoexp_cf0433f838ae_GNI_CallStaticFloatMethodV _cgoexp_cf0433f838ae_GNI_CallStaticFloatMethodV
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallStaticFloatMethodV
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallStaticFloatMethodV(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallStaticFloatMethodV
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallStaticFloatMethodV(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jmethodID, p3 *_Ctype_va_list) (r0 _Ctype_jfloat) {
	return GNI_CallStaticFloatMethodV(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallStaticFloatMethodA
//go:linkname _cgoexp_cf0433f838ae_GNI_CallStaticFloatMethodA _cgoexp_cf0433f838ae_GNI_CallStaticFloatMethodA
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallStaticFloatMethodA
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallStaticFloatMethodA(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallStaticFloatMethodA
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallStaticFloatMethodA(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jmethodID, p3 *_Ctype_jvalue) (r0 _Ctype_jfloat) {
	return GNI_CallStaticFloatMethodA(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallStaticDoubleMethodV
//go:linkname _cgoexp_cf0433f838ae_GNI_CallStaticDoubleMethodV _cgoexp_cf0433f838ae_GNI_CallStaticDoubleMethodV
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallStaticDoubleMethodV
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallStaticDoubleMethodV(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallStaticDoubleMethodV
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallStaticDoubleMethodV(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jmethodID, p3 *_Ctype_va_list) (r0 _Ctype_jdouble) {
	return GNI_CallStaticDoubleMethodV(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallStaticDoubleMethodA
//go:linkname _cgoexp_cf0433f838ae_GNI_CallStaticDoubleMethodA _cgoexp_cf0433f838ae_GNI_CallStaticDoubleMethodA
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallStaticDoubleMethodA
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallStaticDoubleMethodA(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallStaticDoubleMethodA
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallStaticDoubleMethodA(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jmethodID, p3 *_Ctype_jvalue) (r0 _Ctype_jdouble) {
	return GNI_CallStaticDoubleMethodA(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallStaticVoidMethodV
//go:linkname _cgoexp_cf0433f838ae_GNI_CallStaticVoidMethodV _cgoexp_cf0433f838ae_GNI_CallStaticVoidMethodV
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallStaticVoidMethodV
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallStaticVoidMethodV(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallStaticVoidMethodV
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallStaticVoidMethodV(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jmethodID, p3 *_Ctype_va_list) {
	GNI_CallStaticVoidMethodV(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_CallStaticVoidMethodA
//go:linkname _cgoexp_cf0433f838ae_GNI_CallStaticVoidMethodA _cgoexp_cf0433f838ae_GNI_CallStaticVoidMethodA
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_CallStaticVoidMethodA
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_CallStaticVoidMethodA(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_CallStaticVoidMethodA
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_CallStaticVoidMethodA(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jmethodID, p3 *_Ctype_jvalue) {
	GNI_CallStaticVoidMethodA(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_GetStaticFieldID
//go:linkname _cgoexp_cf0433f838ae_GNI_GetStaticFieldID _cgoexp_cf0433f838ae_GNI_GetStaticFieldID
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetStaticFieldID
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetStaticFieldID(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetStaticFieldID
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetStaticFieldID(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 *_Ctype_char, p3 *_Ctype_char) (r0 _Ctype_jfieldID) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_GetStaticFieldID(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_GetStaticObjectField
//go:linkname _cgoexp_cf0433f838ae_GNI_GetStaticObjectField _cgoexp_cf0433f838ae_GNI_GetStaticObjectField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetStaticObjectField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetStaticObjectField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetStaticObjectField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetStaticObjectField(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jfieldID) (r0 _Ctype_jobject) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_GetStaticObjectField(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_GetStaticBooleanField
//go:linkname _cgoexp_cf0433f838ae_GNI_GetStaticBooleanField _cgoexp_cf0433f838ae_GNI_GetStaticBooleanField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetStaticBooleanField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetStaticBooleanField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetStaticBooleanField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetStaticBooleanField(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jfieldID) (r0 _Ctype_jboolean) {
	return GNI_GetStaticBooleanField(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_GetStaticByteField
//go:linkname _cgoexp_cf0433f838ae_GNI_GetStaticByteField _cgoexp_cf0433f838ae_GNI_GetStaticByteField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetStaticByteField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetStaticByteField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetStaticByteField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetStaticByteField(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jfieldID) (r0 _Ctype_jbyte) {
	return GNI_GetStaticByteField(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_GetStaticCharField
//go:linkname _cgoexp_cf0433f838ae_GNI_GetStaticCharField _cgoexp_cf0433f838ae_GNI_GetStaticCharField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetStaticCharField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetStaticCharField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetStaticCharField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetStaticCharField(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jfieldID) (r0 _Ctype_jchar) {
	return GNI_GetStaticCharField(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_GetStaticShortField
//go:linkname _cgoexp_cf0433f838ae_GNI_GetStaticShortField _cgoexp_cf0433f838ae_GNI_GetStaticShortField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetStaticShortField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetStaticShortField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetStaticShortField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetStaticShortField(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jfieldID) (r0 _Ctype_jshort) {
	return GNI_GetStaticShortField(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_GetStaticIntField
//go:linkname _cgoexp_cf0433f838ae_GNI_GetStaticIntField _cgoexp_cf0433f838ae_GNI_GetStaticIntField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetStaticIntField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetStaticIntField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetStaticIntField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetStaticIntField(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jfieldID) (r0 _Ctype_jint) {
	return GNI_GetStaticIntField(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_GetStaticLongField
//go:linkname _cgoexp_cf0433f838ae_GNI_GetStaticLongField _cgoexp_cf0433f838ae_GNI_GetStaticLongField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetStaticLongField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetStaticLongField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetStaticLongField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetStaticLongField(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jfieldID) (r0 _Ctype_jlong) {
	return GNI_GetStaticLongField(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_GetStaticFloatField
//go:linkname _cgoexp_cf0433f838ae_GNI_GetStaticFloatField _cgoexp_cf0433f838ae_GNI_GetStaticFloatField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetStaticFloatField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetStaticFloatField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetStaticFloatField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetStaticFloatField(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jfieldID) (r0 _Ctype_jfloat) {
	return GNI_GetStaticFloatField(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_GetStaticDoubleField
//go:linkname _cgoexp_cf0433f838ae_GNI_GetStaticDoubleField _cgoexp_cf0433f838ae_GNI_GetStaticDoubleField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetStaticDoubleField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetStaticDoubleField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetStaticDoubleField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetStaticDoubleField(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jfieldID) (r0 _Ctype_jdouble) {
	return GNI_GetStaticDoubleField(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_SetStaticObjectField
//go:linkname _cgoexp_cf0433f838ae_GNI_SetStaticObjectField _cgoexp_cf0433f838ae_GNI_SetStaticObjectField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_SetStaticObjectField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_SetStaticObjectField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_SetStaticObjectField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_SetStaticObjectField(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jfieldID, p3 _Ctype_jobject) {
	GNI_SetStaticObjectField(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_SetStaticBooleanField
//go:linkname _cgoexp_cf0433f838ae_GNI_SetStaticBooleanField _cgoexp_cf0433f838ae_GNI_SetStaticBooleanField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_SetStaticBooleanField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_SetStaticBooleanField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_SetStaticBooleanField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_SetStaticBooleanField(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jfieldID, p3 _Ctype_jboolean) {
	GNI_SetStaticBooleanField(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_SetStaticByteField
//go:linkname _cgoexp_cf0433f838ae_GNI_SetStaticByteField _cgoexp_cf0433f838ae_GNI_SetStaticByteField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_SetStaticByteField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_SetStaticByteField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_SetStaticByteField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_SetStaticByteField(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jfieldID, p3 _Ctype_jbyte) {
	GNI_SetStaticByteField(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_SetStaticCharField
//go:linkname _cgoexp_cf0433f838ae_GNI_SetStaticCharField _cgoexp_cf0433f838ae_GNI_SetStaticCharField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_SetStaticCharField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_SetStaticCharField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_SetStaticCharField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_SetStaticCharField(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jfieldID, p3 _Ctype_jchar) {
	GNI_SetStaticCharField(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_SetStaticShortField
//go:linkname _cgoexp_cf0433f838ae_GNI_SetStaticShortField _cgoexp_cf0433f838ae_GNI_SetStaticShortField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_SetStaticShortField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_SetStaticShortField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_SetStaticShortField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_SetStaticShortField(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jfieldID, p3 _Ctype_jshort) {
	GNI_SetStaticShortField(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_SetStaticIntField
//go:linkname _cgoexp_cf0433f838ae_GNI_SetStaticIntField _cgoexp_cf0433f838ae_GNI_SetStaticIntField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_SetStaticIntField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_SetStaticIntField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_SetStaticIntField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_SetStaticIntField(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jfieldID, p3 _Ctype_jint) {
	GNI_SetStaticIntField(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_SetStaticLongField
//go:linkname _cgoexp_cf0433f838ae_GNI_SetStaticLongField _cgoexp_cf0433f838ae_GNI_SetStaticLongField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_SetStaticLongField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_SetStaticLongField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_SetStaticLongField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_SetStaticLongField(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jfieldID, p3 _Ctype_jlong) {
	GNI_SetStaticLongField(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_SetStaticFloatField
//go:linkname _cgoexp_cf0433f838ae_GNI_SetStaticFloatField _cgoexp_cf0433f838ae_GNI_SetStaticFloatField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_SetStaticFloatField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_SetStaticFloatField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_SetStaticFloatField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_SetStaticFloatField(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jfieldID, p3 _Ctype_jfloat) {
	GNI_SetStaticFloatField(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_SetStaticDoubleField
//go:linkname _cgoexp_cf0433f838ae_GNI_SetStaticDoubleField _cgoexp_cf0433f838ae_GNI_SetStaticDoubleField
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_SetStaticDoubleField
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_SetStaticDoubleField(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_SetStaticDoubleField
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_SetStaticDoubleField(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 _Ctype_jfieldID, p3 _Ctype_jdouble) {
	GNI_SetStaticDoubleField(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_NewString
//go:linkname _cgoexp_cf0433f838ae_GNI_NewString _cgoexp_cf0433f838ae_GNI_NewString
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_NewString
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_NewString(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_NewString
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_NewString(p0 *_Ctype_JNIEnv, p1 *_Ctype_jchar, p2 _Ctype_jsize) (r0 _Ctype_jstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_NewString(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_GetStringLength
//go:linkname _cgoexp_cf0433f838ae_GNI_GetStringLength _cgoexp_cf0433f838ae_GNI_GetStringLength
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetStringLength
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetStringLength(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetStringLength
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetStringLength(p0 *_Ctype_JNIEnv, p1 _Ctype_jstring) (r0 _Ctype_jsize) {
	return GNI_GetStringLength(p0, p1)
}
//go:cgo_export_dynamic GNI_GetStringChars
//go:linkname _cgoexp_cf0433f838ae_GNI_GetStringChars _cgoexp_cf0433f838ae_GNI_GetStringChars
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetStringChars
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetStringChars(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetStringChars
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetStringChars(p0 *_Ctype_JNIEnv, p1 _Ctype_jstring, p2 *_Ctype_jboolean) (r0 *_Ctype_jchar) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_GetStringChars(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_ReleaseStringChars
//go:linkname _cgoexp_cf0433f838ae_GNI_ReleaseStringChars _cgoexp_cf0433f838ae_GNI_ReleaseStringChars
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_ReleaseStringChars
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_ReleaseStringChars(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_ReleaseStringChars
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_ReleaseStringChars(p0 *_Ctype_JNIEnv, p1 _Ctype_jstring, p2 _Ctype_jchar) {
	GNI_ReleaseStringChars(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_NewStringUTF
//go:linkname _cgoexp_cf0433f838ae_GNI_NewStringUTF _cgoexp_cf0433f838ae_GNI_NewStringUTF
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_NewStringUTF
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_NewStringUTF(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_NewStringUTF
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_NewStringUTF(p0 *_Ctype_JNIEnv, p1 *_Ctype_char) (r0 _Ctype_jstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_NewStringUTF(p0, p1)
}
//go:cgo_export_dynamic GNI_GetStringUTFLength
//go:linkname _cgoexp_cf0433f838ae_GNI_GetStringUTFLength _cgoexp_cf0433f838ae_GNI_GetStringUTFLength
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetStringUTFLength
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetStringUTFLength(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetStringUTFLength
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetStringUTFLength(p0 *_Ctype_JNIEnv, p1 _Ctype_jstring) (r0 _Ctype_jsize) {
	return GNI_GetStringUTFLength(p0, p1)
}
//go:cgo_export_dynamic GNI_GetStringUTFChars
//go:linkname _cgoexp_cf0433f838ae_GNI_GetStringUTFChars _cgoexp_cf0433f838ae_GNI_GetStringUTFChars
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetStringUTFChars
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetStringUTFChars(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetStringUTFChars
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetStringUTFChars(p0 *_Ctype_JNIEnv, p1 _Ctype_jstring, p2 *_Ctype_jboolean) (r0 *_Ctype_char) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_GetStringUTFChars(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_ReleaseStringUTFChars
//go:linkname _cgoexp_cf0433f838ae_GNI_ReleaseStringUTFChars _cgoexp_cf0433f838ae_GNI_ReleaseStringUTFChars
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_ReleaseStringUTFChars
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_ReleaseStringUTFChars(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_ReleaseStringUTFChars
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_ReleaseStringUTFChars(p0 *_Ctype_JNIEnv, p1 _Ctype_jstring, p2 *_Ctype_char) {
	GNI_ReleaseStringUTFChars(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_GetArrayLength
//go:linkname _cgoexp_cf0433f838ae_GNI_GetArrayLength _cgoexp_cf0433f838ae_GNI_GetArrayLength
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetArrayLength
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetArrayLength(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetArrayLength
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetArrayLength(p0 *_Ctype_JNIEnv, p1 *_Ctype_jarray) (r0 _Ctype_jsize) {
	return GNI_GetArrayLength(p0, p1)
}
//go:cgo_export_dynamic GNI_NewObjectArray
//go:linkname _cgoexp_cf0433f838ae_GNI_NewObjectArray _cgoexp_cf0433f838ae_GNI_NewObjectArray
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_NewObjectArray
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_NewObjectArray(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_NewObjectArray
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_NewObjectArray(p0 *_Ctype_JNIEnv, p1 _Ctype_jsize, p2 _Ctype_jclass, p3 _Ctype_jobject) (r0 _Ctype_jobjectArray) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_NewObjectArray(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_GetObjectArrayElement
//go:linkname _cgoexp_cf0433f838ae_GNI_GetObjectArrayElement _cgoexp_cf0433f838ae_GNI_GetObjectArrayElement
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetObjectArrayElement
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetObjectArrayElement(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetObjectArrayElement
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetObjectArrayElement(p0 *_Ctype_JNIEnv, p1 _Ctype_jobjectArray, p2 _Ctype_jsize) (r0 _Ctype_jobject) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_GetObjectArrayElement(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_SetObjectArrayElement
//go:linkname _cgoexp_cf0433f838ae_GNI_SetObjectArrayElement _cgoexp_cf0433f838ae_GNI_SetObjectArrayElement
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_SetObjectArrayElement
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_SetObjectArrayElement(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_SetObjectArrayElement
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_SetObjectArrayElement(p0 *_Ctype_JNIEnv, p1 _Ctype_jobjectArray, p2 _Ctype_jsize, p3 _Ctype_jobject) {
	GNI_SetObjectArrayElement(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_NewBooleanArray
//go:linkname _cgoexp_cf0433f838ae_GNI_NewBooleanArray _cgoexp_cf0433f838ae_GNI_NewBooleanArray
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_NewBooleanArray
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_NewBooleanArray(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_NewBooleanArray
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_NewBooleanArray(p0 *_Ctype_JNIEnv, p1 _Ctype_jsize) (r0 _Ctype_jbooleanArray) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_NewBooleanArray(p0, p1)
}
//go:cgo_export_dynamic GNI_NewByteArray
//go:linkname _cgoexp_cf0433f838ae_GNI_NewByteArray _cgoexp_cf0433f838ae_GNI_NewByteArray
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_NewByteArray
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_NewByteArray(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_NewByteArray
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_NewByteArray(p0 *_Ctype_JNIEnv, p1 _Ctype_jsize) (r0 _Ctype_jbyteArray) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_NewByteArray(p0, p1)
}
//go:cgo_export_dynamic GNI_NewCharArray
//go:linkname _cgoexp_cf0433f838ae_GNI_NewCharArray _cgoexp_cf0433f838ae_GNI_NewCharArray
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_NewCharArray
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_NewCharArray(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_NewCharArray
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_NewCharArray(p0 *_Ctype_JNIEnv, p1 _Ctype_jsize) (r0 _Ctype_jcharArray) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_NewCharArray(p0, p1)
}
//go:cgo_export_dynamic GNI_NewShortArray
//go:linkname _cgoexp_cf0433f838ae_GNI_NewShortArray _cgoexp_cf0433f838ae_GNI_NewShortArray
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_NewShortArray
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_NewShortArray(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_NewShortArray
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_NewShortArray(p0 *_Ctype_JNIEnv, p1 _Ctype_jsize) (r0 _Ctype_jshortArray) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_NewShortArray(p0, p1)
}
//go:cgo_export_dynamic GNI_NewIntArray
//go:linkname _cgoexp_cf0433f838ae_GNI_NewIntArray _cgoexp_cf0433f838ae_GNI_NewIntArray
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_NewIntArray
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_NewIntArray(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_NewIntArray
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_NewIntArray(p0 *_Ctype_JNIEnv, p1 _Ctype_jsize) (r0 _Ctype_jintArray) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_NewIntArray(p0, p1)
}
//go:cgo_export_dynamic GNI_NewLongArray
//go:linkname _cgoexp_cf0433f838ae_GNI_NewLongArray _cgoexp_cf0433f838ae_GNI_NewLongArray
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_NewLongArray
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_NewLongArray(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_NewLongArray
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_NewLongArray(p0 *_Ctype_JNIEnv, p1 _Ctype_jsize) (r0 _Ctype_jlongArray) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_NewLongArray(p0, p1)
}
//go:cgo_export_dynamic GNI_NewFloatArray
//go:linkname _cgoexp_cf0433f838ae_GNI_NewFloatArray _cgoexp_cf0433f838ae_GNI_NewFloatArray
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_NewFloatArray
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_NewFloatArray(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_NewFloatArray
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_NewFloatArray(p0 *_Ctype_JNIEnv, p1 _Ctype_jsize) (r0 _Ctype_jfloatArray) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_NewFloatArray(p0, p1)
}
//go:cgo_export_dynamic GNI_NewDoubleArray
//go:linkname _cgoexp_cf0433f838ae_GNI_NewDoubleArray _cgoexp_cf0433f838ae_GNI_NewDoubleArray
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_NewDoubleArray
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_NewDoubleArray(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_NewDoubleArray
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_NewDoubleArray(p0 *_Ctype_JNIEnv, p1 _Ctype_jsize) (r0 _Ctype_jdoubleArray) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_NewDoubleArray(p0, p1)
}
//go:cgo_export_dynamic GNI_GetBooleanArrayElements
//go:linkname _cgoexp_cf0433f838ae_GNI_GetBooleanArrayElements _cgoexp_cf0433f838ae_GNI_GetBooleanArrayElements
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetBooleanArrayElements
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetBooleanArrayElements(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetBooleanArrayElements
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetBooleanArrayElements(p0 *_Ctype_JNIEnv, p1 _Ctype_jbooleanArray, p2 *_Ctype_jboolean) (r0 _Ctype_jboolean) {
	return GNI_GetBooleanArrayElements(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_GetByteArrayElements
//go:linkname _cgoexp_cf0433f838ae_GNI_GetByteArrayElements _cgoexp_cf0433f838ae_GNI_GetByteArrayElements
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetByteArrayElements
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetByteArrayElements(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetByteArrayElements
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetByteArrayElements(p0 *_Ctype_JNIEnv, p1 _Ctype_jbyteArray, p2 *_Ctype_jboolean) (r0 _Ctype_jbyte) {
	return GNI_GetByteArrayElements(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_GetCharArrayElements
//go:linkname _cgoexp_cf0433f838ae_GNI_GetCharArrayElements _cgoexp_cf0433f838ae_GNI_GetCharArrayElements
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetCharArrayElements
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetCharArrayElements(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetCharArrayElements
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetCharArrayElements(p0 *_Ctype_JNIEnv, p1 _Ctype_jcharArray, p2 *_Ctype_jboolean) (r0 _Ctype_jchar) {
	return GNI_GetCharArrayElements(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_GetShortArrayElements
//go:linkname _cgoexp_cf0433f838ae_GNI_GetShortArrayElements _cgoexp_cf0433f838ae_GNI_GetShortArrayElements
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetShortArrayElements
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetShortArrayElements(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetShortArrayElements
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetShortArrayElements(p0 *_Ctype_JNIEnv, p1 _Ctype_jshortArray, p2 *_Ctype_jboolean) (r0 _Ctype_jshort) {
	return GNI_GetShortArrayElements(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_GetIntArrayElements
//go:linkname _cgoexp_cf0433f838ae_GNI_GetIntArrayElements _cgoexp_cf0433f838ae_GNI_GetIntArrayElements
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetIntArrayElements
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetIntArrayElements(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetIntArrayElements
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetIntArrayElements(p0 *_Ctype_JNIEnv, p1 _Ctype_jintArray, p2 *_Ctype_jboolean) (r0 _Ctype_jint) {
	return GNI_GetIntArrayElements(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_GetLongArrayElements
//go:linkname _cgoexp_cf0433f838ae_GNI_GetLongArrayElements _cgoexp_cf0433f838ae_GNI_GetLongArrayElements
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetLongArrayElements
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetLongArrayElements(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetLongArrayElements
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetLongArrayElements(p0 *_Ctype_JNIEnv, p1 _Ctype_jlongArray, p2 *_Ctype_jboolean) (r0 _Ctype_jlong) {
	return GNI_GetLongArrayElements(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_GetFloatArrayElements
//go:linkname _cgoexp_cf0433f838ae_GNI_GetFloatArrayElements _cgoexp_cf0433f838ae_GNI_GetFloatArrayElements
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetFloatArrayElements
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetFloatArrayElements(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetFloatArrayElements
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetFloatArrayElements(p0 *_Ctype_JNIEnv, p1 _Ctype_jfloatArray, p2 *_Ctype_jboolean) (r0 _Ctype_jfloat) {
	return GNI_GetFloatArrayElements(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_GetDoubleArrayElements
//go:linkname _cgoexp_cf0433f838ae_GNI_GetDoubleArrayElements _cgoexp_cf0433f838ae_GNI_GetDoubleArrayElements
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetDoubleArrayElements
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetDoubleArrayElements(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetDoubleArrayElements
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetDoubleArrayElements(p0 *_Ctype_JNIEnv, p1 _Ctype_jdoubleArray, p2 *_Ctype_jboolean) (r0 _Ctype_jdouble) {
	return GNI_GetDoubleArrayElements(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_ReleaseBooleanArrayElements
//go:linkname _cgoexp_cf0433f838ae_GNI_ReleaseBooleanArrayElements _cgoexp_cf0433f838ae_GNI_ReleaseBooleanArrayElements
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_ReleaseBooleanArrayElements
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_ReleaseBooleanArrayElements(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_ReleaseBooleanArrayElements
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_ReleaseBooleanArrayElements(p0 *_Ctype_JNIEnv, p1 _Ctype_jbooleanArray, p2 *_Ctype_jboolean, p3 _Ctype_jint) {
	GNI_ReleaseBooleanArrayElements(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_ReleaseByteArrayElements
//go:linkname _cgoexp_cf0433f838ae_GNI_ReleaseByteArrayElements _cgoexp_cf0433f838ae_GNI_ReleaseByteArrayElements
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_ReleaseByteArrayElements
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_ReleaseByteArrayElements(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_ReleaseByteArrayElements
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_ReleaseByteArrayElements(p0 *_Ctype_JNIEnv, p1 _Ctype_jbyteArray, p2 *_Ctype_jbyte, p3 _Ctype_jint) {
	GNI_ReleaseByteArrayElements(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_ReleaseCharArrayElements
//go:linkname _cgoexp_cf0433f838ae_GNI_ReleaseCharArrayElements _cgoexp_cf0433f838ae_GNI_ReleaseCharArrayElements
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_ReleaseCharArrayElements
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_ReleaseCharArrayElements(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_ReleaseCharArrayElements
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_ReleaseCharArrayElements(p0 *_Ctype_JNIEnv, p1 _Ctype_jcharArray, p2 *_Ctype_jchar, p3 _Ctype_jint) {
	GNI_ReleaseCharArrayElements(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_ReleaseShortArrayElements
//go:linkname _cgoexp_cf0433f838ae_GNI_ReleaseShortArrayElements _cgoexp_cf0433f838ae_GNI_ReleaseShortArrayElements
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_ReleaseShortArrayElements
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_ReleaseShortArrayElements(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_ReleaseShortArrayElements
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_ReleaseShortArrayElements(p0 *_Ctype_JNIEnv, p1 _Ctype_jshortArray, p2 *_Ctype_jshort, p3 _Ctype_jint) {
	GNI_ReleaseShortArrayElements(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_ReleaseIntArrayElements
//go:linkname _cgoexp_cf0433f838ae_GNI_ReleaseIntArrayElements _cgoexp_cf0433f838ae_GNI_ReleaseIntArrayElements
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_ReleaseIntArrayElements
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_ReleaseIntArrayElements(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_ReleaseIntArrayElements
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_ReleaseIntArrayElements(p0 *_Ctype_JNIEnv, p1 _Ctype_jintArray, p2 *_Ctype_jint, p3 _Ctype_jint) {
	GNI_ReleaseIntArrayElements(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_ReleaseLongArrayElements
//go:linkname _cgoexp_cf0433f838ae_GNI_ReleaseLongArrayElements _cgoexp_cf0433f838ae_GNI_ReleaseLongArrayElements
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_ReleaseLongArrayElements
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_ReleaseLongArrayElements(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_ReleaseLongArrayElements
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_ReleaseLongArrayElements(p0 *_Ctype_JNIEnv, p1 _Ctype_jlongArray, p2 *_Ctype_jlong, p3 _Ctype_jint) {
	GNI_ReleaseLongArrayElements(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_ReleaseFloatArrayElements
//go:linkname _cgoexp_cf0433f838ae_GNI_ReleaseFloatArrayElements _cgoexp_cf0433f838ae_GNI_ReleaseFloatArrayElements
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_ReleaseFloatArrayElements
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_ReleaseFloatArrayElements(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_ReleaseFloatArrayElements
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_ReleaseFloatArrayElements(p0 *_Ctype_JNIEnv, p1 _Ctype_jfloatArray, p2 *_Ctype_jfloat, p3 _Ctype_jint) {
	GNI_ReleaseFloatArrayElements(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_ReleaseDoubleArrayElements
//go:linkname _cgoexp_cf0433f838ae_GNI_ReleaseDoubleArrayElements _cgoexp_cf0433f838ae_GNI_ReleaseDoubleArrayElements
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_ReleaseDoubleArrayElements
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_ReleaseDoubleArrayElements(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_ReleaseDoubleArrayElements
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_ReleaseDoubleArrayElements(p0 *_Ctype_JNIEnv, p1 _Ctype_jdoubleArray, p2 *_Ctype_jdouble, p3 _Ctype_jint) {
	GNI_ReleaseDoubleArrayElements(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_GetBooleanArrayRegion
//go:linkname _cgoexp_cf0433f838ae_GNI_GetBooleanArrayRegion _cgoexp_cf0433f838ae_GNI_GetBooleanArrayRegion
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetBooleanArrayRegion
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetBooleanArrayRegion(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetBooleanArrayRegion
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetBooleanArrayRegion(p0 *_Ctype_JNIEnv, p1 _Ctype_jbooleanArray, p2 _Ctype_jsize, p3 _Ctype_jsize, p4 *_Ctype_jboolean) {
	GNI_GetBooleanArrayRegion(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_GetByteArrayRegion
//go:linkname _cgoexp_cf0433f838ae_GNI_GetByteArrayRegion _cgoexp_cf0433f838ae_GNI_GetByteArrayRegion
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetByteArrayRegion
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetByteArrayRegion(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetByteArrayRegion
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetByteArrayRegion(p0 *_Ctype_JNIEnv, p1 _Ctype_jbyteArray, p2 _Ctype_jsize, p3 _Ctype_jsize, p4 *_Ctype_jbyte) {
	GNI_GetByteArrayRegion(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_GetCharArrayRegion
//go:linkname _cgoexp_cf0433f838ae_GNI_GetCharArrayRegion _cgoexp_cf0433f838ae_GNI_GetCharArrayRegion
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetCharArrayRegion
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetCharArrayRegion(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetCharArrayRegion
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetCharArrayRegion(p0 *_Ctype_JNIEnv, p1 _Ctype_jcharArray, p2 _Ctype_jsize, p3 _Ctype_jsize, p4 *_Ctype_jchar) {
	GNI_GetCharArrayRegion(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_GetShortArrayRegion
//go:linkname _cgoexp_cf0433f838ae_GNI_GetShortArrayRegion _cgoexp_cf0433f838ae_GNI_GetShortArrayRegion
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetShortArrayRegion
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetShortArrayRegion(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetShortArrayRegion
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetShortArrayRegion(p0 *_Ctype_JNIEnv, p1 _Ctype_jshortArray, p2 _Ctype_jsize, p3 _Ctype_jsize, p4 *_Ctype_jshort) {
	GNI_GetShortArrayRegion(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_GetIntArrayRegion
//go:linkname _cgoexp_cf0433f838ae_GNI_GetIntArrayRegion _cgoexp_cf0433f838ae_GNI_GetIntArrayRegion
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetIntArrayRegion
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetIntArrayRegion(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetIntArrayRegion
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetIntArrayRegion(p0 *_Ctype_JNIEnv, p1 _Ctype_jintArray, p2 _Ctype_jsize, p3 _Ctype_jsize, p4 *_Ctype_jint) {
	GNI_GetIntArrayRegion(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_GetLongArrayRegion
//go:linkname _cgoexp_cf0433f838ae_GNI_GetLongArrayRegion _cgoexp_cf0433f838ae_GNI_GetLongArrayRegion
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetLongArrayRegion
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetLongArrayRegion(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetLongArrayRegion
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetLongArrayRegion(p0 *_Ctype_JNIEnv, p1 _Ctype_jlongArray, p2 _Ctype_jsize, p3 _Ctype_jsize, p4 *_Ctype_jlong) {
	GNI_GetLongArrayRegion(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_GetFloatArrayRegion
//go:linkname _cgoexp_cf0433f838ae_GNI_GetFloatArrayRegion _cgoexp_cf0433f838ae_GNI_GetFloatArrayRegion
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetFloatArrayRegion
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetFloatArrayRegion(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetFloatArrayRegion
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetFloatArrayRegion(p0 *_Ctype_JNIEnv, p1 _Ctype_jfloatArray, p2 _Ctype_jsize, p3 _Ctype_jsize, p4 *_Ctype_jfloat) {
	GNI_GetFloatArrayRegion(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_GetDoubleArrayRegion
//go:linkname _cgoexp_cf0433f838ae_GNI_GetDoubleArrayRegion _cgoexp_cf0433f838ae_GNI_GetDoubleArrayRegion
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetDoubleArrayRegion
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetDoubleArrayRegion(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetDoubleArrayRegion
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetDoubleArrayRegion(p0 *_Ctype_JNIEnv, p1 _Ctype_jdoubleArray, p2 _Ctype_jsize, p3 _Ctype_jsize, p4 *_Ctype_jdouble) {
	GNI_GetDoubleArrayRegion(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_SetBooleanArrayRegion
//go:linkname _cgoexp_cf0433f838ae_GNI_SetBooleanArrayRegion _cgoexp_cf0433f838ae_GNI_SetBooleanArrayRegion
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_SetBooleanArrayRegion
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_SetBooleanArrayRegion(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_SetBooleanArrayRegion
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_SetBooleanArrayRegion(p0 *_Ctype_JNIEnv, p1 _Ctype_jbooleanArray, p2 _Ctype_jsize, p3 _Ctype_jsize, p4 *_Ctype_jboolean) {
	GNI_SetBooleanArrayRegion(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_SetByteArrayRegion
//go:linkname _cgoexp_cf0433f838ae_GNI_SetByteArrayRegion _cgoexp_cf0433f838ae_GNI_SetByteArrayRegion
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_SetByteArrayRegion
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_SetByteArrayRegion(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_SetByteArrayRegion
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_SetByteArrayRegion(p0 *_Ctype_JNIEnv, p1 _Ctype_jbyteArray, p2 _Ctype_jsize, p3 _Ctype_jsize, p4 *_Ctype_jbyte) {
	GNI_SetByteArrayRegion(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_SetCharArrayRegion
//go:linkname _cgoexp_cf0433f838ae_GNI_SetCharArrayRegion _cgoexp_cf0433f838ae_GNI_SetCharArrayRegion
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_SetCharArrayRegion
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_SetCharArrayRegion(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_SetCharArrayRegion
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_SetCharArrayRegion(p0 *_Ctype_JNIEnv, p1 _Ctype_jcharArray, p2 _Ctype_jsize, p3 _Ctype_jsize, p4 *_Ctype_jchar) {
	GNI_SetCharArrayRegion(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_SetShortArrayRegion
//go:linkname _cgoexp_cf0433f838ae_GNI_SetShortArrayRegion _cgoexp_cf0433f838ae_GNI_SetShortArrayRegion
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_SetShortArrayRegion
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_SetShortArrayRegion(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_SetShortArrayRegion
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_SetShortArrayRegion(p0 *_Ctype_JNIEnv, p1 _Ctype_jshortArray, p2 _Ctype_jsize, p3 _Ctype_jsize, p4 *_Ctype_jshort) {
	GNI_SetShortArrayRegion(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_SetIntArrayRegion
//go:linkname _cgoexp_cf0433f838ae_GNI_SetIntArrayRegion _cgoexp_cf0433f838ae_GNI_SetIntArrayRegion
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_SetIntArrayRegion
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_SetIntArrayRegion(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_SetIntArrayRegion
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_SetIntArrayRegion(p0 *_Ctype_JNIEnv, p1 _Ctype_jintArray, p2 _Ctype_jsize, p3 _Ctype_jsize, p4 *_Ctype_jint) {
	GNI_SetIntArrayRegion(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_SetLongArrayRegion
//go:linkname _cgoexp_cf0433f838ae_GNI_SetLongArrayRegion _cgoexp_cf0433f838ae_GNI_SetLongArrayRegion
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_SetLongArrayRegion
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_SetLongArrayRegion(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_SetLongArrayRegion
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_SetLongArrayRegion(p0 *_Ctype_JNIEnv, p1 _Ctype_jlongArray, p2 _Ctype_jsize, p3 _Ctype_jsize, p4 *_Ctype_jlong) {
	GNI_SetLongArrayRegion(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_SetFloatArrayRegion
//go:linkname _cgoexp_cf0433f838ae_GNI_SetFloatArrayRegion _cgoexp_cf0433f838ae_GNI_SetFloatArrayRegion
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_SetFloatArrayRegion
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_SetFloatArrayRegion(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_SetFloatArrayRegion
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_SetFloatArrayRegion(p0 *_Ctype_JNIEnv, p1 _Ctype_jfloatArray, p2 _Ctype_jsize, p3 _Ctype_jsize, p4 *_Ctype_jfloat) {
	GNI_SetFloatArrayRegion(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_SetDoubleArrayRegion
//go:linkname _cgoexp_cf0433f838ae_GNI_SetDoubleArrayRegion _cgoexp_cf0433f838ae_GNI_SetDoubleArrayRegion
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_SetDoubleArrayRegion
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_SetDoubleArrayRegion(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_SetDoubleArrayRegion
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_SetDoubleArrayRegion(p0 *_Ctype_JNIEnv, p1 _Ctype_jdoubleArray, p2 _Ctype_jsize, p3 _Ctype_jsize, p4 *_Ctype_jdouble) {
	GNI_SetDoubleArrayRegion(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_RegisterNatives
//go:linkname _cgoexp_cf0433f838ae_GNI_RegisterNatives _cgoexp_cf0433f838ae_GNI_RegisterNatives
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_RegisterNatives
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_RegisterNatives(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_RegisterNatives
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_RegisterNatives(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass, p2 *_Ctype_struct___0, p3 _Ctype_jint) (r0 _Ctype_jint) {
	return GNI_RegisterNatives(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_UnregisterNatives
//go:linkname _cgoexp_cf0433f838ae_GNI_UnregisterNatives _cgoexp_cf0433f838ae_GNI_UnregisterNatives
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_UnregisterNatives
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_UnregisterNatives(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_UnregisterNatives
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_UnregisterNatives(p0 *_Ctype_JNIEnv, p1 _Ctype_jclass) (r0 _Ctype_jint) {
	return GNI_UnregisterNatives(p0, p1)
}
//go:cgo_export_dynamic GNI_MonitorEnter
//go:linkname _cgoexp_cf0433f838ae_GNI_MonitorEnter _cgoexp_cf0433f838ae_GNI_MonitorEnter
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_MonitorEnter
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_MonitorEnter(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_MonitorEnter
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_MonitorEnter(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject) (r0 _Ctype_jint) {
	return GNI_MonitorEnter(p0, p1)
}
//go:cgo_export_dynamic GNI_MonitorExit
//go:linkname _cgoexp_cf0433f838ae_GNI_MonitorExit _cgoexp_cf0433f838ae_GNI_MonitorExit
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_MonitorExit
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_MonitorExit(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_MonitorExit
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_MonitorExit(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject) (r0 _Ctype_jint) {
	return GNI_MonitorExit(p0, p1)
}
//go:cgo_export_dynamic GNI_GetJavaVM
//go:linkname _cgoexp_cf0433f838ae_GNI_GetJavaVM _cgoexp_cf0433f838ae_GNI_GetJavaVM
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetJavaVM
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetJavaVM(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetJavaVM
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetJavaVM(p0 *_Ctype_JNIEnv, p1 unsafe.Pointer) (r0 _Ctype_jint) {
	return GNI_GetJavaVM(p0, p1)
}
//go:cgo_export_dynamic GNI_GetStringRegion
//go:linkname _cgoexp_cf0433f838ae_GNI_GetStringRegion _cgoexp_cf0433f838ae_GNI_GetStringRegion
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetStringRegion
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetStringRegion(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetStringRegion
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetStringRegion(p0 *_Ctype_JNIEnv, p1 _Ctype_jstring, p2 _Ctype_jsize, p3 _Ctype_jsize, p4 *_Ctype_jchar) {
	GNI_GetStringRegion(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_GetStringUTFRegion
//go:linkname _cgoexp_cf0433f838ae_GNI_GetStringUTFRegion _cgoexp_cf0433f838ae_GNI_GetStringUTFRegion
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetStringUTFRegion
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetStringUTFRegion(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetStringUTFRegion
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetStringUTFRegion(p0 *_Ctype_JNIEnv, p1 _Ctype_jstring, p2 _Ctype_jsize, p3 _Ctype_jsize, p4 *_Ctype_char) {
	GNI_GetStringUTFRegion(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic GNI_GetPrimitiveArrayCritical
//go:linkname _cgoexp_cf0433f838ae_GNI_GetPrimitiveArrayCritical _cgoexp_cf0433f838ae_GNI_GetPrimitiveArrayCritical
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetPrimitiveArrayCritical
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetPrimitiveArrayCritical(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetPrimitiveArrayCritical
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetPrimitiveArrayCritical(p0 *_Ctype_JNIEnv, p1 _Ctype_jarray, p2 *_Ctype_jboolean) {
	GNI_GetPrimitiveArrayCritical(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_ReleasePrimitiveArrayCritical
//go:linkname _cgoexp_cf0433f838ae_GNI_ReleasePrimitiveArrayCritical _cgoexp_cf0433f838ae_GNI_ReleasePrimitiveArrayCritical
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_ReleasePrimitiveArrayCritical
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_ReleasePrimitiveArrayCritical(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_ReleasePrimitiveArrayCritical
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_ReleasePrimitiveArrayCritical(p0 *_Ctype_JNIEnv, p1 _Ctype_jarray, p2 unsafe.Pointer, p3 _Ctype_jint) {
	GNI_ReleasePrimitiveArrayCritical(p0, p1, p2, p3)
}
//go:cgo_export_dynamic GNI_GetStringCritical
//go:linkname _cgoexp_cf0433f838ae_GNI_GetStringCritical _cgoexp_cf0433f838ae_GNI_GetStringCritical
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetStringCritical
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetStringCritical(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetStringCritical
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetStringCritical(p0 *_Ctype_JNIEnv, p1 _Ctype_jstring, p2 *_Ctype_jboolean) (r0 *_Ctype_jchar) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_GetStringCritical(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_ReleaseStringCritical
//go:linkname _cgoexp_cf0433f838ae_GNI_ReleaseStringCritical _cgoexp_cf0433f838ae_GNI_ReleaseStringCritical
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_ReleaseStringCritical
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_ReleaseStringCritical(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_ReleaseStringCritical
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_ReleaseStringCritical(p0 *_Ctype_JNIEnv, p1 _Ctype_jstring, p2 *_Ctype_jchar) {
	GNI_ReleaseStringCritical(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_NewWeakGlobalRef
//go:linkname _cgoexp_cf0433f838ae_GNI_NewWeakGlobalRef _cgoexp_cf0433f838ae_GNI_NewWeakGlobalRef
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_NewWeakGlobalRef
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_NewWeakGlobalRef(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_NewWeakGlobalRef
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_NewWeakGlobalRef(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject) (r0 _Ctype_jweak) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_NewWeakGlobalRef(p0, p1)
}
//go:cgo_export_dynamic GNI_DeleteWeakGlobalRef
//go:linkname _cgoexp_cf0433f838ae_GNI_DeleteWeakGlobalRef _cgoexp_cf0433f838ae_GNI_DeleteWeakGlobalRef
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_DeleteWeakGlobalRef
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_DeleteWeakGlobalRef(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_DeleteWeakGlobalRef
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_DeleteWeakGlobalRef(p0 *_Ctype_JNIEnv, p1 _Ctype_jweak) {
	GNI_DeleteWeakGlobalRef(p0, p1)
}
//go:cgo_export_dynamic GNI_ExceptionCheck
//go:linkname _cgoexp_cf0433f838ae_GNI_ExceptionCheck _cgoexp_cf0433f838ae_GNI_ExceptionCheck
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_ExceptionCheck
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_ExceptionCheck(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_ExceptionCheck
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_ExceptionCheck(p0 *_Ctype_JNIEnv) (r0 _Ctype_jboolean) {
	return GNI_ExceptionCheck(p0)
}
//go:cgo_export_dynamic GNI_NewDirectByteBuffer
//go:linkname _cgoexp_cf0433f838ae_GNI_NewDirectByteBuffer _cgoexp_cf0433f838ae_GNI_NewDirectByteBuffer
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_NewDirectByteBuffer
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_NewDirectByteBuffer(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_NewDirectByteBuffer
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_NewDirectByteBuffer(p0 *_Ctype_JNIEnv, p1 unsafe.Pointer, p2 _Ctype_jlong) (r0 _Ctype_jobject) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_NewDirectByteBuffer(p0, p1, p2)
}
//go:cgo_export_dynamic GNI_GetDirectBufferAddress
//go:linkname _cgoexp_cf0433f838ae_GNI_GetDirectBufferAddress _cgoexp_cf0433f838ae_GNI_GetDirectBufferAddress
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetDirectBufferAddress
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetDirectBufferAddress(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetDirectBufferAddress
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetDirectBufferAddress(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject) (r0 unsafe.Pointer) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return GNI_GetDirectBufferAddress(p0, p1)
}
//go:cgo_export_dynamic GNI_GetDirectBufferCapacity
//go:linkname _cgoexp_cf0433f838ae_GNI_GetDirectBufferCapacity _cgoexp_cf0433f838ae_GNI_GetDirectBufferCapacity
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetDirectBufferCapacity
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetDirectBufferCapacity(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetDirectBufferCapacity
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetDirectBufferCapacity(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject) (r0 _Ctype_jlong) {
	return GNI_GetDirectBufferCapacity(p0, p1)
}
//go:cgo_export_dynamic GNI_GetObjectRefType
//go:linkname _cgoexp_cf0433f838ae_GNI_GetObjectRefType _cgoexp_cf0433f838ae_GNI_GetObjectRefType
//go:cgo_export_static _cgoexp_cf0433f838ae_GNI_GetObjectRefType
//go:nosplit
//go:norace
func _cgoexp_cf0433f838ae_GNI_GetObjectRefType(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_cf0433f838ae_GNI_GetObjectRefType
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_cf0433f838ae_GNI_GetObjectRefType(p0 *_Ctype_JNIEnv, p1 _Ctype_jobject) (r0 _Ctype_jobjectRefType) {
	return GNI_GetObjectRefType(p0, p1)
}
