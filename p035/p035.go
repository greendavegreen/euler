package main

import (
	"strconv"
	"fmt"
	"strings"
	"github.com/greendavegreen/factor"
)

func main() {

	ul := 1000000
	ft := factor.NewFactorTable(ul)

	cCount := 1

	for i:= 3;i<ul;i +=2{
		isCircular,_ := ft.IsPrime(i)

		if isCircular {
			for j:=rotate(i); isCircular && j!=i; j=rotate(j) {
				isP, _ := ft.IsPrime(j)
				isCircular = isCircular && isP
			}
		}
		if isCircular {
			//fmt.Println(i)
			cCount += 1
		}
	}
	fmt.Println(cCount)
}


func rotate(i int) int {
	str := strings.Builder{}

	s := strconv.Itoa(i)
	str.WriteString(s[len(s)-1:])
	str.WriteString(s[:len(s)-1 ])

	v, _ := strconv.Atoi(str.String())
	return v
}