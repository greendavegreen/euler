package main

import "testing"

func TestPrimeFactors(t *testing.T) {
	ft := NewFactorTable(25)
	cases := []struct {
		input           int
		expectedFactors []int
	}{
		{1, []int{}},
		{2, []int{2}},
		{5, []int{5}},
		{7, []int{7}},
		{8, []int{2, 2, 2}},
		{9, []int{3, 3}},
		{10, []int{2, 5}},
		{11, []int{11}},
		{14, []int{2, 7}},
		{22, []int{2, 11}},
		{25, []int{5, 5}},
	}
	for _, c := range cases {
		f := ft.factor(c.input)
		primeFactors := f.primeFactors();
		if len(primeFactors) != len(c.expectedFactors) {
			t.Errorf("input: %d - expected %v, actual %v", c.input, c.expectedFactors, primeFactors)
		} else {
			for i, v := range c.expectedFactors {
				if v != primeFactors[i] {
					t.Errorf("input: %d - expected %v, actual %v", c.input, v, primeFactors[i])
				}
			}
		}

	}
}

func TestValue(t *testing.T) {
	ft := NewFactorTable(25)
	cases := []struct {
		input         int
		expectedValue int
	}{
		{7, 7},
	}
	for _, c := range cases {
		f := ft.factor(c.input)

		if f.value() != c.expectedValue {
			t.Errorf("input: %d - expected %v, actual %v", c.input, c.expectedValue, f.value())
		}

	}
}

func TestDivisorCount(t *testing.T) {
	ft := NewFactorTable(25)
	cases := []struct {
		input         int
		expectedValue int
	}{
		{2, 2},
		{4, 3},
		{6, 4},
		{12, 6},
		{16, 5},
		{17, 2},
		{24, 8},
	}
	for _, c := range cases {
		f := ft.factor(c.input)

		if dc := f.divisorCount(); dc != c.expectedValue {
			t.Errorf("input: %d - expected %v, actual %v", c.input, c.expectedValue, dc)
		}

	}
}

func TestFtLCM(t *testing.T) {
	ft := NewFactorTable(200)

	cases := []struct {
		input    []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{7}, 7},
		{[]int{4, 8}, 8},
		{[]int{4, 18}, 36},
		{[]int{60, 9}, 180},

		{[]int{19, 23}, 19 * 23},
		{[]int{2, 3, 197, 199}, 2 * 3 * 197 * 199},
	}
	for _, c := range cases {
		items := make([]factored, len(c.input))
		for i, v := range c.input {
			items[i] = *ft.factor(v)
		}

		actual := ft.lcm(items...)
		if actual.v != c.expected {
			t.Errorf("input: %d - expected %v, actual %v", c.input, c.expected, actual.v)
		}
	}
}
