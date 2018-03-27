package main

import "fmt"

func palindrome(x int) bool {
	r := []rune(fmt.Sprintf("%d", x))

	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		if r[i] != r[j] { return false}
	}

	return true
}
