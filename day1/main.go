package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func scanFile(filename string, lines []string) []string {
	// Scans a file, strips newlines and stores lines in slice
	// Returns slice of lines
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		read_line := s.Text()
		lines = append(lines, read_line)
	}
	return lines
}

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

func infoElfMostSnacks(elves []Elf) {
	elfi, largestTotal := findLargestTotal(elves)
	fattestElf := elves[elfi]
	fmt.Println("Largest total: ", largestTotal)
	fmt.Println("Items: ", fattestElf.items)
	fmt.Println("Largest total from elf: ", fattestElf.total)

}

func main() {
	lines := scanFile("input.txt", []string{})
	elves := makeElves(lines, []Elf{})
	// Day 1
	infoElfMostSnacks(elves)
}
