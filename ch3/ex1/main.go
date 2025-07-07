package main

import "fmt"

func main() {
	greetings := []string{"Hello", "Hola", "नमस्ते", "こんにちは", "привет"}
	a := greetings[:2]
	b := greetings[2:5]
	c := greetings[4:]

	fmt.Print("greetings = ", greetings, "\na = ", a, "\nb = ", b, "\nc = ", c, "\n")
}
