package main

import "fmt"

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

func detector() (func([] int), func() int) {
	count := 0

	test := func(a []int) {

		if p3 := a[7]*100 + a[8]*10 + a[9]; p3%17 == 0 {
			if (a[6]*100+a[7]*10+a[8])%13 == 0 {
				if (a[5]*100+a[6]*10+a[7])%11 == 0 {
					if p2 := a[4]*100 + a[5]*10 + a[6]; p2%7 == 0 {
						if (a[3]*100+a[4]*10+a[5])%5 == 0 {
							if (a[2]*100+a[3]*10+a[4])%3 == 0 {
								if p1 := a[1]*100 + a[2]*10 + a[3]; p1%2 == 0 {
									count +=  p3 + p2 * 1000 + p1 * 1000000 + a[0] * 1000000000
								}
							}
						}
					}
				}
			}
		}
	}

	score := func() int {
		return count
	}

	return test, score
}

func main() {

	detect, score := detector()
	visit([] int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, detect)

	fmt.Println(score())
}
