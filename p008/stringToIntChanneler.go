package main

import (
	"unicode"
)

func streamDigits(input string, c chan int) {
	for _, r := range []rune(input) {
		if unicode.IsNumber(r) {
			digit := int(r - '0')
			c <- digit
		}
	}
	close(c)
}
