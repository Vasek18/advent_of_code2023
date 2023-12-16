package main

import (
	"fmt"
	"os"
	"strings"
)

type Cell struct {
	y        int
	x        int
	prevCell *Cell
}

var ourMap [][]string

var cacheForLoops = map[string]bool{}
var visitedCells = map[string]bool{}

func readMap(input string) [][]string {
	for _, r := range strings.Split(input, "\n") {
		ourMap = append(ourMap, strings.Split(r, ""))
	}

	return ourMap
}

func (cell Cell) isOutOfBounds() bool {
	if cell.y < 0 {
		return true
	}
	if cell.x < 0 {
		return true
	}
	if cell.y > len(ourMap)-1 {
		return true
	}
	if cell.x > len(ourMap[0])-1 {
		return true
	}

	return false
}

func (cell Cell) getCacheKey() string {
	if cell.prevCell == nil {
		return "I've never been here before"
	}

	return fmt.Sprintf("%d_%d_%d_%d", cell.prevCell.y, cell.prevCell.x, cell.y, cell.x)
}

func getNextCells(cell Cell) []Cell {
	var nextCells []Cell

	switch ourMap[cell.y][cell.x] {
	case "|":
		if cell.prevCell.x != cell.x {
			nextCells = append(nextCells, Cell{cell.y - 1, cell.x, &cell})
			nextCells = append(nextCells, Cell{cell.y + 1, cell.x, &cell})

			return nextCells
		} else {
			if cell.prevCell.y < cell.y {
				return []Cell{{cell.y + 1, cell.x, &cell}}
			} else {
				return []Cell{{cell.y - 1, cell.x, &cell}}
			}
		}

		break
	case "-":
		if cell.prevCell.y != cell.y {
			nextCells = append(nextCells, Cell{cell.y, cell.x - 1, &cell})
			nextCells = append(nextCells, Cell{cell.y, cell.x + 1, &cell})

			return nextCells
		} else {
			if cell.prevCell.x < cell.x {
				return []Cell{{cell.y, cell.x + 1, &cell}}
			} else {
				return []Cell{{cell.y, cell.x - 1, &cell}}
			}
		}

		break
	case "/":
		if cell.prevCell.x < cell.x {
			return []Cell{{cell.y - 1, cell.x, &cell}}
		}
		if cell.prevCell.x > cell.x {
			return []Cell{{cell.y + 1, cell.x, &cell}}
		}
		if cell.prevCell.y < cell.y {
			return []Cell{{cell.y, cell.x - 1, &cell}}
		}
		if cell.prevCell.y > cell.y {
			return []Cell{{cell.y, cell.x + 1, &cell}}
		}

		break
	case "\\":
		if cell.prevCell == nil {
			return []Cell{{cell.y + 1, cell.x, &cell}}
		}
		if cell.prevCell.x < cell.x {
			return []Cell{{cell.y + 1, cell.x, &cell}}
		}
		if cell.prevCell.x > cell.x {
			return []Cell{{cell.y - 1, cell.x, &cell}}
		}
		if cell.prevCell.y < cell.y {
			return []Cell{{cell.y, cell.x + 1, &cell}}
		}
		if cell.prevCell.y > cell.y {
			return []Cell{{cell.y, cell.x - 1, &cell}}
		}

		break
	default:
		if cell.prevCell == nil {
			return []Cell{{cell.y, cell.x + 1, &cell}}
		}
		if cell.prevCell.x < cell.x {
			return []Cell{{cell.y, cell.x + 1, &cell}}
		}
		if cell.prevCell.x > cell.x {
			return []Cell{{cell.y, cell.x - 1, &cell}}
		}
		if cell.prevCell.y < cell.y {
			return []Cell{{cell.y + 1, cell.x, &cell}}
		}
		if cell.prevCell.y > cell.y {
			return []Cell{{cell.y - 1, cell.x, &cell}}
		}

		break
	}

	return nextCells
}

func countEnergizedCells(currentCell Cell) int {
	cacheKeyWithDirection := currentCell.getCacheKey()
	if _, exists := cacheForLoops[cacheKeyWithDirection]; exists {
		return 0
	}
	cacheForLoops[cacheKeyWithDirection] = true

	count := 0

	cellNumber := fmt.Sprintf("%d_%d", currentCell.y, currentCell.x)
	if _, visited := visitedCells[cellNumber]; !visited {
		count++
		visitedCells[cellNumber] = true
	}

	nextCells := getNextCells(currentCell)

	// fmt.Println(ourMap[currentCell.y][currentCell.x], cellNumber, nextCells)

	for _, cell := range nextCells {
		if cell.isOutOfBounds() {
			continue
		}

		count += countEnergizedCells(cell)
	}

	return count
}

func solve() int {
	answer := 0

	startingCell := Cell{0, 0, nil}

	answer = countEnergizedCells(startingCell)

	return answer
}

func main() {
	input, _ := os.ReadFile("./16.1/input.txt")
	readMap(string(input))

	answer := solve()

	fmt.Println(answer)
}
