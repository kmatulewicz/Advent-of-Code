package main

import "slices"

type loc struct {
	x, y int
}

type Santa struct {
	c loc   // current location
	l []loc // visited houses locations
}

// New returns Santa with a starting position
func New() Santa {
	return Santa{loc{}, []loc{{0, 0}}}
}

// Move moves Santa to the pointed direction
func (s *Santa) Move(direction rune) {
	switch {
	case direction == '^':
		s.c.y++
	case direction == 'v':
		s.c.y--
	case direction == '<':
		s.c.x--
	case direction == '>':
		s.c.x++
	default:
		panic("wrong instruction")
	}

	if !slices.Contains(s.l, s.c) {
		s.l = append(s.l, s.c)
	}
}
