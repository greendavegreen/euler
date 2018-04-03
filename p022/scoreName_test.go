package main

import (
	"testing"
	"strings"
	"github.com/greendavegreen/euler/scan"
)

func TestScore(t *testing.T) {
	cases := []struct {
		input         string
		expectedValue int
	}{
		{"A", 1},
		{"a", 1},
		{"b", 2},
	}
	for _, c := range cases {

		if r := score(c.input); r != c.expectedValue {
			t.Errorf("input: %v - expected %v, actual %v", c.input, c.expectedValue, r)
		}

	}
}

func TestScoreAll(t *testing.T) {
	cases := []struct {
		input         string
		expectedValue int
	}{
		{"A", 1},
		{"a,b", 5},
	}
	for _, c := range cases {

		scanner := scan.Comma(strings.NewReader(c.input))
		if r := scoreAll(scanner); r != c.expectedValue {
			t.Errorf("input: %v - expected %v, actual %v", c.input, c.expectedValue, r)
		}

	}
}
