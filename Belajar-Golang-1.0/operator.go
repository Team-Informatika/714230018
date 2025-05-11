package main

// AddInts returns the sum of two integers
func AddInts(a, b int) int {
	return a + b
}

// SubtractInts returns the difference of two integers
func SubtractInts(a, b int) int {
	return a - b
}

// MultiplyInts returns the product of two integers
func MultiplyInts(a, b int) int {
	return a * b
}

// DivideInts returns the quotient of two integers
// Returns 0 if divisor is 0
func DivideInts(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

// AddFloats returns the sum of two float64 numbers
func AddFloats(a, b float64) float64 {
	return a + b
}

// SubtractFloats returns the difference of two float64 numbers
func SubtractFloats(a, b float64) float64 {
	return a - b
}

// MultiplyFloats returns the product of two float64 numbers
func MultiplyFloats(a, b float64) float64 {
	return a * b
}

// DivideFloats returns the quotient of two float64 numbers
// Returns 0 if divisor is 0
func DivideFloats(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}
