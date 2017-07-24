package jago


func register_sun_misc_Unsafe() {
	register("sun/misc/Unsafe.registerNatives()V", Java_sun_misc_Unsafe_registerNatives)
	register("sun/misc/Unsafe.arrayBaseOffset(Ljava/lang/Class;)I", Java_sun_misc_Unsafe_arrayBaseOffset)
	register("sun/misc/Unsafe.arrayIndexScale(Ljava/lang/Class;)I", Java_sun_misc_Unsafe_arrayIndexScale)
	register("sun/misc/Unsafe.addressSize()I", Java_sun_misc_Unsafe_addressSize)
	register("sun/misc/Unsafe.objectFieldOffset(Ljava/lang/reflect/Field;)J", Java_sun_misc_Unsafe_objectFieldOffset)
	register("sun/misc/Unsafe.compareAndSwapObject(Ljava/lang/Object;JLjava/lang/Object;Ljava/lang/Object;)Z", Java_sun_misc_Unsafe_compareAndSwapObject)
	register("sun/misc/Unsafe.getIntVolatile(Ljava/lang/Object;J)I", Java_sun_misc_Unsafe_getIntVolatile)
	register("sun/misc/Unsafe.compareAndSwapInt(Ljava/lang/Object;JII)Z", Java_sun_misc_Unsafe_compareAndSwapInt)
	register("sun/misc/Unsafe.allocateMemory(J)J", Java_sun_misc_Unsafe_allocateMemory)
	register("sun/misc/Unsafe.putLong(JJ)V", Java_sun_misc_Unsafe_putLong)
	register("sun/misc/Unsafe.getByte(J)B", Java_sun_misc_Unsafe_getByte)
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

func Java_sun_misc_Unsafe_objectFieldOffset(this Reference, fieldObject JavaLangReflectField) Long {
	slot := fieldObject.GetInstanceVariableByName("slot", "I").(Int)
	return Long(slot)
}

func Java_sun_misc_Unsafe_compareAndSwapObject(this Reference, obj Reference, offset Long, expected Reference, newVal Reference) Boolean {
	if obj.IsNull() {
		Throw("NullPointerException", "")
	}

	slots := obj.oop.values
	current := slots[offset]
	if current == expected {
		slots[offset] = newVal
		return TRUE
	}

	return FALSE
}

func Java_sun_misc_Unsafe_compareAndSwapInt(this Reference, obj Reference, offset Long, expected Int, newVal Int) Boolean {
	if obj.IsNull() {
		Throw("NullPointerException", "")
	}

	slots := obj.oop.values
	current := slots[offset]
	if current == expected {
		slots[offset] = newVal
		return TRUE
	}

	return FALSE
}

func Java_sun_misc_Unsafe_getIntVolatile(this Reference, obj Reference, offset Long) Int {
	if obj.IsNull() {
		Throw("NullPointerException", "")
	}

	slots := obj.oop.values
	return slots[offset].(Int)
}

func Java_sun_misc_Unsafe_allocateMemory(this Reference, size Long) Long {
	//TODO
	return size
}

func Java_sun_misc_Unsafe_putLong(this Reference, address Long, val Long) {
	//TODO
}

func Java_sun_misc_Unsafe_getByte(this Reference, address Long) Byte {
	//TODO
	return Byte(0)
}