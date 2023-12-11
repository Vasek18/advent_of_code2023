package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Galaxy struct {
	n   int
	lat int
	lon int
}

func readMap(input string) ([]Galaxy, []int, []int) {
	var ourMap []Galaxy

	inputRows := strings.Split(input, "\n")

	rowsMap := make([]bool, len(inputRows))
	colummnsMap := make([]bool, len(inputRows[0]))

	n := 0
	for lon, r := range inputRows {
		for lat, c := range strings.Split(r, "") {
			if c == "#" {
				n++
				ourMap = append(ourMap, Galaxy{n, lat, lon})

				rowsMap[lon] = true
				colummnsMap[lat] = true
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
	manhattanDistance := int(math.Abs(float64(g1.lat-g2.lat)) + math.Abs(float64(g1.lon-g2.lon)))

	for _, r := range emptyRows {
		if r > g1.lon && r < g2.lon {
			manhattanDistance++
		}
		if r < g1.lon && r > g2.lon {
			manhattanDistance++
		}
	}

	for _, c := range emptyColumns {
		if c > g1.lat && c < g2.lat {
			manhattanDistance++
		}
		if c < g1.lat && c > g2.lat {
			manhattanDistance++
		}
	}

	return manhattanDistance
}

func solve(ourMap []Galaxy, emptyRows []int, emptyColumns []int) int {
	answer := 0

	for i := 0; i < len(ourMap)-1; i++ {
		g1 := ourMap[i]
		for j := i + 1; j < len(ourMap); j++ {
			g2 := ourMap[j]

			answer += calcDistance(g1, g2, emptyRows, emptyColumns)
		}
	}

	return answer
}

func main() {
	input, _ := os.ReadFile("./11.1/input.txt")
	ourMap, emptyRows, emptyColumns := readMap(string(input))

	answer := solve(ourMap, emptyRows, emptyColumns)

	fmt.Println(answer)
}
