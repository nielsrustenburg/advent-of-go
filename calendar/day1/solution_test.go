package main

import "testing"

func TestFindSumPairs(t *testing.T){
	nums := []int{1721,979,366,299,675,1456}
	a,b := FindSumPairs(2020,nums)
	expected := 514579
	got := a*b
	if(got != expected){
		t.Errorf("Expected: %v got %v", expected, got)
	}
}

func TestFindSumTrios(t *testing.T){
	nums := []int{1721,979,366,299,675,1456}
	a,b,c := FindSumTrios(2020,nums)
	expected := 241861950
	got := a*b*c
	if(got != expected){
		t.Errorf("Expected: %v got %v", expected, got)
	}
}