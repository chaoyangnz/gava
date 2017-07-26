package jago

import "strings"

func register_java_lang_Class() {
	register("java/lang/Class.registerNatives()V", Java_jang_lang_Class_registerNatives)
	register("java/lang/Class.getPrimitiveClass(Ljava/lang/String;)Ljava/lang/Class;", Java_java_lang_Class_getPrimitiveClass)
	register("java/lang/Class.desiredAssertionStatus0(Ljava/lang/Class;)Z", Java_java_lang_Class_desiredAssertionStatus0)
	register("java/lang/Class.getDeclaredFields0(Z)[Ljava/lang/reflect/Field;", Java_java_lang_Class_getDeclaredFields0)
	register("java/lang/Class.isPrimitive()Z", Java_java_lang_Class_isPrimitive)
	register("java/lang/Class.isAssignableFrom(Ljava/lang/Class;)Z", Java_java_lang_Class_isAssignableFrom)
	register("java/lang/Class.getName0()Ljava/lang/String;", Java_java_lang_Class_getName0)
	register("java/lang/Class.forName0(Ljava/lang/String;ZLjava/lang/ClassLoader;Ljava/lang/Class;)Ljava/lang/Class;", Java_java_lang_Class_forName0)
	register("java/lang/Class.isInterface()Z", Java_java_lang_Class_isInterface)
	register("java/lang/Class.getDeclaredConstructors0(Z)[Ljava/lang/reflect/Constructor;", Java_java_lang_Class_getDeclaredConstructors0)
	register("java/lang/Class.getModifiers()I", Java_java_lang_Class_getModifiers)
	register("java/lang/Class.getSuperclass()Ljava/lang/Class;", Java_java_lang_Class_getSuperclass)
	register("java/lang/Class.isArray()Z", Java_java_lang_Class_isArray)
	register("java/lang/Class.getComponentType()Ljava/lang/Class;", Java_java_lang_Class_getComponentType)
}

// private static void registerNatives()
func Java_jang_lang_Class_registerNatives()  {}

// static Class getPrimitiveClass(String name)
func Java_java_lang_Class_getPrimitiveClass(name JavaLangString) JavaLangClass {
	switch name.toNativeString() {
	case "byte": return BYTE_TYPE.ClassObject()
	case "short": return SHORT_TYPE.ClassObject()
	case "char": return CHAR_TYPE.ClassObject()
	case "int": return INT_TYPE.ClassObject()
	case "long": return LONG_TYPE.ClassObject()
	case "float": return FLOAT_TYPE.ClassObject()
	case "double": return DOUBLE_TYPE.ClassObject()
	case "boolean": return BOOLEAN_TYPE.ClassObject()
	default:
		PseudoThrow("java/lang/RuntimeException", "Not a primitive type")
	}
	return nil
}

// private static boolean desiredAssertionStatus0(Class javaClass)
func Java_java_lang_Class_desiredAssertionStatus0(clazz JavaLangClass) Boolean {
	// Always disable assertions
	return FALSE
}


func Java_java_lang_Class_getDeclaredFields0(this JavaLangClass, publicOnly Boolean) ArrayRef {
	class := this.GetExtra().(*Class)
	fields := class.GetDeclaredFields(publicOnly.IsTrue())
	fieldObjectArr := NewArray("[Ljava/lang/reflect/Field;", Int(len(fields)))
	for i, field := range fields {
		fieldObjectArr.SetElement(Int(i), NewJavaLangReflectField(field))
	}

	return fieldObjectArr
}

func Java_java_lang_Class_isPrimitive(this JavaLangClass) Boolean {
	type_ := this.GetExtra().(Type)
	if _, ok := type_.(*Class); ok {
		return FALSE
	}
	return TRUE
}

func Java_java_lang_Class_isAssignableFrom(this JavaLangClass, cls JavaLangClass) Boolean {
	thisClass := this.GetExtra().(*Class)
	clsClass := cls.GetExtra().(*Class)

	assignable := FALSE
	if thisClass.IsAssignableFrom(clsClass) {
		assignable = TRUE
	}
	return assignable
}

func Java_java_lang_Class_getName0(this JavaLangClass) JavaLangString {
	classNameJavaStyle := strings.Replace(this.Class().name, "/", ".", -1)

	return NewJavaLangString(classNameJavaStyle)
}

func Java_java_lang_Class_forName0(name JavaLangString, initialize Boolean, loader JavaLangClassLoader, caller JavaLangClass) JavaLangClass {
	className := strings.Replace(name.toNativeString(), ".", "/", -1)
	return BOOTSTRAP_CLASSLOADER.CreateClass(className).ClassObject()
}

func Java_java_lang_Class_isInterface(this JavaLangClass) Boolean {
	if this.GetExtra().(*Class).IsInterface() {
		return TRUE
	}
	return FALSE
}

func Java_java_lang_Class_getDeclaredConstructors0(this JavaLangClass, publicOnly Boolean) ArrayRef {
	class := this.GetExtra().(*Class)


	constructors := []*Method{}
	for _,method := range class.methods {
		if method.name == "<init>" && method.returnDescriptor == JVM_SIGNATURE_VOID {
			if publicOnly.IsTrue() {
				if (method.accessFlags & JVM_ACC_PUBLIC) > 0 {
					constructors = append(constructors, method)
				}
			} else {
				constructors = append(constructors, method)
			}
		}
	}

	constructorArr := NewArray("[Ljava/lang/reflect/Constructor;", Int(len(constructors)))
	for i,constructor := range constructors {
		constructorArr.SetElement(Int(i), NewJavaLangReflectConstructor(constructor))
	}

	return constructorArr
}

func Java_java_lang_Class_getModifiers(this JavaLangClass) Int {
	return Int(u16toi32((this.GetExtra().(*Class).accessFlags)))
}

func Java_java_lang_Class_getSuperclass(this JavaLangClass) JavaLangClass {
	class := this.GetExtra().(*Class)
	if class.name == "java/lang/Object" {
		return NULL
	}
	return class.superClass.ClassObject()
}

func Java_java_lang_Class_isArray(this JavaLangClass) Boolean {
	type0 := this.GetExtra().(Type)
	switch type0.(type) {
	case *Class:
		if type0.(*Class).IsArray() {
			return TRUE
		}
	}
	return FALSE
}

func Java_java_lang_Class_getComponentType(this JavaLangClass) JavaLangClass {
	class := this.GetExtra().(*Class)
	if !class.IsArray() {
		Fatal("%s is not array type", this.Class().name)
	}

	return class.componentType.ClassObject()
}
