package main

import (
	"fmt"
	"math"
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
	sum := 0

	for _, ticketString := range inputStrings {
		ticketParts := strings.Split(ticketString, ":")
		power := countWinnigNumbers(ticketParts[1])

		sum += int(math.Pow(2, float64(power-1)))
	}

	return sum
}

func main() {
	input, _ := os.ReadFile("./04.1/input.txt")

	totalSum := solve(strings.Split(string(input), "\n"))

	fmt.Println(totalSum)
}
