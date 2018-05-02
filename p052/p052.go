package main

import "fmt"

func main() {

	var i uint64 = 1

	outer:
	for ;;i++ {
		j := i
		base, err := digiSlice(i)
		if err != nil {
			continue
		}

		for c := 2; c<7; c++{
			j += i
			b2, err := digiSlice(j)
			if err != nil || !digiEqual(base, b2){
				continue outer
			}
		}
		fmt.Println("success ", i)
	}
}
