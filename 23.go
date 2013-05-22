/*
   Problem 23
   Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.

   Note: This one takes ~1 second to execute, and could really use some optimization

*/

package main

import (
	"fmt"
)

func isAbundant(num int) bool {
	s := 1
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			s += i
		}
		if s > num {
			return true
		}
	}
	return false
}

func calcNonAbundantsSum(max int) int {
	// calculate abundant numbers up to max
	var abundants []int
	for i := 2; i < max; i++ {
		if isAbundant(i) {
			abundants = append(abundants, i)
		}
	}

	// see which integers cannot be made
	var summable [28123]int

	for a := 0; a < len(abundants); a++ {
		for b := 0; b < len(abundants); b++ {
			s := abundants[a] + abundants[b]
			if s < len(summable) {
				summable[s] = 1
			}
		}
	}

	holesSum := 0
	for i := 0; i < len(summable); i++ {
		if summable[i] != 1 {
			holesSum += i
		}
	}

	return holesSum
}

func main() {
	max := 28124
	fmt.Println(calcNonAbundantsSum(max))
}
