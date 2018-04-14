package main

import (
	"strconv"
	"fmt"
)

func main() {
	var i, sum int64
	for i = 1; i < 1000000; i++ {
		tString := strconv.FormatInt(i, 10)
		if palindrome(tString) {
			bString := strconv.FormatInt(i, 2)
			if palindrome(bString) {
				sum += i
				fmt.Println(i, bString)
			}
		}
	}
	fmt.Println(sum)
}

func palindrome(s string) bool {
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
