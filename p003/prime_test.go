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


func TestFactor(t *testing.T) {
	cases := []struct {
		itemToFactor int
		desiredFactors [] int
	}{
		{34, []int{2, 17}},
		{29, []int{29}},
		{9, []int{3, 3}},
		{12, []int{2, 2, 3}},
		{24, []int{2,2,2,3}},
		{13195, []int{5, 7, 13, 29}},
		{600851475143, []int{71, 839, 1471, 6857}},
	}
	for _, c := range cases {
		actual := pfactor(c.itemToFactor)
		if len(actual) != len(c.desiredFactors) {
			t.Errorf("expected %v, actual %v", c.desiredFactors, actual)
		} else {
			for i, v := range c.desiredFactors {
				if v != actual[i] {
					t.Errorf("expected %d, actual %d", v, actual[i])
				}
			}
		}
	}
}



func TestPrimesCovering(t *testing.T) {
	cases := []struct {
		maxPrime  int
		desired [] int
	}{
		{2, []int{2}},
		{4, []int{2, 3, 5}},
		{5, []int{2, 3, 5}},
		{18, []int{2, 3, 5, 7, 11, 13, 17, 19}},
	}
	for _, c := range cases {
		actual := primesCovering(c.maxPrime)
		if len(actual) != len(c.desired) {
			t.Errorf("expected %v, actual %v", c.desired, actual)
		}
		for i, v := range c.desired {
			if v != actual[i] {
				t.Errorf("expected %d, actual %d", v, actual[i])
			}
		}
	}
}