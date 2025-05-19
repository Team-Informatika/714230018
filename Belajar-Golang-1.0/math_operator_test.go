package main

import (
	"strings"
	"testing"
)

func TestModulusInts(t *testing.T) {
	t.Run("Valid modulus", func(t *testing.T) {
		ErrorCounter = 0
		result, err := ModulusInts(10, 3)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if result != 1 {
			t.Errorf("Expected 1, got %d", result)
		}
	})

	t.Run("Division by zero", func(t *testing.T) {
		ErrorCounter = 0
		_, err := ModulusInts(10, 0)
		if err == nil {
			t.Errorf("Expected error for division by zero, got nil")
		} else {
			if !strings.Contains(err.Error(), "division by zero") {
				t.Errorf("Unexpected error message: %v", err)
			}
		}
		if ErrorCounter != 1 {
			t.Errorf("Expected ErrorCounter to be 1, got %d", ErrorCounter)
		}
	})
}

func TestPowerInts(t *testing.T) {
	t.Run("Positive exponent", func(t *testing.T) {
		result := PowerInts(2, 3)
		if result != 8 {
			t.Errorf("Expected 8, got %d", result)
		}
	})

	t.Run("Zero exponent", func(t *testing.T) {
		result := PowerInts(5, 0)
		if result != 1 {
			t.Errorf("Expected 1, got %d", result)
		}
	})

	t.Run("Negative exponent", func(t *testing.T) {
		result := PowerInts(2, -1)
		if result != 0 {
			t.Errorf("Expected 0 for negative exponent, got %d", result)
		}
	})
}

func TestPowerFloats(t *testing.T) {
	tests := []struct {
		name     string
		base     float64
		exponent float64
		expected float64
	}{
		{"Positive exponent", 2.0, 3.0, 8.0},
		{"Zero exponent", 5.0, 0.0, 1.0},
		{"Negative exponent", 2.0, -1.0, 0.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PowerFloats(tt.base, tt.exponent)
			if result != tt.expected {
				t.Errorf("Expected %f, got %f", tt.expected, result)
			}
		})
	}
}

func TestAbsInt(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"Negative input", -5, 5},
		{"Positive input", 5, 5},
		{"Zero input", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AbsInt(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestAbsFloat(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"Negative input", -5.5, 5.5},
		{"Positive input", 5.5, 5.5},
		{"Zero input", 0.0, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AbsFloat(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %f, got %f", tt.expected, result)
			}
		})
	}
}

func TestSqrtFloat(t *testing.T) {
	t.Run("Valid square root", func(t *testing.T) {
		ErrorCounter = 0
		result, err := SqrtFloat(16.0)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if result != 4.0 {
			t.Errorf("Expected 4.0, got %f", result)
		}
	})

	t.Run("Square root of negative number", func(t *testing.T) {
		ErrorCounter = 0
		_, err := SqrtFloat(-1.0)
		if err == nil {
			t.Errorf("Expected error for square root of negative number, got nil")
		} else {
			if !strings.Contains(err.Error(), "square root of negative number") {
				t.Errorf("Unexpected error message: %v", err)
			}
		}
		if ErrorCounter != 1 {
			t.Errorf("Expected ErrorCounter to be 1, got %d", ErrorCounter)
		}
	})
}
