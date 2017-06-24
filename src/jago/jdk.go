package jago

import "strings"

/* ---------- JDK Native methods implementation ---*/

func Java_GVM_println(o ObjectRef, s ObjectRef) {
	println(s.toString())
}

func Java_GVM_toUpper(s ObjectRef) ObjectRef {
	return NewJavaLangString(strings.ToUpper(s.toString()))
}

func Java_java_lang_Object_registerNatives() {

}

func Java_java_lang_System_registerNatives() {

}

func Java_java_lang_Class_registerNatives()  {
	
}

func Java_java_lang_Class_getPrimitiveClass(jobject ObjectRef) ObjectRef {
	return NewJavaLangClass()
}

func Java_java_lang_Class_desiredAssertionStatus0(clazz ObjectRef) Int {
	return TRUE.ToInt()
}

func Java_java_lang_Float_floatToRawIntBits(f Float) Int {
	return 1
}

func Java_java_lang_Double_doubleToRawLongBits(d Double) Long {
	return 1
}
