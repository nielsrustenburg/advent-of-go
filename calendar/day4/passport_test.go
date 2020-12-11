package main

import (
	"strings"
	"testing"
)

func TestIsvalid(t *testing.T){
	testcases := []struct {
		input Field
		expected bool
	}{
		{Field{"byr", "2002"}, true},
		{Field{"byr", "2003"}, false},
		{Field{"hgt", "60in"}, true},
		{Field{"hgt", "190cm"}, true},
		{Field{"hgt", "190in"}, false},
		{Field{"hgt", "190"}, false},
		{Field{"hcl", "#123abc"}, true},
		{Field{"hcl", "#123abz"}, false},
		{Field{"hcl", "123abc"}, false},
		{Field{"ecl", "brn"}, true},
		{Field{"ecl", "wat"}, false},
		{Field{"pid", "000000001"}, true},
		{Field{"pid", "0123456789"}, false},
	}

	for _, tcase := range testcases{
		got := tcase.input.IsValid()
		if(got != tcase.expected){
			t.Errorf("%v ... expected: %v got %v", tcase.input, tcase.expected, got)
		}
	}
}



func TestCreatePassports(t *testing.T){
	lines := strings.Split(string(testpassports1), "\n")
	passports := CreatePassports(lines)
	got := len(passports)
	expected := 4
	if(got != expected){
		t.Errorf("expected: %v got %v \n %v", expected, got, passports)
	}
}

func TestContainsRequiredFields(t *testing.T){
	lines := strings.Split(string(testpassports1), "\n")
	passports := CreatePassports(lines)
	valid := FilterPassports(passports, ContainsRequiredFields)
	got := len(valid)
	expected := 2
	if(got != expected){
		t.Errorf("expected: %v got %v \n %v", expected, got, valid)
	}
}

func TestAllFieldsValidWithInvalid(t *testing.T){
	lines := strings.Split(string(invalidFields), "\n")
	shouldBeInvalid := CreatePassports(lines)
	shouldBeEmpty := FilterPassports(shouldBeInvalid, AllFieldsValid)
	got := len(shouldBeEmpty)
	expected := 0
	if(got != expected){
		t.Errorf("expected: %v got %v \n %v", expected, got, shouldBeEmpty)
	}
}

func TestAllFieldsValidWithValid(t *testing.T){
	lines := strings.Split(string(validFields), "\n")
	shouldBeValid := CreatePassports(lines)
	shouldBeFull := FilterPassports(shouldBeValid, AllFieldsValid)
	got := len(shouldBeFull)
	expected := 4
	if(got != expected){
		t.Errorf("expected: %v got %v \n %v", expected, got, shouldBeFull)
	}
}

func TestGeertensTestCases(t *testing.T){
	lines := strings.Split(string(lotsOfPassports), "\n")
	alines := strings.Split(string(lotsOfAnswersInDutchThanksGeerten), "\n")
	passports := CreatePassports(lines)

	for i, passport := range passports{
		got := ContainsRequiredFields(passport) && AllFieldsValid(passport)
		expected := alines[i] == "goed"
		if(got != expected){
			t.Errorf("expected: %v got %v \n %v", expected, got, passport)
		}
	}
}


const testpassports1 = `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`

const invalidFields = `
eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007`

const validFields = `pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719`
