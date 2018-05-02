package main

import (
	"strconv"
	"errors"
)

// sums all items that are evenly divisible any divisor and less than limit

func digiSlice(input uint64) (map[rune]bool, error) {
	pstring := strconv.FormatUint(input, 10)

	aMap := make(map[rune]bool)
	for _, v := range []rune(pstring) {
		aMap[v] = true
	}

	if len(aMap) != len(pstring) {
		return nil, errors.New("digits non-distinct")
	}

	return aMap, nil
}

func digiEqual(a map[rune]bool, b map[rune]bool) bool {
	if len(a) != len(b) {
		return false
	}

	for k, _ := range a {
		if !b[k] {
			return false
		}
	}

	for k, _ := range b {
		if !a[k] {
			return false
		}
	}

	return true
}
