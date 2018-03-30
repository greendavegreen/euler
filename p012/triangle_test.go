package main

import "testing"

func TestTriangles(t *testing.T) {

	//t := triangles()
	gen := triangles()

	actual := [] int{gen(), gen(), gen(), gen(), gen(), gen(), gen()}
	expected := []int{1, 3, 6, 10, 15, 21, 28}

	for i, v := range actual {
		if v != expected[i] {
			t.Errorf("expected %d, actual %d", expected[i], v)
		}
	}
}
