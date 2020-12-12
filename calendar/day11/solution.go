package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const floor,empty,taken = byte(46), byte(76), byte(35)

func main(){
	Solve("input.txt")
}

func Solve(file string){
	input, err := ioutil.ReadFile(file)
	if(err != nil){
		panic(err)
	}
	lines := strings.Split(string(input), "\n")
	p1 := SolvePartOne(lines)
	fmt.Println(p1)
	p2 := SolvePartTwo(lines)
	fmt.Println(p2)
}

func SolvePartOne(lines []string)int{
	grid := InitializeGrid(lines)
	vmap := CreateVisionMap(false, grid)
	stable := Simulate(false, grid, vmap)
	return CountTakenSeats(stable)
}

func SolvePartTwo(lines []string)int{
	grid := InitializeGrid(lines)
	vmap := CreateVisionMap(true, grid)
	stable := Simulate(true, grid, vmap)
	return CountTakenSeats(stable)
}

func InitializeGrid(lines []string)[][]byte{
	grid := make([][]byte, len(lines))
	for i, _ := range grid{
		grid[i] = []byte(lines[i])
	}
	return grid
}

func CountTakenSeats(grid [][]byte) int {
	count := 0
	for _, row := range grid{
		for _, val := range row{
			if(val == taken){
				count++
			}
		}
	}
	return count
}

func Simulate(distantnb bool, grid [][]byte, vmap map[Coordinate][]Coordinate)[][]byte{
	nextgrid :=  make([][]byte, len(grid))
	stable := true
	
	for rown, row := range grid{
		nextgrid[rown] = make([]byte, len(row))
		for coln, val := range row{
			coord := Coordinate{rown, coln}
			nbcount := 0
			nextgrid[rown][coln] = grid[rown][coln]

			if(val != floor){
				for _,nb := range vmap[coord]{
					if(grid[nb.Row][nb.Col] == taken){
						nbcount++
					}
				}
				if(val == empty && nbcount == 0){
					nextgrid[rown][coln] = taken
					stable = false
				}
				if(val == taken && (nbcount >= 5 || !distantnb && nbcount >= 4)){
					nextgrid[rown][coln] = empty
					stable = false
				}
			}
		}
	}

	if(stable){
		return grid
	} else {
		return Simulate(distantnb, nextgrid, vmap) 
	}
}

func CountActiveNeighbours(col,row int, grid [][]byte) int{
	wok := IsWithinBounds(col-1,0,len(grid[0]))
	eok := IsWithinBounds(col+1,0,len(grid[0]))
	nok := IsWithinBounds(row-1,0,len(grid))
	sok := IsWithinBounds(row+1,0,len(grid))
	sum := 0
	sum = IncrementByOneIfTrue(sum, wok && grid[row][col-1] == taken)
	sum = IncrementByOneIfTrue(sum, eok && grid[row][col+1] == taken)
	sum = IncrementByOneIfTrue(sum, nok && grid[row-1][col] == taken)
	sum = IncrementByOneIfTrue(sum, sok && grid[row+1][col] == taken)
	sum = IncrementByOneIfTrue(sum, wok && nok && grid[row-1][col-1] == taken)
	sum = IncrementByOneIfTrue(sum, wok && sok && grid[row+1][col-1] == taken)
	sum = IncrementByOneIfTrue(sum, eok && nok && grid[row-1][col+1] == taken)
	sum = IncrementByOneIfTrue(sum, eok && sok && grid[row+1][col+1] == taken)
	return sum
}

func CreateVisionMap(distant bool, grid [][]byte) map[Coordinate][]Coordinate{
	visionmap := make(map[Coordinate][]Coordinate)
	for rown, row := range grid{
		for coln, val := range row{
			if(val != floor){
				neighbours := make([]Coordinate, 0)
				coord := Coordinate{rown, coln}
				neighbours = AddNeighbourInDirection(distant, coord, 1,0,grid,neighbours)
				neighbours = AddNeighbourInDirection(distant, coord, -1,0,grid,neighbours)
				neighbours = AddNeighbourInDirection(distant, coord, 0,1,grid,neighbours)
				neighbours = AddNeighbourInDirection(distant, coord, 0,-1,grid,neighbours)
				neighbours = AddNeighbourInDirection(distant, coord, 1,-1,grid,neighbours)
				neighbours = AddNeighbourInDirection(distant, coord, 1,1,grid,neighbours)
				neighbours = AddNeighbourInDirection(distant, coord, -1,1,grid,neighbours)
				neighbours = AddNeighbourInDirection(distant, coord, -1,-1,grid,neighbours)
				visionmap[coord] = neighbours
			}
		}
	}
	return visionmap
}

func AddNeighbourInDirection(distant bool, coord Coordinate, rowdir, coldir int, grid [][]byte, neighbours []Coordinate) []Coordinate{
	for step := 1; true; step++{
		rowok := IsWithinBounds(coord.Row + rowdir*step, 0, len(grid))
		colok := IsWithinBounds(coord.Col + coldir*step, 0, len(grid[0]))
		if(!(rowok && colok)){
			return neighbours
		}
		other := Coordinate{coord.Row + rowdir*step, coord.Col + coldir*step}
		if(grid[other.Row][other.Col] != floor){
			return append(neighbours, other)
		}
		if(!distant){
			return neighbours
		}
	}
	panic ("DID NOT EXPECT TO GET HERE")
}

func IncrementByOneIfTrue(num int, val bool) int{
	if(val){
		return num+1
	}
	return num
}

func IsWithinBounds(k, min, max int) bool{
	return k >= min && k < max
}

func PrintGrid(grid [][]byte){
	fmt.Println()
	fmt.Printf("%v by %v \n", len(grid[5]), len(grid))
	for _, row := range grid{
		fmt.Println(string(row))
	}
}

type Coordinate struct {
	Row int
	Col int
}