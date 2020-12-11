package main

import (
	"strings"
	"testing"
)

func TestSolvePartOne(t *testing.T){
	lines := strings.Split(string(testinput), "\n")
	got := SolvePartOne(lines, 5)
	expected := 127

	if(got != expected){
		t.Errorf("expected: %v got %v", expected, got)
	}
}

func TestIsSumOfPredecessors(t *testing.T){
	predecessors := []int {35,20,20,25,47}

	got := IsSumOfTwoPredecessors(predecessors, 40)
	expected := true

	if(got != expected){
		t.Errorf("expected: %v got %v", expected, got)
	}

	got2 := IsSumOfTwoPredecessors(predecessors, 40)
	expected2 := true

	
	if(got2 != expected2){
		t.Errorf("expected: %v got %v", expected2, got2)
	}
}

func TestFindContiguousSumSet(t *testing.T){
	lines := strings.Split(string(testinput), "\n")
	got := SolvePartTwo(lines, 127)
	expected := 62

	if(got != expected){
		t.Errorf("expected: %v got %v", expected, got)
	}
}

const testinput = `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`