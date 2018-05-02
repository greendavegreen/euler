package main

import (
	"fmt"
	"github.com/kavehmz/prime"
)

func genOddComposites(ulimit uint64, c chan uint64) {
	var i, j, oddComp uint64
	outer:
		for i = 1;; i++ {
		j = i
		oddComp = 4*i*j + 2*(i+j) + 1
		for {
			if oddComp > ulimit {
				if j == i {
					break outer
				}
				break
			}
			c <- oddComp
			j++
			oddComp = 4*i*j + 2*(i+j) + 1
		}
	}
	close(c)
}

func main() {
	var base uint64 = 100
	maxValueTested := base * base * 2

	pr := prime.Primes(maxValueTested)
	sq := make([]uint64, base)
	var i uint64

	for i = 0; i < base; i++ {
		sq[i] = 2 * (i + 1) * (i + 1)
	}

	oddComposites := make(chan uint64)
	go genOddComposites(maxValueTested, oddComposites)

outer:
	for oc := range oddComposites {
		for _, p:=range pr {
			if p < oc {

				for _, s := range sq {
					if s + p == oc {
						//fmt.Println(oc, p, s)
						continue outer
					}
				}

			} else {
				break
			}
		}
		fmt.Println(oc, "-- UNFOUND")

	}
}

func main2() {
	var i, j, oddComp uint64
	var base uint64 = 5
	maxValueTested := base * base * 2

	//pr := prime.Primes(maxValueTested)
	sq := make([]uint64, base)


	for i = 0; i < base; i++ {
		sq[i] = 2 * (i + 1) * (i + 1)
	}

	for i = 1; j != i; i++ {
		j = i
		oddComp = 4*i*j + 2*(i+j) + 1
		for {
			if oddComp > maxValueTested {
				break
			}

			// oddComp

			j++
			oddComp = 4*i*j + 2*(i+j) + 1
		}
	}
}
