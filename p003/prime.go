package main
import "math"

func oddSieve() func() int {
	primes := make([]int, 0, 50000)
	candidate := 3

	return func() int {
	FactorFound:
		sqrtLimit := int(math.Sqrt(float64(candidate)) + 1)
		for _, p := range primes {
			// a candidate will have a factor <= to its sqrt or be prime
			if p > sqrtLimit {
				break
			}

			// not prime, try next
			if candidate%p == 0 {
				candidate += 2
				goto FactorFound
			}
		}
		primes = append(primes, candidate)
		result := candidate
		candidate += 2
		return result
	}
}
