package main

import "fmt"

func main() {
	var di DoubleInt = 10
	var di2 *DoubleInt = &di
	DoublerCompare(&di, di2)
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
