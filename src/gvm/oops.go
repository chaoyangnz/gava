package gvm


/*
Java object VM representation type hierarchy:

t_any (interface)
   |- t_byte
   |- t_char
   |- t_short
   |- t_int
   |- t_long
   |- t_float
   |- t_double
   |- t_boolean
   |- t_reference (interface)
        |- *t_array
        |- *t_object
            |- *java_lang_object
            |- *java_lang_string
            |- *java_lang_class
            |- *java_lang_classloader
            |- *java_lang_thread
            |- ....

All the pointer type starting with java_** is the equivalent to jdk referenceType for convenient operations
 */

type (
	// these types are for vm internal use only
	// basically they are mapped to 10 types mentioned in JVM specification
	t_any interface {
		isReference()   bool
		defaultValue()  t_any
	}

	t_byte          int8
	t_char          uint16
	t_short         int16
	t_int           int32
	t_long          int64
	t_boolean       int32
	t_float         float32
	t_double        float64
	t_reference Reference
	t_object        struct {
					//header part
					class *ClassType
					flags uint32
					locks uint32
					//fields
					fields []t_any
	}
	t_array         struct {
					//header part
					atype  Type
					flags  uint32
					locks  uint32
					length t_int
					//fields
					elements []t_any
	}
)


const (
	boolean_true  = t_boolean(0x1)
	boolean_false = t_boolean(0x0)

	byte_default = t_byte(0)
	short_default = t_short(0)
	char_default = t_char(0)
	int_default = t_int(0)
	long_default = t_long(0)
	float_default = t_float(0.0)
	double_default = t_double(0.0)
	boolean_default = boolean_false
)
var reference_default Reference = nil
var object_default *t_object = (*t_object)(nil)
var array_default *t_array = (*t_array)(nil)

func (this t_byte) isReference() bool  {
	return false
}

func (this t_byte) defaultValue() t_any  {
	return byte_default
}

func (this t_short) isReference() bool  {
	return false
}

func (this t_short) defaultValue() t_any  {
	return short_default
}

func (this t_char) isReference() bool  {
	return false
}

func (this t_char) defaultValue() t_any  {
	return char_default
}

func (this t_int) isReference() bool  {
	return false
}

func (this t_int) defaultValue() t_any  {
	return int_default
}

func (this t_long) isReference() bool  {
	return false
}

func (this t_long) defaultValue() t_any  {
	return long_default
}

func (this t_float) isReference() bool  {
	return false
}

func (this t_float) defaultValue() t_any  {
	return float_default
}

func (this t_double) isReference() bool  {
	return false
}

func (this t_double) defaultValue() t_any  {
	return double_default
}

func (this t_boolean) isReference() bool  {
	return false
}

func (this t_boolean) defaultValue() t_any  {
	return boolean_default
}

type Reference interface {
	isReference() bool
	defaultValue() t_any
	isArray() bool
}

func (this *t_object)isReference() bool  {
	return true
}

func (this *t_object) defaultValue() t_any {
	return (*t_object)(nil)
}

func (this *t_object)isArray() bool  {
	return false
}

func (this *t_array)isReference() bool  {
	return true
}

func (this *t_array) defaultValue() t_any {
	return (*t_array)(nil)
}

func (this *t_array)isArray() bool  {
	return true
}

func newArray(componentType ReferenceType, length t_int) *t_array {
	elements := make([]t_any, length)
	for i, _ := range elements {
		switch componentType.(type) {
		case *ClassType: elements[i] = object_default
		case *ArrayType: elements[i] = array_default
		default:
			fatal("Not a reference type")
		}
	}
	return &t_array{atype: componentType, length: length, elements: elements}
}

func newByteArray(length t_int) *t_array {
	elements := make([]t_any, length)
	for i, _:= range elements {
		elements[i] = byte_default
	}
	return &t_array{atype: BYTE_TYPE, length: length, elements: elements}
}

func newShortArray(length t_int) *t_array {
	elements := make([]t_any, length)
	for i, _:= range elements {
		elements[i] = short_default
	}
	return &t_array{atype: SHORT_TYPE, length: length, elements: elements}
}

func newCharArray(length t_int) *t_array {
	elements := make([]t_any, length)
	for i, _:= range elements {
		elements[i] = char_default
	}
	return &t_array{atype: CHAR_TYPE, length: length, elements: elements}
}

func newIntArray(length t_int) *t_array {
	elements := make([]t_any, length)
	for i, _:= range elements {
		elements[i] = int_default
	}
	return &t_array{atype: INT_TYPE, length: length, elements: elements}
}

func newLongArray(length t_int) *t_array {
	elements := make([]t_any, length)
	for i, _:= range elements {
		elements[i] = long_default
	}
	return &t_array{atype: LONG_TYPE, length: length, elements: elements}
}

func newFloatArray(length t_int) *t_array {
	elements := make([]t_any, length)
	for i, _:= range elements {
		elements[i] = float_default
	}
	return &t_array{atype: FLOAT_TYPE, length: length, elements: elements}
}

func newDoubleArray(length t_int) *t_array {
	elements := make([]t_any, length)
	for i, _:= range elements {
		elements[i] = double_default
	}
	return &t_array{atype: DOUBLE_TYPE, length: length, elements: elements}
}

func newBooleanArray(length t_int) *t_array {
	elements := make([]t_any, length)
	for i, _:= range elements {
		elements[i] = boolean_default
	}
	return &t_array{atype: BOOLEAN_TYPE, length: length, elements: elements}
}

func newObjectArray(componentClass *ClassType, length t_int) *t_array {
	elements := make([]t_any, length)
	for i, _:= range elements {
		elements[i] = object_default
	}
	return &t_array{atype: componentClass, length: length, elements: elements}
}

func newArrayArray(arrayComponentType *ArrayType, length t_int) *t_array {
	elements := make([]t_any, length)
	for i, _:= range elements {
		elements[i] = object_default
	}
	return &t_array{atype: arrayComponentType, length: length, elements: elements}
}


/*--------------------Extend JDK Class here--------------------------*/
// there objects are actually raw t_object, here provide some utility methods
type (
	java_lang_object struct {*t_object}
	java_lang_string struct {*t_object}
	java_lang_class struct {*t_object}
	java_lang_classloader struct {*t_object}
	java_lang_thread struct {*t_object}
)

func newJavaLangString(str string) *java_lang_string {
	java_lang_string_Class := bootstrapClassLoader.load("java/lang/String")
	object := java_lang_string_Class.newObject()
	// convert a utf8 string to utf-16 using as Java String
	chars := []t_char{}
	for _, codepoint := range str {
		if codepoint <= 0xFFFF {
			chars = append(chars, t_char(codepoint))
		} else {
			/*
				H = (S - 10000) / 400 + D800
				L = (S - 10000) % 400 + DC00
			 */
			high_surrogate := t_char((uint32(codepoint) - 0x10000) / 0x400 + 0xD800)
			low_surrogate := t_char((uint32(codepoint) - 0x10000) % 0x400 + 0xDC00)
			chars = append(chars, high_surrogate, low_surrogate)
		}
	}
	charArray := newCharArray(t_int(len(chars)))
	for i, _ := range chars {
		charArray.elements[i] = chars[i]
	}
	object.fields[0] = charArray
	return &java_lang_string{object}
}

func (this *java_lang_string) toString() string  {
	runes := []rune{}
	chars := this.fields[0].(*t_array).elements
	for i:=0; i < len(chars); i++ {
		char := chars[i].(t_char)
		if char >= 0xD800 && char <= 0xDBFF {
			h := char
			if i+1 < len(chars) && chars[i+1].(t_char) >= 0xDC00 && chars[i+1].(t_char) <= 0xDFFF {
				l := chars[i+1].(t_char)
				//1000016 + (H − D80016) × 40016 + (L − DC0016)
				codepoint := 0x1000 + (h - 0xD800) * 0x400 + (l - 0xDC00)
				runes = append(runes, rune(codepoint))
			} else {
				panic("Illegal UTF-16 string: only high surrogate")
			}
			i++
		} else if char >= 0xDC00 && char <= 0xDFFF {
			panic("Illegal UTF-16 string: only low surrogate")
		} else {
			runes = append(runes, rune(char))
		}
	}
	return string(runes)
}

func newJavaLangClass() *java_lang_class {
	java_lang_class_Class := bootstrapClassLoader.load("java/lang/Class")
	return &java_lang_class{java_lang_class_Class.newObject()}
}