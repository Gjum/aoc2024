package main

import (
	"fmt"
	"os"
	"strings"
)

func readFile(inpath string) ([]antenna, xy, error) {
	bytes, err := os.ReadFile(inpath)
	if err != nil {
		return nil, xy{0, 0}, err
	}
	lines := strings.Split(strings.Trim(string(bytes[:]), "\n"), "\n")
	var antennas []antenna
	for y, line := range lines {
		for x, c := range line {
			if c != '.' {
				antennas = append(antennas, antenna{x, y, c})
			}
		}
	}
	return antennas, xy{len(lines[0]), len(lines)}, nil
}

type antenna struct {
	x, y int
	freq rune
}

type xy struct {
	x, y int
}

func part1(input []antenna, size xy) int {
	antinodes := make(map[xy]bool)
	for _, a := range input {
		for _, b := range input {
			if a == b || a.freq != b.freq {
				continue
			}
			x := b.x + b.x - a.x
			y := b.y + b.y - a.y
			if x < 0 || y < 0 || x >= size.x || y >= size.y {
				continue
			}
			antinodes[xy{x, y}] = true
		}
	}
	return len(antinodes)
}

func part2(input []antenna, size xy) int {
	antinodes := make(map[xy]bool)
	for _, a := range input {
		for _, b := range input {
			if a == b || a.freq != b.freq {
				continue
			}
			for i := 0; i < size.x+size.y; i++ {
				x := b.x + i*(b.x-a.x)
				y := b.y + i*(b.y-a.y)
				if x < 0 || y < 0 || x >= size.x || y >= size.y {
					break
				}
				antinodes[xy{x, y}] = true
			}
		}
	}
	return len(antinodes)
}

func main() {
	input, size, err := readFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("p1:", part1(input, size))
	fmt.Println("p2:", part2(input, size))
}
