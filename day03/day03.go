package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
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

func main() {
	memory, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	sum := part1(memory)
	fmt.Println("p1:", sum)
}
