package main

import "testing"

func Test_findLowestSuffix(t *testing.T) {

	tests := []struct {
		name string
		key  string
		want int
	}{
		{"1", "abcdef", 609043},
		{"2", "pqrstuv", 1048970},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLowestSuffix(tt.key, 5); got != tt.want {
				t.Errorf("findLowestSuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}
