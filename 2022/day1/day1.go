package day1

import (
	"aoc/fileutil"
	"fmt"
)

func Day1() {
	ans := fileutil.Import("2022/day1/input.txt")
	if ans[0] == "is this thing on" {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
