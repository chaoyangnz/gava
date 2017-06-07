package gvm

import "strings"

func u2b(u1s []u1) []uint8 {
	bytes := make([]uint8, len(u1s))
	for i := 0; i < len(bytes); i++ {
		bytes[i] = uint8(u1s[i])
	}
	return bytes
}

func b2u(bytes []uint8) []u1 {
	bs := make([]u1, len(bytes))
	for i := 0; i < len(bytes); i++ {
		bs[i] = u1(bytes[i])
	}
	return bs
}

func u2s(u1s []u1) string {
	return string(u2b(u1s))
}

func bytes2uint16(bytes []uint8) uint16 {
	return uint16((bytes[0] << 8) | bytes[1])
}

func bytes2uint32(bytes []uint8) uint32 {
	return uint32((bytes[0] << 24) | (bytes[1] << 16) | (bytes[2] << 8) | bytes[3])
}

/*
H = (S - 10000) / 400 + D800
L = (S - 10000) % 400 + DC00
 */
func rune2utf16(codepoint rune) []java_char {
	if codepoint <= 0xFFFF {
		return []java_char{ java_char(codepoint)}
	}
	high_surrogate := (uint32(codepoint) - 0x10000) / 0x400 + 0xD800
	low_surrogate := (uint32(codepoint) - 0x10000) % 0x400 + 0xDC00
	return []java_char{java_char(high_surrogate), java_char(low_surrogate)}
}

func bytes2JavaChars(bytes []uint8) []java_char  {
	return string2JavaChars(string(bytes))
}

func string2JavaChars(str string) []java_char  {
	runes := []rune(str)

	chars := []java_char{}
	for i := 0; i < len(runes); i++ {
		arr := rune2utf16(runes[i])
		for j := 0; j < len(arr); j++ {
			chars = append(chars, arr[j])
		}
	}
	return chars
}

func parametersAndReturn(descriptor string) ([]string, string) {
	arr := strings.Split(descriptor, ")")
	parametersStr := arr[0][1:]
	returnStr := arr[1]

	var parametersSlice []string

	for i := 0; i < len(parametersStr); {
		ch := rune(parametersStr[i])
		switch ch {
		case 'B', 'C', 'D', 'F', 'I', 'J', 'S', 'Z':
			parametersSlice = append(parametersSlice, string(ch))
			i++
		case 'L':
		Ref: for j := i+1; j < len(parametersStr); j++ {
			switch rune(parametersStr[j]) {
			case ';':
				parametersSlice = append(parametersSlice, string(parametersStr[i:j+1]))
				i = j+1
				break Ref
			}
		}
		case '[':
		Arr: for j := i+1; j < len(parametersStr); j++ {
			switch rune(parametersStr[j]) {
			case '[':
				continue
			case 'B', 'C', 'D', 'F', 'I', 'J', 'S', 'Z':
				parametersSlice = append(parametersSlice, string(parametersStr[i:j+1]))
				i = j+1
				break Arr
			case 'L':
				for k := j+1; j < len(parametersStr); k++ {
					switch rune(parametersStr[k]) {
					case ';':
						parametersSlice = append(parametersSlice, string(parametersStr[i:k+1]))
						i = k+1
						break Arr
					}
				}
			}
		}
		}
	}

	return parametersSlice, returnStr
}