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

	charCountsPerGroup := CountCharsPerGroup(lines)
	sum := 0
	for _, charCounts := range charCountsPerGroup{
		sum += len(charCounts)
	}
	fmt.Println(sum)

	peoplePerGroup := CountPeoplePerGroup(lines)

	sum2 := 0
	for i, charCounts := range charCountsPerGroup{
		for _, value := range charCounts{
			if(value == peoplePerGroup[i]){
				sum2++
			}
		}
	}

	fmt.Println(sum2)
}

func CountPeoplePerGroup(lines []string)[]int{
	result := make([]int, 0)
	lower := 0
	for i, line := range lines{
		if(len(line) == 0){
			result = append(result, i-lower)
			lower = i+1
		}
	}
	return result
}

func CountCharsPerGroup(lines []string) []map[rune]int{
	result := make([]map[rune]int, 0)
	currentMap := make(map[rune]int)
	for _, line := range lines{
		if(len(line) == 0){
			if(len(currentMap) > 0){
				result = append(result, currentMap)
				currentMap = make(map[rune]int)
			}
		}else {
			for _, bt := range line{
				currentVal, _ := currentMap[bt]
				currentMap[bt] = currentVal+1
			}
		}
	}
	return result
}