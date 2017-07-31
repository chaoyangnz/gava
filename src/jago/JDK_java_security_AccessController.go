package jago

func register_java_security_AccessController() {
	VM.RegisterNative("java/security/AccessController.doPrivileged(Ljava/security/PrivilegedExceptionAction;)Ljava/lang/Object;", JDK_java_security_AccessController_doPrivileged)
	VM.RegisterNative("java/security/AccessController.doPrivileged(Ljava/security/PrivilegedAction;)Ljava/lang/Object;", JDK_java_security_AccessController_doPrivileged)
	VM.RegisterNative("java/security/AccessController.getStackAccessControlContext()Ljava/security/AccessControlContext;", JDK_java_security_AccessController_getStackAccessControlContext)
	VM.RegisterNative("java/security/AccessController.doPrivileged(Ljava/security/PrivilegedExceptionAction;Ljava/security/AccessControlContext;)Ljava/lang/Object;", JDK_java_security_AccessController_doPrivilegedContext)
}

// because here need to call java method, so the return value will automatically be placed in the stack
func JDK_java_security_AccessController_doPrivileged(action Reference) Reference {
	method := action.Class().FindMethod("run", "()Ljava/lang/Object;")
	return VM.InvokeMethod(method, action).(Reference)
}

func JDK_java_security_AccessController_getStackAccessControlContext() Reference {
	//TODO
	return NULL
}

func JDK_java_security_AccessController_doPrivilegedContext(action Reference, context Reference) Reference  {
	return JDK_java_security_AccessController_doPrivileged(action)
}
