package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/adammccartney/advent-of-code/2022/lib/scanner"
)

func isnumber(line string) bool {
	ret := false
	for _, char := range line {
		ret = char >= '0' && char <= '9'
	}
	return ret
}

func convertStringToInt(input string) int {
	number, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	return number
}

func makeElves(lines []string, elves []Elf) []Elf {
	elf := initElf()
	for i, line := range lines {
		if line == "" || i == (len(lines)-1) { // check for empty line or last line
			elf.total = sumElfCalories(elf)
			elves = appendElf(elves, elf)
			elf = initElf()
			continue
		}
		if isnumber(line) {
			calories := convertStringToInt(line)
			appendToItems(&elf, calories)
		}
	}
	return elves
}

func findLargestTotal(elves []Elf) (int, int) {
	largest := 0
	elfi := 0
	for i, elf := range elves {
		if elf.total > largest {
			elfi = i
			largest = elf.total
		}
	}
	return elfi, largest
}

type Elf struct {
	items []int
	total int
}

func appendToItems(elf *Elf, item int) {
	elf.items = append(elf.items, item)
}

func appendElf(elves []Elf, elf Elf) []Elf {
	elves = append(elves, elf)
	return elves
}

func initElf() Elf {
	return Elf{[]int{}, 0}
}

func sumElfCalories(elf Elf) int {
	return sumTotalInArray(elf.items)
}

func mapSumTotalCalories(elves []Elf) []Elf {
	for i, elf := range elves {
		elves[i].total = sumElfCalories(elf)
	}
	return elves
}

func sumTotalInArray(array []int) int {
	total := 0
	for _, item := range array {
		total += item
	}
	return total
}

func quickSort(array []Elf) []Elf {
	if len(array) <= 1 {
		return array
	}
	pivot := array[0]
	var left, right []Elf
	for _, item := range array[1:] {
		if item.total < pivot.total {
			left = append(left, item)
		} else {
			right = append(right, item)
		}
	}
	return append(append(quickSort(left), pivot), quickSort(right)...)
}

// Day 1 info
func infoElfMostSnacks(elves []Elf) {
	elfi, largestTotal := findLargestTotal(elves)
	fattestElf := elves[elfi]
	fmt.Println("Largest total: ", largestTotal)
	fmt.Println("Items: ", fattestElf.items)
	fmt.Println("Largest total from elf: ", fattestElf.total)
}

// Day 2 info
func infoTopThreeElves(elves []Elf) {
	sortedElves := quickSort(elves)
	total := 0
	for i := len(sortedElves) - 1; i >= len(sortedElves)-3; i-- {
		fmt.Println(sortedElves[i].total)
		total += sortedElves[i].total
	}
	fmt.Println("Total of top three elves: ", total)
}

func main() {
	lines := scanner.ScanFileStrings("input.txt", []string{})
	elves := makeElves(lines, []Elf{})
	// Day 1
	infoElfMostSnacks(elves)
	// Day 2
	infoTopThreeElves(elves)
}
