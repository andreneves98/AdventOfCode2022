package day3

import (
	"advent-of-code/utils"
	"log"
	"strings"
)

func RunPuzzle1() {
	// Get rucksacks content
	rucksacks, err := utils.ReadFile("./src/day3/input.txt")
	if err != nil {
		log.Printf("failed to read file: %v", err)
		return
	}

	var prioritiesSum int
	for _, rucksack := range rucksacks {
		firstComp := strings.Split(rucksack[:len(rucksack)/2], "")
		secondComp := strings.Split(rucksack[len(rucksack)/2:], "")
		duplicatedItemType := findDuplicatedItemTypes(firstComp, secondComp)
		prioritiesSum += itemTypeToPriority([]rune(duplicatedItemType)[0])
	}

	log.Printf("Result: %d", prioritiesSum)
}

// itemTypeToPriority returns the priority of an item type.
// We are converting the item type, which is a rune (Go's equivalent of char)
// to its ascii code. We subtract 96 or 38 to lowercases and uppercases,
// respectivelly, in order to map the priorities correctly to 1-26 and 27-52.
func itemTypeToPriority(itemType rune) int {
	// a-z
	if int(itemType) >= 97 && int(itemType) <= 122 {
		return int(itemType) - 96
		// A-Z
	} else if int(itemType) >= 65 && int(itemType) <= 90 {
		return int(itemType) - 38
	}

	return 0
}

// findDuplicateItemTypes returns a slice duplicated item types
// amongst the two compartments of a rucksack.
func findDuplicatedItemTypes(firstComp, secondComp []string) string {
	for _, itemType := range firstComp {
		if contains(secondComp, itemType) {
			return itemType
		}
	}

	return ""
}

// contains is an helper function to check if a string
// slice contains an element.
func contains(slice []string, elem string) bool {
	for _, s := range slice {
		if s == elem {
			return true
		}
	}

	return false
}
