package main

import "fmt"

var primes []int64 = []int64{17, 13, 11, 7, 5, 3, 2}
var digits []int64 = []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var numbers []int64 = []int64{}

func pow(n int64, p int64) int64 {
	if p == 0 {
		return 1
	}
	num := n
	var i int64
	for i = 1; i < p; i++ {
		num *= n
	}
	return num
}

func hasAnyDigit(num int64, digits int64) bool {
	for {
		dig := digits % 10
		n := num
		for {
			d := n % 10
			if d == dig {
				return true
			}
			n /= 10
			if n == 0 {
				return false
			}
		}
		digits /= 10
		if digits == 0 {
			return false
		}
	}
	return false
}

func recurse(n int64, index int64, hasZero bool) {
	if n < 0 {
		return
	}
	if index >= int64(len(primes)) {
		numbers = append(numbers, n)
		return
	}

	for _, i := range digits {
		num := n
		if !hasAnyDigit(num, i) && !(i == 0 && hasZero) {
			num = i*pow(10, index+3) + num
			if index == int64(len(primes)-1) {
				numbers = append(numbers, num)
				return
			} else if (num/pow(10, index+1))%primes[index+1] == 0 {
				if i == 0 {
					recurse(num, index+1, true)
				} else {
					recurse(num, index+1, false)
				}

			}
		}
	}
}

func main() {
	for u := primes[0]; u < 1000; u += primes[0] {
		if (u/10)%10 != (u/100)%10 && (u%10) != u/100 {
			recurse(u, 0, false)
		}
	}
	var sum int64 = 0
	for _, i := range numbers {
		sum += i
	}
	fmt.Println(sum)
}
