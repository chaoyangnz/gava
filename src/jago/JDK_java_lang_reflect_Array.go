package jago

func register_java_lang_reflect_Array()  {
	VM.RegisterNative("java/lang/reflect/Array.newArray(Ljava/lang/Class;I)Ljava/lang/Object;", JDK_java_lang_reflect_Array_newArray)
}

func JDK_java_lang_reflect_Array_newArray(componentClassObject JavaLangClass, length Int) ArrayRef {
	componentType := componentClassObject.retrieveType()

	return VM.NewArrayOfComponent(componentType, length)
}