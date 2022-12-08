package day8

import (
	"aoc/fileutil"
	"fmt"
	"strconv"
)

type Forest [][]Tree

type Tree struct {
	height      int
	visable     bool
	scenicScore int
}

func Part1() {
	trees := importTrees()
	trees.calculateVisibility()
	fmt.Println(trees.totalVisable())
}

func Part2() {
	trees := importTrees()
	trees.calculateScenicScore()
	fmt.Println(trees.highestScenicScore())
}

func importTrees() Forest {
	var trees Forest

	input := fileutil.Import("2022/day8/input.txt")
	for i := 0; i < len(input); i++ {
		var treeRow []Tree
		for j := 0; j < len(input[i]); j++ {
			treeRow = append(treeRow, Tree{height: mustInt(input[i][j])})
		}
		trees = append(trees, treeRow)
	}

	return trees
}

func (f Forest) calculateVisibility() {
	for i := 0; i < len(f); i++ {
		for j := 0; j < len(f[i]); j++ {
			if f.minHeight(i, j) < f[i][j].height {
				f[i][j].visable = true
			}
		}
	}
}

func (f Forest) minHeight(i, j int) int {
	return minInt([]int{
		f.northHeight(i, j),
		f.eastHeight(i, j),
		f.southHeight(i, j),
		f.westHeight(i, j),
	})
}

func (f Forest) northHeight(i, j int) int {
	h := -1
	for k := 0; k < i; k++ {
		if f[k][j].height > h {
			h = f[k][j].height
		}
	}
	return h
}

func (f Forest) eastHeight(i, j int) int {
	h := -1
	for k := j + 1; k < len(f[i]); k++ {
		if f[i][k].height > h {
			h = f[i][k].height
		}
	}
	return h
}

func (f Forest) southHeight(i, j int) int {
	h := -1
	for k := i + 1; k < len(f); k++ {
		if f[k][j].height > h {
			h = f[k][j].height
		}
	}
	return h
}

func (f Forest) westHeight(i, j int) int {
	h := -1
	for k := 0; k < j; k++ {
		if f[i][k].height > h {
			h = f[i][k].height
		}
	}
	return h
}

func (f Forest) totalVisable() int {
	count := 0
	for i := 0; i < len(f); i++ {
		for j := 0; j < len(f[i]); j++ {
			if f[i][j].visable {
				count++
			}
		}
	}
	return count
}

func (f Forest) calculateScenicScore() {
	for i := 0; i < len(f); i++ {
		for j := 0; j < len(f[i]); j++ {
			score := f.northScore(i, j)
			score *= f.eastScore(i, j)
			score *= f.southScore(i, j)
			score *= f.westScore(i, j)
			f[i][j].scenicScore = score
		}
	}
}

func (f Forest) northScore(i, j int) int {
	if i == 0 {
		return 0
	}

	k := i - 1
	for k >= 0 {
		if f[k][j].height >= f[i][j].height {
			return i - k
		}
		k--
	}
	return i - k - 1
}

func (f Forest) eastScore(i, j int) int {
	if j == len(f[i])-1 {
		return 0
	}

	k := j + 1
	for k < len(f[i]) {
		if f[i][k].height >= f[i][j].height {
			return k - j
		}
		k++
	}
	return k - j - 1
}

func (f Forest) southScore(i, j int) int {
	if i == len(f)-1 {
		return 0
	}

	k := i + 1
	for k < len(f) {
		if f[k][j].height >= f[i][j].height {
			return k - i
		}
		k++
	}
	return k - i - 1
}

func (f Forest) westScore(i, j int) int {
	if j == 0 {
		return 0
	}

	k := j - 1
	for k >= 0 {
		if f[i][k].height >= f[i][j].height {
			return j - k
		}
		k--
	}
	return j - k - 1
}

func (f Forest) highestScenicScore() int {
	max := 0
	for i := 0; i < len(f); i++ {
		for j := 0; j < len(f[i]); j++ {
			if f[i][j].scenicScore > max {
				max = f[i][j].scenicScore
			}
		}
	}
	return max
}

func mustInt(c uint8) int {
	i, _ := strconv.Atoi(string(c))
	return i
}

func minInt(ints []int) int {
	var min int
	for i, val := range ints {
		if i == 0 || val < min {
			min = val
		}
	}
	return min
}
