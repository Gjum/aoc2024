package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part1(memb []byte) int {
	memory := string(memb[:])
	sum := 0
	re := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	for _, match := range re.FindAllStringSubmatch(memory, -1) {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		sum += a * b
	}
	return sum
}

func part2(memb []byte) int {
	memory := string(memb[:])
	sum := 0
	re := regexp.MustCompile(`do(n't)?\(\)|mul\(([0-9]+),([0-9]+)\)`)
	enabled := true
	for _, match := range re.FindAllStringSubmatch(memory, -1) {
		switch strings.Split(match[0], "(")[0] {
		case "do":
			enabled = true
		case "don't":
			enabled = false
		case "mul":
			if enabled {
				a, _ := strconv.Atoi(match[2])
				b, _ := strconv.Atoi(match[3])
				sum += a * b
			}
		}
	}
	return sum
}

func main() {
	memory, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	sum := part1(memory)
	fmt.Println("p1:", sum)

	sum2 := part2(memory)
	fmt.Println("p2:", sum2)
}
