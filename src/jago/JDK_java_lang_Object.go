package jago

func register_java_lang_Object() {
	VM.RegisterNative("java/lang/Object.registerNatives()V", JDK_jang_lang_Object_registerNatives)
	VM.RegisterNative("java/lang/Object.hashCode()I", JDK_java_lang_Object_hashCode)
	VM.RegisterNative("java/lang/Object.getClass()Ljava/lang/Class;", JDK_java_lang_Object_getClass)
	VM.RegisterNative("java/lang/Object.clone()Ljava/lang/Object;", JDK_java_lang_Object_clone)
	VM.RegisterNative("java/lang/Object.notifyAll()V", JDK_java_lang_Object_notifyAll)
	VM.RegisterNative("java/lang/Object.wait(J)V", JDK_java_lang_Object_wait)
	VM.RegisterNative("java/lang/Object.notify()V", JDK_java_lang_Object_notifyAll)
}

// private static void registerNatives()
func JDK_jang_lang_Object_registerNatives() {}

func JDK_java_lang_Object_hashCode(this Reference) Int {
	return VM.IHashCode(this)
}

func JDK_java_lang_Object_getClass(this Reference) JavaLangClass {
	return this.Class().ClassObject()
}

func JDK_java_lang_Object_clone(this Reference) Reference {
	cloneable := VM.ResolveClass("java/lang/Cloneable", TRIGGER_BY_CHECK_OBJECT_TYPE)
	if !cloneable.IsAssignableFrom(this.Class()) {
		VM.Throw("java/lang/CloneNotSupportedException", "Not implement java.lang.Cloneable")
	}

	return VM.Clone(this)
}

func JDK_java_lang_Object_wait(this Reference, millis Long) {
	// TODO timeout
	monitor := this.Monitor()
	if !monitor.HasOwner(VM.CurrentThread()) {
		VM.Throw("java/lang/IllegalMonitorStateException", "Cannot wait() when not holding a monitor")
	}

	interrupted := monitor.Wait(int64(millis))
	if interrupted {
		VM.Throw("java/lang/InterruptedException", "wait interrupted")
	}
}

func JDK_java_lang_Object_notifyAll(this Reference) {
	monitor := this.Monitor()
	if !monitor.HasOwner(VM.CurrentThread()) {
		VM.Throw("java/lang/IllegalMonitorStateException", "Cannot notifyAll() when not holding a monitor")
	}

	monitor.NotifyAll()
}
