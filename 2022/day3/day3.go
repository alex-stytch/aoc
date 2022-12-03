package day3

import (
	"aoc/fileutil"
	"fmt"
	"strings"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Rucksack struct {
	FirstCompartment  string
	SecondCompartment string
	BothCompartments  string
}

func Part1() {
	rucksacks := getRucksacks()
	priorities := 0
	for _, r := range rucksacks {
		priorities += r.priority()
	}
	fmt.Println(priorities)
}

func Part2() {
	rucksacks := getRucksacks()
	priorities := 0
	for i := 0; i < len(rucksacks); i += 3 {
		for j := 0; j < len(letters); j++ {
			if strings.Contains(rucksacks[i].BothCompartments, string(letters[j])) &&
				strings.Contains(rucksacks[i+1].BothCompartments, string(letters[j])) &&
				strings.Contains(rucksacks[i+2].BothCompartments, string(letters[j])) {
				priorities += j + 1
			}
		}
	}
	fmt.Println(priorities)
}

func getRucksacks() []Rucksack {
	input := fileutil.Import("2022/day3/input.txt")
	var rs []Rucksack
	for _, line := range input {
		r := Rucksack{
			FirstCompartment:  line[0 : len(line)/2],
			SecondCompartment: line[len(line)/2:],
			BothCompartments:  line,
		}

		rs = append(rs, r)
	}

	return rs
}

func (r *Rucksack) priority() int {
	for i := 0; i < len(r.FirstCompartment); i++ {
		if strings.Contains(r.SecondCompartment, string(r.FirstCompartment[i])) {
			return priority(string(r.FirstCompartment[i]))
		}
	}
	panic("priority not found")
}

func priority(char string) int {
	score := rune(strings.ToLower(char)[0]) - 'a' + 1

	if strings.ToUpper(char) == char {
		score += 26
	}

	return int(score)
}
