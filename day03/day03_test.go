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
