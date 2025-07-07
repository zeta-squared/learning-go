package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
	}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
	}
}

func main() {
	person := MakePerson("Zeaiter", "Zeaiter", 34)
	fmt.Println(person)
	personPointer := MakePersonPointer("Zeaiter", "Zeaiter", 34)
	fmt.Println(personPointer)
}
