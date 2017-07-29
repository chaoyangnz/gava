package jago

func register_java_lang_Object() {
	register("java/lang/Object.registerNatives()V", JDK_jang_lang_Object_registerNatives)
	register("java/lang/Object.hashCode()I", JDK_java_lang_Object_hashCode)
	register("java/lang/Object.getClass()Ljava/lang/Class;", JDK_java_lang_Object_getClass)
	register("java/lang/Object.clone()Ljava/lang/Object;", JDK_java_lang_Object_clone)
	register("java/lang/Object.notifyAll()V", JDK_java_lang_Object_notifyAll)
	register("java/lang/Object.wait(J)V", JDK_java_lang_Object_wait)
}

// private static void registerNatives()
func JDK_jang_lang_Object_registerNatives()  {}

func JDK_java_lang_Object_hashCode(this Reference) Int  {
	return VM_iHashCode(this)
}

func JDK_java_lang_Object_getClass(this Reference) JavaLangClass {
	return this.Class().ClassObject()
}

func JDK_java_lang_Object_clone(this Reference) Reference {
	cloneable := BOOTSTRAP_CLASSLOADER.CreateClass("java/lang/Cloneable", TRIGGER_BY_CHECK_OBJECT_TYPE)
	if !cloneable.IsAssignableFrom(this.Class()) {
		VM_throw("java/lang/CloneNotSupportedException", "Not implement java.lang.Cloneable")
	}

	return VM_clone(this)
}

func JDK_java_lang_Object_wait(this Reference, timeout Long) {
	// TODO timeout
	monitor := this.Monitor()
	if !monitor.HasOwner(VM_currentThread()) {
		VM_throw("java/lang/IllegalMonitorStateException", "Cannot wait() when not holding a monitor")
	}

	monitor.Wait()
}

func JDK_java_lang_Object_notifyAll(this Reference) {
	monitor := this.Monitor()
	if !monitor.HasOwner(VM_currentThread()) {
		VM_throw("java/lang/IllegalMonitorStateException", "Cannot notifyAll() when not holding a monitor")
	}

	monitor.NotifyAll()
}

