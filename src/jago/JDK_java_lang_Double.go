package jago

func register_java_lang_Float() {
	register("java/lang/Float.floatToRawIntBits(F)I", Java_java_lang_Float_floatToRawIntBits)
}

// public static native int floatToRawIntBits(float value)
func Java_java_lang_Float_floatToRawIntBits(value Float) Int {
	return Int(int32(float32(value)))
}