package main

import (
	"fmt"
	"os"
	"strings"
)

func readInput(input string) [][]string {
	var platform [][]string

	for _, r := range strings.Split(input, "\n") {
		platform = append(platform, strings.Split(r, ""))
	}

	return platform
}

func solve(platform [][]string) int {
	cache := make(map[string]int)

	spinCount := 1000000000

	for i := 0; i < spinCount; i++ {
		platform = aroundTheWorld(platform)

		cacheKey := getCacheKey(platform)
		repeatedIteration, wasHereBefore := cache[cacheKey]

		if wasHereBefore {
			period := i - repeatedIteration
			canSkipSpinsCount := (spinCount - i) / period * period
			i += canSkipSpinsCount

			continue
		}

		cache[cacheKey] = i
	}

	return countLoad(platform)
}

func getCacheKey(platform [][]string) string {
	key := ""

	for _, r := range platform {
		key += strings.Join(r, "")
	}

	return key
}

func aroundTheWorld(platform [][]string) [][]string {
	for i := 0; i < 4; i++ {
		platform = moveStonesNorth(platform)
		platform = rotateMatrixClockwise(platform)
	}
	return platform
}

func rotateMatrixClockwise(platform [][]string) [][]string {
	rowCount := len(platform)
	colCount := len(platform[0])

	rotatedPlatform := make([][]string, colCount)
	for x := 0; x < colCount; x++ {
		rotatedPlatform[x] = make([]string, rowCount)
	}

	for y, r := range platform {
		for x, c := range r {
			rotatedPlatform[x][rowCount-y-1] = c
		}
	}

	return rotatedPlatform
}

func moveStonesNorth(platform [][]string) [][]string {
	for x := 0; x < len(platform[0]); x++ {
		for y := 0; y < len(platform); y++ {
			for y1 := 1; y1 < len(platform); y1++ {
				if platform[y1][x] == "O" && platform[y1-1][x] == "." {
					platform[y1][x] = "."
					platform[y1-1][x] = "O"
				}
			}
		}
	}

	return platform
}

func countLoad(platform [][]string) int {
	answer := 0

	multiplier := 1
	for i := len(platform) - 1; i >= 0; i-- {
		answer += countStonesInRow(platform[i]) * multiplier
		multiplier++
	}

	return answer
}

func countStonesInRow(row []string) int {
	count := 0

	for _, c := range row {
		if c == "O" {
			count++
		}
	}

	return count
}

func main() {
	input, _ := os.ReadFile("./14.2/input.txt")
	rows := readInput(string(input))

	answer := solve(rows)

	fmt.Println(answer)
}
