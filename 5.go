package main

import "fmt"

func main() {
	var factors = [...]int{2, 3, 2, 5, 2, 7, 3, 11, 13, 2, 17, 19}
	var num int = 1
	for key, _ := range factors {
		num *= factors[key]
	}
	fmt.Println(num)
}
