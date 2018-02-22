package jago

func register_sun_reflect_NativeConstructorAccessorImpl() {
	VM.RegisterNative("sun/reflect/NativeConstructorAccessorImpl.newInstance0(Ljava/lang/reflect/Constructor;[Ljava/lang/Object;)Ljava/lang/Object;", JDK_sun_reflect_NativeConstructorAccessorImpl_newInstance0)
}

func JDK_sun_reflect_NativeConstructorAccessorImpl_newInstance0(constructor JavaLangReflectConstructor, args ArrayRef) ObjectRef {

	classObject := constructor.GetInstanceVariableByName("clazz", "Ljava/lang/Class;").(JavaLangClass)
	class := classObject.retrieveType().(*Class)
	descriptor := constructor.GetInstanceVariableByName("signature", "Ljava/lang/String;").(JavaLangString).toNativeString()

	method := class.GetConstructor(descriptor)

	objeref := VM.NewObject(class)
	allArgs := []Value{objeref}
	if !args.IsNull() {
		allArgs = append(allArgs, args.oop.slots...)
	}

	VM.InvokeMethod(method, allArgs...)

	return objeref
}
