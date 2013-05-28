package main

import "fmt"

func pow(n int64, p int64) int64 {
	if p == 0 {
		return 1
	}
	num := n
	var i int64
	for i = 1; i < p; i++ {
		num *= n
		num %= 10000000000
	}
	return num
}

func main() {
	var n int64 = 0
	for i := int64(1); i <= 1000; i++ {
		n += pow(i, i)
		n %= 10000000000
	}
	fmt.Println(n)
}
