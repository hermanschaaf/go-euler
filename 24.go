/*

   Problem 24:
   What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?

   We will solve this using the [Factoradic number system](http://en.wikipedia.org/wiki/Factoradic)
   It's blisteringly fast: O(N) time, where N is number of digits in the number we are looking for 
   (in this case, 10). 

*/

package main

import (
	"fmt"
	"strconv"
)

func factorial(n int) (factorial int) {
	factorial = 1
	for i := n; i >= 2; i-- {
		factorial *= i
	}
	return factorial
}

func main() {
	// define the base
	base := 10

	// compute the base
	var baseNums [10]int
	for i := base - 1; i >= 0; i-- {
		baseNums[base-1-i] = factorial(i)
	}

	// find the digits of the Factoradic number for 999999
	var digits [10]int
	for i, n := base-1, 999999; i >= 0; i-- {
		remain := n % baseNums[base-1-i]
		digits[base-1-i] = (n - remain) / baseNums[base-1-i]
		n = remain
	}

	// now, display these numbers:
	numbers := []int{}
	for i := 0; i < base; i++ {
		numbers = append(numbers, i)
	}

	ans := ""
	for i := 0; i < len(digits); i++ {
		ans += strconv.FormatInt(int64(numbers[digits[i]]), 10)
		numbers = append(numbers[:digits[i]], numbers[digits[i]+1:]...)
	}

	fmt.Println(ans)
}
