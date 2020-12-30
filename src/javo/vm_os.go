package javo

import (
	"github.com/fatih/color"
	"os"
	"time"
)

type OS struct{}

func (this *OS) CurrentTimeMillis() Long {
	return Long(time.Now().UnixNano() / int64(time.Millisecond))
}

func (this *OS) CurrentTimeNano() Long {
	return Long(time.Now().UnixNano())
}

func (this *OS) StdoutPrintf(format string, args ...interface{}) {
	color.New(color.FgGreen).Fprintf(os.Stdout, format, args...)
}

func (this *OS) StderrPrintf(format string, args ...interface{}) {
	color.New(color.FgRed).Fprintf(os.Stderr, format, args...)
}
