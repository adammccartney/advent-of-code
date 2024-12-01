package main

import (
	"bufio"
	"log"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	// read from stdin 
	scanner := bufio.NewScanner(os.Stdin)

	var lhs []string
	nocRhs := make(map[string]int)
	for scanner.Scan() {
		// split here, make two slices
		parts := strings.Split(scanner.Text(), " ")
		l := len(parts)
		lhs = append(lhs, parts[0])
		k := parts[l-1]
		nocRhs[k] += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("something went wrong", err)
	}


	// iterate over first array and check for similarity score in second
	// similarity score: lhs * nocrhs

	sum := 0
	for i := range lhs {
		key := lhs[i]
		n, _ := strconv.Atoi(key)
		occursRhs := nocRhs[key]
		simscore := n * occursRhs
		sum+=simscore
	}
	fmt.Printf("Similarity score sum: %d\n", sum)
}
