package main

import (
	"unicode"
	"fmt"
	"strconv"
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
	bigSum := 0

	for i := 10; i < 10000000; i++ {
		ch := make(chan int)
		go streamDigits(strconv.Itoa(i), ch)

		sum := 0
		for digit := range ch {
			sum += digit * digit * digit * digit * digit
		}
		if sum == i {
			bigSum += sum
			fmt.Println(sum)
		}
	}

	fmt.Printf("The sum of these is : %v", bigSum)

}
