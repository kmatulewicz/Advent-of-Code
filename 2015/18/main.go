// Description of the task: https://adventofcode.com/2015/day/18

package main

import (
	"fmt"
	"helpers"
)

const DEBUG = false

func main() {
	lines := helpers.LoadLines("input")
	initialState := parseInput(lines)

	steps := 100
	count := countLights(initialState, 100, false)

	fmt.Printf("After %d steps, there are %d lights on.\n", steps, count)

	count = countLights(initialState, 100, true)
	fmt.Printf("In part 2, after %d steps, there are %d lights on.\n", steps, count)

}

func countLights(initialState [][]bool, steps int, partTwo bool) int {

	// turn on the lights in corners
	if partTwo {
		part2(initialState)
	}

	// print the initial state
	if DEBUG {
		fmt.Println("Initial state")
		printState(initialState)
	}

	newState := [][]bool{}
	for y := 0; y < len(initialState); y++ {
		newRow := []bool{}
		for x := 0; x < len(initialState[y]); x++ {
			sum := countNeighbors(initialState, x, y)
			switch {
			case initialState[y][x] && sum != 2 && sum != 3:
				// A light which is on stays on when 2 or 3 neighbors are on, and turns off otherwise.
				newRow = append(newRow, false)
			case !initialState[y][x] && sum == 3:
				//A light which is off turns on if exactly 3 neighbors are on, and stays off otherwise.
				newRow = append(newRow, true)
			default:
				newRow = append(newRow, initialState[y][x])
			}
		}
		newState = append(newState, newRow)
	}

	// turn on the lights in corners
	if partTwo {
		part2(newState)
	}

	// print the new state
	if DEBUG {
		fmt.Println("New state")
		printState(newState)
	}

	if steps != 1 {
		// there are some steps to do
		return countLights(newState, steps-1, partTwo)
	} else {
		// calculate the final sum
		sum := 0
		for y := 0; y < len(newState); y++ {
			for x := 0; x < len(newState[y]); x++ {
				if newState[y][x] {
					sum++
				}
			}
		}
		return sum
	}
}

func countNeighbors(state [][]bool, x, y int) int {
	sum := 0

	// top row
	if y+1 < len(state) {
		sum = countRowNeighbors(state[y+1], x)
	}
	// middle row
	if x-1 >= 0 && state[y][x-1] {
		sum++
	}
	if x+1 < len(state[y]) && state[y][x+1] {
		sum++
	}
	//bottom row
	if y-1 >= 0 {
		sum += countRowNeighbors(state[y-1], x)
	}
	return sum
}

// countRowNeighbors counts lightened neighbors in the given row
func countRowNeighbors(row []bool, x int) int {
	sum := 0
	for i := x - 1; i <= x+1; i++ {
		if i < 0 || i > len(row)-1 || !row[i] {
			continue
		}
		sum++
	}

	return sum
}

// part2 turns on the lights in corners
func part2(s [][]bool) {
	s[0][0] = true
	s[len(s)-1][0] = true
	s[len(s)-1][len(s[len(s)-1])-1] = true
	s[0][len(s[len(s)-1])-1] = true
}

// parseInput transforms the input into a grid of lights
func parseInput(lines []string) [][]bool {
	out := [][]bool{}
	for _, line := range lines {
		tmp := []bool{}
		for _, c := range line {
			switch c {
			case '.':
				// lamp off
				tmp = append(tmp, false)
			case '#':
				// lamp on
				tmp = append(tmp, true)
			default:
				panic("wrong input")
			}
		}
		out = append(out, tmp)
	}
	return out
}

// printState prints the given state for debugging purposes
func printState(state [][]bool) {
	fmt.Println()
	for y := len(state) - 1; y >= 0; y-- {
		for x := 0; x < len(state[y]); x++ {
			if state[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
