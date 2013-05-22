/*

   Problem: 15
   How many paths are there through a square lattice of size 20x20,
   starting at the upper left and only moving right or down?

*/

package main

import (
	"fmt"
	"math/big"
)

func main() {
	N := 20

	// drawing it out, I found that it
	// comes down to pascal's triangle,
	// which simplifies to N choose R:
	// N! / (R! * (N - R)!)
	// where R = N / 2
	// so:
	// N! / ((N/2)! * (N/2)!)

	// So the solution is easy:
	a := new(big.Int)

	fmt.Println(a.Binomial((int64)(N*2), (int64)(N)))
}
