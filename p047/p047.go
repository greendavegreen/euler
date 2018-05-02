package main

import (
	"github.com/greendavegreen/factor"
	"fmt"
)

func main() {
	ulimit := 1000000

	ft := factor.NewFactorTable(ulimit)

	streak :=0
	for i:=2; i<ulimit; i++ {
		f, _ := ft.Factor(i)
		if f.DistinctPrimeCount() == 4 {
			streak++
			//fmt.Println(i)
		} else {
			streak = 0
		}
		if streak == 4 {
			fmt.Println("found 4 ending with ", i)
			fmt.Println(f.PrimeFactors())
		}
	}
}
