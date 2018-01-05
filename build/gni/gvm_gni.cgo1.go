// Created by cgo - DO NOT EDIT

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:5
package gvm
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:13

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:12
import "unsafe"
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:16

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:15
func GVM_print(s java_lang_string) {
	println(s.toString())
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:23

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:22
func GNI_GetVersion(env *_Ctype_JNIEnv) _Ctype_jint {
	fatal("Not implemented")
	return _Ctype_jint(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:29

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:28
func GNI_DefineClass(env *_Ctype_JNIEnv, name *_Ctype_char, loader _Ctype_jobject, buf *_Ctype_jbyte, len _Ctype_jsize) _Ctype_jclass {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:35

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:34
func GNI_FindClass(env *_Ctype_JNIEnv, name *_Ctype_char) _Ctype_jclass {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:41

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:40
func GNI_FromReflectedMethod(env *_Ctype_JNIEnv, method _Ctype_jobject) _Ctype_jmethodID {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:47

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:46
func GNI_FromReflectedField(env *_Ctype_JNIEnv, field _Ctype_jobject) _Ctype_jfieldID {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:53

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:52
func GNI_ToReflectedMethod(env *_Ctype_JNIEnv, cls _Ctype_jclass, methodID _Ctype_jmethodID, isStatic _Ctype_jboolean) _Ctype_jobject {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:59

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:58
func GNI_GetSuperclass(env *_Ctype_JNIEnv, sub _Ctype_jclass) _Ctype_jclass {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:65

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:64
func GNI_IsAssignableFrom(env *_Ctype_JNIEnv, sub _Ctype_jclass, sup _Ctype_jclass) _Ctype_jboolean {
	fatal("Not implemented")
	return _Ctype_jboolean(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:71

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:70
func GNI_ToReflectedField(env *_Ctype_JNIEnv, cls _Ctype_jclass, fieldID _Ctype_jfieldID, isStatic _Ctype_jboolean) _Ctype_jobject {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:77

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:76
func GNI_Throw(env *_Ctype_JNIEnv, obj _Ctype_jthrowable) _Ctype_jint {
	fatal("Not implemented")
	return _Ctype_jint(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:83

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:82
func GNI_ThrowNew(env *_Ctype_JNIEnv, clazz _Ctype_jclass, msg *_Ctype_char) _Ctype_jint {
	fatal("Not implemented")
	return _Ctype_jint(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:89

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:88
func GNI_ExceptionOccurred(env *_Ctype_JNIEnv) _Ctype_jthrowable {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:95

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:94
func GNI_ExceptionDescribe(env *_Ctype_JNIEnv) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:100

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:99
func GNI_ExceptionClear(env *_Ctype_JNIEnv) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:105

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:104
func GNI_FatalError(env *_Ctype_JNIEnv, msg *_Ctype_char) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:110

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:109
func GNI_PushLocalFrame(env *_Ctype_JNIEnv, capacity _Ctype_jint) _Ctype_jint {
	fatal("Not implemented")
	return _Ctype_jint(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:116

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:115
func GNI_PopLocalFrame(env *_Ctype_JNIEnv, result _Ctype_jobject) _Ctype_jobject {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:122

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:121
func GNI_NewGlobalRef(env *_Ctype_JNIEnv, lobj _Ctype_jobject) _Ctype_jobject {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:128

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:127
func GNI_DeleteGlobalRef(env *_Ctype_JNIEnv, gref _Ctype_jobject) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:133

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:132
func GNI_DeleteLocalRef(env *_Ctype_JNIEnv, obj _Ctype_jobject) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:138

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:137
func GNI_IsSameObject(env *_Ctype_JNIEnv, obj1 _Ctype_jobject, obj2 _Ctype_jobject) _Ctype_jboolean {
	fatal("Not implemented")
	return _Ctype_jboolean(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:144

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:143
func GNI_NewLocalRef(env *_Ctype_JNIEnv, ref _Ctype_jobject) _Ctype_jobject {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:150

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:149
func GNI_EnsureLocalCapacity(env *_Ctype_JNIEnv, capacity _Ctype_jint) _Ctype_jint {
	fatal("Not implemented")
	return _Ctype_jint(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:156

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:155
func GNI_AllocObject(env *_Ctype_JNIEnv, clazz _Ctype_jclass) _Ctype_jobject {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:167

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:166
func GNI_NewObjectV(env *_Ctype_JNIEnv, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_va_list) _Ctype_jobject {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:173

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:172
func GNI_NewObjectA(env *_Ctype_JNIEnv, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_jvalue) _Ctype_jobject {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:179

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:178
func GNI_GetObjectClass(env *_Ctype_JNIEnv, obj _Ctype_jobject) _Ctype_jclass {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:185

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:184
func GNI_IsInstanceOf(env *_Ctype_JNIEnv, obj _Ctype_jobject, clazz _Ctype_jclass) _Ctype_jboolean {
	fatal("Not implemented")
	return _Ctype_jboolean(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:191

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:190
func GNI_GetMethodID(env *_Ctype_JNIEnv, clazz _Ctype_jclass, name *_Ctype_char, sig *_Ctype_char) _Ctype_jmethodID {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:202

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:201
func GNI_CallObjectMethodV(env *_Ctype_JNIEnv, obj _Ctype_jobject, methodID _Ctype_jmethodID, args *_Ctype_va_list) _Ctype_jobject {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:208

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:207
func GNI_CallObjectMethodA(env *_Ctype_JNIEnv, obj _Ctype_jobject, methodID _Ctype_jmethodID, args *_Ctype_jvalue) _Ctype_jobject {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:219

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:218
func GNI_CallBooleanMethodV(env *_Ctype_JNIEnv, obj _Ctype_jobject, methodID _Ctype_jmethodID, args *_Ctype_va_list) _Ctype_jboolean {
	fatal("Not implemented")
	return _Ctype_jboolean(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:225

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:224
func GNI_CallBooleanMethodA(env *_Ctype_JNIEnv, obj _Ctype_jobject, methodID _Ctype_jmethodID, args *_Ctype_jvalue) _Ctype_jboolean {
	fatal("Not implemented")
	return _Ctype_jboolean(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:236

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:235
func GNI_CallByteMethodV(env *_Ctype_JNIEnv, obj _Ctype_jobject, methodID _Ctype_jmethodID, args *_Ctype_va_list) _Ctype_jbyte {
	fatal("Not implemented")
	return _Ctype_jbyte(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:242

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:241
func GNI_CallByteMethodA(env *_Ctype_JNIEnv, obj _Ctype_jobject, methodID _Ctype_jmethodID, args *_Ctype_jvalue) _Ctype_jbyte {
	fatal("Not implemented")
	return _Ctype_jbyte(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:253

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:252
func GNI_CallCharMethodV(env *_Ctype_JNIEnv, obj _Ctype_jobject, methodID _Ctype_jmethodID, args *_Ctype_va_list) _Ctype_jchar {
	fatal("Not implemented")
	return _Ctype_jchar(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:259

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:258
func GNI_CallCharMethodA(env *_Ctype_JNIEnv, obj _Ctype_jobject, methodID _Ctype_jmethodID, args *_Ctype_jvalue) _Ctype_jchar {
	fatal("Not implemented")
	return _Ctype_jchar(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:270

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:269
func GNI_CallShortMethodV(env *_Ctype_JNIEnv, obj _Ctype_jobject, methodID _Ctype_jmethodID, args *_Ctype_va_list) _Ctype_jshort {
	fatal("Not implemented")
	return _Ctype_jshort(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:276

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:275
func GNI_CallShortMethodA(env *_Ctype_JNIEnv, obj _Ctype_jobject, methodID _Ctype_jmethodID, args *_Ctype_jvalue) _Ctype_jshort {
	fatal("Not implemented")
	return _Ctype_jshort(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:287

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:286
func GNI_CallIntMethodV(env *_Ctype_JNIEnv, obj _Ctype_jobject, methodID _Ctype_jmethodID, args *_Ctype_va_list) _Ctype_jint {
	fatal("Not implemented")
	return _Ctype_jint(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:293

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:292
func GNI_CallIntMethodA(env *_Ctype_JNIEnv, obj _Ctype_jobject, methodID _Ctype_jmethodID, args *_Ctype_jvalue) _Ctype_jint {
	fatal("Not implemented")
	return _Ctype_jint(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:304

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:303
func GNI_CallLongMethodV(env *_Ctype_JNIEnv, obj _Ctype_jobject, methodID _Ctype_jmethodID, args *_Ctype_va_list) _Ctype_jlong {
	fatal("Not implemented")
	return _Ctype_jlong(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:310

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:309
func GNI_CallLongMethodA(env *_Ctype_JNIEnv, obj _Ctype_jobject, methodID _Ctype_jmethodID, args *_Ctype_jvalue) _Ctype_jlong {
	fatal("Not implemented")
	return _Ctype_jlong(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:321

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:320
func GNI_CallFloatMethodV(env *_Ctype_JNIEnv, obj _Ctype_jobject, methodID _Ctype_jmethodID, args *_Ctype_va_list) _Ctype_jfloat {
	fatal("Not implemented")
	return _Ctype_jfloat(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:327

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:326
func GNI_CallFloatMethodA(env *_Ctype_JNIEnv, obj _Ctype_jobject, methodID _Ctype_jmethodID, args *_Ctype_jvalue) _Ctype_jfloat {
	fatal("Not implemented")
	return _Ctype_jfloat(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:338

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:337
func GNI_CallDoubleMethodV(env *_Ctype_JNIEnv, obj _Ctype_jobject, methodID _Ctype_jmethodID, args *_Ctype_va_list) _Ctype_jdouble {
	fatal("Not implemented")
	return _Ctype_jdouble(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:344

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:343
func GNI_CallDoubleMethodA(env *_Ctype_JNIEnv, obj _Ctype_jobject, methodID _Ctype_jmethodID, args *_Ctype_jvalue) _Ctype_jdouble {
	fatal("Not implemented")
	return _Ctype_jdouble(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:354

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:353
func GNI_CallVoidMethodV(env *_Ctype_JNIEnv, obj _Ctype_jobject, methodID _Ctype_jmethodID, args *_Ctype_va_list) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:359

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:358
func GNI_CallVoidMethodA(env *_Ctype_JNIEnv, obj _Ctype_jobject, methodID _Ctype_jmethodID, args *_Ctype_jvalue) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:369

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:368
func GNI_CallNonvirtualObjectMethodV(env *_Ctype_JNIEnv, obj _Ctype_jobject, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_va_list) _Ctype_jobject {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:375

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:374
func GNI_CallNonvirtualObjectMethodA(env *_Ctype_JNIEnv, obj _Ctype_jobject, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_jvalue) _Ctype_jobject {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:386

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:385
func GNI_CallNonvirtualBooleanMethodV(env *_Ctype_JNIEnv, obj _Ctype_jobject, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_va_list) _Ctype_jboolean {
	fatal("Not implemented")
	return _Ctype_jboolean(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:392

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:391
func GNI_CallNonvirtualBooleanMethodA(env *_Ctype_JNIEnv, obj _Ctype_jobject, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_jvalue) _Ctype_jboolean {
	fatal("Not implemented")
	return _Ctype_jboolean(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:403

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:402
func GNI_CallNonvirtualByteMethodV(env *_Ctype_JNIEnv, obj _Ctype_jobject, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_va_list) _Ctype_jbyte {
	fatal("Not implemented")
	return _Ctype_jbyte(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:409

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:408
func GNI_CallNonvirtualByteMethodA(env *_Ctype_JNIEnv, obj _Ctype_jobject, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_jvalue) _Ctype_jbyte {
	fatal("Not implemented")
	return _Ctype_jbyte(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:420

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:419
func GNI_CallNonvirtualCharMethodV(env *_Ctype_JNIEnv, obj _Ctype_jobject, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_va_list) _Ctype_jchar {
	fatal("Not implemented")
	return _Ctype_jchar(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:426

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:425
func GNI_CallNonvirtualCharMethodA(env *_Ctype_JNIEnv, obj _Ctype_jobject, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_jvalue) _Ctype_jchar {
	fatal("Not implemented")
	return _Ctype_jchar(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:437

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:436
func GNI_CallNonvirtualShortMethodV(env *_Ctype_JNIEnv, obj _Ctype_jobject, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_va_list) _Ctype_jshort {
	fatal("Not implemented")
	return _Ctype_jshort(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:443

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:442
func GNI_CallNonvirtualShortMethodA(env *_Ctype_JNIEnv, obj _Ctype_jobject, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_jvalue) _Ctype_jshort {
	fatal("Not implemented")
	return _Ctype_jshort(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:454

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:453
func GNI_CallNonvirtualIntMethodV(env *_Ctype_JNIEnv, obj _Ctype_jobject, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_va_list) _Ctype_jint {
	fatal("Not implemented")
	return _Ctype_jint(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:460

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:459
func GNI_CallNonvirtualIntMethodA(env *_Ctype_JNIEnv, obj _Ctype_jobject, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_jvalue) _Ctype_jint {
	fatal("Not implemented")
	return _Ctype_jint(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:471

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:470
func GNI_CallNonvirtualLongMethodV(env *_Ctype_JNIEnv, obj _Ctype_jobject, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_va_list) _Ctype_jlong {
	fatal("Not implemented")
	return _Ctype_jlong(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:477

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:476
func GNI_CallNonvirtualLongMethodA(env *_Ctype_JNIEnv, obj _Ctype_jobject, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_jvalue) _Ctype_jlong {
	fatal("Not implemented")
	return _Ctype_jlong(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:488

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:487
func GNI_CallNonvirtualFloatMethodV(env *_Ctype_JNIEnv, obj _Ctype_jobject, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_va_list) _Ctype_jfloat {
	fatal("Not implemented")
	return _Ctype_jfloat(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:494

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:493
func GNI_CallNonvirtualFloatMethodA(env *_Ctype_JNIEnv, obj _Ctype_jobject, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_jvalue) _Ctype_jfloat {
	fatal("Not implemented")
	return _Ctype_jfloat(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:505

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:504
func GNI_CallNonvirtualDoubleMethodV(env *_Ctype_JNIEnv, obj _Ctype_jobject, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_va_list) _Ctype_jdouble {
	fatal("Not implemented")
	return _Ctype_jdouble(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:511

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:510
func GNI_CallNonvirtualDoubleMethodA(env *_Ctype_JNIEnv, obj _Ctype_jobject, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_jvalue) _Ctype_jdouble {
	fatal("Not implemented")
	return _Ctype_jdouble(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:521

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:520
func GNI_CallNonvirtualVoidMethodV(env *_Ctype_JNIEnv, obj _Ctype_jobject, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_va_list) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:526

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:525
func GNI_CallNonvirtualVoidMethodA(env *_Ctype_JNIEnv, obj _Ctype_jobject, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_jvalue) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:531

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:530
func GNI_GetFieldID(env *_Ctype_JNIEnv, clazz _Ctype_jclass, name *_Ctype_char, sig *_Ctype_char) _Ctype_jfieldID {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:537

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:536
func GNI_GetObjectField(env *_Ctype_JNIEnv, obj _Ctype_jobject, fieldID _Ctype_jfieldID) _Ctype_jobject {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:543

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:542
func GNI_GetBooleanField(env *_Ctype_JNIEnv, obj _Ctype_jobject, fieldID _Ctype_jfieldID) _Ctype_jboolean {
	fatal("Not implemented")
	return _Ctype_jboolean(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:549

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:548
func GNI_GetByteField(env *_Ctype_JNIEnv, obj _Ctype_jobject, fieldID _Ctype_jfieldID) _Ctype_jbyte {
	fatal("Not implemented")
	return _Ctype_jbyte(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:555

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:554
func GNI_GetCharField(env *_Ctype_JNIEnv, obj _Ctype_jobject, fieldID _Ctype_jfieldID) _Ctype_jchar {
	fatal("Not implemented")
	return _Ctype_jchar(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:561

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:560
func GNI_GetShortField(env *_Ctype_JNIEnv, obj _Ctype_jobject, fieldID _Ctype_jfieldID) _Ctype_jshort {
	fatal("Not implemented")
	return _Ctype_jshort(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:567

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:566
func GNI_GetIntField(env *_Ctype_JNIEnv, obj _Ctype_jobject, fieldID _Ctype_jfieldID) _Ctype_jint {
	fatal("Not implemented")
	return _Ctype_jint(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:573

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:572
func GNI_GetLongField(env *_Ctype_JNIEnv, obj _Ctype_jobject, fieldID _Ctype_jfieldID) _Ctype_jlong {
	fatal("Not implemented")
	return _Ctype_jlong(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:579

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:578
func GNI_GetFloatField(env *_Ctype_JNIEnv, obj _Ctype_jobject, fieldID _Ctype_jfieldID) _Ctype_jfloat {
	fatal("Not implemented")
	return _Ctype_jfloat(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:585

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:584
func GNI_GetDoubleField(env *_Ctype_JNIEnv, obj _Ctype_jobject, fieldID _Ctype_jfieldID) _Ctype_jdouble {
	fatal("Not implemented")
	return _Ctype_jdouble(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:591

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:590
func GNI_SetObjectField(env *_Ctype_JNIEnv, obj _Ctype_jobject, fieldID _Ctype_jfieldID, val _Ctype_jobject) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:596

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:595
func GNI_SetBooleanField(env *_Ctype_JNIEnv, obj _Ctype_jobject, fieldID _Ctype_jfieldID, val _Ctype_jboolean) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:601

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:600
func GNI_SetByteField(env *_Ctype_JNIEnv, obj _Ctype_jobject, fieldID _Ctype_jfieldID, val _Ctype_jbyte) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:606

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:605
func GNI_SetCharField(env *_Ctype_JNIEnv, obj _Ctype_jobject, fieldID _Ctype_jfieldID, val _Ctype_jchar) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:611

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:610
func GNI_SetShortField(env *_Ctype_JNIEnv, obj _Ctype_jobject, fieldID _Ctype_jfieldID, val _Ctype_jshort) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:616

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:615
func GNI_SetIntField(env *_Ctype_JNIEnv, obj _Ctype_jobject, fieldID _Ctype_jfieldID, val _Ctype_jint) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:621

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:620
func GNI_SetLongField(env *_Ctype_JNIEnv, obj _Ctype_jobject, fieldID _Ctype_jfieldID, val _Ctype_jlong) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:626

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:625
func GNI_SetFloatField(env *_Ctype_JNIEnv, obj _Ctype_jobject, fieldID _Ctype_jfieldID, val _Ctype_jfloat) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:631

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:630
func GNI_SetDoubleField(env *_Ctype_JNIEnv, obj _Ctype_jobject, fieldID _Ctype_jfieldID, val _Ctype_jdouble) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:636

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:635
func GNI_GetStaticMethodID(env *_Ctype_JNIEnv, clazz _Ctype_jclass, name *_Ctype_char, sig *_Ctype_char) _Ctype_jmethodID {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:647

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:646
func GNI_CallStaticObjectMethodV(env *_Ctype_JNIEnv, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_va_list) _Ctype_jobject {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:653

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:652
func GNI_CallStaticObjectMethodA(env *_Ctype_JNIEnv, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_jvalue) _Ctype_jobject {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:664

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:663
func GNI_CallStaticBooleanMethodV(env *_Ctype_JNIEnv, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_va_list) _Ctype_jboolean {
	fatal("Not implemented")
	return _Ctype_jboolean(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:670

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:669
func GNI_CallStaticBooleanMethodA(env *_Ctype_JNIEnv, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_jvalue) _Ctype_jboolean {
	fatal("Not implemented")
	return _Ctype_jboolean(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:681

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:680
func GNI_CallStaticByteMethodV(env *_Ctype_JNIEnv, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_va_list) _Ctype_jbyte {
	fatal("Not implemented")
	return _Ctype_jbyte(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:687

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:686
func GNI_CallStaticByteMethodA(env *_Ctype_JNIEnv, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_jvalue) _Ctype_jbyte {
	fatal("Not implemented")
	return _Ctype_jbyte(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:698

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:697
func GNI_CallStaticCharMethodV(env *_Ctype_JNIEnv, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_va_list) _Ctype_jchar {
	fatal("Not implemented")
	return _Ctype_jchar(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:704

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:703
func GNI_CallStaticCharMethodA(env *_Ctype_JNIEnv, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_jvalue) _Ctype_jchar {
	fatal("Not implemented")
	return _Ctype_jchar(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:720

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:719
func GNI_CallStaticShortMethodA(env *_Ctype_JNIEnv, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_jvalue) _Ctype_jshort {
	fatal("Not implemented")
	return _Ctype_jshort(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:731

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:730
func GNI_CallStaticIntMethodV(env *_Ctype_JNIEnv, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_va_list) _Ctype_jint {
	fatal("Not implemented")
	return _Ctype_jint(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:737

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:736
func GNI_CallStaticIntMethodA(env *_Ctype_JNIEnv, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_jvalue) _Ctype_jint {
	fatal("Not implemented")
	return _Ctype_jint(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:748

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:747
func GNI_CallStaticLongMethodV(env *_Ctype_JNIEnv, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_va_list) _Ctype_jlong {
	fatal("Not implemented")
	return _Ctype_jlong(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:754

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:753
func GNI_CallStaticLongMethodA(env *_Ctype_JNIEnv, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_jvalue) _Ctype_jlong {
	fatal("Not implemented")
	return _Ctype_jlong(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:765

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:764
func GNI_CallStaticFloatMethodV(env *_Ctype_JNIEnv, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_va_list) _Ctype_jfloat {
	fatal("Not implemented")
	return _Ctype_jfloat(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:771

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:770
func GNI_CallStaticFloatMethodA(env *_Ctype_JNIEnv, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_jvalue) _Ctype_jfloat {
	fatal("Not implemented")
	return _Ctype_jfloat(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:782

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:781
func GNI_CallStaticDoubleMethodV(env *_Ctype_JNIEnv, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_va_list) _Ctype_jdouble {
	fatal("Not implemented")
	return _Ctype_jdouble(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:788

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:787
func GNI_CallStaticDoubleMethodA(env *_Ctype_JNIEnv, clazz _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_jvalue) _Ctype_jdouble {
	fatal("Not implemented")
	return _Ctype_jdouble(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:798

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:797
func GNI_CallStaticVoidMethodV(env *_Ctype_JNIEnv, cls _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_va_list) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:803

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:802
func GNI_CallStaticVoidMethodA(env *_Ctype_JNIEnv, cls _Ctype_jclass, methodID _Ctype_jmethodID, args *_Ctype_jvalue) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:808

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:807
func GNI_GetStaticFieldID(env *_Ctype_JNIEnv, clazz _Ctype_jclass, name *_Ctype_char, sig *_Ctype_char) _Ctype_jfieldID {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:814

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:813
func GNI_GetStaticObjectField(env *_Ctype_JNIEnv, clazz _Ctype_jclass, fieldID _Ctype_jfieldID) _Ctype_jobject {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:820

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:819
func GNI_GetStaticBooleanField(env *_Ctype_JNIEnv, clazz _Ctype_jclass, fieldID _Ctype_jfieldID) _Ctype_jboolean {
	fatal("Not implemented")
	return _Ctype_jboolean(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:826

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:825
func GNI_GetStaticByteField(env *_Ctype_JNIEnv, clazz _Ctype_jclass, fieldID _Ctype_jfieldID) _Ctype_jbyte {
	fatal("Not implemented")
	return _Ctype_jbyte(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:832

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:831
func GNI_GetStaticCharField(env *_Ctype_JNIEnv, clazz _Ctype_jclass, fieldID _Ctype_jfieldID) _Ctype_jchar {
	fatal("Not implemented")
	return _Ctype_jchar(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:838

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:837
func GNI_GetStaticShortField(env *_Ctype_JNIEnv, clazz _Ctype_jclass, fieldID _Ctype_jfieldID) _Ctype_jshort {
	fatal("Not implemented")
	return _Ctype_jshort(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:844

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:843
func GNI_GetStaticIntField(env *_Ctype_JNIEnv, clazz _Ctype_jclass, fieldID _Ctype_jfieldID) _Ctype_jint {
	fatal("Not implemented")
	return _Ctype_jint(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:850

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:849
func GNI_GetStaticLongField(env *_Ctype_JNIEnv, clazz _Ctype_jclass, fieldID _Ctype_jfieldID) _Ctype_jlong {
	fatal("Not implemented")
	return _Ctype_jlong(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:856

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:855
func GNI_GetStaticFloatField(env *_Ctype_JNIEnv, clazz _Ctype_jclass, fieldID _Ctype_jfieldID) _Ctype_jfloat {
	fatal("Not implemented")
	return _Ctype_jfloat(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:862

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:861
func GNI_GetStaticDoubleField(env *_Ctype_JNIEnv, clazz _Ctype_jclass, fieldID _Ctype_jfieldID) _Ctype_jdouble {
	fatal("Not implemented")
	return _Ctype_jdouble(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:868

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:867
func GNI_SetStaticObjectField(env *_Ctype_JNIEnv, clazz _Ctype_jclass, fieldID _Ctype_jfieldID, value _Ctype_jobject) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:873

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:872
func GNI_SetStaticBooleanField(env *_Ctype_JNIEnv, clazz _Ctype_jclass, fieldID _Ctype_jfieldID, value _Ctype_jboolean) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:878

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:877
func GNI_SetStaticByteField(env *_Ctype_JNIEnv, clazz _Ctype_jclass, fieldID _Ctype_jfieldID, value _Ctype_jbyte) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:883

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:882
func GNI_SetStaticCharField(env *_Ctype_JNIEnv, clazz _Ctype_jclass, fieldID _Ctype_jfieldID, value _Ctype_jchar) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:888

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:887
func GNI_SetStaticShortField(env *_Ctype_JNIEnv, clazz _Ctype_jclass, fieldID _Ctype_jfieldID, value _Ctype_jshort) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:893

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:892
func GNI_SetStaticIntField(env *_Ctype_JNIEnv, clazz _Ctype_jclass, fieldID _Ctype_jfieldID, value _Ctype_jint) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:898

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:897
func GNI_SetStaticLongField(env *_Ctype_JNIEnv, clazz _Ctype_jclass, fieldID _Ctype_jfieldID, value _Ctype_jlong) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:903

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:902
func GNI_SetStaticFloatField(env *_Ctype_JNIEnv, clazz _Ctype_jclass, fieldID _Ctype_jfieldID, value _Ctype_jfloat) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:908

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:907
func GNI_SetStaticDoubleField(env *_Ctype_JNIEnv, clazz _Ctype_jclass, fieldID _Ctype_jfieldID, value _Ctype_jdouble) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:913

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:912
func GNI_NewString(env *_Ctype_JNIEnv, unicode *_Ctype_jchar, len _Ctype_jsize) _Ctype_jstring {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:919

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:918
func GNI_GetStringLength(env *_Ctype_JNIEnv, str _Ctype_jstring) _Ctype_jsize {
	fatal("Not implemented")
	return _Ctype_jsize(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:925

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:924
func GNI_GetStringChars(env *_Ctype_JNIEnv, str _Ctype_jstring, isCopy *_Ctype_jboolean) *_Ctype_jchar {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:931

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:930
func GNI_ReleaseStringChars(env *_Ctype_JNIEnv, str _Ctype_jstring, chars _Ctype_jchar) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:936

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:935
func GNI_NewStringUTF(env *_Ctype_JNIEnv, utf *_Ctype_char) _Ctype_jstring {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:942

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:941
func GNI_GetStringUTFLength(env *_Ctype_JNIEnv, str _Ctype_jstring) _Ctype_jsize {
	fatal("Not implemented")
	return _Ctype_jsize(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:948

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:947
func GNI_GetStringUTFChars(env *_Ctype_JNIEnv, str _Ctype_jstring, isCopy *_Ctype_jboolean) *_Ctype_char {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:954

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:953
func GNI_ReleaseStringUTFChars(env *_Ctype_JNIEnv, str _Ctype_jstring, chars *_Ctype_char) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:959

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:958
func GNI_GetArrayLength(env *_Ctype_JNIEnv, array *_Ctype_jarray) _Ctype_jsize {
	fatal("Not implemented")
	return _Ctype_jsize(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:965

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:964
func GNI_NewObjectArray(env *_Ctype_JNIEnv, len _Ctype_jsize, clazz _Ctype_jclass, init _Ctype_jobject) _Ctype_jobjectArray {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:971

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:970
func GNI_GetObjectArrayElement(env *_Ctype_JNIEnv, array _Ctype_jobjectArray, index _Ctype_jsize) _Ctype_jobject {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:977

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:976
func GNI_SetObjectArrayElement(env *_Ctype_JNIEnv, array _Ctype_jobjectArray, index _Ctype_jsize, val _Ctype_jobject) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:982

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:981
func GNI_NewBooleanArray(env *_Ctype_JNIEnv, len _Ctype_jsize) _Ctype_jbooleanArray {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:988

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:987
func GNI_NewByteArray(env *_Ctype_JNIEnv, len _Ctype_jsize) _Ctype_jbyteArray {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:994

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:993
func GNI_NewCharArray(env *_Ctype_JNIEnv, len _Ctype_jsize) _Ctype_jcharArray {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1000

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:999
func GNI_NewShortArray(env *_Ctype_JNIEnv, len _Ctype_jsize) _Ctype_jshortArray {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1006

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1005
func GNI_NewIntArray(env *_Ctype_JNIEnv, len _Ctype_jsize) _Ctype_jintArray {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1012

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1011
func GNI_NewLongArray(env *_Ctype_JNIEnv, len _Ctype_jsize) _Ctype_jlongArray {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1018

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1017
func GNI_NewFloatArray(env *_Ctype_JNIEnv, len _Ctype_jsize) _Ctype_jfloatArray {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1024

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1023
func GNI_NewDoubleArray(env *_Ctype_JNIEnv, len _Ctype_jsize) _Ctype_jdoubleArray {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1030

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1029
func GNI_GetBooleanArrayElements(env *_Ctype_JNIEnv, array _Ctype_jbooleanArray, isCopy *_Ctype_jboolean) _Ctype_jboolean {
	fatal("Not implemented")
	return _Ctype_jboolean(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1036

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1035
func GNI_GetByteArrayElements(env *_Ctype_JNIEnv, array _Ctype_jbyteArray, isCopy *_Ctype_jboolean) _Ctype_jbyte {
	fatal("Not implemented")
	return _Ctype_jbyte(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1042

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1041
func GNI_GetCharArrayElements(env *_Ctype_JNIEnv, array _Ctype_jcharArray, isCopy *_Ctype_jboolean) _Ctype_jchar {
	fatal("Not implemented")
	return _Ctype_jchar(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1048

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1047
func GNI_GetShortArrayElements(env *_Ctype_JNIEnv, array _Ctype_jshortArray, isCopy *_Ctype_jboolean) _Ctype_jshort {
	fatal("Not implemented")
	return _Ctype_jshort(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1054

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1053
func GNI_GetIntArrayElements(env *_Ctype_JNIEnv, array _Ctype_jintArray, isCopy *_Ctype_jboolean) _Ctype_jint {
	fatal("Not implemented")
	return _Ctype_jint(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1060

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1059
func GNI_GetLongArrayElements(env *_Ctype_JNIEnv, array _Ctype_jlongArray, isCopy *_Ctype_jboolean) _Ctype_jlong {
	fatal("Not implemented")
	return _Ctype_jlong(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1066

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1065
func GNI_GetFloatArrayElements(env *_Ctype_JNIEnv, array _Ctype_jfloatArray, isCopy *_Ctype_jboolean) _Ctype_jfloat {
	fatal("Not implemented")
	return _Ctype_jfloat(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1072

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1071
func GNI_GetDoubleArrayElements(env *_Ctype_JNIEnv, array _Ctype_jdoubleArray, isCopy *_Ctype_jboolean) _Ctype_jdouble {
	fatal("Not implemented")
	return _Ctype_jdouble(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1078

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1077
func GNI_ReleaseBooleanArrayElements(env *_Ctype_JNIEnv, array _Ctype_jbooleanArray, elems *_Ctype_jboolean, mode _Ctype_jint) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1083

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1082
func GNI_ReleaseByteArrayElements(env *_Ctype_JNIEnv, array _Ctype_jbyteArray, elems *_Ctype_jbyte, mode _Ctype_jint) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1088

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1087
func GNI_ReleaseCharArrayElements(env *_Ctype_JNIEnv, array _Ctype_jcharArray, elems *_Ctype_jchar, mode _Ctype_jint) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1093

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1092
func GNI_ReleaseShortArrayElements(env *_Ctype_JNIEnv, array _Ctype_jshortArray, elems *_Ctype_jshort, mode _Ctype_jint) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1098

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1097
func GNI_ReleaseIntArrayElements(env *_Ctype_JNIEnv, array _Ctype_jintArray, elems *_Ctype_jint, mode _Ctype_jint) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1103

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1102
func GNI_ReleaseLongArrayElements(env *_Ctype_JNIEnv, array _Ctype_jlongArray, elems *_Ctype_jlong, mode _Ctype_jint) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1108

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1107
func GNI_ReleaseFloatArrayElements(env *_Ctype_JNIEnv, array _Ctype_jfloatArray, elems *_Ctype_jfloat, mode _Ctype_jint) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1113

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1112
func GNI_ReleaseDoubleArrayElements(env *_Ctype_JNIEnv, array _Ctype_jdoubleArray, elems *_Ctype_jdouble, mode _Ctype_jint) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1118

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1117
func GNI_GetBooleanArrayRegion(env *_Ctype_JNIEnv, array _Ctype_jbooleanArray, start _Ctype_jsize, l _Ctype_jsize, buf *_Ctype_jboolean) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1123

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1122
func GNI_GetByteArrayRegion(env *_Ctype_JNIEnv, array _Ctype_jbyteArray, start _Ctype_jsize, len _Ctype_jsize, buf *_Ctype_jbyte) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1128

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1127
func GNI_GetCharArrayRegion(env *_Ctype_JNIEnv, array _Ctype_jcharArray, start _Ctype_jsize, len _Ctype_jsize, buf *_Ctype_jchar) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1133

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1132
func GNI_GetShortArrayRegion(env *_Ctype_JNIEnv, array _Ctype_jshortArray, start _Ctype_jsize, len _Ctype_jsize, buf *_Ctype_jshort) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1138

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1137
func GNI_GetIntArrayRegion(env *_Ctype_JNIEnv, array _Ctype_jintArray, start _Ctype_jsize, len _Ctype_jsize, buf *_Ctype_jint) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1143

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1142
func GNI_GetLongArrayRegion(env *_Ctype_JNIEnv, array _Ctype_jlongArray, start _Ctype_jsize, len _Ctype_jsize, buf *_Ctype_jlong) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1148

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1147
func GNI_GetFloatArrayRegion(env *_Ctype_JNIEnv, array _Ctype_jfloatArray, start _Ctype_jsize, len _Ctype_jsize, buf *_Ctype_jfloat) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1153

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1152
func GNI_GetDoubleArrayRegion(env *_Ctype_JNIEnv, array _Ctype_jdoubleArray, start _Ctype_jsize, len _Ctype_jsize, buf *_Ctype_jdouble) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1158

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1157
func GNI_SetBooleanArrayRegion(env *_Ctype_JNIEnv, array _Ctype_jbooleanArray, start _Ctype_jsize, l _Ctype_jsize, buf *_Ctype_jboolean) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1163

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1162
func GNI_SetByteArrayRegion(env *_Ctype_JNIEnv, array _Ctype_jbyteArray, start _Ctype_jsize, len _Ctype_jsize, buf *_Ctype_jbyte) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1168

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1167
func GNI_SetCharArrayRegion(env *_Ctype_JNIEnv, array _Ctype_jcharArray, start _Ctype_jsize, len _Ctype_jsize, buf *_Ctype_jchar) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1173

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1172
func GNI_SetShortArrayRegion(env *_Ctype_JNIEnv, array _Ctype_jshortArray, start _Ctype_jsize, len _Ctype_jsize, buf *_Ctype_jshort) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1178

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1177
func GNI_SetIntArrayRegion(env *_Ctype_JNIEnv, array _Ctype_jintArray, start _Ctype_jsize, len _Ctype_jsize, buf *_Ctype_jint) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1183

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1182
func GNI_SetLongArrayRegion(env *_Ctype_JNIEnv, array _Ctype_jlongArray, start _Ctype_jsize, len _Ctype_jsize, buf *_Ctype_jlong) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1188

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1187
func GNI_SetFloatArrayRegion(env *_Ctype_JNIEnv, array _Ctype_jfloatArray, start _Ctype_jsize, len _Ctype_jsize, buf *_Ctype_jfloat) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1193

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1192
func GNI_SetDoubleArrayRegion(env *_Ctype_JNIEnv, array _Ctype_jdoubleArray, start _Ctype_jsize, len _Ctype_jsize, buf *_Ctype_jdouble) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1198

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1197
func GNI_RegisterNatives(env *_Ctype_JNIEnv, clazz _Ctype_jclass, methods *_Ctype_struct___0, nMethods _Ctype_jint) _Ctype_jint {
	fatal("Not implemented")
	return _Ctype_jint(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1204

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1203
func GNI_UnregisterNatives(env *_Ctype_JNIEnv, clazz _Ctype_jclass) _Ctype_jint {
	fatal("Not implemented")
	return _Ctype_jint(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1210

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1209
func GNI_MonitorEnter(env *_Ctype_JNIEnv, obj _Ctype_jobject) _Ctype_jint {
	fatal("Not implemented")
	return _Ctype_jint(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1216

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1215
func GNI_MonitorExit(env *_Ctype_JNIEnv, obj _Ctype_jobject) _Ctype_jint {
	fatal("Not implemented")
	return _Ctype_jint(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1222

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1221
func GNI_GetJavaVM(env *_Ctype_JNIEnv, vm unsafe.Pointer) _Ctype_jint {
	fatal("Not implemented")
	return _Ctype_jint(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1228

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1227
func GNI_GetStringRegion(env *_Ctype_JNIEnv, str _Ctype_jstring, start _Ctype_jsize, len _Ctype_jsize, buf *_Ctype_jchar) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1233

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1232
func GNI_GetStringUTFRegion(env *_Ctype_JNIEnv, str _Ctype_jstring, start _Ctype_jsize, len _Ctype_jsize, buf *_Ctype_char) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1238

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1237
func GNI_GetPrimitiveArrayCritical(env *_Ctype_JNIEnv, array _Ctype_jarray, isCopy *_Ctype_jboolean) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1243

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1242
func GNI_ReleasePrimitiveArrayCritical(env *_Ctype_JNIEnv, array _Ctype_jarray, carray unsafe.Pointer, mode _Ctype_jint) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1248

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1247
func GNI_GetStringCritical(env *_Ctype_JNIEnv, string _Ctype_jstring, isCopy *_Ctype_jboolean) *_Ctype_jchar {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1254

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1253
func GNI_ReleaseStringCritical(env *_Ctype_JNIEnv, string _Ctype_jstring, cstring *_Ctype_jchar) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1259

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1258
func GNI_NewWeakGlobalRef(env *_Ctype_JNIEnv, obj _Ctype_jobject) _Ctype_jweak {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1265

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1264
func GNI_DeleteWeakGlobalRef(env *_Ctype_JNIEnv, ref _Ctype_jweak) {
	fatal("Not implemented")
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1270

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1269
func GNI_ExceptionCheck(env *_Ctype_JNIEnv) _Ctype_jboolean {
	fatal("Not implemented")
	return _Ctype_jboolean(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1276

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1275
func GNI_NewDirectByteBuffer(env *_Ctype_JNIEnv, address unsafe.Pointer, capacity _Ctype_jlong) _Ctype_jobject {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1282

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1281
func GNI_GetDirectBufferAddress(env *_Ctype_JNIEnv, buf _Ctype_jobject) unsafe.Pointer {
	fatal("Not implemented")
	return nil
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1288

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1287
func GNI_GetDirectBufferCapacity(env *_Ctype_JNIEnv, buf _Ctype_jobject) _Ctype_jlong {
	fatal("Not implemented")
	return _Ctype_jlong(0)
}
//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1295

//line /Users/Chao/GoglandProjects/gvm/src/gvm/gni.go:1294
func GNI_GetObjectRefType(env *_Ctype_JNIEnv, obj _Ctype_jobject) _Ctype_jobjectRefType {
	fatal("Not implemented")
	return _Ctype_jobjectRefType(0)
}
