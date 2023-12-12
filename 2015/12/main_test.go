// Description of the task: https://adventofcode.com/2015/day/12

package main

import "testing"

func Test_sumDigits(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		removeRed bool
		want      float64
	}{
		{"[1,2,3]", "[1,2,3]", false, 6},
		{`{"a":2,"b":4}`, `{"a":2,"b":4}`, false, 6},
		{`[[[3]]]`, `[[[3]]]`, false, 3},
		{`{"a":{"b":4},"c":-1}`, `{"a":{"b":4},"c":-1}`, false, 3},
		{`{"a":[-1,1]}`, `{"a":[-1,1]}`, false, 0},
		{`[-1,{"a":1}]`, `[-1,{"a":1}]`, false, 0},
		{`[]`, `[]`, false, 0},
		{`{}`, `{}`, false, 0},
		{`[1,2,3]`, `[1,2,3]`, true, 6},
		{`[1,{"c":"red","b":2},3]`, `[1,{"c":"red","b":2},3]`, true, 4},
		{`{"d":"red","e":[1,2,3,4],"f":5}`, `{"d":"red","e":[1,2,3,4],"f":5}`, true, 0},
		{`[1,"red",5]`, `[1,"red",5]`, true, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumDigits(tt.input, tt.removeRed); got != tt.want {
				t.Errorf("sumDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
