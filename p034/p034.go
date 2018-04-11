package main

import (
	"strconv"
	"fmt"
	"strings"
)

func main() {
	i := 7654
	fmt.Println(i,rotate(i))
}

func rotate(i int) int {
	str := strings.Builder{}

	s := strconv.Itoa(i)
	str.WriteString(s[len(s)-1:])
	str.WriteString(s[:len(s)-1 ])

	v, _ := strconv.Atoi(str.String())
	return v
}