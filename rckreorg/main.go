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

// Iterate and find the common character pairs
func findCommonRune(a, b string) (int, int) {
	for i, r := range a {
		if r == rune(b[i]) {
			return i, getRunePriority(r)
		}
	}
	return -1, -1
}

func sortString(s string) string {
	// Convert the string to a rune slice
	r := []rune(s)

	// Sort the slice
	for i := 0; i < len(r)-1; i++ {
		for j := i + 1; j < len(r); j++ {
			if getRunePriority(r[i]) > getRunePriority(r[j]) {
				r[i], r[j] = r[j], r[i]
			}
		}
	}

	// Convert the slice back to a string
	return string(r)
}

// Read in the lines, and split them into two compartments
// Sort the compartments, return the rucksack
func fillRucksack(lines []string) [][2]string {
	r := make([][2]string, len(lines))
	for i, line := range lines {
		a, b := splitLine(line)
		r[i][0] = sortString(a)
		r[i][1] = sortString(b)
	}
	return r
}

func main() {

	lines := scanner.ScanFileStrings("test.txt", []string{})
	r := fillRucksack(lines)
	fmt.Println(r)
}
