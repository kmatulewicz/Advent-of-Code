// Description of the task: https://adventofcode.com/2015/day/20

package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		limit int
		want  int
	}{
		{"1", 10, 1},
		{"4", 70, 4},
		{"8", 130, 8},
		{"12", 240, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(1, tt.limit); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {

	tests := []struct {
		name  string
		limit int
		want  int
	}{
		{"1", 11, 1},
		{"2", 33, 2},
		{"3", 44, 3},
		{"4", 77, 4},
		{"4", 66, 4},
		{"6", 132, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(1, tt.limit); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
