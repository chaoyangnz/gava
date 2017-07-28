package jago


func register_sun_misc_Unsafe() {
	register("sun/misc/Unsafe.registerNatives()V", JDK_sun_misc_Unsafe_registerNatives)
	register("sun/misc/Unsafe.arrayBaseOffset(Ljava/lang/Class;)I", JDK_sun_misc_Unsafe_arrayBaseOffset)
	register("sun/misc/Unsafe.arrayIndexScale(Ljava/lang/Class;)I", JDK_sun_misc_Unsafe_arrayIndexScale)
	register("sun/misc/Unsafe.addressSize()I", JDK_sun_misc_Unsafe_addressSize)
	register("sun/misc/Unsafe.objectFieldOffset(Ljava/lang/reflect/Field;)J", JDK_sun_misc_Unsafe_objectFieldOffset)
	register("sun/misc/Unsafe.compareAndSwapObject(Ljava/lang/Object;JLjava/lang/Object;Ljava/lang/Object;)Z", JDK_sun_misc_Unsafe_compareAndSwapObject)
	register("sun/misc/Unsafe.getIntVolatile(Ljava/lang/Object;J)I", JDK_sun_misc_Unsafe_getIntVolatile)
	register("sun/misc/Unsafe.getObjectVolatile(Ljava/lang/Object;J)Ljava/lang/Object;", JDK_sun_misc_Unsafe_getObjectVolatile)
	register("sun/misc/Unsafe.putObjectVolatile(Ljava/lang/Object;JLjava/lang/Object;)V", JDK_sun_misc_Unsafe_putObjectVolatile)

	register("sun/misc/Unsafe.compareAndSwapInt(Ljava/lang/Object;JII)Z", JDK_sun_misc_Unsafe_compareAndSwapInt)
	register("sun/misc/Unsafe.compareAndSwapLong(Ljava/lang/Object;JJJ)Z", JDK_sun_misc_Unsafe_compareAndSwapLong)
	register("sun/misc/Unsafe.allocateMemory(J)J", JDK_sun_misc_Unsafe_allocateMemory)
	register("sun/misc/Unsafe.putLong(JJ)V", JDK_sun_misc_Unsafe_putLong)
	register("sun/misc/Unsafe.getByte(J)B", JDK_sun_misc_Unsafe_getByte)
	register("sun/misc/Unsafe.freeMemory(J)V", JDK_sun_misc_Unsafe_freeMemory)

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
		VM_throw("java/lang/NullPointerException", "")
	}

	slots := obj.oop.values
	current := slots[offset]
	if current == expected {
		slots[offset] = newVal
		return TRUE
	}

	return FALSE
}

func JDK_sun_misc_Unsafe_compareAndSwapInt(this Reference, obj Reference, offset Long, expected Int, newVal Int) Boolean {
	if obj.IsNull() {
		VM_throw("java/lang/NullPointerException", "")
	}

	slots := obj.oop.values
	current := slots[offset]
	if current == expected {
		slots[offset] = newVal
		return TRUE
	}

	return FALSE
}

func JDK_sun_misc_Unsafe_compareAndSwapLong(this Reference, obj Reference, offset Long, expected Long, newVal Long) Boolean {
	if obj.IsNull() {
		VM_throw("java/lang/NullPointerException", "")
	}

	slots := obj.oop.values
	current := slots[offset]
	if current == expected {
		slots[offset] = newVal
		return TRUE
	}

	return FALSE
}

func JDK_sun_misc_Unsafe_getIntVolatile(this Reference, obj Reference, offset Long) Int {
	if obj.IsNull() {
		VM_throw("java/lang/NullPointerException", "")
	}

	slots := obj.oop.values
	return slots[offset].(Int)
}

func JDK_sun_misc_Unsafe_getObjectVolatile(this Reference, obj Reference, offset Long) Reference {
	slots := obj.oop.values
	return slots[offset].(Reference)
}

func JDK_sun_misc_Unsafe_putObjectVolatile(this Reference, obj Reference, offset Long, val Reference)  {
	slots := obj.oop.values
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

func JDK_sun_misc_Unsafe_freeMemory(this Reference, size Long)  {
	// do nothing
}

