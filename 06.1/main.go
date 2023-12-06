package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getDistanceFromTime(chargeTime int, totalTime int) int {
	return chargeTime * (totalTime - chargeTime)
}

func solve(inputStrings []string) int {
	ratio := 1

	times := strings.Fields(strings.Split(inputStrings[0], ":")[1])
	distances := strings.Fields(strings.Split(inputStrings[1], ":")[1])

	for i, timeS := range times {
		time, _ := strconv.Atoi(timeS)
		record, _ := strconv.Atoi(distances[i])

		recordsCount := 0

		for t := 0; t <= time; t++ {
			if getDistanceFromTime(t, time) > record {
				recordsCount++
			}
		}

		ratio *= recordsCount
	}

	return ratio
}

func main() {
	input, _ := os.ReadFile("./06.1/input.txt")

	totalSum := solve(strings.Split(string(input), "\n"))

	fmt.Println(totalSum)
}
