package main

import (
	"fmt"
)

func main() {
	for a := 1; a < 1000; a++ {
		asq := a*a

		for b := a; b < 1000 - a; b++ {
			bsq:=b*b

			c := 1000 - a - b
			csq := c*c

			if asq + bsq == csq {
				fmt.Printf("a:%v  b:%v  c:%v\n", a,b,c)
				fmt.Printf("abc product %v\n", a*b*c)
			}
		}
	}
}
