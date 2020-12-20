package main

import (
	"fmt"
)

func main(){
	Solve("input.txt")
}

func Solve(file string){
	// input, err := ioutil.ReadFile(file)
	// if(err != nil){
	// 	panic(err)
	// }
	// lines := strings.Split(string(input), "\n")
	// fmt.Println(lines[0])
	startNums := []int{20,0,1,11,6,3}
	p1 := SolvePartOne(startNums)
	fmt.Println(p1)
	// p2 := SolvePartTwo(lines)
	// fmt.Println(p2)
}

func SolvePartOne(targetNum int, startNums []int) int{
	turnSpoken := make(map[int]int)

	for i, num := range startNums{
		turnSpoken[num] = i
	}

	spokenDiff := 0 //only true if no duplicates in startnums
	for i := len(startNums); i <= targetNum-2; i++{
		lastSpoken, exists := turnSpoken[spokenDiff]
		turnSpoken[spokenDiff] = i
		if(exists){
			spokenDiff = i-lastSpoken
		} else {
			spokenDiff = 0
		}

		if(i >= 29999995){
			fmt.Println(spokenDiff)
		}
	}
	return spokenDiff
}
