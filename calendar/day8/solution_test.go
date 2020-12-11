package main

import (
	"strings"
	"testing"
)

func TestFindLoop(t *testing.T){

	lines := strings.Split(string(testinput), "\n")
	hh := MakeHandheld(lines)
	got := FindLoop(hh)
	expected := 7
	if(got != expected){
		t.Errorf("expected: %v got %v", expected, got)
	}
}

const testinput = `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`