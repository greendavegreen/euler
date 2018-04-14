package main

import (
	"strconv"
	"fmt"
	"github.com/greendavegreen/factor"
)

func main() {
	ulimit := 1000000
	ft := factor.NewFactorTable(ulimit)
	grandSum := 0

	for i := 23; i < ulimit; i += 2 {
		shiftable, _ := ft.IsPrime(i)
		if shiftable {
			for j := i / 10; shiftable && j > 0; j = j / 10 {
				isP, _ := ft.IsPrime(j)
				shiftable = shiftable && isP
			}
		}
		if shiftable {
			for j := lshift(i); shiftable && j > 0; j = lshift(j) {
				isP, _ := ft.IsPrime(j)
				shiftable = shiftable && isP
			}
		}
		if shiftable {
			grandSum += i
			fmt.Println(i)
		}
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()

	fmt.Println(grandSum)
}

func lshift(i int) int {
	if i < 10 {
		return 0
	}
	s := strconv.Itoa(i)
	chopped := s[1:]
	v, _ := strconv.Atoi(chopped)
	return v
}
