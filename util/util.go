package util

import (
	"math"
	"strconv"
)

func MustInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func Abs(val int) int {
	return int(math.Abs(float64(val)))
}
