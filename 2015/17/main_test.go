// Description of the task: https://adventofcode.com/2015/day/17

package main

import "testing"

func Test_countCombinations(t *testing.T) {

	tests := []struct {
		name      string
		amount    int
		inventory []int
		want1     int
		want2     int
	}{
		{"1", 25, []int{20, 15, 10, 5, 5}, 4, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got1, got2 := countCombinations(tt.amount, tt.inventory); got1 != tt.want1 || got2 != tt.want2 {
				t.Errorf("countCombinations() = %v, want %v and %v, want %v", got1, tt.want1, got2, tt.want2)
			}
		})
	}
}
