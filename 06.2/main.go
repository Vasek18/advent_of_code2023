package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getDistanceFromTime(chargeTime int, totalTime int) int {
	return chargeTime * (totalTime - chargeTime)
}

func getNumberFromInputString(input string) int {
	reg, _ := regexp.Compile("[^0-9]+")
	number, _ := strconv.Atoi(reg.ReplaceAllString(input, ""))
	return number
}

func solve(inputStrings []string) int {
	time := getNumberFromInputString(inputStrings[0])
	record := getNumberFromInputString(inputStrings[1])

	recordsCount := 0

	for t := 0; t <= time; t++ {
		if getDistanceFromTime(t, time) > record {
			recordsCount++
		}
	}

	return recordsCount
}

func main() {
	input, _ := os.ReadFile("./06.2/input.txt")

	totalSum := solve(strings.Split(string(input), "\n"))

	fmt.Println(totalSum)
}
