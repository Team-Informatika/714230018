package main

import "testing"

func TestMajority(t *testing.T) {
	tests := []struct {
		a, b, c bool
		want    bool
	}{
		{true, true, true, true},
		{true, true, false, true},
		{true, false, true, true},
		{false, true, true, true},
		{true, false, false, false},
		{false, true, false, false},
		{false, false, true, false},
		{false, false, false, false},
	}
	for _, tt := range tests {
		got := Majority(tt.a, tt.b, tt.c)
		if got != tt.want {
			t.Errorf("Majority(%v, %v, %v) = %v; want %v", tt.a, tt.b, tt.c, got, tt.want)
		}
	}
}

func TestImplies(t *testing.T) {
	tests := []struct {
		a, b bool
		want bool
	}{
		{true, true, true},
		{true, false, false},
		{false, true, true},
		{false, false, true},
	}
	for _, tt := range tests {
		got := Implies(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("Implies(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestParity(t *testing.T) {
	tests := []struct {
		a, b, c bool
		want    bool
	}{
		{true, true, true, true},
		{true, true, false, false},
		{true, false, true, false},
		{false, true, true, false},
		{true, false, false, true},
		{false, true, false, true},
		{false, false, true, true},
		{false, false, false, false},
	}
	for _, tt := range tests {
		got := Parity(tt.a, tt.b, tt.c)
		if got != tt.want {
			t.Errorf("Parity(%v, %v, %v) = %v; want %v", tt.a, tt.b, tt.c, got, tt.want)
		}
	}
}

func TestConditional(t *testing.T) {
	tests := []struct {
		a, b, c bool
		want    bool
	}{
		{true, true, false, true},
		{true, false, true, false},
		{false, true, false, false},
		{false, false, true, true},
	}
	for _, tt := range tests {
		got := Conditional(tt.a, tt.b, tt.c)
		if got != tt.want {
			t.Errorf("Conditional(%v, %v, %v) = %v; want %v", tt.a, tt.b, tt.c, got, tt.want)
		}
	}
}

func TestComplexLogic(t *testing.T) {
	tests := []struct {
		a, b, c bool
		want    bool
	}{
		{true, true, true, false},
		{true, true, false, true},
		{true, false, true, false},
		{true, false, false, true},
		{false, true, true, true},
		{false, true, false, false},
		{false, false, true, false},
		{false, false, false, false},
	}
	for _, tt := range tests {
		got := ComplexLogic(tt.a, tt.b, tt.c)
		if got != tt.want {
			t.Errorf("ComplexLogic(%v, %v, %v) = %v; want %v", tt.a, tt.b, tt.c, got, tt.want)
		}
	}
}
