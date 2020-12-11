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
	validCount := 0 
	valid2Count := 0
	for _, line := range lines{
		if(IsValidPassword(line)){
			validCount++
		}
		if(IsValidPasswordTwo(line)){
			valid2Count++
		}
	}
	fmt.Println(validCount)
	fmt.Println(valid2Count)
}

func IsValidPassword(pwline string) bool {
	policy_pw := strings.Split(pwline,": ")
	policy := CreatePolicy(policy_pw[0])
	password := policy_pw[1]
	occurrences := strings.Count(password,policy.Character)
	return occurrences >= policy.Min && occurrences <= policy.Max
}

func IsValidPasswordTwo(pwline string) bool{
	policy_pw := strings.Split(pwline,": ")
	policy := CreatePolicy(policy_pw[0])
	password := policy_pw[1]
	one := (string(password[policy.Min-1]) ==  policy.Character)
	two := (string(password[policy.Max-1]) == policy.Character)
	return (one || two) && !( one && two)
}

func CreatePolicy(pline string) Policy {
	splitChar := strings.Split(pline, " ")
	char := splitChar[1]
	splitMinMax := strings.Split(splitChar[0], "-")
	min, _ := strconv.Atoi(splitMinMax[0])
	max, _ := strconv.Atoi(splitMinMax[1])
	return Policy{min,max,char}
}

type Policy struct {
	Min int
	Max int
	Character string
}