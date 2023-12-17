// Description of the task: https://adventofcode.com/2015/day/19

package main

import (
	"fmt"
	"helpers"
	"slices"
	"strings"
)

func main() {

	lines := helpers.LoadLines("input")

	replacements, molecule := parseInput(lines)

	nCombinations, _ := calculateMoleculesNumber(replacements, molecule)
	fmt.Printf("%d distinct molecules can be created.\n", nCombinations)

	step := createMolecule(parseInput2(lines))
	fmt.Printf("%d steps are needed to obtain the molecule.\n", step)
}

// createMolecule returns a number of steps needed to create the molecule using provided replacements
func createMolecule(repG1 map[string][]string, repG2 map[string][]string, molecule string) int {
	step := 0
loop:
	for ; ; step++ {

		// done
		if molecule == "e" {
			break loop
		}

		// do replacements with a greater reduction.
		for k, v := range repG2 {
			for _, vv := range v {
				if strings.Index(molecule, vv) != -1 {
					molecule = strings.Replace(molecule, vv, k, 1)
					continue loop
				}
			}
		}

		// if there is no other possibility do replacements with a smaller reduction.
		for k, v := range repG1 {
			for _, vv := range v {
				if strings.Index(molecule, vv) != -1 {
					molecule = strings.Replace(molecule, vv, k, 1)
					continue loop
				}
			}
		}

		// no replacement was possible
		panic("no replace possible")
	}

	return step
}

// calculateMoleculesNumber calculates the number of molecules that can be created and returns a slice of all unique combinations
func calculateMoleculesNumber(replacements map[string][]string, molecule string) (int, map[string]bool) {
	results := make(map[string]bool)

	for k, v := range replacements {
		for _, vv := range v {
			n := strings.Count(molecule, k)
			index := 0
			for ; n > 0; n-- {
				// set the index to the new change point
				index = strings.Index(molecule[index:], k) + index
				// copy the unchanged part of string
				s := strings.Clone(molecule[:index])
				// add the rest of the base string with the change
				s += strings.Replace(molecule[index:], k, vv, 1)
				// append to the slice of results
				results[s] = true
				// move index after the point of change
				index += len(k)
			}
		}
	}

	return len(results), results
}

func parseInput(lines []string) (map[string][]string, string) {
	r := make(map[string][]string)
	m := ""

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, " ")
		if len(parts) == 3 {
			v, ok := r[parts[0]]
			if ok {
				v = append(v, parts[2])
				r[parts[0]] = v
			} else {
				r[parts[0]] = []string{parts[2]}
			}
			continue
		}

		// the molecule line is not empty and cant be split
		m = line
	}

	return r, m
}

// split replacements into 2 groups
func parseInput2(lines []string) (map[string][]string, map[string][]string, string) {
	g1 := make(map[string][]string)
	g2 := make(map[string][]string)

	m := ""

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, " ")
		if len(parts) == 3 {

			if strings.Index(parts[2], "Rn") == -1 {
				v, ok := g1[parts[0]]
				if ok {
					v = append(v, parts[2])
					g1[parts[0]] = v
				} else {
					g1[parts[0]] = []string{parts[2]}
				}
			} else {
				v, ok := g2[parts[0]]
				if ok {
					v = append(v, parts[2])
					g2[parts[0]] = v
				} else {
					g2[parts[0]] = []string{parts[2]}
				}
			}
			continue
		}

		// the molecule line is not empty and cant be split
		m = line
	}

	return g1, g2, m
}

func unique(input []string) []string {
	// make results unique
	unique := []string{}
	for i := range input {
		if slices.Index(unique, input[i]) == -1 {
			unique = append(unique, input[i])
		}
	}
	return unique
}
