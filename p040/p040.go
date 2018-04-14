package main

import (
	"github.com/greendavegreen/feed"
	"fmt"
)

func main() {
	i:= 1
	target := 1
	climit := 1000000
	product := 1
	for  r := range feed.AllInts() {
		if i == target {
			fmt.Println(i,r-'0')
			product *= int(r-'0')
			target *= 10
			if target > climit {
				break
			}
		}
		i++
	}
	fmt.Println("----------")
	fmt.Println(product)
}
