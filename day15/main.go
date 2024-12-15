package main

import (
	"fmt"
	"os"
	"strings"
)

type xy struct {
	x, y int
}

func (pos xy) plus(x, y int) xy {
	return xy{pos.x + x, pos.y + y}
}

func (pos xy) plusDir(dir rune) xy {
	switch dir {
	case '>':
		return xy{pos.x + 1, pos.y}
	case '<':
		return xy{pos.x - 1, pos.y}
	case 'v':
		return xy{pos.x, pos.y + 1}
	case '^':
		return xy{pos.x, pos.y - 1}
	}
	panic("invalid dir")
}

func getworld(world [][]rune, pos xy) rune {
	if pos.x < 0 || pos.y < 0 || pos.x >= len(world[0]) || pos.y >= len(world) {
		return '~'
	}
	return world[pos.y][pos.x]
}

func setworld(world [][]rune, pos xy, tile rune) {
	if pos.x < 0 || pos.y < 0 || pos.x >= len(world[0]) || pos.y >= len(world) {
		return
	}
	world[pos.y][pos.x] = tile
}

type Puzzle struct {
	world [][]rune
	robot xy
	moves []rune
}

func readFile(inpath string) (Puzzle, error) {
	bytes, err := os.ReadFile(inpath)
	if err != nil {
		return Puzzle{nil, xy{}, nil}, err
	}
	sections := strings.Split(strings.Trim(string(bytes[:]), "\n"), "\n\n")
	var world [][]rune
	var robot xy
	var moves []rune
	for y, line := range strings.Split(sections[0], "\n") {
		var row []rune
		for x, c := range line {
			if c == '@' {
				robot.x = x
				robot.y = y
				c = '.'
			}
			row = append(row, c)
		}
		world = append(world, row)
	}
	for _, line := range strings.Split(sections[1], "\n") {
		for _, c := range line {
			moves = append(moves, c)
		}
	}
	return Puzzle{world, robot, moves}, nil
}

func pushInto1(puzzle *Puzzle, pos xy, dir rune) bool {
	switch getworld(puzzle.world, pos) {
	case '#':
		return false
	case '.':
		return true
	case 'O':
		into := pos.plusDir(dir)
		if pushInto1(puzzle, into, dir) {
			setworld(puzzle.world, into, getworld(puzzle.world, pos))
			setworld(puzzle.world, pos, '.')
			return true
		}
		return false
	default:
		return false
	}
}

func part1(puzzle Puzzle) int {
	for _, dir := range puzzle.moves {
		into := puzzle.robot.plusDir(dir)
		if pushInto1(&puzzle, into, dir) {
			setworld(puzzle.world, into, '.')
			puzzle.robot = into
		}
	}

	for y, row := range puzzle.world {
		for x, c := range row {
			if puzzle.robot.x == x && puzzle.robot.y == y {
				fmt.Print("@")
			} else {
				fmt.Printf("%c", c)
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")

	output := 0
	for y, row := range puzzle.world {
		for x, c := range row {
			if c == 'O' {
				output += x + y*100
			}
		}
	}
	return output
}

func part2(puzzle Puzzle) int {
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
	inEx, err := readFile("day15/example.in")
	if err != nil {
		fmt.Println(err)
		return
	}
	expect(10092, part1(inEx))
	expect(9021, part2(inEx))

	inCh, err := readFile("day15/challenge.in")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("p1:", part1(inCh))
	fmt.Println("p2:", part2(inCh))
}
