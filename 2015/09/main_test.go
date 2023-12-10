// Description of the task: https://adventofcode.com/2015/day/9

package main

import (
	"helpers"
	"testing"
)

func Test_calcShortestDistance(t *testing.T) {

	lines := helpers.LoadLines("input.test")
	if got := calcDistance(lines, func(v1, v2 int) int { return min(v1, v2) }); got != 605 {
		t.Errorf("calcShortestDistance() = %v, want %v", got, 605)
	}

	if got := calcDistance(lines, func(v1, v2 int) int { return max(v1, v2) }); got != 982 {
		t.Errorf("calcShortestDistance() = %v, want %v", got, 982)
	}

}
