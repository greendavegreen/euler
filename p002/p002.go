package main

import "fmt"

func main() {
	limit := 4000000
	result := evenFibs(limit)
	fmt.Printf("Sum even fibs less than %v = %v\n", limit, result)
}
