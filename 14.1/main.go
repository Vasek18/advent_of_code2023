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

	multiplier := 1
	for i := len(rows) - 1; i >= 0; i-- {
		answer += strCount(rows[i], "O") * multiplier
		fmt.Println(strCount(rows[i], "O"), multiplier)
		multiplier++
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
