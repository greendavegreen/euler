package main

import (
	"fmt"
)

func main() {
	limit := 10001
	result := nthPrime(limit)
	fmt.Printf("The %vth prime = %v\n", limit, result)
}
