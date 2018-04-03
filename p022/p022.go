package main

import (
	"log"
	"os"
	"fmt"
	"github.com/greendavegreen/feed"
	"github.com/greendavegreen/scan"
	"bufio"
	"sort"
	"strings"
)

func score(name string) int {
	sum := 0
	for digit := range feed.AsRank(name) {
		sum += digit
	}
	return sum
}

func scoreSlice(names []string) int {
	sort.Strings(names)
	sum := 0
	for i, n := range names {
		sum += score(n) * (i + 1)
	}
	return sum
}


func scoreAll(sc *bufio.Scanner) int {
	names := []string{}
	for sc.Scan() {
		item := strings.Trim(sc.Text(),"\"")
		names = append(names, item)
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	return scoreSlice(names)
}

func main() {
	file, err := os.Open("./p022_names.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	value := scoreAll(scan.Comma(file))

	fmt.Printf("score %v\n", value)
}
