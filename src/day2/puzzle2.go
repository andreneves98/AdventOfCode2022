package day2

import (
	"advent-of-code/utils"
	"log"
	"strings"
)

// decryptionMap associates each shape to an index
// to be used with the plays matrix.
var decryptionMap map[string]int = map[string]int{
	"A": 0,
	"B": 1,
	"C": 2,
	"X": 0,
	"Y": 1,
	"Z": 2,
}

var outcomePoints map[string]int = map[string]int{
	"X": Lose,
	"Y": Draw,
	"Z": Win,
}

// guide defines what i need to play given the elf's play.
// It is arranged with the columns being the outcomes (X Y Z) and
// the rows being the elf's plays (A B C).
var guide [][]string = [][]string{
	// X Y Z
	{"C", "A", "B"}, // A
	{"A", "B", "C"}, // B
	{"B", "C", "A"}, // C
}

func RunPuzzle2() {
	var totalPoints int

	// // Start by reading the strategy guide line by line
	strategyGuide, err := utils.ReadFile("./src/day2/input.txt")
	if err != nil {
		log.Printf("failed to read file: %v", err)
	}

	for _, round := range strategyGuide {
		// Split the round
		roundSplit := strings.Split(round, " ")
		// Sum total points
		totalPoints += sumRoundPoints(roundSplit[0], roundSplit[1])
	}

	log.Printf("Total points: %d", totalPoints)
}

func sumRoundPoints(elfPlay, outcome string) int {
	myPlay := guide[decryptionMap[elfPlay]][decryptionMap[outcome]]
	myPlayPoints := decryptionMap[myPlay] + 1
	return myPlayPoints + outcomePoints[outcome]
}
