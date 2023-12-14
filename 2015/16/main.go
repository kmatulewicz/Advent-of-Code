// Description of the task: https://adventofcode.com/2015/day/16

package main

import (
	"fmt"
	"helpers"
	"strconv"
	"strings"
)

func main() {
	lines := helpers.LoadLines("input")
	facts := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}

	aunts := parseInput(lines)

	i := indexOfAunt(aunts, facts)

	fmt.Println("Part 1 -> Sue number:", i)

	i = indexOfAunt2(aunts, facts)

	fmt.Println("Part 2 -> Sue number:", i)

}

// parseInput parses an input into a slice of maps
// where an index is an aunt's number, a key is the compound's name, and the value is the value of the compound.
func parseInput(lines []string) []map[string]int {

	aunts := make([]map[string]int, 501)

	for _, line := range lines {

		// get aunt number
		parts := strings.SplitN(line, " ", 3)
		num, err := strconv.Atoi(parts[1][:len(parts[1])-1])
		if err != nil {
			panic(err.Error())
		}

		// get compounds
		m := make(map[string]int)
		rest := parts[2]
		for {
			colon := strings.Index(rest, ":")
			comma := strings.Index(rest, ",")

			// end of a current expression
			end := comma
			if comma == -1 {
				end = len(rest)
			}

			val, err := strconv.Atoi(rest[colon+2 : end])
			if err != nil {
				panic(err.Error())
			}
			m[rest[:colon]] = val

			// this is the last expression
			if comma == -1 {
				break
			}

			// cut evaluated part from string
			rest = rest[comma+2:]
		}
		aunts[num] = m

	}

	return aunts
}

// indexOfAunt returns the index of the aunt for whom the greatest number of facts match
func indexOfAunt(aunts []map[string]int, facts map[string]int) int {

	maxScore := 0
	savedIndex := -1

	// go for all aunts
	for i := range aunts {
		score := 0
		// check all facts
		for k, v1 := range facts {
			v2, ok := aunts[i][k]
			if !ok {
				// if aunt has no information abut that fact
				continue
			}

			if v1 == v2 {
				// if the fact match
				score++
			}
		}

		// save the best score and the index of current aunt
		if score > maxScore {
			maxScore = score
			savedIndex = i
		}

	}

	return savedIndex
}

// indexOfAunt returns the index of the aunt for whom the greatest number of facts match
// for part 2
func indexOfAunt2(aunts []map[string]int, facts map[string]int) int {

	maxScore := 0
	savedIndex := -1

	// go for all aunts
	for i := range aunts {
		score := 0
		// check all facts
		for k, v1 := range facts {
			v2, ok := aunts[i][k]
			if !ok {
				// if aunt has no information abut that fact
				continue
			}
			switch k {
			case "cats", "trees":
				if v1 < v2 {
					// if the fact match
					score++
				}
			case "pomeranians", "goldfish":
				if v1 > v2 {
					// if the fact match
					score++
				}
			default:
				if v1 == v2 {
					// if the fact match
					score++
				}
			}
		}

		// save the best score and the index of current aunt
		if score > maxScore {
			maxScore = score
			savedIndex = i
		}

	}

	return savedIndex
}
