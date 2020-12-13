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
	target, _:= strconv.Atoi(lines[0])
	bussesAndXes := strings.Split(lines[1], ",")
	busses := make([]int, 0)
	for _, busOrX := range bussesAndXes{
		bus, err := strconv.Atoi(busOrX)
		if(err == nil){
			busses = append(busses, bus)
		}
	}
	bussesToWaitTimes := GetBussesToWaitingTimesFromTimestamp(target, busses)
	minWait := target
	earliestBus := -1
	for bus, wait := range bussesToWaitTimes{
		if(wait < minWait){
			minWait = wait
			earliestBus = bus
		}
	}
	p1 := earliestBus * minWait
	fmt.Println(p1)
	p2 := SolvePartTwo(bussesAndXes)
	fmt.Println(p2)
}

func GetBussesToWaitingTimesFromTimestamp(timestamp int, busses []int)map[int]int{
	bussesToWaitTimes := make(map[int]int)
	for _, bus := range busses{
		timeUnderWay := timestamp % bus
		waitingTime := bus - timeUnderWay
		bussesToWaitTimes[bus] = waitingTime
	}
	return bussesToWaitTimes
}

func SolvePartTwo(bussesOrXes []string) int{
	previous := 1
	timestamp := 0
	for i, busOrX := range bussesOrXes{
		bus, err := strconv.Atoi(busOrX)
		if(err == nil){
			var nextstamp int
			for j := 0; true; j++{
				nextstamp = timestamp + previous*j
				if(nextstamp % bus == Mod(bus-i, bus)){
					break
				}
			}
			previous = previous * bus // Should be LCM if the input weren't all primes I think
			timestamp = nextstamp
		}
	}
	return timestamp
}

func Mod(a, m int) int{
	remainder := a % m
	if(remainder < 0){
		remainder += m
	}
	return remainder
}