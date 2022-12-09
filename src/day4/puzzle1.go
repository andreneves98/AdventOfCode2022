package day4

import (
	"advent-of-code/utils"
	"log"
	"strconv"
	"strings"
)

func RunPuzzle1() {
	// Each pair represents the cleaning sections of 2 elves
	listOfPairs, err := utils.ReadFile("./src/day4/input.txt")
	if err != nil {
		log.Printf("failed to read file: %v", err)
		return
	}

	var numberOfFullOverlaps int
	for _, pair := range listOfPairs {
		sections := strings.Split(pair, ",")
		firstElfSectionRange := strings.Split(sections[0], "-")
		secondElfSectionRange := strings.Split(sections[1], "-")
		smallerRange, biggerRange := getBiggerRangeExtremities(firstElfSectionRange, secondElfSectionRange)

		if totallyContains(smallerRange, biggerRange) {
			numberOfFullOverlaps++
		}
	}

	log.Printf("Result: %d", numberOfFullOverlaps)
}

func getBiggerRangeExtremities(firstElfSection, secondElfSection []string) (smallerRange, biggerRange []int) {
	firstSecLower, _ := strconv.Atoi(firstElfSection[0])
	firstSecHigher, _ := strconv.Atoi(firstElfSection[1])
	secondSecLower, _ := strconv.Atoi(secondElfSection[0])
	secondSecHigher, _ := strconv.Atoi(secondElfSection[1])

	if firstSecHigher-firstSecLower <= secondSecHigher-secondSecLower {
		return []int{firstSecLower, firstSecHigher}, []int{secondSecLower, secondSecHigher}
	}

	return []int{secondSecLower, secondSecHigher}, []int{firstSecLower, firstSecHigher}
}

func totallyContains(smallerRange, biggerRange []int) bool {
	return smallerRange[0] >= biggerRange[0] && smallerRange[1] <= biggerRange[1]
}
