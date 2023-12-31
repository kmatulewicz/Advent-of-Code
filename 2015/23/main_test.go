// Description of the task: https://adventofcode.com/2015/day/23

package main

import (
	"helpers"
	"testing"
)

func TestProgram_Init(t *testing.T) {
	lines := helpers.LoadLines("input.test")
	p := &Program{}
	p.Init1(lines)

	if p.r['a'] != 2 {
		t.Fail()
	}
}
