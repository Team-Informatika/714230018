package main

import "fmt"

// Add returns the sum of two integers
func Add(a int, b int) int {
	return a + b
}

// Greet prints a greeting message
func Greet() {
	fmt.Println("Hello, welcome to the function.go example!")
}

// Swap swaps two strings and returns them
func Swap(x, y string) (string, string) {
	return y, x
}

// ExampleUsage demonstrates the usage of the functions in this file
func ExampleUsage() {
	fmt.Println("Add(3, 4):", Add(3, 4))
	Greet()
	a, b := Swap("first", "second")
	fmt.Println("Swap(\"first\", \"second\"):", a, b)
}
