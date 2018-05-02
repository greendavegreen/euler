package main

import (
	"math/big"
	"fmt"
)

func main() {
	var p big.Int

	sum := big.NewInt(0)
	var i int64

	for i = 1; i <= 1000; i++ {
		p.Exp(big.NewInt(i), big.NewInt(i), nil)
		sum = sum.Add(sum, &p)
	}
	s := sum.String()
	fmt.Println(s[len(s)-10:])
}
