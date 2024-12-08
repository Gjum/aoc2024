package main

import (
	"testing"
)

func TestDay08p1(t *testing.T) {
	antennas, size, err := readFile("example.in")
	if err != nil {
		t.Fatal(err)
	}
	solution := part1(antennas, size)
	if solution != 14 {
		t.Fatalf(`solution = %d, expected 14`, solution)
	}
}

func TestDay08p2(t *testing.T) {
	antennas, size, err := readFile("example.in")
	if err != nil {
		t.Fatal(err)
	}
	solution := part2(antennas, size)
	if solution != 34 {
		t.Fatalf(`solution = %d, expected 34`, solution)
	}
}
