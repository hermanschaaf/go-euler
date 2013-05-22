package main

import "fmt"

func main() {
	var numberOfPrimes, sumOfPrimes, maxLimit int64 = 1, 2, 2000000
	var sieve [2000000]bool
	var i, j int64 = 0, 0
	for i = 3; i < maxLimit; i += 2 {
		if sieve[i] == true {
			continue
		}

		numberOfPrimes++
		sumOfPrimes += i

		for j = i * i; j < maxLimit; j += 2 * i {
			sieve[j] = true
		}
	}
	fmt.Println(sumOfPrimes)
}
