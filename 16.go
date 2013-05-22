package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	ans := new(big.Int)
	s := "1"
	for i := 0; i < 1000; i++ {
		s += "0"
	}
	ans.SetString(s, 2)
	var n int64 = 0
	for _, c := range ans.String() {
		a, _ := strconv.ParseInt((string)(c), 10, 0)
		n += a
	}
	fmt.Println(n)
}
