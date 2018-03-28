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
			r.counts[i]+=1
		}
	}
}

// raise exponents of r if they are lower than t
func (r factorMap) maxExponents(t factorMap) {
	for i, d := range t.counts {
		if r.counts[i] < d {
			r.counts[i] = d
		}
	}
}

// raise exponents of r if they are lower than t
func (r factorMap) value() int {
	result := 1
	for i, base := range r.cover {
		for j:=r.counts[i]; j>0; j-- {
			result *= base
		}
	}
	return result
}

// returns fm of lcm of input fms
func lcm(fMaps... factorMap) factorMap {
	if len(fMaps) == 0 {
		// return smallest empty factorMap
		cover := primesCovering(2)
		result := factorMap{cover, make([]int, 1, 1)}
		return result
	} else {
		//init empty map with similar cover
		first := fMaps[0]
		coverSize := len(first.cover)
		result := factorMap{first.cover, make([]int, coverSize, coverSize)}

		for _, fm := range fMaps {
			result.maxExponents(fm)
		}
		return result
	}

}