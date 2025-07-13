package main

import (
	"fmt"
)

func main() {
	var a MyInt = 10
	var b MyFloat = 3.87298
	PrintToScreen(a)
	PrintToScreen(b)
}

func PrintToScreen[T Printable](p T) {
	fmt.Println(p.String())
}

type Printable interface {
	~int | ~float64
	fmt.Stringer
}

type MyInt int

func (x MyInt) String() string {
	return fmt.Sprintf("The value of the Printable is %d", x)
}

type MyFloat float64

func (f MyFloat) String() string {
	return fmt.Sprintf("The value of the Printable is %f", f)
}
