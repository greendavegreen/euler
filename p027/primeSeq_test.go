package main

import (
	"testing"
	"github.com/greendavegreen/factor"
)

func TestFact(t *testing.T) {
	cases := []struct {
		a         int
		b         int
		expectedValue int
	}{
		{1, 41, 40},
		{-79, 1601, 80},
	}
	ft := factor.NewFactorTable(1000000)

	for _, c := range cases {
		actual, err := maxPrimeSequence( ft, c.a, c.b)
		if (err != nil) {
			t.Errorf("unexpected error during Sequence Explore %v", err)
		}
		if actual != c.expectedValue {
			t.Errorf("input: %d,%d - expected %v, actual %v", c.a, c.b, c.expectedValue, actual)
		}

	}
}
