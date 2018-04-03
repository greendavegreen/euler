package main

import "testing"

func TestCollatz(t *testing.T) {
	c := collatz(13)

	actual := [] int{c(), c(), c(), c(), c(), c(), c(), c(), c(), c()}

	expected := [] int{13, 40, 20, 10, 5, 16, 8, 4, 2, 1}

	for i, v := range actual {
		if v != expected[i] {
			t.Errorf("expected %d, actual %d", expected[i], v)
		}
	}
}

func TestCollatzLength(t *testing.T) {
	c := collatzLen(13)
	if c!= 10 {
		t.Errorf("expected %d, actual %d", 10, c)
	}
}