package jago

func register_java_security_AccessController() {
	register("java/security/AccessController.doPrivileged(Ljava/security/PrivilegedExceptionAction;)Ljava/lang/Object;", Java_java_security_AccessController_doPrivileged)
	register("java/security/AccessController.doPrivileged(Ljava/security/PrivilegedAction;)Ljava/lang/Object;", Java_java_security_AccessController_doPrivileged)
	register("java/security/AccessController.getStackAccessControlContext()Ljava/security/AccessControlContext;", Java_java_security_AccessController_getStackAccessControlContext)
}

// because here need to call java method, so the return value will automatically be placed in the stack
func Java_java_security_AccessController_doPrivileged(action Reference) /*Reference*/ {
	method := action.Class().FindMethod("run", "()Ljava/lang/Object;")
	VM_invokeMethod(method, action)
}

func Java_java_security_AccessController_getStackAccessControlContext() Reference {
	//TODO
	return NULL
}
