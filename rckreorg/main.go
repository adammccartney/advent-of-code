package main

import (
	"fmt"

	"github.com/adammccartney/aoc2022/scanner"
)

// split the lines in half, forming two "compartment"
func splitLine(line string) (string, string) {
	mid := len(line) / 2
	return line[:mid], line[mid:]
}

func convertStringToRunes(s string, converter map[rune]rune) map[rune]rune {
	for _, r := range s {
		converter[r] = r
	}
	return converter
}

func stringToIntArray(s string) []int {
	var arr []int
	for _, r := range s {
		prio := getRunePriority(r)
		arr = append(arr, prio)
	}
	return arr
}

func isLowercase(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

// Convert a rune to an int
func getRunePriority(r rune) int {
	// a = 1, ..., z = 26
	// A = 27, ..., Z = 52
	if isLowercase(r) {
		return int(r) - 96
	}
	if isUpper(r) {
		return int(r) - 38
	}
	return -1
}

func getPriorityRune(p int) rune {
	if p > 26 {
		return rune(p + 38)
	}
	return rune(p + 96)
}

type Compartment struct {
	contents   string
	priorities []int
}

func initCompartment(l int) Compartment {
	return Compartment{
		contents:   "",
		priorities: make([]int, l),
	}
}

func translatePriorities(priorities []int, out string) string {
	for _, p := range priorities {
		out += string(getPriorityRune(p))
	}
	return out
}

// Basic find common ints routine
// runs in O(n^2) time
// Possible Optimizations:
// 1. Sort the arrays, then use a binary search to find the common ints (O(nlogn))
// 2. Use a map to store the ints, then iterate over the other array to find
// the common ints (O(n))
func findCommonInt(a []int, b []int) []int {
	var out []int
	for _, v := range a {
		for _, w := range b {
			if v == w {
				out = append(out, v)
			}
		}
	}
	return out
}

func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

type Rucksack struct {
	lcomparts []Compartment
	rcomparts []Compartment
	common    []int
}

func initRucksack(l int) Rucksack {
	return Rucksack{
		lcomparts: make([]Compartment, l),
		rcomparts: make([]Compartment, l),
		common:    make([]int, l),
	}
}

func fillCompartment(line string, c Compartment) Compartment {
	c.contents = line
	c.priorities = stringToIntArray(line)
	return c
}

func makeCompartments(line string) []Compartment {
	l, r := splitLine(line)
	lc := initCompartment(len(l))
	rc := initCompartment(len(r))
	lc = fillCompartment(l, lc)
	rc = fillCompartment(r, rc)
	return []Compartment{lc, rc}
}

// Read in the lines, and split them into two compartments
// Sort the compartments, return the rucksack
func fillRucksack(lines []string, r Rucksack) Rucksack {
	r = initRucksack(len(lines))
	for i, line := range lines {
		comparts := makeCompartments(line)
		r.lcomparts[i] = comparts[0]
		r.rcomparts[i] = comparts[1]
		common := findCommonInt(comparts[0].priorities, comparts[1].priorities)
		r.common[i] = common[0]
	}
	return r
}

func sumCommon(r Rucksack) int {
	sum := 0
	for _, v := range r.common {
		sum += v
	}
	return sum
}

func calculateTotalFromInput(lines []string) int {
	r := initRucksack(len(lines))
	r = fillRucksack(lines, r)
	return sumCommon(r)
}

func main() {
	lines := scanner.ScanFileStrings("input.txt", []string{})
	total := calculateTotalFromInput(lines)
	fmt.Println("Total: ", total)
}
