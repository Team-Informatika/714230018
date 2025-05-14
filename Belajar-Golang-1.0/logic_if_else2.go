package main

import "fmt"

// ComplexIfElseExample demonstrates complex if-else logic with multiple conditions and nested ifs
func ComplexIfElseExample(nilai int, kehadiran int) string {
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

// ExampleComplex demonstrates usage of ComplexIfElseExample
func ExampleComplex() {
	var nilai, kehadiran int
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nilai)
	fmt.Print("Masukkan persentase kehadiran: ")
	fmt.Scan(&kehadiran)
	result := ComplexIfElseExample(nilai, kehadiran)
	fmt.Println(result)
}
