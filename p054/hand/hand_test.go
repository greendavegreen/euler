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
		{[]string{"2C", "3C", "4C", "6C", "7S"},HIGH},
		{[]string{"3C", "3D", "AH", "KS", "2H"},ONE_PAIR},
		{[]string{"3C", "3D", "AH", "2S", "2H"},TWO_PAIR},
		{[]string{"3C", "3D", "3H", "AS", "2H"},THREE_OF_A_KIND},
		{[]string{"2C", "3C", "4C", "5C", "6S"},STRAIGHT},
		{[]string{"2C", "3C", "4C", "6C", "7C"},FLUSH},
		{[]string{"3C", "3D", "3H", "2S", "2H"},FULL_HOUSE},
		{[]string{"3C", "3D", "3H", "3S", "2H"},FOUR_OF_A_KIND},
		{[]string{"2C", "3C", "4C", "5C", "6C"},STRAIGHT_FLUSH},

	}
	for _, c := range cases {
		h := New(c.hand)
		if htype(h.rank.vals[0]) != c.val {
			fmt.Printf("incorrect rank %v = %v EXPECTED %v", c.hand, htype(h.rank.vals[0]).toString(), c.val.toString(), )
		}
	}
}