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

type JavaLangClass interface {
	ObjectRef
	attachType(type0 Type)
	retrieveType() Type
}

func (this Reference) attachType(type0 Type)  {
	this.SetExtra(type0)
}
func (this Reference) retrieveType() Type {
	return this.GetExtra().(Type)
}

type JavaLangReflectField interface {ObjectRef}
type JavaLangReflectConstructor interface{ObjectRef}
type JavaLangReflectMethod interface {ObjectRef}
type JavaLangClassLoader interface{ObjectRef}

// if need to attach extra *Class, do it yourself!!
func NewJavaLangClass(_type Type) JavaLangClass {
	object := NewObject(JAVA_LANG_CLASS).(JavaLangClass)

	object.attachType(_type) // attach Class to it

	return object
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

	fieldObject.SetInstanceVariableByName("type", JVM_SIGNATURE_CLASS + JAVA_LANG_CLASS + JVM_SIGNATURE_ENDCLASS, VM_getTypeClass(field.descriptor))
	fieldObject.SetInstanceVariableByName("modifiers", JVM_SIGNATURE_INT, Int(field.accessFlags))
	fieldObject.SetInstanceVariableByName("slot", JVM_SIGNATURE_INT, Int(field.slot))
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
		parameterTypes.SetElement(Int(i), VM_getTypeClass(parameterDescriptor))
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

type JavaLangThread interface {
	ObjectRef
	attatchThread(thread *Thread)
	retrieveThread() *Thread
}

func (this Reference) attatchThread(thread *Thread)  {
	this.SetExtra(thread)
}

func (this Reference) retrieveThread() *Thread  {
	return this.GetExtra().(*Thread)
}

func NewJavaLangThread(name string) JavaLangThread {
	jThread := NewObject(JAVA_LANG_THREAD)

	// ThreadGroup just need to set its name; no parent
	jThreadGroup := NewObject("java/lang/ThreadGroup")
	jThreadGroup.SetInstanceVariableByName("name", "Ljava/lang/String;", NewJavaLangString("main"))

	jThread.SetInstanceVariableByName("name", "Ljava/lang/String;", NewJavaLangString(name))
	jThread.SetInstanceVariableByName("group", "Ljava/lang/ThreadGroup;", jThreadGroup)
	jThread.SetInstanceVariableByName("priority", "I", Int(5))
	// no initialization here

	return jThread.(JavaLangThread)
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

