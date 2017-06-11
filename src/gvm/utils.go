package gvm

import (
	"fmt"
	"strings"
)

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

const (
	ALL     = 0
	TRACE   = 1
	DEBUG	= 2
	INFO    = 3
	WARN    = 4
	ERROR   = 5
	FATAL   = 6
)

const logLevel = TRACE

func log(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	fmt.Println("")
}

func all(format string, args ...interface{})   {
	if logLevel <= ALL {
		log(format, args...)
	}
}

func trace(format string, args ...interface{})   {
	if logLevel <= TRACE {
		log(format, args...)
	}
}

func debug(format string, args ...interface{})   {
	if logLevel <= DEBUG {
		log(format, args...)
	}
}

func info(format string, args ...interface{})   {
	if logLevel <= INFO {
		log(format, args...)
	}
}

func warn(format string, args ...interface{})   {
	if logLevel <= WARN {
		log(format, args...)
	}
}

func error(format string, args ...interface{})   {
	if logLevel <= ERROR {
		log(format, args...)
	}
}

func fatal(format string, args ...interface{})   {
	//if logLevel <= FATAL {
		panic(fmt.Sprintf(format, args...))
	//}
}

