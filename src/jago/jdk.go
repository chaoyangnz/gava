package jago

import "strings"

/* ---------- JDK Native methods implementation ---*/

func Java_GVM_println(o jobject, s JavaLangString) {
	println(s.toString())
}

func Java_GVM_toUpper(s JavaLangString) JavaLangString {
	return NewJavaLangString(strings.ToUpper(s.toString()))
}

func Java_java_lang_Object_registerNatives() {

}
