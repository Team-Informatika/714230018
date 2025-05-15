package main

import "testing"

func TestComplexIfElseExample(t *testing.T) {
	tests := []struct {
		name      string
		nilai     int
		kehadiran int
		want      string
	}{
		{"Nilai 85, Kehadiran 95", 85, 95, "Sangat Baik: A+"},
		{"Nilai 85, Kehadiran 80", 85, 80, "Baik Sekali: A"},
		{"Nilai 85, Kehadiran 70", 85, 70, "Baik: A-"},
		{"Nilai 75, Kehadiran 85", 75, 85, "Baik: B+"},
		{"Nilai 75, Kehadiran 70", 75, 70, "Cukup: B"},
		{"Nilai 65, Kehadiran 75", 65, 75, "Cukup: C+"},
		{"Nilai 65, Kehadiran 65", 65, 65, "Kurang: C"},
		{"Nilai 50, Kehadiran 65", 50, 65, "Buruk: D"},
		{"Nilai 50, Kehadiran 50", 50, 50, "Gagal: E"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ComplexIfElseExample(tt.nilai, tt.kehadiran)
			if got != tt.want {
				t.Errorf("ComplexIfElseExample(%d, %d) = %q; want %q", tt.nilai, tt.kehadiran, got, tt.want)
			}
		})
	}
}
