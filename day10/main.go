package main

import (
	"fmt"
	"os"
	"strings"
)

type xy struct{ x, y int }

func readFile(inpath string) ([][]rune, error) {
	bytes, err := os.ReadFile(inpath)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(strings.Trim(string(bytes[:]), "\n"), "\n")
	var input [][]rune
	for _, line := range lines {
		row := []rune{}
		for _, c := range line {
			row = append(row, c)
		}
		input = append(input, row)
	}
	return input, nil
}

func traceTrail(input [][]rune, peaks map[xy]bool, x, y int) int {
	c := input[y][x]
	if c == '9' {
		peaks[xy{x, y}] = true
		return 1
	}
	rating := 0
	if x+1 < len(input[0]) && input[y][x+1] == c+1 {
		rating += traceTrail(input, peaks, x+1, y)
	}
	if x > 0 && input[y][x-1] == c+1 {
		rating += traceTrail(input, peaks, x-1, y)
	}
	if y+1 < len(input) && input[y+1][x] == c+1 {
		rating += traceTrail(input, peaks, x, y+1)
	}
	if y > 0 && input[y-1][x] == c+1 {
		rating += traceTrail(input, peaks, x, y-1)
	}
	return rating
}

func part1(input [][]rune) int {
	output := 0
	for y, line := range input {
		for x, c := range line {
			if c == '0' {
				peaks := make(map[xy]bool)
				traceTrail(input, peaks, x, y)
				output += len(peaks)
			}
		}
	}
	return output
}

func part2(input [][]rune) int {
	output := 0
	for y, line := range input {
		for x, c := range line {
			if c == '0' {
				peaks := make(map[xy]bool)
				output += traceTrail(input, peaks, x, y)
			}
		}
	}
	return output
}

func expect(expected, actual int) bool {
	if expected != actual {
		fmt.Printf("got %d but expected %d\n", actual, expected)
	}
	return expected != actual
}

func main() {
	inEx, err := readFile("day10/example.in")
	if err != nil {
		fmt.Println(err)
		return
	}
	if expect(36, part1(inEx)) {
		return
	}
	if expect(81, part2(inEx)) {
		return
	}

	inCh, err := readFile("day10/challenge.in")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("p1:", part1(inCh))
	fmt.Println("p2:", part2(inCh))
}
