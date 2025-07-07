package main

func main() {
	// Using make does not show the effect on the garbage collector as the memory is allocated at compile
	// time.

	// people := make([]Person, 0, 10_000_000)
	var people []Person
	for i := 0; i < 10_000_000; i++ {
		people = append(people, MakePerson("Zeaiter", "Zeaiter", 34))
	}
}

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
