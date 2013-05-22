package main

import (
	"fmt"
)

func calcFactorialSum(num int) int {
	// implement multiplication in base 100000,
	// then sum the digits

	var dig [10000]int
	first, last, carry, n, x, sum := 0, 0, 0, 0, 0, 0
	dig[0] = 1
	for n = 2; n <= num; n++ {
		carry = 0
		for x = first; x <= last; x++ {
			carry = dig[x]*n + carry
			dig[x] = carry % 100000
			if x == first && carry%100000 == 0 {
				first += 1
			}
			carry /= 100000
		}
		if carry != 0 {
			last += 1
			dig[last] = carry
		}
	}
	for x = first; x <= last; x++ {
		sum += dig[x]%10 + (dig[x]/10)%10 + (dig[x]/100)%10 + (dig[x]/1000)%10 + (dig[x]/10000)%10
	}
	return sum
}

func main() {
	fmt.Println(calcFactorialSum(100))
}
