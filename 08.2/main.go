package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Node struct {
	name string
	l    string
	r    string
}

func parseInput(input string) (string, map[string]Node) {
	instructions := ""
	nodes := make(map[string]Node)

	for i, s := range strings.Split(string(input), "\n") {
		if i == 0 {
			instructions = s
		}

		if i < 2 {
			continue
		}

		r, _ := regexp.Compile(`^(?P<nodeName>\w+)\s\=\s\((?P<leftNode>\w+),\s(?P<rightNode>\w+)\)`)
		m := r.FindStringSubmatch(s)
		result := make(map[string]string)
		for i, name := range r.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = m[i]
			}
		}

		nodes[result["nodeName"]] = Node{result["nodeName"], result["leftNode"], result["rightNode"]}
	}

	return instructions, nodes
}

func solve(instructions string, nodes map[string]Node) int {
	var cycleLengths []int

	for _, node := range nodes {
		if node.name[2:] == "A" {
			cycleLengths = append(cycleLengths, countSteps(instructions, nodes, node, 0))
		}
	}

	return minProductOfFactors(cycleLengths)
}

func countSteps(instructions string, nodes map[string]Node, currentNode Node, stepsCount int) int {
	for _, step := range instructions {
		stepsCount++

		if step == 'L' {
			currentNode = nodes[currentNode.l]
		}
		if step == 'R' {
			currentNode = nodes[currentNode.r]
		}

		if currentNode.name[2:] == "Z" {
			return stepsCount
		}
	}

	return countSteps(instructions, nodes, currentNode, stepsCount)
}

func getFactors(num int) []int {
	factors := []int{}
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			factors = append(factors, i)
		}
	}

	return factors
}

func minProductOfFactors(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	result := nums[0]

	for _, num := range nums[1:] {
		factors := getFactors(num)

		for _, factor := range factors {
			if result%factor == 0 {
				result /= factor
			}
		}

		result *= num
	}

	return result
}

func main() {
	input, _ := os.ReadFile("./08.2/input.txt")

	instructions, nodes := parseInput(string(input))

	answer := solve(instructions, nodes)

	fmt.Println(answer)
}

func convertToInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
