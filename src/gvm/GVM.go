package gvm

/* ---------- JDK Native methods implementation ---*/

func Java_GVM_print(o *j_object, s java_lang_string) {
	print(s.toString())
}