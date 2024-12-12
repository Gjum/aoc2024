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
	lines := strings.Split(strings.Trim(string(bytes[:]), "\n"), "\n")
	var input [][]rune
	for _, line := range lines {
		var row []rune
		for _, c := range line {
			row = append(row, c)
		}
		input = append(input, row)
	}
	return input, nil
}

type xy struct {
	x, y int
}

func part1(input [][]rune) int {
	output := 0
	for y, row := range input {
		for x, c := range row {
			if input[y][x] > 'Z' {
				continue // already visited
			}
			area := 0
			perim := 0
			visited := c - 'A' + 'a'
			queue := []xy{{x, y}}
			for len(queue) > 0 {
				p := queue[0]
				queue = queue[1:]
				if p.x < 0 || p.y < 0 || p.x >= len(input[0]) || p.y >= len(input) {
					// out of bounds
					perim++
					continue
				}
				if input[p.y][p.x] == visited {
					continue // already visited
				}
				if input[p.y][p.x] != c {
					// other region
					perim++
					continue
				}
				area++
				input[p.y][p.x] = visited // mark visited
				queue = append(queue, xy{p.x + 1, p.y}, xy{p.x - 1, p.y}, xy{p.x, p.y + 1}, xy{p.x, p.y - 1})
			}
			// fmt.Printf("%c %c %v %v %v\n", c, lower, area, perim, area*perim)
			output += area * perim
		}
	}
	return output
}

func part2(input [][]rune) int {
	output := 0
	return output
}

func expect(expected, actual int) {
	if expected != actual {
		panic(fmt.Sprintf("got %d but expected %d\n", actual, expected))
	}
	fmt.Printf("OK %d\n", actual)
}

func main() {
	inEx, err := readFile("day12/example.in")
	if err != nil {
		fmt.Println(err)
		return
	}
	expect(1930, part1(inEx))
	expect(0, part2(inEx))

	inCh, err := readFile("day12/challenge.in")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("p1:", part1(inCh))
	fmt.Println("p2:", part2(inCh))
}
