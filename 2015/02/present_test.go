package main

import (
	"testing"
)

func Test_present_CountMinSurface(t *testing.T) {
	type fields struct {
		l int
		w int
		h int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"1", fields{1, 2, 3}, 2},
		{"2", fields{3, 2, 1}, 2},
		{"3", fields{2, 2, 2}, 4},
		{"4", fields{2, 3, 4}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := present{
				l: tt.fields.l,
				w: tt.fields.w,
				h: tt.fields.h,
			}
			if got := p.CountMinSurface(); got != tt.want {
				t.Errorf("present.CountMinSurface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_present_CountSurface(t *testing.T) {
	type fields struct {
		l int
		w int
		h int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"1", fields{1, 2, 3}, 22},
		{"2", fields{3, 2, 1}, 22},
		{"3", fields{2, 2, 2}, 24},
		{"4", fields{2, 3, 4}, 52},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := present{
				l: tt.fields.l,
				w: tt.fields.w,
				h: tt.fields.h,
			}
			if got := p.CountSurface(); got != tt.want {
				t.Errorf("present.CountSurface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_present_CountRibbon(t *testing.T) {
	type fields struct {
		l int
		w int
		h int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"1", fields{1, 2, 3}, 12},
		{"2", fields{3, 2, 1}, 12},
		{"3", fields{2, 2, 2}, 16},
		{"4", fields{2, 3, 4}, 34},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := present{
				l: tt.fields.l,
				w: tt.fields.w,
				h: tt.fields.h,
			}
			if got := p.CountRibbon(); got != tt.want {
				t.Errorf("present.CountRibbon() = %v, want %v", got, tt.want)
			}
		})
	}
}
