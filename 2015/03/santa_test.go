package main

import "testing"

func TestSanta_Move(t *testing.T) {
	tests := []struct {
		name         string
		instructions string
		visited      int
	}{
		{"1", ">", 2},
		{"2", "^>v<", 4},
		{"3", "^v^v^v^v^v", 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			for _, m := range tt.instructions {
				s.Move(m)
			}
			if len(s.l) != tt.visited {
				t.Fail()
			}
		})
	}
}
