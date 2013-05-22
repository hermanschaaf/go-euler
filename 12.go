/*

   Problem 12:
   What is the value of the first triangle number to have over five hundred divisors?

*/

package main

import "fmt"

func numDivisors(n int) int {
	// find number of divisors of number
	if n%2 == 0 {

		// ignore one divisor of 2 because F(n) = n(n+1)/2
		n = n / 2
	}
	divisors := 1
	count := 0

	// get count of factors of 2
	for {
		if n%2 != 0 {
			break
		}
		count += 1
		n = n / 2
	}

	divisors = divisors * (count + 1)

	// now get count of factors for all other numbers
	// until n is 1:
	p := 3
	for {
		if n == 1 {
			break
		}

		count = 0
		for {
			if n%p != 0 {
				break
			}
			count += 1
			n = n / p
		}
		divisors = divisors * (count + 1)
		p += 2
	}
	return divisors
}

func findTriangularIndex(factorLimit int) int {
	// go through all n until we find one with 
	// factor total > factorLimit
	n := 1

	// total number of divisors is the product
	// of number of divisors for n and n+1, because
	// F(n) = n(n+1) / 2
	lnum, rnum := numDivisors(n), numDivisors(n+1)
	for {
		if lnum*rnum > factorLimit {
			break
		}
		n += 1
		lnum, rnum = rnum, numDivisors(n+1)
	}
	return n
}

func main() {
	var minFactors int = 500
	num := findTriangularIndex(minFactors)

	// now use the formula to calculate value
	triangle := (num * (num + 1)) / 2
	fmt.Println(triangle)
}
