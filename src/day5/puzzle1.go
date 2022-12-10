package day5

import (
	"advent-of-code/utils"
	"log"
	"regexp"
	"strconv"
)

var stacks [][]string

func RunPuzzle1() {
	input, err := utils.ReadFile("./src/day5/input.txt")
	if err != nil {
		log.Printf("failed to read file: %v", err)
		return
	}

	stacks = make([][]string, 9)
	for line := 0; input[line] != ""; line++ {
		log.Println(input[line])

		stackIdx := 0
		for c := 0; c < len(input[line]); c += 4 {
			if input[line][c] != 32 { // 32 is the ascii code for SPACE
				// log.Println("crate: ", input[line][c+1:c+2])
				stacks[stackIdx] = append(stacks[stackIdx], input[line][c+1:c+2])
			}
			stackIdx++
		}
	}

	// reverse each slice's order to form actual stacks
	for idx := range stacks {
		reverseSlice(stacks[idx])
	}

	log.Printf("Initial stacks: %v", stacks)
	// read the rearrangement procedure
	processRearrangementProcedure(input, stacks)
	log.Printf("Stacks after rearrangement: %v", stacks)

	// get the top crate on each stack at the end of the procedure
	var result string
	for idx := range stacks {
		result += stacks[idx][len(stacks[idx])-1]
		log.Printf("Top crate at stack %d: %s\n", idx+1, stacks[idx][len(stacks[idx])-1])
	}

	log.Printf("Result: %s", result)
}

// https://www.golangprograms.com/regular-expression-to-extract-numbers-from-a-string-in-golang.html
func processRearrangementProcedure(procedure []string, stacks [][]string) {
	for line := 10; line < len(procedure); line++ {
		var procedureData []int
		re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
		findAllString := re.FindAllString(procedure[line], -1)
		for _, elem := range findAllString {
			elem, err := strconv.Atoi(elem)
			if err != nil {
				log.Printf("failed to convert: %v", err)
				return
			}
			procedureData = append(procedureData, elem)
		}

		swap(procedureData[0], procedureData[1], procedureData[2])
	}
}

func swap(number, src, dest int) {
	for i := 0; i < number; i++ {
		crate := stacks[src-1][len(stacks[src-1])-1]
		stacks[src-1] = stacks[src-1][:len(stacks[src-1])-1]
		stacks[dest-1] = append(stacks[dest-1], crate)
	}
}

func reverseSlice(slice []string) {
	s := slice
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
