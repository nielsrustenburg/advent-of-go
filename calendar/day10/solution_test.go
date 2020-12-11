package main

import (
	"strings"
	"testing"
	"strconv"
)

func TestSolvePartOne(t *testing.T){
	lines := strings.Split(string(testinput), "\n")
	adapters := make([]int,0)
	for _, line := range lines{
		adapter, _ := strconv.Atoi(line)
		adapters = append(adapters, adapter)
	}
	got := SolvePartOne(adapters)
	expected := 220

	if(got != expected){
		t.Errorf("expected: %v got %v", expected, got)
	}
}

const testinput = `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`