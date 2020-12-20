package main

import (
	"fmt"
	"strings"
	"strconv"
	"io/ioutil"
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
	fields, myticket, othertickets := ParseLines(lines)
	p1, validTickets := CalculateTicketScanningErrorRate(fields, othertickets)
	fmt.Println(p1)
	sortedFields := SortTicketFields(fields, validTickets)
	p2 := CalculateDepartureValue(myticket, sortedFields)
	fmt.Println(p2)
}

func CalculateDepartureValue(ticket []int, fields map[int]Field) int{
	val := 1
	for i, field := range fields{
		if(strings.Contains(field.Name, "departure")){
			val = val * ticket[i]
		}
	}
	return val
}

func SortTicketFields(fields map[string]Field, tickets [][]int) map[int]Field{
	potentialFieldsPerColumn := FindPotentialFieldsPerColumn(fields, tickets)
	colsToFields := make(map[int]Field)

	for {
		found := -1
		for col, potentialFields := range potentialFieldsPerColumn{
			_, alreadyFound := colsToFields[col]
			if(!alreadyFound && len(potentialFields) == 1){
				found = col
				for _, foundField := range potentialFields{
					colsToFields[col] = foundField
				}
				break
			}
		}
		if(found > -1){
			for i, potentialFields := range potentialFieldsPerColumn{
				_, isPot := potentialFields[colsToFields[found].Name]
				if(isPot && i != found){
					delete(potentialFields, colsToFields[found].Name)
				}
			}
		}
		if(found == -1){
			break
		}
	}

	fmt.Println(colsToFields)
	return colsToFields
}

func FindPotentialFieldsPerColumn(fields map[string]Field, tickets [][]int) []map[string]Field{

	allFields := make([]Field,0)
	for _, field := range fields{
		allFields = append(allFields, field)
	}

	potFieldsPerCol := make([]map[string]Field, len(allFields))
	for i := 0; i < len(allFields); i++{
		allFieldsMapped := make(map[string]Field)
		for _, field := range allFields{
			allFieldsMapped[field.Name] = field
		}
		potFieldsPerCol[i] = allFieldsMapped
	}

	for _, ticket := range tickets{
		for col, val := range ticket {
			removePotFields := make([]string, 0)
			for fi, field := range potFieldsPerCol[col]{
				fits := false
				for _, rng := range field.Ranges{
					if(rng.Fits(val)){
						fits = true
						break
					}
				}
				if(!fits){
					removePotFields = append(removePotFields, fi)
				}
			}

			for _, remove := range removePotFields{
				delete(potFieldsPerCol[col], remove)
			}
		}
	}

	return potFieldsPerCol
}

func(rng Range) Fits(val int) bool{
	return val >= rng.Low && val <= rng.Up
}

func CalculateTicketScanningErrorRate(fields map[string]Field, tickets [][]int) (int, [][]int){
	errorRate := 0
	validTickets := make([][]int, 0)
	for _, ticket := range tickets{
		ticketValid := true
		for _, tval := range ticket{
			valid := false
			for _, field := range fields{
				if(!valid){
					// if it fits we can skip to the next tval
					for _, rng := range field.Ranges{
						if(rng.Fits(tval)){
							valid = true
							break
						}
					}
				}
				if(valid){
					break
				}
			}
			if(!valid){
				ticketValid = false
				errorRate += tval
			}
		}
		if(ticketValid){
			validTickets = append(validTickets, ticket)
		} 
	}
	return errorRate, validTickets
}

func ParseLines(lines []string) (map[string]Field, []int, [][]int) {
	mode := 0
	fields := make(map[string]Field)
	var myticket []int
	othertickets := make([][]int, 0)
	for _, line := range lines {
		switch line {
		case "your ticket:":
			mode = 1
		case "nearby tickets:":
			mode = 2
		}
		if(len(line) > 0 && !strings.Contains(line, "ticket")){
			switch mode {
			case 0:
				fld := ParseField(line)
				fields[fld.Name] = fld
			case 1:
				myticket = ParseTicket(line)
			case 2:
				otherticket := ParseTicket(line)
				othertickets = append(othertickets, otherticket)
			}
		}
	}
	return fields, myticket, othertickets
}

func ParseTicket(line string) []int{
	values := strings.Split(line, ",")
	ticketvalues := make([]int, 0)
	for _, val := range values {
		intval, _ := strconv.Atoi(val)
		ticketvalues = append(ticketvalues, intval)
	}
	return ticketvalues
}

func ParseField(line string)Field{
	nameAndRanges := strings.Split(line, ": ")
	name := nameAndRanges[0]
	ranges := strings.Split(nameAndRanges[1], " or ")
	frngs := make([]Range, 0)
	for _, r := range ranges {
		rng := parseRange(r)
		frngs = append(frngs, rng)
	}
	return Field{name, frngs}
}

func parseRange(rangeLine string)Range{
	lwrAndUpr := strings.Split(rangeLine, "-")
	lower, _ := strconv.Atoi(lwrAndUpr[0])
	upper, _ := strconv.Atoi(lwrAndUpr[1])
	return Range{lower, upper}
}

type Field struct {
	Name string
	Ranges []Range
}

type Range struct {
	Low int
	Up int
}
