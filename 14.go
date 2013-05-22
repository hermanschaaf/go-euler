/*

   Problem: 14
   Which starting number, under one million, produces the longest Collatz chain?

   We'll tackle this one with a DP solution.

*/

package main

import (
	"fmt"
)

var chains = make([]int64, 1000000)

func getChainLength(n int64) int64 {
	if n < (int64)(len(chains)) {
		if chains[n] != 0 {
			return chains[n]
		}
	} else if n == 1 {
		return 0
	}
	var count int64 = 0
	for {
		if n%2 == 1 {
			break
		}
		count += 1
		n /= 2
	}
	if n == 1 {
		return count
	}
	n = n*3 + 1
	count += 1
	return count + getChainLength(n)
}

func getLongestChain(max int64) int64 {
	var maxLength, startingNumber, i int64 = 0, 0, 2
	for i = 2; i < max; i++ {
		chainLength := getChainLength(i)
		if i < 1000000 {
			chains[i] = chainLength
		}
		if chainLength > maxLength {
			maxLength = chainLength
			startingNumber = i
		}
	}
	return startingNumber
}

func main() {
	var maxNumber int64 = 1000000
	fmt.Println(getLongestChain(maxNumber))
}
