package main

import (
	"fmt"
	"os"
	"strings"

	"main.go/utils"
)

func readInput(input string) [][]string {
	var parsedInput [][]string

	for _, rawPattern := range strings.Split(input, "\n\n") {
		pattern := strings.Split(rawPattern, "\n")
		parsedInput = append(parsedInput, pattern)
	}

	return parsedInput
}

func findReflectionCandidates(s string) [][]int {
	var pairs [][]int

	for i := 0; i < len(s)-1; i++ {
		isReflection := checkPairInString(s, []int{i, i + 1})

		if isReflection {
			pairs = append(pairs, []int{i, i + 1})
		}
	}

	return pairs
}

func checkPairInString(s string, pair []int) bool {
	for j := 0; j < utils.Min(pair[0]+1, len(s)-pair[0]-1); j++ {
		leftChar := s[pair[0]-j]
		rightChar := s[pair[1]+j]

		if leftChar != rightChar {
			return false
		}
	}

	return true
}

func getColumn(pattern []string, columnNumber int) []rune {
	var firstChars []rune

	for _, str := range pattern {
		firstChars = append(firstChars, rune(str[columnNumber]))
	}

	return firstChars
}

func filterPairs(pattern []string, pairs [][]int, forRow bool, iterationNumber int) [][]int {
	if len(pairs) == 0 {
		return [][]int{}
	}

	maxIteration := len(pattern)
	if forRow {
		maxIteration = len(pattern[0])
	}
	if iterationNumber >= maxIteration {
		return pairs
	}

	var checkedPairs [][]int
	for _, pair := range pairs {
		var s string
		if forRow {
			s = string(getColumn(pattern, iterationNumber))
		} else {
			s = pattern[iterationNumber]
		}

		if checkPairInString(s, pair) {
			checkedPairs = append(checkedPairs, pair)
		}
	}

	return filterPairs(pattern, checkedPairs, forRow, iterationNumber+1)
}

func solve(patterns [][]string) int {
	answer := 0

	for _, pattern := range patterns {
		columnReflectionPairs := findReflectionCandidates(pattern[0])
		columnReflectionPairs = filterPairs(pattern, columnReflectionPairs, false, 1)

		if len(columnReflectionPairs) > 0 {
			for _, pair := range columnReflectionPairs {
				answer += pair[1]
			}
		}

		rowReflectionPairs := findReflectionCandidates(string(getColumn(pattern, 0)))
		rowReflectionPairs = filterPairs(pattern, rowReflectionPairs, true, 1)
		if len(rowReflectionPairs) > 0 {
			for _, pair := range rowReflectionPairs {
				answer += pair[1] * 100
			}
		}
	}

	return answer
}

func main() {
	input, _ := os.ReadFile("./13.1/input.txt")
	patterns := readInput(string(input))

	answer := solve(patterns)

	fmt.Println(answer)
}
