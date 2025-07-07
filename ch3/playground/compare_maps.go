package main

import (
	"fmt"
	"maps"
)

func compareMaps() {
	m := map[string]int{
		"hello": 10,
		"world": 5,
	}

	n := map[string]int{
		"hello": 10,
		"world": 5,
	}

	o := map[string]int{
		"hello": 10,
		"word":  5,
	}

	p := map[string]int{
		"hello": 10,
		"world": 10,
	}

	fmt.Println(m, n, o, p)
	fmt.Println("m = n ? ", maps.Equal(m, n))
	fmt.Println("m = o ? ", maps.Equal(m, o))
	fmt.Println("m = p ? ", maps.Equal(m, p))
}
