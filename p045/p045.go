package main

import "fmt"

func triangles() func() int {
	inc, i := 1, 0
	return func() int {
		i, inc = i+inc, inc+1
		return i
	}
}

func pentagons() func() int {
	inc, i := 1, 0
	return func() int {
		i, inc = i+inc, inc+3
		return i
	}
}

func hexagons() func() int {
	inc, i := 1, 0
	return func() int {
		i, inc = i+inc, inc+4
		return i
	}
}

func main() {
	var x int

	previousMatched := false
	last := 0
	hitCount := 0
	for f, g, h := triangles(), pentagons(), hexagons(); hitCount < 3 ; f, g, h = g, h, f {
		for x = f(); x < last; x = f() { // advance sequence f until match or new value
		}
		if x == last && previousMatched {  // match following previous match means f=g=h at value x
			fmt.Println("solution", x)
			hitCount++
		}
		previousMatched = x == last
		last = x
	}
}
