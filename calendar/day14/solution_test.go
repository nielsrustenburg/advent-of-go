package main

import (
	"strings"
	"testing"
)

func TestSolvePartTwo(t *testing.T){
	lines := strings.Split(testinput, "\n")
	got := SolvePartOne(lines)
	expected := int64(165)

	if(got != expected){
		t.Errorf("expected: %v got %v", expected, got)
	}
}

const testinput = `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`