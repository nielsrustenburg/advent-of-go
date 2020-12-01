package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
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
	var nums []int

	for _, s := range lines{
		n, _ := strconv.Atoi(s)
		nums = append(nums,n)
	}

	a,b := FindSumNums(2020, nums)
    p1 := a*b
	fmt.Printf("Part one: %v", p1)
}

func FindSumNums(target int, nums []int) (int, int) {
	m := make(map[int]int)

	for i:=0 ; i < len(nums); i++{
		next := nums[i]
		other := target - next

		_, ok := m[other]
		if(ok){
			return next,other
		}

		m[next] = other
	}
	return 0, target
}