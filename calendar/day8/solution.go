package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	//"utils/utils"
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

	hh := MakeHandheld(lines)
	_, reached := hh.FindLoop()
	fmt.Println(hh.Accumulator)

	part2 := FindCorruption(hh, reached)

	if(hh.Code[part2].Op == "jmp"){
		lines[part2] = "nop" + lines[part2][3:]
	} else{
		lines[part2] = "jmp" + lines[part2][3:]
	}
	hh2 := MakeHandheld(lines)
	hh2.StepUntilEnd()

	fmt.Println(hh2.Accumulator)

	//utils.HelloWorld()
}

func MakeHandheld(input []string) Handheld{
	instructions := make([]Instruction,0)
	for _, line := range input{
		inst := ParseInstruction(line)
		instructions = append(instructions, inst)
	}
	handheld := Handheld{instructions,0,0}
	return handheld
}

func(handheld *Handheld) FindLoop() (int, map[int]int){
	count := 0
	pointersReached := make(map[int]int)
	for true{
		pointersReached[handheld.Pointer] = 1
		handheld.Step()
		count++
		_, loopFound := pointersReached[handheld.Pointer]
		if(loopFound){
			// keys := make([]int, len(pointersReached))
			// i := 0
			// for pointer := range pointersReached {
			// 	keys[i] = pointer
			// 	i++
			// }
			return count, pointersReached
		}
	}
	panic("WHY AM I OUTSIDE AN INFINITE LOOP")
}

func FindCorruption(handheld Handheld, reachedIndices map[int]int) int{
	landingzone := FindLandingZone(handheld)

	//If landingzone-blocker has been reached flip to nop --> done
	landingBlocker := len(handheld.Code)-len(landingzone)-1
	_, lastJmpReached := reachedIndices[landingBlocker]
	if(lastJmpReached){
		return landingBlocker
	}

	//For every unreached jmp, note their destination
	jmpDestinations, nopDestinations := FindJumpDestinations(handheld, reachedIndices)

	//Keep track of jmps that enter landing zone
		//for each --> add to landingzone + previous instructions until next jmp
		//(check if any jmps have this new landingzone as their destination, add to pile)
		//if jmp was reached --> flip it --> done
		//if jmp was not reached continue to next jmp
	blocker := ExpandLandingZone(jmpDestinations, landingzone, reachedIndices, handheld)
	if(blocker != -1){
		return blocker
	}

	//Look for nops that can reach landingzone by flipping directly (It has to reach the landingzone or indirectly by reaching unreached code that reaches landingzone.....)
	for nope, dest := range nopDestinations{
		_, destPartOfLandingZone := landingzone[dest]
		if(destPartOfLandingZone){
			return nope
		}
	}

	panic("WE HAVE BEEN LIED TO, NO SOLUTION EXISTS (or I implemented this poorly)")
}

func FindLandingZone(handheld Handheld) map[int]int{
	landingzone := make(map[int]int)
	for i := len(handheld.Code)-1; i >= 0; i--{
		inst := handheld.Code[i]
		if(inst.Op == "jmp" && inst.Arg < 0){
			break
		}
		landingzone[i] = 1
	}
	return landingzone
}

func FindJumpDestinations(handheld Handheld, reachedIndices map[int]int) (map[int]int, map[int]int){
	jmpDestinations := make(map[int]int)
	nopDestinations := make(map[int]int)

	for i := 0; i < len(handheld.Code); i++{
		_, reached := reachedIndices[i]
		inst := handheld.Code[i]
		if(!reached){
			dest := i + inst.Arg
			if(inst.Op == "jmp"){
				jmpDestinations[i] = dest
			}
			if(inst.Op == "nop"){
				nopDestinations[i] = dest
			}
		}
	}
	return jmpDestinations, nopDestinations
}

func ExpandLandingZone(jmpDestinations, landingzone, reachedIndices map[int]int, hh Handheld) int{
	expandables := make([]int,0)
	for jmpIndex, dest := range jmpDestinations{
		_, alreadyInLandingzone := landingzone[jmpIndex] 
		_, destInLandingzone := landingzone[dest]
		if(!alreadyInLandingzone && destInLandingzone){
			expandables = append(expandables,jmpIndex)
		}
	}

	if(len(expandables) > 0){
		//expand landing zone
		for _, jmpIndex := range expandables{
			prevJumpReached := false
			count := 0
			for !prevJumpReached{
				landingzone[jmpIndex-count] = 1
				count++
				if(hh.Code[jmpIndex-count].Op == "jmp"){
					_, previousReached := reachedIndices[jmpIndex-count]
					if(previousReached){
						return jmpIndex-count
					}
					prevJumpReached = true
				}
			}
		}
		return ExpandLandingZone(jmpDestinations, landingzone, reachedIndices, hh)
	}
	return -1
}

func(handheld *Handheld) StepUntilEnd(){
	done := false
	for !done{
		handheld.Step()
		done = handheld.Pointer >= len(handheld.Code)
	}
}

type Handheld struct {
	Code        []Instruction
	Accumulator int
	Pointer     int
}

type Instruction struct {
	Op  string
	Arg int
}

func (handheld *Handheld) Step() {
	// TODO: Let the Instructions contain the relevant functions instead of strings?
	currentInstruction := handheld.Code[handheld.Pointer]
	switch op := currentInstruction.Op; op {
	case "acc":
		handheld.Acc(currentInstruction.Arg)
	case "jmp":
		handheld.Jmp(currentInstruction.Arg)
	case "nop":
		handheld.Nop(currentInstruction.Arg)
	}
}

func (handheld *Handheld) Nop (_ int) {
	handheld.Pointer = handheld.Pointer + 1
}


func (handheld *Handheld) Acc(increment int) {
	handheld.Accumulator += increment
	handheld.Pointer = handheld.Pointer + 1
}


func (handheld *Handheld) Jmp(diff int) {
	handheld.Pointer = handheld.Pointer + diff
}

func ParseInstruction(input string) Instruction{
	split := strings.Split(input, " ")
	inc, _ := strconv.Atoi(split[1])
	instr := Instruction {split[0], inc}
	return instr
}