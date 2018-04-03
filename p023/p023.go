package main

import (
    "github.com/greendavegreen/factor"
    "fmt"
)

func main() {
    ul := 28124
    ft := factor.NewFactorTable(ul)

    // build a slice of D(i) for integers < 28124 where D(i) is the sum of the proper divisors of i
    ds := make([]int, ul)
    for i := 1; i < ul; i++ {
        ds[i] = ft.Factor(i).D()
    }

    // build slice of just the abundant integers
    abundants := []int{}
    for j := 0; j < ul; j++ {
        if ds[j] > j {
            abundants = append(abundants, j)
        }
    }

    // for each integer, test if sum of abundant pair
    sumOfUnfound := 0
    for i := 1; i < ul; i++ {
        unfound := true
        // pairs made of abundant[j] abundant[k] where j<= k
        outer:
        for j:=0; abundants[j] < i; j++ {
            for k:=j; abundants[k] <= i - abundants[j]; k++ {
                if abundants[k] + abundants[j] == i {
                    unfound = false
                    break outer
                }
            }
        }
        if unfound {
            sumOfUnfound += i
        }
    }
    fmt.Printf("The sum of positive integers less than %d that are not abundant sums is %v\n", ul, sumOfUnfound)

}
