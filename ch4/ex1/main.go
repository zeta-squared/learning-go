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

	fmt.Println(x)
}
