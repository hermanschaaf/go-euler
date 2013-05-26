/*

   Problem 25:
   What is the first term in the Fibonacci sequence to contain 1000 digits?

   Math based on [this wikipedia paragraph](http://en.wikipedia.org/wiki/Fibonacci_number#Computation_by_rounding)

*/

package main

import (
	"fmt"
	"math"
)

func main() {
	phi := (1 + math.Sqrt(5.0)) / 2
	ans := int(math.Floor((999.0+0.5*math.Log10(5.0))/math.Log10(phi) + 0.5))
	fmt.Println(ans)
}
