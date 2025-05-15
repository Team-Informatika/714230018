package main

import (
	"math"
	"testing"
)

func TestAddInts(t *testing.T) {
	result := AddInts(2, 3)
	if result != 5 {
		t.Errorf("AddInts(2, 3) = %d; want 5", result)
	}
}

func TestSubtractInts(t *testing.T) {
	result := SubtractInts(5, 3)
	if result != 2 {
		t.Errorf("SubtractInts(5, 3) = %d; want 2", result)
	}
}

func TestDivideInts(t *testing.T) {
	result := DivideInts(10, 2)
	if result != 5 {
		t.Errorf("DivideInts(10, 2) = %d; want 5", result)
	}
	result = DivideInts(10, 0)
	if result != 0 {
		t.Errorf("DivideInts(10, 0) = %d; want 0", result)
	}
}

func TestAddFloats(t *testing.T) {
	result := AddFloats(2.5, 3.5)
	if result != 6.0 {
		t.Errorf("AddFloats(2.5, 3.5) = %f; want 6.0", result)
	}
}

func TestSubtractFloats(t *testing.T) {
	result := SubtractFloats(5.5, 3.0)
	if result != 2.5 {
		t.Errorf("SubtractFloats(5.5, 3.0) = %f; want 2.5", result)
	}
}
func TestDivideFloats(t *testing.T) {
	result := DivideFloats(10.0, 2.0)
	if result != 5.0 {
		t.Errorf("DivideFloats(10.0, 2.0) = %f; want 5.0", result)
	}
	result = DivideFloats(10.0, 0.0)
	if result != 0.0 {
		t.Errorf("DivideFloats(10.0, 0.0) = %f; want 0.0", result)
	}
}

func TestCompareInts(t *testing.T) {
	if !EqualInts(5, 5) {
		t.Errorf("EqualInts(5, 5) = false; want true")
	}
	if NotEqualInts(5, 3) == false {
		t.Errorf("NotEqualInts(5, 3) = false; want true")
	}
	if !GreaterThanInts(5, 3) {
		t.Errorf("GreaterThanInts(5, 3) = false; want true")
	}
	if !LessThanInts(3, 5) {
		t.Errorf("LessThanInts(3, 5) = false; want true")
	}
	if !GreaterThanOrEqualInts(5, 5) {
		t.Errorf("GreaterThanOrEqualInts(5, 5) = false; want true")
	}
	if !LessThanOrEqualInts(5, 5) {
		t.Errorf("LessThanOrEqualInts(5, 5) = false; want true")
	}
}
func TestCompareFloats(t *testing.T) {
	if !EqualFloats(5.5, 5.5) {
		t.Errorf("EqualFloats(5.5, 5.5) = false; want true")
	}
	if NotEqualFloats(5.5, 3.3) == false {
		t.Errorf("NotEqualFloats(5.5, 3.3) = false; want true")
	}
	if !GreaterThanFloats(5.5, 3.3) {
		t.Errorf("GreaterThanFloats(5.5, 3.3) = false; want true")
	}
	if !LessThanFloats(3.3, 5.5) {
		t.Errorf("LessThanFloats(3.3, 5.5) = false; want true")
	}
	if !GreaterThanOrEqualFloats(5.5, 5.5) {
		t.Errorf("GreaterThanOrEqualFloats(5.5, 5.5) = false; want true")
	}
	if !LessThanOrEqualFloats(5.5, 5.5) {
		t.Errorf("LessThanOrEqualFloats(5.5, 5.5) = false; want true")
	}

	// Edge cases
	if !EqualFloats(-0.0, 0.0) {
		t.Errorf("EqualFloats(-0.0, 0.0) = false; want true")
	}
	if EqualFloats(1e-10, 0.0) {
		t.Errorf("EqualFloats(1e-10, 0.0) = true; want false")
	}
	if !GreaterThanFloats(1e10, 1e9) {
		t.Errorf("GreaterThanFloats(1e10, 1e9) = false; want true")
	}
	if !LessThanFloats(-1e10, 1e9) {
		t.Errorf("LessThanFloats(-1e10, 1e9) = false; want true")
	}
	if !NotEqualFloats(math.NaN(), math.NaN()) {
		t.Errorf("NotEqualFloats(NaN, NaN) = false; want true")
	}
	// NaN comparisons
	// Use math.NaN() to get NaN value and math package's IsNaN to check
	_ = math.NaN()
}
