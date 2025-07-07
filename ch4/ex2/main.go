package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := make([]int, 100)

	for i := range 100 {
		x[i] = rand.Intn(100)
	}

	for _, v := range x {
		switch {
		case v%6 == 0:
			fmt.Println("Six!")
		case v%3 == 0:
			fmt.Println("Three!")
		case v%2 == 0:
			fmt.Println("Two!")
		default:
			fmt.Println("Never mind")
		}
	}
}
