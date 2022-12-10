package day6

import (
	"advent-of-code/utils"
	"log"
)

func RunPuzzle1() {
	datastream, err := utils.ReadFile("./src/day6/input.txt")
	if err != nil {
		log.Printf("failed to read file: %v", err)
		return
	}

	processedChars := 0
	section := []rune{}
	for idx := 0; idx < len(datastream[0]); idx++ {
		if idx < 4 {
			section = append(section, rune(datastream[0][idx]))
			continue
		}
		section = nextSection(section, rune(datastream[0][idx]))
		if isMarker(section) {
			processedChars = idx + 1
			break
		}
	}

	log.Println("Result:", processedChars)
}

func isMarker(section []rune) bool {
	countPerChar := make(map[rune]int)
	for _, char := range section {
		countPerChar[char]++
	}

	for _, count := range countPerChar {
		if count > 1 {
			return false
		}
	}

	return true
}

func nextSection(section []rune, char rune) []rune {
	section = section[1:]           // remove the first character
	section = append(section, char) // add the new one
	return section
}
