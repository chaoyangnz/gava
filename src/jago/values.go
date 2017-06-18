package jago

type Value interface {
	Type() Type
}

type (
	jbyte int8
	jshort int16
	jchar uint16
	jint int32
    jlong int64
    jfloat float32
    jdouble float64
	jboolean int16
)

func (this jbyte) Type() Type { return BYTE_TYPE }
func (this jshort) Type() Type { return SHORT_TYPE }
func (this jchar) Type() Type { return CHAR_TYPE }
func (this jint) Type() Type { return INT_TYPE }
func (this jlong) Type() Type { return LONG_TYPE }
func (this jfloat) Type() Type { return FLOAT_TYPE }
func (this jdouble) Type() Type { return DOUBLE_TYPE }
func (this jboolean) Type() Type { return BOOLEAN_TYPE }



type Reference interface {
	Value
	Class() ClassType
}

type Object struct {
	class        *Class
	instanceVars []Value
}

func (this *Object) Type() Type {
	return this.Class()
}

func (this *Object) Class() ClassType {
	return this.class
}

type Array struct {
	class       *ArrayClass
	length      jint
	elements    []Value
}

func (this *Array) Type() Type {
	return this.Class()
}

func (this *Array) Class() ClassType {
	return this.class
}

type (
	JavaLangClass struct {*Object}
	JavaLangObject struct {*Object}
	JavaLangString struct {*Object}
)

func (this JavaLangString) toString() string  {
	runes := []rune{}
	chars := this.instanceVars[0].(*Array).elements
	for i:=0; i < len(chars); i++ {
		char := chars[i].(jchar)
		if char >= 0xD800 && char <= 0xDBFF {
			h := char
			if i+1 < len(chars) && chars[i+1].(jchar) >= 0xDC00 && chars[i+1].(jchar) <= 0xDFFF {
				l := chars[i+1].(jchar)
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

func NewJavaLangClass() JavaLangClass {
	class := BOOTSTRAP_CLASSLOADER.CreateClass("java/lang/Class").(*Class)
	object := class.NewObject()

	return JavaLangClass{object}
}

func NewJavaLangString(str string) JavaLangString {
	// check string table
	if strObj, found := STRING_TABLE[str]; found {
		return strObj
	}
	class := BOOTSTRAP_CLASSLOADER.CreateClass("java/lang/String").(*Class)
	object := class.NewObject()

	// convert rune (int32) to Java char (UTF-16 with surrogate)
	chars := []jchar{}
	for _, codepoint := range str {
		if codepoint <= 0xFFFF {
			chars = append(chars, jchar(codepoint))
		} else {
			/*
				H = (S - 10000) / 400 + D800
				L = (S - 10000) % 400 + DC00
			 */
			high_surrogate := jchar((uint32(codepoint) - 0x10000) / 0x400 + 0xD800)
			low_surrogate := jchar((uint32(codepoint) - 0x10000) % 0x400 + 0xDC00)
			chars = append(chars, high_surrogate, low_surrogate)
		}
	}
	// create value field
	charArrayClass := BOOTSTRAP_CLASSLOADER.CreateClass("[C").(*ArrayClass)
	values := charArrayClass.NewArray(jint(len(chars)))
	for i, _ := range values.elements {
		values.elements[i] = chars[i]
	}
	object.instanceVars[0] = values

	// create hashing field?
	// TODO

	return JavaLangString{object}
}