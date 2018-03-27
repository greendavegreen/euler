package main

import "testing"


func TestOddSieve(t *testing.T) {
	f := oddSieve()
	actual := [] int{f(), f(), f(), f(), f(), f(), f(), f(), f(), f()}
	expected := []int{3, 5, 7, 11, 13, 17, 19, 23, 29, 31}

	for i, v := range actual {
		if v != expected[i] {
			t.Errorf("expected %d, actual %d", expected[i], v)
		}
	}
}
