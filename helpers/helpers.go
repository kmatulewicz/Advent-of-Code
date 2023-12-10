package helpers

import (
	"os"
	"strings"
)

// Load loads specified file and return its content as string
func Load(s string) string {
	input, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	return string(input)
}

// LoadLines loads specified file and return its content as slice of lines.
// If last line is empty it will be removed from slice
func LoadLines(s string) []string {
	input := Load(s)
	lines := strings.Split(input, "\n")
	if lines[len(lines)-1] == "" {
		return lines[:len(lines)-1]
	}
	return lines
}
