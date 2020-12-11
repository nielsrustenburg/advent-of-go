package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main(){
	Solve("input.txt")
}

func Solve(file string){
	input, err := ioutil.ReadFile(file)
	if(err != nil){
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	slope1 := Slope{3, 1}
	part1 := TreesEncountered(slope1, lines)
	fmt.Println(part1)

	part2 := part1
	slopes := []Slope{
		Slope{1,1},
		Slope{5,1},
		Slope{7,1},
		Slope{1,2},
	}

	for _, slope2 := range slopes{
		part2 = part2 * TreesEncountered(slope2, lines)
	}

	fmt.Println(part2)
}

func TreesEncountered(slope Slope, hill []string) int {
	tree := []byte("#")[0]
	hillHeight := len(hill)
	hillWidth := len(hill[0])
	position := Coordinate {0, 0}
	treesEncountered := 0
	for position.Y < hillHeight{
		if(hill[position.Y][position.X] == tree){
			treesEncountered++
		}
		position.X = (position.X + slope.Right) % hillWidth
		position.Y += slope.Down 
	}
	return treesEncountered
}

type Coordinate struct {
	X int
	Y int
}

type Slope struct{
	Right int
	Down int
}