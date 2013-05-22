package main

import (
	"fmt"
)

var words = [...]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var tens = [...]string{"zero", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty",
	"ninety"}
var hundred = "hundred"
var thousand = "thousand"
var teens = [...]string{"zero", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen",
	"seventeen", "eighteen", "nineteen", "twenty"}

func below100(i int) string {
	i = i % 100
	var word string
	if i < 10 {
		word = words[i]
	} else if i%10 == 0 {
		word = tens[i/10]
	} else if i < 20 {
		word = teens[i-10]
	} else if i < 100 {
		word = tens[i/10] + words[i%10]
	}
	return word
}

func main() {
	totalLength := 0
	for i := 1; i <= 1000; i++ {
		var word string
		if i < 100 {
			word = below100(i)
		} else if i < 1000 {
			if i%100 == 0 {
				word = words[i/100] + hundred
			} else {
				word = words[i/100] + hundred + "and" + below100(i)
			}
		} else {
			word = "one" + thousand
		}
		totalLength += len(word)
	}
	fmt.Println(totalLength)
}
