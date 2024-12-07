package main

import (
	"testing"
)

func TestDay06p1(t *testing.T) {
	themap, err := readFile("example.in")
	if err != nil {
		t.Fatal(err)
	}
	solution := part1(themap)
	if solution != 41 {
		t.Fatalf(`solution = %d, expected 41`, solution)
	}
}

func TestDay06p2(t *testing.T) {
	themap, err := readFile("example.in")
	if err != nil {
		t.Fatal(err)
	}
	solution := part2(themap)
	if solution != 6 {
		t.Fatalf(`solution = %d, expected 6`, solution)
	}
}
