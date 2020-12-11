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
	passports := CreatePassports(lines)
	part1pps := FilterPassports(passports, ContainsRequiredFields)
	fmt.Println(len(part1pps))
	part2pps := FilterPassports(passports, AllFieldsValid)
	fmt.Println(len(part2pps)) // Too high 250
}

func CreatePassports(lines []string) []Passport{
	passports := make([]Passport, 0)

	var currentPassport = Passport {make(map[string]Field)}
	for _, line := range lines{
		if(len(line) > 0){
			fs := strings.Split(line, " ")
			for _, f := range fs {
				fSpl := strings.Split(f, ":")
				field := Field {fSpl[0], fSpl[1]}
				currentPassport.Fields[field.Name] = field
			}
		} else {
			if(len(currentPassport.Fields) > 0){
				passports = append(passports, currentPassport)
			}
			currentPassport = Passport {make(map[string]Field)}
		}
	}
	if(len(currentPassport.Fields) > 0){
		passports = append(passports, currentPassport)
	}
	return passports
}
