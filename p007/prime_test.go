package main

import "testing"


func TestNthPrime(t *testing.T) {
	cases := []struct {
		prime_number  int
		desired int
	}{
		{1, 2},
		{2, 3},
		{3, 5},
		{5, 11},
		{6, 13},
		{ 5000, 48611},
	}
	for _, c := range cases {
		actual := nthPrime(c.prime_number)
		if actual != c.desired {
			t.Errorf("expected %d, actual %d", c.desired, actual)
		}
	}
}

func TestPrimeList(t *testing.T) {
	cases := []struct {
		prime_number  int
		desired [] int
	}{
		{2, []int{2,3}},
		{7, []int{2,3,5,7,11,13,17}},
	}
	for _, c := range cases {
		actual := primeList(c.prime_number)
		for i, v := range actual {
			if v != c.desired[i] {
				t.Errorf("expected %d, actual %d", c.desired[i], v)
			}
		}
	}
}
