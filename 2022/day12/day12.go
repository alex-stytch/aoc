package day12

import (
	"aoc/fileutil"
	"fmt"
	"math"
)

type grid struct {
	elevation [][]int
	minSteps  [][]int

	startingX, startingY int
	endingX, endingY     int
}

func Part1() {
	g := importGrid()
	g.startCrawl()
	fmt.Println(g.solveSteps())
}

func Part2() {
	g := importGrid()
	g.startCrawl()
	fmt.Println(g.solveTrail())
}

func importGrid() *grid {
	g := &grid{}
	input := fileutil.Import("2022/day12/input.txt")

	for i, line := range input {
		g.elevation = append(g.elevation, []int{})
		g.minSteps = append(g.minSteps, []int{})
		for j, loc := range line {
			g.elevation[i] = append(g.elevation[i], toElevation(loc))
			g.minSteps[i] = append(g.minSteps[i], math.MaxInt)

			if loc == 'S' {
				g.startingX = j
				g.startingY = i
			}
			if loc == 'E' {
				g.endingX = j
				g.endingY = i
			}
		}
	}

	return g
}

func (g *grid) startCrawl() {
	g.minSteps[g.endingY][g.endingX] = 0

	g.crawl(g.endingX-1, g.endingY, 1, 25)
	g.crawl(g.endingX+1, g.endingY, 1, 25)
	g.crawl(g.endingX, g.endingY-1, 1, 25)
	g.crawl(g.endingX, g.endingY+1, 1, 25)
}

func (g *grid) crawl(x, y, minSteps, elevation int) {
	if y < 0 || y >= len(g.elevation) || x < 0 || x >= len(g.elevation[y]) {
		return
	}

	if g.minSteps[y][x] <= minSteps {
		return
	}

	if g.elevation[y][x] < elevation-1 {
		return
	}

	g.minSteps[y][x] = minSteps

	g.crawl(x-1, y, minSteps+1, g.elevation[y][x])
	g.crawl(x+1, y, minSteps+1, g.elevation[y][x])
	g.crawl(x, y-1, minSteps+1, g.elevation[y][x])
	g.crawl(x, y+1, minSteps+1, g.elevation[y][x])
}

func (g *grid) solveSteps() int {
	return g.minSteps[g.startingY][g.startingX]
}

func (g *grid) solveTrail() int {
	min := math.MaxInt

	for i := range g.elevation {
		for j := range g.elevation[i] {
			if g.elevation[i][j] == 0 && min > g.minSteps[i][j] {
				min = g.minSteps[i][j]
			}
		}
	}

	return min
}

func toElevation(i int32) int {
	switch i {
	case 'S':
		return 0
	case 'E':
		return 25
	default:
		return 25 - int('z'-i)
	}
}
