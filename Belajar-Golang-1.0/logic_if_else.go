package main

import "fmt"

// LogicIfElseExample evaluates the grade based on nilai and returns the grade string
func LogicIfElseExample(nilai int) string {
	if nilai >= 80 {
		return "Nilai Anda A"
	} else if nilai >= 70 {
		return "Nilai Anda B"
	} else if nilai >= 60 {
		return "Nilai Anda C"
	} else if nilai >= 50 {
		return "Nilai Anda D"
	} else {
		return "Nilai Anda E"
	}
}

// Example usage function to demonstrate the logic
func Example() {
	var nilai int
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nilai)
	grade := LogicIfElseExample(nilai)
	fmt.Println(grade)
}
