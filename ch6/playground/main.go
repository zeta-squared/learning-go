package main

import "fmt"

func main() {
	x := make([]int, 4, 4)
	x[0] = 1
	x[1] = 2
	x[2] = 3
	x[3] = 4

	x = append(x, []int{1, 2, 3, 4}...)

	fmt.Println(x)
}
