package main

import "errors"

// Number is a constraint that permits any integer or float type.
type Number interface {
	~int | ~float64
}

// GenAdd returns the sum of two numbers.
func GenAdd[T Number](a, b T) T {
	return a + b
}

// GenSubtract returns the difference of two numbers.
func GenSubtract[T Number](a, b T) T {
	return a - b
}

// GenMultiply returns the product of two numbers.
func GenMultiply[T Number](a, b T) T {
	return a * b
}

// GenDivide returns the quotient of two numbers.
// Returns an error if divisor is zero.
func GenDivide[T Number](a, b T) (T, error) {
	var zero T
	if b == zero {
		return zero, errors.New("division by zero")
	}
	return a / b, nil
}

// Wrapper functions for backward compatibility

// AddInts returns the sum of two integers.
func AddInts(a, b int) int {
	return GenAdd[int](a, b)
}

// SubtractInts returns the difference of two integers.
func SubtractInts(a, b int) int {
	return GenSubtract[int](a, b)
}

// MultiplyInts returns the product of two integers.
func MultiplyInts(a, b int) int {
	return GenMultiply[int](a, b)
}

// DivideInts returns the quotient of two integers.
// Returns an error if divisor is zero.
func DivideInts(a, b int) (int, error) {
	return GenDivide[int](a, b)
}

// AddFloats returns the sum of two float64 numbers.
func AddFloats(a, b float64) float64 {
	return GenAdd[float64](a, b)
}

// SubtractFloats returns the difference of two float64 numbers.
func SubtractFloats(a, b float64) float64 {
	return GenSubtract[float64](a, b)
}

// MultiplyFloats returns the product of two float64 numbers.
func MultiplyFloats(a, b float64) float64 {
	return GenMultiply[float64](a, b)
}

// DivideFloats returns the quotient of two float64 numbers.
// Returns an error if divisor is zero.
func DivideFloats(a, b float64) (float64, error) {
	return GenDivide[float64](a, b)
}

func EqualInts(a, b int) bool {
	return a == b
}

func NotEqualInts(a, b int) bool {
	return a != b
}

func GreaterThanInts(a, b int) bool {
	return a > b
}

func LessThanInts(a, b int) bool {
	return a < b
}

func GreaterThanOrEqualInts(a, b int) bool {
	return a >= b
}

func LessThanOrEqualInts(a, b int) bool {
	return a <= b
}

func EqualFloats(a, b float64) bool {
	return a == b
}

func NotEqualFloats(a, b float64) bool {
	return a != b
}

func GreaterThanFloats(a, b float64) bool {
	return a > b
}

func LessThanFloats(a, b float64) bool {
	return a < b
}

func GreaterThanOrEqualFloats(a, b float64) bool {
	return a >= b
}

func LessThanOrEqualFloats(a, b float64) bool {
	return a <= b
}
