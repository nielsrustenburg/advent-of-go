package main

import (
	"fmt"
	"strings"
	"io/ioutil"
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
	grid := ParseLines(lines)
	afterSix := CycleNTimes(grid, 6, false)
	p1 := len(afterSix)
	fmt.Println(p1)
	afterSix4D := CycleNTimes(grid, 6, true)
	p2 := len(afterSix4D)
	fmt.Println(p2)
}

func ParseLines(lines []string) map[Coord]bool{
	grid := make(map[Coord]bool)
	for y, row := range lines{
		for x, char := range row{
			if(char == rune("#"[0])){
				grid[Coord{x,y,0,0}] = true
			}
		}
	}
	return grid
}

func CycleNTimes(grid map[Coord]bool, n int, fourD bool) map[Coord]bool{
	for i := 0; i < n; i++{
		grid = Next(grid, fourD)
	}
	return grid
}

func Next(grid map[Coord]bool, fourD bool) map[Coord]bool{
	nextGrid := make(map[Coord]bool)
	worthChecking := make(map[Coord]int)

	for coord, _ := range grid{
		nbs := coord.Neighbours(fourD)
		for _, nb := range nbs{
			current, _ := worthChecking[nb]
			worthChecking[nb] = current + 1
		}
	}

	for coord, nbCount := range worthChecking {
		switch now, _ := grid[coord]; now{
		case true:
			if(nbCount == 2 || nbCount == 3){
				nextGrid[coord] = true
			}
		case false:
			if(nbCount == 3){
				nextGrid[coord] = true
			}
		}
	}

	return nextGrid
}

//Does not take into account borders of grid.. BUT WE DON'T NEED TO HAHAAAAAAA
func(this Coord) Neighbours(fourD bool) []Coord{
	nbs := make([]Coord,0)
	for x := -1; x <= 1; x++{
		for y := -1; y <= 1; y++{
			for z := -1; z <= 1; z++{
				if(fourD){
					for f := -1; f <= 1; f++{
						if(f != 0 || x != 0 || y != 0 || z != 0){
							nbs = append(nbs, Coord{this.X+x, this.Y+y, this.Z+z, this.FOUR+f})
						}
					}
				}else {
					if(x != 0 || y != 0 || z != 0){
						nbs = append(nbs, Coord{this.X+x, this.Y+y, this.Z+z, 0})
					}
				}
			}
		}
	}
	return nbs
}

type Coord struct {
	X int
	Y int
	Z int
	FOUR int
}