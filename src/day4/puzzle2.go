package day4

import (
	"advent-of-code/utils"
	"log"
	"strings"
)

func RunPuzzle2() {
	// Each pair represents the cleaning sections of 2 elves
	listOfPairs, err := utils.ReadFile("./src/day4/input.txt")
	if err != nil {
		log.Printf("failed to read file: %v", err)
		return
	}

	var totalOverlaps int
	for _, pair := range listOfPairs {
		sections := strings.Split(pair, ",")
		firstElfSectionRange := strings.Split(sections[0], "-")
		secondElfSectionRange := strings.Split(sections[1], "-")
		smallerRange, biggerRange := getBiggerRangeExtremities(firstElfSectionRange, secondElfSectionRange)

		if overlaps(smallerRange, biggerRange) {
			totalOverlaps++
		}
	}

	log.Printf("Result: %d", totalOverlaps)
}

func overlaps(smallerRange, biggerRange []int) bool {
	return smallerRange[0] <= biggerRange[1] && smallerRange[1] >= biggerRange[0]
}
