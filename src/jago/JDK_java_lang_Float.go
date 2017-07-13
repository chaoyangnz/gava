package jago

func register_java_lang_Double() {
	register("java/lang/Double.doubleToRawLongBits(D)J", Java_jang_lang_Double_doubleToRawLongBits)
	register("java/lang/Double.longBitsToDouble(J)D", Java_jang_lang_Double_longBitsToDouble)
}

// public static native int floatToRawIntBits(float value)
func Java_jang_lang_Double_doubleToRawLongBits(value Double) Long {
	return Long(int64(float64(value)))
}

// public static native int floatToRawIntBits(float value)
func Java_jang_lang_Double_longBitsToDouble(value Long) Double {
	return Double(float64(int64(value)))
}