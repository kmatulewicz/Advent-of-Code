// Description of the task: https://adventofcode.com/2015/day/13

package main

import (
	"helpers"
	"testing"
)

func Test_optimalArrangementValue(t *testing.T) {

	input := helpers.LoadLines("input.test")
	guests := parseInput(input)

	tests := []struct {
		name   string
		guests []Guest
		want   int
	}{
		{"1", guests, 330},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := optimalArrangementValue(tt.guests); got != tt.want {
				t.Errorf("optimalArrangementValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
