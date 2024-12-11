package main

import (
	"fmt"
	"os"
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

func blink(n int, times int) int {
	if times == 0 {
		return 1
	}
	if n == 0 {
		return blink(1, times-1)
	}
	half, power := splitEvenDigits(n)
	if n >= power/10 {
		left := n / half
		right := n % half
		return blink(left, times-1) + blink(right, times-1)
	}
	return blink(n*2024, times-1)
}

func part1(input []int) int {
	output := 0
	for _, n := range input {
		output += blink(n, 25)
	}
	return output
}

func part2(input []int) int {
	output := 0
	for _, n := range input {
		output += blink(n, 75)
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
	expect(55312, part1(inEx))

	inCh, err := readFile("day11/challenge.in")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("p1:", part1(inCh))
	fmt.Println("p2:", part2(inCh))
}
