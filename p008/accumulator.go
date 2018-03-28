package main

import "fmt"

type accumulator struct {
	buffer          []int
	head            int
}

func NewAccumulator(size int) *accumulator {
	b := make([]int, size, size)
	for i, _ := range b {
		b[i] = 1
	}
	return &accumulator{
		buffer:          b,
		head:            0,
	}
}

func (a *accumulator) dump() {
	i := a.head
	sentry := i
	size := len(a.buffer)
	for {
		fmt.Printf("%v ", a.buffer[i])
		i = (i + 1) % size
		if i == sentry {
			fmt.Println("")
			break
		}
	}
}

func (a *accumulator) value() int {
	result := 1
	for _, v := range a.buffer {
		result *= v
	}
	return result
}


func (a *accumulator) ingest(x int) {
	a.buffer[a.head] = x
	a.head = (a.head + 1) % len(a.buffer)
}
