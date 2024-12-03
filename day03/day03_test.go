package main

import (
	"os"
	"testing"
)

func TestDay03p1(t *testing.T) {
	memory, err := os.ReadFile("example.in")
	if err != nil {
		t.Fatal(err)
	}
	sum := part1(memory)
	if sum != 161 {
		t.Fatalf(`sum = %d, expected 161`, sum)
	}
}

func TestDay03p2(t *testing.T) {
	memory, err := os.ReadFile("example2.in")
	if err != nil {
		t.Fatal(err)
	}
	sum := part2(memory)
	if sum != 48 {
		t.Fatalf(`sum = %d, expected 48`, sum)
	}
}
