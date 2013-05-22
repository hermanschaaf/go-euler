package main

import (
	"fmt"
	"io"
	"math/big"
	"os"
	"strings"
)

func readFile(fileName string) (string, error) {
	// read input file
	fi, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	var result []byte
	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := fi.Read(buf)
		result = append(result, buf[0:n]...)
		if err != nil && err != io.EOF {
			return "", err // fi will be closed if we return here.
		}
		if n == 0 {
			break
		}
	}
	return string(result), nil
}

func main() {
	var result, _ = readFile("inputs/13.in")

	var lines []string = strings.Split(result, "\n")
	var nums []*big.Int
	total := new(big.Int)
	for _, line := range lines {
		n := new(big.Int)
		n.SetString((string)(line), 10)
		nums = append(nums, n)
		total.Add(n, total)
	}
	fmt.Println(total.String()[:10])
}
