package day15

import (
	"aoc/fileutil"
	"aoc/util"
	"fmt"
)

type Sensor struct {
	sensorX, sensorY               int
	closestBeaconX, closestBeaconY int
}

func Part1() {
	sensors := importSensors()
	fmt.Println(part1(sensors))
}

func Part2() {
	sensors := importSensors()
	fmt.Println(part2(sensors))
}

func importSensors() []*Sensor {
	var sensors []*Sensor
	input := fileutil.Import("2022/day15/input.txt")
	for _, line := range input {
		sensor := &Sensor{}
		fmt.Sscanf(
			line,
			"Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&sensor.sensorX,
			&sensor.sensorY,
			&sensor.closestBeaconX,
			&sensor.closestBeaconY,
		)
		sensors = append(sensors, sensor)
	}
	return sensors
}

func part1(sensors []*Sensor) int {
	magicY := 2_000_000

	spots := map[int]bool{}
	for _, s := range sensors {
		dx := 0
		for manhattanDistance(s.sensorX, s.sensorX+dx, s.sensorY, magicY) <= s.closestBeaconDistance() {
			spots[s.sensorX-dx] = true
			spots[s.sensorX+dx] = true
			dx++
		}
	}

	for _, s := range sensors {
		if s.sensorY == magicY {
			spots[s.sensorX] = false
		}
		if s.closestBeaconY == magicY {
			spots[s.closestBeaconX] = false
		}
	}

	ans := 0
	for _, val := range spots {
		if val {
			ans++
		}
	}

	return ans
}

func part2(sensors []*Sensor) int {
	x := 0
	boundX := 4_000_000
	boundY := 4_000_000
	for x <= boundX {
		y := 0
		for y <= boundY {
			wiggleRoom := -1
			for _, s := range sensors {
				dDist := s.closestBeaconDistance() - manhattanDistance(s.sensorX, x, s.sensorY, y)
				if dDist > wiggleRoom {
					wiggleRoom = dDist
				}
			}
			if wiggleRoom == -1 {
				return x*4000000 + y
			}

			y += wiggleRoom + 1
		}
		x++
	}

	return -1
}

func (s *Sensor) closestBeaconDistance() int {
	return manhattanDistance(s.sensorX, s.closestBeaconX, s.sensorY, s.closestBeaconY)
}

func manhattanDistance(x1, x2, y1, y2 int) int {
	return util.Abs(x1-x2) + util.Abs(y1-y2)
}
