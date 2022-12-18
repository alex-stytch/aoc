package day14

import (
	"aoc/fileutil"
	"aoc/util"
	"fmt"
	"strings"
)

type Board map[int]map[int]Area
type Area string

const (
	AreaRock = "#"
	AreaSand = "o"
	AreaAir  = "."

	AbyssLevel = 1_000
)

func Part1() {
	board := importBoard()
	board.dropSand(true)
	fmt.Println(board.totalSand())
	//board.print(490, 510, 0, 10)
}

func Part2() {
	board := importBoard()
	board.addFloor()
	board.dropSand(false)
	fmt.Println(board.totalSand())
	//board.print(490, 510, 0, 13)
}

func importBoard() *Board {
	board := &Board{}
	input := fileutil.Import("2022/day14/input.txt")
	for _, line := range input {
		line = strings.ReplaceAll(line, " -> ", ",")
		cors := strings.Split(line, ",")
		for len(cors) >= 4 {
			board.addRockLine(
				util.MustInt(cors[0]),
				util.MustInt(cors[1]),
				util.MustInt(cors[2]),
				util.MustInt(cors[3]),
			)
			cors = cors[2:]
		}
	}

	return board
}

func (b *Board) addRockLine(x1, y1, x2, y2 int) {
	dx := 0
	if x1 > x2 {
		dx = -1
	} else if x1 < x2 {
		dx = 1
	}
	dy := 0
	if y1 > y2 {
		dy = -1
	} else if y1 < y2 {
		dy = 1
	}

	b.placeArea(x1, y1, AreaRock)
	for x1 != x2 || y1 != y2 {
		x1 += dx
		y1 += dy
		b.placeArea(x1, y1, AreaRock)
	}
}

func (b *Board) addFloor() {
	maxY := 0
	for y := range *b {
		if y > maxY {
			maxY = y
		}
	}

	(*b)[maxY+2] = map[int]Area{}
	x := 0
	for x < 1_000 {
		(*b)[maxY+2][x] = AreaRock
		x++
	}
}

func (b *Board) dropSand(abyss bool) {
	if b.dropSingleSand(500, 0, abyss) {
		b.dropSand(abyss)
	}
}

func (b *Board) dropSingleSand(x, y int, abyss bool) bool {
	if abyss && y >= AbyssLevel {
		return false
	}

	if b.freeArea(x, y+1) {
		return b.dropSingleSand(x, y+1, abyss)
	}
	if b.freeArea(x-1, y+1) {
		return b.dropSingleSand(x-1, y+1, abyss)
	}
	if b.freeArea(x+1, y+1) {
		return b.dropSingleSand(x+1, y+1, abyss)
	}

	b.placeArea(x, y, AreaSand)

	if !abyss && x == 500 && y == 0 {
		return false
	}
	return true
}

func (b Board) freeArea(x, y int) bool {
	if b[y] == nil {
		return true
	}
	return b[y][x] == ""
}

func (b *Board) placeArea(x, y int, area Area) {
	if (*b)[y] == nil {
		(*b)[y] = map[int]Area{}
	}
	(*b)[y][x] = area
}

func (b Board) totalSand() int {
	count := 0
	for y := range b {
		for x := range b[y] {
			if b[y][x] == AreaSand {
				count++
			}
		}
	}

	return count
}

func (b Board) print(x1, x2, y1, y2 int) {
	for y1 <= y2 {
		xo := x1
		for xo <= x2 {
			if b[y1] == nil || b[y1][xo] == "" {
				fmt.Print(AreaAir)
			} else {
				fmt.Print(b[y1][xo])
			}
			xo++
		}
		fmt.Println()
		y1++
	}
}
