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
	seatNrs := make(map[int]bool)
	min := 1023
	max := 0
	for _, line := range lines{
		seatNr := ToSeatNumber(line)
		seatNrs[seatNr] = true
		if(seatNr > max) {
			max = seatNr
		}
		if(seatNr < min){
			min = seatNr
		}
	}

	fmt.Println(max)
	part2 := FindMissingNr(min, max, seatNrs)
	fmt.Println(part2)
}

func FindMissingNr(min, max int, nrs map[int]bool)int{
	for i := min+1; i < max; i++ {
		_, ok := nrs[i]
		if(!ok){
			return i
		}
	}
	panic("no missing nr found")
}

func ToSeatNumber(bpass string)int{
	bpass = strings.ReplaceAll(bpass,"F", "0")
	bpass = strings.ReplaceAll(bpass,"B", "1")
	bpass = strings.ReplaceAll(bpass,"L", "0")
	bpass = strings.ReplaceAll(bpass,"R", "1")
	ans, _ := strconv.ParseInt(bpass, 2, 11)
	return int(ans)
}

