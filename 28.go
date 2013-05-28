/*
   Problem 28

   Starting with the number 1 and moving to the right in a clockwise direction a 5 by 5 spiral is formed as follows:

   21 22 23 24 25
   20  7  8  9 10
   19  6  1  2 11
   18  5  4  3 12
   17 16 15 14 13

   It can be verified that the sum of the numbers on the diagonals is 101.

   What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral formed in the same way?

   Solution:
   ---------

   Looking at an extended grid, this is what we see:

    43 44 45 46 47 48 49 
    42 21 22 23 24 25 26
    41 20  7  8  9 10 27
    40 19  6  1  2 11 28
    39 18  5  4  3 12 29
    38 17 16 15 14 13 30
    37 36 35 34 33 32 31

     1
   + 3^2 + 5^2 + 7^2 + ...
   + 2^2+1 + 4^2+1 + 6^2+1 + ... = 
   + 3^2-2 + 5^2-4 + 7^2-6 + ... = 
   + 2^2-1 + 4^2-3 + 6^2-5 + ... = 

   Adding the first series with the second, and the third with the fourth, 
   we get:

   S1 = n(n+1)(2n+1)/6 - 1 + n/2
   S2 = n(n+1)(2n+1)/6 - (n-1)(n)/2 - 1

   So the answer is f(n) = 1 + S1(n) + S2(n) with n = 1001!

*/

package main

import "fmt"

func main() {
	var n int64 = 1001
	var S1 int64 = n*(n+1)*(2*n+1)/6 + n/2 - 1
	var S2 int64 = n*(n+1)*(2*n+1)/6 - (n-1)*n/2 - 1
	var ans int64 = 1 + S1 + S2
	fmt.Println(ans)
}
