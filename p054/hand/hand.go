package hand

import (
	"fmt"
)

type htype int

const (
	HIGH htype = iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	STRAIGHT
	FLUSH
	FULL_HOUSE
	FOUR_OF_A_KIND
	STRAIGHT_FLUSH
)

func (ht htype) toString() string {
	var res string
	switch ht {
	case FOUR_OF_A_KIND:
		res = "Four of a Kind"
	case FULL_HOUSE:
		res = "Full House"
	case THREE_OF_A_KIND:
		res = "Three of a Kind"
	case TWO_PAIR:
		res = "Two Pair"
	case ONE_PAIR:
		res = "One Pair"
	case FLUSH:
		res = "Flush"
	case STRAIGHT:
		res = "Straight"
	case HIGH:
		res = "High"
	default:
		res = "unknown"
	}
	return res
}

type power struct {
	vals []int
}

func (p power) Display() {
	fmt.Println(htype(p.vals[0]).toString(), "-",p.vals[1])
}

type hand struct {
	Raw    []string
	rank   power
	values map[int]int
	suits  map[rune]int
}

func (h hand) Display() {
	fmt.Println(h.Raw, htype(h.rank.vals[0]).toString(), "-",h.rank.vals)
}

func sinsert(v int, items []int) {
	for i := 0; i<len(items);i++{
		if v>items[i] {
			v, items[i] = items[i],v
		}
	}
}

func New(input []string) hand {
	runeToValue := map[rune]int{
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'T': 10,
		'J': 11,
		'Q': 12,
		'K': 13,
		'A': 14,
	}
	values, s := buildMaps(input, runeToValue)

	p := power{[]int{0, 0, 0, 0, 0, 0}}
	switch c := len(values); c {
	case 2:
		for k, v := range values {
			if v == 4 {
				p.vals[0] = int(FOUR_OF_A_KIND)
				p.vals[1] = k
				break
			}
			if v == 3 {
				p.vals[0] = int(FULL_HOUSE)
				p.vals[1] = k
				break
			}
		}
	case 3:
		for k, v := range values {
			if v == 3 {
				p.vals[0] = int(THREE_OF_A_KIND)
				p.vals[1] = k
				break
			}
			if v == 2 {
				p.vals[0] = int(TWO_PAIR)
				if k > p.vals[1] {
					p.vals[1], p.vals[2] = k, p.vals[1]
				} else {
					p.vals[2] = k
				}
			}
			if v == 1 {
				sinsert(k, p.vals[3:])
			}
		}
	case 4:
		for k, v := range values {
			if v == 2 {
				p.vals[0] = int(ONE_PAIR)
				p.vals[1] = k
			}
			if v == 1 {
				sinsert(k, p.vals[2:])
			}
		}
	default:
		for k, _ := range values {
			sinsert(k, p.vals[1:])
		}
		if p.vals[1] - p.vals[5] == 4 {
			if len(s) == 1 {
				p.vals[0] = int(STRAIGHT_FLUSH)
			} else {
				p.vals[0] = int(STRAIGHT)
			}
		} else {
			if len(s) == 1 {
				p.vals[0] = int(FLUSH)
			} else {
				p.vals[0] = int(HIGH)
			}
		}
	}

	h := hand{input, p, values, s}
	return h
}

func buildMaps(input []string, runeToValue map[rune]int) (map[int]int, map[rune]int) {
	values := make(map[int]int)
	s := make(map[rune]int)
	for _, c := range input {
		a := []rune(c)
		values[runeToValue[a[0]]]++
		s[a[1]]++
	}
	return values, s
}
