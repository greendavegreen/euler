package main

import (
	"fmt"
	"github.com/etnz/permute"
)

func maina() {

	x := []int{0, 1, 2, 3, 4}
	positions := permute.New(1)
	fmt.Println(positions)
	for permute.SubsetLexNext(positions, len(x)) {
		fmt.Println(positions)
	}

}
