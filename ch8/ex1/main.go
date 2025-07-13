package main

import "fmt"

func main() {
	fmt.Println(Double(10))
	fmt.Println(Double(7.82))
}

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func Double[T Number](x T) T {
	return 2 * x
}
