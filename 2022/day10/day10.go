package day10

import (
	"aoc/fileutil"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Instructions []Instruction

type Instruction struct {
	typ   InstructionType
	value int
}

type InstructionType string

const (
	noopInstructionType = "noop"
	addxInstructionType = "addx"
)

func Part1() {
	instructions := getInstructions()
	cyclesPerTime := instructions.calculateCycles()
	fmt.Println(getStrength(cyclesPerTime))
}

func Part2() {
	instructions := getInstructions()
	cyclesPerTime := instructions.calculateCycles()
	printCycles(cyclesPerTime)
}

func getInstructions() Instructions {
	var instructions Instructions

	input := fileutil.Import("2022/day10/input.txt")
	for _, line := range input {
		if strings.HasPrefix(line, "noop") {
			instructions = append(instructions, Instruction{
				typ: noopInstructionType,
			})
		} else {
			instructions = append(instructions, Instruction{
				typ:   addxInstructionType,
				value: mustInt(line[5:]),
			})
		}
	}

	return instructions
}

func (i Instructions) calculateCycles() []int {
	cycles := []int{1}

	for _, inst := range i {
		cycles = append(cycles, cycles[len(cycles)-1])
		if inst.typ == addxInstructionType {
			cycles = append(cycles, cycles[len(cycles)-1]+inst.value)
		}
	}

	return cycles
}

func getStrength(input []int) int {
	s := 0
	for _, l := range []int{20, 60, 100, 140, 180, 220} {
		s += l * input[l-1]
	}
	return s
}

func printCycles(input []int) {
	for i := 0; i < 240; i++ {
		if i%40 == 0 {
			fmt.Println()
		}

		if abs(input[i]-(i%40)) <= 1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

func mustInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func abs(val int) int {
	return int(math.Abs(float64(val)))
}
