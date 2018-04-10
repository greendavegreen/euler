package main

import "fmt"

// generate 3 7 13 21 ...
func d1_gen() func() int {
	a, b := 1, 2
	return func() int {
		a, b = a+b, b+2
		return a
	}
}

// generate 1 5 9 17 25 ...
func d2_gen() func() int {
	z, a, b, c, d := 1, 0, 4, 0, 4
	return func() int {
		z, a, b, c, d = z+a, b, b+c, d, c
		return z
	}
}

func main() {
	f := d1_gen()
	g := d2_gen()

	sum := g()
	for i := 0; i < 500; i++ {
		sum += f()
		sum += f()
		sum += g()
		sum += g()
	}

	fmt.Printf("sum %v", sum)
}
