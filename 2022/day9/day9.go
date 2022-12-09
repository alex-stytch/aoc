package day9

import (
	"aoc/fileutil"
	"fmt"
	"math"
)

type Directions []Direction
type Direction string

const (
	LeftDirection  = "L"
	RightDirection = "R"
	UpDirection    = "U"
	DownDirection  = "D"
)

type Rope []*RopePart

type RopePart struct {
	leaderRopePart  *RopePart // the head will have a nil value
	X, Y            int
	vistedLocations map[int]map[int]bool // x loc to y loc to visited
}

func Part1() {
	fmt.Println(solveTailLocations(2))
}

func Part2() {
	fmt.Println(solveTailLocations(10))
}

func solveTailLocations(length int) int {
	directions := getDirections()
	rope := newRope(length)
	for _, dir := range directions {
		rope.update(dir)
	}
	return rope[length-1].calucateVisitedLocations()
}

func getDirections() Directions {
	var dirs Directions

	input := fileutil.Import("2022/day9/input.txt")
	for _, line := range input {
		var dir Direction
		var distance int
		out, err := fmt.Sscanf(line, "%s %d", &dir, &distance)
		if err != nil || out != 2 {
			panic(line)
		}

		for distance > 0 {
			dirs = append(dirs, dir)
			distance--
		}
	}

	return dirs
}

func newRope(length int) Rope {
	if length < 1 {
		panic("all code assumes length of at least 1")
	}

	rope := append(Rope{}, &RopePart{
		X:               0,
		Y:               0,
		vistedLocations: map[int]map[int]bool{0: {0: true}},
	})

	for index := 1; index < length; index++ {
		rope = append(rope, &RopePart{
			leaderRopePart:  rope[index-1],
			X:               0,
			Y:               0,
			vistedLocations: map[int]map[int]bool{0: {0: true}},
		})
	}

	return rope
}

func (r *Rope) update(dir Direction) {
	(*r)[0].updateHead(dir)
	for i := 1; i < len(*r); i++ {
		(*r)[i].update()
		if (*r)[i].vistedLocations[(*r)[i].X] == nil {
			(*r)[i].vistedLocations[(*r)[i].X] = map[int]bool{}
		}
		(*r)[i].vistedLocations[(*r)[i].X][(*r)[i].Y] = true
	}
}

func (rp *RopePart) updateHead(dir Direction) {
	switch dir {
	case LeftDirection:
		rp.X--
	case RightDirection:
		rp.X++
	case UpDirection:
		rp.Y--
	case DownDirection:
		rp.Y++
	}
}

func (rp *RopePart) update() {
	// don't need to move
	if abs(rp.X-rp.leaderRopePart.X) < 2 && abs(rp.Y-rp.leaderRopePart.Y) < 2 {
		return
	}

	moveX := 0
	if rp.X < rp.leaderRopePart.X {
		moveX = 1
	} else if rp.leaderRopePart.X < rp.X {
		moveX = -1
	}

	moveY := 0
	if rp.Y < rp.leaderRopePart.Y {
		moveY = 1
	} else if rp.leaderRopePart.Y < rp.Y {
		moveY = -1
	}

	rp.X += moveX
	rp.Y += moveY
}

func (rp *RopePart) calucateVisitedLocations() int {
	count := 0

	for i := range rp.vistedLocations {
		for j := range rp.vistedLocations[i] {
			if rp.vistedLocations[i][j] {
				count++
			}
		}
	}

	return count
}

func abs(val int) int {
	return int(math.Abs(float64(val)))
}
