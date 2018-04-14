package main

import "fmt"

func intSolutionCount(perimeter int, verbose bool) int {
	result:= 0
	for hyp :=1; hyp <= perimeter / 2; hyp++ {
		for a:=1; a<= (perimeter - hyp) / 2; a++ {
			b := perimeter - hyp - a
			if a*a + b*b == hyp*hyp {
				if verbose {
					fmt.Println(a,b,hyp)
				}
				result += 1
			}
		}
	}
	return result
}

func main() {

	maxSol := 0

	for p:=3;p<=1000;p++ {
		if sc:=intSolutionCount(p, false); sc > 0 {
			if maxSol < sc {
				maxSol = sc
				//intSolutionCount(p,true)
				fmt.Println(p, sc)
				fmt.Println("--------")
			}
		}
	}


}
