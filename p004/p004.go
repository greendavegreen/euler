package main

import "fmt"

func main() {
	// largest palindrome product of 2 3-digit numbers
	bestProd, besti, bestj := -1, -1, -1

	for i := 999; i > 99; i-- {
		for j := i; j > 99; j-- {
			product := i * j
			if product <= bestProd {
				break
			}
			if palindrome(product) {
				bestProd, besti, bestj = product, i, j
			}
		}
	}
	fmt.Printf("best product: %d - %d x %d", bestProd, besti, bestj)
}
