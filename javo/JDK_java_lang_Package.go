package javo

import "strings"

func register_java_lang_Package() {
	VM.RegisterNative("java/lang/Package.getSystemPackage0(Ljava/lang/String;)Ljava/lang/String;", JDK_java_lang_Package_getSystemPackage0)
}

func JDK_java_lang_Package_getSystemPackage0(vmPackageName JavaLangString) JavaLangString {
	for nl, class := range VM.MethodArea.DefinedClasses {

		if nl.L == nil && strings.HasPrefix(class.Name(), vmPackageName.toNativeString()) {
			return vmPackageName
		}
	}
	return NULL
}
