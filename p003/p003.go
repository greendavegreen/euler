package main

import (
	"fmt"
	"time"
)

func main() {
	limit := 5000

	start := time.Now()

	result := nthPrime(limit)

	elapsed := time.Since(start)

	fmt.Printf("The %vth prime = %v\n", limit, result)
	fmt.Printf("Time: %s", elapsed)
}
