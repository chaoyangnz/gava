package gvm


/*
Java object VM representation type hierarchy:

t_any (interface)
   |- j_byte
   |- j_char
   |- j_short
   |- j_int
   |- j_long
   |- j_float
   |- j_double
   |- j_boolean
   |- j_reference (interface)
        |- *j_array
        |- *j_object
            < the following is java references>
            |- java_lang_object
            |- java_lang_string
            |- java_lang_class
            |- java_lang_classloader
            |- java_lang_thread
            |- ....

All the pointer type starting with java_** is the equivalent to jdk referenceType for convenient operations
 */

type (
	// these types are for vm internal use only
	// basically they are mapped to 10 types mentioned in JVM specification
	t_any           interface {
					isReference()   bool
					defaultValue()  t_any
	}

	j_byte          int8
	j_char          uint16
	j_short         int16
	j_int           int32
	j_long          int64
	j_boolean       int32
	j_float         float32
	j_double        float64

	j_reference     interface {
					isReference() bool
					defaultValue() t_any
					isArray() bool
	}
	j_object        struct {
					//header part
					class *ClassType
					flags uint32
					locks uint32
					//fields
					fields []t_any
	}
	j_array         struct {
					//header part
					atype  Type
					flags  uint32
					locks  uint32
					length j_int
					//fields
					elements []t_any
	}
)


const (
	boolean_true  = j_boolean(0x1)
	boolean_false = j_boolean(0x0)

	byte_default = j_byte(0)
	short_default = j_short(0)
	char_default = j_char(0)
	int_default = j_int(0)
	long_default = j_long(0)
	float_default = j_float(0.0)
	double_default = j_double(0.0)
	boolean_default = boolean_false
)
var (
	object_null = (*j_object)(nil)
	array_null = (*j_array)(nil)

	reference_default j_reference = nil
	object_default    *j_object   = object_null
	array_default     *j_array    = array_null
)

func (this j_byte) isReference() bool  {
	return false
}

func (this j_byte) defaultValue() t_any  {
	return byte_default
}

func (this j_short) isReference() bool  {
	return false
}

func (this j_short) defaultValue() t_any  {
	return short_default
}

func (this j_char) isReference() bool  {
	return false
}

func (this j_char) defaultValue() t_any  {
	return char_default
}

func (this j_int) isReference() bool  {
	return false
}

func (this j_int) defaultValue() t_any  {
	return int_default
}

func (this j_long) isReference() bool  {
	return false
}

func (this j_long) defaultValue() t_any  {
	return long_default
}

func (this j_float) isReference() bool  {
	return false
}

func (this j_float) defaultValue() t_any  {
	return float_default
}

func (this j_double) isReference() bool  {
	return false
}

func (this j_double) defaultValue() t_any  {
	return double_default
}

func (this j_boolean) isReference() bool  {
	return false
}

func (this j_boolean) defaultValue() t_any  {
	return boolean_default
}



func (this *j_object)isReference() bool  {
	return true
}

func (this *j_object) defaultValue() t_any {
	return (*j_object)(nil)
}

func (this *j_object)isArray() bool  {
	return false
}

func (this *j_array)isReference() bool  {
	return true
}

func (this *j_array) defaultValue() t_any {
	return (*j_array)(nil)
}

func (this *j_array)isArray() bool  {
	return true
}

func newArray(componentType ReferenceType, length j_int) *j_array {
	elements := make([]t_any, length)
	for i, _ := range elements {
		switch componentType.(type) {
		case *ClassType: elements[i] = object_default
		case *ArrayType: elements[i] = array_default
		default:
			fatal("Not a reference type")
		}
	}
	return &j_array{atype: componentType, length: length, elements: elements}
}

func newByteArray(length j_int) *j_array {
	elements := make([]t_any, length)
	for i, _:= range elements {
		elements[i] = byte_default
	}
	return &j_array{atype: BYTE_TYPE, length: length, elements: elements}
}

func newShortArray(length j_int) *j_array {
	elements := make([]t_any, length)
	for i, _:= range elements {
		elements[i] = short_default
	}
	return &j_array{atype: SHORT_TYPE, length: length, elements: elements}
}

func newCharArray(length j_int) *j_array {
	elements := make([]t_any, length)
	for i, _:= range elements {
		elements[i] = char_default
	}
	return &j_array{atype: CHAR_TYPE, length: length, elements: elements}
}

func newIntArray(length j_int) *j_array {
	elements := make([]t_any, length)
	for i, _:= range elements {
		elements[i] = int_default
	}
	return &j_array{atype: INT_TYPE, length: length, elements: elements}
}

func newLongArray(length j_int) *j_array {
	elements := make([]t_any, length)
	for i, _:= range elements {
		elements[i] = long_default
	}
	return &j_array{atype: LONG_TYPE, length: length, elements: elements}
}

func newFloatArray(length j_int) *j_array {
	elements := make([]t_any, length)
	for i, _:= range elements {
		elements[i] = float_default
	}
	return &j_array{atype: FLOAT_TYPE, length: length, elements: elements}
}

func newDoubleArray(length j_int) *j_array {
	elements := make([]t_any, length)
	for i, _:= range elements {
		elements[i] = double_default
	}
	return &j_array{atype: DOUBLE_TYPE, length: length, elements: elements}
}

func newBooleanArray(length j_int) *j_array {
	elements := make([]t_any, length)
	for i, _:= range elements {
		elements[i] = boolean_default
	}
	return &j_array{atype: BOOLEAN_TYPE, length: length, elements: elements}
}

func newObjectArray(componentClass *ClassType, length j_int) *j_array {
	elements := make([]t_any, length)
	for i, _:= range elements {
		elements[i] = object_default
	}
	return &j_array{atype: componentClass, length: length, elements: elements}
}

func newArrayArray(arrayComponentType *ArrayType, length j_int) *j_array {
	elements := make([]t_any, length)
	for i, _:= range elements {
		elements[i] = object_default
	}
	return &j_array{atype: arrayComponentType, length: length, elements: elements}
}


/*--------------------Extend JDK Class here--------------------------*/
// there objects are actually reference to raw j_object, here provide some utility methods
// these structs themselves only contain raw pointer, so they are reference -> just like other value type.
// They are designed as READ-ONLY, because we don't want referee can be changed.
// If we need a new reference, just create a new one with no effort. So never use their pointers
// all these structs are always used receiver, instead its pointer.
// reference will never be null, but its referee may be null.
// create a null reference: java_lang_object{nil}
type (
	java_lang_object struct {*j_object }
	java_lang_string struct {*j_object }
	java_lang_class struct {*j_object }
	java_lang_classloader struct {*j_object }
	java_lang_thread struct {*j_object }
)

func newJavaLangString(str string) java_lang_string {
	java_lang_string_Class := bootstrapClassLoader.load("java/lang/String")
	object := java_lang_string_Class.newObject()
	// convert a utf8 string to utf-16 using as Java String
	chars := []j_char{}
	for _, codepoint := range str {
		if codepoint <= 0xFFFF {
			chars = append(chars, j_char(codepoint))
		} else {
			/*
				H = (S - 10000) / 400 + D800
				L = (S - 10000) % 400 + DC00
			 */
			high_surrogate := j_char((uint32(codepoint) - 0x10000) / 0x400 + 0xD800)
			low_surrogate := j_char((uint32(codepoint) - 0x10000) % 0x400 + 0xDC00)
			chars = append(chars, high_surrogate, low_surrogate)
		}
	}
	charArray := newCharArray(j_int(len(chars)))
	for i, _ := range chars {
		charArray.elements[i] = chars[i]
	}
	object.fields[0] = charArray
	return java_lang_string{object}
}

func (this java_lang_string) toString() string  {
	runes := []rune{}
	chars := this.fields[0].(*j_array).elements
	for i:=0; i < len(chars); i++ {
		char := chars[i].(j_char)
		if char >= 0xD800 && char <= 0xDBFF {
			h := char
			if i+1 < len(chars) && chars[i+1].(j_char) >= 0xDC00 && chars[i+1].(j_char) <= 0xDFFF {
				l := chars[i+1].(j_char)
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

func newJavaLangClass() java_lang_class {
	java_lang_class_Class := bootstrapClassLoader.load("java/lang/Class")
	return java_lang_class{java_lang_class_Class.newObject()}
}