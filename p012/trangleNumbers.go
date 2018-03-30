package main

func triangles() func() int {
	a, b := 0, 1
	return func() int {
		a, b = a + b, b + 1
		return a
	}
}

