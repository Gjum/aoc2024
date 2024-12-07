package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(inpath string) ([][]int, error) {
	bytes, err := os.ReadFile(inpath)
	if err != nil {
		return nil, err
	}
	var eqns [][]int
	for _, line := range strings.Split(strings.Trim(string(bytes[:]), "\n"), "\n") {
		var eqn []int
		for i, word := range strings.Split(line, " ") {
			if i == 0 {
				word = word[:len(word)-1]
			}
			n, err := strconv.Atoi(word)
			if err != nil {
				return nil, err
			}
			eqn = append(eqn, n)
		}
		eqns = append(eqns, eqn)
	}
	return eqns, nil
}

func computes(soln int, tally int, eqn []int) bool {
	if len(eqn) <= 0 {
		return soln == tally
	}
	return computes(soln, tally+eqn[0], eqn[1:]) || computes(soln, tally*eqn[0], eqn[1:])
}

func part1(eqns [][]int) int {
	sum := 0
	for _, eqn := range eqns {
		if computes(eqn[0], eqn[1], eqn[2:]) {
			sum += eqn[0]
		}
	}
	return sum
}

func concat(a, b int) int {
	power := 1
	for ; power <= b; power *= 10 {
	}
	return power*a + b
}

func computes2(soln int, tally int, eqn []int) bool {
	if len(eqn) <= 0 {
		return soln == tally
	}
	return computes2(soln, tally+eqn[0], eqn[1:]) || computes2(soln, tally*eqn[0], eqn[1:]) || computes2(soln, concat(tally, eqn[0]), eqn[1:])
}

func part2(eqns [][]int) int {
	sum := 0
	for _, eqn := range eqns {
		if computes2(eqn[0], eqn[1], eqn[2:]) {
			sum += eqn[0]
		}
	}
	return sum
}

func main() {
	eqns, err := readFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("p1:", part1(eqns))
	fmt.Println("p2:", part2(eqns))
}
