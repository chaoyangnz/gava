package jago


func register_sun_reflect_Reflection() {
	register("sun/reflect/Reflection.getCallerClass()Ljava/lang/Class;", Java_sun_reflect_Reflection_getCallerClass)
}

func Java_sun_reflect_Reflection_getCallerClass() JavaLangClass {
	//todo

	vmStack := THREAD_MANAGER.currentThread.vmStack
	if len(vmStack) == 1 {
		return NULL
	} else {
		return vmStack[len(vmStack)-2].method.class.classObject
	}
}

