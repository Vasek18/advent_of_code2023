package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("./02.2/input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		gameStringParts := strings.Split(scanner.Text(), ":")

		sum += getGamePower(gameStringParts[1])

	}

	fmt.Println(sum)
}

func getGamePower(setsString string) int {
	counts := map[string]int{
		"green": 0,
		"red":   0,
		"blue":  0,
	}
	sets := strings.Split(setsString, ";")

	for _, set := range sets {
		for _, cubeCount := range strings.Split(set, ",") {
			cubeCountParts := strings.Split(strings.Trim(cubeCount, " "), " ")
			count, _ := strconv.Atoi(cubeCountParts[0])
			color := strings.Trim(cubeCountParts[1], " ")

			if count > counts[color] {
				counts[color] = count
			}
		}
	}

	result := 1
	for _, count := range counts {
		result *= count
	}

	return result
}
