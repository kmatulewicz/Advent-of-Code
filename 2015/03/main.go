// Description of the task: https://adventofcode.com/2015/day/3

package main

import (
	"fmt"
	"helpers"
	"slices"
)

func main() {
	input := helpers.Load("input")

	// Part 1
	santa := New()
	for _, move := range string(input) {
		santa.Move(move)
	}
	fmt.Printf("Part 1: %d houses received at least one present\n", len(santa.l))

	// Part 2
	santa = New()
	roboSanta := New()
	for i, move := range string(input) {
		if i%2 == 0 {
			santa.Move(move)
		} else {
			roboSanta.Move(move)
		}
	}

	combined := append(santa.l, roboSanta.l...)
	unique := sliceUnique(combined)
	fmt.Printf("Part 2: %d houses received at least one present\n", len(unique))
}

func sliceUnique[T comparable](a []T) []T {
	u := make([]T, 0)
	for i := range a {
		if !slices.Contains(u, a[i]) {
			u = append(u, a[i])
		}
	}
	return u
}
