package main

import "testing"

func TestDigitToText(t *testing.T) {
	cases := []struct {
		digit  int
		desired string
	}{
		{1, "one"},
		{10, "ten"},
		{19, "nineteen"},
	}
	for _, c := range cases {
		actual := digitToText(c.digit)
		if actual != c.desired {
			t.Errorf("desired %d, actual %d", c.desired, actual)
		}
	}
}

func TestCentsToText(t *testing.T) {
	cases := []struct {
		digit  int
		desired string
	}{
		{20, "twenty"},
		{31, "thirty-one"},
		{42, "forty-two"},
	}
	for _, c := range cases {
		actual := centsToText(c.digit)
		if actual != c.desired {
			t.Errorf("desired %s, actual %s", c.desired, actual)
		}
	}
}

func TestHundredToText(t *testing.T) {
	cases := []struct {
		digit  int
		desired string
	}{
		{3, "three hundred"},
		{5, "five hundred"},
	}
	for _, c := range cases {
		actual := hToText(c.digit)
		if actual != c.desired {
			t.Errorf("desired %s, actual %s", c.desired, actual)
		}
	}
}


func TestIntegerToText(t *testing.T) {
	cases := []struct {
		digit  int
		desired string
	}{
		{1, "one"},
		{19, "nineteen"},
		{20, "twenty"},
		{21, "twentyone"},
		{99, "ninetynine"},
		{100, "onehundred"},
		{101, "onehundredandone"},
		{155, "onehundredandfiftyfive"},
		{1000, "onethousand"},

	}
	for _, c := range cases {
		actual := iToText(c.digit)
		if actual != c.desired {
			t.Errorf("desired %s, actual %s", c.desired, actual)
		}
	}
}


func TestIToTextLength(t *testing.T) {
	cases := []struct {
		seqLimit int
		length int
	}{
		{342,23},
		{115,20},
	}
	for _, c := range cases {
		actual := len(iToText(c.seqLimit))
		if actual != c.length {
			t.Errorf("desired %d, actual %d", c.length, actual)
		}
	}
}

func TestSeqLength(t *testing.T) {
	cases := []struct {
		seqLimit int
		length int
	}{
		{1, 3},
		{5, 19},
	}
	for _, c := range cases {
		actual := seqTextLength(c.seqLimit)
		if actual != c.length {
			t.Errorf("desired %d, actual %d", c.length, actual)
		}
	}
}