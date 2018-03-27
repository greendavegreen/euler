package main

// sums all items that are evenly divisible any divisor and less than limit

func multiples(divisors []int, limit int) int {
	sum := 0
	for i := 1; i < limit; i++ {
		for _, d := range divisors {
			if i%d == 0 {
				sum += i
				break
			}
		}
	}
	return sum
}
