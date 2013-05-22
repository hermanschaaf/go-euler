package main

import "fmt"

func getSum(n int) int {
	return n * (n + 1) / 2
}

func getSumOfSquares(n int) int {
	return n * (2*n + 1) * (n + 1) / 6
}

func main() {
	var n = 100
	var squareOfSums = getSum(n) * getSum(n)
	var sumOfSquares = getSumOfSquares(n)
	fmt.Println(squareOfSums - sumOfSquares)
}
