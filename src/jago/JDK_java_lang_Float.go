package jago

import "math"

func register_java_lang_Float() {
	register("java/lang/Float.floatToRawIntBits(F)I", Java_java_lang_Float_floatToRawIntBits)
	register("java/lang/Float.intBitsToFloat(I)F", Java_java_lang_Float_intBitsToFloat)
}

// public static native int floatToRawIntBits(float value)
func Java_java_lang_Float_floatToRawIntBits(value Float) Int {
	bits := math.Float32bits(float32(value))
	return Int(int32(bits))
}

func Java_java_lang_Float_intBitsToFloat(bits Int) Float {
	value := math.Float32frombits(uint32(bits))
	return Float(value)
}