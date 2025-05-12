package main

import "math"

// ErrorCounter counts the number of errors occurred in math operations
var ErrorCounter int

// ModulusInts returns the remainder of the division of two integers
// Returns 0 if divisor is 0 and increments ErrorCounter
func ModulusInts(a, b int) int {
	if b == 0 {
		ErrorCounter++
		return 0
	}
	return a % b
}

// PowerInts returns a raised to the power of b
func PowerInts(a, b int) int {
	result := 1
	for i := 0; i < b; i++ {
		result *= a
	}
	return result
}

// PowerFloats returns a raised to the power of b
func PowerFloats(a, b float64) float64 {
	return math.Pow(a, b)
}

// AbsInt returns the absolute value of an integer
func AbsInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// AbsFloat returns the absolute value of a float64
func AbsFloat(a float64) float64 {
	return math.Abs(a)
}

// SqrtFloat returns the square root of a float64
// Returns 0 if input is negative and increments ErrorCounter
func SqrtFloat(a float64) float64 {
	if a < 0 {
		ErrorCounter++
		return 0
	}
	return math.Sqrt(a)
}
