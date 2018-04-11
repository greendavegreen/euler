package main

import "fmt"

func main() {
	grandTotal:=0
	nineFactorial := 362880

	for i :=3;i<7*nineFactorial;i++ {
		if i == fSum(i) {
			fmt.Println(i)
			grandTotal+=i
		}
	}
	fmt.Println(grandTotal)
}

	func fSum(a int) int {
		f:= [] int {
			1,
			1,
			2,
			6,
			24,
			120,
			720,
			5040,
			40320,
			362880,
		}
		var rfunc func(int)int
		rfunc = func (a int) int {
			if a > 0 {
				return f[a%10] + rfunc(a/10)
			}
			return 0
		}

		return rfunc(a)
	}