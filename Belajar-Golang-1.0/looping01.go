package main

import "fmt"

// MaxValue returns the maximum value in a slice of integers
func MaxValue(nums []int) int {
	if len(nums) == 0 {
		return 0 // or error value
	}

	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

// PrintElements prints all elements of a slice of strings
func PrintElements(elements []string) {
	for i := 0; i < len(elements); i++ {
		fmt.Println(elements[i])
	}
}

// SumEvenNumbers sums all even numbers from 0 up to n (exclusive)
func SumEvenNumbers(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	return sum
}

// NestedLoopsExample demonstrates nested loops by printing a multiplication table
func NestedLoopsExample(size int) {
	for i := 1; i <= size; i++ {
		for j := 1; j <= size; j++ {
			fmt.Printf("%d\t", i*j)
		}
		fmt.Println()
	}
}

// WhileLikeLoop demonstrates a while-like loop using for
func WhileLikeLoop(limit int) {
	i := 0
	for i < limit {
		fmt.Println(i)
		i++
	}
}
