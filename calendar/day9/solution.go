package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
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
	p1 := SolvePartOne(lines, 25)
	fmt.Println(p1)
	p2 := SolvePartTwo(lines,p1)
	fmt.Println(p2)
}

func SolvePartOne(lines []string, n int) int{
	intlines := make([]int, 0)
	for _, line := range lines{
		num, err := strconv.Atoi(line)
		if(err != nil){
			panic(err)
		}
		intlines = append(intlines, num)
	}

	return FindWeakness(intlines, n)
}

func SolvePartTwo(lines []string, p1 int) int{
	intlines := make([]int, 0)
	for _, line := range lines{
		num, err := strconv.Atoi(line)
		if(err != nil){
			panic(err)
		}
		intlines = append(intlines, num)
	}
	lower, upper := FindContiguousSumSet(intlines,p1)

	smallest := p1
	largest := 0
	for i := lower; i <= upper; i++{
		if(intlines[i] < smallest){
			smallest = intlines[i]
		}
		if(intlines[i] > largest){
			largest = intlines[i]
		}
	}
	return smallest+largest
}

func FindWeakness(sequence []int, n int) int{
	fmt.Println(sequence)
	for i := n; i <= len(sequence); i++{
		smallSlice := sequence[i-n:i]
		isSum := IsSumOfTwoPredecessors(smallSlice, sequence[i])
		if(!isSum){
			return sequence[i]
		}
	}
	return -1
}

func IsSumOfTwoPredecessors(predecessors []int, num int) bool{
	for i, p1 := range predecessors{
		for j, p2 := range predecessors{
			if(i != j && p1 + p2 == num){
				return true
			}
		}
	}
	return false
}

func FindContiguousSumSet(nums []int, sumNum int) (int,int){
	lower := 0
	upper := 1
	sum := nums[0] + nums[1]
	for true {
		fmt.Println(upper)
		if(sum == sumNum){
			fmt.Printf("low: %v    high: %v \n", lower, upper)
			return lower, upper
		}
		if(sum < sumNum){
			upper++
			sum = sum + nums[upper]
		} else{
			sum = sum - nums[lower]
			lower++
		}
	}
	panic("WEE WOO SHOULDN'T BE HERE")
}