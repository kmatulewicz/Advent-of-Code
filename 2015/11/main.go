// Description of the task: https://adventofcode.com/2015/day/11

package main

import (
	"fmt"
	"strings"
)

func main() {
	current := "vzbxkghb"
	new := nextPass(current)
	fmt.Printf("New password after %s is: %s\n", current, new)
	current = new
	new = nextPass(current)
	fmt.Printf("New password after %s is: %s\n", current, new)

}

// nextPass generates next password fulfilling all rules
func nextPass(current string) string {
	rules := []func(string) bool{rule1, rule2, rule3}
	for {
		new := incPass(current)
		if checkRules(rules, new) {
			return new
		}
		current = new
	}
}

// incPass increments current password by 1 and returns it
func incPass(current string) string {
	b := []byte(current)
	for i := len(b) - 1; i >= 0; i-- {
		b[i]++
		if b[i] <= 'z' {
			break
		}
		b[i] = 'a'
	}
	return string(b)
}

// checkRules returns true if password fulfill all provided rules, otherwise returns false
func checkRules(checks []func(string) bool, pass string) bool {
	for _, f := range checks {
		if f(pass) == false {
			return false
		}
	}

	return true
}

// rule1 checks if password fulfills rule:
// Passwords must include one increasing straight of at least three letters.
func rule1(pass string) bool {
	b := []byte(pass)
	for i := 0; i < len(b)-3; i++ {
		if b[i] == (b[i+1]-1) && b[i] == (b[i+2]-2) {
			return true
		}
	}

	return false
}

// rule2 checks if password fulfills rule:
// Passwords may not contain the letters 'i', 'o', or 'l'.
func rule2(pass string) bool {
	if strings.Index(pass, "i") != -1 {
		return false
	}
	if strings.Index(pass, "o") != -1 {
		return false
	}
	if strings.Index(pass, "l") != -1 {
		return false
	}
	return true
}

// rule3 checks if password fulfills rule:
// Passwords must contain at least two different, non-overlapping pairs of letters.
func rule3(pass string) bool {
	b := []byte(pass)
	lastPair := byte(0)
	for i := 0; i < len(b)-1; i++ {
		if b[i] == b[i+1] {
			if lastPair != 0 && lastPair != b[i] {
				return true
			}
			lastPair = b[i]
			i++
		}

	}

	return false
}
