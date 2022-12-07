package scanner

import (
	"bufio"
	"os"
)

func scanFile(input string, lines []rune) []rune {
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
