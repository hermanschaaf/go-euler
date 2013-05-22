package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
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
	var result, _ = readFile("inputs/8.in")

	// remove newlines
	result = strings.Replace(result, "\n", "", -1)
	var biggest, num int64 = 0, 0

	for i, _ := range result {
		if i > len(result)-5 {
			fmt.Println(biggest)
			return
		}
		num = 1
		for u := 0; u < 5; u++ {
			c, _ := strconv.ParseInt((string)(result[i+u]), 0, 0)
			num *= c
		}
		if num > biggest {
			biggest = num
		}
	}
}
