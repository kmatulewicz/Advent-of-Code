// Description of the task: https://adventofcode.com/2015/day/5

package main

import (
	"testing"
)

func Test_isStringNice(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"1", "ugknbfddgicrmopn", true},
		{"2", "aaa", true},
		{"3", "jchzalrnumimnmhp", false},
		{"4", "haegwjzuvuyypxyu", false},
		{"5", "dvszwmarrgswjxmb", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStringNiceP1(tt.s); got != tt.want {
				t.Logf("vowels: %t, double letters: %t, forbidden: %t", checkForThreeVowels(tt.s), checkForDoubleLetters(tt.s), checkForForbiddenSubstrings(tt.s))
				t.Errorf("isStringNiceP1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isStringNiceP2(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"1", "qjhvhtzxzqqjkmpb", true},
		{"2", "xxyxx", true},
		{"3", "uurcxstgmygtbstg", false},
		{"4", "ieodomkazucvgmuy", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStringNiceP2(tt.s); got != tt.want {
				t.Logf("rule 1: %t, rule 2: %t", checkForRepeatingPairOfLetters(tt.s), checkForDoubledLettersSeparatedByLetter(tt.s))
				t.Errorf("isStringNiceP2() = %v, want %v", got, tt.want)
			}
		})
	}
}
