package main

import (
	"fmt"
	"math"
)

func main() {
	var sum int = 1000
	for a := 1; a < 333; a += 1 {
		for b := a; b < 1000-b-a; b++ {
			if sum*sum-2*sum*a-2*sum*b+2*a*b == 0 {
				var c int = (int)(math.Sqrt((float64)(a*a + b*b)))
				fmt.Println(a * b * c)
			}
		}
	}
}
