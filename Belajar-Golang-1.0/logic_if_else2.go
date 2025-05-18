package main

import "fmt"

// Grader interface defines a method to grade based on nilai and kehadiran
type Grader interface {
	Grade(nilai int, kehadiran int) string
}

// ComplexGrader implements Grader interface with complex if-else logic
type ComplexGrader struct{}

// Grade returns the grade string based on nilai and kehadiran
func (cg ComplexGrader) Grade(nilai int, kehadiran int) string {
	if nilai >= 80 {
		if kehadiran >= 90 {
			return "Sangat Baik: A+"
		} else if kehadiran >= 75 {
			return "Baik Sekali: A"
		} else {
			return "Baik: A-"
		}
	} else if nilai >= 70 {
		if kehadiran >= 80 {
			return "Baik: B+"
		} else {
			return "Cukup: B"
		}
	} else if nilai >= 60 {
		if kehadiran >= 70 {
			return "Cukup: C+"
		} else {
			return "Kurang: C"
		}
	} else {
		if kehadiran >= 60 {
			return "Buruk: D"
		} else {
			return "Gagal: E"
		}
	}
}

// ExampleComplex demonstrates usage of ComplexGrader
func ExampleComplex() {
	var nilai, kehadiran int
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nilai)
	fmt.Print("Masukkan persentase kehadiran: ")
	fmt.Scan(&kehadiran)
	var grader Grader = ComplexGrader{}
	result := grader.Grade(nilai, kehadiran)
	fmt.Println(result)
}
