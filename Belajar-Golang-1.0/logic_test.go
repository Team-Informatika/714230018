package main

import "testing"

func TestAnd(t *testing.T) {
	tests := []struct {
		name string
		a, b bool
		want bool
	}{
		{"true and true", true, true, true},
		{"true and false", true, false, false},
		{"false and true", false, true, false},
		{"false and false", false, false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := And(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("And(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestOr(t *testing.T) {
	tests := []struct {
		name string
		a, b bool
		want bool
	}{
		{"true or true", true, true, true},
		{"true or false", true, false, true},
		{"false or true", false, true, true},
		{"false or false", false, false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Or(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Or(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
func TestNot(t *testing.T) {
	tests := []struct {
		name string
		a    bool
		want bool
	}{
		{"not true", true, false},
		{"not false", false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Not(tt.a)
			if got != tt.want {
				t.Errorf("Not(%v) = %v; want %v", tt.a, got, tt.want)
			}
		})
	}
}

func TestXor(t *testing.T) {
	tests := []struct {
		name string
		a, b bool
		want bool
	}{
		{"true xor true", true, true, false},
		{"true xor false", true, false, true},
		{"false xor true", false, true, true},
		{"false xor false", false, false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Xor(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Xor(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestNand(t *testing.T) {
	tests := []struct {
		name string
		a, b bool
		want bool
	}{
		{"true nand true", true, true, false},
		{"true nand false", true, false, true},
		{"false nand true", false, true, true},
		{"false nand false", false, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Nand(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Nand(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestNor(t *testing.T) {
	tests := []struct {
		name string
		a, b bool
		want bool
	}{
		{"true nor true", true, true, false},
		{"true nor false", true, false, false},
		{"false nor true", false, true, false},
		{"false nor false", false, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Nor(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Nor(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestXnor(t *testing.T) {
	tests := []struct {
		name string
		a, b bool
		want bool
	}{
		{"true xnor true", true, true, true},
		{"true xnor false", true, false, false},
		{"false xnor true", false, true, false},
		{"false xnor false", false, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Xnor(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Xnor(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
