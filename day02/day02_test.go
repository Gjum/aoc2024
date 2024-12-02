package main

import (
	"testing"
)

func TestDay02p1(t *testing.T) {
	reports, err := readFile("example.in")
	if err != nil {
		t.Fatal(err)
	}
	numSafe := part1(reports)
	if err != nil {
		t.Fatal(err)
	}
	if numSafe != 2 {
		t.Fatalf(`numSafe = %d, expected 2`, numSafe)
	}
}

func TestDay02p2(t *testing.T) {
	reports, err := readFile("example.in")
	if err != nil {
		t.Fatal(err)
	}
	numSafe := part2(reports)
	if err != nil {
		t.Fatal(err)
	}
	if numSafe != 4 {
		t.Fatalf(`numSafe = %d, expected 4`, numSafe)
	}
}
