package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readFile(inpath string) ([]int, error) {
	bytes, err := os.ReadFile(inpath)
	if err != nil {
		return nil, err
	}
	line := strings.Split(strings.Trim(string(bytes[:]), "\n"), " ")
	var input []int
	for _, word := range line {
		row, err := strconv.Atoi(word)
		if err != nil {
			return nil, err
		}
		input = append(input, row)
	}
	return input, nil
}

func splitEvenDigits(n int) (int, int) {
	power := 100
	half := 10
	for n > power {
		power *= 100
		half *= 10
	}
	return half, power
}

type nextsplit struct {
	// when number immediately splits into a,b then depth=1
	depth, a, b int
}

func findNextSplit(nOrig int, cache map[int]nextsplit) nextsplit {
	encounters := make([]int, 0, 8) // TODO reuse
	var found nextsplit
	n := nOrig
	for {
		if cached, exists := cache[n]; exists {
			found = cached
			break
		}
		if n == 0 {
			encounters = append(encounters, n)
			n = 1
			continue
		}
		half, power := splitEvenDigits(n)
		if n < power/10 {
			encounters = append(encounters, n)
			n = n * 2024
			continue
		}
		found = nextsplit{1, n / half, n % half}
		cache[n] = found
		break
	}
	slices.Reverse(encounters)
	for d, n := range encounters {
		cache[n] = nextsplit{d + 2, found.a, found.b}
	}
	return cache[nOrig]
}

type todo struct {
	n, depth int
}

func run(input []int, maxDepth int) int {
	cache := make(map[int]nextsplit)
	todos := make([]todo, 0, 1<<16)
	for _, n := range input {
		todos = append(todos, todo{n, maxDepth})
	}
	output := 0
	for len(todos) > 0 {
		n, depth := todos[0].n, todos[0].depth
		todos = todos[1:] // poll

		ns := findNextSplit(n, cache)

		if depth < ns.depth {
			output++
			continue
		}
		if depth == ns.depth {
			output += 2
			continue
		}
		depth -= ns.depth

		todos = append(todos, todo{ns.a, depth}, todo{ns.b, depth})
	}
	return output
}

func expect(expected, actual int) {
	if expected != actual {
		fmt.Printf("got %d but expected %d\n", actual, expected)
		panic(1)
	}
	fmt.Printf("OK %d\n", actual)
}

func main() {
	inEx, err := readFile("day11/example.in")
	if err != nil {
		fmt.Println(err)
		return
	}
	expect(55312, run(inEx, 25))

	inCh, err := readFile("day11/challenge.in")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("p1:", run(inCh, 25))
	fmt.Println("p2:", run(inCh, 75))
}
