package gvm

type GNIEnv struct {}

/* ---------- JDK Native methods implementation ---*/
func GVM_print(s java_lang_string) {
	println(s.toString())
}




