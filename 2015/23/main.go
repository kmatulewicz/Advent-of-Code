// Description of the task: https://adventofcode.com/2015/day/23

package main

import (
	"fmt"
	"helpers"
	"strconv"
)

func main() {
	lines := helpers.LoadLines("input")
	p := &Program{}
	p.Init1(lines)
	p.Init2(lines)
}

type Program struct {
	pointer int           //current instruction
	r       map[byte]uint // registers
}

// hlf sets register r to half its current value, then continues with the next instruction
func (p *Program) hlf(r byte) {
	p.r[r] /= 2
	p.pointer++
}

// tmp sets register r to triple its current value, then continues with the next instruction
func (p *Program) tpl(r byte) {
	p.r[r] *= 3
	p.pointer++
}

// inc increments register r, adding 1 to it, then continues with the next instruction
func (p *Program) inc(r byte) {
	p.r[r]++
	p.pointer++
}

// jmp is a jump; it continues with the instruction offset away relative to itself
func (p *Program) jmp(o int) {
	p.pointer += o
}

// jie is like jmp, but only jumps if register r is even
func (p *Program) jie(r byte, o int) {
	if p.r[r]%2 == 0 {
		p.jmp(o)
	} else {
		p.pointer++
	}
}

// jio is like jmp, but only jumps if register r is 1
func (p *Program) jio(r byte, o int) {
	if p.r[r] == 1 {
		p.jmp(o)
	} else {
		p.pointer++
	}
}

// doInst parses instruction
func (p *Program) doInst(inst string) {
	switch inst[0:3] {
	case "hlf":
		p.hlf(inst[4])
	case "tpl":
		p.tpl(inst[4])
	case "inc":
		p.inc(inst[4])
	case "jmp":
		p.jmp(getOffset(inst[4:]))
	case "jie":
		p.jie(inst[4], getOffset(inst[7:]))
	case "jio":
		p.jio(inst[4], getOffset(inst[7:]))
	default:
		panic("wrong instruction")
	}

}

// getOffset converts string to int
func getOffset(s string) int {
	offset, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return offset
}

// Init1 runs instructions in i for the part 1
func (p *Program) Init1(i []string) {

	p.r = map[byte]uint{'a': 0, 'b': 0}
	p.pointer = 0

	for p.pointer < len(i) && p.pointer >= 0 {
		p.doInst(i[p.pointer])
	}

	fmt.Println("Part 1: a:", p.r['a'], "b:", p.r['b'])
}

// Init2 runs instructions in i for the part 2
func (p *Program) Init2(i []string) {

	p.r = map[byte]uint{'a': 1, 'b': 0}
	p.pointer = 0

	for p.pointer < len(i) && p.pointer >= 0 {
		p.doInst(i[p.pointer])
	}

	fmt.Println("Part 2: a:", p.r['a'], "b:", p.r['b'])
}
