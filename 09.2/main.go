package main

import (
	"fmt"
	"os"
	"strings"

	"main.go/utils"
)

func getDiffArray(values []int) ([]int, bool) {
	var diffs []int
	allZeros := true

	for i := 1; i < len(values); i++ {
		v := values[i]
		prevV := values[i-1]

		diff := v - prevV

		if diff != 0 {
			allZeros = false
		}

		diffs = append(diffs, diff)
	}

	return diffs, allZeros
}

func getDiffMatrix(values []string) [][]int {
	var matrix [][]int

	var diffs []int
	for _, v := range values {
		diffs = append(diffs, utils.ConvertStringToInt(v))
	}
	matrix = append(matrix, diffs)

	allZeros := false

	for {
		diffs, allZeros = getDiffArray(diffs)

		if allZeros {
			break
		}

		matrix = append(matrix, diffs)
	}

	return matrix
}

func getNextNumber(values []string) int {
	diffMatrix := getDiffMatrix(values)

	diff := diffMatrix[len(diffMatrix)-1][0]
	curentValue := 0

	for i := len(diffMatrix) - 2; i >= 0; i-- {
		curentValue = diff - curentValue
		diff = diffMatrix[i][0]
	}

	return diff - curentValue
}

func solve(rows []string) int {
	sum := 0

	for _, row := range rows {
		sum += getNextNumber(strings.Fields(row))
	}

	return sum
}

func main() {
	input, _ := os.ReadFile("./09.2/input.txt")

	answer := solve(strings.Split(string(input), "\n"))

	fmt.Println(answer)
}
