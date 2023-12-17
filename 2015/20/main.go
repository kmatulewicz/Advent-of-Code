// Description of the task: https://adventofcode.com/2015/day/20

package main

import (
	"fmt"
)

func main() {

	// input
	const limit = 36_000_000

	num := part1(1, limit)
	fmt.Printf("Part 1: The lowest house number of the house to get at least %d presents is: %d\n", limit, num)

	num = part2(1, limit)
	fmt.Printf("Part 2: The lowest house number of the house to get at least %d presents is: %d\n", limit, num)
}

// optimized brute force method
// TODO: Sieve of Eratosthenes - max is always in the next step after a first number
func part1(start, limit int) int {
	if start%2 != 0 {
		start++
	}
	// Special case
	if limit <= 10 {
		return 1
	}

	for i := start; ; i += 2 {
		sum := i * 10
		for j := 1; j <= int(i/2); j++ {
			if i%j == 0 {
				sum += j * 10
			}
		}

		if sum >= limit {
			return i
		}
	}
}

func part2(start, limit int) int {
	if start%2 != 0 {
		start++
	}
	// Special case
	if limit <= 11 {
		return 1
	}
	for i := start; ; i += 2 {
		sum := i * 11
		for j := max(1, int(i/50)); j <= int(i/2); j++ {
			if i%j == 0 {
				sum += j * 11
			}
		}
		if sum >= limit {
			return i
		}
	}
}
