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

type factorTable struct {
	cover       []int
	maxPromised int
}

func NewFactorTable(upperLimit int) *factorTable {
	tableBounds := int(math.Sqrt(float64(upperLimit)))
	cover := primesCovering(tableBounds)
	maxPrime := cover[len(cover)-1]
	return &factorTable{primesCovering(tableBounds), maxPrime * maxPrime}
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

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (ft factorTable) lcm(fMaps ... factored) *factored {
	if len(fMaps) == 0 {
		return &factored{0, []int{}, []int{}, 1}
	}
	// baseline is zero of everything
	counts := make([]int, len(ft.cover))

	remainsFoundList := make([]int, 0)
	remainsBuilder := 1

	// if any factormap uses more copies of a factor, up our count to that
	for _, fm := range fMaps {
		if len(ft.cover) != len(fm.cover) {
			fmt.Errorf("item passed to factortable.lcm has illegal table")
		}
		for i, d := range fm.counts {
			if counts[i] < d {
				counts[i] = d
			}
		}
		if !contains(remainsFoundList, fm.remains) {
			remainsFoundList = append(remainsFoundList, fm.remains)
			remainsBuilder = remainsBuilder * fm.remains
		}
	}

	rawValue := 1
	for i, base := range ft.cover {
		for j := counts[i]; j > 0; j-- {
			rawValue *= base
		}
	}
	rawValue *= remainsBuilder

	return &factored{rawValue, ft.cover, counts, remainsBuilder}
}

type factored struct {
	v   int
	cover   []int
	counts  []int
	remains int
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




