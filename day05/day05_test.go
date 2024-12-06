package main

import (
	"testing"
)

func TestDay05p1(t *testing.T) {
	precedences, books, err := readFile("example.in")
	if err != nil {
		t.Fatal(err)
	}
	solution := part1(precedences, books)
	if solution != 143 {
		t.Fatalf(`solution = %d, expected 143`, solution)
	}
}

func TestDay05p2(t *testing.T) {
	precedences, books, err := readFile("example.in")
	if err != nil {
		t.Fatal(err)
	}
	solution := part2(precedences, books)
	if solution != 123 {
		t.Fatalf(`solution = %d, expected 123`, solution)
	}
}
