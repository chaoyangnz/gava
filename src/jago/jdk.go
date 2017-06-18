package jago

import "strings"

/* ---------- JDK Native methods implementation ---*/

func Java_GVM_println(o jobject, s jobject) {
	println(s.toString())
}

func Java_GVM_toUpper(s jobject) jobject {
	return NewJavaLangString(strings.ToUpper(s.toString()))
}

func Java_java_lang_Object_registerNatives() {

}

func Java_java_lang_System_registerNatives() {

}

func Java_java_lang_Class_registerNatives()  {
	
}
