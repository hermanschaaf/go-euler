package main

import "fmt"

func isProduct(num int) bool {
	for i := 999; i >= 100; i-- {
		if (num%i == 0) && (num/i >= 100) && (num/i <= 999) {
			return true
		}
	}
	return false
}

func main() {
	var palindrome int = 0
	for n := 999; n >= 100; n-- {
		palindrome = 1100*n - 990*(n/10) - 99*(n/100)
		if isProduct(palindrome) {
			fmt.Println(palindrome)
			return
		}
	}

}
