/*

   Problem 30
   Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.

   Solution:
   The solution employed here is simply testing all possible numbers up to a maximum of
   9^5 * 5, and it runs so darn fast in Go that I'm not going to bother writing a better solution.

   It can be improved, however, by noting that there are only 10 possible digits, each with a 
   specific power of 5. So there are only a maximum of 10!/(10-5)! = 30240 sums to be found. 
   If we find these, we can just test each individually to see if it is equal to the sum of the 
   5th powers of its digits. This would be faster than testing all 9^5 * 5 numbers individually,
   like we are in this solution.

   Still, this runs in 30ms. Good enough, I'd say.
*/

package main

import (
	"fmt"
	"time"
)

var sum int64 = 0
var numbers []int64 = []int64{}
var pows []int64 = []int64{}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Println("-- Took", elapsed)
}

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

func testNumber(n int64) {
	var s int64 = 0
	var num int64 = n
	for {
		s += pows[n%10]
		n /= 10
		if n == 0 || s > num {
			break
		}
	}
	if s == num {
		sum += s
		numbers = append(numbers, s)
	}
}

func main() {
	defer timeTrack(time.Now())
	MAX := int(pow(9, 5)*5 + 1)

	// calculate powers
	for i := 0; i < 10; i++ {
		pows = append(pows, pow(int64(i), 5))
	}

	for i := 2; i < MAX; i++ {
		testNumber(int64(i))
	}

	fmt.Println(sum)
}
