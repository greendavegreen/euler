package main

import "fmt"

func main() {
	limit := 100

	sumOfSquares := 0
	sumOfNaturals := 0
	for i := 1; i <= limit; i++ {
		sumOfSquares += i * i
		sumOfNaturals += i
	}

	squareOfSums := sumOfNaturals * sumOfNaturals

	fmt.Printf("Sum of squares of first %v natural numbers = %v\n", limit, sumOfSquares)
	fmt.Printf("Square of sum of first %v natural numbers = %v\n", limit, squareOfSums)
	fmt.Printf("Diff: %v", squareOfSums-sumOfSquares)
}
