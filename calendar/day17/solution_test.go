package main

import (
	"testing"
)

func TestNeighbours(t *testing.T){
	coord := Coord{0,0,0,0}
	got := len(coord.Neighbours(true))
	expected := 80

	if(got != expected){
		t.Errorf("expected: %v got %v", expected, got)
	}
}
