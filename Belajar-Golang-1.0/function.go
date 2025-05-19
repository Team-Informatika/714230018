package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"sync"
	"time"
)

// DivideWithRemainder returns the quotient and remainder of two integers
func DivideWithRemainder(a int, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, errors.New("division by zero")
	}
	return a / b, a % b, nil
}

// Greet prints a greeting message
func Greet() {
	fmt.Println("Hello, welcome to the function.go example!")
}

// Swap swaps two strings and returns them
func Swap(x, y string) (string, string) {
	return y, x
}

// Min returns the minimum of two integers
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// IsEven returns true if the number is even
func IsEven(n int) bool {
	return n%2 == 0
}

// IsOdd returns true if the number is odd
func IsOdd(n int) bool {
	return n%2 != 0
}

// Fibonacci returns the nth Fibonacci number
func Fibonacci(n int) (int, error) {
	if n < 0 {
		return -1, errors.New("negative input")
	}
	if n == 0 {
		return 0, nil
	}
	if n == 1 {
		return 1, nil
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b, nil
}

// ReverseString returns the reversed string of input s
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// CountWords counts the number of words in a string
func CountWords(s string) int {
	words := strings.Fields(s)
	return len(words)
}

// Contains checks if a slice contains a given integer
func Contains(slice []int, val int) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

// SumSlice returns the sum of all integers in a slice
func SumSlice(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

// FilterEven returns a slice containing only even numbers from the input slice
func FilterEven(slice []int) []int {
	var evens []int
	for _, v := range slice {
		if IsEven(v) {
			evens = append(evens, v)
		}
	}
	return evens
}

// MapStringLengths returns a map of string to its length for a slice of strings
func MapStringLengths(slice []string) map[string]int {
	m := make(map[string]int)
	for _, s := range slice {
		m[s] = len(s)
	}
	return m
}

// Struct example: Person
type Person struct {
	Name string
	Age  int
}

// NewPerson creates a new Person instance
func NewPerson(name string, age int) *Person {
	return &Person{Name: name, Age: age}
}

// Birthday increments the age of the person by 1
func (p *Person) Birthday() {
	p.Age++
}

// String returns a string representation of the person
func (p *Person) String() string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

// Interface example: Shape
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle struct implements Shape
type Rectangle struct {
	Width, Height float64
}

// Area returns the area of the rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter returns the perimeter of the rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle struct implements Shape
type Circle struct {
	Radius float64
}

// Area returns the area of the circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter returns the perimeter (circumference) of the circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// PrintShapeInfo prints area and perimeter of a shape
func PrintShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

// Concurrency example: Print numbers with goroutines
func PrintNumbers() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println("Number:", n)
		}(i)
	}
	wg.Wait()
}

// Channel example: Sum numbers sent on a channel
func SumChannel(nums []int) int {
	ch := make(chan int)
	go func() {
		sum := 0
		for _, n := range nums {
			sum += n
		}
		ch <- sum
	}()
	return <-ch
}

// Error handling example: Safe division
func SafeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Recursive function: Power calculates a^b
func Power(a, b int) int {
	if b == 0 {
		return 1
	}
	return a * Power(a, b-1)
}

// Variadic function: Sum all integers
func Sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Closure example: Adder returns a function that adds a fixed number
func Adder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

// Anonymous function example: Filter slice with a predicate
func Filter(slice []int, predicate func(int) bool) []int {
	var result []int
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Timer example: Measure execution time of a function
func MeasureExecutionTime(f func()) time.Duration {
	start := time.Now()
	f()
	return time.Since(start)
}

// Max returns the maximum of two integers
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Example usage demonstrating many functions
func ExampleUsage() {
	fmt.Println("Add(3, 4):", Add(3.0, 4.0))
	fmt.Println("Subtract(10, 5):", Subtract(10.0, 5.0))
	fmt.Println("Multiply(2, 3):", Multiply(2.0, 3.0))
	q, r, err := DivideWithRemainder(10, 3)
	if err != nil {
		fmt.Println("Divide error:", err)
	} else {
		fmt.Printf("Divide(10, 3): quotient=%d remainder=%d\n", q, r)
	}
	Greet()
	a, b := Swap("first", "second")
	fmt.Println("Swap(\"first\", \"second\"):", a, b)
	fmt.Println("Max(10, 20):", Max(10, 20))
	fmt.Println("Min(10, 20):", Min(10, 20))
	fmt.Println("IsEven(4):", IsEven(4))
	fmt.Println("IsOdd(5):", IsOdd(5))
	fib, err := Fibonacci(7)
	if err != nil {
		fmt.Println("Fibonacci error:", err)
	} else {
		fmt.Println("Fibonacci(7):", fib)
	}
	fmt.Println("ReverseString(\"hello\"):", ReverseString("hello"))
	fmt.Println("CountWords(\"hello world\"):", CountWords("hello world"))
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Contains(slice, 3):", Contains(slice, 3))
	fmt.Println("SumSlice(slice):", SumSlice(slice))
	fmt.Println("FilterEven(slice):", FilterEven(slice))
	strSlice := []string{"apple", "banana", "cherry"}
	fmt.Println("MapStringLengths(strSlice):", MapStringLengths(strSlice))
	p := NewPerson("Alice", 30)
	fmt.Println(p)
	p.Birthday()
	fmt.Println("After birthday:", p)
	var rect Shape = Rectangle{Width: 3, Height: 4}
	var c Shape = Circle{Radius: 5}
	PrintShapeInfo(rect)
	PrintShapeInfo(c)
	PrintNumbers()
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("SumChannel(nums):", SumChannel(nums))
	divResult, divErr := SafeDivide(10, 0)
	if divErr != nil {
		fmt.Println("SafeDivide error:", divErr)
	} else {
		fmt.Println("SafeDivide result:", divResult)
	}
	fmt.Println("Power(2, 3):", Power(2, 3))
	fmt.Println("Sum(1, 2, 3, 4):", Sum(1, 2, 3, 4))
	adder := Adder(5)
	fmt.Println("Adder(5)(10):", adder(10))
	filtered := Filter(slice, func(n int) bool { return n > 2 })
	fmt.Println("Filter(slice, n > 2):", filtered)
	duration := MeasureExecutionTime(func() {
		time.Sleep(100 * time.Millisecond)
	})
	fmt.Println("MeasureExecutionTime:", duration)
}
