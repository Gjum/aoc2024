package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Sign(x int) int {
	if x < 0 {
		return -1
	}
	if x > 0 {
		return 1
	}
	return 0
}

func readFile(inpath string) ([][]int, error) {
	fp, err := os.Open(inpath)
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	scanner := bufio.NewScanner(fp)
	reports := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		report := []int{}
		for _, word := range words {
			n, err := strconv.Atoi(word)
			if err != nil {
				return nil, err
			}
			report = append(report, n)
		}
		reports = append(reports, report)
	}
	return reports, nil
}

func isSafe(report []int) bool {
	incRep := report[0] < report[1]
	for i := 1; i < len(report); i++ {
		diff := Abs(report[i-1] - report[i])
		incPair := report[i-1] < report[i]
		if diff < 1 || diff > 3 || incPair != incRep {
			return false
		}
	}
	return true
}

func part1(reports [][]int) int {
	numSafe := 0
	for _, report := range reports {
		if isSafe(report) {
			numSafe++
		}
	}
	return numSafe
}

func part2(reports [][]int) int {
	numSafe := 0
	for _, report := range reports {
		if isSafe(report) {
			numSafe++
			continue
		}
		for i := range report {
			without := []int{}
			without = append(without, report[:i]...)
			without = append(without, report[i+1:]...)
			if isSafe(without) {
				numSafe++
				break
			}
		}
	}
	return numSafe
}

func main() {
	left, err := readFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	numSafe := part1(left)
	fmt.Println("p1:", numSafe)

	numSafe2 := part2(left)
	fmt.Println("p2:", numSafe2)
}
