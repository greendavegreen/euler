package main

import "fmt"

var digitNames = []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
var tensNames = []string{"", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

func digitToText(i int) string {
	return digitNames[i]
}

func tToText(t int) string {
	return tensNames[t]
}

func centsToText(c int) string {
	switch {
	case c < 20:
		{
			return digitToText(c)
		}
	default:
		{
			return tToText(c/10) + digitToText(c%10)
		}
	}
}

func hToText(h int) string {
	if h > 0 {
		return digitToText(h) + "hundred"
	} else {
		return ""
	}
}

func thousandToText(h int) string {
	if h > 0 {
		return digitToText(h) + "thousand"
	} else {
		return ""
	}
}

func iToText(i int) string {
	result := thousandToText(i / 1000)

	result += hToText((i % 1000) / 100)

	if r := i % 100; r > 0 {
		if len(result) > 0 {
			result += "and"
		}
		result += centsToText(r)
	}

	return result
}

func seqTextLength(seqLength int) int {
	sum := 0
	for i := 1; i < seqLength+1; i++ {
		sum += len(iToText(i))
	}
	return sum
}

func main() {

	res := seqTextLength(1000)
	fmt.Printf("The lettercount is: %d", res)
}
