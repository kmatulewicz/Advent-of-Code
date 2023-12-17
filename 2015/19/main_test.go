// Description of the task: https://adventofcode.com/2015/day/19

package main

import (
	"testing"
)

func Test_calculateMoleculesNumber(t *testing.T) {
	input := []string{
		"H => HO",
		"H => OH",
		"O => HH",
		"R => HH",
	}
	r, _ := parseInput(input)

	tests := []struct {
		name         string
		replacements map[string][]string
		molecule     string
		want         int
	}{
		{"HOH", r, "HOH", 4},
		{"HOHOHO", r, "HOHOHO", 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := calculateMoleculesNumber(tt.replacements, tt.molecule); got != tt.want {
				t.Errorf("calculateMoleculesNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
