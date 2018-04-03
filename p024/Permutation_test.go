package main

import (
	"testing"
)

func TestFact(t *testing.T) {
	cases := []struct {
		input         int
		expectedValue int
	}{
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
	}
	for _, c := range cases {
		actual := fact(c.input)

		if actual != c.expectedValue {
			t.Errorf("input: %d - expected %v, actual %v", c.input, c.expectedValue, actual)
		}

	}
}

func TestLexPermutation(t *testing.T) {
	cases := []struct {
		input    []int
		pnum     int
		expected []int
	}{
		{[]int{0, 1, 2, 3}, 1, []int{0, 1, 2, 3}},
		{[]int{0, 1, 2, 3}, 6, []int{0, 3, 2, 1}},
		{[]int{0, 1, 2, 3}, 7, []int{1, 0, 2, 3}},
	}
	for _, c := range cases {
		perm := lexPermutation(c.input, c.pnum)
		if len(perm) != len(c.expected) {
			t.Errorf("input: %d - lengths differ expected %v, actual %v", c.input, c.expected, perm)
		} else {
			for i, v := range c.expected {
				if v != perm[i] {
					t.Errorf("input: %d - position %v expected %v, actual %v", c.input, i, v, perm[i])
				}
			}
		}

	}
}
