package main

import "fmt"


var lookup[][]int

func search(m int, n int) int {
	if val := lookup[m][n]; val > 0 {
		return val
	}
	computed := search(m-1, n) + search(m, n-1)
	lookup[m][n] = computed
	return computed
}


func main() {
	m := 20
	n := 20
	lookup = make([][]int, m+1)

	for i:=0; i< m+1; i++ {
		lookup[i] = make([]int, n+1)
		lookup[i][0] = 1
	}
	for i:=0; i< n+1; i++ {
		lookup[0][i] = 1
	}


	count := search(20,20)

	fmt.Printf("paths in %dx%d grid = %d", m, n, count)
}
