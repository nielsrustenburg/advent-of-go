package main

import (
	"strings"
	"testing"
)

func TestTreesEncountered(t *testing.T){
	const input = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

	hill := strings.Split(string(input), "\n")
	slope := Slope {3, 1}

	got := TreesEncountered(slope, hill)
	expected := 7 

	if(got != expected){
		t.Errorf("Expected: %v got %v", expected, got)
	}
}