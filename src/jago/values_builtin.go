package jago

import "fmt"

var NULL = Reference{nil}

const (
	JAVA_LANG_STRING = "java/lang/String"
	JAVA_LANG_CLASS = "java/lang/Class"
	JAVA_LANG_REFLECT_FIELD = "java/lang/reflect/Field"
	JAVA_LANG_REFLECT_CONSTRUCTOR = "java/lang/reflect/Constructor"
	JAVA_LANG_REFLECT_METHOD = "java/lang/reflect/Method"
	JAVA_LANG_THREAD = "java/lang/Thread"
)


type JavaLangString interface {
	ObjectRef
	toNativeString() string
}

func (this Reference) toNativeString() string  {
	if this.IsNull() {
		VM_throw("java/lang/NullPointerException", "")
	}
	this.assertObject()
	if this.AsObjectRef().Class().name != JAVA_LANG_STRING {
		Bug("It is not a java/lang/String")
	}

	runes := []rune{}
	chars := this.GetInstanceVariableByName("value", "[C").(ArrayRef).Elements()
	for i:=0; i < len(chars); i++ {
		char := chars[i].(Char)
		if char >= 0xD800 && char <= 0xDBFF {
			h := char
			if i+1 < len(chars) && chars[i+1].(Char) >= 0xDC00 && chars[i+1].(Char) <= 0xDFFF {
				l := chars[i+1].(Char)
				//1000016 + (H − D80016) × 40016 + (L − DC0016)
				codepoint := 0x1000 + (h - 0xD800) * 0x400 + (l - 0xDC00)
				runes = append(runes, rune(codepoint))
			} else {
				VM_throw("java/lang/UnknownError","Illegal UTF-16 string: only high surrogate")
			}
			i++
		} else if char >= 0xDC00 && char <= 0xDFFF {
			VM_throw("java/lang/UnknownError","Illegal UTF-16 string: only low surrogate")
		} else {
			runes = append(runes, rune(char))
		}
	}
	return string(runes)
}

/*
Create interned java.lang.String
 */
func NewJavaLangString(str string) JavaLangString {
	// check string table
	if strObj, found := STRING_TABLE[str]; found {
		return strObj
	}
	object := NewObject(JAVA_LANG_STRING)

	// convert rune (int32) to Java char (UTF-16 with surrogate)
	chars := []Char{}
	for _, codepoint := range str {
		if codepoint <= 0xFFFF {
			chars = append(chars, Char(codepoint))
		} else {
			/*
				H = (S - 10000) / 400 + D800
				L = (S - 10000) % 400 + DC00
			 */
			high_surrogate := Char((uint32(codepoint) - 0x10000) / 0x400 + 0xD800)
			low_surrogate := Char((uint32(codepoint) - 0x10000) % 0x400 + 0xDC00)
			chars = append(chars, high_surrogate, low_surrogate)
		}
	}
	// create value field
	values := NewArray("[C", Int(len(chars)))
	for i, _ := range values.Elements() {
		values.SetElement(Int(i), chars[i])
	}
	object.SetInstanceVariableByName("value", "[C", values)

	// create hashing field?
	// TODO

	// intern
	return VM_intern_String(object.(Reference))
}

type JavaLangClass interface {ObjectRef}
type JavaLangReflectField interface {ObjectRef}
type JavaLangReflectConstructor interface{ObjectRef}
type JavaLangReflectMethod interface {ObjectRef}
type JavaLangClassLoader interface{ObjectRef}

// if need to attach extra *Class, do it yourself!!
func NewJavaLangClass(_type Type) JavaLangClass {
	object := NewObject(JAVA_LANG_CLASS)

	object.SetExtra(_type) // attach Class to it

	return object
}

func __getTypeClass(descriptor string) JavaLangClass {
	var typeClass JavaLangClass
	switch string(descriptor[0]) {
	case JVM_SIGNATURE_CLASS: {
		fieldTypeClassName := descriptor[1:len(descriptor)-1]
		typeClass = BOOTSTRAP_CLASSLOADER.CreateClass(fieldTypeClassName).ClassObject()
	}
	case JVM_SIGNATURE_ARRAY: {
		fieldTypeClassName := descriptor
		typeClass = BOOTSTRAP_CLASSLOADER.CreateClass(fieldTypeClassName).ClassObject()
	}
	case JVM_SIGNATURE_BYTE: typeClass = BYTE_TYPE.ClassObject()
	case JVM_SIGNATURE_SHORT: typeClass = SHORT_TYPE.ClassObject()
	case JVM_SIGNATURE_CHAR: typeClass = CHAR_TYPE.ClassObject()
	case JVM_SIGNATURE_INT: typeClass = INT_TYPE.ClassObject()
	case JVM_SIGNATURE_LONG: typeClass = LONG_TYPE.ClassObject()
	case JVM_SIGNATURE_FLOAT: typeClass = FLOAT_TYPE.ClassObject()
	case JVM_SIGNATURE_DOUBLE: typeClass = DOUBLE_TYPE.ClassObject()
	case JVM_SIGNATURE_BOOLEAN: typeClass = BOOLEAN_TYPE.ClassObject()
	default:
		Fatal("type %s is not a unsupported type", descriptor)
	}

	return typeClass
}

/*
Field(Class<?> declaringClass,
      String name,
      Class<?> type,
      int modifiers,
      int slot,
      String signature,
      byte[] annotations)

      	"(Ljava/lang/Class;" +
	"Ljava/lang/String;" +
	"Ljava/lang/Class;" +
	"II" +
	"Ljava/lang/String;" +
	"[B)V"
*/
func NewJavaLangReflectField(field *Field) JavaLangReflectField {
	fieldObject := NewObject(JAVA_LANG_REFLECT_FIELD)
	fieldObject.SetInstanceVariableByName("clazz", JVM_SIGNATURE_CLASS + JAVA_LANG_CLASS + JVM_SIGNATURE_ENDCLASS, field.class.classObject)
	fieldObject.SetInstanceVariableByName("name", JVM_SIGNATURE_CLASS + JAVA_LANG_STRING + JVM_SIGNATURE_ENDCLASS, NewJavaLangString(field.name))

	fieldObject.SetInstanceVariableByName("type", JVM_SIGNATURE_CLASS + JAVA_LANG_CLASS + JVM_SIGNATURE_ENDCLASS, __getTypeClass(field.descriptor))
	fieldObject.SetInstanceVariableByName("modifiers", JVM_SIGNATURE_INT, Int(field.accessFlags))
	fieldObject.SetInstanceVariableByName("slot", JVM_SIGNATURE_INT, Int(field.index))
	fieldObject.SetInstanceVariableByName("signature", JVM_SIGNATURE_CLASS + JAVA_LANG_STRING + JVM_SIGNATURE_ENDCLASS, NewJavaLangString(field.descriptor))

	annotations := NewArray("[B", 0) //TODO
	fieldObject.SetInstanceVariableByName("annotations", "[B", annotations)

	fieldObject.SetExtra(field)
	return fieldObject
}

/*
Constructor(Class<T> declaringClass,
            Class<?>[] parameterTypes,
            Class<?>[] checkedExceptions,
            int modifiers,
            int slot,
            String signature,
            byte[] annotations,
            byte[] parameterAnnotations)
}
*/
func NewJavaLangReflectConstructor(method *Method) JavaLangReflectConstructor {
	constructorObject := NewObject(JAVA_LANG_REFLECT_CONSTRUCTOR)
	constructorObject.SetInstanceVariableByName("clazz", JVM_SIGNATURE_CLASS + JAVA_LANG_CLASS + JVM_SIGNATURE_ENDCLASS, method.class.classObject)

	parameterTypes := NewArray("[Ljava/lang/Class;", Int(len(method.parameterDescriptors)))
	for i, parameterDescriptor := range method.parameterDescriptors {
		parameterTypes.SetElement(Int(i), __getTypeClass(parameterDescriptor))
	}
	constructorObject.SetInstanceVariableByName("parameterTypes", "[Ljava/lang/Class;", parameterTypes)
	// TODO
	exceptionTypes := NewArray("[Ljava/lang/Class;", 0)
	constructorObject.SetInstanceVariableByName("exceptionTypes", "[Ljava/lang/Class;", exceptionTypes)
	constructorObject.SetInstanceVariableByName("modifiers", JVM_SIGNATURE_INT, Int(method.accessFlags))
	constructorObject.SetInstanceVariableByName("slot", JVM_SIGNATURE_INT, Int(0)) // TODO method slot
	constructorObject.SetInstanceVariableByName("signature", "Ljava/lang/String;", NewJavaLangString(method.descriptor))

	annotations := NewArray("[B", 0) //TODO
	constructorObject.SetInstanceVariableByName("annotations", "[B", annotations)

	parameterAnnotations := NewArray("[B", 0) //TODO
	constructorObject.SetInstanceVariableByName("parameterAnnotations", "[B", parameterAnnotations)

	constructorObject.SetExtra(method)
	return constructorObject
}

type JavaLangThread interface {ObjectRef}

func NewJavaLangThread() JavaLangThread {
	jThread := NewObject(JAVA_LANG_THREAD)

	jGroup := NewObject("java/lang/ThreadGroup")

	jThread.SetInstanceVariableByName("group", "Ljava/lang/ThreadGroup;", jGroup)
	jThread.SetInstanceVariableByName("priority", "I", Int(1))

	return jThread
}

func NewThrowable(exception string, message string, args ...interface{}) Reference {
	msg := fmt.Sprintf(message, args...)

	throwable := NewObject(exception).(Reference)
	constructorWithMessage := throwable.Class().GetConstructor("(Ljava/lang/String;)V")
	if constructorWithMessage != nil {
		VM_invokeMethod(constructorWithMessage, throwable, NewJavaLangString(msg))
	} else {
		constructorDefault := throwable.Class().GetConstructor( "()V")
		if constructorDefault == nil {
			Fatal("%s has no default constructor")
		}
		VM_invokeMethod(constructorWithMessage, throwable)
	}

	return throwable
}

/////////////////////// Type Conversion //////////////////////////////////
func (this Boolean) ToInt() Int {
	return Int(this)
}

func (this Boolean) ToByte() Byte {
	return Byte(this)
}

func (this Int) ToBoolean() Boolean {
	return Boolean(this)
}

func (this Byte) ToBoolean() Boolean {
	return Boolean(this)
}
