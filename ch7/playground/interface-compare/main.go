package main

import "fmt"

func main() {
	var di DoubleInt = 10
	var di2 DoubleInt = 10
	// var di2 *DoubleInt = &di
	var db DoubleBrie = 10
	fmt.Println(di == di2)   // User types are comparable by type and value
	DoublerCompare(&di, &db) // Interfaces allow comparison functions to operate on more than one user type
}

func DoublerCompare(d1, d2 Doubler) {
	fmt.Println(d1 == d2)
}

type Doubler interface {
	Double()
}

type DoubleInt int

func (d *DoubleInt) Double() {
	*d = 2 * *d
}

type DoubleBrie int

func (d *DoubleBrie) Double() {
	*d = 2 * *d
}
