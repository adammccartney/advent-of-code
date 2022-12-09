package main

import (
	"testing"

	"github.com/adammccartney/aoc2022/scanner"
)

func TestGetStringPriority(t *testing.T) {
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
		got := getRunePriority(c.in)
		if got != c.want {
			t.Errorf("getStringPriority(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestFillCompartment(t *testing.T) {
	cases := []struct {
		in   string
		want Compartment
	}{
		{"abc", Compartment{"abc", []int{1, 2, 3}}},
	}
	for _, c := range cases {
		got := fillCompartment(c.in, c.want)
		if got.contents != c.want.contents {
			t.Errorf("fillCompartment(%q).contents == %q, want %q", c.in, got.contents, c.want.contents)
		}
		if len(got.priorities) != len(c.want.priorities) {
			t.Errorf("fillCompartment(%q).priorities == %q, want %q", c.in, got.priorities, c.want.priorities)
		}
		for i := range got.priorities {
			if got.priorities[i] != c.want.priorities[i] {
				t.Errorf("fillCompartment(%q).priorities == %q, want %q", c.in, got.priorities, c.want.priorities)
			}
		}
	}
}

func TestFindCommonInt(t *testing.T) {
	cases := []struct {
		in1  []int
		in2  []int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 1}, []int{1}},
		{[]int{1, 2, 3, 4, 5}, []int{10, 11, 12, 13, 4}, []int{4}},
		{[]int{1, 2, 3, 4, 5}, []int{10, 11, 12, 13, 14}, []int{}},
	}
	for _, c := range cases {
		got := findCommonInt(c.in1, c.in2)
		if !sliceEqual(got, c.want) {
			t.Errorf("findCommonElt(%q, %q) == %q, want %q", c.in1, c.in2, got, c.want)
		}
	}
}

func TestCalculateFromInput(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"test.txt", 157},
	}
	for _, c := range cases {
		lines := scanner.ScanFileStrings(c.in, []string{})
		total := calculateTotalFromInput(lines)
		if total != c.want {
			t.Errorf("calculateFromInput(%q) == %q, want %q", c.in, total, c.want)
		}
	}
}
