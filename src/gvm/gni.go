package gvm

type GNIEnv struct {}

type (
	// these types are the vm bridge for java language types
	//jbyte                   t_byte
	//jshort                  t_short
	//jchar                   t_char
	//jint                    t_int
	//jlong                   t_long
	//jfloat                  t_float
	//jdouble                 t_double
	//
	//jarray                  t_array

	java_lang_object        t_object
	java_lang_string        t_object
	java_lang_class         t_object
	java_lang_classloader   t_object
)

func (this *GNIEnv) GetObjectClass(object java_lang_object) java_lang_class {
	return object.class.classObject
}

/*------------java_lang_string----------------------------------*/
//func newJavaLangString(chars []t_char) java_lang_string {
//	stringClass := bootstrapClassLoader.loadClass("java/lang/String")
//	object := stringClass.new()
//	object.fields[0] = newCharArray(chars)
//	return java_lang_string(object)
//}

func newJavaLangString(str string) java_lang_string {
	chars := _string2jchars(str)
	stringClass := bootstrapClassLoader.load("java/lang/String")
	object := stringClass.new()
	object.fields[0] = newCharArray(chars)
	return java_lang_string(object)
}

/*
H = (S - 10000) / 400 + D800
L = (S - 10000) % 400 + DC00
 */
func _rune2jchar(codepoint rune) []t_char {
	if codepoint <= 0xFFFF {
		return []t_char{t_char(codepoint)}
	}
	high_surrogate := (uint32(codepoint) - 0x10000) / 0x400 + 0xD800
	low_surrogate := (uint32(codepoint) - 0x10000) % 0x400 + 0xDC00
	return []t_char{t_char(high_surrogate), t_char(low_surrogate)}
}

func _string2jchars(str string) []t_char {
	runes := []rune(str)

	chars := []t_char{}
	for i := 0; i < len(runes); i++ {
		arr := _rune2jchar(runes[i])
		for k := 0; k < len(arr); k++ {
			chars = append(chars, arr[k])
		}
	}
	return chars
}

func javaLangString2string(jstring java_lang_string) string  {
	runes := []rune{}
	chars := jstring.fields[0].(t_array).elements
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

/* -----------java_lang_class ----------------*/

func newJavaLangClass() java_lang_class {
	class := bootstrapClassLoader.load("java/lang/Class")
	return java_lang_class(class.new())
}

/* ---------- JDK Native methods implementation ---*/
func GVM_print(s java_lang_string) {
	println(javaLangString2string(s))
}




