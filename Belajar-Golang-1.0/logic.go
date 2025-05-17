// Additional varied logic functions for learning coding

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

// Majority returns true if at least two of the three inputs are true
func Majority(a, b, c bool) bool {
	return (a && b) || (a && c) || (b && c)
}

// Implies returns the logical implication (if a then b)
func Implies(a, b bool) bool {
	return !a || b
}

// Parity returns true if an odd number of inputs are true (odd parity)
func Parity(a, b, c bool) bool {
	return (a != b) != c
}

// Conditional returns b if a is true, otherwise c (if-then-else logic)
func Conditional(a, b, c bool) bool {
	if a {
		return b
	}
	return c
}

// ComplexLogic demonstrates a custom logic combining multiple operations
func ComplexLogic(a, b, c bool) bool {
	return (a && !b) || (b && c) && !(a && c)
}
