package main

import "fmt"

func fact(x int) int {
	res := 1
	for i := x; i > 1; i-- {
		res *= i
	}
	return res
}

func lexPermutation(items []int, pnum int) []int {
	itemcount := len(items)
	if itemcount == 1 {
		return items
	}

	subProblemSize := fact(itemcount - 1)
	subProblem := (pnum - 1) / subProblemSize
	subCut := subProblem * subProblemSize

	chosen := []int{items[subProblem]}
	remaining := append(items[:subProblem], items[subProblem+1:]...)

	return append(chosen, lexPermutation(remaining, pnum-subCut)...)
}

func main() {
	items := []int{0,1,2,3,4,5,6,7,8,9}

	fmt.Println(lexPermutation(items, 1000000))
}
