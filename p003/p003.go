package main

import (
	"fmt"
	"time"
)

func main() {
//	limit := 1000000

	start := time.Now()

	item := 600851475143
	factors := pfactor(item)
	fmt.Printf("Factors of %d, %v \n\n", item, factors)
//	result := nthPrime(limit)

	elapsed := time.Since(start)

//	fmt.Printf("The %vth prime = %v\n", limit, result)
	fmt.Printf("Time: %s", elapsed)
}
