package main

import (
	"fmt"
	"os"
	"strings"
)

func readInput(input string) []string {
	return strings.Split(input, "\n")
}

func solve(rows []string) int {
	answer := 0

	for j := 0; j <= len(rows[0])-1; j++ {
		multiplier := len(rows)
		for i := 0; i <= len(rows)-1; i++ {
			char := rows[i][j]

			if char == 'O' {
				answer += multiplier
				multiplier--
			}

			if char == '#' {
				multiplier = len(rows) - i - 1
			}
		}
	}

	return answer
}

func strCount(input string, char string) int {
	count := 0

	for _, c := range input {
		if string(c) == char {
			count++
		}
	}
	return count
}

func main() {
	input, _ := os.ReadFile("./14.1/input.txt")
	rows := readInput(string(input))

	answer := solve(rows)

	fmt.Println(answer)
}
