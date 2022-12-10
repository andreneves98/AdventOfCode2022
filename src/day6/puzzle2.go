package day6

import (
	"advent-of-code/utils"
	"log"
)

func RunPuzzle2() {
	datastream, err := utils.ReadFile("./src/day6/input.txt")
	if err != nil {
		log.Printf("failed to read file: %v", err)
		return
	}

	processedChars := 0
	section := []rune{}
	for idx := 0; idx < len(datastream[0]); idx++ {
		if idx < 14 {
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
