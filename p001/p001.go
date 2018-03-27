package main

import "fmt"

func main() {
	divisors := []int{3,5}
	limit := 1000
	result := multiples(divisors, limit)
	fmt.Printf("Sum of multiples of any %v less than %v = %v\n", divisors, limit, result)
}
