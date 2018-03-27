package main

import "testing"

func TestMultiples(t *testing.T) {
	cases := []struct {
		divisors []int
		limit    int
		wanted     int
	}{
		{[]int{3, 5}, 10, 23},
	}
	for _, c := range cases {

		actual := multiples(c.divisors, c.limit)
		if actual != c.wanted {
			t.Errorf("expected %d, actual %d", actual, c.wanted)
		}
	}
}
