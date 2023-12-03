package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func getNumberFromCoords(x int, row string) int {
	numStr := string(row[x])

	for x1 := x - 1; x1 >= 0; x1-- {
		isNumber := unicode.IsDigit(rune(row[x1]))

		if isNumber {
			numStr = string(row[x1]) + numStr
		} else {
			break
		}
	}

	for x1 := x + 1; x1 <= len(row)-1; x1++ {
		isNumber := unicode.IsDigit(rune(row[x1]))

		if isNumber {
			numStr += string(row[x1])
		} else {
			break
		}
	}

	number, _ := strconv.Atoi(numStr)

	return number
}

func getConnectedNUmbers(x int, y int, engineSchema []string) []int {
	var numbers []int
	isNumber := false

	for y1 := y - 1; y1 <= y+1; y1++ {
		prevIsNumber := false

		for x1 := x - 1; x1 <= x+1; x1++ {
			if y1 < 0 || x1 < 0 {
				continue
			}
			if y1 >= len(engineSchema) || x1 >= len(engineSchema[y1]) {
				continue
			}

			isNumber = unicode.IsDigit(rune(engineSchema[y1][x1]))

			if prevIsNumber && isNumber {
				continue
			}
			if !isNumber {
				prevIsNumber = false
				continue
			}

			prevIsNumber = true

			number := getNumberFromCoords(x1, engineSchema[y1])
			numbers = append(numbers, number)
		}
	}

	return numbers
}

func solve(engineSchema []string) int {
	sum := 0

	for y, row := range engineSchema {
		for x, char := range row {
			if string(char) != "*" {
				continue
			}

			connectedNumbers := getConnectedNUmbers(x, y, engineSchema)
			if len(connectedNumbers) < 2 {
				continue
			}

			ratio := 1
			for _, n := range connectedNumbers {
				ratio *= n
			}

			sum += ratio
		}
	}

	return sum
}

func main() {
	input, _ := os.ReadFile("./03.2/input.txt")

	totalSum := solve(strings.Split(string(input), "\n"))

	fmt.Println(totalSum)
}
