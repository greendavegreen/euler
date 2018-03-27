package main

import "testing"

func TestOddSieve(t *testing.T) {
	f := oddSieve(9)
	actual := [] int{f(), f(), f(), f(), f(), f(), f(), f(), f(), f()}
	expected := []int{5, 7, 11, 13, 17, 19, 23, 29, 31, 37}

	for i, v := range actual {
		if v != expected[i] {
			t.Errorf("expected %d, actual %d", expected[i], v)
		}
	}
}

func TestNthPrime(t *testing.T) {
	cases := []struct {
		prime_number  int
		desired int
	}{
		{1, 2},
		{2, 3},
		{3, 5},
		{5, 11},
		{ 5000, 48611},
	}
	for _, c := range cases {
		actual := nthPrime(c.prime_number)
		if actual != c.desired {
			t.Errorf("expected %d, actual %d", c.desired, actual)
		}
	}
}
