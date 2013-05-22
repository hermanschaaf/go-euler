package main

import "fmt"

func main() {
	var numberOfPrimes, number, maxLimit int64 = 1, 1, 114320
	var sieve [114320]bool
	var i, j int64 = 0, 0
	for i = 3; i < maxLimit; i += 2 {
		if sieve[i] == true {
			continue
		}

		numberOfPrimes++

		if numberOfPrimes == 10001 {
			number = i
			break
		}

		for j = i * i; j < maxLimit; j += 2 * i {
			sieve[j] = true
		}
	}
	fmt.Println(number)
}
