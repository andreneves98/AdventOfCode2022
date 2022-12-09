package day3

import (
	"advent-of-code/utils"
	"log"
	"strings"
)

func RunPuzzle2() {
	// Get rucksacks content
	rucksacks, err := utils.ReadFile("./src/day3/input.txt")
	if err != nil {
		log.Printf("failed to read file: %v", err)
		return
	}

	var badgesSum int
	for group := 0; group < len(rucksacks); group += 3 {
		// Grab 3 elves' rucksacks at a time
		groupRucksacks := rucksacks[group : group+3]
		// Split first elf item types
		firstElfItemTypes := strings.Split(groupRucksacks[0], "")

		// mutualItemTypes1st2nd is a slice with the item types that
		// both the 1st and 2nd elves of the group have in their rucksack
		var mutualItemTypes1st2nd []string
		for _, secondElfItemType := range strings.Split(groupRucksacks[1], "") {
			if contains(firstElfItemTypes, secondElfItemType) {
				mutualItemTypes1st2nd = append(mutualItemTypes1st2nd, secondElfItemType)
			}
		}

		var badge string
		for _, thirdElfItemType := range strings.Split(groupRucksacks[2], "") {
			if contains(mutualItemTypes1st2nd, thirdElfItemType) {
				badge = thirdElfItemType
			}
		}

		badgesSum += itemTypeToPriority([]rune(badge)[0])
	}

	log.Printf("Result: %d", badgesSum)
}
