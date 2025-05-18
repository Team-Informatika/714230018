package main

import "testing"

func TestEqualInts(t *testing.T) {
	tests := []struct {
		a, b int
		want bool
	}{
		{1, 1, true},
		{1, 2, false},
		{0, 0, true},
		{-1, -1, true},
		{-1, 1, false},
	}
	for _, tt := range tests {
		got := EqualInts(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("EqualInts(%d, %d) = %v; want %v", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestNotEqualInts(t *testing.T) {
	tests := []struct {
		a, b int
		want bool
	}{
		{1, 1, false},
		{1, 2, true},
		{0, 0, false},
		{-1, -1, false},
		{-1, 1, true},
	}
	for _, tt := range tests {
		got := NotEqualInts(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("NotEqualInts(%d, %d) = %v; want %v", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestGreaterThanInts(t *testing.T) {
	tests := []struct {
		a, b int
		want bool
	}{
		{2, 1, true},
		{1, 2, false},
		{1, 1, false},
		{0, -1, true},
		{-1, 0, false},
	}
	for _, tt := range tests {
		got := GreaterThanInts(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("GreaterThanInts(%d, %d) = %v; want %v", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestLessThanInts(t *testing.T) {
	tests := []struct {
		a, b int
		want bool
	}{
		{1, 2, true},
		{2, 1, false},
		{1, 1, false},
		{-1, 0, true},
		{0, -1, false},
	}
	for _, tt := range tests {
		got := LessThanInts(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("LessThanInts(%d, %d) = %v; want %v", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestGreaterThanOrEqualInts(t *testing.T) {
	tests := []struct {
		a, b int
		want bool
	}{
		{2, 1, true},
		{1, 2, false},
		{1, 1, true},
		{0, 0, true},
		{-1, -2, true},
		{-2, -1, false},
	}
	for _, tt := range tests {
		got := GreaterThanOrEqualInts(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("GreaterThanOrEqualInts(%d, %d) = %v; want %v", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestLessThanOrEqualInts(t *testing.T) {
	tests := []struct {
		a, b int
		want bool
	}{
		{1, 2, true},
		{2, 1, false},
		{1, 1, true},
		{0, 0, true},
		{-2, -1, true},
		{-1, -2, false},
	}
	for _, tt := range tests {
		got := LessThanOrEqualInts(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("LessThanOrEqualInts(%d, %d) = %v; want %v", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestEqualFloats(t *testing.T) {
	tests := []struct {
		a, b float64
		want bool
	}{
		{1.0, 1.0, true},
		{1.0, 2.0, false},
		{0.0, 0.0, true},
		{-1.0, -1.0, true},
		{-1.0, 1.0, false},
	}
	for _, tt := range tests {
		got := EqualFloats(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("EqualFloats(%f, %f) = %v; want %v", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestNotEqualFloats(t *testing.T) {
	tests := []struct {
		a, b float64
		want bool
	}{
		{1.0, 1.0, false},
		{1.0, 2.0, true},
		{0.0, 0.0, false},
		{-1.0, -1.0, false},
		{-1.0, 1.0, true},
	}
	for _, tt := range tests {
		got := NotEqualFloats(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("NotEqualFloats(%f, %f) = %v; want %v", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestGreaterThanFloats(t *testing.T) {
	tests := []struct {
		a, b float64
		want bool
	}{
		{2.0, 1.0, true},
		{1.0, 2.0, false},
		{1.0, 1.0, false},
		{0.0, -1.0, true},
		{-1.0, 0.0, false},
	}
	for _, tt := range tests {
		got := GreaterThanFloats(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("GreaterThanFloats(%f, %f) = %v; want %v", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestLessThanFloats(t *testing.T) {
	tests := []struct {
		a, b float64
		want bool
	}{
		{1.0, 2.0, true},
		{2.0, 1.0, false},
		{1.0, 1.0, false},
		{-1.0, 0.0, true},
		{0.0, -1.0, false},
	}
	for _, tt := range tests {
		got := LessThanFloats(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("LessThanFloats(%f, %f) = %v; want %v", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestGreaterThanOrEqualFloats(t *testing.T) {
	tests := []struct {
		a, b float64
		want bool
	}{
		{2.0, 1.0, true},
		{1.0, 2.0, false},
		{1.0, 1.0, true},
		{0.0, 0.0, true},
		{-1.0, -2.0, true},
		{-2.0, -1.0, false},
	}
	for _, tt := range tests {
		got := GreaterThanOrEqualFloats(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("GreaterThanOrEqualFloats(%f, %f) = %v; want %v", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestLessThanOrEqualFloats(t *testing.T) {
	tests := []struct {
		a, b float64
		want bool
	}{
		{1.0, 2.0, true},
		{2.0, 1.0, false},
		{1.0, 1.0, true},
		{0.0, 0.0, true},
		{-2.0, -1.0, true},
		{-1.0, -2.0, false},
	}
	for _, tt := range tests {
		got := LessThanOrEqualFloats(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("LessThanOrEqualFloats(%f, %f) = %v; want %v", tt.a, tt.b, got, tt.want)
		}
	}
}
