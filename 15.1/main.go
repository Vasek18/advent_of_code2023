package main

import (
	"fmt"
	"os"
	"strings"
)

func readInput(input string) []string {
	return strings.Split(input, ",")
}

func solve(steps []string) int {
	answer := 0

	for _, step := range steps {
		answer += getHash(step)
	}

	return answer
}

func getHash(s string) int {
	count := 0

	for _, c := range s {
		count = countHshForLetter(c, count)
	}

	return count
}

func countHshForLetter(c rune, current int) int {
	current += int(c)
	current *= 17
	return current % 256
}

func main() {
	input, _ := os.ReadFile("./15.1/input.txt")
	steps := readInput(string(input))

	answer := solve(steps)

	fmt.Println(answer)
}
