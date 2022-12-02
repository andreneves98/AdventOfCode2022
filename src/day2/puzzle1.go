package day2

import (
	"advent-of-code/utils"
	"log"
	"strings"
)

// shapeToIdx1 associates each shape to an index
// to be used with the plays matrix.
var shapeToIdx map[string]int = map[string]int{
	"A": 0,
	"B": 1,
	"C": 2,
	"X": 0,
	"Y": 1,
	"Z": 2,
}

// playsMatrix1 defines the outcomes of all the possible plays.
// It is arranged with the columns being our plays (X Y Z) and
// the rows being the elf's plays (A B C).
var playsMatrix [][]int = [][]int{
	// X Y Z
	{Draw, Win, Lose}, // A
	{Lose, Draw, Win}, // B
	{Win, Lose, Draw}, // C
}

func RunPuzzle1() {
	var totalPoints int

	// Start by reading the strategy guide line by line
	strategyGuide, err := utils.ReadFile("./src/day2/input.txt")
	if err != nil {
		log.Printf("failed to read file: %v", err)
	}

	for _, round := range strategyGuide {
		// Split the round into shapes
		shapes := strings.Split(round, " ")
		// Sum total points
		totalPoints += getRoundWinner(shapes[0], shapes[1])
	}

	log.Printf("Total points: %d", totalPoints)
}

// getRoundWinner checks the round against the plays matrix to know
// which situation occurred and sums the points of the outcome and
// the played shape.
func getRoundWinner(elfPlay, myPlay string) int {
	return shapeToIdx[myPlay] + 1 + playsMatrix[shapeToIdx[elfPlay]][shapeToIdx[myPlay]]
}
