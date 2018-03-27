package main

import "testing"

func TestPalindrome(t *testing.T) {
	cases := []struct {
		input    int
		expected bool
	}{
		{1, true},
		{11, true},
		{12, false},
		{111, true},
		{121, true},
		{221, false},
	}
	for _, c := range cases {
		actual := palindrome(c.input)
		if actual != c.expected {
			t.Errorf("input: %d expected %v, actual %v", c.input, c.expected, actual)
		}
	}
}
