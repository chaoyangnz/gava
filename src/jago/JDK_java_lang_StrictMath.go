package jago

import "math"

func register_java_lang_StrictMath() {
	VM.RegisterNative("java/lang/StrictMath.pow(DD)D", JDK_java_lang_StrictMath_pow)
}

// private static void registers()
func JDK_java_lang_StrictMath_pow(base Double, exponent Double) Double {
	return Double(math.Pow(float64(base), float64(exponent)))
}
