package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	sqrtCached := sqrtMapCached()
	for i := range 100_000 {
		if i%1000 == 0 {
			fmt.Println(sqrtCached[i])
		}
	}
}

func sqrtMap() map[int]float64 {
	sqrt := make(map[int]float64, 100_000)
	for i := range 100_000 {
		sqrt[i] = math.Sqrt(float64(i))
	}

	return sqrt
}

var sqrtMapCached func() map[int]float64 = sync.OnceValue(sqrtMap)
