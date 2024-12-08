package main

import (
	"fmt"
	"os"
	"strings"
)

func readFile(inpath string) ([][]rune, error) {
	bytes, err := os.ReadFile(inpath)
	if err != nil {
		return nil, err
	}
	var themap [][]rune
	for _, line := range strings.Split(strings.Trim(string(bytes[:]), "\n"), "\n") {
		var row []rune
		for _, c := range line {
			row = append(row, c)
		}
		themap = append(themap, row)
	}
	return themap, nil
}

func copymap(origmap [][]rune) [][]rune {
	var themap [][]rune
	for _, origrow := range origmap {
		row := append(make([]rune, 0, len(origrow)), origrow...)
		themap = append(themap, row)
	}
	return themap
}

func getmap(themap [][]rune, x, y int) rune {
	if x < 0 || y < 0 || x >= len(themap[0]) || y >= len(themap) {
		return '_'
	}
	return themap[y][x]
}

type guard struct {
	x, y int
	dir  int
}

func (guard *guard) move(themap [][]rune) {
	// turn while blocked
	for {
		var ahead rune
		switch guard.dir {
		case 0:
			ahead = getmap(themap, guard.x, guard.y-1)
		case 1:
			ahead = getmap(themap, guard.x+1, guard.y)
		case 2:
			ahead = getmap(themap, guard.x, guard.y+1)
		case 3:
			ahead = getmap(themap, guard.x-1, guard.y)
		}
		if ahead != '#' {
			break
		}
		guard.dir = (guard.dir + 1) % 4
	}
	// walk
	switch guard.dir {
	case 0:
		guard.y--
	case 1:
		guard.x++
	case 2:
		guard.y++
	case 3:
		guard.x--
	}
}

func part1(origmap [][]rune) int {
	themap := copymap(origmap)

	var guard guard
outer:
	for y, row := range themap {
		for x, c := range row {
			if c == '^' {
				guard.x = x
				guard.y = y
				guard.dir = 0
				themap[y][x] = '1'
				break outer
			}
		}
	}

	visited := 1
	for {
		guard.move(themap)
		if guard.x < 0 || guard.y < 0 || guard.x >= len(themap[0]) || guard.y >= len(themap) {
			break
		}
		if themap[guard.y][guard.x] == '.' {
			themap[guard.y][guard.x] = '0'
			visited++
		}
		themap[guard.y][guard.x]++
		if themap[guard.y][guard.x] == '4' {
			// visited over 4 times -> loop
			return -1
		}
	}
	return visited
}

func part2(themap [][]rune) int {
	numObstacles := 0
	for y, row := range themap {
		for x, c := range row {
			if c == '.' {
				themap[y][x] = '#'
				if part1(themap) == -1 {
					numObstacles++
				}
				themap[y][x] = '.'
			}
		}
	}
	return numObstacles
}

func main() {
	themap, err := readFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("p1:", part1(themap))
	fmt.Println("p2:", part2(themap))
}
