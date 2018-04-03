package main

import (
	"fmt"
	"math/big"
	"unicode"
)

func streamDigits(input string, c chan int) {
	for _, r := range []rune(input) {
		if unicode.IsNumber(r) {
			c <- int(r - '0')
		}
	}
	close(c)
}

func main() {
	bit := big.NewInt(1)
	bit.Lsh(bit, 1000)

	ch := make(chan int)
	go streamDigits(bit.String(), ch)

	sum := 0
	for digit := range ch {
		sum += digit
	}
	fmt.Printf("sum of digits in 2^%d = %v\n", 1000, sum)
}
