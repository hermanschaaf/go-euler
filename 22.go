//Problem 22
// What is the total of all the name scores in the file
// Lifted from https://github.com/Narsil/go-euler/blob/master/problems/problem22.go
// out of laziness and it being a lame problem

package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func value(s string) int {
	sum := 0
	for _, char := range s {
		sum += (int)(char) - 64
	}
	return sum
}

func main() {
	sum := 0
	file, error := ioutil.ReadFile("inputs/names.txt")
	if error != nil {
		fmt.Println("Cannot open names.txt, make sure you have it in the current directory")
	}
	names := strings.SplitN(string(file), ",", -1)
	sort.Strings(names)
	for index, name := range names {
		sum += value(strings.Trim(name, "\"")) * (index + 1)
	}
	fmt.Println(sum)
}
