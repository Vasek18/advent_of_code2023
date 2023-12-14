package utils

import (
	"fmt"
	"strconv"
)

func ConvertStringToInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func Max(x, y int) int {
	if x < y {
		return y
	}

	return x
}

func Min(x, y int) int {
	if x > y {
		return y
	}

	return x
}

func substrCount(input string, char string) int {
	count := 0

	for _, c := range input {
		if string(c) == char {
			count++
		}
	}
	return count
}

func printMatrix(platform [][]string) {
	for _, r := range platform {
		fmt.Println(r)
	}
}

func replaceAtIndex(str string, replacement rune, index int) string {
	return str[:index] + string(replacement) + str[index+1:]
}
