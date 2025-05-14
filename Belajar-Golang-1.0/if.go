package main

import "fmt"

// BasicIfExample demonstrates simple if logic
func BasicIfExample(nilai int) string {
	if nilai >= 75 {
		return "Lulus"
	}
	return "Tidak Lulus"
}

// ExampleBasic demonstrates usage of BasicIfExample
func ExampleBasic() {
	var nilai int
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nilai)
	result := BasicIfExample(nilai)
	fmt.Println(result)
}
