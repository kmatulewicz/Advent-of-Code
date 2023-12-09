// Description of the task: https://adventofcode.com/2015/day/5

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\n")

	// Part 1
	counter := 0
	for _, line := range lines {
		if isStringNiceP1(line) {
			counter++
		}
	}

	fmt.Printf("Input contains %d nice strings according to rules in part 1.\n", counter)

	// Part 2
	counter = 0
	for _, line := range lines {
		if isStringNiceP2(line) {
			counter++
		}
	}

	fmt.Printf("Input contains %d nice strings according to rules in part 2.\n", counter)
}

// isStringNiceP1 checks if string is nice according to rules in part 1
func isStringNiceP1(s string) bool {
	if checkForThreeVowels(s) && checkForDoubleLetters(s) && checkForForbiddenSubstrings(s) {
		return true
	}

	return false
}

// isStringNiceP2 checks if string is nice according to rules in part 1
func isStringNiceP2(s string) bool {
	if checkForRepeatingPairOfLetters(s) && checkForDoubledLettersSeparatedByLetter(s) {
		return true
	}

	return false
}

// checkForThreeVowels checks if string contains at least three vowels
func checkForThreeVowels(s string) bool {
	count := 0
	for _, ch := range s {
		switch ch {
		case 'a', 'e', 'i', 'o', 'u':
			if count == 2 {
				return true
			}
			count++
		}
	}

	return false
}

// checkForDoubleLetters checks if string contains at least one letter that appears twice in a row
func checkForDoubleLetters(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			return true
		}
	}

	return false
}

// checkForForbiddenSubstrings checks if string does not contain the strings ab, cd, pq, or xy
func checkForForbiddenSubstrings(s string) bool {
	for i := 1; i < len(s); i++ {
		switch subStr := s[i-1 : i+1]; subStr {
		case "ab", "cd", "pq", "xy":
			return false
		}
	}

	return true
}

// checkForRepeatingPairOfLetters checks if string contains a pair of any two letters that appears at least twice in the string without overlapping
func checkForRepeatingPairOfLetters(s string) bool {
	for i := 1; i < len(s); i++ {
		if strings.Count(s, s[i-1:i+1]) > 1 {
			return true
		}
	}

	return false
}

// checkForDoubledLettersSeparatedByLetter checks if string contains at least one letter which repeats with exactly one letter between them
func checkForDoubledLettersSeparatedByLetter(s string) bool {
	for i := 2; i < len(s); i++ {
		if s[i] == s[i-2] {
			return true
		}
	}

	return false
}
