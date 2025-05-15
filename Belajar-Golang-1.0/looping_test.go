package main

import "testing"

func TestSumRange(t *testing.T) {
	tests := []struct {
		start, end int
		want       int
	}{
		{1, 5, 15},
		{0, 0, 0},
		{3, 3, 3},
		{5, 1, 0}, // invalid range, expect 0 or no sum
	}

	for _, tt := range tests {
		got := SumRange(tt.start, tt.end)
		if got != tt.want {
			t.Errorf("SumRange(%d, %d) = %d; want %d", tt.start, tt.end, got, tt.want)
		}
	}
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 1},
		{1, 1},
		{5, 120},
		{-1, -1}, // error case
	}

	for _, tt := range tests {
		got := Factorial(tt.n)
		if got != tt.want {
			t.Errorf("Factorial(%d) = %d; want %d", tt.n, got, tt.want)
		}
	}
}

func TestCountVowels(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"hello", 2},
		{"HELLO", 2},
		{"xyz", 0},
		{"aeiouAEIOU", 10},
		{"", 0},
	}

	for _, tt := range tests {
		got := CountVowels(tt.input)
		if got != tt.want {
			t.Errorf("CountVowels(%q) = %d; want %d", tt.input, got, tt.want)
		}
	}
}
