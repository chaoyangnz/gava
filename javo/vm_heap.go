package javo

import (
	"fmt"
	"hash/fnv"
	"strconv"
	"unsafe"
)

/*
|--------------------------------------------------------------------------
| Heap structure
|--------------------------------------------------------------------------
|
| Heap stores all the Java objects. Here, its major functionality includes
| factories of instantiation of many well-known classes: java.lang.Object,
| java.lang.Class, java.lang.String, java.lang.Thread, java.lang.Throwable,
| and Array, because these classes are necessary when bootstrapping the VM.
| As per the object location, we don't have the control of where the objects
| are located like realistic VM has a dedicated area to hold objects.
*/

type Heap struct {
}

func (this *Heap) NewObjectOfName(className string) ObjectRef {
	class := VM.ResolveClass(className, TRIGGER_BY_NEW_INSTANCE)

	return this.NewObject(class)
}

/*
arrayClassName is the full array class, not its component type
*/
func (this *Heap) NewArrayOfName(arrayClassName string, length Int) ArrayRef {
	arrayClass := VM.ResolveClass(arrayClassName, TRIGGER_BY_NEW_INSTANCE)

	return this.NewArray(arrayClass, length)
}

func hashcode(object *Object) Int {
	h := fnv.New32a()
	address := uintptr(unsafe.Pointer(object))

	h.Write([]byte(strconv.Itoa(int(address))))

	//ts := time.Now().UnixNano()
	//h.Write([]byte(strconv.Itoa(int(ts))))

	return Int(h.Sum32())
}

func (this *Heap) NewObject(class *Class) ObjectRef {
	if class.IsArray() {
		VM.Throw("java/lang/IllegalArgumentException", "Class %s must be an Non-Array class", class.name)
	}
	object := &Object{}
	object.header = Header{
		class:    class,
		hashCode: hashcode(object),
		monitor:  NewMonitor(object)}
	object.slots = make([]Value, class.instanceVarsCount)
	// Initialize instance variables
	clazz := class
	for clazz != nil {
		for _, field := range clazz.fields {
			if !field.IsStatic() {
				object.slots[field.slot] = field.DefaultValue()
			}
		}
		clazz = clazz.superClass
	}

	// verify initialization
	for _, instanceVar := range object.slots {
		if instanceVar == nil {
			Fatal("Something wrong, unfinished instance variable initialization")
		}
	}

	return Reference{object}
}

func (this *Heap) NewArrayOfComponent(arrayComponentType Type, length Int) ArrayRef {
	arrayClass := VM.ResolveClass(JVM_SIGNATURE_ARRAY+arrayComponentType.Descriptor(), TRIGGER_BY_NEW_INSTANCE)

	return this.NewArray(arrayClass, length)
}

func (this *Heap) NewArray(arrayClass *Class, length Int) ArrayRef {
	if !arrayClass.IsArray() {
		VM.Throw("java/lang/IllegalArgumentException", "Class %s must be an Array class", arrayClass.name)
	}
	elements := make([]Value, length)
	for i := range elements {
		switch arrayClass.componentType.(type) {
		case *ByteType:
			elements[i] = Byte(0)
		case *ShortType:
			elements[i] = Short(0)
		case *CharType:
			elements[i] = Char(0)
		case *IntType:
			elements[i] = Int(0)
		case *LongType:
			elements[i] = Long(0)
		case *FloatType:
			elements[i] = Float(0.0)
		case *DoubleType:
			elements[i] = Long(0.0)
		case *BooleanType:
			elements[i] = FALSE
		case *Class:
			elements[i] = NULL
		default:
			Fatal("Not a valid component type")
		}
	}

	object := &Object{}
	object.header = Header{
		class:    arrayClass,
		hashCode: hashcode(object),
		monitor:  NewMonitor(object)}
	object.slots = elements

	return Reference{object}
}

func (this *Heap) NewMultiDimensionalArray(arrayClass *Class, lengths []Int) ArrayRef {
	count := lengths[0]
	arr := this.NewArray(arrayClass, count)

	if len(lengths) > 1 {
		elements := arr.ArrayElements()
		for i := range elements {
			elements[i] = this.NewMultiDimensionalArray(arrayClass.componentType.(*Class), lengths[1:])
		}
	}

	return arr
}

/*
Create interned java.lang.String
*/
func (this *Heap) NewJavaLangString(str string) JavaLangString {
	// check string table
	if strObj, found := VM.GetStringInPool(str); found {
		return strObj
	}

	object := VM.NewObjectOfName(JAVA_LANG_STRING)

	// convert rune (int32) to Java char (UTF-16 with surrogate)
	var chars []Char
	for _, codepoint := range str {
		if codepoint <= 0xFFFF {
			chars = append(chars, Char(codepoint))
		} else {
			/*
				H = (S - 10000) / 400 + D800
				L = (S - 10000) % 400 + DC00
			*/
			high_surrogate := Char((uint32(codepoint)-0x10000)/0x400 + 0xD800)
			low_surrogate := Char((uint32(codepoint)-0x10000)%0x400 + 0xDC00)
			chars = append(chars, high_surrogate, low_surrogate)
		}
	}
	// create value field
	values := VM.NewArrayOfName("[C", Int(len(chars)))
	for i := range values.ArrayElements() {
		values.SetArrayElement(Int(i), chars[i])
	}
	object.SetInstanceVariableByName("value", "[C", values)

	// create hashing field?
	// TODO

	// intern
	return VM.InternString(object)
}

func (this *Heap) NewJavaLangClass(type0 Type) JavaLangClass {

	classObject := VM.NewObjectOfName(JAVA_LANG_CLASS)

	switch type0.(type) {
	case *BooleanType:
		type0.(*BooleanType).classObject = classObject
	case *ByteType:
		type0.(*ByteType).classObject = classObject
	case *ShortType:
		type0.(*ShortType).classObject = classObject
	case *IntType:
		type0.(*IntType).classObject = classObject
	case *LongType:
		type0.(*LongType).classObject = classObject
	case *FloatType:
		type0.(*FloatType).classObject = classObject
	case *DoubleType:
		type0.(*DoubleType).classObject = classObject
	case *Class:
		type0.(*Class).classObject = classObject
		VM.Info(">>> create java.lang.Class for %s *c=%p and jc=%p \n", type0.Name(), type0.(*Class), classObject.oop)
	}

	classObject.attachType(type0) // attach Class to it

	return classObject
}

func (this *Heap) NewJavaLangThread(name string) JavaLangThread {
	jThread := VM.NewObjectOfName(JAVA_LANG_THREAD)

	// ThreadGroup just need to set its name; no parent
	jThreadGroup := VM.NewObjectOfName("java/lang/ThreadGroup")
	jThreadGroup.SetInstanceVariableByName("name", "Ljava/lang/String;", VM.NewJavaLangString("main"))

	jThread.SetInstanceVariableByName("name", "Ljava/lang/String;", VM.NewJavaLangString(name))
	jThread.SetInstanceVariableByName("group", "Ljava/lang/ThreadGroup;", jThreadGroup)
	jThread.SetInstanceVariableByName("priority", "I", Int(5))
	// no initialization here

	return jThread
}

func (this *Heap) NewThrowable(exception string, message string, args ...interface{}) Reference {
	msg := fmt.Sprintf(message, args...)

	throwable := VM.NewObjectOfName(exception)
	constructorWithMessage := throwable.Class().GetConstructor("(Ljava/lang/String;)V")
	if constructorWithMessage != nil {
		VM.InvokeMethod(constructorWithMessage, throwable, VM.NewJavaLangString(msg))
	} else {
		constructorDefault := throwable.Class().GetConstructor("()V")
		if constructorDefault == nil {
			Fatal("%s has no default constructor")
		}
		VM.InvokeMethod(constructorWithMessage, throwable)
	}

	return throwable
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
func (this *Heap) NewJavaLangReflectField(field *Field) JavaLangReflectField {
	var fieldObject = VM.NewObjectOfName(JAVA_LANG_REFLECT_FIELD)
	fieldObject.SetInstanceVariableByName("clazz", JVM_SIGNATURE_CLASS+JAVA_LANG_CLASS+JVM_SIGNATURE_ENDCLASS, field.class.ClassObject())
	fieldObject.SetInstanceVariableByName("name", JVM_SIGNATURE_CLASS+JAVA_LANG_STRING+JVM_SIGNATURE_ENDCLASS, VM.NewJavaLangString(field.name))

	fieldObject.SetInstanceVariableByName("type", JVM_SIGNATURE_CLASS+JAVA_LANG_CLASS+JVM_SIGNATURE_ENDCLASS, VM.GetTypeClass(field.descriptor))
	fieldObject.SetInstanceVariableByName("modifiers", JVM_SIGNATURE_INT, Int(field.accessFlags))
	fieldObject.SetInstanceVariableByName("slot", JVM_SIGNATURE_INT, Int(field.slot))
	fieldObject.SetInstanceVariableByName("signature", JVM_SIGNATURE_CLASS+JAVA_LANG_STRING+JVM_SIGNATURE_ENDCLASS, VM.NewJavaLangString(field.descriptor))

	annotations := VM.NewArrayOfName("[B", 0) //TODO
	fieldObject.SetInstanceVariableByName("annotations", "[B", annotations)

	//fieldObject.SetExtra(field)
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
func (this *Heap) NewJavaLangReflectConstructor(method *Method) JavaLangReflectConstructor {
	constructorObject := VM.NewObjectOfName(JAVA_LANG_REFLECT_CONSTRUCTOR)
	constructorObject.SetInstanceVariableByName("clazz", JVM_SIGNATURE_CLASS+JAVA_LANG_CLASS+JVM_SIGNATURE_ENDCLASS, method.class.ClassObject())

	parameterTypes := VM.NewArrayOfName("[Ljava/lang/Class;", Int(len(method.parameterDescriptors)))
	for i, parameterDescriptor := range method.parameterDescriptors {
		parameterTypes.SetArrayElement(Int(i), VM.GetTypeClass(parameterDescriptor))
	}
	constructorObject.SetInstanceVariableByName("parameterTypes", "[Ljava/lang/Class;", parameterTypes)
	// TODO
	exceptionTypes := VM.NewArrayOfName("[Ljava/lang/Class;", 0)
	constructorObject.SetInstanceVariableByName("exceptionTypes", "[Ljava/lang/Class;", exceptionTypes)
	constructorObject.SetInstanceVariableByName("modifiers", JVM_SIGNATURE_INT, Int(method.accessFlags))
	constructorObject.SetInstanceVariableByName("slot", JVM_SIGNATURE_INT, Int(0)) // TODO method slot
	constructorObject.SetInstanceVariableByName("signature", "Ljava/lang/String;", VM.NewJavaLangString(method.descriptor))

	annotations := VM.NewArrayOfName("[B", 0) //TODO
	constructorObject.SetInstanceVariableByName("annotations", "[B", annotations)

	parameterAnnotations := VM.NewArrayOfName("[B", 0) //TODO
	constructorObject.SetInstanceVariableByName("parameterAnnotations", "[B", parameterAnnotations)

	//constructorObject.SetExtra(method)
	return constructorObject
}
