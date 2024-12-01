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

func isLowercase(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

// Convert a rune to an int
func getStringPriority(r rune) int {
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

func binaryChop(c int, sorted []int) bool {
	low := 0
	high := len(sorted) - 1
	for low <= high {
		mid := (low + high) / 2
		if sorted[mid] == c {
			return true
		}
		if sorted[mid] < c {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

func mergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	mid := len(a) / 2
	return merge(mergeSort(a[:mid]), mergeSort(a[mid:]))
}

func merge(a, b []int) []int {
	final := []int{}
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
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

func sortStringToIntArray(s string, priorities []int) []int {
	// Convert the string to an array of priorities
	fmt.Println("s: ", s)
	for i := 0; i < len(s); i++ {
		char := string(s[i])
		fmt.Println("s[i]: ", string(char))
		priorities = append(priorities, getRunePriority(rune(char)))
	}
	// Merge sort the priorities
	ret := mergeSort(priorities)
	return ret
}

func findCommonElt(a, b []int) int {
	for _, v := range a {
		fmt.Println("v: ", v)
		if binaryChop(v, b) {
			return v
		}
	}
	return -1
}

type Rucksack struct {
	lcomparts  []Compartment
	rcomparts  []Compartment
	priorities []int
}

func initRucksack(l int) Rucksack {
	return Rucksack{
		lcomparts:  make([]Compartment, l),
		rcomparts:  make([]Compartment, l),
		priorities: make([]int, l),
	}
}

// Read in the lines, and split them into two compartments
// Sort the compartments, return the rucksack
func fillRucksack(lines []string, r Rucksack) Rucksack {
	r = initRucksack(len(lines))
	for i, line := range lines {
		a, b := splitLine(line)
		fmt.Println("line a: ", a, "\tline b: ", b)
		compA := initCompartment(len(a))
		compA.priorities = sortStringToIntArray(a, compA.priorities)
		compA.contents = translatePriorities(compA.priorities, "")
		fmt.Println("compA contents: ", compA.contents)
		compB := initCompartment(len(b))
		compB.priorities = sortStringToIntArray(b, compB.priorities)
		compB.contents = translatePriorities(compB.priorities, "")
		fmt.Println("compB contents: ", compB.contents)
		r.lcomparts[i] = compA
		r.rcomparts[i] = compB
		commonElt := findCommonElt(compA.priorities, compB.priorities)
		if commonElt != -1 {
			r.priorities = append(r.priorities, commonElt)
		}
	}
	return r
}

func main() {

	lines := scanner.ScanFileStrings("test.txt", []string{})
	var r Rucksack
	sortedR := fillRucksack(lines, r)
	fmt.Println(sortedR)
}
