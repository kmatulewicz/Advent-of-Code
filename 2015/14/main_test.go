// Description of the task: https://adventofcode.com/2015/day/14

package main

import (
	"helpers"
	"reflect"
	"testing"
)

var r []Reindeer

func init() {
	r = []Reindeer{
		{"Comet", 10, 127, 14},
		{"Dancer", 11, 162, 16},
	}
}

func Test_calcDistance(t *testing.T) {
	type args struct {
		time int
		r    Reindeer
	}
	tests := []struct {
		name string
		time int
		r    Reindeer
		want int
	}{
		{"Comet", 1000, r[0], 1120},
		{"Dancer", 1000, r[1], 1056},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcDistance(tt.time, tt.r); got != tt.want {
				t.Errorf("calcDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseInput(t *testing.T) {
	lines := helpers.LoadLines("input.test")

	if got := parseInput(lines); !reflect.DeepEqual(got, r) {
		t.Errorf("parseInput() = %v, want %v", got, r)
	}
}

func Test_assignPoints(t *testing.T) {
	time := 1000
	want := map[string]int{
		"Dancer": 689,
		"Comet":  312,
	}
	if got := assignPoints(time, r); !reflect.DeepEqual(got, want) {
		t.Errorf("assignPoints() = %v, want %v", got, want)
	}

}
