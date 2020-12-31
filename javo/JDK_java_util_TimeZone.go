package javo

import "time"

func register_java_util_TimeZone() {
	VM.RegisterNative("java/util/TimeZone.getSystemTimeZoneID(Ljava/lang/String;)Ljava/lang/String;", JDK_java_util_TimeZone_getSystemTimeZoneID)
}

func JDK_java_util_TimeZone_getSystemTimeZoneID(javaHome JavaLangString) JavaLangString {
	loc := time.Local
	return VM.NewJavaLangString(loc.String())
}
