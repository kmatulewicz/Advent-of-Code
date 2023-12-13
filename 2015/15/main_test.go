// Description of the task: https://adventofcode.com/2015/day/15

package main

import (
	"helpers"
	"testing"
)

func Test_calculateScores(t *testing.T) {
	lines := helpers.LoadLines("input.test")
	ingredients := parseInput(lines)
	want := 62842880
	want500 := 57600000

	if got1, got2 := calculateScores(ingredients, 100, 0); got1 != want || got2 != want500 {
		t.Errorf("calculateScores() = %v, %v, want %v, %v", got1, got2, want, want500)
	}
}
