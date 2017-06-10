/* Created by cgo - DO NOT EDIT. */
#include <stdlib.h>
#include "_cgo_export.h"

extern void crosscall2(void (*fn)(void *, int, __SIZE_TYPE__), void *, int, __SIZE_TYPE__);
extern __SIZE_TYPE__ _cgo_wait_runtime_init_done();
extern void _cgo_release_context(__SIZE_TYPE__);

extern char* _cgo_topofstack(void);
#define CGO_NO_SANITIZE_THREAD
#define _cgo_tsan_acquire()
#define _cgo_tsan_release()

extern void _cgoexp_cf0433f838ae_GNI_GetVersion(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jint GNI_GetVersion(JNIEnv* p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jint r0;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetVersion, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_DefineClass(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jclass GNI_DefineClass(JNIEnv* p0, char* p1, jobject p2, jbyte* p3, jsize p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		char* p1;
		jobject p2;
		jbyte* p3;
		jsize p4;
		char __pad0[4];
		jclass r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_DefineClass, &a, 48, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_FindClass(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jclass GNI_FindClass(JNIEnv* p0, char* p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		char* p1;
		jclass r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_FindClass, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_FromReflectedMethod(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jmethodID GNI_FromReflectedMethod(JNIEnv* p0, jobject p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jmethodID r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_FromReflectedMethod, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_FromReflectedField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jfieldID GNI_FromReflectedField(JNIEnv* p0, jobject p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jfieldID r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_FromReflectedField, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_ToReflectedMethod(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jobject GNI_ToReflectedMethod(JNIEnv* p0, jclass p1, jmethodID p2, jboolean p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jmethodID p2;
		jboolean p3;
		char __pad0[7];
		jobject r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_ToReflectedMethod, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetSuperclass(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jclass GNI_GetSuperclass(JNIEnv* p0, jclass p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jclass r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetSuperclass, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_IsAssignableFrom(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jboolean GNI_IsAssignableFrom(JNIEnv* p0, jclass p1, jclass p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jclass p2;
		jboolean r0;
		char __pad0[7];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_IsAssignableFrom, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_ToReflectedField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jobject GNI_ToReflectedField(JNIEnv* p0, jclass p1, jfieldID p2, jboolean p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jfieldID p2;
		jboolean p3;
		char __pad0[7];
		jobject r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_ToReflectedField, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_Throw(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jint GNI_Throw(JNIEnv* p0, jthrowable p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jthrowable p1;
		jint r0;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_Throw, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_ThrowNew(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jint GNI_ThrowNew(JNIEnv* p0, jclass p1, char* p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		char* p2;
		jint r0;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_ThrowNew, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_ExceptionOccurred(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jthrowable GNI_ExceptionOccurred(JNIEnv* p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jthrowable r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_ExceptionOccurred, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_ExceptionDescribe(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_ExceptionDescribe(JNIEnv* p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_ExceptionDescribe, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_ExceptionClear(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_ExceptionClear(JNIEnv* p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_ExceptionClear, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_FatalError(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_FatalError(JNIEnv* p0, char* p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		char* p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_FatalError, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_PushLocalFrame(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jint GNI_PushLocalFrame(JNIEnv* p0, jint p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jint p1;
		char __pad0[4];
		jint r0;
		char __pad1[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_PushLocalFrame, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_PopLocalFrame(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jobject GNI_PopLocalFrame(JNIEnv* p0, jobject p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jobject r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_PopLocalFrame, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_NewGlobalRef(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jobject GNI_NewGlobalRef(JNIEnv* p0, jobject p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jobject r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_NewGlobalRef, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_DeleteGlobalRef(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_DeleteGlobalRef(JNIEnv* p0, jobject p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_DeleteGlobalRef, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_DeleteLocalRef(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_DeleteLocalRef(JNIEnv* p0, jobject p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_DeleteLocalRef, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_IsSameObject(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jboolean GNI_IsSameObject(JNIEnv* p0, jobject p1, jobject p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jobject p2;
		jboolean r0;
		char __pad0[7];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_IsSameObject, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_NewLocalRef(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jobject GNI_NewLocalRef(JNIEnv* p0, jobject p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jobject r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_NewLocalRef, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_EnsureLocalCapacity(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jint GNI_EnsureLocalCapacity(JNIEnv* p0, jint p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jint p1;
		char __pad0[4];
		jint r0;
		char __pad1[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_EnsureLocalCapacity, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_AllocObject(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jobject GNI_AllocObject(JNIEnv* p0, jclass p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jobject r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_AllocObject, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_NewObjectV(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jobject GNI_NewObjectV(JNIEnv* p0, jclass p1, jmethodID p2, va_list* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jmethodID p2;
		va_list* p3;
		jobject r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_NewObjectV, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_NewObjectA(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jobject GNI_NewObjectA(JNIEnv* p0, jclass p1, jmethodID p2, jvalue* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jmethodID p2;
		jvalue* p3;
		jobject r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_NewObjectA, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetObjectClass(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jclass GNI_GetObjectClass(JNIEnv* p0, jobject p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jclass r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetObjectClass, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_IsInstanceOf(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jboolean GNI_IsInstanceOf(JNIEnv* p0, jobject p1, jclass p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jclass p2;
		jboolean r0;
		char __pad0[7];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_IsInstanceOf, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetMethodID(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jmethodID GNI_GetMethodID(JNIEnv* p0, jclass p1, char* p2, char* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		char* p2;
		char* p3;
		jmethodID r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetMethodID, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallObjectMethodV(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jobject GNI_CallObjectMethodV(JNIEnv* p0, jobject p1, jmethodID p2, va_list* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jmethodID p2;
		va_list* p3;
		jobject r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallObjectMethodV, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallObjectMethodA(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jobject GNI_CallObjectMethodA(JNIEnv* p0, jobject p1, jmethodID p2, jvalue* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jmethodID p2;
		jvalue* p3;
		jobject r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallObjectMethodA, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallBooleanMethodV(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jboolean GNI_CallBooleanMethodV(JNIEnv* p0, jobject p1, jmethodID p2, va_list* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jmethodID p2;
		va_list* p3;
		jboolean r0;
		char __pad0[7];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallBooleanMethodV, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallBooleanMethodA(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jboolean GNI_CallBooleanMethodA(JNIEnv* p0, jobject p1, jmethodID p2, jvalue* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jmethodID p2;
		jvalue* p3;
		jboolean r0;
		char __pad0[7];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallBooleanMethodA, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallByteMethodV(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jbyte GNI_CallByteMethodV(JNIEnv* p0, jobject p1, jmethodID p2, va_list* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jmethodID p2;
		va_list* p3;
		jbyte r0;
		char __pad0[7];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallByteMethodV, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallByteMethodA(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jbyte GNI_CallByteMethodA(JNIEnv* p0, jobject p1, jmethodID p2, jvalue* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jmethodID p2;
		jvalue* p3;
		jbyte r0;
		char __pad0[7];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallByteMethodA, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallCharMethodV(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jchar GNI_CallCharMethodV(JNIEnv* p0, jobject p1, jmethodID p2, va_list* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jmethodID p2;
		va_list* p3;
		jchar r0;
		char __pad0[6];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallCharMethodV, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallCharMethodA(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jchar GNI_CallCharMethodA(JNIEnv* p0, jobject p1, jmethodID p2, jvalue* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jmethodID p2;
		jvalue* p3;
		jchar r0;
		char __pad0[6];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallCharMethodA, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallShortMethodV(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jshort GNI_CallShortMethodV(JNIEnv* p0, jobject p1, jmethodID p2, va_list* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jmethodID p2;
		va_list* p3;
		jshort r0;
		char __pad0[6];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallShortMethodV, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallShortMethodA(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jshort GNI_CallShortMethodA(JNIEnv* p0, jobject p1, jmethodID p2, jvalue* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jmethodID p2;
		jvalue* p3;
		jshort r0;
		char __pad0[6];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallShortMethodA, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallIntMethodV(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jint GNI_CallIntMethodV(JNIEnv* p0, jobject p1, jmethodID p2, va_list* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jmethodID p2;
		va_list* p3;
		jint r0;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallIntMethodV, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallIntMethodA(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jint GNI_CallIntMethodA(JNIEnv* p0, jobject p1, jmethodID p2, jvalue* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jmethodID p2;
		jvalue* p3;
		jint r0;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallIntMethodA, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallLongMethodV(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jlong GNI_CallLongMethodV(JNIEnv* p0, jobject p1, jmethodID p2, va_list* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jmethodID p2;
		va_list* p3;
		jlong r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallLongMethodV, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallLongMethodA(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jlong GNI_CallLongMethodA(JNIEnv* p0, jobject p1, jmethodID p2, jvalue* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jmethodID p2;
		jvalue* p3;
		jlong r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallLongMethodA, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallFloatMethodV(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jfloat GNI_CallFloatMethodV(JNIEnv* p0, jobject p1, jmethodID p2, va_list* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jmethodID p2;
		va_list* p3;
		jfloat r0;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallFloatMethodV, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallFloatMethodA(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jfloat GNI_CallFloatMethodA(JNIEnv* p0, jobject p1, jmethodID p2, jvalue* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jmethodID p2;
		jvalue* p3;
		jfloat r0;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallFloatMethodA, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallDoubleMethodV(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jdouble GNI_CallDoubleMethodV(JNIEnv* p0, jobject p1, jmethodID p2, va_list* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jmethodID p2;
		va_list* p3;
		jdouble r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallDoubleMethodV, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallDoubleMethodA(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jdouble GNI_CallDoubleMethodA(JNIEnv* p0, jobject p1, jmethodID p2, jvalue* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jmethodID p2;
		jvalue* p3;
		jdouble r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallDoubleMethodA, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallVoidMethodV(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_CallVoidMethodV(JNIEnv* p0, jobject p1, jmethodID p2, va_list* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jmethodID p2;
		va_list* p3;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallVoidMethodV, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_CallVoidMethodA(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_CallVoidMethodA(JNIEnv* p0, jobject p1, jmethodID p2, jvalue* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jmethodID p2;
		jvalue* p3;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallVoidMethodA, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_CallNonvirtualObjectMethodV(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jobject GNI_CallNonvirtualObjectMethodV(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, va_list* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jclass p2;
		jmethodID p3;
		va_list* p4;
		jobject r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallNonvirtualObjectMethodV, &a, 48, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallNonvirtualObjectMethodA(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jobject GNI_CallNonvirtualObjectMethodA(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, jvalue* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jclass p2;
		jmethodID p3;
		jvalue* p4;
		jobject r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallNonvirtualObjectMethodA, &a, 48, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallNonvirtualBooleanMethodV(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jboolean GNI_CallNonvirtualBooleanMethodV(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, va_list* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jclass p2;
		jmethodID p3;
		va_list* p4;
		jboolean r0;
		char __pad0[7];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallNonvirtualBooleanMethodV, &a, 48, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallNonvirtualBooleanMethodA(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jboolean GNI_CallNonvirtualBooleanMethodA(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, jvalue* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jclass p2;
		jmethodID p3;
		jvalue* p4;
		jboolean r0;
		char __pad0[7];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallNonvirtualBooleanMethodA, &a, 48, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallNonvirtualByteMethodV(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jbyte GNI_CallNonvirtualByteMethodV(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, va_list* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jclass p2;
		jmethodID p3;
		va_list* p4;
		jbyte r0;
		char __pad0[7];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallNonvirtualByteMethodV, &a, 48, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallNonvirtualByteMethodA(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jbyte GNI_CallNonvirtualByteMethodA(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, jvalue* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jclass p2;
		jmethodID p3;
		jvalue* p4;
		jbyte r0;
		char __pad0[7];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallNonvirtualByteMethodA, &a, 48, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallNonvirtualCharMethodV(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jchar GNI_CallNonvirtualCharMethodV(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, va_list* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jclass p2;
		jmethodID p3;
		va_list* p4;
		jchar r0;
		char __pad0[6];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallNonvirtualCharMethodV, &a, 48, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallNonvirtualCharMethodA(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jchar GNI_CallNonvirtualCharMethodA(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, jvalue* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jclass p2;
		jmethodID p3;
		jvalue* p4;
		jchar r0;
		char __pad0[6];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallNonvirtualCharMethodA, &a, 48, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallNonvirtualShortMethodV(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jshort GNI_CallNonvirtualShortMethodV(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, va_list* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jclass p2;
		jmethodID p3;
		va_list* p4;
		jshort r0;
		char __pad0[6];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallNonvirtualShortMethodV, &a, 48, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallNonvirtualShortMethodA(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jshort GNI_CallNonvirtualShortMethodA(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, jvalue* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jclass p2;
		jmethodID p3;
		jvalue* p4;
		jshort r0;
		char __pad0[6];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallNonvirtualShortMethodA, &a, 48, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallNonvirtualIntMethodV(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jint GNI_CallNonvirtualIntMethodV(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, va_list* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jclass p2;
		jmethodID p3;
		va_list* p4;
		jint r0;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallNonvirtualIntMethodV, &a, 48, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallNonvirtualIntMethodA(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jint GNI_CallNonvirtualIntMethodA(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, jvalue* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jclass p2;
		jmethodID p3;
		jvalue* p4;
		jint r0;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallNonvirtualIntMethodA, &a, 48, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallNonvirtualLongMethodV(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jlong GNI_CallNonvirtualLongMethodV(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, va_list* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jclass p2;
		jmethodID p3;
		va_list* p4;
		jlong r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallNonvirtualLongMethodV, &a, 48, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallNonvirtualLongMethodA(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jlong GNI_CallNonvirtualLongMethodA(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, jvalue* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jclass p2;
		jmethodID p3;
		jvalue* p4;
		jlong r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallNonvirtualLongMethodA, &a, 48, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallNonvirtualFloatMethodV(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jfloat GNI_CallNonvirtualFloatMethodV(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, va_list* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jclass p2;
		jmethodID p3;
		va_list* p4;
		jfloat r0;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallNonvirtualFloatMethodV, &a, 48, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallNonvirtualFloatMethodA(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jfloat GNI_CallNonvirtualFloatMethodA(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, jvalue* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jclass p2;
		jmethodID p3;
		jvalue* p4;
		jfloat r0;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallNonvirtualFloatMethodA, &a, 48, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallNonvirtualDoubleMethodV(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jdouble GNI_CallNonvirtualDoubleMethodV(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, va_list* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jclass p2;
		jmethodID p3;
		va_list* p4;
		jdouble r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallNonvirtualDoubleMethodV, &a, 48, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallNonvirtualDoubleMethodA(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jdouble GNI_CallNonvirtualDoubleMethodA(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, jvalue* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jclass p2;
		jmethodID p3;
		jvalue* p4;
		jdouble r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallNonvirtualDoubleMethodA, &a, 48, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallNonvirtualVoidMethodV(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_CallNonvirtualVoidMethodV(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, va_list* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jclass p2;
		jmethodID p3;
		va_list* p4;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallNonvirtualVoidMethodV, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_CallNonvirtualVoidMethodA(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_CallNonvirtualVoidMethodA(JNIEnv* p0, jobject p1, jclass p2, jmethodID p3, jvalue* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jclass p2;
		jmethodID p3;
		jvalue* p4;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallNonvirtualVoidMethodA, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_GetFieldID(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jfieldID GNI_GetFieldID(JNIEnv* p0, jclass p1, char* p2, char* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		char* p2;
		char* p3;
		jfieldID r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetFieldID, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetObjectField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jobject GNI_GetObjectField(JNIEnv* p0, jobject p1, jfieldID p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jfieldID p2;
		jobject r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetObjectField, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetBooleanField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jboolean GNI_GetBooleanField(JNIEnv* p0, jobject p1, jfieldID p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jfieldID p2;
		jboolean r0;
		char __pad0[7];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetBooleanField, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetByteField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jbyte GNI_GetByteField(JNIEnv* p0, jobject p1, jfieldID p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jfieldID p2;
		jbyte r0;
		char __pad0[7];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetByteField, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetCharField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jchar GNI_GetCharField(JNIEnv* p0, jobject p1, jfieldID p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jfieldID p2;
		jchar r0;
		char __pad0[6];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetCharField, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetShortField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jshort GNI_GetShortField(JNIEnv* p0, jobject p1, jfieldID p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jfieldID p2;
		jshort r0;
		char __pad0[6];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetShortField, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetIntField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jint GNI_GetIntField(JNIEnv* p0, jobject p1, jfieldID p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jfieldID p2;
		jint r0;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetIntField, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetLongField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jlong GNI_GetLongField(JNIEnv* p0, jobject p1, jfieldID p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jfieldID p2;
		jlong r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetLongField, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetFloatField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jfloat GNI_GetFloatField(JNIEnv* p0, jobject p1, jfieldID p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jfieldID p2;
		jfloat r0;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetFloatField, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetDoubleField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jdouble GNI_GetDoubleField(JNIEnv* p0, jobject p1, jfieldID p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jfieldID p2;
		jdouble r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetDoubleField, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_SetObjectField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_SetObjectField(JNIEnv* p0, jobject p1, jfieldID p2, jobject p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jfieldID p2;
		jobject p3;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_SetObjectField, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_SetBooleanField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_SetBooleanField(JNIEnv* p0, jobject p1, jfieldID p2, jboolean p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jfieldID p2;
		jboolean p3;
		char __pad0[7];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_SetBooleanField, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_SetByteField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_SetByteField(JNIEnv* p0, jobject p1, jfieldID p2, jbyte p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jfieldID p2;
		jbyte p3;
		char __pad0[7];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_SetByteField, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_SetCharField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_SetCharField(JNIEnv* p0, jobject p1, jfieldID p2, jchar p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jfieldID p2;
		jchar p3;
		char __pad0[6];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_SetCharField, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_SetShortField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_SetShortField(JNIEnv* p0, jobject p1, jfieldID p2, jshort p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jfieldID p2;
		jshort p3;
		char __pad0[6];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_SetShortField, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_SetIntField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_SetIntField(JNIEnv* p0, jobject p1, jfieldID p2, jint p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jfieldID p2;
		jint p3;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_SetIntField, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_SetLongField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_SetLongField(JNIEnv* p0, jobject p1, jfieldID p2, jlong p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jfieldID p2;
		jlong p3;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_SetLongField, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_SetFloatField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_SetFloatField(JNIEnv* p0, jobject p1, jfieldID p2, jfloat p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jfieldID p2;
		jfloat p3;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_SetFloatField, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_SetDoubleField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_SetDoubleField(JNIEnv* p0, jobject p1, jfieldID p2, jdouble p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jfieldID p2;
		jdouble p3;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_SetDoubleField, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_GetStaticMethodID(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jmethodID GNI_GetStaticMethodID(JNIEnv* p0, jclass p1, char* p2, char* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		char* p2;
		char* p3;
		jmethodID r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetStaticMethodID, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallStaticObjectMethodV(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jobject GNI_CallStaticObjectMethodV(JNIEnv* p0, jclass p1, jmethodID p2, va_list* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jmethodID p2;
		va_list* p3;
		jobject r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallStaticObjectMethodV, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallStaticObjectMethodA(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jobject GNI_CallStaticObjectMethodA(JNIEnv* p0, jclass p1, jmethodID p2, jvalue* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jmethodID p2;
		jvalue* p3;
		jobject r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallStaticObjectMethodA, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallStaticBooleanMethodV(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jboolean GNI_CallStaticBooleanMethodV(JNIEnv* p0, jclass p1, jmethodID p2, va_list* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jmethodID p2;
		va_list* p3;
		jboolean r0;
		char __pad0[7];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallStaticBooleanMethodV, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallStaticBooleanMethodA(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jboolean GNI_CallStaticBooleanMethodA(JNIEnv* p0, jclass p1, jmethodID p2, jvalue* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jmethodID p2;
		jvalue* p3;
		jboolean r0;
		char __pad0[7];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallStaticBooleanMethodA, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallStaticByteMethodV(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jbyte GNI_CallStaticByteMethodV(JNIEnv* p0, jclass p1, jmethodID p2, va_list* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jmethodID p2;
		va_list* p3;
		jbyte r0;
		char __pad0[7];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallStaticByteMethodV, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallStaticByteMethodA(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jbyte GNI_CallStaticByteMethodA(JNIEnv* p0, jclass p1, jmethodID p2, jvalue* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jmethodID p2;
		jvalue* p3;
		jbyte r0;
		char __pad0[7];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallStaticByteMethodA, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallStaticCharMethodV(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jchar GNI_CallStaticCharMethodV(JNIEnv* p0, jclass p1, jmethodID p2, va_list* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jmethodID p2;
		va_list* p3;
		jchar r0;
		char __pad0[6];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallStaticCharMethodV, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallStaticCharMethodA(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jchar GNI_CallStaticCharMethodA(JNIEnv* p0, jclass p1, jmethodID p2, jvalue* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jmethodID p2;
		jvalue* p3;
		jchar r0;
		char __pad0[6];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallStaticCharMethodA, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallStaticShortMethodA(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jshort GNI_CallStaticShortMethodA(JNIEnv* p0, jclass p1, jmethodID p2, jvalue* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jmethodID p2;
		jvalue* p3;
		jshort r0;
		char __pad0[6];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallStaticShortMethodA, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallStaticIntMethodV(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jint GNI_CallStaticIntMethodV(JNIEnv* p0, jclass p1, jmethodID p2, va_list* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jmethodID p2;
		va_list* p3;
		jint r0;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallStaticIntMethodV, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallStaticIntMethodA(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jint GNI_CallStaticIntMethodA(JNIEnv* p0, jclass p1, jmethodID p2, jvalue* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jmethodID p2;
		jvalue* p3;
		jint r0;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallStaticIntMethodA, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallStaticLongMethodV(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jlong GNI_CallStaticLongMethodV(JNIEnv* p0, jclass p1, jmethodID p2, va_list* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jmethodID p2;
		va_list* p3;
		jlong r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallStaticLongMethodV, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallStaticLongMethodA(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jlong GNI_CallStaticLongMethodA(JNIEnv* p0, jclass p1, jmethodID p2, jvalue* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jmethodID p2;
		jvalue* p3;
		jlong r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallStaticLongMethodA, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallStaticFloatMethodV(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jfloat GNI_CallStaticFloatMethodV(JNIEnv* p0, jclass p1, jmethodID p2, va_list* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jmethodID p2;
		va_list* p3;
		jfloat r0;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallStaticFloatMethodV, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallStaticFloatMethodA(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jfloat GNI_CallStaticFloatMethodA(JNIEnv* p0, jclass p1, jmethodID p2, jvalue* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jmethodID p2;
		jvalue* p3;
		jfloat r0;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallStaticFloatMethodA, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallStaticDoubleMethodV(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jdouble GNI_CallStaticDoubleMethodV(JNIEnv* p0, jclass p1, jmethodID p2, va_list* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jmethodID p2;
		va_list* p3;
		jdouble r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallStaticDoubleMethodV, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallStaticDoubleMethodA(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jdouble GNI_CallStaticDoubleMethodA(JNIEnv* p0, jclass p1, jmethodID p2, jvalue* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jmethodID p2;
		jvalue* p3;
		jdouble r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallStaticDoubleMethodA, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_CallStaticVoidMethodV(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_CallStaticVoidMethodV(JNIEnv* p0, jclass p1, jmethodID p2, va_list* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jmethodID p2;
		va_list* p3;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallStaticVoidMethodV, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_CallStaticVoidMethodA(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_CallStaticVoidMethodA(JNIEnv* p0, jclass p1, jmethodID p2, jvalue* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jmethodID p2;
		jvalue* p3;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_CallStaticVoidMethodA, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_GetStaticFieldID(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jfieldID GNI_GetStaticFieldID(JNIEnv* p0, jclass p1, char* p2, char* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		char* p2;
		char* p3;
		jfieldID r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetStaticFieldID, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetStaticObjectField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jobject GNI_GetStaticObjectField(JNIEnv* p0, jclass p1, jfieldID p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jfieldID p2;
		jobject r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetStaticObjectField, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetStaticBooleanField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jboolean GNI_GetStaticBooleanField(JNIEnv* p0, jclass p1, jfieldID p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jfieldID p2;
		jboolean r0;
		char __pad0[7];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetStaticBooleanField, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetStaticByteField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jbyte GNI_GetStaticByteField(JNIEnv* p0, jclass p1, jfieldID p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jfieldID p2;
		jbyte r0;
		char __pad0[7];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetStaticByteField, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetStaticCharField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jchar GNI_GetStaticCharField(JNIEnv* p0, jclass p1, jfieldID p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jfieldID p2;
		jchar r0;
		char __pad0[6];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetStaticCharField, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetStaticShortField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jshort GNI_GetStaticShortField(JNIEnv* p0, jclass p1, jfieldID p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jfieldID p2;
		jshort r0;
		char __pad0[6];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetStaticShortField, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetStaticIntField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jint GNI_GetStaticIntField(JNIEnv* p0, jclass p1, jfieldID p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jfieldID p2;
		jint r0;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetStaticIntField, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetStaticLongField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jlong GNI_GetStaticLongField(JNIEnv* p0, jclass p1, jfieldID p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jfieldID p2;
		jlong r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetStaticLongField, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetStaticFloatField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jfloat GNI_GetStaticFloatField(JNIEnv* p0, jclass p1, jfieldID p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jfieldID p2;
		jfloat r0;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetStaticFloatField, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetStaticDoubleField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jdouble GNI_GetStaticDoubleField(JNIEnv* p0, jclass p1, jfieldID p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jfieldID p2;
		jdouble r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetStaticDoubleField, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_SetStaticObjectField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_SetStaticObjectField(JNIEnv* p0, jclass p1, jfieldID p2, jobject p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jfieldID p2;
		jobject p3;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_SetStaticObjectField, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_SetStaticBooleanField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_SetStaticBooleanField(JNIEnv* p0, jclass p1, jfieldID p2, jboolean p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jfieldID p2;
		jboolean p3;
		char __pad0[7];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_SetStaticBooleanField, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_SetStaticByteField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_SetStaticByteField(JNIEnv* p0, jclass p1, jfieldID p2, jbyte p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jfieldID p2;
		jbyte p3;
		char __pad0[7];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_SetStaticByteField, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_SetStaticCharField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_SetStaticCharField(JNIEnv* p0, jclass p1, jfieldID p2, jchar p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jfieldID p2;
		jchar p3;
		char __pad0[6];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_SetStaticCharField, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_SetStaticShortField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_SetStaticShortField(JNIEnv* p0, jclass p1, jfieldID p2, jshort p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jfieldID p2;
		jshort p3;
		char __pad0[6];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_SetStaticShortField, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_SetStaticIntField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_SetStaticIntField(JNIEnv* p0, jclass p1, jfieldID p2, jint p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jfieldID p2;
		jint p3;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_SetStaticIntField, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_SetStaticLongField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_SetStaticLongField(JNIEnv* p0, jclass p1, jfieldID p2, jlong p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jfieldID p2;
		jlong p3;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_SetStaticLongField, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_SetStaticFloatField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_SetStaticFloatField(JNIEnv* p0, jclass p1, jfieldID p2, jfloat p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jfieldID p2;
		jfloat p3;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_SetStaticFloatField, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_SetStaticDoubleField(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_SetStaticDoubleField(JNIEnv* p0, jclass p1, jfieldID p2, jdouble p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jfieldID p2;
		jdouble p3;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_SetStaticDoubleField, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_NewString(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jstring GNI_NewString(JNIEnv* p0, jchar* p1, jsize p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jchar* p1;
		jsize p2;
		char __pad0[4];
		jstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_NewString, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetStringLength(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jsize GNI_GetStringLength(JNIEnv* p0, jstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jstring p1;
		jsize r0;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetStringLength, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetStringChars(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jchar* GNI_GetStringChars(JNIEnv* p0, jstring p1, jboolean* p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jstring p1;
		jboolean* p2;
		jchar* r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetStringChars, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_ReleaseStringChars(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_ReleaseStringChars(JNIEnv* p0, jstring p1, jchar p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jstring p1;
		jchar p2;
		char __pad0[6];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_ReleaseStringChars, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_NewStringUTF(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jstring GNI_NewStringUTF(JNIEnv* p0, char* p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		char* p1;
		jstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_NewStringUTF, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetStringUTFLength(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jsize GNI_GetStringUTFLength(JNIEnv* p0, jstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jstring p1;
		jsize r0;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetStringUTFLength, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetStringUTFChars(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
char* GNI_GetStringUTFChars(JNIEnv* p0, jstring p1, jboolean* p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jstring p1;
		jboolean* p2;
		char* r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetStringUTFChars, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_ReleaseStringUTFChars(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_ReleaseStringUTFChars(JNIEnv* p0, jstring p1, char* p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jstring p1;
		char* p2;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_ReleaseStringUTFChars, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_GetArrayLength(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jsize GNI_GetArrayLength(JNIEnv* p0, jarray* p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jarray* p1;
		jsize r0;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetArrayLength, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_NewObjectArray(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jobjectArray GNI_NewObjectArray(JNIEnv* p0, jsize p1, jclass p2, jobject p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jsize p1;
		char __pad0[4];
		jclass p2;
		jobject p3;
		jobjectArray r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_NewObjectArray, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetObjectArrayElement(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jobject GNI_GetObjectArrayElement(JNIEnv* p0, jobjectArray p1, jsize p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobjectArray p1;
		jsize p2;
		char __pad0[4];
		jobject r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetObjectArrayElement, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_SetObjectArrayElement(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_SetObjectArrayElement(JNIEnv* p0, jobjectArray p1, jsize p2, jobject p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobjectArray p1;
		jsize p2;
		char __pad0[4];
		jobject p3;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_SetObjectArrayElement, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_NewBooleanArray(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jbooleanArray GNI_NewBooleanArray(JNIEnv* p0, jsize p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jsize p1;
		char __pad0[4];
		jbooleanArray r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_NewBooleanArray, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_NewByteArray(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jbyteArray GNI_NewByteArray(JNIEnv* p0, jsize p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jsize p1;
		char __pad0[4];
		jbyteArray r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_NewByteArray, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_NewCharArray(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jcharArray GNI_NewCharArray(JNIEnv* p0, jsize p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jsize p1;
		char __pad0[4];
		jcharArray r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_NewCharArray, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_NewShortArray(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jshortArray GNI_NewShortArray(JNIEnv* p0, jsize p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jsize p1;
		char __pad0[4];
		jshortArray r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_NewShortArray, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_NewIntArray(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jintArray GNI_NewIntArray(JNIEnv* p0, jsize p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jsize p1;
		char __pad0[4];
		jintArray r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_NewIntArray, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_NewLongArray(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jlongArray GNI_NewLongArray(JNIEnv* p0, jsize p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jsize p1;
		char __pad0[4];
		jlongArray r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_NewLongArray, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_NewFloatArray(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jfloatArray GNI_NewFloatArray(JNIEnv* p0, jsize p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jsize p1;
		char __pad0[4];
		jfloatArray r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_NewFloatArray, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_NewDoubleArray(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jdoubleArray GNI_NewDoubleArray(JNIEnv* p0, jsize p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jsize p1;
		char __pad0[4];
		jdoubleArray r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_NewDoubleArray, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetBooleanArrayElements(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jboolean GNI_GetBooleanArrayElements(JNIEnv* p0, jbooleanArray p1, jboolean* p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jbooleanArray p1;
		jboolean* p2;
		jboolean r0;
		char __pad0[7];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetBooleanArrayElements, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetByteArrayElements(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jbyte GNI_GetByteArrayElements(JNIEnv* p0, jbyteArray p1, jboolean* p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jbyteArray p1;
		jboolean* p2;
		jbyte r0;
		char __pad0[7];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetByteArrayElements, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetCharArrayElements(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jchar GNI_GetCharArrayElements(JNIEnv* p0, jcharArray p1, jboolean* p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jcharArray p1;
		jboolean* p2;
		jchar r0;
		char __pad0[6];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetCharArrayElements, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetShortArrayElements(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jshort GNI_GetShortArrayElements(JNIEnv* p0, jshortArray p1, jboolean* p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jshortArray p1;
		jboolean* p2;
		jshort r0;
		char __pad0[6];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetShortArrayElements, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetIntArrayElements(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jint GNI_GetIntArrayElements(JNIEnv* p0, jintArray p1, jboolean* p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jintArray p1;
		jboolean* p2;
		jint r0;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetIntArrayElements, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetLongArrayElements(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jlong GNI_GetLongArrayElements(JNIEnv* p0, jlongArray p1, jboolean* p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jlongArray p1;
		jboolean* p2;
		jlong r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetLongArrayElements, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetFloatArrayElements(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jfloat GNI_GetFloatArrayElements(JNIEnv* p0, jfloatArray p1, jboolean* p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jfloatArray p1;
		jboolean* p2;
		jfloat r0;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetFloatArrayElements, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetDoubleArrayElements(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jdouble GNI_GetDoubleArrayElements(JNIEnv* p0, jdoubleArray p1, jboolean* p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jdoubleArray p1;
		jboolean* p2;
		jdouble r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetDoubleArrayElements, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_ReleaseBooleanArrayElements(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_ReleaseBooleanArrayElements(JNIEnv* p0, jbooleanArray p1, jboolean* p2, jint p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jbooleanArray p1;
		jboolean* p2;
		jint p3;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_ReleaseBooleanArrayElements, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_ReleaseByteArrayElements(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_ReleaseByteArrayElements(JNIEnv* p0, jbyteArray p1, jbyte* p2, jint p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jbyteArray p1;
		jbyte* p2;
		jint p3;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_ReleaseByteArrayElements, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_ReleaseCharArrayElements(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_ReleaseCharArrayElements(JNIEnv* p0, jcharArray p1, jchar* p2, jint p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jcharArray p1;
		jchar* p2;
		jint p3;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_ReleaseCharArrayElements, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_ReleaseShortArrayElements(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_ReleaseShortArrayElements(JNIEnv* p0, jshortArray p1, jshort* p2, jint p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jshortArray p1;
		jshort* p2;
		jint p3;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_ReleaseShortArrayElements, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_ReleaseIntArrayElements(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_ReleaseIntArrayElements(JNIEnv* p0, jintArray p1, jint* p2, jint p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jintArray p1;
		jint* p2;
		jint p3;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_ReleaseIntArrayElements, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_ReleaseLongArrayElements(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_ReleaseLongArrayElements(JNIEnv* p0, jlongArray p1, jlong* p2, jint p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jlongArray p1;
		jlong* p2;
		jint p3;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_ReleaseLongArrayElements, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_ReleaseFloatArrayElements(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_ReleaseFloatArrayElements(JNIEnv* p0, jfloatArray p1, jfloat* p2, jint p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jfloatArray p1;
		jfloat* p2;
		jint p3;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_ReleaseFloatArrayElements, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_ReleaseDoubleArrayElements(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_ReleaseDoubleArrayElements(JNIEnv* p0, jdoubleArray p1, jdouble* p2, jint p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jdoubleArray p1;
		jdouble* p2;
		jint p3;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_ReleaseDoubleArrayElements, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_GetBooleanArrayRegion(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_GetBooleanArrayRegion(JNIEnv* p0, jbooleanArray p1, jsize p2, jsize p3, jboolean* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jbooleanArray p1;
		jsize p2;
		jsize p3;
		jboolean* p4;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetBooleanArrayRegion, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_GetByteArrayRegion(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_GetByteArrayRegion(JNIEnv* p0, jbyteArray p1, jsize p2, jsize p3, jbyte* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jbyteArray p1;
		jsize p2;
		jsize p3;
		jbyte* p4;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetByteArrayRegion, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_GetCharArrayRegion(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_GetCharArrayRegion(JNIEnv* p0, jcharArray p1, jsize p2, jsize p3, jchar* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jcharArray p1;
		jsize p2;
		jsize p3;
		jchar* p4;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetCharArrayRegion, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_GetShortArrayRegion(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_GetShortArrayRegion(JNIEnv* p0, jshortArray p1, jsize p2, jsize p3, jshort* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jshortArray p1;
		jsize p2;
		jsize p3;
		jshort* p4;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetShortArrayRegion, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_GetIntArrayRegion(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_GetIntArrayRegion(JNIEnv* p0, jintArray p1, jsize p2, jsize p3, jint* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jintArray p1;
		jsize p2;
		jsize p3;
		jint* p4;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetIntArrayRegion, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_GetLongArrayRegion(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_GetLongArrayRegion(JNIEnv* p0, jlongArray p1, jsize p2, jsize p3, jlong* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jlongArray p1;
		jsize p2;
		jsize p3;
		jlong* p4;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetLongArrayRegion, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_GetFloatArrayRegion(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_GetFloatArrayRegion(JNIEnv* p0, jfloatArray p1, jsize p2, jsize p3, jfloat* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jfloatArray p1;
		jsize p2;
		jsize p3;
		jfloat* p4;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetFloatArrayRegion, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_GetDoubleArrayRegion(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_GetDoubleArrayRegion(JNIEnv* p0, jdoubleArray p1, jsize p2, jsize p3, jdouble* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jdoubleArray p1;
		jsize p2;
		jsize p3;
		jdouble* p4;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetDoubleArrayRegion, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_SetBooleanArrayRegion(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_SetBooleanArrayRegion(JNIEnv* p0, jbooleanArray p1, jsize p2, jsize p3, jboolean* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jbooleanArray p1;
		jsize p2;
		jsize p3;
		jboolean* p4;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_SetBooleanArrayRegion, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_SetByteArrayRegion(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_SetByteArrayRegion(JNIEnv* p0, jbyteArray p1, jsize p2, jsize p3, jbyte* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jbyteArray p1;
		jsize p2;
		jsize p3;
		jbyte* p4;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_SetByteArrayRegion, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_SetCharArrayRegion(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_SetCharArrayRegion(JNIEnv* p0, jcharArray p1, jsize p2, jsize p3, jchar* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jcharArray p1;
		jsize p2;
		jsize p3;
		jchar* p4;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_SetCharArrayRegion, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_SetShortArrayRegion(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_SetShortArrayRegion(JNIEnv* p0, jshortArray p1, jsize p2, jsize p3, jshort* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jshortArray p1;
		jsize p2;
		jsize p3;
		jshort* p4;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_SetShortArrayRegion, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_SetIntArrayRegion(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_SetIntArrayRegion(JNIEnv* p0, jintArray p1, jsize p2, jsize p3, jint* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jintArray p1;
		jsize p2;
		jsize p3;
		jint* p4;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_SetIntArrayRegion, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_SetLongArrayRegion(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_SetLongArrayRegion(JNIEnv* p0, jlongArray p1, jsize p2, jsize p3, jlong* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jlongArray p1;
		jsize p2;
		jsize p3;
		jlong* p4;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_SetLongArrayRegion, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_SetFloatArrayRegion(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_SetFloatArrayRegion(JNIEnv* p0, jfloatArray p1, jsize p2, jsize p3, jfloat* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jfloatArray p1;
		jsize p2;
		jsize p3;
		jfloat* p4;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_SetFloatArrayRegion, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_SetDoubleArrayRegion(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_SetDoubleArrayRegion(JNIEnv* p0, jdoubleArray p1, jsize p2, jsize p3, jdouble* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jdoubleArray p1;
		jsize p2;
		jsize p3;
		jdouble* p4;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_SetDoubleArrayRegion, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_RegisterNatives(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jint GNI_RegisterNatives(JNIEnv* p0, jclass p1, JNINativeMethod* p2, jint p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		JNINativeMethod* p2;
		jint p3;
		char __pad0[4];
		jint r0;
		char __pad1[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_RegisterNatives, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_UnregisterNatives(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jint GNI_UnregisterNatives(JNIEnv* p0, jclass p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jclass p1;
		jint r0;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_UnregisterNatives, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_MonitorEnter(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jint GNI_MonitorEnter(JNIEnv* p0, jobject p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jint r0;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_MonitorEnter, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_MonitorExit(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jint GNI_MonitorExit(JNIEnv* p0, jobject p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jint r0;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_MonitorExit, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetJavaVM(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jint GNI_GetJavaVM(JNIEnv* p0, void* p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		void* p1;
		jint r0;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetJavaVM, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetStringRegion(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_GetStringRegion(JNIEnv* p0, jstring p1, jsize p2, jsize p3, jchar* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jstring p1;
		jsize p2;
		jsize p3;
		jchar* p4;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetStringRegion, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_GetStringUTFRegion(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_GetStringUTFRegion(JNIEnv* p0, jstring p1, jsize p2, jsize p3, char* p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jstring p1;
		jsize p2;
		jsize p3;
		char* p4;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetStringUTFRegion, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_GetPrimitiveArrayCritical(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_GetPrimitiveArrayCritical(JNIEnv* p0, jarray p1, jboolean* p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jarray p1;
		jboolean* p2;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetPrimitiveArrayCritical, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_ReleasePrimitiveArrayCritical(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_ReleasePrimitiveArrayCritical(JNIEnv* p0, jarray p1, void* p2, jint p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jarray p1;
		void* p2;
		jint p3;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_ReleasePrimitiveArrayCritical, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_GetStringCritical(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jchar* GNI_GetStringCritical(JNIEnv* p0, jstring p1, jboolean* p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jstring p1;
		jboolean* p2;
		jchar* r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetStringCritical, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_ReleaseStringCritical(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_ReleaseStringCritical(JNIEnv* p0, jstring p1, jchar* p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jstring p1;
		jchar* p2;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_ReleaseStringCritical, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_NewWeakGlobalRef(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jweak GNI_NewWeakGlobalRef(JNIEnv* p0, jobject p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jweak r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_NewWeakGlobalRef, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_DeleteWeakGlobalRef(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GNI_DeleteWeakGlobalRef(JNIEnv* p0, jweak p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jweak p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_DeleteWeakGlobalRef, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_cf0433f838ae_GNI_ExceptionCheck(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jboolean GNI_ExceptionCheck(JNIEnv* p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jboolean r0;
		char __pad0[7];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_ExceptionCheck, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_NewDirectByteBuffer(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jobject GNI_NewDirectByteBuffer(JNIEnv* p0, void* p1, jlong p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		void* p1;
		jlong p2;
		jobject r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_NewDirectByteBuffer, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetDirectBufferAddress(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void* GNI_GetDirectBufferAddress(JNIEnv* p0, jobject p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		void* r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetDirectBufferAddress, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetDirectBufferCapacity(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jlong GNI_GetDirectBufferCapacity(JNIEnv* p0, jobject p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jlong r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetDirectBufferCapacity, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_cf0433f838ae_GNI_GetObjectRefType(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
jobjectRefType GNI_GetObjectRefType(JNIEnv* p0, jobject p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		JNIEnv* p0;
		jobject p1;
		jobjectRefType r0;
		char __pad0[4];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_cf0433f838ae_GNI_GetObjectRefType, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
