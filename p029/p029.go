package main

import (
	"math/big"
	"sort"
	"fmt"
)

func main () {
	var bns []*big.Int

	for a := 2; a <= 100; a++ {
		for b := 2; b <= 100; b++ {
			var limit big.Int
			limit.Exp(big.NewInt(int64(a)), big.NewInt(int64(b)), nil)
			bns = append(bns, &limit)
		}
	}

	sort.Slice(bns, func(i, j int) bool {
		return bns[i].Cmp(bns[j]) < 0
	})

	distinct := 0
	last := big.NewInt(0)

	for _, n := range bns {
		if n.Cmp(last) != 0 {
			distinct += 1
			last = n
		}
	}
	fmt.Println(distinct)
}

