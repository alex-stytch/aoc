package day1

import (
	"aoc/fileutil"
	"fmt"
	"sort"
	"strconv"
)

func Part1() {
	fmt.Println(getMax(getCaloriesPerElf()))
}

func Part2() {
	max3 := getMaxN(getCaloriesPerElf(), 3)
	fmt.Println(max3[0] + max3[1] + max3[2])
}

func getCaloriesPerElf() []int {
	input := fileutil.Import("2022/day1/input.txt")
	var caloriesPerElf []int
	var elfCalories int
	for _, line := range input {
		if line == "" {
			caloriesPerElf = append(caloriesPerElf, elfCalories)
			elfCalories = 0
		} else {
			lineInt, _ := strconv.Atoi(line) // Assume no issues converting
			elfCalories += lineInt
		}
	}

	caloriesPerElf = append(caloriesPerElf, elfCalories)
	return caloriesPerElf
}

// Assumes length of at least 1
// O(N)
func getMax(list []int) int {
	max := list[0]
	for i := 0; i < len(list); i++ {
		if max < list[i] {
			max = list[i]
		}
	}

	return max
}

// Assumes length of at least n
// O(NlogN)
func getMaxN(list []int, n int) []int {
	copyList := make([]int, len(list))
	copy(copyList, list)
	sort.Ints(copyList)

	return copyList[len(copyList)-n:]
}
