package main

import (
	"advent-of-code/src/day1"
	"advent-of-code/src/day2"
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
	}
}
