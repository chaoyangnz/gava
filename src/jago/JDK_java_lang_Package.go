package jago

import "strings"

func register_java_lang_Package()  {
	VM.RegisterNative("java/lang/Package.getSystemPackage0(Ljava/lang/String;)Ljava/lang/String;", JDK_java_lang_Package_getSystemPackage0)
}

func JDK_java_lang_Package_getSystemPackage0(name JavaLangString) JavaLangString {
	for _, cal := range VM.ClassCache.Items() {
		classAndLoader := cal.(*ClassAndLoader)
		if classAndLoader.classLoader.IsNull() && strings.HasPrefix(classAndLoader.class.Name(), name.toNativeString()) {
			return name
		}
	}
	return NULL
}