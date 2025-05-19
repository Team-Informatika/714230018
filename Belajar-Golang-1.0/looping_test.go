package main

import "testing"

func TestSumRange(t *testing.T) {
	tests := []struct {
		name       string
		start, end int
		want       int
	}{
		{"normal range", 1, 5, 15},
		{"zero range", 0, 0, 0},
		{"single element range", 3, 3, 3},
		{"invalid range", 5, 1, 0},   // invalid range, expect 0 or no sum
		{"negative range", -3, 3, 0}, // additional test case, expect 0 because loop won't run
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SumRange(tt.start, tt.end)
			if got != tt.want {
				t.Errorf("SumRange(%d, %d) = %d; want %d", tt.start, tt.end, got, tt.want)
			}
		})
	}
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"zero factorial", 0, 1},
		{"one factorial", 1, 1},
		{"five factorial", 5, 120},
		{"negative input", -1, -1},   // error case
		{"large input", 10, 3628800}, // additional test case
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Factorial(tt.n)
			if got != tt.want {
				t.Errorf("Factorial(%d) = %d; want %d", tt.n, got, tt.want)
			}
		})
	}
}
func TestCountVowels(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"lowercase vowels", "hello", 2},
		{"uppercase vowels", "HELLO", 2},
		{"no vowels", "xyz", 0},
		{"all vowels", "aeiouAEIOU", 10},
		{"empty string", "", 0},
		{"mixed string", "Golang is fun!", 4}, // additional test case
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountVowels(tt.input)
			if got != tt.want {
				t.Errorf("CountVowels(%q) = %d; want %d", tt.input, got, tt.want)
			}
		})
	}
}

// TestGenerateRange tests the GenerateRange function
func TestGenerateRange(t *testing.T) {
	tests := []struct {
		name       string
		start, end int
		want       []int
	}{
		{"normal range", 1, 5, []int{1, 2, 3, 4, 5}},
		{"single element", 3, 3, []int{3}},
		{"empty range", 5, 1, []int{}},
		{"negative to positive", -2, 2, []int{-2, -1, 0, 1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateRange(tt.start, tt.end)
			if len(got) != len(tt.want) {
				t.Errorf("GenerateRange(%d, %d) = %v; want %v", tt.start, tt.end, got, tt.want)
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("GenerateRange(%d, %d)[%d] = %d; want %d", tt.start, tt.end, i, got[i], tt.want[i])
				}
			}
		})
	}
}

// TestMaxInSlice tests the MaxInSlice function
func TestMaxInSlice(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"normal slice", []int{1, 3, 2, 5, 4}, 5},
		{"single element", []int{7}, 7},
		{"empty slice", []int{}, 0},
		{"all negative", []int{-5, -1, -3}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxInSlice(tt.nums)
			if got != tt.want {
				t.Errorf("MaxInSlice(%v) = %d; want %d", tt.nums, got, tt.want)
			}
		})
	}
}
