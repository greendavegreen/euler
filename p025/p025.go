package main

import (
	"math/big"
	"fmt"
)

func main() {
	// 10^999 is first integer with 1000 digits
	var limit big.Int
	limit.Exp(big.NewInt(10), big.NewInt(999), nil)

	i, f := 1, fib()
	for v := f(); v.Cmp(&limit) < 0; v = f() {
		i += 1
	}
	fmt.Println(i)
}
