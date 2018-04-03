package main

import "fmt"
import "github.com/greendavegreen/factor"
//Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
//If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.
//
//For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.
//
//Evaluate the sum of all the amicable numbers under 10000.

func main() {
	ul := 10000
	ft := factor.NewFactorTable(ul)

	sumOfAmicables := 0
	for i := 2; i <= ul; i++ {
		if d1 := ft.Factor(i).D(); d1 > i {
			if d2 := ft.Factor(d1).D(); d2 == i {
				fmt.Printf("Proper Pair:\n    d(%d)=%d\n    d(%d)=%d\n\n", i, d1, d1, d2)
				sumOfAmicables += i + d1
			}
		}
	}

	fmt.Printf("The sum of amicable numbers below %d is %d\n", ul, sumOfAmicables)
}
