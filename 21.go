package main

import "fmt"

func sum(ar []int) int {
	var ans int = 0
	for i := range ar {
		ans += ar[i]
	}
	return ans
}

func calcAmicableNumbersSum(max int) int {
	var numbers [10000]int
	for i := 2; i < max; i++ {
		for u := 0; u < max; u++ {
			if u != i && u%i == 0 {
				numbers[u] += i
			} else if u == i {
				numbers[u] += 1
			}
		}
	}

	var amicables []int
	for i := range numbers {
		if numbers[i] < len(numbers) && numbers[numbers[i]] == i && numbers[i] != i {
			amicables = append(amicables, i)
		}
	}

	return sum(amicables)
}

func main() {
	max := 10000
	fmt.Println(calcAmicableNumbersSum(max))
}
