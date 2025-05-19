package main

import "testing"

func TestNor(t *testing.T) {
	tests := []struct {
		a, b bool
		want bool
	}{
		{false, false, true},
		{false, true, false},
		{true, false, false},
		{true, true, false},
	}
	for _, tt := range tests {
		got := Nor(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("Nor(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestNand(t *testing.T) {
	tests := []struct {
		a, b bool
		want bool
	}{
		{false, false, true},
		{false, true, true},
		{true, false, true},
		{true, true, false},
	}
	for _, tt := range tests {
		got := Nand(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("Nand(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestXor(t *testing.T) {
	tests := []struct {
		a, b bool
		want bool
	}{
		{false, false, false},
		{false, true, true},
		{true, false, true},
		{true, true, false},
	}
	for _, tt := range tests {
		got := Xor(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("Xor(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestXnor(t *testing.T) {
	tests := []struct {
		a, b bool
		want bool
	}{
		{false, false, true},
		{false, true, false},
		{true, false, false},
		{true, true, true},
	}
	for _, tt := range tests {
		got := Xnor(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("Xnor(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestAnd3(t *testing.T) {
	tests := []struct {
		a, b, c bool
		want    bool
	}{
		{true, true, true, true},
		{true, true, false, false},
		{true, false, true, false},
		{false, true, true, false},
		{false, false, false, false},
	}
	for _, tt := range tests {
		got := And3(tt.a, tt.b, tt.c)
		if got != tt.want {
			t.Errorf("And3(%v, %v, %v) = %v; want %v", tt.a, tt.b, tt.c, got, tt.want)
		}
	}
}

func TestOr3(t *testing.T) {
	tests := []struct {
		a, b, c bool
		want    bool
	}{
		{true, true, true, true},
		{true, true, false, true},
		{true, false, true, true},
		{false, true, true, true},
		{false, false, false, false},
	}
	for _, tt := range tests {
		got := Or3(tt.a, tt.b, tt.c)
		if got != tt.want {
			t.Errorf("Or3(%v, %v, %v) = %v; want %v", tt.a, tt.b, tt.c, got, tt.want)
		}
	}
}

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
