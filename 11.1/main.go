package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Galaxy struct {
	lat int
	lon int
}

func readMap(input string) ([]Galaxy, []int, []int) {
	var ourMap []Galaxy

	inputRows := strings.Split(input, "\n")

	rowsMap := make([]bool, len(inputRows))
	colummnsMap := make([]bool, len(inputRows[0]))

	for lat, r := range inputRows {
		for lon, c := range strings.Split(r, "") {
			if c == "#" {
				ourMap = append(ourMap, Galaxy{lat, lon})

				rowsMap[lat] = true
				colummnsMap[lon] = true
			}
		}
	}

	var emptyRows []int
	for i, r := range rowsMap {
		if r == false {
			emptyRows = append(emptyRows, i)
		}
	}

	var emptyColumns []int
	for i, c := range colummnsMap {
		if c == false {
			emptyColumns = append(emptyColumns, i)
		}
	}

	return ourMap, emptyRows, emptyColumns
}

func calcDistance(g1 Galaxy, g2 Galaxy, emptyRows []int, emptyColumns []int) int {
	manhattanDistance := int(math.Abs(float64(g2.lat-g1.lat)) + math.Abs(float64(g2.lon-g1.lon)))

	return manhattanDistance
}

func solve(ourMap []Galaxy, emptyRows []int, emptyColumns []int) int {
	answer := 0

	for _, g1 := range ourMap {
		for _, g2 := range ourMap {
			if g1.lat == g2.lat && g1.lon == g2.lon {
				continue
			}

			answer += calcDistance(g1, g2, emptyRows, emptyColumns)
		}
	}

	return answer
}

func main() {
	input, _ := os.ReadFile("./11.1/input.txt")
	ourMap, emptyRows, emptyColumns := readMap(string(input))

	fmt.Println(ourMap, emptyRows, emptyColumns)

	answer := solve(ourMap, emptyRows, emptyColumns)

	fmt.Println(answer)
}
