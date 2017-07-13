package jago


func register_sun_misc_Unsafe() {
	register("sun/misc/Unsafe.registerNatives()V", Java_sun_misc_Unsafe_registerNatives)
	register("sun/misc/Unsafe.arrayBaseOffset(Ljava/lang/Class;)I", Java_sun_misc_Unsafe_arrayBaseOffset)
	register("sun/misc/Unsafe.arrayIndexScale(Ljava/lang/Class;)I", Java_sun_misc_Unsafe_arrayIndexScale)
	register("sun/misc/Unsafe.addressSize()I", Java_sun_misc_Unsafe_addressSize)
}

// private static void registerNatives()
func Java_sun_misc_Unsafe_registerNatives() {}

func Java_sun_misc_Unsafe_arrayBaseOffset(this Reference, arrayClass JavaLangClass) Int {
	//todo
	return Int(0)
}

func Java_sun_misc_Unsafe_arrayIndexScale(this Reference, arrayClass JavaLangClass) Int {
	//todo
	return Int(0)
}

func Java_sun_misc_Unsafe_addressSize(this Reference) Int {
	//todo
	return Int(0)
}