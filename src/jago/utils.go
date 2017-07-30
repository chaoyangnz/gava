package jago

import (
	"strings"
	"fmt"
	"runtime"
	"bytes"
	"strconv"
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
					Bug("Jago project has never non-throwable panic. " +
						"There is some 3rd-party package doing a non-throwable panic, check it. " +
						"Original panic: \n%v", r)
				}
			}
		}()
	}
	tcf.try()
}

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func vmName2JavaName(name string) string {
	return strings.Replace(name, "/", ".", -1)
}