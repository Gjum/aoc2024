package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func day01p1(left, right []int) (int, error) {
	sort.Ints(left)
	sort.Ints(right)

	sum := 0
	for i := range left {
		sum += Abs(right[i] - left[i])
	}
	return sum, nil
}

func day01p2(left, right []int) (int, error) {
	sort.Ints(left)
	sort.Ints(right)

	counts := make(map[int]int)
	for i := range right {
		counts[right[i]] += 1
	}

	total := 0
	for i := range left {
		total += left[i] * counts[left[i]]
	}
	return total, nil
}

func readFile(inpath string) ([]int, []int, error) {
	fp, err := os.Open(inpath)
	if err != nil {
		return nil, nil, err
	}
	defer fp.Close()
	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanWords)
	var left, right []int
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, nil, err
		}
		if len(left) <= len(right) {
			left = append(left, n)
		} else {
			right = append(right, n)
		}
	}
	if len(left) != len(right) {
		return nil, nil, fmt.Errorf("different lengths: %d %d", len(left), len(right))
	}
	return left, right, nil
}

func main() {
	left, right, err := readFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	sum, err := day01p1(left, right)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("p1:", sum)

	score, err := day01p2(left, right)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("p2:", score)
}
