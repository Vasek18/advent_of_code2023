package main

import (
	"fmt"
	"os"
	"strings"

	"main.go/utils"
)

type Row struct {
	cells  string
	blocks []int
}

func readInput(input string) []Row {
	var parsedInput []Row

	for _, r := range strings.Split(input, "\n") {
		rParts := strings.Split(r, " ")

		var blocks []int
		rawBlocks := strings.Split(rParts[1], ",")
		for _, v := range rawBlocks {
			blocks = append(blocks, utils.ConvertStringToInt(v))
		}

		parsedInput = append(parsedInput, Row{rParts[0], blocks})
	}

	return parsedInput
}

func countCombinations(row Row, cellIndex int, blockIndex int, currentBlockLength int) int {
	if cellIndex >= len(row.cells) {
		if blockIndex >= len(row.blocks) && currentBlockLength == 0 { // ran out of blocks
			return 1
		}

		if blockIndex == len(row.blocks)-1 && row.blocks[blockIndex] == currentBlockLength { // on the end of last block
			return 1
		}

		return 0 // no more cells
	}

	numberOfCombinations := 0

	if row.cells[cellIndex] == '.' || row.cells[cellIndex] == '?' {
		if currentBlockLength == 0 { // block start
			numberOfCombinations += countCombinations(row, cellIndex+1, blockIndex, 0)
		}
		if currentBlockLength > 0 && blockIndex < len(row.blocks) && row.blocks[blockIndex] == currentBlockLength { // block end
			numberOfCombinations += countCombinations(row, cellIndex+1, blockIndex+1, 0)
		}
	}

	if row.cells[cellIndex] == '#' || row.cells[cellIndex] == '?' {
		numberOfCombinations += countCombinations(row, cellIndex+1, blockIndex, currentBlockLength+1) // still in block
	}

	return numberOfCombinations
}

func solve(rows []Row) int {
	answer := 0

	for _, r := range rows {
		answer += countCombinations(r, 0, 0, 0)
	}

	return answer
}

func main() {
	input, _ := os.ReadFile("./12.1/input.txt")
	rows := readInput(string(input))

	answer := solve(rows)

	fmt.Println(answer)
}
