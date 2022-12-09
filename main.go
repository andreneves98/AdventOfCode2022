package main

import (
	"advent-of-code/src/day1"
	"advent-of-code/src/day2"
	"advent-of-code/src/day3"
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
	}
}
