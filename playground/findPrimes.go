package main

func findPrimes(n int) []int {
	primes := make([]int, 1, n)
	primes[0] = 2

	for i := 3; i <= n; i++ {
		isPrime := true

		for _, v := range primes {
			if i%v == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primes = append(primes, i)
		}
	}

	return primes
}
