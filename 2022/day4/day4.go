package day4

import (
	"aoc/fileutil"
	"fmt"
	"strconv"
	"strings"
)

type Assignment struct {
	left1, left2, right1, right2 int
}

func Part1() {
	ans := 0
	assignments := getAssignments()
	for _, a := range assignments {
		if a.HasFullOverlap() {
			ans++
		}
	}
	fmt.Println(ans)
}

func Part2() {
	ans := 0
	assignments := getAssignments()
	for _, a := range assignments {
		if a.HasOverlap() {
			ans++
		}
	}
	fmt.Println(ans)
}

func getAssignments() []Assignment {
	input := fileutil.Import("2022/day4/input.txt")
	var as []Assignment
	for _, line := range input {
		// 94-97,31-95 --> ["94", "97", "31", "95"]
		vals := strings.Split(strings.ReplaceAll(line, "-", ","), ",")

		a := Assignment{
			left1:  MustInt(vals[0]),
			right1: MustInt(vals[1]),
			left2:  MustInt(vals[2]),
			right2: MustInt(vals[3]),
		}
		as = append(as, a)
	}

	return as
}

func (a *Assignment) HasFullOverlap() bool {
	// 1 is fully overlapping 2
	if a.left1 <= a.left2 && a.right1 >= a.right2 {
		return true
	}

	// 2 is fully overlapping 1
	if a.left2 <= a.left1 && a.right2 >= a.right1 {
		return true
	}

	return false
}

func (a *Assignment) HasOverlap() bool {
	// 1 is left of 2
	if a.right1 < a.left2 {
		return false
	}

	// 2 is left 1
	if a.right2 < a.left1 {
		return false
	}

	return true
}

func MustInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
