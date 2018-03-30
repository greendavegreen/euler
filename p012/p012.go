package main

import (
	"fmt"
)

func main() {
	ulimit := 10000000000

	ft := NewFactorTable(ulimit)
	max := 0

	fmt.Printf("Cover calculated: size=%v \n", len(ft.cover))

	gen := triangles()
	for v := gen(); v < ulimit; v = gen() {
		f := ft.factor(v)
		if dc := f.divisorCount(); dc > max {
			fmt.Printf("%v -> %v\n", v, dc)
			max = dc
		}
	}
}
