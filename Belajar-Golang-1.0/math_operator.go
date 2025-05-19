package main

import (
	"errors"
	"math"
)

// ErrorCounter counts the number of errors encountered in tests
var ErrorCounter int

// ModulusInts returns the remainder of the division of two integers.
// Returns an error if divisor is 0.
func ModulusInts(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a % b, nil
}

// PowerInts returns a raised to the power of b.
// Supports negative exponents by returning 0 for now.
func PowerInts(a, b int) int {
	if b < 0 {
		// For negative exponents, return 0 as integer power is not defined here.
		return 0
	}
	result := 1
	for i := 0; i < b; i++ {
		result *= a
	}
	return result
}

// PowerFloats returns a raised to the power of b.
func PowerFloats(a, b float64) float64 {
	return math.Pow(a, b)
}

// AbsInt returns the absolute value of an integer.
func AbsInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// AbsFloat returns the absolute value of a float64.
func AbsFloat(a float64) float64 {
	return math.Abs(a)
}

// SqrtFloat returns the square root of a float64.
// Returns an error if input is negative.
func SqrtFloat(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("square root of negative number")
	}
	return math.Sqrt(a), nil
}
