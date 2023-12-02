package main

import (
	"fmt"
	"strconv"
	"strings"
)

var maxCounts = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	games := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
	Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
	Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
	Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

	sum := 0

	for _, gameString := range strings.Split(games, "\n") {
		gameStringParts := strings.Split(gameString, ":")

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
