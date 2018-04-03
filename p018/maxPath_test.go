package main

import (
	"testing"
)

func TestMaxPath(t *testing.T) {
	vals := [][]int{
		{1},
		{2, 3},
		{4, 5, 6},
	}
	expected := 10

	v := maxPath(vals)
	if v != expected {
		t.Errorf("desired %d, actual %d", expected, v)
	}
}
