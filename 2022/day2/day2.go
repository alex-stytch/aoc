package day2

import (
	"aoc/fileutil"
	"fmt"
)

type Game struct {
	OpponentPlay Play
	YourPlay     Play
}

type Play int

const (
	PlayRock     Play = 1
	PlayPaper    Play = 2
	PlayScissors Play = 3
)

func Part1() {
	games := getGamesPart1()
	totalScore := 0
	for _, g := range games {
		totalScore += g.score()
	}
	fmt.Println(totalScore)
}

func Part2() {
	games := getGamesPart2()
	totalScore := 0
	for _, g := range games {
		totalScore += g.score()
	}
	fmt.Println(totalScore)
}

func getGamesPart1() []Game {
	input := fileutil.Import("2022/day2/input.txt")
	var games []Game
	for _, line := range input {
		game := Game{
			OpponentPlay: convertColumnToPlay(line[0:1]),
			YourPlay:     convertColumnToPlay(line[2:3]),
		}

		games = append(games, game)
	}

	return games
}

func getGamesPart2() []Game {
	input := fileutil.Import("2022/day2/input.txt")
	var games []Game
	for _, line := range input {
		opponentPlay := convertColumnToPlay(line[0:1])
		yourPlay := resolvePlay(opponentPlay, line[2:3])

		game := Game{
			OpponentPlay: opponentPlay,
			YourPlay:     yourPlay,
		}

		games = append(games, game)
	}

	return games
}

func (g *Game) score() int {
	score := int(g.YourPlay)

	if g.isWin() {
		return score + 6
	} else if g.isDraw() {
		return score + 3
	} else {
		return score
	}
}

func (g *Game) isWin() bool {
	return (int(g.YourPlay)-int(g.OpponentPlay)+3)%3 == 1
}

func (g *Game) isDraw() bool {
	return g.OpponentPlay == g.YourPlay
}

func convertColumnToPlay(val string) Play {
	switch val {
	case "A", "X":
		return PlayRock
	case "B", "Y":
		return PlayPaper
	case "C", "Z":
		return PlayScissors
	default:
		return PlayRock
	}
}

func resolvePlay(opponentPlay Play, val string) Play {
	p := int(opponentPlay)
	switch val {
	case "X": // loss
		return Play((p+1)%3 + 1)
	case "Y": // draw
		return opponentPlay
	case "Z": // win
		return Play((p+3)%3 + 1)
	default:
		return PlayRock
	}
}
