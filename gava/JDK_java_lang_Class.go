package gava

func register_java_lang_Class() {
	VM.RegisterNative("java/lang/Class.registerNatives()V", JDK_jang_lang_Class_registerNatives)
	VM.RegisterNative("java/lang/Class.getPrimitiveClass(Ljava/lang/String;)Ljava/lang/Class;", JDK_java_lang_Class_getPrimitiveClass)
	VM.RegisterNative("java/lang/Class.desiredAssertionStatus0(Ljava/lang/Class;)Z", JDK_java_lang_Class_desiredAssertionStatus0)
	VM.RegisterNative("java/lang/Class.getDeclaredFields0(Z)[Ljava/lang/reflect/Field;", JDK_java_lang_Class_getDeclaredFields0)
	VM.RegisterNative("java/lang/Class.isPrimitive()Z", JDK_java_lang_Class_isPrimitive)
	VM.RegisterNative("java/lang/Class.isAssignableFrom(Ljava/lang/Class;)Z", JDK_java_lang_Class_isAssignableFrom)
	VM.RegisterNative("java/lang/Class.getName0()Ljava/lang/String;", JDK_java_lang_Class_getName0)
	//VM.RegisterNative("java/lang/Class.forName0(Ljava/lang/String;ZLjava/lang/ClassLoader;Ljava/lang/Class;)Ljava/lang/Class;", JDK_java_lang_Class_forName0)
	VM.RegisterNative("java/lang/Class.isInterface()Z", JDK_java_lang_Class_isInterface)
	VM.RegisterNative("java/lang/Class.getDeclaredConstructors0(Z)[Ljava/lang/reflect/Constructor;", JDK_java_lang_Class_getDeclaredConstructors0)
	VM.RegisterNative("java/lang/Class.getModifiers()I", JDK_java_lang_Class_getModifiers)
	VM.RegisterNative("java/lang/Class.getSuperclass()Ljava/lang/Class;", JDK_java_lang_Class_getSuperclass)
	VM.RegisterNative("java/lang/Class.isArray()Z", JDK_java_lang_Class_isArray)
	VM.RegisterNative("java/lang/Class.getComponentType()Ljava/lang/Class;", JDK_java_lang_Class_getComponentType)
	VM.RegisterNative("java/lang/Class.getEnclosingMethod0()[Ljava/lang/Object;", JDK_java_lang_Class_getEnclosingMethod0)
	VM.RegisterNative("java/lang/Class.getDeclaringClass0()Ljava/lang/Class;", JDK_java_lang_Class_getDeclaringClass0)
	VM.RegisterNative("java/lang/Class.forName0(Ljava/lang/String;ZLjava/lang/ClassLoader;Ljava/lang/Class;)Ljava/lang/Class;", JDK_java_lang_Class_forName0)
}

// private static void registerNatives()
func JDK_jang_lang_Class_registerNatives() {}

// static Class getPrimitiveClass(String name)
func JDK_java_lang_Class_getPrimitiveClass(name JavaLangString) JavaLangClass {
	switch name.toNativeString() {
	case "byte":
		return BYTE_TYPE.ClassObject()
	case "short":
		return SHORT_TYPE.ClassObject()
	case "char":
		return CHAR_TYPE.ClassObject()
	case "int":
		return INT_TYPE.ClassObject()
	case "long":
		return LONG_TYPE.ClassObject()
	case "float":
		return FLOAT_TYPE.ClassObject()
	case "double":
		return DOUBLE_TYPE.ClassObject()
	case "boolean":
		return BOOLEAN_TYPE.ClassObject()
	default:
		VM.Throw("java/lang/RuntimeException", "Not a primitive type")
	}
	return NULL
}

// private static boolean desiredAssertionStatus0(Class javaClass)
func JDK_java_lang_Class_desiredAssertionStatus0(clazz JavaLangClass) Boolean {
	// Always disable assertions
	return FALSE
}

func JDK_java_lang_Class_getDeclaredFields0(this JavaLangClass, publicOnly Boolean) ArrayRef {
	class := this.retrieveType().(*Class)
	fields := class.GetDeclaredFields(publicOnly.IsTrue())
	fieldObjectArr := VM.NewArrayOfName("[Ljava/lang/reflect/Field;", Int(len(fields)))
	for i, field := range fields {
		fieldObjectArr.SetArrayElement(Int(i), VM.NewJavaLangReflectField(field))
	}

	return fieldObjectArr
}

func JDK_java_lang_Class_isPrimitive(this JavaLangClass) Boolean {
	type_ := this.retrieveType()
	if _, ok := type_.(*Class); ok {
		return FALSE
	}
	return TRUE
}

func JDK_java_lang_Class_isAssignableFrom(this JavaLangClass, cls JavaLangClass) Boolean {
	thisClass := this.retrieveType().(*Class)
	clsClass := cls.retrieveType().(*Class)

	assignable := FALSE
	if thisClass.IsAssignableFrom(clsClass) {
		assignable = TRUE
	}
	return assignable
}

func JDK_java_lang_Class_getName0(this JavaLangClass) JavaLangString {
	return binaryNameToJavaName(this.retrieveType().Name())
}

func JDK_java_lang_Class_forName0(name JavaLangString, initialize Boolean, loader JavaLangClassLoader, caller JavaLangClass) JavaLangClass {
	className := javaNameToBinaryName(name)
	return VM.ResolveClass(className, TRIGGER_BY_JAVA_REFLECTION).ClassObject()
}

func JDK_java_lang_Class_isInterface(this JavaLangClass) Boolean {
	if this.retrieveType().(*Class).IsInterface() {
		return TRUE
	}
	return FALSE
}

func JDK_java_lang_Class_getDeclaredConstructors0(this JavaLangClass, publicOnly Boolean) ArrayRef {
	class := this.retrieveType().(*Class)

	constructors := class.GetConstructors(publicOnly.IsTrue())

	constructorArr := VM.NewArrayOfName("[Ljava/lang/reflect/Constructor;", Int(len(constructors)))
	for i, constructor := range constructors {
		constructorArr.SetArrayElement(Int(i), VM.NewJavaLangReflectConstructor(constructor))
	}

	return constructorArr
}

func JDK_java_lang_Class_getModifiers(this JavaLangClass) Int {
	return Int(u16toi32(this.retrieveType().(*Class).accessFlags))
}

func JDK_java_lang_Class_getSuperclass(this JavaLangClass) JavaLangClass {
	class := this.retrieveType().(*Class)
	if class.name == "java/lang/Object" {
		return NULL
	}
	return class.superClass.ClassObject()
}

func JDK_java_lang_Class_isArray(this JavaLangClass) Boolean {
	type0 := this.retrieveType().(Type)
	switch type0.(type) {
	case *Class:
		if type0.(*Class).IsArray() {
			return TRUE
		}
	}
	return FALSE
}

func JDK_java_lang_Class_getComponentType(this JavaLangClass) JavaLangClass {
	class := this.retrieveType().(*Class)
	if !class.IsArray() {
		Fatal("%s is not array type", this.Class().name)
	}

	return class.componentType.ClassObject()
}

func JDK_java_lang_Class_getEnclosingMethod0(this JavaLangClass) ArrayRef {

	//TODO
	return NULL
}

func JDK_java_lang_Class_getDeclaringClass0(this JavaLangClass) JavaLangClass {

	//TODO
	return NULL
}
