package main

import "fmt"

func stringSlice() {
	x := make([]string, 0, 4)
	x = append(x, "my", "first", "string", "slice")
	fmt.Println(x)

	y := x[1:]
	fmt.Println(cap(y), cap(x))
}
