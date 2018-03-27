package main

import "math"


// returns a slice prime numbers that equals or exceeds maxPrime
// useful if you have to factor a primesCovering(sqrt(x)) is useful for factoring x
func primesCovering(maxPrime int) [] int {
	knownPrimes := make([]int, 0, 10000)
	if maxPrime > 2 {
		candidate := 3
		knownPrimes = append(knownPrimes, candidate)

	FactorFound:
		for knownPrimes[len(knownPrimes) - 1] < maxPrime {
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


// returns a slice of howMany primes
func primeList(howMany int) [] int {
	knownPrimes := make([]int, 0, howMany)
	if howMany > 1 {
		candidate := 3
		knownPrimes = append(knownPrimes, candidate)

	FactorFound:
		for len(knownPrimes) < howMany - 1 {
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

//returns a generator function that produces successive primes on each call beginning with 5
func oddSieve(capacityEstimate int) func() int {
	knownPrimes := make([]int, 1, capacityEstimate)
	knownPrimes[0] = 3 // start with single known prime as last item found
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

func pfactor(x int) [] int {
	residue := x
	limit := int(math.Sqrt(float64(x)))
	candidates := primesCovering(limit)
	var factors []int

	for _, d := range candidates{
		for residue % d == 0 {
			residue = residue / d
			factors = append(factors, d)
		}
	}
	if residue > 1 {
		factors = append(factors, residue)
	}
	return factors
}

func nthPrime(index int) int {
	pList := primeList(index)
	return pList[len(pList)-1]
}
