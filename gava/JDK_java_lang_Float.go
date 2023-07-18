package gava

import "math"

func register_java_lang_Float() {
	VM.RegisterNative("java/lang/Float.floatToRawIntBits(F)I", JDK_java_lang_Float_floatToRawIntBits)
	VM.RegisterNative("java/lang/Float.intBitsToFloat(I)F", JDK_java_lang_Float_intBitsToFloat)
}

// public static native int floatToRawIntBits(float value)
func JDK_java_lang_Float_floatToRawIntBits(value Float) Int {
	bits := math.Float32bits(float32(value))
	return Int(int32(bits))
}

func JDK_java_lang_Float_intBitsToFloat(bits Int) Float {
	value := math.Float32frombits(uint32(bits))
	return Float(value)
}
