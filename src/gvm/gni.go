package gvm

type GNIEnv struct {}

type (
	jbyte       java_byte
	jshort      java_short
	jchar       java_char
	jint        java_int
	jlong       java_long
	jfloat      java_float
	jdouble     java_double

	jobject     java_object
	jstring     java_object
	jclass      java_object
	jarray      java_array
)

func (this *GNIEnv) GetObjectClass(object jobject) jclass {
	return object.class.classObject
}

/*------------jstring----------------------------------*/
func NewJavaString(chars []java_char) jstring  {
	stringClass := classCache["java/lang/String"]
	if stringClass == nil {
		stringClass = bootstrapClassLoader.load("java/lang/String")
	}
	object := stringClass.new()
	object.fields[0] = NewCharArray(chars)
	return jstring(object)
}

func jtring2string(jstring jstring) string {
	runes := []rune{}
	chars := jstring.fields[0].(java_array).elements
	for i:=0; i < len(chars); i++ {
		char := chars[i].(java_char)
		if char >= 0xD800 && char <= 0xDBFF {
			h := char
			if i+1 < len(chars) && chars[i+1].(java_char) >= 0xDC00 && chars[i+1].(java_char) <= 0xDFFF {
				l := chars[i+1].(java_char)
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

func GVM_print(s jstring) {
	println(jtring2string(s))
}




