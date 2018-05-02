package main

import (
	"math/big"
	"fmt"
)

func fact(x int64) *big.Int {
	bi := big.NewInt(1)
	for x > 0 {
		bi.Mul(big.NewInt(x), bi)
		x = x - 1
	}
	return bi
}

func choose(n int64, r int64) *big.Int {
	var f big.Int
	var p big.Int
	p.Mul(fact(r), fact(n-r))
	f.Div(fact(n), &p)

	return &f
}

func main() {
	count := 0
	bi := big.NewInt(1000000)
	var n,r int64
	for n = 1; n <= 100; n++ {
		for r=1;r<=n;r++ {
			if bi.Cmp(choose(n,r)) < 0 {
				count++
			}
		}
	}

	fmt.Println(count, "greater than 1 million")
}
