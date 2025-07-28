package main

import "fmt"

func sliceCopy() {
	x := [][]int{{1, 2}, {3, 4}, {5, 6}}
	y := make([][]int, 3)
	copy(y, x)

	fmt.Println("x: ", x)
	fmt.Println("y: ", y)
	y[2][0] = 6 // copy is a shallow copy
	fmt.Println("x: ", x)
	fmt.Println("y: ", y)
}
