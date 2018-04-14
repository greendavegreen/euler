package main

import "fmt"

func pNumberGenerator() func() int {
	inc := 1
	i := 0
	return func() int {
		i, inc = i+inc, inc+3
		return i
	}
}

func main() {
	p := pNumberGenerator()
	pMap := make(map[int]bool)

	knownPs := []int{p(), p()}
	pMap[knownPs[0]] = true
	pMap[knownPs[1]] = true


	for s := p(); s < 10000000; s = p() {
		for _, d := range knownPs {
			pj := (s - d) / 2
			pk := (s + d) / 2

			if pMap[pj] && pMap[pk] {
				fmt.Println("super", d, pj, pk, s)
				break
			}
		}
		knownPs = append(knownPs, s)
		pMap[s] = true
	}
}
