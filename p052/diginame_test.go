package main

import (
	"testing"
	"fmt"
)

func TestDigiSlice(t *testing.T) {
	cases := []struct {
		a  uint64
	}{
		{33},
		{1241},

	}
	for _, c := range cases {
		_,err := digiSlice(c.a)
		if err == nil {
			fmt.Println("expected error for duplicate characters", c.a)
		}
	}
}


func TestDigiUnequal(t *testing.T) {
	cases := []struct {
		a  uint64
		b uint64
	}{
		{1,2},
		{123, 12},
		{12, 123},
	}
	for _, c := range cases {
		a,_ := digiSlice(c.a)
		b,_ := digiSlice( c.b)
		if digiEqual(a,b) {
			fmt.Println("expected unequal", c.a, c.b)
		}
	}
}

func TestDigiEqual(t *testing.T) {
	cases := []struct {
		a  uint64
		b uint64
	}{
		//{1,2},
		{12, 21},
		{216, 621},
		{21, 1212},
	}
	for _, c := range cases {
		a,_ := digiSlice(c.a)
		b,_ := digiSlice( c.b)
		if !digiEqual(a,b) {
			fmt.Println("failed to equal", c.a, c.b)
		}
	}
}