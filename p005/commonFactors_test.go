package main

import (
	"testing"
)

func TestFactorMap(t *testing.T) {
	// setup test for anything up to the number 11
	cover := primesCovering(11)
	r := factorMap{cover: cover}
	cases := []struct {
		input          int
		expectedCounts []int
	}{
		{1, []int{0, 0, 0, 0, 0}},
		{2, []int{1, 0, 0, 0, 0}},
		{8, []int{3, 0, 0, 0, 0}},
		{9, []int{0, 2, 0, 0, 0}},
		{10, []int{1, 0, 1, 0, 0}},
		{11, []int{0, 0, 0, 0, 1}},
	}
	for _, c := range cases {
		r.factor(c.input)
		for i, v := range c.expectedCounts {
			if v != r.counts[i] {
				t.Errorf("input: %d - expected %v, actual %v", c.input, c.expectedCounts, r.counts)
			}
		}
	}
}


func TestLCM(t *testing.T) {
	cover := primesCovering(199)

	cases := []struct {
		input          []int
		expectedCounts []int
	}{
		{[]int{}, []int{0}},
		{[]int{7}, []int{0, 0, 0, 1}},
		{[]int{4,8}, []int{3}},
		{[]int{4,18}, []int{2, 2}},
		{[]int{60,9}, []int{2, 2, 1}},
	}
	for _, c := range cases {
		var maps []factorMap
		for _, i := range c.input {
			m := factorMap{cover: cover}
			m.factor(i)
			maps = append(maps, m)
		}

		actual := lcm(maps...)

		for i, v := range c.expectedCounts {
			if v != actual.counts[i] {
				t.Errorf("input: %d - expected %v, actual %v", c.input, c.expectedCounts, actual.counts[i])
			}
		}
	}
}