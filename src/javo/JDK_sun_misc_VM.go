package javo


func register_sun_misc_VM() {
	VM.RegisterNative("sun/misc/VM.initialize()V", JDK_sun_misc_VM_initialize)
}

// private static void registerNatives()
func JDK_sun_misc_VM_initialize()  {}