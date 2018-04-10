package main

import (
	"fmt"
)

func main() {
 	m := make(map[int]bool)
	visit([] int{1, 2, 3, 4, 5, 6, 7, 8, 9}, test1x4, m)
	visit([] int{1, 2, 3, 4, 5, 6, 7, 8, 9}, test2x3, m)

	grandSum := 0
	for z := range m {
		grandSum += z
	}
	fmt.Println(grandSum)
}

// detect form     abcd = e * fghi
func test1x4(a []int, m map[int]bool) {
	x := a[4]
	y := a[5] * 1000 + a[6] * 100 + a[7] * 10 + a[8]
	prod := a[0]*1000 + a[1] * 100 + a[2] * 10 + a[3]

	if x * y == prod {
		m[prod] = true
	}
}

// detect form     abcd = ef * ghi
func test2x3(a []int, m map[int]bool) {
	x := a[4] * 10 + a[5]
	y := a[6] * 100 + a[7] * 10 + a[8]
	prod := a[0]*1000 + a[1] * 100 + a[2] * 10 + a[3]

	if x * y == prod {
		m[prod] = true
	}
}

// heap's algorithm for permutations with function to invoke on each distinct
func visit(inputs []int, f func([]int, map[int]bool), m map[int]bool) {
	size := len(inputs)
	c := make([]int, size)
	i := 1
	var j int

	f(inputs, m)
	for i < size {
		if c[i] < i {
			if i%2 == 0 {
				j = 0
			} else {
				j = c[i]
			}
			inputs[i], inputs[j] = inputs[j], inputs[i]
			f(inputs, m)
			c[i] += 1
			i = 1
		} else {
			c[i] = 0
			i += 1
		}
	}
}