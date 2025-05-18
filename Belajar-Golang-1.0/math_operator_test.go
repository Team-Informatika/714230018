package main

import "testing"

func TestModulusInts(t *testing.T) {
	ErrorCounter = 0
	result, err := ModulusInts(10, 3)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}

	_, err = ModulusInts(10, 0)
	if err == nil {
		t.Errorf("Expected error for division by zero, got nil")
	}
	if ErrorCounter != 1 {
		t.Errorf("Expected ErrorCounter to be 1, got %d", ErrorCounter)
	}
}
func TestPowerInts(t *testing.T) {
	result := PowerInts(2, 3)
	if result != 8 {
		t.Errorf("Expected 8, got %d", result)
	}

	result = PowerInts(5, 0)
	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
}

func TestPowerFloats(t *testing.T) {
	result := PowerFloats(2.0, 3.0)
	if result != 8.0 {
		t.Errorf("Expected 8.0, got %f", result)
	}

	result = PowerFloats(5.0, 0.0)
	if result != 1.0 {
		t.Errorf("Expected 1.0, got %f", result)
	}
}

func TestAbsInt(t *testing.T) {
	if AbsInt(-5) != 5 {
		t.Errorf("Expected 5, got %d", AbsInt(-5))
	}
	if AbsInt(5) != 5 {
		t.Errorf("Expected 5, got %d", AbsInt(5))
	}
	if AbsInt(0) != 0 {
		t.Errorf("Expected 0, got %d", AbsInt(0))
	}
}

func TestAbsFloat(t *testing.T) {
	if AbsFloat(-5.5) != 5.5 {
		t.Errorf("Expected 5.5, got %f", AbsFloat(-5.5))
	}
	if AbsFloat(5.5) != 5.5 {
		t.Errorf("Expected 5.5, got %f", AbsFloat(5.5))
	}
	if AbsFloat(0.0) != 0.0 {
		t.Errorf("Expected 0.0, got %f", AbsFloat(0.0))
	}
}

func TestSqrtFloat(t *testing.T) {
	ErrorCounter = 0
	result, err := SqrtFloat(16.0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 4.0 {
		t.Errorf("Expected 4.0, got %f", result)
	}

	_, err = SqrtFloat(-1.0)
	if err == nil {
		t.Errorf("Expected error for square root of negative number, got nil")
	}
	if ErrorCounter != 1 {
		t.Errorf("Expected ErrorCounter to be 1, got %d", ErrorCounter)
	}
}
