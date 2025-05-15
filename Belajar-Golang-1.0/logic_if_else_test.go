package main

import "testing"

func TestLogicIfElseExample(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  string
	}{
		{"Grade A", 85, "Nilai Anda A"},
		{"Grade B", 75, "Nilai Anda B"},
		{"Grade C", 65, "Nilai Anda C"},
		{"Grade D", 55, "Nilai Anda D"},
		{"Grade E", 45, "Nilai Anda E"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LogicIfElseExample(tt.input)
			if got != tt.want {
				t.Errorf("LogicIfElseExample(%d) = %q; want %q", tt.input, got, tt.want)
			}
		})
	}
}
