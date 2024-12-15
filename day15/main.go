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

func pushInto(puzzle *Puzzle, pos xy, dir rune, affected map[xy]rune) bool {
	switch getworld(puzzle.world, pos) {
	case '#':
		return false
	case '.':
		return true
	case 'O':
		affected[pos] = 'O'
		into := pos.plusDir(dir)
		return pushInto(puzzle, into, dir, affected)
	case '[':
		pos2 := pos.plus(1, 0)
		if affected[pos] == '[' && affected[pos2] == ']' {
			return true
		}
		affected[pos] = '['
		affected[pos2] = ']'
		into := pos.plusDir(dir)
		into2 := into.plus(1, 0)
		return pushInto(puzzle, into, dir, affected) && pushInto(puzzle, into2, dir, affected)
	case ']':
		pos2 := pos.plus(-1, 0)
		if affected[pos] == ']' && affected[pos2] == '[' {
			return true
		}
		affected[pos] = ']'
		affected[pos2] = '['
		into := pos.plusDir(dir)
		into2 := into.plus(-1, 0)
		return pushInto(puzzle, into, dir, affected) && pushInto(puzzle, into2, dir, affected)
	default:
		return false
	}
}

func part1(puzzle Puzzle) int {
	for _, dir := range puzzle.moves {
		affected := make(map[xy]rune)
		robotNew := puzzle.robot.plusDir(dir)
		if pushInto(&puzzle, robotNew, dir, affected) {
			for from := range affected {
				setworld(puzzle.world, from, '.')
			}
			for from, tile := range affected {
				setworld(puzzle.world, from.plusDir(dir), tile)
			}
			puzzle.robot = robotNew
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
			if c == '[' {
				output += x + y*100
			}
		}
	}
	return output
}

func part2(puzzle Puzzle) int {
	var world [][]rune
	for _, rowOrig := range puzzle.world {
		var row []rune
		for _, c := range rowOrig {
			if c == 'O' {
				row = append(row, '[', ']')
			} else {
				row = append(row, c, c)
			}
		}
		world = append(world, row)
	}
	rx := puzzle.robot.x * 2
	ry := puzzle.robot.y
	return part1(Puzzle{world, xy{rx, ry}, puzzle.moves})
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
	inEx, err = readFile("day15/example.in")
	if err != nil {
		fmt.Println(err)
		return
	}
	expect(9021, part2(inEx))

	inCh, err := readFile("day15/challenge.in")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("p1:", part1(inCh))

	inCh, err = readFile("day15/challenge.in")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("p2:", part2(inCh))
}
