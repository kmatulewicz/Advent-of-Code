// Description of the task: https://adventofcode.com/2015/day/11

package main

import "testing"

func Test_nextPass(t *testing.T) {
	tests := []struct {
		name    string
		current string
		want    string
	}{
		{"abcdefgh", "abcdefgh", "abcdffaa"},
		{"ghijklmn", "ghijklmn", "ghjaabcc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextPass(tt.current); got != tt.want {
				t.Errorf("nextPass() = %v, want %v", got, tt.want)
			}
		})
	}
}
