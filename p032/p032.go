package main

import (
	"fmt"
	"time"
	"log"
)

func main() {
	start := time.Now()

	detect, score := detector()
	visit([] int{1, 2, 3, 4, 5, 6, 7, 8, 9}, detect)
	fmt.Println(score())

	elapsed := time.Since(start)

	log.Printf("Binomial took %s", elapsed)

}

func detector() (func([] int), func() int) {
	m := make(map[int]bool)

	// detect form     abcd = ef * ghi
	test := func(a []int) {
		prod := a[0]*1000 + a[1]*100 + a[2]*10 + a[3]

		x := a[4]*10 + a[5]
		y := a[6]*100 + a[7]*10 + a[8]

		x1 := a[4]
		y1 := a[5]*1000 + a[6]*100 + a[7]*10 + a[8]

		if x*y == prod || x1*y1 == prod {
			m[prod] = true
		}
	}

	score := func() int {
		grandSum := 0
		for z := range m {
			grandSum += z
		}
		return grandSum
	}

	return test, score
}

// heap's algorithm for permutations with function to invoke on each distinct
func visit(inputs []int, f func([]int)) {
	size := len(inputs)
	c := make([]int, size)
	i := 1
	var j int

	f(inputs)
	for i < size {
		if c[i] < i {
			if i%2 == 0 {
				j = 0
			} else {
				j = c[i]
			}
			inputs[i], inputs[j] = inputs[j], inputs[i]
			f(inputs)
			c[i] += 1
			i = 1
		} else {
			c[i] = 0
			i += 1
		}
	}
}
