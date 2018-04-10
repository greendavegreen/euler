package main

import "fmt"

func makeChange(target int, denominations []int) int {
	if len(denominations) == 1 {
		// with one type of coin, you either an make even change, or not
		if target%denominations[0] == 0 {
			return 1
		} else {
			return 0
		}
	} else {
		// result is sum sub-problems using zero of first coin, 1 of first coin, etc
		total := 0
		for target >= 0 {
			total += makeChange(target, denominations[1:])
			target -= denominations[0]
		}
		return total
	}
}

func main() {
	den := []int{200, 100, 50, 20, 10, 5, 2, 1}
	count := makeChange(200, den)
	fmt.Println(count)
}
