package main

import (
	"fmt"
	"os"
	"strings"

	"main.go/utils"
)

type Cell struct {
	s        string
	lat      int
	lon      int
	prevCell *Cell
}

func readMap(input string) [][]string {
	var ourMap [][]string

	for _, r := range strings.Split(input, "\n") {
		ourMap = append(ourMap, strings.Split(r, ""))
	}

	return ourMap
}

func getConnectedCells(cell Cell, ourMap [][]string) []Cell {
	var cells []Cell

	maxLat := len(ourMap[0]) - 1
	for lat := utils.Max(cell.lat-1, 0); lat <= utils.Min(cell.lat+1, maxLat); lat++ {
		for lon := utils.Max(cell.lon-1, 0); lon <= utils.Min(cell.lon+1, maxLat); lon++ {
			char := ourMap[lat][lon]
			newCell := Cell{char, lat, lon, &cell}

			// top
			if lon == cell.lon && lat == cell.lat-1 {
				if char == "|" || char == "F" || char == "7" {
					cells = append(cells, newCell)
				}
			}

			// left
			if lon == cell.lon-1 && lat == cell.lat {
				if char == "-" || char == "L" || char == "F" {
					cells = append(cells, newCell)
				}
			}

			// right
			if lon == cell.lon+1 && lat == cell.lat {
				if char == "-" || char == "J" || char == "7" {
					cells = append(cells, newCell)
				}
			}

			// bottom
			if lon == cell.lon && lat == cell.lat+1 {
				if char == "|" || char == "L" || char == "J" {
					cells = append(cells, newCell)
				}
			}
		}
	}

	return cells
}

func getStartPoint(ourMap [][]string) Cell {
	for lat, _ := range ourMap {
		for lon, char := range ourMap[lat] {
			if char == "S" {
				return Cell{"S", lat, lon, nil}
			}
		}
	}

	panic("No starting point")
}

func getNextCell(cell Cell, ourMap [][]string) Cell {
	var nextLat, nextLon int

	switch cell.s {
	case "|":
		if cell.prevCell.lat < cell.lat {
			nextLat, nextLon = cell.lat+1, cell.lon
		} else {
			nextLat, nextLon = cell.lat-1, cell.lon
		}

		break
	case "-":
		if cell.prevCell.lon < cell.lon {
			nextLat, nextLon = cell.lat, cell.lon+1
		} else {
			nextLat, nextLon = cell.lat, cell.lon-1
		}

		break
	case "L":
		if cell.prevCell.lat < cell.lat {
			nextLat, nextLon = cell.lat, cell.lon+1
		} else {
			nextLat, nextLon = cell.lat-1, cell.lon
		}

		break
	case "J":
		if cell.prevCell.lat < cell.lat {
			nextLat, nextLon = cell.lat, cell.lon-1
		} else {
			nextLat, nextLon = cell.lat-1, cell.lon
		}

		break
	case "7":
		if cell.prevCell.lat == cell.lat {
			nextLat, nextLon = cell.lat+1, cell.lon
		} else {
			nextLat, nextLon = cell.lat, cell.lon-1
		}

		break
	case "F":
		if cell.prevCell.lat == cell.lat {
			nextLat, nextLon = cell.lat+1, cell.lon
		} else {
			nextLat, nextLon = cell.lat, cell.lon+1
		}

		break
	}

	return Cell{ourMap[nextLat][nextLon], nextLat, nextLon, &cell}
}

func (cell *Cell) isVertex() bool {
	if cell.s == "L" {
		return true
	}
	if cell.s == "J" {
		return true
	}
	if cell.s == "7" {
		return true
	}
	if cell.s == "F" {
		return true
	}
	if cell.s == "S" {
		return true
	}

	return false
}

func (cell *Cell) isStartingPoint() bool {
	if cell.s == "S" {
		return true
	}

	return false
}

func getVertices(currentCell Cell, ourMap [][]string, vertices []Cell, loopCount int) ([]Cell, int) {
	loopCount++

	if currentCell.isVertex() {
		vertices = append(vertices, currentCell)
	}

	nextCell := getNextCell(currentCell, ourMap)
	if nextCell.isStartingPoint() {
		return vertices, loopCount
	}

	return getVertices(nextCell, ourMap, vertices, loopCount)
}

func countAreaShoelace(vertices []Cell) int {
	area := 0

	for i := 0; i < len(vertices); i++ {
		curr := vertices[i]

		var next Cell
		if i+1 >= len(vertices) {
			next = vertices[0]
		} else {
			next = vertices[i+1]
		}

		area += curr.lat * next.lon
	}

	for i := 0; i < len(vertices); i++ {
		curr := vertices[i]

		var next Cell
		if i+1 >= len(vertices) {
			next = vertices[0]
		} else {
			next = vertices[i+1]
		}

		area -= curr.lon * next.lat
	}

	return area / 2
}

func countAreaPick(areaGauss int, loopCount int) int {
	return areaGauss - loopCount/2 + 1
}

func solve(ourMap [][]string) int {
	answer := 0

	startingPoint := getStartPoint(ourMap)
	startingCells := getConnectedCells(startingPoint, ourMap)

	vertices := []Cell{startingPoint}
	loopCount := 0

	vertices, loopCount = getVertices(startingCells[1], ourMap, vertices, 1)
	answer = countAreaShoelace(vertices)
	answer = countAreaPick(answer, loopCount)

	return answer
}

func main() {
	input, _ := os.ReadFile("./10.2/input.txt")
	ourMap := readMap(string(input))

	answer := solve(ourMap)

	fmt.Println(answer)
}
