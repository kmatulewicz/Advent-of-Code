// Description of the task: https://adventofcode.com/2015/day/1

package main

import (
	"fmt"
	"helpers"
	"io"
	"os"
)

func main() {

	lastFloor, basementEntrance := countFromString("input")
	fmt.Printf("Last floor: %d\n", lastFloor)
	fmt.Printf("Basement entrance at position: %d\n", basementEntrance)

	// slower version
	lastFloor, basementEntrance = countByteByByte("input")
	fmt.Printf("Last floor: %d\n", lastFloor)
	fmt.Printf("Basement entrance at position: %d\n", basementEntrance)

}

func countFromString(name string) (lastFloor, basementEntrance int) {
	floor, basementEntrance := 0, 0
	input := helpers.Load(name)

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

func countByteByByte(name string) (lastFloor, basementEntrance int) {
	floor := 0

	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b := make([]byte, 1)
	for i := 1; ; i++ {
		_, err := f.Read(b)
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			break
		}
		switch {
		case b[0] == '(':
			floor++
		case b[0] == ')':
			floor--
		default:
			panic("wrong rune")
		}

		// Catch only the first entrance to the basement
		if basementEntrance == 0 && floor < 0 {
			basementEntrance = i
		}
	}

	return floor, basementEntrance
}
