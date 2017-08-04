package jago

func register_sun_misc_URLClassPath()  {
	VM.RegisterNative("sun/misc/URLClassPath.getLookupCacheURLs(Ljava/lang/ClassLoader;)[Ljava/net/URL;", JDK_sun_misc_URLClassPath_getLookupCacheURLs)
}

func JDK_sun_misc_URLClassPath_getLookupCacheURLs(classloader JavaLangClassLoader) ArrayRef {
	return VM.NewArrayOfName("[Ljava/net/URL;", 0)
}