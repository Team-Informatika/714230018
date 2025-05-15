package main

// SumRange returns the sum of all integers from start to end inclusive
func SumRange(start, end int) int {
	sum := 0
	for i := start; i <= end; i++ {
		sum += i
	}
	return sum
}

// Factorial calculates the factorial of a non-negative integer n
func Factorial(n int) int {
	if n < 0 {
		return -1 // error case, factorial not defined for negative numbers
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// CountVowels counts the number of vowels in a given string
func CountVowels(s string) int {
	count := 0
	for _, ch := range s {
		switch ch {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			count++
		}
	}
	return count
}
