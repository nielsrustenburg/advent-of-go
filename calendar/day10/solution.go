package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"sort"
)

func main(){
	Solve("input.txt")
}

func Solve(file string){
	input, err := ioutil.ReadFile(file)
	if(err != nil){
		panic(err)
	}
	adapters := make([]int, 0)
	lines := strings.Split(string(input), "\n")
	for _, line := range lines{
		adapter, _ := strconv.Atoi(line)
		adapters = append(adapters, adapter)
	}

	p1, adapters := SolvePartOne(adapters)
	fmt.Println(p1)
	p2 := SolvePartTwo(adapters)
	fmt.Println(p2)
}

func SolvePartOne(adapters []int) (int, []int){
	adapters = append([]int{0}, adapters...) //Add outlet to the chain
	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3) //Add your device to the chain
	diffs := []int {0,0,0,0}

	for i, adapter := range adapters[1:]{
		diff := adapter - adapters[i]
		diffs[diff] = diffs[diff]+1
	}

	return diffs[1] * diffs[3], adapters
}

func SolvePartTwo(adapters []int)int{
	cache := make(map[int]int)
	cache[len(adapters)-1] = 1 //The device has only one way of connecting to the next device (none)
	return SolveDynamically(0, adapters, cache)
}

func SolveDynamically(index int, adapters []int, cache map[int]int)int{
	cachedResult, exists := cache[index]
	if(exists){
		return cachedResult
	}
	result := 0
	for i:= 1; i <= 3; i++{
		fits := index+i < len(adapters) && adapters[index+i] - adapters[index] <= 3
		if(fits){
			result = result + SolveDynamically(index+i, adapters, cache)
		} else {
			break
		}
	}
	cache[index] = result
	return result
}