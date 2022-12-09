package scanner

import (
	"bufio"
	"log"
	"os"
)

func ScanFileRunes(input string, lines []rune) []rune {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		read_line := s.Text()
		asrune := []rune(read_line)
		lines = append(lines, asrune[0])
	}
	return lines
}

func ScanFileStrings(filename string, lines []string) []string {
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
