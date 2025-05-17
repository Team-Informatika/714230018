package main

import "testing"

func TestCheckEvenOdd(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  string
	}{
		{"even positive", 2, "even"},
		{"odd positive", 3, "odd"},
		{"zero", 0, "even"},
		{"odd negative", -1, "odd"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckEvenOdd(tt.input)
			if got != tt.want {
				t.Errorf("CheckEvenOdd(%d) = %q; want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestGradeClassifier(t *testing.T) {
	tests := []struct {
		name  string
		score int
		want  string
	}{
		{"grade A", 95, "A"},
		{"grade B", 85, "B"},
		{"grade C", 75, "C"},
		{"grade D", 65, "D"},
		{"grade F", 50, "F"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GradeClassifier(tt.score)
			if got != tt.want {
				t.Errorf("GradeClassifier(%d) = %q; want %q", tt.score, got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"a less than b", 1, 2, 2},
		{"a greater than b", 5, 3, 5},
		{"a equals b", 0, 0, 0},
		{"negative values", -1, -5, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Max(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Max(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestIsPositive(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want bool
	}{
		{"positive number", 5, true},
		{"zero", 0, false},
		{"negative number", -3, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsPositive(tt.num)
			if got != tt.want {
				t.Errorf("IsPositive(%d) = %v; want %v", tt.num, got, tt.want)
			}
		})
	}
}

func TestIsLeapYear(t *testing.T) {
	tests := []struct {
		name string
		year int
		want bool
	}{
		{"year divisible by 400", 2000, true},
		{"year divisible by 100 but not 400", 1900, false},
		{"year divisible by 4 but not 100", 2012, true},
		{"year not divisible by 4", 2019, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsLeapYear(tt.year)
			if got != tt.want {
				t.Errorf("IsLeapYear(%d) = %v; want %v", tt.year, got, tt.want)
			}
		})
	}
}

func TestDescribeNumber(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want string
	}{
		{"zero", 0, "zero"},
		{"positive", 5, "positive"},
		{"negative", -3, "negative"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DescribeNumber(tt.num)
			if got != tt.want {
				t.Errorf("DescribeNumber(%d) = %q; want %q", tt.num, got, tt.want)
			}
		})
	}
}

func TestTemperatureClassification(t *testing.T) {
	tests := []struct {
		name string
		temp float64
		want string
	}{
		{"freezing", -5.0, "Freezing"},
		{"cold", 10.0, "Cold"},
		{"warm", 20.0, "Warm"},
		{"hot", 30.0, "Hot"},
		{"unknown", 999.0, "Unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TemperatureClassification(tt.temp)
			if got != tt.want {
				t.Errorf("TemperatureClassification(%f) = %q; want %q", tt.temp, got, tt.want)
			}
		})
	}
}
