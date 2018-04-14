package main

import (
	"github.com/greendavegreen/feed"
	"fmt"
	"strings"
	"log"
	"os"
	"github.com/greendavegreen/scan"
)

func triangles() func () int {
	inc := 1
	i := 0
	return func() int {
		i, inc = i+inc, inc + 1
		return i
	}
}

func score(name string) int {
	sum := 0
	for digit := range feed.AsRank(name) {
		sum += digit
	}
	return sum
}

func main() {
	file, err := os.Open("./p042_words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	maxWordLength := 15
	t:=triangles()
	tmap := make(map[int]bool)
	for i:=t(); i<maxWordLength*26+1; i=t() {
		tmap[i]=true
	}

	tCount := 0
	for sc := scan.Comma(file);sc.Scan(); {
		item := strings.Trim(sc.Text(),"\"")
		if len(item) > maxWordLength {
			fmt.Println("max word length exceeded ", item)
			break
		}
		s := score(item)
		if tmap[s] {
			tCount++
			fmt.Println(item, "is traiangular")
		}
	}

	fmt.Println(tCount, " triangle words in file.")
}
