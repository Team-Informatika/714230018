package main

// Ordered is a constraint that permits any ordered type: int, float64, etc.
type Ordered interface {
	~int | ~float64
}

// Equal returns true if a equals b
func Equal[T Ordered](a, b T) bool {
	return a == b
}

// NotEqual returns true if a does not equal b
func NotEqual[T Ordered](a, b T) bool {
	return a != b
}

// GreaterThan returns true if a is greater than b
func GreaterThan[T Ordered](a, b T) bool {
	return a > b
}

// LessThan returns true if a is less than b
func LessThan[T Ordered](a, b T) bool {
	return a < b
}

// GreaterThanOrEqual returns true if a is greater than or equal to b
func GreaterThanOrEqual[T Ordered](a, b T) bool {
	return a >= b
}

// LessThanOrEqual returns true if a is less than or equal to b
func LessThanOrEqual[T Ordered](a, b T) bool {
	return a <= b
}
