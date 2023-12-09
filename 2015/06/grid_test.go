package main

import (
	"reflect"
	"testing"
)

var g *grid

// get new greed filled with alternating 0 nad 1
func prep() {
	g = New()
	change := false
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if change {
				g.g[x][y] = 1
			}
			change = !change
		}
	}
}

// Test part 1
func Test_grid1(t *testing.T) {
	tests := []struct {
		name    string
		command string
		want    int
	}{
		{"1", "turn on 0,0 through 999,999", 1000000},
		{"2", "toggle 0,0 through 999,0", 500000},
		{"3", "turn off 499,499 through 500,500", 499998},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prep()
			if g.Count() != 500000 {
				t.Errorf("not prepared")
			}
			g.Command(tt.command)
			if got := g.Count(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("g.Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test part 2
func Test_grid2(t *testing.T) {
	tests := []struct {
		name    string
		command string
		want    int
	}{
		{"1", "turn on 0,0 through 999,999", 1500000},
		{"2", "toggle 0,0 through 999,0", 502000},
		{"3", "turn off 499,499 through 500,500", 499998},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prep()
			if g.Count() != 500000 {
				t.Errorf("not prepared")
			}
			g.Command2(tt.command)
			if got := g.Count(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("g.Count() = %v, want %v", got, tt.want)
			}
		})
	}
}
