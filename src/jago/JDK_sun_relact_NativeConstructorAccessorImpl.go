package jago

func register_sun_reflect_NativeConstructorAccessorImpl() {
	register("sun/reflect/NativeConstructorAccessorImpl.newInstance0(Ljava/lang/reflect/Constructor;[Ljava/lang/Object;)Ljava/lang/Object;", Java_sun_reflect_NativeConstructorAccessorImpl_newInstance0)
}

func Java_sun_reflect_NativeConstructorAccessorImpl_newInstance0(constructor JavaLangReflectConstructor, args ArrayRef) /*Reference*/ {

	class := constructor.GetInstanceVariableByName("clazz", "Ljava/lang/Class;")
	method := constructor.GetExtra().(*Method)
	VM_invokeJavaMethod(THREAD_MANAGER.currentThread, method, args.(Reference).oop.values...)
}
