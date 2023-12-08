package main

import (
	"reflect"
	"testing"
)

func Test_sliceUnique(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"1", []int{1, 1, 1}, []int{1}},
		{"2", []int{1, 1, 2, 2, 3, 3, 4, 4}, []int{1, 2, 3, 4}},
		{"3", []int{6, 2, 1, 6, 2, 1}, []int{6, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sliceUnique(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendUniques() = %v, want %v", got, tt.want)
			}
		})
	}
}
