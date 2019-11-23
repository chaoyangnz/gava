package javo

import (
	"strings"
	"fmt"
	"github.com/petermattis/goid"
)

func u16toi32(i uint16) int32 {
	return int32(uint32(i))
}

func numberWithSign(i int32) string {
	if i >= 0 {
		return fmt.Sprintf("%s%d", "+", i)
	} else {
		return fmt.Sprintf("%s%d", "-", -i)
	}
}

func repeat(str string, times int) string {
	return strings.Repeat(str, times)
}

/*
A Java try {} catch() {} finally {} block
 */
type Block struct {
	try     func()
	catch   func(throwable Reference) // throwable never be nil
	finally func()
}

func (tcf Block) Do() {
	if tcf.finally != nil {
		defer tcf.finally()
	}
	if tcf.catch != nil {
		defer func() {
			if r := recover(); r != nil {
				if throwable, ok := r.(Reference); ok {
					tcf.catch(throwable)
				} else {
					// otherwise, the whole project has non-throwable panic,
					// But some 3rd-party package can panic other non-throwable
					Bug("Javo project has never non-throwable panic. "+
						"There is some 3rd-party package doing a non-throwable panic, check it. "+
						"Original panic: \n%v", r)
				}
			}
		}()
	}
	tcf.try()
}

func getGID() int64 {
	return goid.Get()
}

func binaryName2JavaName(name string) JavaLangString {
	return VM.NewJavaLangString(strings.Replace(name, "/", ".", -1))
}

func javaName2BinaryName(name JavaLangString) string {
	return javaName2BinaryName0(name.toNativeString())
}

func javaName2BinaryName0(name string) string {
	return strings.Replace(name, ".", "/", -1)
}
