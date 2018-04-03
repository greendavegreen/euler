package main

func collatz(start int) func() int {
	n := start
	return func() (res int) {
		res = n
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		return
	}
}

func collatzLen(start int) (length int) {
	length = 1
	for c := collatz(start); c() != 1; length++ {
	}
	return
}
