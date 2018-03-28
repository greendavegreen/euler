package main

import "math"

// returns a slice of howMany primes
func primeList(howMany int) [] int {
	candidate := 1
	knownPrimes := make([]int, 0, howMany)

FactorFound:
	for len(knownPrimes) < howMany-1 {
		candidate += 2
		sqrtLimit := int(math.Sqrt(float64(candidate)) + 1)
		for _, p := range knownPrimes {
			if p > sqrtLimit {
				break // any non-prime will have factor <= sqrt(candidate)
			}
			if candidate%p == 0 {
				continue FactorFound // next candidate
			}
		}
		knownPrimes = append(knownPrimes, candidate)
	}

	return append([] int{2}, knownPrimes...)
}

func nthPrime(n int) int {
	pList := primeList(n)
	return pList[len(pList)-1]
}
