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

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	var result, _ = readFile("inputs/67.in")

	var lines []string = strings.Split(result, "\n")
	var nums [][]int
	for _, line := range lines {
		var row []int
		for _, num := range strings.Split(line, " ") {
			i, _ := strconv.ParseInt((string)(num), 10, 0)
			row = append(row, (int)(i))
		}
		nums = append(nums, row)
	}

	var maxis [][]int
	l := len(nums)
	for y := l - 1; y >= 0; y-- {
		var row []int
		for x := range nums[y] {
			if y == l-1 {
				row = append(row, nums[y][x])
			} else {
				if len(maxis[l-y-2]) > x+1 {
					row = append(row, nums[y][x]+max(maxis[l-y-2][x], maxis[l-y-2][x+1]))
				} else {
					row = append(row, nums[y][x]+maxis[l-y-2][x])
				}
			}
		}
		maxis = append(maxis, row)
	}
	fmt.Println(maxis[len(maxis)-1][0])
}
