package day5

import (
	"aoc/fileutil"
	"fmt"
	"strings"
)

type crateStacks []crateStack
type crateStack []string
type move struct {
	quanity, fromRow, toRow int
}

func Part1() {
	cs, ms := importCrateStacksAndMoves()
	for _, m := range ms {
		cs.moveOneAtATime(m)
	}
	fmt.Println(cs.reportTop())
}

func Part2() {
	cs, ms := importCrateStacksAndMoves()
	for _, m := range ms {
		cs.moveAllAtOnce(m)
	}
	fmt.Println(cs.reportTop())
}

func importCrateStacksAndMoves() (crateStacks, []move) {
	input := fileutil.Import("2022/day5/input.txt")

	stacks := make(crateStacks, numCrateStacksFromInput(input))
	var moves []move

	for _, line := range input {
		if stackLine(line) {
			stacks.importStack(line)
		} else if moveLine(line) {
			moves = append(moves, toMove(line))
		}
	}

	return stacks, moves
}

func (cs crateStacks) moveOneAtATime(m move) {
	for i := 0; i < m.quanity; i++ {
		poppedVal := cs[m.fromRow][len(cs[m.fromRow])-1]
		cs[m.fromRow] = cs[m.fromRow][:len(cs[m.fromRow])-1]
		cs[m.toRow] = append(cs[m.toRow], poppedVal)
	}
}

func (cs crateStacks) moveAllAtOnce(m move) {
	poppedVals := cs[m.fromRow][len(cs[m.fromRow])-m.quanity:]
	cs[m.fromRow] = cs[m.fromRow][:len(cs[m.fromRow])-m.quanity]
	cs[m.toRow] = append(cs[m.toRow], poppedVals...)
}

func (cs crateStacks) reportTop() string {
	ans := ""
	for _, cs := range cs {
		ans += cs[len(cs)-1]
	}
	return ans
}

func (cs crateStacks) importStack(line string) {
	for i := 0; i < len(line); i++ {
		if i%4 != 1 {
			continue
		}

		char := string(line[i])
		if char != " " {
			cs[i/4] = append([]string{char}, cs[i/4]...)
		}
	}
}

// Stole from Max's day 4 solution :)
func toMove(line string) move {
	var m move
	out, err := fmt.Sscanf(line, "move %d from %d to %d", &m.quanity, &m.fromRow, &m.toRow)
	if err != nil || out != 3 {
		panic(line)
	}

	// zero index rows
	m.fromRow--
	m.toRow--

	return m
}

func stackLine(line string) bool {
	return strings.Contains(line, "[")
}

func moveLine(line string) bool {
	return strings.Contains(line, "move")
}

func numCrateStacksFromInput(input []string) int {
	return (len(input[0]) + 1) / 4
}
