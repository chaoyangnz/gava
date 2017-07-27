package jago

import (
	"fmt"
	"errors"
	"os"
	"bufio"
	"runtime/debug"
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



func NewLog(category string, level int, logfile string) *Log {
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
		Fatal("File does not exists or cannot be created")
	}

	w := bufio.NewWriter(f)

	log := &Log{category, level, w}
	return log
}

type Log struct {
	category string
	level int
	writer *bufio.Writer
}

func (this *Log) _log(format string, args ...interface{}) {
	fmt.Fprintf(this.writer, format, args...)
	this.writer.Flush()
}

func (this *Log)  All(format string, args ...interface{})   {
	if this.level <= ALL {
		this._log(format, args...)
	}
}

func (this *Log)  Trace(format string, args ...interface{})   {
	if this.level <= TRACE {
		this._log(format, args...)
	}
}

func (this *Log)   Debug(format string, args ...interface{})   {
	if this.level <= DEBUG {
		this._log(format, args...)
	}
}

func (this *Log)  Info(format string, args ...interface{})   {
	if this.level <= INFO {
		this._log(format, args...)
	}
}

func (this *Log)  Warn(format string, args ...interface{})   {
	if this.level <= WARN {
		this._log(format, args...)
	}
}

func (this *Log)  Error(format string, args ...interface{})   {
	if this.level <= ERROR {
		this._log(format, args...)
	}
}

func Fatal(format string, args ...interface{})   {
	fmt.Fprintf(os.Stderr, "VM internal error: ")
	fmt.Fprintf(os.Stderr, format+ "\n\n ------------------------\n", args...)
	os.Stderr.Write(debug.Stack())
	os.Exit(2)
}

func Bug(format string, args ...interface{})   {
	fmt.Fprintf(os.Stderr, "VM implmentation bug: ")
	fmt.Fprintf(os.Stderr, format + "\n\n ------------------------\n", args...)

	os.Stderr.Write(debug.Stack())
	os.Exit(3)
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
