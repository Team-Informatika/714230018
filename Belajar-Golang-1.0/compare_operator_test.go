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
		got := Equal[int](tt.a, tt.b)
		if got != tt.want {
			t.Errorf("Equal[int](%d, %d) = %v; want %v", tt.a, tt.b, got, tt.want)
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
		got := NotEqual[int](tt.a, tt.b)
		if got != tt.want {
			t.Errorf("NotEqual[int](%d, %d) = %v; want %v", tt.a, tt.b, got, tt.want)
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
		got := GreaterThan[int](tt.a, tt.b)
		if got != tt.want {
			t.Errorf("GreaterThan[int](%d, %d) = %v; want %v", tt.a, tt.b, got, tt.want)
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
		got := LessThan[int](tt.a, tt.b)
		if got != tt.want {
			t.Errorf("LessThan[int](%d, %d) = %v; want %v", tt.a, tt.b, got, tt.want)
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
		got := GreaterThanOrEqual[int](tt.a, tt.b)
		if got != tt.want {
			t.Errorf("GreaterThanOrEqual[int](%d, %d) = %v; want %v", tt.a, tt.b, got, tt.want)
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
		got := LessThanOrEqual[int](tt.a, tt.b)
		if got != tt.want {
			t.Errorf("LessThanOrEqual[int](%d, %d) = %v; want %v", tt.a, tt.b, got, tt.want)
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
		got := Equal[float64](tt.a, tt.b)
		if got != tt.want {
			t.Errorf("Equal[float64](%f, %f) = %v; want %v", tt.a, tt.b, got, tt.want)
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
		got := NotEqual[float64](tt.a, tt.b)
		if got != tt.want {
			t.Errorf("NotEqual[float64](%f, %f) = %v; want %v", tt.a, tt.b, got, tt.want)
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
		got := GreaterThan[float64](tt.a, tt.b)
		if got != tt.want {
			t.Errorf("GreaterThan[float64](%f, %f) = %v; want %v", tt.a, tt.b, got, tt.want)
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
		got := LessThan[float64](tt.a, tt.b)
		if got != tt.want {
			t.Errorf("LessThan[float64](%f, %f) = %v; want %v", tt.a, tt.b, got, tt.want)
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
		got := GreaterThanOrEqual[float64](tt.a, tt.b)
		if got != tt.want {
			t.Errorf("GreaterThanOrEqual[float64](%f, %f) = %v; want %v", tt.a, tt.b, got, tt.want)
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
		got := LessThanOrEqual[float64](tt.a, tt.b)
		if got != tt.want {
			t.Errorf("LessThanOrEqual[float64](%f, %f) = %v; want %v", tt.a, tt.b, got, tt.want)
		}
	}
}
