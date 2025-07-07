package main

import "fmt"

func sliceAppend() {
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")   // x = ["a" "b" "c" "d"]
	y := x[:2:2]                        // y = ["a" "b"]
	z := x[2:4:4]                       // z = ["c" "d"]
	fmt.Println(cap(x), cap(y), cap(z)) // 5, 5, 3
	y = append(y, "i", "j", "k")        // x = ["a" "b" "i" "j"], y = ["a" "b" "i" "j" "k"], z = ["i" "j"]
	x = append(x, "x")                  // x = ["a" "b" "i" "j" "x"], y = ["a" "b" "i" "j" "x"], z = ["i" "j"]
	z = append(z, "y")                  // x = ["a" "b" "i" "j" "y"], y = ["a" "b" "i" "j" "y"], z = ["i" "j" "y"]
	fmt.Println("x: ", x)               // x: ["a" "b" "i" "j" "y"]
	fmt.Println("y: ", y)               // y: ["a" "b" "i" "j" "y"]
	fmt.Println("z: ", z)               // z: ["i" "j" "y"]
}
