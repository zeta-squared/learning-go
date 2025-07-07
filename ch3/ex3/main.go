package main

import "fmt"

func main() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	a := Employee{
		"John",
		"Smith",
		32,
	}
	b := Employee{
		firstName: "Joe",
		lastName:  "Mill",
		id:        34,
	}
	var c Employee
	c.firstName = "Zeaiter"
	c.lastName = "Zeaiter"
	c.id = 1

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
}
