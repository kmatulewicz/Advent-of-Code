// Description of the task: https://adventofcode.com/2015/day/1

package main

import (
	"fmt"
	"os"
)

func main() {
	input := loadFileToString("input")

	lastFloor, basementEntrance := countFromString(input)
	fmt.Printf("Last floor: %d\n", lastFloor)
	fmt.Printf("Basement entrance at position: %d\n", basementEntrance)
}

func loadFileToString(name string) string {
	out, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}

	return string(out)
}

func countFromString(input string) (lastFloor, basementEntrance int) {
	floor, basementEntrance := 0, 0

	for i, move := range input {
		switch {
		case move == '(':
			floor++
		case move == ')':
			floor--
		default:
			panic("wrong rune")
		}

		// Catch only the first entrance to the basement
		if basementEntrance == 0 && floor < 0 {
			// Moves starts from 1
			basementEntrance = i + 1
		}
	}

	return floor, basementEntrance
}
