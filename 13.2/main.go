package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"main.go/utils"
)

type Pair struct {
	left     int
	right    int
	mistakes int
}

func readInput(input string) [][]string {
	var parsedInput [][]string

	for _, rawPattern := range strings.Split(input, "\n\n") {
		pattern := strings.Split(rawPattern, "\n")
		parsedInput = append(parsedInput, pattern)
	}

	return parsedInput
}

func findReflectionCandidates(s string) []Pair {
	var pairs []Pair

	for i := 0; i < len(s)-1; i++ {
		numberOfErrors := checkPairInString(s, Pair{i, i + 1, 0})

		if numberOfErrors <= 1 {
			pairs = append(pairs, Pair{i, i + 1, numberOfErrors})
		}
	}

	return pairs
}

func checkPairInString(s string, pair Pair) int {
	numberOfErrors := 0

	for j := 0; j < utils.Min(pair.left+1, len(s)-pair.left-1); j++ {
		leftChar := s[pair.left-j]
		rightChar := s[pair.right+j]

		if leftChar != rightChar {
			numberOfErrors++
		}
	}

	return numberOfErrors
}

func getColumn(pattern []string, columnNumber int) []rune {
	var firstChars []rune

	for _, str := range pattern {
		firstChars = append(firstChars, rune(str[columnNumber]))
	}

	return firstChars
}

func filterPairs(pattern []string, pairs []Pair, forRow bool, iterationNumber int) []Pair {
	if len(pairs) == 0 {
		return []Pair{}
	}

	maxIteration := len(pattern)
	if forRow {
		maxIteration = len(pattern[0])
	}
	if iterationNumber >= maxIteration {
		return pairs
	}

	var checkedPairs []Pair
	for _, pair := range pairs {
		var s string
		if forRow {
			s = string(getColumn(pattern, iterationNumber))
		} else {
			s = pattern[iterationNumber]
		}

		numberOfErrors := checkPairInString(s, pair)
		if numberOfErrors+pair.mistakes <= 1 {
			pair.mistakes = numberOfErrors + pair.mistakes
			checkedPairs = append(checkedPairs, pair)
		}
	}

	return filterPairs(pattern, checkedPairs, forRow, iterationNumber+1)
}

func solve(patterns [][]string) int {
	var wg sync.WaitGroup
	var mu sync.Mutex
	ch := make(chan int)

	answer := 0

	for _, pattern := range patterns {
		wg.Add(1)

		go func(pattern []string) {
			defer wg.Done()

			columnReflectionPairs := findReflectionCandidates(pattern[0])
			columnReflectionPairs = filterPairs(pattern, columnReflectionPairs, false, 1)

			if len(columnReflectionPairs) > 0 {
				for _, pair := range columnReflectionPairs {
					if pair.mistakes == 1 {
						mu.Lock()
						answer += pair.right
						mu.Unlock()
					}
				}
			}

			rowReflectionPairs := findReflectionCandidates(string(getColumn(pattern, 0)))
			rowReflectionPairs = filterPairs(pattern, rowReflectionPairs, true, 1)
			if len(rowReflectionPairs) > 0 {
				for _, pair := range rowReflectionPairs {
					if pair.mistakes == 1 {
						mu.Lock()
						answer += pair.right * 100
						mu.Unlock()
					}
				}
			}

			ch <- 1
		}(pattern)
	}

	return answer
}

func main() {
	start := time.Now()

	input, _ := os.ReadFile("./13.2/input.txt")
	patterns := readInput(string(input))

	answer := solve(patterns)

	fmt.Println(answer)

	elapsed := time.Since(start)
	fmt.Printf("Binomial took %s\n", elapsed)
}
