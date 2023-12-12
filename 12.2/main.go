package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"main.go/utils"
)

type Row struct {
	cells  string
	blocks []int
}

var cache = map[string]int{}

func readInput(input string) []Row {
	var parsedInput []Row

	for _, r := range strings.Split(input, "\n") {
		rParts := strings.Split(r, " ")

		expandedCells := strings.Join([]string{rParts[0], rParts[0], rParts[0], rParts[0], rParts[0]}, "?")

		var blocks []int
		expandedBlocks := strings.Join([]string{rParts[1], rParts[1], rParts[1], rParts[1], rParts[1]}, ",")
		rawBlocks := strings.Split(expandedBlocks, ",")
		for _, v := range rawBlocks {
			blocks = append(blocks, utils.ConvertStringToInt(v))
		}

		parsedInput = append(parsedInput, Row{expandedCells, blocks})
	}

	return parsedInput
}

func countCombinations(row Row, cellIndex int, blockIndex int, currentBlockLength int) int {
	cacheKey := fmt.Sprintf("%s_%s_%s", strconv.Itoa(cellIndex), strconv.Itoa(blockIndex), strconv.Itoa(currentBlockLength))
	if value, exists := cache[cacheKey]; exists {
		return value
	}

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

	cache[cacheKey] = numberOfCombinations

	return numberOfCombinations
}

func solve(rows []Row) int {
	answer := 0

	for _, r := range rows {
		cache = map[string]int{}

		answer += countCombinations(r, 0, 0, 0)
	}

	return answer
}

func main() {
	input, _ := os.ReadFile("./12.2/input.txt")
	rows := readInput(string(input))

	answer := solve(rows)

	fmt.Println(answer)
}
