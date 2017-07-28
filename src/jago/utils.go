package jago

import (
	"strings"
	"fmt"
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