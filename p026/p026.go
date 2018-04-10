package main

import "fmt"
// motivation -find decimal repeats
// technique -
// carry out long division algorithm step by step
// https://en.wikipedia.org/wiki/Long_division

// After an iteration,
//      - If you see a residue a second time, a cycle exists.
//      - If the reside is zero, then the decimal ends and no cycle exists.
//
// example:
//      1.0000000/7
//
//      	 .142857
//      7	1.000000    <- residue 1       <- 6 steps --
//      	  7                                         |
//      	  30        <- residue 3                    |
//      	  28                                        |
//      	   20       <- residue 2                    |
//      	   14                                       |
//      	    60      <- residue 6                    |
//      	    56                                      |
//      	     40     <- residue 4                    |
//      	     35                                     |
//      	      50    <- residue 5                    |
//      	      49                                    |
//      	       1    <- residue 1    - Cycle Found! -
//
func main() {
	maxCycleLength, maxCycleDivisor := 0, 0

	// find largest cycle in first 1000 items
	// iterate downward since the maximum cycle that can be produced by divisor i is i-1
	for divisor := 1000; divisor > 1; divisor-- {
		if maxCycleLength >= divisor {
			break       // past point where a new longest cycle is possible
		}
		// index = remainder numerator, numerator = iteration that started with that numerator
		residueMemory := make([]int, divisor)
		numerator, iteration := 1, 1
		for residueMemory[numerator] == 0 && numerator != 0 { // empty numerator means decimal terminates
			residueMemory[numerator] = iteration

			// do long division machinery - add zero, pull out as many copies of divisor as possible
			numerator *= 10
			numerator %= divisor
			iteration++
		}


		if numerator > 0 && iteration-residueMemory[numerator] > maxCycleLength {
			maxCycleLength = iteration - residueMemory[numerator]
			maxCycleDivisor = divisor
		}
	}

	fmt.Printf("maxCycle  1/%v has length %v", maxCycleDivisor, maxCycleLength)
}
