package day1

import (
	"advent-of-code/utils"
	"log"
	"sort"
	"strconv"

	"github.com/pkg/errors"
)

func RunPuzzle1() (map[int]int, error) {
	log.Println("*****************************************")
	log.Println("* Advent of Code 2022: Day 1 - Puzzle 1 *")
	log.Println("*****************************************")
	inputFile := "./src/day1/input.txt"

	// Read input file
	lines, err := utils.ReadFile(inputFile)
	if err != nil {
		log.Printf("failed to read file: %v", err)
		return map[int]int{}, errors.Wrap(err, "failed to run puzzle 1")
	}

	// Count calories by elf
	caloriesByElf, err := CountCaloriesByElf(lines)
	if err != nil {
		log.Printf("failed to count calories from file: %v", err)
		return map[int]int{}, errors.Wrap(err, "failed to run puzzle 1")
	}

	// Check which elf carries the most number of calories and how much
	calories, elf := GetMostCarryingElfCalories(caloriesByElf)
	log.Printf("The elf carrying most calories is elf #%d with %d calories!",
		elf, calories)
	log.Println("")

	return caloriesByElf, nil
}

func RunPuzzle2(caloriesByElf map[int]int) {
	log.Println("")
	log.Println("*****************************************")
	log.Println("* Advent of Code 2022: Day 1 - Puzzle 2 *")
	log.Println("*****************************************")

	// Append calories by elf to slice
	var caloriesByElfSlice []int
	for _, calories := range caloriesByElf {
		caloriesByElfSlice = append(caloriesByElfSlice, calories)
	}

	// Sort the list of calories
	sort.Ints(caloriesByElfSlice)

	// Count top 3 total calories
	top3TotalCalories := CountTop3TotalCalories(caloriesByElfSlice[len(caloriesByElfSlice)-3:])
	log.Printf("The total calories carried by the top 3 elves combined are: %d",
		top3TotalCalories)
}

// CountTop3TotalCalories returns the combined total calories
// carried by the top three elves.
func CountTop3TotalCalories(top3Calories []int) int {
	var total int
	for _, calories := range top3Calories {
		total += calories
	}

	return total
}

// CountCalories counts all the calories carried by the elves
// and stores them in a map with both keys and values as integers.
// If successful, returns the map as the first return argument
// and an error as the second one.
func CountCaloriesByElf(lines []string) (map[int]int, error) {
	log.Println("Counting elve's calories...")
	var (
		caloriesByElf = make(map[int]int)
		elf           = 1
	)

	for _, line := range lines {
		if line != "" {
			calories, err := strconv.Atoi(line)
			if err != nil {
				log.Printf("failed to convert string to int: %v", err)
				return map[int]int{}, errors.Wrap(err, "failed to count calories")
			}
			caloriesByElf[elf] += calories
		} else {
			elf++
			continue
		}
	}

	log.Println("Finished counting all the calories carried by the elves...")
	return caloriesByElf, nil
}

// GetMostCarryingElfCalories returns the number of calories
// carried by the elf who carries the most and that elf's number,
// respectively.
func GetMostCarryingElfCalories(caloriesByElf map[int]int) (int, int) {
	var (
		maxCalories int
		maxElf      int
	)
	for elf, calories := range caloriesByElf {
		if calories > maxCalories {
			maxCalories = calories
			maxElf = elf
		}
	}

	return maxCalories, maxElf
}
