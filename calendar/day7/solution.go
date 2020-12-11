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

	outersToInners, innersToOuters := BuildSolutionTree(lines)

	p1 := BagsContainingShinyGold(innersToOuters)
	p2 := BagsInsideBag("shiny gold bags", outersToInners)
	fmt.Println(p1)
	fmt.Println(p2)
}

func BagsInsideBag(bag string, outersToInners map[string]map[string]int) int{
	sum := 0
	for innerbag, quantity := range outersToInners[bag]{
		sum = sum + quantity + quantity * BagsInsideBag(innerbag, outersToInners)
	}
	return sum
}

func BagsContainingShinyGold(innersToOuters map[string]map[string]int) int{
	bagsToCheck := []string {"shiny gold bags"}
	bagsChecked := make(map[string]bool)
	for len(bagsToCheck) > 0 {
		checkBag := bagsToCheck[0]
		_, alreadyChecked := bagsChecked[checkBag]
		if(!alreadyChecked){
			newBagsToCheck, _ := innersToOuters[checkBag]
			for k, _ := range newBagsToCheck{
				bagsToCheck = append(bagsToCheck, k)
			}
			bagsChecked[checkBag] = true
		}
		bagsToCheck = bagsToCheck[1:len(bagsToCheck)]
	}
	return len(bagsChecked)-1
}

func BuildSolutionTree(lines []string) (map[string]map[string]int, map[string]map[string]int){
	outerToInnerMap := make(map[string]map[string]int)
	innerToOuterMap := make(map[string]map[string]int)

	for _, line := range lines{
		outer, inner := ParseInputLine(line)
		outerToInnerMap[outer] = inner

		for innerbag, quantity := range inner{
			_, ok := innerToOuterMap[innerbag]
			if(!ok){
				innerToOuterMap[innerbag] = make(map[string]int)
			}
			innerToOuterMap[innerbag][outer] = quantity
		}
	}

	return outerToInnerMap, innerToOuterMap
}

func ParseInputLine(line string) (string, map[string]int){
	noPeriod := line[0:(len(line)-1)]
	splitContains := strings.Split(noPeriod, " contain ") 
	outerBag := splitContains[0]
	innerBags := strings.Split(splitContains[1], ", ")
	bagsToQuantities := make(map[string]int) 
	for _, b := range innerBags{
		space := strings.Index(b, " ")
		quantity, _ := strconv.Atoi(b[0:space])
		bag := b[space+1:len(b)]
		if(quantity == 1){
			bag = bag + "s"
		}
		bagsToQuantities[bag] = quantity
	}
	return outerBag, bagsToQuantities
}