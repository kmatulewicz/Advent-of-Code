package main

import (
	"sort"
)

type present struct {
	l, w, h int
}

// CountSurface counts surface of all sides
func (p present) CountSurface() int {
	return 2*p.l*p.w + 2*p.w*p.h + 2*p.h*p.l
}

// CountMinSurface counts the smallest side surface
func (p present) CountMinSurface() int {
	min := p.l * p.w
	if p.w*p.h < min {
		min = p.w * p.h
	}
	if p.h*p.l < min {
		min = p.h * p.l
	}
	return min
}

// CountRibbon counts the total length of ribbon needed
func (p present) CountRibbon() int {
	a, b := p.getTwoShortestDimensions()
	ribbon := 2*a + 2*b + p.l*p.w*p.h
	return ribbon
}

func (p present) getTwoShortestDimensions() (int, int) {
	ints := []int{p.l, p.w, p.h}
	sort.Ints(ints)
	return ints[0], ints[1]
}
