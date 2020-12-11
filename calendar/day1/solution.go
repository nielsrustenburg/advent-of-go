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

	a,b := FindSumPairs(2020, nums)
    p1 := a*b
	fmt.Printf("Part one: %v\n", p1)

	c,d,e := FindSumTrios(2020,nums)
	p2 := c*d*e

	fmt.Printf("Part two: %v\n", p2)
}

func FindSumPairs(target int, nums []int) (int, int) {
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
	return 0, 0
}

func FindSumTrios(target int, nums []int) (int, int, int){
	for _, num := range nums{
		pairTarget := target - num
		a, b := FindSumPairs(pairTarget, nums)
		if(a != 0){
			return a,b,num
		}
	}
	return 0,0,0
}