package main

import "fmt"

func main() {
	//f := func(a []int) {
	//	fmt.Println(a)
	//}
	f := func(a []int) {
		prod := a[0] * 10 + a[1]
		if a[2] * a[3] == prod && a[2] < a[3] {
			fmt.Println(a)
		}
	}

	pperm([] int{1, 2, 3, 4}, f)
}

func pperm(inputs []int, f func([]int)) {
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

//
//for perm := range GeneratePermutations([]int{1,2,3}){
//fmt.Println(perm)
//}
func GeneratePermutations(data []int) <-chan []int {
	c := make(chan []int)
	go func(c chan []int) {
		defer close(c)
		permutate(c, data)
	}(c)
	return c
}

func permutate(c chan []int, inputs []int) {
	output := make([]int, len(inputs))
	copy(output, inputs)
	c <- output

	size := len(inputs)
	p := make([]int, size+1)
	for i := range p {
		p[i] = i
	}
	for i := 1; i < size; {
		p[i]--
		j := 0
		if i%2 == 1 {
			j = p[i]
		}
		inputs[i], inputs[j] = inputs[j], inputs[i]

		output := make([]int, len(inputs))
		copy(output, inputs)
		c <- output

		for i = 1; p[i] == 0; i++ {
			p[i] = i
		}
	}
}
