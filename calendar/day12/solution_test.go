package main

import (
	"strings"
	"testing"
)

func TestSolvePartOne(t *testing.T){
	lines := strings.Split(string(testinput), "\n")
	got := SolvePartOne(lines)
	expected := 25

	if(got != expected){
		t.Errorf("expected: %v got %v", expected, got)
	}
}

func TestSolvePartTwo(t *testing.T){
	lines := strings.Split(string(testinput), "\n")
	got := SolvePartTwo(lines)
	expected := 286

	if(got != expected){
		t.Errorf("expected: %v got %v", expected, got)
	}
}

const testinput = `F10
N3
F7
R90
F11`