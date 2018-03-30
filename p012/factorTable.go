package main

import (
	"math"
	"fmt"
)

// returns a slice prime numbers that equals or exceeds maxPrime
// useful if you have to factor a primesCovering(sqrt(x)) is useful for factoring x
func primesCovering(maxPrime int) [] int {
	knownPrimes := make([]int, 0, 10000)
	if maxPrime > 2 {
		candidate := 3
		knownPrimes = append(knownPrimes, candidate)

	FactorFound:
		for knownPrimes[len(knownPrimes)-1] < maxPrime {
			candidate += 2
			sqrtLimit := int(math.Sqrt(float64(candidate)) + 1)
			for _, p := range knownPrimes {
				if p > sqrtLimit {
					break // any non-prime will have factor <= sqrt(candidate), therefore this one is prime
				}
				if candidate%p == 0 {
					continue FactorFound // next candidate
				}
			}
			knownPrimes = append(knownPrimes, candidate)
		}
	}

	return append([] int{2}, knownPrimes...)
}


func NewFactorTable(upperLimit int) *factorTable {
	tableBounds := int(math.Sqrt(float64(upperLimit)))
	cover := primesCovering(tableBounds)
	maxPrime := cover[len(cover)-1]
	return &factorTable{primesCovering(tableBounds), maxPrime * maxPrime}
}

type factorTable struct {
	cover       []int
	maxPromised int
}

type factored struct {
	v   int
	cover   []int
	counts  []int
	remains int
}

// set exponents in r to be factors of x
func (t factorTable) factor(x int) *factored {
	if x > t.maxPromised {
		fmt.Errorf("input value too large for this factorTable %v", x)
	}

	exps := make([]int, len(t.cover))

	residue := x
	for i, d := range t.cover {
		for residue%d == 0 {
			residue = residue / d
			exps[i] += 1
		}
		if residue == 1 {
			break
		}
	}
	return &factored{x, t.cover, exps, residue}
}



func (f factored) primeFactors() []int {
	var res []int
	for i, base := range f.cover {
		for j := 0; j < f.counts[i]; j++ {
			res = append(res, base)
		}
	}
	if f.remains > 1 {
		res = append(res, f.remains)
	}
	return res
}

func (f factored) value() int {
	return f.v
}

func (f factored) divisorCount() int {
	result := 1
	for _, v := range f.counts {
		result *= (v + 1)
	}
	if f.remains > 1 {
		result *= 2
	}
	return result
}



func (t factorTable) lcm(fMaps ... factored) *factored {
	if len(fMaps) == 0 {
		// error
	}

	// baseline is zero of everything
	firstCover := fMaps[0].cover
	counts := make([]int, len(firstCover))

	// if any factormap uses more copies of a factor, up our count to that
	for _, fm := range fMaps {
		for i, d := range fm.counts {
			if counts[i] < d {
				counts[i] = d
			}
		}
	}

	return &factored{0, firstCover, counts, 1}
}
