package utils

import (
	"strconv"
)

func ConvertStringToInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
