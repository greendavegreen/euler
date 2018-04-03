package main

import (
	"fmt"
)

func main() {
	limit := 1000000

	maxLen := 1
	maxStart := 1

	for i:=1; i<limit; i++ {
		if cl:=collatzLen(i); cl > maxLen {
			maxLen, maxStart = cl, i
		}
	}

	fmt.Printf("collatz of %d has length %d", maxStart, maxLen)
}