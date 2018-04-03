package main

import "math/big"

func fib() func() *big.Int {
	a := big.NewInt(0)
	b := big.NewInt(1)

	return func() *big.Int {
		a.Add(a, b)
		a, b = b, a
		return a
	}
}

