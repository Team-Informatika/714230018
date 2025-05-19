// Alternative logical functions for learning coding

package main

// Nor returns the logical NOR of two boolean values
func Nor(a, b bool) bool {
	return !(a || b)
}

// Nand returns the logical NAND of two boolean values
func Nand(a, b bool) bool {
	return !(a && b)
}

// Xor returns the logical exclusive OR of two boolean values
func Xor(a, b bool) bool {
	return a != b
}

// Xnor returns the logical equivalence (XNOR) of two boolean values
func Xnor(a, b bool) bool {
	return a == b
}

// And3 returns the logical AND of three boolean values
func And3(a, b, c bool) bool {
	return a && b && c
}

// Or3 returns the logical OR of three boolean values
func Or3(a, b, c bool) bool {
	return a || b || c
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
