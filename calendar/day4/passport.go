package main

import (
	"encoding/hex"
	"strconv"
)

type Passport struct{
	Fields map[string]Field
}

type filter func(Passport)bool

func ContainsRequiredFields(passport Passport) bool{
	required := []string {"byr","iyr","eyr","hgt","hcl","ecl","pid"}
	for _, name := range required{
		_, ok := passport.Fields[name]
		if(!ok){
			return false
		}
	}
	return true
}

func AllFieldsValid(passport Passport) bool{
	for _, field := range passport.Fields{
		if(!field.IsValid()){
			return false
		}
	}
	return true
}

func FilterPassports(passports []Passport, f filter) []Passport{
	valid := make([]Passport, 0)
	for _, passport := range passports{
		if(f(passport)){
			valid = append(valid, passport)
		}
	}
	return valid
}

type Field struct {
	Name string
	Data string
}

func(field Field) IsValid() bool{
	switch field.Name {
	case "byr":
		return field.IsValidYr(1920,2002)
	case "iyr":
		return field.IsValidYr(2010,2020)
	case "eyr":
		return field.IsValidYr(2020,2030)
	case "hgt":
		return field.IsValidHgt()
	case "hcl":
		return field.IsValidRgb()
	case "ecl":
		return field.IsValidEcl()
	case "pid":
		return field.IsValidPid()
	default:
		return true
	}
}

func(field Field) IsValidYr(lower, upper int) bool{
	val, _ := strconv.Atoi(field.Data)
	return IsBtwn(val,lower,upper)
}

func(field Field) IsValidHgt() bool{
	val, _ := strconv.Atoi(field.Data[:len(field.Data)-2])
	unit := field.Data[len(field.Data)-2:]
	if(unit == "cm"){
		return IsBtwn(val,150,193)
	}
	if(unit == "in"){
		return IsBtwn(val,59,76)
	}
	return false
}

func(field Field) IsValidRgb() bool{
	if(field.Data[0] == "#"[0]){
		hexa, _ := hex.DecodeString(field.Data[1:])
		return len(hexa) == 3
	}
	return false
}

func(field Field) IsValidEcl() bool{
	colors := map[string]int {"amb":1, "blu":1, "brn":1, "gry":1, "grn":1, "hzl":1, "oth":1}
	_, ok := colors[field.Data]
	return ok
}

func(field Field) IsValidPid() bool{
	_, err := strconv.Atoi(field.Data)
	return (err == nil && len(field.Data) == 9)
}

func IsBtwn(val, lower, upper int) bool{
	return val >= lower && val <= upper
}