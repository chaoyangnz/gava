package jago

func register_sun_misc_Unsafe() {
	VM.RegisterNative("sun/misc/Unsafe.registerNatives()V", JDK_sun_misc_Unsafe_registerNatives)
	VM.RegisterNative("sun/misc/Unsafe.arrayBaseOffset(Ljava/lang/Class;)I", JDK_sun_misc_Unsafe_arrayBaseOffset)
	VM.RegisterNative("sun/misc/Unsafe.arrayIndexScale(Ljava/lang/Class;)I", JDK_sun_misc_Unsafe_arrayIndexScale)
	VM.RegisterNative("sun/misc/Unsafe.addressSize()I", JDK_sun_misc_Unsafe_addressSize)
	VM.RegisterNative("sun/misc/Unsafe.objectFieldOffset(Ljava/lang/reflect/Field;)J", JDK_sun_misc_Unsafe_objectFieldOffset)
	VM.RegisterNative("sun/misc/Unsafe.compareAndSwapObject(Ljava/lang/Object;JLjava/lang/Object;Ljava/lang/Object;)Z", JDK_sun_misc_Unsafe_compareAndSwapObject)
	VM.RegisterNative("sun/misc/Unsafe.getIntVolatile(Ljava/lang/Object;J)I", JDK_sun_misc_Unsafe_getIntVolatile)
	VM.RegisterNative("sun/misc/Unsafe.getObjectVolatile(Ljava/lang/Object;J)Ljava/lang/Object;", JDK_sun_misc_Unsafe_getObjectVolatile)
	VM.RegisterNative("sun/misc/Unsafe.putObjectVolatile(Ljava/lang/Object;JLjava/lang/Object;)V", JDK_sun_misc_Unsafe_putObjectVolatile)

	VM.RegisterNative("sun/misc/Unsafe.compareAndSwapInt(Ljava/lang/Object;JII)Z", JDK_sun_misc_Unsafe_compareAndSwapInt)
	VM.RegisterNative("sun/misc/Unsafe.compareAndSwapLong(Ljava/lang/Object;JJJ)Z", JDK_sun_misc_Unsafe_compareAndSwapLong)
	VM.RegisterNative("sun/misc/Unsafe.allocateMemory(J)J", JDK_sun_misc_Unsafe_allocateMemory)
	VM.RegisterNative("sun/misc/Unsafe.putLong(JJ)V", JDK_sun_misc_Unsafe_putLong)
	VM.RegisterNative("sun/misc/Unsafe.getByte(J)B", JDK_sun_misc_Unsafe_getByte)
	VM.RegisterNative("sun/misc/Unsafe.freeMemory(J)V", JDK_sun_misc_Unsafe_freeMemory)

	VM.RegisterNative("sun/misc/Unsafe.ensureClassInitialized(Ljava/lang/Class;)V", JDK_sun_misc_Unsafe_ensureClassInitialized)

}

// private static void registerNatives()
func JDK_sun_misc_Unsafe_registerNatives() {}

func JDK_sun_misc_Unsafe_arrayBaseOffset(this Reference, arrayClass JavaLangClass) Int {
	//todo
	return Int(0)
}

func JDK_sun_misc_Unsafe_arrayIndexScale(this Reference, arrayClass JavaLangClass) Int {
	//todo
	return Int(1)
}

func JDK_sun_misc_Unsafe_addressSize(this Reference) Int {
	//todo
	return Int(8)
}

func JDK_sun_misc_Unsafe_objectFieldOffset(this Reference, fieldObject JavaLangReflectField) Long {
	slot := fieldObject.GetInstanceVariableByName("slot", "I").(Int)
	return Long(slot)
}

func JDK_sun_misc_Unsafe_compareAndSwapObject(this Reference, obj Reference, offset Long, expected Reference, newVal Reference) Boolean {
	if obj.IsNull() {
		VM.Throw("java/lang/NullPointerException", "")
	}

	slots := obj.oop.slots
	current := slots[offset]
	if current == expected {
		slots[offset] = newVal
		return TRUE
	}

	return FALSE
}

func JDK_sun_misc_Unsafe_compareAndSwapInt(this Reference, obj Reference, offset Long, expected Int, newVal Int) Boolean {
	if obj.IsNull() {
		VM.Throw("java/lang/NullPointerException", "")
	}

	slots := obj.oop.slots
	current := slots[offset]
	if current == expected {
		slots[offset] = newVal
		return TRUE
	}

	return FALSE
}

func JDK_sun_misc_Unsafe_compareAndSwapLong(this Reference, obj Reference, offset Long, expected Long, newVal Long) Boolean {
	if obj.IsNull() {
		VM.Throw("java/lang/NullPointerException", "")
	}

	slots := obj.oop.slots
	current := slots[offset]
	if current == expected {
		slots[offset] = newVal
		return TRUE
	}

	return FALSE
}

func JDK_sun_misc_Unsafe_getIntVolatile(this Reference, obj Reference, offset Long) Int {
	if obj.IsNull() {
		VM.Throw("java/lang/NullPointerException", "")
	}

	slots := obj.oop.slots
	return slots[offset].(Int)
}

func JDK_sun_misc_Unsafe_getObjectVolatile(this Reference, obj Reference, offset Long) Reference {
	slots := obj.oop.slots
	return slots[offset].(Reference)
}

func JDK_sun_misc_Unsafe_putObjectVolatile(this Reference, obj Reference, offset Long, val Reference) {
	slots := obj.oop.slots
	slots[offset] = val
}

func JDK_sun_misc_Unsafe_allocateMemory(this Reference, size Long) Long {
	//TODO
	return size
}

func JDK_sun_misc_Unsafe_putLong(this Reference, address Long, val Long) {
	//TODO
}

func JDK_sun_misc_Unsafe_getByte(this Reference, address Long) Byte {
	//TODO
	return Byte(0x08) //0x01 big_endian
}

func JDK_sun_misc_Unsafe_freeMemory(this Reference, size Long) {
	// do nothing
}

func JDK_sun_misc_Unsafe_ensureClassInitialized(this Reference, class JavaLangClass) {
	// LOCK ???
	if class.retrieveType().(*Class).initialized != INITIALIZED {
		VM.Throw("java/lang/AssertionError", "Class has not been initialized")
	}
}
