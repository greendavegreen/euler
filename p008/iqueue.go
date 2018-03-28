package main

type IQueue []int

func (q *IQueue) Push(n int) {
	*q = append(*q, n)
}

func (q *IQueue) Pop() (n int) {
	n = (*q)[0]
	*q = (*q)[1:]
	return
}

func (q *IQueue) Len() int {
	return len(*q)
}