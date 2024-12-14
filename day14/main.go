package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type xy struct {
	x, y int
}

type Robot struct {
	p, v xy
}

func readFile(inpath string) ([]Robot, error) {
	bytes, err := os.ReadFile(inpath)
	if err != nil {
		return nil, err
	}
	texts := strings.Split(strings.Trim(string(bytes[:]), "\n"), "\n")
	var input []Robot
	re := regexp.MustCompile(`p=([-0-9]+),([-0-9]+) v=([-0-9]+),([-0-9]+)`)
	for _, text := range texts {
		matches := re.FindStringSubmatch(text)
		px, _ := strconv.Atoi(matches[1])
		py, _ := strconv.Atoi(matches[2])
		vx, _ := strconv.Atoi(matches[3])
		vy, _ := strconv.Atoi(matches[4])
		input = append(input, Robot{xy{px, py}, xy{vx, vy}})
	}
	return input, nil
}

func part1(input []Robot, roomsize xy) int {
	nw := 0
	ne := 0
	sw := 0
	se := 0
	for _, robot := range input {
		x := (robot.p.x + robot.v.x*100 + roomsize.x*100) % roomsize.x
		y := (robot.p.y + robot.v.y*100 + roomsize.y*100) % roomsize.y
		if x < roomsize.x/2 {
			if y < roomsize.y/2 {
				nw++
			}
			if y > roomsize.y/2 {
				sw++
			}
		}
		if x > roomsize.x/2 {
			if y < roomsize.y/2 {
				ne++
			}
			if y > roomsize.y/2 {
				se++
			}
		}
	}
	return nw * ne * sw * se
}

func printMap(step int, input []Robot, roomsize xy) {
	fmt.Print(step, "\n")
	s := ""
	for y := 0; y < roomsize.y; y++ {
		for x := 0; x < roomsize.x; x++ {
			occupied := false
			for _, robot := range input {
				if robot.p.x == x && robot.p.y == y {
					occupied = true
				}
			}
			if occupied {
				s += "#"
			} else {
				s += " "
			}
		}
		s += "\n"
	}
	fmt.Print(s)
}

func part2(input []Robot, roomsize xy) {
	for step := 1; ; step++ {
		for ri, robot := range input {
			input[ri].p.x = (robot.p.x + robot.v.x + roomsize.x) % roomsize.x
			input[ri].p.y = (robot.p.y + robot.v.y + roomsize.y) % roomsize.y
		}
		// place horizontally adjacent robots adjacent in the array
		// so we don't need to plot every frame
		sort.Slice(input, func(a, b int) bool {
			if input[a].p.y == input[b].p.y {
				return input[a].p.x < input[b].p.x
			}
			return input[a].p.y < input[b].p.y
		})
		// find straight horizontal lines
		maxrun := 0
		run := 0
		for ri, robot := range input {
			if ri < 1 {
				continue
			}
			prev := input[ri-1]
			if prev.p.x == robot.p.x-1 {
				run++
			} else {
				run = 0
			}
			maxrun = max(maxrun, run)
		}
		if maxrun > 20 {
			printMap(step, input, roomsize)
		}
	}
}

func expect(expected, actual int) {
	if expected != actual {
		panic(fmt.Sprintf("got %d but expected %d\n", actual, expected))
	}
	fmt.Printf("OK %d\n", actual)
}

func main() {
	inEx, err := readFile("day14/example.in")
	if err != nil {
		fmt.Println(err)
		return
	}
	expect(12, part1(inEx, xy{11, 7}))

	inCh, err := readFile("day14/challenge.in")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("p1:", part1(inCh, xy{101, 103}))
	part2(inCh, xy{101, 103})
}
