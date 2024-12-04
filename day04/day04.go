package main

import (
	"fmt"
	"os"
	"strings"
)

func readFile(inpath string) ([]string, error) {
	bytes, err := os.ReadFile(inpath)
	if err != nil {
		return nil, err
	}
	grid := strings.Split(strings.Trim(string(bytes[:]), "\n"), "\n")
	return grid, nil
}

type dir struct {
	dx int
	dy int
}

func getc(grid []string, x, y int) byte {
	if x < 0 || y < 0 || x >= len(grid[0]) || y >= len(grid) {
		return '_'
	}
	return grid[y][x]
}

func part1(grid []string) int {
	dirs := []dir{}
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx != 0 || dy != 0 {
				dirs = append(dirs, dir{dx, dy})
			}
		}
	}
	numFound := 0
	word := "XMAS"
	for y0, line := range grid {
		for x0 := range line {
		nextDir:
			for _, dir := range dirs {
				for i := range word {
					x := x0 + dir.dx*i
					y := y0 + dir.dy*i
					if getc(grid, x, y) != word[i] {
						continue nextDir
					}
				}
				numFound++
			}
		}
	}
	return numFound
}

func part2(grid []string) int {
	numFound := 0
	patterns := [][]string{
		{"M.S", ".A.", "M.S"},
		{"M.M", ".A.", "S.S"},
		{"S.M", ".A.", "S.M"},
		{"S.S", ".A.", "M.M"},
	}
	for y0, line := range grid {
		for x0 := range line {
		next_pattern:
			for _, pattern := range patterns {
				for dy, pline := range pattern {
					for dx, pc := range pline {
						c := getc(grid, x0+dx, y0+dy)
						if pc != '.' && pline[dx] != c {
							continue next_pattern
						}
					}
				}
				numFound++
			}
		}
	}
	return numFound
}

func main() {
	grid, err := readFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	numFound := part1(grid)
	fmt.Println("p1:", numFound)

	numFound2 := part2(grid)
	fmt.Println("p2:", numFound2)
}
