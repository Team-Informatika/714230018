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

// Xor returns the logical exclusive OR of two boolean values
func Xor(a, b bool) bool {
	return (a || b) && !(a && b)
}

// Nand returns the logical NAND of two boolean values
func Nand(a, b bool) bool {
	return !(a && b)
}

// Nor returns the logical NOR of two boolean values
func Nor(a, b bool) bool {
	return !(a || b)
}

// Xnor returns the logical equivalence (XNOR) of two boolean values
func Xnor(a, b bool) bool {
	return a == b
}
