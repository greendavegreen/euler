package main

import (
	"github.com/kavehmz/prime"
	"fmt"
)

func panDigital(x uint64) bool {
	i := 0
	unUsed := []bool{false, true, true, true, true, true, true, true, true, true}
	for product := x; product > 0; product = product / 10 {
		res := product % 10
		unUsed[res] = false
		i++
	}
	for i > 0 {
		if unUsed[i] {
			return false
		}
		i--
	}
	return true
}


func main() {
	x := prime.Primes(7654321)
	fmt.Println("got primes")
	for i:= len(x) - 1; i>0;i-- {
		if panDigital(x[i]) {
			fmt.Println(x[i])
			break
		}
	}
}
