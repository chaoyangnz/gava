package jago


func register_sun_reflect_Reflection() {
	register("sun/reflect/Reflection.getCallerClass()Ljava/lang/Class;", Java_sun_reflect_Reflection_getCallerClass)
	register("sun/reflect/Reflection.getClassAccessFlags(Ljava/lang/Class;)I", Java_sun_reflect_Reflection_getClassAccessFlags)
}

func Java_sun_reflect_Reflection_getCallerClass() JavaLangClass {
	//todo

	vmStack := VM_getCurrentThread().vmStack
	if len(vmStack) == 1 {
		return NULL
	} else {
		return vmStack[len(vmStack)-2].method.class.classObject
	}
}

func Java_sun_reflect_Reflection_getClassAccessFlags(classObj JavaLangClass) Int {
	return Int(u16toi32(classObj.GetExtra().(*Class).accessFlags))
}

