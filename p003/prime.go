package main

import "math"


//returns a generator function that produces successive primes on each call beginning with 5
func oddSieve(capacityEstimate int) func() int {
	knownPrimes := make([]int, 1, capacityEstimate)
	knownPrimes[0] = 3  // start with single known prime as last item found
	candidate := 3

	return func() int {
	FactorFound:
		candidate += 2
		sqrtLimit := int(math.Sqrt(float64(candidate)) + 1)
		for _, p := range knownPrimes {
			if p > sqrtLimit {
				break // any non-prime will have factor <= sqrt(candidate), therefore this one is prime
			}
			if candidate%p == 0 {
				goto FactorFound // next candidate
			}
		}
		knownPrimes = append(knownPrimes, candidate)
		return candidate
	}
}


func nthPrime(index int) int {
	if index == 1 {
		return 2
	}
	if index == 2 {
		return 3
	}

	f := oddSieve(index)
	for i := 1; i < index-2; i++ {
		f()
	}

	return f()
}
