package day11

import (
	"aoc/fileutil"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	items         []int
	operation     operation
	operationVal  int
	divisibleTest int
	trueMonkey    int
	falseMonkey   int

	inspectCount int
}

type operation string

const (
	operationPlus     = "+"
	operationMultiply = "*"
	operationSquare   = "^"
)

func Part1() {
	monkeys := importMonkeys()
	runRounds(monkeys, 20, true)
	fmt.Println(calculateMonkeyBusiness(monkeys))
}

func Part2() {
	monkeys := importMonkeys()
	runRounds(monkeys, 10000, false)
	fmt.Println(calculateMonkeyBusiness(monkeys))
}

func importMonkeys() []*monkey {
	var monkeys []*monkey
	input := fileutil.Import("2022/day11/input.txt")
	curIndex := 0
	for curIndex < len(input) {
		curMonkey := &monkey{}

		itemInput := strings.Split(input[curIndex+1], "Starting items: ")[1]
		itemSplit := strings.Split(itemInput, ", ")
		for i := range itemSplit {
			curMonkey.items = append(curMonkey.items, MustInt(itemSplit[i]))
		}

		op := strings.Split(input[curIndex+2], "  Operation: new = old ")[1]
		if op == "* old" {
			curMonkey.operation = operationSquare
		} else {
			if strings.Contains(op, "*") {
				curMonkey.operation = operationMultiply
			} else {
				curMonkey.operation = operationPlus
			}
			curMonkey.operationVal = MustInt(op[2:])
		}

		fmt.Sscanf(input[curIndex+3], "  Test: divisible by %d", &curMonkey.divisibleTest)
		fmt.Sscanf(input[curIndex+4], "    If true: throw to monkey %d", &curMonkey.trueMonkey)
		fmt.Sscanf(input[curIndex+5], "    If false: throw to monkey %d", &curMonkey.falseMonkey)

		monkeys = append(monkeys, curMonkey)
		curIndex += 7
	}

	return monkeys
}

func runRounds(monkeys []*monkey, rounds int, divide bool) {
	for rounds > 0 {
		runRound(monkeys, divide)
		rounds--
	}
}

func runRound(monkeys []*monkey, divide bool) {
	for i := range monkeys {
		runTurn(monkeys, i, divide)
	}
}

func runTurn(monkeys []*monkey, monkeyIndex int, divide bool) {
	curMonkey := monkeys[monkeyIndex]
	for len(curMonkey.items) > 0 {
		item := curMonkey.items[0]
		curMonkey.items = curMonkey.items[1:]

		switch curMonkey.operation {
		case operationPlus:
			item += curMonkey.operationVal
		case operationMultiply:
			item *= curMonkey.operationVal
		case operationSquare:
			item *= item
		}

		if divide {
			item /= 3
		}

		item %= leastCommonMultiple(monkeys)

		toMonkeyIndex := curMonkey.falseMonkey
		if item%curMonkey.divisibleTest == 0 {
			toMonkeyIndex = curMonkey.trueMonkey
		}

		monkeys[toMonkeyIndex].items = append(monkeys[toMonkeyIndex].items, item)

		curMonkey.inspectCount++
	}
}

func calculateMonkeyBusiness(monkeys []*monkey) int {
	inspectCount := make([]int, len(monkeys))
	for i := range monkeys {
		inspectCount[i] = monkeys[i].inspectCount
	}
	sort.Ints(inspectCount)
	return inspectCount[len(inspectCount)-2] * inspectCount[len(inspectCount)-1]
}

func leastCommonMultiple(monkeys []*monkey) int {
	lcm := 1
	for i := range monkeys {
		lcm *= monkeys[i].divisibleTest
	}
	return lcm
}

func MustInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
