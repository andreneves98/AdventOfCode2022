package main

import (
	"advent-of-code/src/day1"
	"advent-of-code/src/day2"
	"advent-of-code/src/day3"
	"advent-of-code/src/day4"
	"advent-of-code/src/day5"
	"os"
)

func main() {
	switch os.Args[1] {
	case "1":
		day1.Run()
	case "2.1":
		day2.RunPuzzle1()
	case "2.2":
		day2.RunPuzzle2()
	case "3.1":
		day3.RunPuzzle1()
	case "3.2":
		day3.RunPuzzle2()
	case "4.1":
		day4.RunPuzzle1()
	case "4.2":
		day4.RunPuzzle2()
	case "5.1":
		day5.RunPuzzle1()
	case "5.2":
		day5.RunPuzzle2()
	}
}
