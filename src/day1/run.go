package day1

import "log"

func Run() {
	caloriesByElf, err := RunPuzzle1()
	if err != nil {
		log.Fatalf("Failed running puzzle 1: %v", err)
	}

	RunPuzzle2(caloriesByElf)
}
