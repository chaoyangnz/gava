package jago

import "fmt"

var NULL = Reference{nil}

type Reference struct {
	oop *Object
}

type ObjectRef = Reference
type ArrayRef = Reference

func (this Reference) Type() Type {
	return this.Class()
}
func (this Reference) IsNull() bool  { return this.oop == nil }
func (this Reference) IsArray() bool { return this.oop != nil && this.oop.header.class.IsArray() }
func (this Reference) IsEqual(reference Reference) bool {
	if this.IsNull() && reference.IsNull() {
		return true
	}
	if !this.IsNull() && !reference.IsNull() {
		return this.oop == reference.oop
	}
	return false
}

func (this Reference) Class() *Class {
	return this.oop.header.class
}
func (this Reference) Monitor() *Monitor {
	return this.oop.header.monitor
}

func (this Reference) assertObject() {
	if this.IsNull() {
		VM.Throw("java/lang/NullPointerException", "")
	}
	if this.Class().IsArray() {
		Bug("It is not an ObjectRef")
	}
}

func (this Reference) assertArray() {
	if this.IsNull() {
		VM.Throw("java/lang/NullPointerException", "")
	}
	if !this.Class().IsArray() {
		Bug("It is not an ArrayRef")
	}
}

func (this Reference) IHashCode() Int {
	if this.IsNull() {
		return Int(0)
	}
	return this.oop.header.hashCode
}

func (this Reference) Clone() Reference {

	var clone Reference
	if !this.Class().IsArray() {
		clone = VM.NewObject(this.Class())
	} else {
		clone = VM.NewArray(this.Class(), this.ArrayLength())
	}

	for i, value := range this.oop.slots {
		clone.oop.slots[i] = value
	}

	return clone
}

func (this ObjectRef) GetInstanceVariable(index Int) Value {
	this.assertObject()
	return this.oop.slots[index]
}
func (this ObjectRef) SetInstanceVariable(index Int, value Value) {
	this.assertObject()
	this.oop.slots[index] = value
}
func (this ObjectRef) GetInstanceVariableByName(name string, descriptor string) Value {
	this.assertObject()

	objectref := this.oop
	field := objectref.header.class.FindField(name, descriptor)
	if field.IsStatic() {
		Fatal("not a instance variable")
	}
	return objectref.slots[field.slot]
}
func (this ObjectRef) SetInstanceVariableByName(name string, descriptor string, value Value) {
	this.assertObject()

	objectref := this.oop
	field := objectref.header.class.FindField(name, descriptor)
	if field == nil {
		Fatal("Field (%s %s)  cannot be found in %s", name, descriptor, objectref.header.class.Name())
	}
	if field.IsStatic() {
		Fatal("not a instance variable")
	}
	objectref.slots[field.slot] = value
}

func (this Reference) dump() {
	if !this.IsNull() {
		VM.Debug("Dump object (%s): {", this.Class().Name())
		for _, field := range this.Class().fields {
			if !field.IsStatic() {
				value := this.GetInstanceVariable(Int(field.slot))
				VM.Debug("\t%s: %v", field.name, value)
			}
		}
		VM.Debug("}\n")
	}
}

func (this ArrayRef) ArrayElements() []Value {
	this.assertArray()
	return this.oop.slots
}
func (this ArrayRef) ArrayLength() Int {
	this.assertArray()
	return Int(len(this.oop.slots))
}
func (this ArrayRef) GetArrayElement(index Int) Value {
	this.assertArray()
	return this.oop.slots[index]
}
func (this ArrayRef) SetArrayElement(index Int, value Value) {
	this.assertArray()
	if index >= this.ArrayLength() {
		VM.Throw("java/lang/ArrayIndexOutOfBoundsException", "%d exceeded array boundary %d", index, this.ArrayLength())
	}
	this.oop.slots[index] = value
}

/**
 * Well-known reference types
 */
const (
	JAVA_LANG_STRING              = "java/lang/String"
	JAVA_LANG_CLASS               = "java/lang/Class"
	JAVA_LANG_REFLECT_FIELD       = "java/lang/reflect/Field"
	JAVA_LANG_REFLECT_CONSTRUCTOR = "java/lang/reflect/Constructor"
	JAVA_LANG_REFLECT_METHOD      = "java/lang/reflect/Method"
	JAVA_LANG_THREAD              = "java/lang/Thread"
)

type JavaLangString = Reference

func (this JavaLangString) toNativeString() string {
	if this.IsNull() {
		VM.Throw("java/lang/NullPointerException", "")
	}
	this.assertObject()
	if this.Class().name != JAVA_LANG_STRING {
		Bug("It is not a java/lang/String")
	}

	var runes []rune
	chars := this.GetInstanceVariableByName("value", "[C").(ArrayRef).ArrayElements()
	for i := 0; i < len(chars); i++ {
		char := chars[i].(Char)
		if char >= 0xD800 && char <= 0xDBFF {
			h := char
			if i+1 < len(chars) && chars[i+1].(Char) >= 0xDC00 && chars[i+1].(Char) <= 0xDFFF {
				l := chars[i+1].(Char)
				//1000016 + (H − D80016) × 40016 + (L − DC0016)
				codepoint := 0x1000 + (h-0xD800)*0x400 + (l - 0xDC00)
				runes = append(runes, rune(codepoint))
			} else {
				VM.Throw("java/lang/UnknownError", "Illegal UTF-16 string: only high surrogate")
			}
			i++
		} else if char >= 0xDC00 && char <= 0xDFFF {
			VM.Throw("java/lang/UnknownError", "Illegal UTF-16 string: only low surrogate")
		} else {
			runes = append(runes, rune(char))
		}
	}
	return string(runes)
}

type JavaLangClass = Reference

func (this JavaLangClass) attachType(type0 Type) {
	var C *Class
	if c, ok := type0.(*Class); ok {
		C = c
	}
	VM.Info(":::%s *Class *c=%p attach to classobject jc=%p\n", type0.Name(), C, this.oop)
	this.oop.header._type = type0
}

func (this JavaLangClass) retrieveType() Type {

	return this.oop.header._type
}

type JavaLangThread = Reference

func (this JavaLangThread) attachThread(thread *Thread) {
	this.oop.header._thread = thread
}

func (this JavaLangThread) retrieveThread() *Thread {
	return this.oop.header._thread
}

type JavaLangThrowable = JavaLangThread

func (this JavaLangThrowable) attachStacktrace(stacktrace []StackTraceElement) {
	this.oop.header._backTrace = stacktrace
}

func (this JavaLangThrowable) retrieveStacktrace() []StackTraceElement {
	return this.oop.header._backTrace
}

type StackTraceElement struct {
	declaringClass string
	methodName     string
	fileName       string
	lineNumber     int
}

func (this StackTraceElement) toString() string {
	return fmt.Sprintf("%s.%s(%s:%d)", this.declaringClass, this.methodName, this.fileName, this.lineNumber)
}

type JavaLangReflectField = Reference
type JavaLangReflectConstructor = Reference
type JavaLangReflectMethod = Reference
type JavaLangClassLoader = Reference