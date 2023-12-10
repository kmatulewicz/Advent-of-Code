// Description of the task: https://adventofcode.com/2015/day/8

package main

import (
	"fmt"
	"helpers"
	"strconv"
)

func main() {
	lines := helpers.LoadLines("input")
	diff := 0
	for _, line := range lines {
		diff += countDiffLenP1(line)
	}

	fmt.Printf("Total difference in part 1 is: %d\n", diff)

	diff = 0
	for _, line := range lines {
		diff += countDiffLenP2(line)
	}

	fmt.Printf("Total difference in part 2 is: %d\n", diff)

}

func countEscapedString(s string) int {
	uq, err := strconv.Unquote(s)
	if err != nil {
		panic(err)
	}
	return len(uq)
}

func countChars(s string) int {
	return len(s)
}

func countDiffLenP1(s string) int {
	return countChars(s) - countEscapedString(s)
}

func countQuotedString(s string) int {
	q := strconv.Quote(s)
	return len(q)
}

func countDiffLenP2(s string) int {
	return countQuotedString(s) - countChars(s)
}
