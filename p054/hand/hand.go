package hand

import (
	"fmt"
)

type htype int

const (
	HIGH htype = iota
	FULL_HOUSE
	THREE_OF_A_KIND
	TWO_PAIR
	ONE_PAIR
	FOUR_OF_A_KIND
)

func (ht htype) toString() string {
	var res string
	switch ht {
	case FOUR_OF_A_KIND:
		res = "four of a kind"
	default:
		res = "unknown"
	}
	return res
}

type power struct {
	t     htype
	hvals []int
}

func (p power) Display() {
	fmt.Println(p.t.toString(), "-",p.hvals[0])
}

type hand struct {
	Raw    []string
	rank   power
	values map[int]int
	suits  map[rune]int
}

func (h hand) Display() {
	fmt.Println(h.Raw, h.values, h.suits)
	h.rank.Display()
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
	values := make(map[int]int)
	s := make(map[rune]int)

	for _, c := range input {
		a := []rune(c)
		values[runeToValue[a[0]]]++
		s[a[1]]++
	}

	p := power{}
	switch c := len(values); c {
	case 2:
		p.t = FOUR_OF_A_KIND
		for k, v := range values {
			if v == 4 {
				p.hvals = []int{k}
			}
		}
	case 3:
	case 4:
		fmt.Println("Linux.")
	default:
		// test for flush
		// test for straight
		// order values and store in values
	}

	h := hand{input, p, values, s}
	return h
}
