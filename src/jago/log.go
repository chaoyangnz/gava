package jago

import (
	"fmt"
	"errors"
)

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

func Log(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	fmt.Println("")
}

func All(format string, args ...interface{})   {
	if logLevel <= ALL {
		Log(format, args...)
	}
}

func Trace(format string, args ...interface{})   {
	if logLevel <= TRACE {
		Log(format, args...)
	}
}

func Debug(format string, args ...interface{})   {
	if logLevel <= DEBUG {
		Log(format, args...)
	}
}

func Info(format string, args ...interface{})   {
	if logLevel <= INFO {
		Log(format, args...)
	}
}

func Warn(format string, args ...interface{})   {
	if logLevel <= WARN {
		Log(format, args...)
	}
}

func Error(format string, args ...interface{})   {
	if logLevel <= ERROR {
		Log(format, args...)
	}
}

func Fatal(format string, args ...interface{})   {
	//if logLevel <= FATAL {
	panic(fmt.Sprintf(format, args...))
	//}
}

func Throw(exception string, message string) error {
	msg := fmt.Sprintf("%s: %s", exception, message)
	Fatal(msg)
	return errors.New(msg)
}

func NewError(format string, args ...interface{}) error {
	return errors.New(fmt.Sprintf(format, args...))
}

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