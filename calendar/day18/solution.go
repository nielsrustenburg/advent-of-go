package main

import (
	"fmt"
	"strings"
	"io/ioutil"
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

	p1 := 0 
	p2 := 0
	for _, line := range lines{
		exp := ParseLine(line)
		p1 += exp.Solve()
		exp2 := ParseLine(Rewrite(line))
		p2 += exp2.Solve()
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func ParseLine(line string) Solvable{
	if(len(line) == 1){
		val, _:= strconv.Atoi(line)
		return Value{val}
	}

	var right Solvable
	var opIndex int
	if(line[len(line)-1] == ")"[0]){
		depth := 0
		for i := len(line)-1; i >= 0; i--{
			if(line[i] == ")"[0]){
				depth++
			}
			if(line[i] == "("[0]){
				depth--
				if(depth == 0){
					right = ParseLine(line[i+1:len(line)-1])
					opIndex = i - 2
					break
				}
			}
		}
	} else {
		val, _ := strconv.Atoi(line[len(line)-1:len(line)])
		right = Value{val}
		opIndex = len(line)-1 - 2
	}

	if(opIndex < 2){
		return right
	}

	left := ParseLine(line[0:opIndex-1])
	return Expression{line[opIndex:opIndex+1], left, right}
}

type Expression struct{
	Operator string
	Left Solvable
	Right Solvable
}

type Solvable interface{
	Solve() int
}

type Value struct {
	Val int
}

func(val Value) Solve() int{
	return val.Val
}

func(exp Expression) Solve() int{
	switch exp.Operator {
	case "+":
		return exp.Left.Solve() + exp.Right.Solve()
	case "*":
		return exp.Left.Solve() * exp.Right.Solve()
	}
	panic("WHY AM I HERE????")
}

func Rewrite(line string) string {
	for{
		plusFound := false
		for i, char := range line{
			if(char == rune("+"[0])){
				plusFound = true
				//Replace + with Minus for now
				line = line[0:i] + "-" + line[i+1:len(line)]

				//Add Right Parenthesis
				if(line[i+2] == "("[0]){
					//Find matching closing parenthesis and add one next to it
					depth := 1
					for j := i+3; j < len(line); j++{
						if(line[j] == ")"[0]){
							depth--
							if(depth == 0){
								// FOUND IT
								line = line[0:j+1] + ")" + line[j+1:len(line)]
								break
							}
						}

						if(line[j] == "("[0]){
							depth++
						}
					}
				} else {
					line = line[0:i+3] + ")" + line[i+3:len(line)] 
				}

				//Add Left Parenthesis
				if(line[i-2] == ")"[0]){
					//Find matching opening parenthesis and add one next to it
					depth := 1
					for j := i-3; j >= 0; j--{
						if(line[j] == "("[0]){
							depth--
							if(depth == 0){
								line = line[0:j] + "(" + line[j:len(line)]
								break
							}
						}

						if(line[j] == ")"[0]){
							depth++
						}
					}
				} else {
					line = line[0:i-2] + "(" + line[i-2:len(line)] 
				}
				break
			}
		}
		if(!plusFound){
			break
		}
	}
	result :=  strings.Replace(line, "-", "+", -1)
	return result
}
