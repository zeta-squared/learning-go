package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	ch := make(chan int)
	var wg1 sync.WaitGroup
	wg1.Add(2)
	for i := range 2 {
		go func() {
			defer wg1.Done()
			for j := range 10 {
				ch <- j + i + rand.Int()
			}
			fmt.Println("Exiting goroutine", i+1)
		}()
	}

	go func() {
		wg1.Wait()
		close(ch)
	}()

	var wg2 sync.WaitGroup
	wg2.Add(1)
	go func() {
		defer wg2.Done()

		for v := range ch {
			fmt.Println("Reading value:", v)
		}
		fmt.Println("Exiting goroutine 3")
	}()
	wg2.Wait()
}
