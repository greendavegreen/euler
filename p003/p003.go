package main

import (
	"fmt"
	"time"
)

func main() {
	f := oddSieve()
	fmt.Printf("#%v %v\n", 1, 2)
	limit := 1000000

	start := time.Now()

	for i := 2; i < limit; i++ {
		f()
	}

	fmt.Printf("#%v %v\n", limit, f())

	elapsed := time.Since(start)
	fmt.Printf("limit primes took %s", elapsed)
}
