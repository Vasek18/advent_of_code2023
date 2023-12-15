package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"

	"main.go/utils"
)

type Lens struct {
	label string
	value int
}

type Box struct {
	lenses []Lens
}

func readInput(input string) []string {
	return strings.Split(input, ",")
}

func parseStep(step string) (Lens, int, string) {
	re := regexp.MustCompile(`(\w+)([\-\=])(\d*)`)
	matches := re.FindStringSubmatch(step)

	label := matches[1]
	boxId := getHash(label)
	action := matches[2]
	value := matches[3]

	lens := Lens{label, utils.ConvertStringToInt(value)}

	return lens, boxId, action
}

func (box Box) deleteLens(lensToDelete Lens) Box {
	box.lenses = slices.DeleteFunc(box.lenses, func(lens Lens) bool {
		return lens.label == lensToDelete.label
	})

	return box
}

func (box Box) updateLens(lensToUpdate Lens) Box {
	wasAmongLenses := false

	for i, lens := range box.lenses {
		if lens.label == lensToUpdate.label {
			box.lenses[i].value = lensToUpdate.value

			wasAmongLenses = true
		}
	}

	if !wasAmongLenses {
		box.lenses = append(box.lenses, lensToUpdate)
	}

	return box
}

func solve(steps []string) int {
	answer := 0

	boxes := make([]Box, 256)
	for _, step := range steps {
		lens, boxId, action := parseStep(step)

		box := boxes[boxId]

		if action == "-" {
			boxes[boxId] = box.deleteLens(lens)
			continue
		}

		boxes[boxId] = box.updateLens(lens)
	}

	for ib, box := range boxes {
		for il, lens := range box.lenses {
			answer += (ib + 1) * (il + 1) * lens.value
		}
	}

	return answer
}

func getHash(s string) int {
	count := 0

	for _, c := range s {
		count = countHshForLetter(c, count)
	}

	return count
}

func countHshForLetter(c rune, current int) int {
	current += int(c)
	current *= 17
	return current % 256
}

func main() {
	input, _ := os.ReadFile("./15.2/input.txt")
	steps := readInput(string(input))

	answer := solve(steps)

	fmt.Println(answer)
}
