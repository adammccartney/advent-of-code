package main

import (
	"testing"
)

func TestgetStringPriority(t *testing.T) {
	cases := []struct {
		in   rune
		want int
	}{
		{'a', 1},
		{'z', 26},
		{'A', 27},
		{'Z', 52},
	}
	for _, c := range cases {
		got := getStringPriority(c.in)
		if got != c.want {
			t.Errorf("getStringPriority(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
