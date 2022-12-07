package day7

import (
	"aoc/fileutil"
	"fmt"
)

const (
	part1Max           = 100_000
	availableDiskSpace = 70_000_000
	requiredDiskSpace  = 30_000_000
)

type Directory struct {
	parentDirectory *Directory
	directories     map[string]*Directory
	files           map[string]int
	size            int
}

func NewDirectory(parentDirectory *Directory) *Directory {
	return &Directory{
		parentDirectory: parentDirectory,
		directories:     map[string]*Directory{},
		files:           map[string]int{},
	}
}

func Part1() {
	dir := importDirectory()
	dir.calculateSize()

	fmt.Println(dir.part1())
}

func Part2() {
	dir := importDirectory()
	dir.calculateSize()

	currentDiskSpace := availableDiskSpace - dir.size
	neededDiskSpace := requiredDiskSpace - currentDiskSpace
	fmt.Println(dir.part2(neededDiskSpace))
}

func importDirectory() *Directory {
	terminal := fileutil.Import("2022/day7/input.txt")
	topDir := NewDirectory(nil)
	curDir := topDir

	index := 0
	for index < len(terminal) {
		input := terminal[index]

		if input == "$ cd .." {
			curDir = curDir.parentDirectory
		} else if input == "$ cd /" {
			curDir = topDir
		} else if input[0:4] == "$ cd" {
			curDir = curDir.directories[input[5:]]
		}
		index++

		if input != "$ ls" {
			continue
		}

		// must be `$ ls`
		for index < len(terminal) && terminal[index][0] != '$' {
			if terminal[index][0:3] == "dir" {
				curDir.directories[terminal[index][4:]] = NewDirectory(curDir)
			} else {
				var fileSize int
				var fileName string
				out, err := fmt.Sscanf(terminal[index], "%d %s", &fileSize, &fileName)
				if err != nil || out != 2 {
					panic(terminal[index])
				}
				curDir.files[fileName] = fileSize
			}
			index++
		}
	}

	return topDir
}

func (d *Directory) calculateSize() {
	size := 0

	for _, d := range d.directories {
		d.calculateSize()
		size += d.size
	}

	for _, s := range d.files {
		size += s
	}

	d.size = size
}

func (d *Directory) part1() int {
	ans := 0

	for _, d := range d.directories {
		ans += d.part1()
	}

	if d.size < part1Max {
		ans += d.size
	}

	return ans
}

func (d *Directory) part2(neededDiskSpace int) int {
	bestDirectorySize := d.size

	for _, d := range d.directories {
		res := d.part2(neededDiskSpace)
		if res != -1 && res < bestDirectorySize {
			bestDirectorySize = res
		}
	}

	if bestDirectorySize < neededDiskSpace {
		return -1
	}

	return bestDirectorySize
}
