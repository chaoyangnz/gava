package jago

func register_java_lang_reflect_Array()  {
	register("java/lang/reflect/Array.newArray(Ljava/lang/Class;I)Ljava/lang/Object;", JDK_java_lang_reflect_Array_newArray)
}

func JDK_java_lang_reflect_Array_newArray(componentClassObject JavaLangClass, length Int) Reference {
	componentClass := componentClassObject.GetExtra().(*Class)
	class := BOOTSTRAP_CLASSLOADER.CreateClass(JVM_SIGNATURE_ARRAY + componentClass.Descriptor(), TRIGGER_BY_NEW_INSTANCE)
	if !class.IsArray() {
		VM_throw("java/lang/IllegalArgumentException", "Class %s is not an Array class", class.name)
	}

	return class.NewArray(length).(Reference)
}