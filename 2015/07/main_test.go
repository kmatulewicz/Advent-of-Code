// Description of the task: https://adventofcode.com/2015/day/7
package main

import (
	"strings"
	"testing"
)

func Test_do(t *testing.T) {
	tests := []struct {
		name string
		dest string
		want uint16
	}{
		{"1", "d", 72},
		{"2", "e", 507},
		{"3", "f", 492},
		{"4", "g", 114},
		{"5", "h", 65412},
		{"6", "i", 65079},
		{"7", "x", 123},
		{"8", "y", 456},
	}

	testString := "123 -> x\n456 -> y\nx AND y -> d\nx OR y -> e\nx LSHIFT 2 -> f\ny RSHIFT 2 -> g\nNOT x -> h\nNOT y -> i"
	lines := strings.Split(testString, "\n")

	commands := []command{}
	for _, line := range lines {
		if line == "" {
			break
		}
		commands = append(commands, parseCommand(line))
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sv := make(map[string]uint16)
			if got := do(commands, sv, tt.dest); got != tt.want {
				t.Errorf("do() = %v, want %v", got, tt.want)
			}
		})
	}
}
