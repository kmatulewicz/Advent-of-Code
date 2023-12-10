// Description of the task: https://adventofcode.com/2015/day/8

package main

import (
	"testing"
)

func Test_countChars(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"1", `""`, 2},
		{"2", `"abc"`, 5},
		{"3", `"aaa\"aaa"`, 10},
		{"4", `"\x27"`, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countChars(tt.s); got != tt.want {
				t.Errorf("countChars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countEscapedString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"1", `""`, 0},
		{"2", `"abc"`, 3},
		{"3", `"aaa\"aaa"`, 7},
		{"4", `"\x27"`, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countEscapedString(tt.s); got != tt.want {
				t.Errorf("countEscapedString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countDiffLenP1(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"1", `""`, 2},
		{"2", `"abc"`, 2},
		{"3", `"aaa\"aaa"`, 3},
		{"4", `"\x27"`, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDiffLenP1(tt.s); got != tt.want {
				t.Errorf("countDiffLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countQuotedString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"1", `""`, 6},
		{"2", `"abc"`, 9},
		{"3", `"aaa\"aaa"`, 16},
		{"4", `"\x27"`, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countQuotedString(tt.s); got != tt.want {
				t.Errorf("countQuotedString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countDiffLenP2(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"1", `""`, 4},
		{"2", `"abc"`, 4},
		{"3", `"aaa\"aaa"`, 6},
		{"4", `"\x27"`, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDiffLenP2(tt.s); got != tt.want {
				t.Errorf("countDiffLenP2() = %v, want %v", got, tt.want)
			}
		})
	}
}
