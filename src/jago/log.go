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
