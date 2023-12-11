// Description of the task: https://adventofcode.com/2015/day/10

package main

import "testing"

func Test_lookAndSay(t *testing.T) {
	tests := []struct {
		name       string
		input      []rune
		iterations int
		want       string
	}{
		{"1", []rune("1"), 1, "11"},
		{"11", []rune("11"), 1, "21"},
		{"21", []rune("21"), 1, "1211"},
		{"1211", []rune("1211"), 1, "111221"},
		{"111221", []rune("111221"), 1, "312211"},
		{"1", []rune("1"), 5, "312211"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lookAndSay(tt.input, tt.iterations); got != tt.want {
				t.Errorf("lookAndSay() = %v, want %v", got, tt.want)
			}
		})
	}
}
