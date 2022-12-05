tpackage main

import (
	"bufio"
	"fmt"
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

type Results struct {
	op []rune
	me []rune
}

func initResults() Results {
	return Results{[]rune{}, []rune{}}
}

func getResults() Results {
	results := initResults()
	opInput := "op.input.txt"
	opLines := scanFile(opInput, []rune{})
	meInput := "me.input.txt"
	meLines := scanFile(meInput, []rune{})
	if len(opLines) != len(meLines) {
		panic("op and me input files are not the same length")
	}
	results.op = opLines
	results.me = meLines
	return results
}

type Score struct {
	// scores are calculated based on the choice
	// and the outcome of the match
	// {"rock":1, "paper":2, "scissors":3}
	// {"lose":0, "draw":3, "win":6}
	choice  int
	outcome int
}

func initScore() Score {
	return Score{0, 0}
}

func compareHand(oponent rune, player rune) Score {
	// Compare an oponent's hand to player's
	// Will return
	m := map[string]int{
		"rock":     1,
		"paper":    2,
		"scissors": 3,
		"lose":     0,
		"draw":     3,
		"win":      6,
	}

	hand := map[rune]string{
		'A': "rock",
		'B': "paper",
		'C': "scissors",
		'X': "rock",
		'Y': "paper",
		'Z': "scissors",
	}

	ohand := hand[oponent]
	phand := hand[player]

	s := initScore()
	if ohand == "rock" && phand == "scissors" {
		s.choice = m["scissors"]
		s.outcome = m["lose"]
	} else if ohand == "rock" && phand == "paper" {
		s.choice = m["paper"]
		s.outcome = m["win"]
	} else if ohand == "rock" && phand == "rock" {
		s.choice = m["rock"]
		s.outcome = m["draw"]
	} else if ohand == "paper" && phand == "scissors" {
		s.choice = m["scissors"]
		s.outcome = m["win"]
	} else if ohand == "paper" && phand == "paper" {
		s.choice = m["paper"]
		s.outcome = m["draw"]
	} else if ohand == "paper" && phand == "rock" {
		s.choice = m["rock"]
		s.outcome = m["lose"]
	} else if ohand == "scissors" && phand == "scissors" {
		s.choice = m["scissors"]
		s.outcome = m["draw"]
	} else if ohand == "scissors" && phand == "paper" {
		s.choice = m["paper"]
		s.outcome = m["lose"]
	} else if ohand == "scissors" && phand == "rock" {
		s.choice = m["rock"]
		s.outcome = m["win"]
	}
	return s
}

func calculateScore(results Results) int {
	// Calculate the score based on the results
	total := 0
	for i := range results.me {
		s := compareHand(results.op[i], results.me[i])
		total += s.choice + s.outcome
	}
	return total
}

func main() {
	results := getResults()
	total := calculateScore(results)
	fmt.Println(total)
}
