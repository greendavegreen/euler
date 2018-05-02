package main

import (
	"fmt"
	"math"
)

func pNumberGenerator() func() uint64 {
	var inc uint64 = 1
	var i uint64 = 0
	return func() uint64 {
		i, inc = i+inc, inc+3
		return i
	}
}

// maintain list of known pentagons and map of pentagon->bool for quick test

// generate next largest pentagon and assume it is S = Pj + Pk
// assume each pentagon smaller could be D
//   given S,D, test of whether Pj, and Pk exist is simple

func main() {
	nextP := pNumberGenerator()
	isPent := make(map[uint64]bool)

	knownPentagons := []uint64{nextP(), nextP()}
	isPent[knownPentagons[0]] = true
	isPent[knownPentagons[1]] = true

	fmt.Println(math.MaxUint32)
	var bestD uint64 = math.MaxUint32

	for s := nextP(); s < math.MaxUint32; s = nextP() {
		for _, d := range knownPentagons {
			if d == bestD {
				break
			}
			pj := (s - d) / 2
			pk := (s + d) / 2
			if isPent[pj] && isPent[pk] {
				fmt.Println("solution", d, pj, pk, s)
				bestD = d
				break
			}
		}
		knownPentagons = append(knownPentagons, s)
		isPent[s] = true
	}


	lastP := knownPentagons[len(knownPentagons)-1]
	cut := lastP /2
	fmt.Println("largest S tested:", lastP)
	fmt.Println("cut:", cut)

	i := 0
	for i = 0; knownPentagons[i] < cut; i++ {
	}
	fmt.Println("ends:", knownPentagons[i-1], knownPentagons[i])
	fmt.Println("gap:", knownPentagons[i]- knownPentagons[i-1])

}
