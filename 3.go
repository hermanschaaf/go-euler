package main

import "fmt"

func isPrime(num int) bool {
	for i := 2; i < num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var num int = 600851475143
	for i := 2; i < num/2; i++ {
		if num%i == 0 {
			var pair int = num / i
			if isPrime(pair) {
				fmt.Println(pair)
				return
			}
		}
	}
}
