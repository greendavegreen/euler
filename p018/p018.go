package main

import (
	"strings"
	"strconv"
	"fmt"
)

var input string = `
75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23
`

func buildSlices(input string) (*[][]int, error) {
	var ss [][]int
	for _, line := range strings.Split(input, "\n") {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		fields := strings.Fields(line)
		row := make([]int, len(fields))
		for i, field := range fields {
			v, err := strconv.Atoi(field)
			if err != nil {
				return nil, err
			}
			row[i] = v
		}
		ss = append(ss, row)
	}
	return &ss, nil
}

func maxI(i int, j int) int {
	if i >= j {
		return i
	} else {
		return j
	}
}

func maxPath(weights [][]int) int {
	sums := make([][]int, len(weights))
	for i := range sums {
		sums[i] = make([]int, len(weights[i]))
	}

	lastRow := len(sums) - 1
	for i := range sums[lastRow] {
		sums[lastRow][i] = weights[lastRow][i]
	}

	for r := len(sums) - 2; r >= 0; r-- {
		for c := range sums[r] {
			sums[r][c] = weights[r][c] + maxI(sums[r+1][c], sums[r+1][c+1])
		}
	}

	return sums[0][0]
}

func main() {
	if triangle, err := buildSlices(input); err == nil {
		m := maxPath(*triangle)
		fmt.Printf("Maximum path =%d", m)
	}
}
