package main

import "testing"

func TestEvenFibs(t *testing.T) {
	cases := []struct {
		limit  int
		wanted int
	}{
		{8, 2},
		{9, 10},
		{34, 10},
		{35, 44},

	}
	for _, c := range cases {

		actual := evenFibs(c.limit)
		if actual != c.wanted {
			t.Errorf("expected %d, actual %d", c.wanted, actual)
		}
	}
}

func TestFib(t *testing.T) {
	f := fib()
	actual := [] int{f(), f(), f(), f(), f(), f(), f()}
	expected := []int{1, 1, 2, 3, 5, 8, 13}

	for i, v := range actual {
		if v != expected[i] {
			t.Errorf("expected %d, actual %d", expected[i], v)
		}
	}
}
