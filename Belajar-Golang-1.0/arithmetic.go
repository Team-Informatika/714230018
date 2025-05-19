package main

import (
	"fmt"
	"log"
)

// Basic arithmetic functions for addition, subtraction, multiplication, division
func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func main() {
	// Demonstrate basic arithmetic
	a, b := 15.0, 4.0

	fmt.Printf("Addition: %.2f + %.2f = %.2f\n", a, b, Add(a, b))
	fmt.Printf("Subtraction: %.2f - %.2f = %.2f\n", a, b, Subtract(a, b))
	fmt.Printf("Multiplication: %.2f * %.2f = %.2f\n", a, b, Multiply(a, b))

	divResult, err := Divide(a, b)
	if err != nil {
		log.Println("Error in division:", err)
	} else {
		fmt.Printf("Division: %.2f / %.2f = %.2f\n", a, b, divResult)
	}

	// Demonstrate modulus using existing function
	modResult, err := ModulusInts(int(a), int(b))
	if err != nil {
		log.Println("Error in modulus:", err)
	} else {
		fmt.Printf("Modulus: %d %% %d = %d\n", int(a), int(b), modResult)
	}

	// Demonstrate power functions
	fmt.Printf("PowerInts: %d ^ %d = %d\n", int(a), int(b), PowerInts(int(a), int(b)))
	fmt.Printf("PowerFloats: %.2f ^ %.2f = %.2f\n", a, b, PowerFloats(a, b))

	// Demonstrate absolute value functions
	fmt.Printf("AbsInt: abs(%d) = %d\n", -int(a), AbsInt(-int(a)))
	fmt.Printf("AbsFloat: abs(%.2f) = %.2f\n", -a, AbsFloat(-a))

	// Demonstrate square root function
	sqrtResult, err := SqrtFloat(a)
	if err != nil {
		log.Println("Error in square root:", err)
	} else {
		fmt.Printf("Square root: sqrt(%.2f) = %.2f\n", a, sqrtResult)
	}

	// Demonstrate error case for square root of negative number
	_, err = SqrtFloat(-a)
	if err != nil {
		fmt.Println("Expected error for sqrt of negative number:", err)
	}
}
