package main

import "fmt"

func UpdateSlice(mySlice []string, myString string) {
	l := len(mySlice)
	mySlice[l-1] = myString
	fmt.Println(mySlice)
}

func GrowSlice(mySlice []string, myString string) {
	mySlice = append(mySlice, myString)
	fmt.Println(mySlice)
}

func main() {
	mySlice := []string{"Zeaiter", "Tupou"}
	myString := "Isaac"
	fmt.Println("Before UpdateSlice:", mySlice)
	UpdateSlice(mySlice, myString)
	fmt.Println("After UpdateSlice:", mySlice)
	fmt.Println("Before GrowSlice:", mySlice)
	GrowSlice(mySlice, myString)
	fmt.Println("After GrowSlice:", mySlice)
}
