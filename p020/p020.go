package main

import (
	"unicode"
	"math/big"
	"fmt"
	"github.com/greendavegreen/euler/feed"
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
	var fact int64 = 100
	var f big.Int

	f.MulRange(1, fact)

	sum := 0
	for digit := range feed.AsDigits(f.String()) {
		sum += digit
	}

	fmt.Printf("sum of digits in %d! = %v\n", fact, sum)
}
