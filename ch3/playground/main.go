package main

import "fmt"

func main() {
	fmt.Print("=====================\n\n")
	fmt.Print("Slice Test:\n\n")
	fmt.Print("=====================\n\n")
	sliceTest()
	fmt.Println()

	fmt.Print("=====================\n\n")
	fmt.Print("String Slice:\n\n")
	fmt.Print("=====================\n\n")
	stringSlice()
	fmt.Println()

	fmt.Print("=====================\n\n")
	fmt.Print("Slice Append:\n\n")
	fmt.Print("=====================\n\n")
	sliceAppend()
	fmt.Println()

	fmt.Print("=====================\n\n")
	fmt.Print("Array Append:\n\n")
	fmt.Print("=====================\n\n")
	arrayAppend()
	fmt.Println()

	fmt.Print("=====================\n\n")
	fmt.Print("Compare Maps:\n\n")
	fmt.Print("=====================\n\n")
	compareMaps()
	fmt.Println()
}
