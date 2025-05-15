package main

import "testing"

func TestCheckEvenOdd(t *testing.T) {
	tests := []struct {
		input int
		want  string
	}{
		{2, "even"},
		{3, "odd"},
		{0, "even"},
		{-1, "odd"},
	}

	for _, tt := range tests {
		got := CheckEvenOdd(tt.input)
		if got != tt.want {
			t.Errorf("CheckEvenOdd(%d) = %q; want %q", tt.input, got, tt.want)
		}
	}
}

func TestGradeClassifier(t *testing.T) {
	tests := []struct {
		score int
		want  string
	}{
		{95, "A"},
		{85, "B"},
		{75, "C"},
		{65, "D"},
		{50, "F"},
	}

	for _, tt := range tests {
		got := GradeClassifier(tt.score)
		if got != tt.want {
			t.Errorf("GradeClassifier(%d) = %q; want %q", tt.score, got, tt.want)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{1, 2, 2},
		{5, 3, 5},
		{0, 0, 0},
		{-1, -5, -1},
	}

	for _, tt := range tests {
		got := Max(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("Max(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
		}
	}
}
