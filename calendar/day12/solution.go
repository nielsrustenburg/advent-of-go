package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
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

func SolvePartOne(lines []string) int{
	ship := Ship{0,0,0,0,0}
	for _, line := range lines{
		ship.Act(line)
	}
	return Abs(ship.Hor) + Abs(ship.Vert)
}

func SolvePartTwo(lines []string) int{
	ship := Ship{0,0,0,10,1}
	for _, line := range lines{
		ship.Act2(line)
	}
	return Abs(ship.Hor) + Abs(ship.Vert)
}
	
type Ship struct {
	Hor int
	Vert int
	Angle int
	WpHor int
	WpVert int
}

func Abs(val int) int{
	if (val < 0){
		return -val
	}
	return val
}

func(ship *Ship) Act(action string){
	val, _ := strconv.Atoi(action[1:])
	switch op := string(action[0]); op{
	case "N":
		ship.Vert += val
	case "S":
		ship.Vert -= val
	case "E":
		ship.Hor += val
	case "W":
		ship.Hor -= val
	case "R":
		ship.Angle = (ship.Angle + val) % 360
	case "L":
		ship.Angle = (ship.Angle - val + 360) % 360
	case "F":
		switch ship.Angle {
		case 0:
			ship.Hor += val
		case 90: 
			ship.Vert -= val
		case 180: 
			ship.Hor -= val
		case 270: 
			ship.Vert += val
		default: 
			panic("DID NOT EXPECT ANGLE THAT ISN'T MULTIPLE OF 90")
		}
	default:
		panic("UNEXPECTED ACTION")
	}
}

func(ship *Ship) Act2(action string){
	val, _ := strconv.Atoi(action[1:])
	turned := false
	switch op := string(action[0]); op{
	case "N":
		ship.WpVert += val
	case "S":
		ship.WpVert -= val
	case "E":
		ship.WpHor += val
	case "W":
		ship.WpHor -= val
	case "R":
		ship.Angle = (val) % 360
		turned = true
	case "L":
		ship.Angle = (-val + 360) % 360
		turned = true
	case "F":
		ship.Vert += val * ship.WpVert
		ship.Hor +=  val * ship.WpHor
	default:
		panic("UNEXPECTED ACTION")
	}

	if(turned){
		currentVert := ship.WpVert
		currentHor := ship.WpHor
		switch ship.Angle {
		case 0:
		case 90: 
			ship.WpVert = -currentHor
			ship.WpHor = currentVert
		case 180: 
			ship.WpVert = -currentVert
			ship.WpHor = -currentHor
		case 270: 
			ship.WpVert = currentHor
			ship.WpHor = -currentVert
		default: 
			panic("DID NOT EXPECT ANGLE THAT ISN'T MULTIPLE OF 90")
		}
		ship.Angle = 0
	}
}