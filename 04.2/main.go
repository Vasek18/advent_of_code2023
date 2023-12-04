package main

import (
	"fmt"
	"os"
	"strings"
)

func countIntersection(winningNumbers []string, ticketNumbers []string) int {
	intersections := make([]string, 0)
	for _, n1 := range winningNumbers {
		for _, n2 := range ticketNumbers {
			if n1 == n2 {
				intersections = append(intersections, n2)
			}
		}
	}

	return len(intersections)
}

func countWinnigNumbers(ticketString string) int {
	ticketParts := strings.Split(ticketString, "|")
	winningNumbers := strings.Fields(ticketParts[0])
	ticketNumbers := strings.Fields(ticketParts[1])

	return countIntersection(winningNumbers, ticketNumbers)
}

func solve(inputStrings []string) int {
	count := 0
	copies := make(map[int]int)

	for i := 0; i <= len(inputStrings)-1; i++ {
		ticketString := inputStrings[i]

		count += copies[i] + 1

		ticketParts := strings.Split(ticketString, ":")
		power := countWinnigNumbers(ticketParts[1])

		for c := 0; c <= power-1; c++ {
			copies[i+c+1] = copies[i+c+1] + (copies[i] + 1)
		}
	}

	return count
}

func main() {
	input, _ := os.ReadFile("./04.2/input.txt")

	totalSum := solve(strings.Split(string(input), "\n"))

	fmt.Println(totalSum)
}
