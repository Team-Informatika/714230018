package main

// And returns the logical AND of two boolean values
func And(a, b bool) bool {
	return a && b
}

// Or returns the logical OR of two boolean values
func Or(a, b bool) bool {
	return a || b
}

// Not returns the logical NOT of a boolean value
func Not(a bool) bool {
	return !a
}
