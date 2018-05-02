package main

import (
	"github.com/kavehmz/prime"
	"fmt"
)

func main() {
	ps := prime.Primes(1000000)
	plen := len(ps)
	pmax := ps[len(ps)-1]

	pmap := make(map[uint64]bool)
	for _, p := range ps {
		pmap[p] = true
	}

	for slen := 7; slen < 1000; slen ++ {
		for i := 0; i < plen-slen; i++ {
			var sum uint64 = 0
			for j := 0; j < slen; j++ {
				sum += ps[i+j]
			}
			if pmap[sum] {
				fmt.Println("seq found [start, len, sum] ", ps[i], slen, sum)
				break
			}
			if sum > pmax {
				continue
			}
		}
	}

}
