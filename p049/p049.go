package main

import (
	"github.com/kavehmz/prime"
	"fmt"
	"strconv"
	"strings"
)

func allIn(a string, b string) bool {
	for i := 0; i < len(a); i++ {
		if strings.Index(b, a[i:i+1]) == -1 {
			return false
		}
	}
	return true
}

func perms(a string, b string) bool {
	return allIn(a, b) && allIn(b, a)
}

func main() {

	ps := prime.Primes(10000)
	i := 0
	for ; ps[i] < 1000; i++ {
	}

	fdps := ps[i:]
	fdpLen := len(fdps)
	fmap := make(map[uint64]bool)
	for _, x := range fdps {
		fmap[x] = true
	}

	c := 0
	for i := 0; i < fdpLen; i++ {
		si := strconv.FormatUint(fdps[i], 10)
		for j := i + 1; j < fdpLen; j++ {
			sj :=  strconv.FormatUint(fdps[j], 10)
			if perms(si,sj) {
				third := 2*fdps[j] - fdps[i]
				if fmap[third] {
					st := strconv.FormatUint(third, 10)
					if perms(si,st) {
						fmt.Println(fdps[i], fdps[j], third)
						c++
					}
				}
			}

		}

	}
	fmt.Println(c)
}
