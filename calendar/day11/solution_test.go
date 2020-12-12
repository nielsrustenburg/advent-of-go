package main

import (
	"strings"
	"testing"
)

func TestSolvePartOne(t *testing.T){
	lines := strings.Split(string(testinput), "\n")
	got := SolvePartOne(lines)
	expected := 37

	if(got != expected){
		t.Errorf("expected: %v got %v", expected, got)
	}
}

func TestSolvePartTwo(t *testing.T){
	lines := strings.Split(string(testinput), "\n")
	got := SolvePartTwo(lines)
	expected := 26

	if(got != expected){
		t.Errorf("expected: %v got %v", expected, got)
	}
}

const testinput = `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`
