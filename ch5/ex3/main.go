package ex3

import "fmt"

func main() {
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))
	fmt.Println(helloPrefix("Maria"))
	fmt.Println()
}

func prefixer(s string) func(string) string {
	return func(a string) string {
		return s + " " + a
	}
}
