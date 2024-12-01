package main

import (
	"testing"
)

func TestDay01p1(t *testing.T) {
	left, right, err := readFile("day01.example.in")
	if err != nil {
		t.Fatalf("%s", err)
	}
	sum, err := day01p1(left, right)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if sum != 11 {
		t.Fatalf(`sum = %d, expected 11`, sum)
	}
}

func TestDay01p2(t *testing.T) {
	left, right, err := readFile("day01.example.in")
	if err != nil {
		t.Fatalf("%s", err)
	}
	score, err := day01p2(left, right)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if score != 31 {
		t.Fatalf(`score = %d, expected 31`, score)
	}
}
