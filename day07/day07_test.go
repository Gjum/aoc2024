package main

import (
	"testing"
)

func TestDay07p1(t *testing.T) {
	themap, err := readFile("example.in")
	if err != nil {
		t.Fatal(err)
	}
	solution := part1(themap)
	if solution != 3749 {
		t.Fatalf(`solution = %d, expected 3749`, solution)
	}
}

func TestConcat(t *testing.T) {
	if concat(5, 1) != 51 {
		t.Fatal(`failed at 1`)
	}
	if concat(5, 9) != 59 {
		t.Fatal(`failed at 99`)
	}
	if concat(5, 10) != 510 {
		t.Fatal(`failed at 10`)
	}
	if concat(5, 99) != 599 {
		t.Fatal(`failed at 99`)
	}
	if concat(5, 100) != 5100 {
		t.Fatal(`failed at 100`)
	}
}

func TestDay07p2(t *testing.T) {
	themap, err := readFile("example.in")
	if err != nil {
		t.Fatal(err)
	}
	solution := part2(themap)
	if solution != 11387 {
		t.Fatalf(`solution = %d, expected 11387`, solution)
	}
}
