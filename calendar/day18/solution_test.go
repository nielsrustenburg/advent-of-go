package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T){
	cases := []Case{exp0,exp1,exp2,exp3}
	for _, casus := range cases{
		fmt.Println(Rewrite(casus.input))
		parsed := ParseLine(casus.input)
		got := parsed.Solve()
		if(got != casus.expected1){
			t.Errorf("expected: %v got %v", casus.expected1, got)
		}
	}
}

func TestSolve2(t *testing.T){
	cases := []Case{exp0,exp1,exp2,exp3}
	for _, casus := range cases{
		parsed := ParseLine(Rewrite(casus.input))
		got := parsed.Solve()
		if(got != casus.expected2){
			t.Errorf("expected: %v got %v", casus.expected2, got)
		}
	}
}

func TestRewrite(t *testing.T){
	input := "6 * 3 + 5 * ((9 * 9 + 3) + (2 + 5 * 6 + 4)) + 7 + 9"
	got := Rewrite(input)
	expected := "6 * (3 + 5) * (((((9 * (9 + 3)) + ((2 + 5) * (6 + 4)))) + 7) + 9)"
	if(got != expected){
		t.Errorf("\n exp: %v\n got: %v", expected, got)
	}
}

var exp0 = Case{`2 * 3 + (4 * 5)`, 26 ,46}
var exp1 = Case{`5 + (8 * 3 + 9 + 3 * 4 * 3)`, 437,1445}
var exp2 = Case{`5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))`, 12240,669060}
var exp3 = Case{`((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2`, 13632,23340}

type Case struct{
	input string
	expected1 int
	expected2 int
}