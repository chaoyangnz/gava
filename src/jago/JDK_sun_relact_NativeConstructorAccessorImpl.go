package jago

func register_sun_reflect_NativeConstructorAccessorImpl() {
	register("sun/reflect/NativeConstructorAccessorImpl.newInstance0(Ljava/lang/reflect/Constructor;[Ljava/lang/Object;)Ljava/lang/Object;", Java_sun_reflect_NativeConstructorAccessorImpl_newInstance0)
}

func Java_sun_reflect_NativeConstructorAccessorImpl_newInstance0(constructor JavaLangReflectConstructor, args ArrayRef) ObjectRef {

	class := constructor.GetInstanceVariableByName("clazz", "Ljava/lang/Class;").(JavaLangClass).GetExtra().(*Class)
	descriptor := constructor.GetInstanceVariableByName("signature", "Ljava/lang/String;").(JavaLangString).toNativeString()

	method := class.GetConstructor(descriptor)

	objeref := class.NewObject()
	allArgs := []Value {objeref}
	if !args.IsNull() {
		allArgs = append(allArgs, args.(Reference).oop.values...)
	}

	VM_invokeMethod(method, allArgs...)

	return objeref
}
