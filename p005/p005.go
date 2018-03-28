package main

import "fmt"

func main() {
	// setup test for anything up to the number 23
	cover := primesCovering(23)

	var fms []factorMap
	limit := 20

	for i:=1;i<=limit;i++ {
		r := factorMap{cover: cover}
		r.factor(i)
		fms = append(fms, r);
	}

	res := lcm(fms...)
	fmt.Printf("LCM of 1 through %d is %d", limit, res.value())
}
