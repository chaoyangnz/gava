package jago

import (
	"fmt"
	"errors"
	"os"
	"bufio"
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



func NewLog(logfile string) *Log {
	var _, err = os.Stat(logfile)

	// create file if not exists
	if os.IsNotExist(err) {
		os.Create(logfile)
	} else {
		os.Remove(logfile)
		os.Create(logfile)
	}

	f, err := os.OpenFile(logfile, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("File does not exists or cannot be created")
		os.Exit(1)
	}

	w := bufio.NewWriter(f)

	log := &Log{writer: w}
	return log
}

type Log struct {
	writer *bufio.Writer
}

func (this *Log) _log(format string, args ...interface{}) {
	fmt.Fprintf(this.writer, format, args...)
	this.writer.Flush()
}

func (this *Log)  All(format string, args ...interface{})   {
	if LOG_LEVEL <= ALL {
		this._log(format, args...)
	}
}

func (this *Log)  Trace(format string, args ...interface{})   {
	if LOG_LEVEL <= TRACE {
		this._log(format, args...)
	}
}

func (this *Log)   Debug(format string, args ...interface{})   {
	if LOG_LEVEL <= DEBUG {
		this._log(format, args...)
	}
}

func (this *Log)  Info(format string, args ...interface{})   {
	if LOG_LEVEL <= INFO {
		this._log(format, args...)
	}
}

func (this *Log)  Warn(format string, args ...interface{})   {
	if LOG_LEVEL <= WARN {
		this._log(format, args...)
	}
}

func (this *Log)  Error(format string, args ...interface{})   {
	if LOG_LEVEL <= ERROR {
		this._log(format, args...)
	}
}

func Fatal(format string, args ...interface{})   {
	//if logLevel <= FATAL {
	panic(fmt.Sprintf(format, args...))
	//}
}

func Bug(format string, args ...interface{})   {
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

func JavaOutPrintf(format string, args ...interface{})  {
	fmt.Fprintf(os.Stdout, format, args...)
}

func JavaErrPrintf(format string, args ...interface{})  {
	fmt.Fprintf(os.Stderr, format, args...)
}
