// Description of the task: https://adventofcode.com/2015/day/17

package main

import (
	"log"
	"slices"
	"testing"
)

func Test_countLights(t *testing.T) {
	isString := []string{
		".#.#.#",
		"...##.",
		"#....#",
		"..#...",
		"#.#..#",
		"####.."}
	slices.Reverse(isString)

	is := parseInput(isString)

	tests := []struct {
		name         string
		initialState [][]bool
		steps        int
		want         int
		partTwo      bool
	}{
		{"1", is, 1, 11, false},
		{"2", is, 2, 8, false},
		{"3", is, 3, 4, false},
		{"4", is, 4, 4, false},
		{"P2-1", is, 1, 18, true},
		{"P2-2", is, 2, 18, true},
		{"P2-3", is, 3, 18, true},
		{"P2-4", is, 4, 14, true},
		{"P2-5", is, 5, 17, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Println("Test: ", tt.name)
			if got := countLights(tt.initialState, tt.steps, tt.partTwo); got != tt.want {
				t.Errorf("countLights() = %v, want %v", got, tt.want)
			}
		})
	}
}
