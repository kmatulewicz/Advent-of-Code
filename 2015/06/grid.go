package main

import (
	"strconv"
	"strings"
)

type grid struct {
	g [][]int
}

// New returns new initialized with 0 grid
func New() *grid {
	g := grid{}
	for y := 0; y < 1000; y++ {
		row := []int{}
		for x := 0; x < 1000; x++ {
			row = append(row, 0)
		}
		g.g = append(g.g, row)
	}

	return &g
}

// Count counts number of 1 in greed
func (g grid) Count() int {
	c := 0
	for y := 0; y < 1000; y++ {
		for x := 0; x < 1000; x++ {
			c += g.g[y][x]
		}
	}

	return c
}

// Command perform specified command on specified range in part 1
func (g *grid) Command(s string) {
	var fromX, fromY, toX, toY int
	var fn func(val int) int

	c := strings.Split(s, " ")

	// set coordinates and function
	switch {
	case len(c) == 4:
		//toggle
		fromX, fromY = splitCoordinates(c[1])
		toX, toY = splitCoordinates(c[3])
		fn = func(val int) int {
			if val == 1 {
				return 0
			}
			return 1
		}
	case c[1] == "on":
		//turn on
		fromX, fromY = splitCoordinates(c[2])
		toX, toY = splitCoordinates(c[4])
		fn = func(val int) int { return 1 }
	case c[1] == "off":
		//turn off
		fromX, fromY = splitCoordinates(c[2])
		toX, toY = splitCoordinates(c[4])
		fn = func(val int) int { return 0 }
	}

	// make changes in selected range
	for y := fromY; y <= toY; y++ {
		for x := fromX; x <= toX; x++ {
			g.g[y][x] = fn(g.g[y][x])
		}
	}
}

// Command2 perform specified command on specified range in part 2
func (g *grid) Command2(s string) {
	var fromX, fromY, toX, toY int
	var fn func(val int) int

	c := strings.Split(s, " ")

	// set coordinates and function
	switch {
	case len(c) == 4:
		//toggle
		fromX, fromY = splitCoordinates(c[1])
		toX, toY = splitCoordinates(c[3])
		fn = func(val int) int {
			val += 2
			return val
		}
	case c[1] == "on":
		//turn on
		fromX, fromY = splitCoordinates(c[2])
		toX, toY = splitCoordinates(c[4])
		fn = func(val int) int { val++; return val }
	case c[1] == "off":
		//turn off
		fromX, fromY = splitCoordinates(c[2])
		toX, toY = splitCoordinates(c[4])
		fn = func(val int) int {
			val--
			if val < 0 {
				val = 0
			}
			return val
		}
	}

	// make changes in selected range
	for y := fromY; y <= toY; y++ {
		for x := fromX; x <= toX; x++ {
			g.g[y][x] = fn(g.g[y][x])
		}
	}
}

func splitCoordinates(s string) (int, int) {
	c := strings.Split(s, ",")
	x, err := strconv.Atoi(c[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(c[1])
	if err != nil {
		panic(err)
	}
	return x, y
}
