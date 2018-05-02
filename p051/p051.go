package main

import (
	"fmt"
	"github.com/kavehmz/prime"
	"strconv"
	"github.com/etnz/permute"
)

func main() {
	ps, pmap := newPmap()

	var maxP uint64 = 0
	maxVar := 7

	// iterate 6 digit primes
	for _, p := range ps {
		if p > 100000 && p < 1000000 {
			pstring := strconv.FormatUint(p, 10)
			digits := []rune(pstring)

			// iterate number of digits to replace
			for k := 1; k< len(digits); k++ {

				positions := permute.New(k)

				maxVar, maxP = testperm(digits, positions, pmap, maxVar, maxP, p)
				for permute.SubsetLexNext(positions, len(digits)) {
					maxVar, maxP = testperm(digits, positions, pmap, maxVar, maxP, p)
				}
			}

		}
	}
}

func testperm(digits []rune, positions []int, pmap map[string]bool, maxVar int, maxP uint64, p uint64) (int, uint64) {
	temp := make([]rune, len(digits))
	copy(temp, digits)
	pvariations := primePerms(temp, digits, positions, pmap)
	if pvariations > maxVar {
		maxVar = pvariations
		maxP = p
		fmt.Println(p, pvariations, positions)
	}
	return maxVar, maxP
}

func newPmap() ([]uint64, map[string]bool) {
	ps := prime.Primes(1000000)
	pmap := make(map[string]bool)
	for _, p := range ps {
		if p > 100000 && p < 1000000 {
			pmap[strconv.FormatUint(p, 10)] = true
		}
	}
	return ps, pmap
}

func primePerms(temp []rune, digits []rune, positions []int, pmap map[string]bool) int {
	pvariations := 0
	for i := 0; i < 10; i++ {
		r := '0' + i
		for _, j := range positions {
			temp[j] = rune(r)
		}
		transmuted := string(temp)
		if pmap[transmuted] {
			pvariations++
		}
	}
	return pvariations
}
