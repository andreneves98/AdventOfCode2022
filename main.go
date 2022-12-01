package main

import (
	"advent-of-code/src/day1"
	"os"
)

func main() {
	switch os.Args[1] {
	case "1":
		day1.Run()
	}
}
