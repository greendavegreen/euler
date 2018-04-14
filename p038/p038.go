package main

import (
	"fmt"
	"strings"
	"strconv"
	"time"
)

func panMult(x int) int {
	i := 0
	used := []bool{true, false, false, false, false, false, false, false, false, false}
	dups := false
	var multiplier int
	for multiplier = 1; i < 9; multiplier++ {
		for product := x * multiplier; product > 0; product = product / 10 {
			res := product % 10
			dups = dups || used[res]
			used[res] = true
			i++
		}
	}
	if i == 9 && !dups {
		return multiplier - 1
	} else {
		return 0
	}
}

func genMult(x int, m int) int {
	var str strings.Builder
	for i:=1; i<=m;i++ {
		str.WriteString(strconv.Itoa(x * i))
	}
	v,_ := strconv.Atoi(str.String())
	return v
}

func main() {
	start := time.Now()
	max:=0
	for i := 1; i < 10000; i++ {
		if r:= panMult(i); r > 0 {
			z := genMult(i,r)
			if z > max {
				max = z
			}
		}
	}
	fmt.Println(max)
	elapsed := time.Since(start)
	fmt.Printf("Binomial took %s", elapsed)
}
