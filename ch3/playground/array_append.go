package main

import "fmt"

func arrayAppend() {
	x := [4]int{1, 2, 3, 4}
	y := x[:2]
	z := x[2:]

	fmt.Println("x: ", x)
	fmt.Println("y: ", y)
	fmt.Println("z: ", z)

	y[1] = 1
	z = append(z, 5)

	fmt.Println("x: ", x)
	fmt.Println("y: ", y)
	fmt.Println("z: ", z)
}
