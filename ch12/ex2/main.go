package main

import (
	"fmt"
	"math/rand"
)

func main() {
	myFunc()
}

func myFunc() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		defer close(ch1)
		for i := range 10 {
			ch1 <- i + rand.Int()
		}
	}()
	go func() {
		defer close(ch2)
		for i := range 10 {
			ch2 <- i + rand.Int()
		}
	}()

	count := 2
	for count != 0 {
		select {
		case out, ok := <-ch1:
			if !ok {
				count--
				ch1 = nil
				break
			}
			fmt.Println("From goroutine 1:", out)
		case out, ok := <-ch2:
			if !ok {
				count--
				ch2 = nil
				break
			}
			fmt.Println("From goroutine 2:", out)
		}
	}
}
