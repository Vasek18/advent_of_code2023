package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var maxCounts = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	file, _ := os.Open("./02.2/input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		gameStringParts := strings.Split(scanner.Text(), ":")

		if isGamePossible(gameStringParts[1]) {
			gameTitleParts := strings.Split(gameStringParts[0], " ")
			gameId, _ := strconv.Atoi(gameTitleParts[1])
			sum += gameId
		}

	}

	fmt.Println(sum)
}

func isGamePossible(setsString string) bool {
	sets := strings.Split(setsString, ";")

	for _, set := range sets {
		for _, cubeCount := range strings.Split(set, ",") {
			cubeCountParts := strings.Split(strings.Trim(cubeCount, " "), " ")
			count, _ := strconv.Atoi(cubeCountParts[0])
			color := strings.Trim(cubeCountParts[1], " ")

			if count > maxCounts[color] {
				return false
			}
		}
	}

	return true
}
