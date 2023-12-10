package utils

import (
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
