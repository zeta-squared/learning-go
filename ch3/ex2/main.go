package main

import "fmt"

func main() {
	message := "Hi 👩 and 👨"
	var c rune = []rune(message)[3]

	fmt.Println(string(c))
}
