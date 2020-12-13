package main

import (
	"strings"
	"testing"
)

func TestSolvePartTwo(t *testing.T){
	inputs := []string{test1,test2,test3,test4,test5,test6}
	expecteds := []int{1068781,3417,754018,779210,1261476,1202161486}

	// inputs := []string{test2}
	// expecteds := []int{3417}

	for i, input := range inputs {
		bx := strings.Split(input, ",")
		got := SolvePartTwo(bx)
		expected := expecteds[i]
	
		if(got != expected){
			t.Errorf("expected: %v got %v", expected, got)
		}
	}
}

const test1 = `7,13,x,x,59,x,31,19`
const test2 = `17,x,13,19`
const test3 = `67,7,59,61`
const test4 = `67,x,7,59,61`
const test5 = `67,7,x,59,61`
const test6 = `1789,37,47,1889`