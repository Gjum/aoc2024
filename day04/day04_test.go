package main

import (
	"testing"
)

func TestDay04p1(t *testing.T) {
	grid, err := readFile("example.in")
	if err != nil {
		t.Fatal(err)
	}
	numFound := part1(grid)
	if numFound != 18 {
		t.Fatalf(`numFound = %d, expected 18`, numFound)
	}
}

func TestDay04p2(t *testing.T) {
	grid, err := readFile("example.in")
	if err != nil {
		t.Fatal(err)
	}
	numFound := part2(grid)
	if numFound != 9 {
		t.Fatalf(`numFound = %d, expected 9`, numFound)
	}
}
