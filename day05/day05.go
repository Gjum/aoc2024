package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type precedence struct {
	a int
	b int
}

func readFile(inpath string) ([]precedence, [][]int, error) {
	bytes, err := os.ReadFile(inpath)
	if err != nil {
		return nil, nil, err
	}
	sections := strings.Split(strings.Trim(string(bytes[:]), "\n"), "\n\n")

	var precedences []precedence
	for _, line := range strings.Split(sections[0], "\n") {
		split := strings.Split(line, "|")
		a, err := strconv.Atoi(split[0])
		if err != nil {
			return nil, nil, err
		}
		b, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, nil, err
		}
		precedences = append(precedences, precedence{a, b})
	}

	var books [][]int
	for _, line := range strings.Split(sections[1], "\n") {
		var pages []int
		for _, pagestr := range strings.Split(line, ",") {
			pagenr, err := strconv.Atoi(pagestr)
			if err != nil {
				return nil, nil, err
			}
			pages = append(pages, pagenr)
		}
		books = append(books, pages)
	}

	return precedences, books, nil
}

func part1(precedences []precedence, books [][]int) int {
	sum := 0
next_book:
	for _, book := range books {
		for iPage, p := range book {
			for iRel, pRel := range book {
				for _, prec := range precedences {
					if prec.a == p && prec.b == pRel && iPage > iRel {
						continue next_book
					}
					if prec.a == pRel && prec.b == p && iRel > iPage {
						continue next_book
					}
				}
			}
		}
		sum += book[len(book)/2]
	}
	return sum
}

// part1 could be solved using sorting like part2 too

func part2(precedences []precedence, books [][]int) int {
	precs := make(map[precedence]bool)
	for _, prec := range precedences {
		precs[prec] = true
	}
	sum := 0
	for _, bookOrig := range books {
		book := make([]int, len(bookOrig))
		copy(book, bookOrig)
		sort.Slice(book, func(a, b int) bool {
			return precs[precedence{book[a], book[b]}]
		})
		if !slices.Equal(book, bookOrig) {
			sum += book[len(book)/2]
		}
	}
	return sum
}

func main() {
	precedences, books, err := readFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("p1:", part1(precedences, books))
	fmt.Println("p2:", part2(precedences, books))
}
