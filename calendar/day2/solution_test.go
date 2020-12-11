package main

import "testing"

func TestCreatePolicy(t *testing.T){
	line := "9-10 b"
	got := CreatePolicy(line)
	expected := Policy{9,10,"b"}
	if(got != expected){
		t.Errorf("Expected: %v got %v", expected, got)
	}
}

func TestIsValidPassword(t *testing.T){
	testcases := []struct {
		input string
		expected bool
	}{
		{"1-3 a: abcde", true},
		{"1-3 b: cdefg", false},
		{"2-9 c: ccccccccc", true},
	}

	for _, testcase := range testcases{
		got := IsValidPassword(testcase.input)
		if(got != testcase.expected){
			t.Errorf("Expected: %v got %v", testcase.expected, got)
		}
	}
}