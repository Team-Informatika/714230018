package main

import (
	"testing"
)

func TestAddUnique(t *testing.T) {
	result := Add(3, 4)
	if result != 7 {
		t.Errorf("Add(3,4) = %d; want 7", result)
	}
}

func TestSubtractUnique(t *testing.T) {
	result := Subtract(10, 5)
	if result != 5 {
		t.Errorf("Subtract(10,5) = %d; want 5", result)
	}
}

func TestMultiplyUnique(t *testing.T) {
	result := Multiply(2, 3)
	if result != 6 {
		t.Errorf("Multiply(2,3) = %d; want 6", result)
	}
}

func TestDivideUnique(t *testing.T) {
	quotient, remainder, err := Divide(10, 3)
	if err != nil {
		t.Errorf("Divide(10,3) returned error: %v", err)
	}
	if quotient != 3 || remainder != 1 {
		t.Errorf("Divide(10,3) = quotient %d, remainder %d; want 3, 1", quotient, remainder)
	}

	_, err = Divide(10, 0)
	if err == nil {
		t.Errorf("Divide(10,0) expected error but got nil")
	}
}

func TestModulusIntsUnique(t *testing.T) {
	result, err := ModulusInts(10, 3)
	if err != nil {
		t.Errorf("ModulusInts(10,3) returned error: %v", err)
	}
	if result != 1 {
		t.Errorf("ModulusInts(10,3) = %d; want 1", result)
	}

	_, err = ModulusInts(10, 0)
	if err == nil {
		t.Errorf("ModulusInts(10,0) expected error but got nil")
	}
}

func TestPowerIntsUnique(t *testing.T) {
	result := PowerInts(2, 3)
	if result != 8 {
		t.Errorf("PowerInts(2,3) = %d; want 8", result)
	}

	result = PowerInts(2, -1)
	if result != 0 {
		t.Errorf("PowerInts(2,-1) = %d; want 0", result)
	}
}

func TestPowerFloatsUnique(t *testing.T) {
	result := PowerFloats(2, 3)
	if result != 8 {
		t.Errorf("PowerFloats(2,3) = %f; want 8", result)
	}
}

func TestAbsIntUnique(t *testing.T) {
	if AbsInt(-5) != 5 {
		t.Errorf("AbsInt(-5) = %d; want 5", AbsInt(-5))
	}
	if AbsInt(5) != 5 {
		t.Errorf("AbsInt(5) = %d; want 5", AbsInt(5))
	}
}

func TestAbsFloatUnique(t *testing.T) {
	if AbsFloat(-5.5) != 5.5 {
		t.Errorf("AbsFloat(-5.5) = %f; want 5.5", AbsFloat(-5.5))
	}
	if AbsFloat(5.5) != 5.5 {
		t.Errorf("AbsFloat(5.5) = %f; want 5.5", AbsFloat(5.5))
	}
}
func TestSqrtFloatUnique(t *testing.T) {
	result, err := SqrtFloat(16)
	if err != nil {
		t.Errorf("SqrtFloat(16) returned error: %v", err)
	}
	if result != 4 {
		t.Errorf("SqrtFloat(16) = %f; want 4", result)
	}

	_, err = SqrtFloat(-1)
	if err == nil {
		t.Errorf("SqrtFloat(-1) expected error but got nil")
	}
}
