// Description of the task: https://adventofcode.com/2015/day/6

package main

import (
	"fmt"
	"helpers"
)

func main() {
	lines := helpers.LoadLines("input")

	// Part 1
	g := New()
	for _, line := range lines {
		g.Command(line)
	}
	fmt.Printf("Part 1: After execution of commands there are %d lightened lights in greed.\n", g.Count())

	// Part 2
	g = New()
	for _, line := range lines {
		g.Command2(line)
	}
	fmt.Printf("Part 2: After execution of commands total brightness of lights is equal to %d.\n", g.Count())
}
