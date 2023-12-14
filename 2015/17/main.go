// Description of the task: https://adventofcode.com/2015/day/17

package main

import (
	"fmt"
	"helpers"
	"strconv"
)

func main() {
	goal := 150

	lines := helpers.LoadLines("input")
	containers := []int{}
	for i := range lines {
		v, err := strconv.Atoi(lines[i])
		if err != nil {
			panic(err.Error())
		}
		containers = append(containers, v)
	}

	allCombinations, combinations := countCombinations(goal, containers)

	fmt.Printf("There is %d different combinations of containers.\n", allCombinations)
	fmt.Printf("There is %d different combinations for minimal number of used containers.\n", combinations)

}

// countCombinations counts all combinations of containers allowing to store the pointed amount of liquid and
// a number of combinations allowing to store the pointed amount of liquid in a minimal number of containers
func countCombinations(amount int, inventory []int) (int, int) {
	used := []int{}
	ch := make(chan int)
	b := block{}

	b.Add()
	go fit(used, inventory, amount, ch, &b)

	// combinations for minimal amount of containers
	combinations := 0
	// all combinations
	allCombinations := 0

	minNum := -1
loop:
	for {
		select {
		case num := <-ch:
			switch {
			case minNum == -1:
				minNum = num
				combinations++
				allCombinations++
			case num < minNum:
				allCombinations++
				combinations = 1
				minNum = num
			case num == minNum:
				combinations++
				allCombinations++
			default:
				allCombinations++
			}
		default:
			if b.Check() {
				break loop
			}
		}
	}

	return allCombinations, combinations
}

// fit fits recursively liquid into available containers
func fit(used []int, available []int, goal int, ch chan int, b *block) {
	// already stored amount
	sumUsed := sumSlice(used)

	// try every available container
	for i := range available {

		// check sum with newly used container
		sum := sumUsed + available[i]

		// exactly fits
		if goal == sum {
			ch <- len(used) + 1
			continue
		}

		// space left unused
		if goal < sum {
			continue
		}

		// some liquid is still waiting for a storage
		if goal > sum {
			uc := make([]int, len(used)+1)
			copy(uc, used)
			uc = append(uc, available[i])

			ac := make([]int, 0)
			ac = append(ac, available[i+1:]...)
			b.Add()

			// try adding some new container
			go fit(uc, ac, goal, ch, b)
		}
	}
	b.Sub()
}

// sumSlice sums all elements in int slice
func sumSlice(s []int) int {
	sum := 0
	for i := range s {
		sum += s[i]
	}

	return sum
}
