package hand

import (
	"testing"
	"fmt"
)



func TestHandScore(t *testing.T) {
	cases := []struct {
		hand []string
		val htype
	}{
		{[]string{"3C", "3D", "3H", "3S", "2H"},FOUR_OF_A_KIND},
	}
	for _, c := range cases {
		h := New(c.hand)
		if h.rank.t != c.val {
			fmt.Println("incorrect rank", c.hand, c.val, h.rank.t)
		}
	}
}