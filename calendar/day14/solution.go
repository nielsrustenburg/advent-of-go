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
	p1 := SolvePartOne(lines)
	fmt.Println(p1)
	p2 := SolvePartTwo(lines)
	fmt.Println(p2)
}

//Too low: 9438959851756
func SolvePartOne(lines []string)int64{
	masks := make(map[int]Mask)
	mem := make(map[int64]int64)
	writes := make([]Write, len(lines))
	for i, line := range lines{
		lr := strings.Split(line, " = ")
		left := lr[0]
		right := lr[1]
		if(left == "mask"){
			masks[i] = ParseMask(right)
		} else {
			writes[i] = ParseWrite(left,right)
		}
	}

	mask := Mask{0,0}
	for i, write := range writes{
		nextmask, ok := masks[i]
		if(ok){
			mask = nextmask
		}else {
			mem[write.To] = (write.Val & mask.And) | mask.Or
		}
	}

	sum := int64(0)
	for _, val := range mem{
		sum += val
	}
	return sum
}

func SolvePartTwo(lines []string)int64{
	masks := make(map[int]Mask2)
	mem := make(map[int64]int64)
	writes := make([]Write, len(lines))
	for i, line := range lines{
		lr := strings.Split(line, " = ")
		left := lr[0]
		right := lr[1]
		if(left == "mask"){
			masks[i] = ParseMask2(right)
		} else {
			writes[i] = ParseWrite(left,right)
		}
	}

	mask := Mask2{0,0,make([]int64,0)}
	for i, write := range writes{
		nextmask, ok := masks[i]
		if(ok){
			mask = nextmask
		}else {
			prepared := (int64(write.To) | mask.Or) &^ mask.Clear
			for _, floatmask := range mask.All{
				address := prepared | floatmask
				mem[address] = write.Val
			}
		}
	}

	sum := int64(0)
	for _, val := range mem{
		sum += val
	}
	return sum
}

func ParseMask(value string) Mask{
	ands := strings.Replace(value, "X", "1", -1)
	ors := strings.Replace(value, "X", "0", -1)
	and, _ := strconv.ParseInt(ands, 2, 64)
	or, _ := strconv.ParseInt(ors, 2, 64)
	return Mask {and, or}
}

func ParseWrite(left,right string) Write{
	to, _ := strconv.ParseInt(left[4:len(left)-1],10,64)
	val, _ := strconv.ParseInt(right, 10, 64)
	return Write{to,val}
}

func ParseMask2(value string) Mask2{
	ors := strings.Replace(value, "X", "0", -1)
	or, _ := strconv.ParseInt(ors, 2, 64)
	floats := make([]int64, 0)
	floater := int64(1)<<35
	clear := int64(0)
	for _, char := range value{
		if(char == []rune("X")[0]){
			floats = append(floats, floater)
			clear += floater
		}
		floater = floater>>1
	} 
	if(floater != 0){
		panic("SANITY CHECK")
	}

	all := AllFloats(floats)
	return Mask2 {or,clear, all}
}

func AllFloats(floats []int64) []int64{
	if(len(floats) == 0){
		return []int64 {0}
	}
	result := AllFloats(floats[1:])
	for _, n := range result{
		result = append(result, n + floats[0])
	}
	return result
}

type Mask struct{
	And int64
	Or int64
}

type Write struct{
	To int64
	Val int64
}

type Mask2 struct {
	Or int64
	Clear int64
	All []int64
}