package main

import (
	"time"
	"fmt"
)

//How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000

func main() {
	start := time.Date(1901, 1, 1, 1, 0, 0, 0, time.UTC)
	end := time.Date(2000, 12, 31, 1, 0, 0, 0, time.UTC)

	count := 0
	for d := start; d.Before(end); d = d.AddDate(0, 1, 0) {
		if d.Weekday() == time.Sunday {
			count += 1
		}
	}

	fmt.Printf("%d sundays between %v and %v", count, start, end)
}
