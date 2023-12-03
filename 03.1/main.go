package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func isEnginePartSymbol(char rune) bool {
	return string(char) != "." && !unicode.IsDigit(char)
}

func isConnectedToPart(x int, y int, engineSchema []string) bool {
	for y1 := y - 1; y1 <= y+1; y1++ {
		for x1 := x - 1; x1 <= x+1; x1++ {
			if y1 < 0 || x1 < 0 {
				continue
			}
			if y1 >= len(engineSchema) || x1 >= len(engineSchema[y1]) {
				continue
			}

			char := engineSchema[y1][x1]

			if isEnginePartSymbol(rune(char)) {
				return true
			}
		}
	}

	return false
}

func solve(engineSchema []string) int {
	sum := 0

	numStr := ""
	num := 0
	connectedToPart := false

	for y, row := range engineSchema {
		for x, char := range row {
			if unicode.IsDigit(char) {
				numStr += string(char)
				if !connectedToPart {
					connectedToPart = isConnectedToPart(x, y, engineSchema)
				}
			} else {
				if len(numStr) != 0 {
					num, _ = strconv.Atoi(numStr)
					numStr = ""

					if connectedToPart {
						sum += num
						connectedToPart = false
					}
				}
			}
		}
	}

	return sum
}

func main() {
	input, _ := os.ReadFile("./03.1/input.txt")

	totalSum := solve(strings.Split(string(input), "\n"))

	fmt.Println(totalSum)
}
