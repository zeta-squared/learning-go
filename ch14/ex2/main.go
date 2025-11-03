package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	sum := 0
	for i := 1; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("sum = %d\nnumber of iterations = %d\nexceeded 2 second time limit\n", sum, i)
			return
		default:
			if sum == 1_234 {
				fmt.Printf("sum = %d\nnumber of iterations = %d\nnumber reached\n", sum, i)
				return
			}

			sum += rand.Intn(100_000_000)
		}
	}
}
