package main

import (
	"math"
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

type factorMap struct {
	cover  []int
	counts []int
}

// set exponents in r to be factors of x
func (r *factorMap) factor(x int) {
	coverSize := len(r.cover)
	r.counts = make([]int, coverSize, coverSize)

	residue := x
	for i, d := range r.cover {
		for residue%d == 0 {
			residue = residue / d
			r.counts[i] += 1
		}
	}
}

func (r factorMap) value() int {
	result := 1
	for i, base := range r.cover {
		for j := r.counts[i]; j > 0; j-- {
			result *= base
		}
	}
	return result
}

// returns fm of lcm of input fms
// operates on the fact that the LCM when factored will consist of
// at least as many copies of each prime as the inputs
// so 18 is 2x3x3,   1 copy of 2, 2 of 3
// and 4 is 2x2      2 copies of 2, 0 of 3
// so LCM(4, 18) must be 2x2x3x3

func lcm(fMaps ... factorMap) factorMap {
	if len(fMaps) == 0 {
		return factorMap{[]int{2}, []int{0}}
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

	return factorMap{firstCover, counts}
}
