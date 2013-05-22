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
	var result, _ = readFile("inputs/11.in")

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
	var biggest, num int = 0, 0
	for y := range nums {
		for x := range nums[y] {
			// horizontal
			if x+3 < len(nums[y]) {
				num = nums[y][x]
				for w := 1; w < 4; w++ {
					num *= nums[y][x+w]
				}
				if num > biggest {
					biggest = num
				}
			}

			// diagonal1
			if x+3 < len(nums[y]) && y+3 < len(nums) {
				num = nums[y][x]
				for w := 1; w < 4; w++ {
					num *= nums[y+w][x+w]
				}
				if num > biggest {
					biggest = num
				}
			}

			// diagonal2
			if x-3 >= 0 && y+3 < len(nums) {
				num = nums[y][x]
				for w := 1; w < 4; w++ {
					num *= nums[y+w][x-w]
				}
				if num > biggest {
					biggest = num
				}
			}

			// vertical
			if y+3 < len(nums) {
				num = nums[y][x]
				for w := 1; w < 4; w++ {
					num *= nums[y+w][x]
				}
				if num > biggest {
					biggest = num
				}
			}
		}
	}
	fmt.Println(biggest)
}
