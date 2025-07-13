package main

import (
	"cmp"
	"fmt"
)

func main() {
	var a uint8 = 255
	b := PlusOneHundred(a)
	fmt.Println(b)
	fmt.Println(cmp.Compare("string1", "string2"))
}

type Integer interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

func Convert[T1, T2 Integer](in T1) T2 {
	return T2(in)
}

func PlusOneHundred[T Integer](in T) T {
	return in + 100
}
