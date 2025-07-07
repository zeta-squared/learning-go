package main

import "fmt"

func sliceTest() {
	var x []int
	fmt.Println(x, len(x), cap(x))

	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150)
	fmt.Println(x, len(x), cap(x))

	clear(x)
	fmt.Println(x, len(x), cap(x))
}
