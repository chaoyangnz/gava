package jago

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
		VM.Throw("java/lang/NullPointerException", "")
	}
	this.assertObject()
	if this.AsObjectRef().Class().name != JAVA_LANG_STRING {
		Bug("It is not a java/lang/String")
	}

	runes := []rune{}
	chars := this.GetInstanceVariableByName("value", "[C").(ArrayRef).ArrayElements()
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
				VM.Throw("java/lang/UnknownError","Illegal UTF-16 string: only high surrogate")
			}
			i++
		} else if char >= 0xDC00 && char <= 0xDFFF {
			VM.Throw("java/lang/UnknownError","Illegal UTF-16 string: only low surrogate")
		} else {
			runes = append(runes, rune(char))
		}
	}
	return string(runes)
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



