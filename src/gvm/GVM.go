package gvm

import "strings"

/* ---------- JDK Native methods implementation ---*/

func Java_GVM_println(o *j_object, s java_lang_string) {
	println(s.toString())
}

func Java_GVM_toUpper(s java_lang_string) java_lang_string {
	return newJavaLangString(strings.ToUpper(s.toString()))
}