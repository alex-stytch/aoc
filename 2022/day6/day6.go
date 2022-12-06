package day6

import (
	"aoc/fileutil"
	"fmt"
	"strings"
)

func Part1() {
	fmt.Println(findMarkerIndex(4))
}

func Part2() {
	fmt.Println(findMarkerIndex(14))
}

func findMarkerIndex(markerLength int) int {
	input := fileutil.Import("2022/day6/input.txt")[0]

	var curVal string
	for i := markerLength; i < len(input); i++ {
		if i == markerLength {
			curVal = input[0:markerLength]
		} else {
			curVal = curVal[1:] + string(input[i-1])
		}

		if winner(curVal) {
			return i
		}
	}
	panic("marker not found")
}

func winner(val string) bool {
	for i := range val {
		if strings.Count(val, string(val[i])) > 1 {
			return false
		}
	}

	return true
}
