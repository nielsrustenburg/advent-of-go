package main

import "testing"

func TestFindSumNums(t *testing.T){
	nums := []int{1721,979,366,299,675,1456}
	a,b := FindSumNums(2020,nums)
	expected := 514579
	got := a*b
	if(got != expected){
		t.Errorf("Expected: 514579 got %v", got)
	}
}